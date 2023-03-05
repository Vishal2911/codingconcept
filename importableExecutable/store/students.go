package store

import "fmt"

type Store struct{}

func (s Store) SaveRecord(record interface{}) {
	s.saveRecord(record)
}

func (s Store)saveRecord(record interface{}) {
	fmt.Println("Record = ", record)
}
