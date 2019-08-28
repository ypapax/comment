package database

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/ypapax/comment/comment"
	"github.com/ypapax/comment/database/psql"
	"time"
)

func DbServices(dbType, addr, user, database string) (comment.Service, error) {
	var companyService comment.Service

	switch dbType {
	case "psql":
		db, err := psql.GetConnection(addr, user, database, 10*time.Second)
		if err != nil {
			err := fmt.Errorf("err: %+v for addr %+v", err, addr)
			logrus.Error(err)
			return nil, err
		}
		companyRepo := psql.NewPostgresCommentRepository(db)
		companyService = comment.NewCommentService(companyRepo)
	default:
		err := fmt.Errorf("db type '%+v' is not supported", dbType)
		return nil, err
	}
	return companyService, nil
}
