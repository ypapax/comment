package comment

type Service interface {
	Insert(*Comment) error
	FindByID(string) (*Comment, error)
	DeleteByID(string) error
	FindByPage(pageID string, pageNumber, limit int) ([]Comment, error)
	FindAll(page, limit int) ([]Comment, error)
}

type commentService struct {
	repo Repository
}

func NewCommentService(repo Repository) Service {
	return &commentService{repo: repo}
}

func (s commentService) Insert(c *Comment) error {
	return s.repo.Insert(c)
}
func (s commentService) FindByID(id string) (*Comment, error) {
	return s.repo.FindByID(id)
}
func (s commentService) DeleteByID(id string) error {
	return s.repo.DeleteByID(id)
}
func (s commentService) FindByPage(pageID string, skip, limit int) ([]Comment, error) {
	return s.repo.FindByPage(pageID, skip, limit)
}

func (s commentService) FindAll(skip, limit int) ([]Comment, error) {
	return s.repo.FindAll(skip, limit)
}
