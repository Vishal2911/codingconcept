package store

type DBOperations interface {
	SaveRecord(record interface{})
}