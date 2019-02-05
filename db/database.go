package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)


type Person struct {
	Name string
	Age int
}

func main() {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()
	cmd := `CREATE TABLE IF NOT EXISTS person(
            name STRING,
            age INT)`
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	//INSERT文
	cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
	_, err = DbConnection.Exec(cmd, "Nancy", 20)
	if err != nil {
		log.Fatalln(err)
	}

	//UPDATE文
	cmd = "UPDATE person SET age = ? WHERE name = ?"
	_, err = DbConnection.Exec(cmd, 30, "Nancy")

	//multiple select
	cmd = "SELECT * FROM person"
	rows, _ := DbConnection.Query(cmd)
	defer rows.Close()
	//構造体Personを複数格納するppを宣言
	var pp []Person
	for rows.Next() {
		var p Person
		//１行の値を取得する
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Println(err)
		}
		//取得した行の値をスライスppに追加する
		pp = append(pp, p)
	}

	//スライスに格納された値を行ごとに出力する
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}

	//single select
	cmd = "SELECT * FROM person WHERE age = ?"
	row := DbConnection.QueryRow(cmd, 30)
	var p Person
	err = row.Scan(&p.Name, &p.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No Row")
		} else {
			fmt.Println(err)
		}
	}
	fmt.Println(p.Name, p.Age)

	//DELETE文
	cmd = "DELETE FROM person WHERE name = ?"
	_, err = DbConnection.Exec(cmd, "Nancy")
	if err != nil {
		log.Fatalln(err)
	}

	//?はSQLインジェクション対策に使う
	//ただし、?はVALUEに対して使えるが、tablenameには　使えない
	tablename := "person"
	cmd = fmt.Sprintf("SELECT * FROM %s", tablename)
	rows, _ = DbConnection.Query(cmd)
	defer rows.Close()
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p)
	}
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}
}
