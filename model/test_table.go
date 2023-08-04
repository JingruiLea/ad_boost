package model

import "time"

type Test struct {
	ID        int
	Name      string
	Age       int
	CreatedAt time.Time
}

func (m *Test) TableName() string {
	return "test_table"
}
