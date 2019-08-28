package comment

import "time"

type Service interface {
	Insert(*Comment) error
	FindByID(int) (*Comment, error)
	DeleteByID(int) error
	FindByPage(pageID, pageNumber, limit int) ([]Comment, error)
	FindAll(page, limit int) ([]Comment, error)
}

type commentService struct {
	repo Repository
}

func NewCommentService(repo Repository) Service {
	return &commentService{repo: repo}
}

func (s commentService) Insert(c *Comment) error {
	c.Created = time.Now()
	return s.repo.Insert(c)
}
func (s commentService) FindByID(id int) (*Comment, error) {
	return s.repo.FindByID(id)
}
func (s commentService) DeleteByID(id int) error {
	return s.repo.DeleteByID(id)
}
func (s commentService) FindByPage(pageID, skip, limit int) ([]Comment, error) {
	return s.repo.FindByPage(pageID, skip, limit)
}

func (s commentService) FindAll(skip, limit int) ([]Comment, error) {
	return s.repo.FindAll(skip, limit)
}
