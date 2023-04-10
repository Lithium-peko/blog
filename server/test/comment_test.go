package test

import (
	"testing"
)

func TestCommentGetList(t *testing.T) {
	BaseSuccessTest(t, Get("/api/comment/list", backRouter))
}

func TestCommentDelete(t *testing.T) {
	ids := []int{11}
	BaseSuccessTest(t, DeleteJson("/api/comment", ids, backRouter))
}

func TestUpdateCommentReview(t *testing.T) {
	param := map[string]any{"ids": []int{11, 12}, "is_review": 1}
	BaseSuccessTest(t, PutJson("/api/comment/review", param, backRouter))
}
