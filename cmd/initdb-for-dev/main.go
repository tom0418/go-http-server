package main

import (
	"flag"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	flagHost     = flag.String("h", "", "Database host")
	flagPort     = flag.Uint("p", 0, "Database port")
	flagDatabase = flag.String("d", "", "Database name")
)

func main() {
	flag.Parse()
	switch {
	case *flagHost == "":
		panic("-h is nil")
	case *flagHost != "mysql8":
		panic("-h is incorrect")
	}

	switch {
	case *flagPort == 0:
		panic("-p is nil")
	case *flagPort != 3306:
		panic("-p is incorrect")
	}

	switch {
	case *flagDatabase == "":
		panic("-d is nil")
	case *flagDatabase != "sample_api":
		panic("-p is incorrect")
	}

	// connect to DB
	dsn := fmt.Sprintf("root:password@tcp(%s:%d)/?charset=utf8mb4&parseTime=True&loc=Local", *flagHost, *flagPort)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect DB")
	}

	// init DB
	db.Exec(fmt.Sprintf("create database if not exists %s", *flagDatabase))
}
