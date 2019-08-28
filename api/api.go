package api

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/ypapax/comment/comment"
	"github.com/ypapax/comment/config"
	"github.com/ypapax/comment/database"
	"net/http"
	"strconv"
)

func Serve(c config.Config) error {
	commentService, err := database.DbServices(c.DbType, c.Db.Addr, c.Db.User, c.Db.Name)
	if err != nil {
		logrus.Error(err)
		return err
	}
	e := echo.New()

	e.POST("/comment", func(c echo.Context) error {
		var com comment.Comment
		if err := c.Bind(&com); err != nil {
			logrus.Error(err)
			return c.String(http.StatusBadRequest, "bad request br989891")
		}
		if err := commentService.Insert(&com); err != nil {
			logrus.Error(err)
			return c.String(http.StatusInternalServerError, "internal server error ise879187090")
		}
		return c.JSON(http.StatusOK, com)
	})

	e.GET("/comment/:id", func(c echo.Context) error {
		id, err := intParam(c, "id")
		if err != nil {
			logrus.Error(err)
			return c.String(http.StatusBadRequest, "bad request br989823")
		}
		com, err := commentService.FindByID(id)
		if err != nil {
			logrus.Error(err)
			return c.String(http.StatusInternalServerError, "internal server error ise879187090")
		}
		return c.JSON(http.StatusOK, com)
	})

	e.DELETE("/comment/:id", func(c echo.Context) error {
		id, err := intParam(c, "id")
		if err != nil {
			logrus.Error(err)
			return c.String(http.StatusBadRequest, "bad request br989823")
		}
		if err := commentService.DeleteByID(id); err != nil {
			logrus.Error(err)
			return c.String(http.StatusInternalServerError, "internal server error isd8092323")
		}
		return c.JSON(http.StatusOK, "ok")
	})

	e.GET("/comment/:page-id", func(c echo.Context) error {
		pageID, err := intParam(c, "page-id")
		if err != nil {
			logrus.Error(err)
			return c.String(http.StatusBadRequest, "bad request br989823")
		}
		page, limit, err := parsePaginationParams(c)
		if err != nil {
			logrus.Error(err)
			return c.String(http.StatusBadRequest, "bad request br23ljlk8")
		}
		cc, err := commentService.FindByPage(pageID, page, limit)
		if err != nil {
			logrus.Error(err)
			return c.String(http.StatusInternalServerError, "internal server error ise23232")
		}
		return c.JSON(http.StatusOK, cc)
	})

	e.GET("/comment", func(c echo.Context) error {
		page, limit, err := parsePaginationParams(c)
		if err != nil {
			logrus.Error(err)
			return c.String(http.StatusBadRequest, "bad request br883838")
		}
		cc, err := commentService.FindAll(page, limit)
		if err != nil {
			logrus.Error(err)
			return c.String(http.StatusInternalServerError, "internal server error ise89820902")
		}
		return c.JSON(http.StatusOK, cc)
	})

	logrus.Printf("listening %+v", c.Api.Bind)
	if err := e.Start(c.Api.Bind); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func parsePaginationParams(c echo.Context) (page int, limit int, err error) {
	var params = []string{"page", "limit"}
	var intParams []int
	for _, p := range params {
		i, err := strconv.ParseInt(c.QueryParam(p), 10, 64)
		if err != nil {
			err := fmt.Errorf("couldn't parse integer from paramater %s", p)
			logrus.Error(err)
			return 0, 0, err
		}
		intParams = append(intParams, int(i))
	}
	return intParams[0], intParams[1], nil
}

func intParam(c echo.Context, name string) (int, error) {
	val := c.Param(name)
	if len(val) == 0 {
		return 0, fmt.Errorf("missing parameter %+v", name)
	}
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		logrus.Error(err)
		return 0, err
	}
	return int(i), err
}
