// *GLDS HANA DB stop
// * Copyright Rushi NS - SAP


package main

// * import SAP Driver and standard packages

import (
// *    "cont!
    "time"
    "fmt"
    "log"
    "database/sql"
    _ "SAP/go-hdb/driver"
)

// * main logic 
// * connection details

func main() {
    // define a connection string
    var connStr string
//  multi container string
//  connStr = "hdb://system:Abcd1234@10.48.146.57:30013?DATABASENAME=GDB" 
connStr = "hdb://system:Abcd1234@10.48.146.57:30013" 
    db, err := sql.Open("hdb", connStr)
    if err != nil {
        log.Fatal(err)
    } else {
        fmt.Println("Connected.", "Date and Time", time.Now())
    }
    defer db.Close()
query(db)
}
// * main query read

func query(db *sql.DB) {
    rows, err := db.Query("ALTER SYSTEM STOP DATABASE TESTDB")
    if err != nil {
        log.Fatal(err)
    } else {
                fmt.Println("Multi Tenant DB halted")
            }
    defer rows.Close()
}
