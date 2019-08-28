package comment

type Repository interface {
	Insert(*Comment) error
	FindByID(int) (*Comment, error)
	DeleteByID(int) error
	FindByPage(pageID int, skip, limit int) ([]Comment, error)
	FindAll(skip, limit int) ([]Comment, error)
}
