package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func initDb() error {
	var err error
	dns := "root:123456@tcp(172.20.35.49:13306)/yqfk"
	DB, err = sql.Open("mysql", dns)
	if err != nil {
		return err
	}
	DB.SetMaxOpenConns(100)
	DB.SetConnMaxIdleTime(16)
	return nil
}

type User struct {
	Id   int64
	Name sql.NullString `db:"string"`
	Age  int            `db:"age"`
}

func testQueryData() {
	sqlstr := "select id ,name,age from user where id >?"
	rows, _ := DB.Query(sqlstr, 0)
	// row := DB.QueryRow(sqlstr, 1)
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	fmt.Println(rows)
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%v age:%d\n", user.Id, user.Name, user.Age)
	}

}

func testInsertData() {
	sqlstr := "insert into user(name,age) values(?,?)"
	result, err := DB.Exec(sqlstr, "tom", 18)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("get last insert id dailed ,err:%v\n", err)
		return
	}
	fmt.Printf("id is %d\n", id)
}

func testupdateData() {
	sqlstr := "update user set name =? where id =?"
	result, err := DB.Exec(sqlstr, "jim", 3)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	fmt.Printf("update db succ ,affected rows:%d\n", affected)

}

func testDeleteData() {
	sqlstr := "delete from user where id =?"
	result, err := DB.Exec(sqlstr, 3)
	if err != nil {
		fmt.Printf("dalete failed,err:%v\n", err)
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return
	}
	fmt.Printf("delete db succ ,affected rows:%d\n", affected)

}

func testPreparData() {
	sqlstr := "select id,name,age from user where id >?"
	stmt, err := DB.Prepare(sqlstr)
	if err != nil {
		fmt.Printf("prepare failed ,err:%v\n", err)
		return
	}
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	rows, err := stmt.Query(0)
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		return
	}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("scan failed ,err:%v\n", err)
			return
		}
		fmt.Printf("user:%#v\n", user)
	}
}

func main() {
	err := initDb()
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	// testQueryData()
	// testInsertData()
	// testupdateData()
	// testDeleteData()
	testPreparData()
}
