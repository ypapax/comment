package page

import "github.com/ypapax/comment/site"

type Page struct {
	ID   string
	Site site.Site
	Path string
}
