package main

import (
	"subLease/src/server"
	"subLease/test/utils/mockDatabase"
)

func main() {
	s := server.Create(mockDatabase.Create())
	s.Run()
}
