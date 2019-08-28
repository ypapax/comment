package comment

import (
	"time"
)

type Comment struct {
	Id      int
	Text    string
	UserID  int
	PageID  int
	Token   string
	Rating  int
	Created time.Time
}
