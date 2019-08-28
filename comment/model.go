package comment

import (
	"github.com/ypapax/comment/page"
	"time"
)

type Comment struct {
	Id      string
	Text    string
	UserID  string
	PageID  string `sql:"page_id"`
	Page    page.Page
	Token   string
	Rating  int
	Created time.Time
}
