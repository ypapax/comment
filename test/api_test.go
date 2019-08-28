package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/ypapax/comment/comment"
	"testing"
)

func TestAddComment(t *testing.T) {
	as := assert.New(t)
	var com = comment.Comment{Text: "some comment text", UserID: 1, PageID: 2}
	status, b, _, err := pathReq("/comment", "POST", com)
	if !as.NoError(err) {
		return
	}
	if !as.Equal(200, status) {
		return
	}
	var respComment comment.Comment
	if !as.NoError(json.Unmarshal(b, &respComment)) {
		return
	}
	if !as.NotZero(respComment.Id) {
		return
	}
	if !as.NotZero(respComment.Created) {
		return
	}
	if !as.Equal(com.PageID, respComment.PageID) {
		return
	}
	if !as.Equal(com.UserID, respComment.UserID) {
		return
	}
}
