// *GLDS HANA DB Connect
//* Copyright Rushi NS

package main

// * import SAP Driver and standard packages

import (
	//*    "context"
        "fmt"
        "log"
        "time"
        "database/sql"
	_ "SAP/go-hdb/driver"
)

//* main logic
//* connection details

func main() {
	// define a connection string
	var connStr string
	connStr = "hdb://system:Abcd1234@10.48.144.154:30215"
	db, err := sql.Open("hdb", connStr)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected.", "Data Reading", time.Now())
	}
	defer db.Close()
	query(db)
}

//* main query read

func query(db *sql.DB) {
	rows, err := db.Query("SELECT ID, NAME FROM GOIMPORT")
	if err != nil {
		log.Fatal(err)
	} else {
		var id int
		var c2 string
		for rows.Next() {
			err := rows.Scan(&id, &c2)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println(id, ", ", c2)
			}
		}
	}
	defer rows.Close()
}

//* select path query
//func queryContext(db *sql.DB) {
//queryStmt, err := db.Prepare("SELECT ID, C2 FROM T1 WHERE ID = ?")
//ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
//defer cancel()
//rows, err := queryStmt.QueryContext(ctx, 1)
