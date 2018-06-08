package main

import (
	"subLease/server"
	"subLease/server/database"
)

func main() {
	s := server.Create(database.Create())
	s.Run()
}
