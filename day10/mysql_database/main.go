package main

import(
	"database/sql"
	"fmt"
	//init初始化的时候像go的基础库添加了识别打开myslq的方法
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB     //全局申明一个sql.DB的对象来方便外部操作
type user struct { //用来接收查询出来的数据
	id   int
	name string
	age  int
}

/*go使用"database/sql" 内置库连接mysql
  1.用户名:密码@tcp(IP:端口)/databaseName
  2.使用sql.Open("数据库类型",sqldns)
	2.1这里打开成功和失败是不会被校验的只要是按照内置库的格式实现就不会error
	2.2如果没有按照格式来写的话就会被报错
  3.db.Ping()的方式尝试连接数据库 err=nil则连接成功
  4.设置最大连接数(连接池的设置)
  5.设置最大空闲连接数 允许最大空闲的连接数量
*/
func intiDB() (err error) {
	//	   用户名:密码@tcp(IP:端口)/databaseName
	dsn := "root:@tcp(127.0.0.1:3306)/sql_test"
	db, err = sql.Open("mysql", dsn) //这里打开成功和失败是不会被校验的只要是按照内置库的格式实现就不会error
	if err != nil {                  //上面的dsbr如果没有按照格式来写的话就会被报错
		fmt.Println("open mysql errr:", err)
		return err
	}
	err = db.Ping() //尝试连接数据库
	if err != nil {
		fmt.Println("Ping failed err:", err)
		return err
	}
	db.SetMaxOpenConns(10) //设置最大连接数
	db.SetMaxIdleConns(2)  //设置最大闲置连接数
	fmt.Println("MySql 连接成功")
	return nil
}

/*queryOne更具id查询单条记录
  1.写好一段条件查询语句更具ID去查询单条
  2.使用db.QueryRow().Scan()非常重要：
    确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
  3.打印获取处理啊的记录
*/
func queryOne(id int) {
	//		  选择  ID,name,age 从user表中 查询条件 id=? (?为占位符)
	sqlstr := `select id,name,age from user where id=?;`
	var u user
	err := db.QueryRow(sqlstr, id).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("sacn failed err:%v\n", err)
		return
	}
	fmt.Printf("id:%d,name:%s,age:%d\n", u.id, u.name, u.age)
}

/*多行查询
多行查询db.Query()执行一次查询，返回多行结果（即Rows），
一般用于执行select命令。参数args表示query中的占位参数。
1.编写SQL语句
2.使用全局的db对象去执行SQ语句条件
3.执行成功获取到rows对象
// Rows is the result of a query. Its cursor starts before the first row
// of the result set. Use Next to advance from row to row.
4.rows.next去循环每条记录循环到结束的时候返回bool值错误.
5.使用ros.Scan()将每个字段写入对应结构体的字段中
6.for循环逐行答应
*/
func queryMultiRowDemo(id int) {
	sqlStr := "select id,name,age from user where id >?"
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed,err%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next(){
		var u user
		err :=rows.Scan(&u.id,&u.name,&u.age)
		if err!=nil{
			fmt.Printf("scan failed,err:%v\n", err)
			return
		}
		fmt.Printf("%+v\n",u)
	}
}

/*单条记录插入
  1.创建sql语句
  2 使用全局的db对象db.Exec(sqlStr,args...)返回(sql.Result, error)
  3.容错
  4.使用result.LastInsertId()//返回新嘻哈如数据的id
  5.容错
  6.打印结果
*/

func insertRow(name string,age int)  {
	sqlStr :="insert into user(name,age)values(?,?)"
	ret,err :=db.Exec(sqlStr,name,age)
	if err!=nil {
		fmt.Println("EXec failed err:",err)
		return
	}
	id,err :=ret.LastInsertId()
	if err!=nil{
		fmt.Printf("inert into failed, err:%v\n",err)
		return
	}
	fmt.Printf("inser into user success id :%d\n", id)
}


/*单条记录修改
  1.创建sql语句
  2 使用全局的db对象db.Exec(sqlStr,args...)返回(sql.Result, error)
  3.容错
  4.使用result.RowsAffected()//返回操作影响的行数
  5.容错
  6.打印结果
*/
func updateRow(age,id int)  {
	sqlStr :="update user set age=? where id=?"
	ret,err :=db.Exec(sqlStr,age,id)
	if err!=nil{
		fmt.Printf("Exec failed err:%v\n",err)
		return
	}
	n,err :=ret.RowsAffected()
	if err!=nil{
		fmt.Printf("upddate failed err:%v\n", err)
		return
	}
	fmt.Printf("update user success affected rowid %v\n",n)
}
/*单条记录删除(线上数据一般不存在真正的数据库删除)
  1.创建sql语句
  2 使用全局的db对象db.Exec(sqlStr,args...)返回(sql.Result, error)
  3.容错
  4.使用result.RowsAffected()//返回操作影响的行数
  5.容错
  6.打印结果
*/
func delRow(id int)  {
	sqlStr :="delete from user where id=?"
	rest,err :=db.Exec(sqlStr,id)
	if err!=nil{
		fmt.Printf("Exec failed err:%v\n",err)
		return
	}
	n,err :=rest.RowsAffected()
	if err!=nil{
		fmt.Printf("upddate failed err:%v\n", err)
		return
	}
	fmt.Printf("update user success affected rowid %v\n",n)
}

func main() {
	err := intiDB()    		//初始化连接
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	// queryOne(2)			//查询单条记录
	// queryMultiRowDemo(0)	//查询多条记录
	// insertRow("del",18)	//插入单条记录
	// updateRow(900,3)		//修改单条记录
	delRow(4)				//删除单条记录
}
