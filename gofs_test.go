package gofs_test

import (
	"testing"

	"github.com/samjtro/gofs"
)

type HelloFuncParams struct {
	world string `json:"world"`
}

func hello(helloFuncParams HelloFuncParams) {}

func Test(t *testing.T) {

	details, err := gofs.GetFunctionDetails(hello)
	if err != nil {
		t.Error(err)
	}
	if details.Name != "hello" {
		t.Errorf("Expected 'hello', got %v", details)
	}

	if details.Parameters[0].Title != "world" {
		t.Errorf("Expected 'world', got %v", details.Parameters[0].Title)
	}

}
