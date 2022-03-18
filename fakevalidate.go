package main

import (
	"fmt"
)

type fakeDB struct {
	email    string
	password string
}
type validate interface {
	validate() bool
}

func fakeLogin(email, pass string) *fakeDB {
	return &fakeDB{email, pass}
}
func (fdb *fakeDB) validate() bool {
	if fdb.email == "Fadhli" && fdb.password == "Abrar" {
		return true
	} else {
		return false
	}
}

func main() {
	test := fakeLogin("Fadhli", "Abrar").validate()

	fmt.Println(test)

}
