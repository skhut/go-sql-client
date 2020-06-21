package sqlclient

import (
	"errors"
	"fmt"
)

type rowsMock struct {
	Columns []string
	Rows    [][]interface{}

	currentIndex int
}

func (m *rowsMock) HasNext() bool {
	return m.currentIndex < len(m.Rows)
}

func (m *rowsMock) Close() error {
	return nil
}

func (m *rowsMock) Scan(destinations ...interface{}) error {
	mockedRow := m.Rows[m.currentIndex]
	if len(mockedRow) != len(destinations) {
		return errors.New("Invalid destination len")
	}

	for index, value := range mockedRow {
		//fmt.Printf("Index:%d  Value:%#v\n", index, value)
		destinations[index] = value
	}
	fmt.Printf("destinations:%#v\n", destinations)
	return nil
}
