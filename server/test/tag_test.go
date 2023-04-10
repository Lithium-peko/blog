package test

import (
	"testing"
)

func TestTagGetList(t *testing.T) {
	BaseSuccessTest(t, Get("/api/tag/list", backRouter))
}

func TestTagDelete(t *testing.T) {
	ids := []int{11, 12}
	BaseSuccessTest(t, DeleteJson("/api/tag", ids, backRouter))
}

func TestTagSaveOrUpdate(t *testing.T) {
	param := map[string]any{"id": 1, "name": "前端"}
	BaseSuccessTest(t, PostJson("/api/tag", param, backRouter))
}
