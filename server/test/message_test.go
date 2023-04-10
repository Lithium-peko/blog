package test

import (
	"testing"
)

func TestMessageGetList(t *testing.T) {
	BaseSuccessTest(t, Get("/api/message/list", backRouter))
}

func TestMessageDelete(t *testing.T) {
	ids := []int{4, 5}
	BaseSuccessTest(t, DeleteJson("/api/message", ids, backRouter))
}

func TestMessageReview(t *testing.T) {
	param := map[string]any{"ids": []int{1, 2}, "is_review": 1}
	BaseSuccessTest(t, PutJson("/api/message/review", param, backRouter))
}
