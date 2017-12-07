//数据库连接池测试
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
	"time"
)

/*
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) DEFAULT '',
  `age` int(11) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4
*/
var db *sql.DB

func init() {
	db, _ = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.Ping()
}

func main() {
	startHttpServer()
}

func startHttpServer() {
	http.HandleFunc("/pool", pool)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func insert() {

	//方式1 insert
	//strconv,int转string:strconv.Itoa(i)
	start := time.Now()
	for i := 1001; i <= 1100; i++ {
		//每次循环内部都会去连接池获取一个新的连接，效率低下
		db.Exec("INSERT INTO user(uid,username,age) values(?,?,?)", i, "user"+strconv.Itoa(i), i-1000)
	}
	end := time.Now()
	fmt.Println("方式1 insert total time:", end.Sub(start).Seconds())

	//方式2 insert
	start = time.Now()
	for i := 1101; i <= 1200; i++ {
		//Prepare函数每次循环内部都会去连接池获取一个新的连接，效率低下
		stm, _ := db.Prepare("INSERT INTO user(uid,username,age) values(?,?,?)")
		stm.Exec(i, "user"+strconv.Itoa(i), i-1000)
		stm.Close()
	}
	end = time.Now()
	fmt.Println("方式2 insert total time:", end.Sub(start).Seconds())

	//方式3 insert
	start = time.Now()
	stm, _ := db.Prepare("INSERT INTO user(uid,username,age) values(?,?,?)")
	for i := 1201; i <= 1300; i++ {
		//Exec内部并没有去获取连接，为什么效率还是低呢？
		stm.Exec(i, "user"+strconv.Itoa(i), i-1000)
	}
	stm.Close()
	end = time.Now()
	fmt.Println("方式3 insert total time:", end.Sub(start).Seconds())

	//方式4 insert
	start = time.Now()
	//Begin函数内部会去获取连接
	tx, _ := db.Begin()
	for i := 1301; i <= 1400; i++ {
		//每次循环用的都是tx内部的连接，没有新建连接，效率高
		tx.Exec("INSERT INTO user(uid,username,age) values(?,?,?)", i, "user"+strconv.Itoa(i), i-1000)
	}
	//最后释放tx内部的连接
	tx.Commit()

	end = time.Now()
	fmt.Println("方式4 insert total time:", end.Sub(start).Seconds())

	//方式5 insert
	start = time.Now()
	for i := 1401; i <= 1500; i++ {
		//Begin函数每次循环内部都会去连接池获取一个新的连接，效率低下
		tx, _ := db.Begin()
		tx.Exec("INSERT INTO user(uid,username,age) values(?,?,?)", i, "user"+strconv.Itoa(i), i-1000)
		//Commit执行后连接也释放了
		tx.Commit()
	}
	end = time.Now()
	fmt.Println("方式5 insert total time:", end.Sub(start).Seconds())
}

func pool(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM user limit 1")
	defer rows.Close()
	checkErr(err)

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	record := make(map[string]string)
	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
	}

	fmt.Println(record)
	fmt.Fprintln(w, "finish")
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
