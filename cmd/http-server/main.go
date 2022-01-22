package main

import database "github.com/tom0418/go-http-server/db"

func main() {
	database.Init()
	database.Close()
}
