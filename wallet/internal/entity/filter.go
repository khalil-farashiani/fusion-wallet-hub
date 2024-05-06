package entity

import "time"

type Paginate struct {
	Cursor int64
	Limit  int64

	AfterDateTime  time.Time
	BeforeDateTime time.Time
}
