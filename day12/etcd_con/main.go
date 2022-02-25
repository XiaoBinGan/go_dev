package main
 
 import (
     "fmt"
    //  etcd_client "github.com/coreos/etcd/clientv3"上面这个无效因为下载不下来 
     etcd_client "go.etcd.io/etcd/clientv3"
     "time"
 )
 /*
                //etcd_client.News生成一个etcd的客户端
 
 */
 func main() {
     cli, err := etcd_client.New(etcd_client.Config{//下面这些时配置内容 
         Endpoints:   []string{"192.168.30.136:2379", "192.168.30.136:22379", "192.168.30.136:32379"},//etcd的端口因为etcd是个集群所以会存在多个不光可以写ip+端口还可以写域名(域名基本不会变而IP可能会变)
         DialTimeout: 5 * time.Second,//连接的超时时间
     })
     if err != nil {
         fmt.Println("connect failed, err:", err)
         return
     }
 
     fmt.Println("connect succ")
     defer cli.Close()//用完记得关闭不然会内存溢出
 }