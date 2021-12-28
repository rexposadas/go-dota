package main

import "fmt"

const (
	RECORDS = "records"
)

func Records(id int) ([]byte, error) {
	path := fmt.Sprintf("/%s/%d", RECORDS, id)
	return Get(path, nil)
}
