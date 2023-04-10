package test

import (
	"testing"
)

func TestCategoryGetList(t *testing.T) {
	BaseSuccessTest(t, Get("/api/category/list", backRouter))
}

func TestCategoryGetFrontList(t *testing.T) {
	BaseSuccessTest(t, Get("/api/front/category/list", frontRouter))
}
