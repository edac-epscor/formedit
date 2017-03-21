package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
//	"io/ioutil"
)

var VERSION string
var CODE string
var CODENAME string

//create a DB handle
var db *sql.DB
var formdb *sql.DB
var UploadDIR string
var cookieid string
var configf = ReadConfig()
func main() {
//	var configf = ReadConfig() //this is in config.go
	UploadDIR = configf.UploadDir
	cookieid = configf.CookieID
	LogFile := configf.LogDir + "formedit.log"
	var err error
	dbuser := configf.DBUsername
	dbpassword := configf.DBPassword
	dbname := configf.DBName
	dsn := dbuser + ":" + dbpassword + "@/" + dbname
	db, err = sql.Open("mysql", dsn) // this does not really open a new connection
	if err != nil {
		log.Fatalf("Error on initializing snort database connection: %s", err.Error())
	}
	db.SetMaxIdleConns(100)
	err = db.Ping() // This DOES open a connection if necessary. This makes sure the database is accessible
	if err != nil {
		log.Fatalf("Error on opening database connection: %s", err.Error())
	}


        formdbuser := configf.FormDBUsername
        formdbpassword := configf.FormDBPassword
        formdbname := configf.FormDBName
        formdsn := formdbuser + ":" + formdbpassword + "@/" + formdbname
        formdb, err = sql.Open("mysql", formdsn) // this does not really open a new connection
        if err != nil {
                log.Fatalf("Error on initializing form database connection: %s", err.Error())
        }
	formdb.SetMaxIdleConns(100)
        err = formdb.Ping() // This DOES open a connection if necessary. This makes sure the database is accessible
        if err != nil {
                log.Fatalf("Error on opening database connection: %s", err.Error())
        }




	f, err := os.OpenFile(LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Cant open log file")
	}
	defer f.Close()

	log.SetOutput(f)

        defer f.Close()


	listensocket := configf.IP + ":" + configf.Port
	router := NewRouter()
	log.Println("formedit running on " + listensocket)
	log.Fatal(http.ListenAndServe(listensocket, router))
}

func logErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
