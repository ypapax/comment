package comment

type Repository interface {
	Insert(*Comment) error
	FindByID(string) (*Comment, error)
	DeleteByID(string) error
	FindByPage(pageID string, skip, limit int) ([]Comment, error)
	FindAll(skip, limit int) ([]Comment, error)
}
