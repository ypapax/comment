package psql

import (
	"fmt"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"github.com/sirupsen/logrus"
	"github.com/ypapax/comment/comment"
	"github.com/ypapax/comment/page"
	"strings"
	"time"
)

func GetConnection(addr, user, database string, timeout time.Duration) (*pg.DB, error) {
	db := pg.Connect(&pg.Options{
		Addr:     addr,
		User:     user,
		Database: database,
	})
	done := make(chan bool)
	var schemaErr error
	go func() {
		for {
			schemaErr = CreateSchema(db)
			if schemaErr != nil {
				logrus.Trace(schemaErr)
				time.Sleep(time.Second)
				continue
			}
			done <- true
			break
		}
	}()
	select {
	case <-time.After(timeout):
		err := fmt.Errorf("timeout: %s, couldn't create schema: %+v", timeout, schemaErr)
		logrus.Error(err)
		return nil, err
	case <-done:
	}
	return db, nil
}

func notFound(err error) bool {
	if err == nil {
		return false
	}
	if !strings.Contains(err.Error(), "no rows in result set") {
		return false
	}
	return true
}

func CreateSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*page.Page)(nil), (*comment.Comment)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists:   true,
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
