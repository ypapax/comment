package psql

import (
	"fmt"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/urlfilter"
	"github.com/sirupsen/logrus"
	"github.com/ypapax/comment/comment"
)

type commentRepository struct {
	db *pg.DB
}

func NewPostgresCommentRepository(db *pg.DB) comment.Repository {
	return &commentRepository{db: db}
}

func (r commentRepository) Insert(c *comment.Comment) error {
	if err := func() error {
		err := r.db.Select(c.Page)
		if err == nil {
			return err
		}
		if !notFound(err) {
			logrus.Error(err)
			return err
		}
		if err := r.db.Insert(c.Page); err != nil {
			logrus.Error(err)
			return err
		}
		return nil
	}(); err != nil {
		logrus.Error(err)
		return err
	}
	if err := r.db.Insert(c); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func (r commentRepository) FindByID(id string) (*comment.Comment, error) {
	var c = comment.Comment{Id: id}
	err := r.db.Select(&c)
	if err != nil {
		if notFound(err) {
			return &c, nil
		}
		logrus.Error(err)
		return nil, err
	}
	return &c, nil
}

func (r commentRepository) DeleteByID(id string) error {
	if err := r.db.Delete(comment.Comment{Id: id}); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
func (r commentRepository) FindByPage(pageID string, page, limit int) ([]comment.Comment, error) {
	values := urlfilter.Values(map[string][]string{
		"page":  {fmt.Sprintf("%d", page)},
		"limit": {fmt.Sprintf("%d", limit)},
	})
	var cc []comment.Comment
	if err := r.db.Model(&cc).Where("page_id=?", pageID).Apply(urlfilter.Pagination(values)).Select(); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return cc, nil
}
func (r commentRepository) FindAll(skip, limit int) ([]comment.Comment, error) {
	var cc []comment.Comment
	if err := r.db.Model(&cc).Select(); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return cc, nil
}
