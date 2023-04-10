package test

import (
	"testing"
)

func TestArticleGetList(t *testing.T) {
	BaseSuccessTest(t, Get("/api/article/list", backRouter))
}

func TestArticleGetInfo(t *testing.T) {
	BaseSuccessTest(t, Get("/api/article/1", backRouter))
}

func TestArticleSoftDelete(t *testing.T) {
	param := map[string]any{"ids": []int{1, 2}, "is_delete": 1}
	BaseSuccessTest(t, PutJson("/api/article/softDelete", param, backRouter))
}

func TestArticleDelete(t *testing.T) {
	ids := []int{11, 23}
	BaseSuccessTest(t, DeleteJson("/api/article", ids, backRouter))
}
