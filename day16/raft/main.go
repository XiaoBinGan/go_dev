package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

//1.实现3节点选举
//2.改造代码成分布式选举代码,加入Rpc调用
//3.演示完整代码 自动选主  日志复制


//定义3节点常量
const raftCount =3

/**
	Leader对象
	Term int			 任期 还没上任为0
	LeaderId int         leaderId 编号 没有编号为-1
 */
type Leader struct {
	Term int
	LeaderId int
}

/**raft的声明

锁											mu					sync.Mutex
节点编号										me					int
当前任期										currentTerm			int
为哪个节点投票	-1代表谁都不投					vatedFor			int
0.follower 1.candidate 2.leader				state				int
发送最后一条消息的时间							lastMessaheTime		int64
当前节点的leader								currentLeader		int
节点间发送信息的通道							message				chan bool
选举通道										electCh				chan bool
心跳信号的通道									heartBeat			chan bool
返回心跳信号的通道								heartbeatRe			chan bool
超时时间										timeout   			int
 */
type Raft struct{
	mu  			sync.Mutex
	me   			int
	currentTerm 	int
	votedFor 		int
	state  			int
	lastMessaheTime int64
	currentLeader   int
	message 		chan bool
	electCh			chan bool
	heartBeat		chan bool
	heartbeatRe		chan bool
	timeout   		int
}

var leader = Leader{0,-1}

func main() {
	//过程:有三个节点,最初都是follower
	//如果有candidate状态,进行投票和拉票
	//会产生leader
	//会产生3个节点
	for i:=0;i<raftCount;i++{
		//穿件三个raft节点
		Make(i)
	}

}

func Make(me int ) *Raft  {
	rf:=&Raft{}
	rf.me=me
	rf.votedFor=-1
	rf.state=0
	rf.timeout=0
	rf.currentLeader=-1
	rf.setTerm(0)
	rf.message=make(chan bool)
	rf.electCh=make(chan bool)
	rf.heartBeat=make(chan bool)
	rf.heartbeatRe=make(chan bool)
	//设置随机种子
	rand.Seed(time.Now().UnixNano())
	//选举协程
     go rf.election()
	//心跳检测协程
	go rf.sendLeaderHeartBeat()
	return rf
}
/** setTerm for currentTerm
	for Raft struct
	params term int which one item number
 */
func (rf *Raft)setTerm(term int ){
	rf.currentTerm=term
}
/** election for leader
 */
func (rf *Raft)election()  {
	var result bool
	for  {
		timeout :=randRage(150,300)
		rf.lastMessaheTime=millisecond()
		select {
		case <-time.After(time.Duration(timeout)*time.Microsecond):
			fmt.Println("当前节点状态为:",rf.state)
		}
		result=false
		for !result{
			result=rf.election_one_round(&leader)
		}
	}
}

/* randRange set time out
 */
func randRage(min,max int64) int64  {
	return  rand.Int63n(max-min)+min
}
/**get the lastMessage send time
 */
func millisecond() int64 {
	return time.Now().UnixNano()/int64(time.Microsecond)
}
/**
implementation of select leader
 */
func (rf *Raft)election_one_round(leader *Leader) bool  {
	//set time out
	var timeout int64
	timeout = 100
	//number of vote
	var vote int
	//heartbeat
	var triggerHeartbeat bool
	var success bool
	//time
	last :=millisecond()
	//change the item be canditate
	rf.mu.Lock()
	//change state
	rf.becomeCandidate()
	rf.mu.Unlock()
	fmt.Println("staet electing leader")
	//range all item
	for{
		for i:=0;i<raftCount;i++ {
			if i!=rf.me{
				// canvass vote
				go func() {
					if leader.LeaderId<0{
							rf.electCh<-true
					}
				}()
			}
		}
		vote  = 1
		for i:=0;i<raftCount;i++{
			//calculate the number of vote
			select {
			case ok:=<-rf.electCh:
				if ok{
					//vote number add 1
					vote++
					//if vote great than node one-second ,successful
					success=vote>raftCount/2
					if success&&!triggerHeartbeat{
						//change be the leader success
						//trigger the hearbeat detection
						triggerHeartbeat =true
						rf.mu.Lock()
						rf.becomeLeader()
						rf.mu.Unlock()
						rf.heartBeat<- true
						fmt.Println(rf.me,"号节点成为leader")
						fmt.Println("leader 开始发送心跳信号")
					}
				}
			}
		}
		//final check
		//if don't time out and the vote great then node one-second ,than impelmention select leader successful.
		if timeout+last<millisecond()||(vote>raftCount/2||rf.currentLeader>-1){
			break
		}else{
			select {
			case <-time.After(time.Duration(10)*time.Microsecond):
			}
		}

	}
	return success

}

/**
change state with candidate
 */
func (rf *Raft)becomeCandidate()  {
	rf.state=1
	rf.setTerm(rf.currentTerm+1)
	rf.votedFor=rf.me
	rf.currentLeader=-1
}

/**
change state with leader
 */
func(rf *Raft)becomeLeader(){
 rf.state=2
 rf.currentLeader=rf.me
}

/**
  leader send heartBeat
  complete data synchronization
 */
func (rf *Raft) sendLeaderHeartBeat()  {
	//infinity loop
	for{
		select {
		case <-rf.heartBeat:
			rf.sendAppendEntriesImpl()
		}
	}
}
/**
return info to the leader
 */
func (rf *Raft)sendAppendEntriesImpl()  {
	if rf.currentLeader == rf.me{
		//the time  is  leader
		//number of record to confirm single node
		var success_count =0
		for i:=0;i<raftCount;i++{
			if i!=rf.me{
				go func() {
					rf.heartbeatRe<-true
				}()
			}
			//complete the sigle node of return confirm
			for i:=0;i<raftCount;i++{
				select {
				case ok:=<-rf.heartbeatRe:
					if ok{
						success_count++
						if success_count>raftCount/2{
							fmt.Println("投票选举成功,心跳信号Ok")
							log.Fatal("程序结束")
						}
					}

				}
			}
		}
	}
}