package main

import (
	//"database/sql"
	//"fmt"
	//_ "github.com/garyburd/redigo/redis"
	//_ "github.com/go-sql-driver/mysql"
	//_ "github.com/nats-io/go-nats"
	//"strings"
	"database/sql"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/nats-io/go-nats"
	"strings"
)

//数据库配置
const (
	userName = "root"
	password = "123456"
	ip       = "192.168.1.124"
	port     = "3306"
	dbName   = "docker"
)

func testMysqlConn() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	var DB *sql.DB
	DB, _ = sql.Open("mysql", path)
	fmt.Printf("%T", DB)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("opon database fail")
		return
	}
	fmt.Println("connnect success")
	sql := "insert into student (id,name) values (2,\"zsd\")"
	result, err := DB.Exec(sql)
	if err != nil {
		fmt.Println("sql执行出错！")
	}
	DB.Close()
	fmt.Println("%T", result)
}

func testRedisConn() {

	conn, err := redis.Dial("tcp", "192.168.1.124:6379", redis.DialPassword(password))
	if err != nil {
		fmt.Println("redis connection fail!", err)
	}
	fmt.Println("redis connection success")
	_, err = conn.Do("set", "test1", "redis1")
	//do, err := conn.Do("set", "test1", "redis1")
	if err != nil {
		fmt.Println("redis insert fail!", err)
	}
	test1, err := redis.String(conn.Do("get", "test1"))
	if err != nil {
		fmt.Println("redis get fail!", err)
	}
	fmt.Println(test1)
	defer conn.Close()
}

func testNatsConn() {
	url := "nats://192.168.3.125:4222"
	nc, err := nats.Connect(url)
	if err != nil {
		fmt.Println("nats connection error,", err)
	}

	nc.Publish("test", []byte("hello nats!"))

}

func main() {
	//testMysqlConn()
	//testRedisConn()
	testNatsConn()
}
