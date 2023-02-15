package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	//连接数据库
	connStr := "root:@tcp(127.0.0.1:3306)/ginsql"

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalln("mysql conn failed:", err.Error())
		return
	}

	defer db.Close()

	//创建表
	sql_pre := "DROP TABLE IF EXISTS person;"
	sql := `
CREATE TABLE person(
id INT AUTO_INCREMENT PRIMARY KEY,
name VARCHAR(16) NOT NULL,
age INT DEFAULT 1
);
`
	_, err = db.Exec(sql_pre)
	if err != nil {
		log.Fatalln("mysql exec failed:", err.Error())
		return
	} else {
		log.Println("数据库去重成功")
	}

	res, err := db.Exec(sql)
	if err != nil {
		log.Fatalln("mysql exec failed:", err.Error())
		return
	} else {
		a, _ := res.LastInsertId()
		b, _ := res.RowsAffected()
		log.Println("数据库建表成功", a, b)
	}
	//mysql 内部输入 desc person;可查看表结构

	//插入数据到数据库表中
	sql = "INSERT INTO person (name,age) VALUE (?,?);"
	res, err = db.Exec(sql, "Adam", 25)
	_, err = db.Exec(sql, "David", 16)
	_, err = db.Exec(sql, "Tom", 18)
	if err != nil {
		log.Fatalln("mysql exec failed:", err.Error())
		return
	} else {
		a, _ := res.LastInsertId()
		b, _ := res.RowsAffected()
		log.Println("插入数据成功", a, b)
	}

	//查询数据表
	sql = "SELECT id,name,age FROM person;"
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatalln("mysql exec failed:", err.Error())
		return
	} else {
		a, _ := res.LastInsertId()
		b, _ := res.RowsAffected()
		log.Println("查询数据成功", a, b)
	}

	for {
		if rows.Next() {
			person := new(Person)
			err = rows.Scan(&person.id, &person.name, &person.age)
			if err != nil {
				log.Fatalln("rows scan failed:", err.Error())
				return
			} else {
				log.Println(person)
			}
		} else {
			break
		}
	}

}

type Person struct {
	id   int
	name string
	age  int
}
