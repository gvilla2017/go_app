package main

import (
	"fmt"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	)

var item_id int
var description string


func mysql(w http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("mysql",
		"root:new_password@tcp(mysql-service:3306)/jewelry")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("select item_id, description from earings")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&item_id, &description)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(item_id, description)
		fmt.Fprintf(w, description)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {
	http.HandleFunc("/hello", hello)
    	http.HandleFunc("/headers", headers)
    	http.HandleFunc("/mysql", mysql)

	http.ListenAndServe(":9200", nil)

}

