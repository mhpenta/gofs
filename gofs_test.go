package gofs_test

import (
	"testing"

	"github.com/mhpenta/gofs"
)

type HelloFuncParams struct {
	world string `json:"world"`
}

func hello(helloFuncParams HelloFuncParams) {}

type HelloFuncParams2 struct {
	world string `json:"worldParam"`
}

func hello2(helloFuncParams HelloFuncParams2) {}

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

	details, err = gofs.GetFunctionDetails(hello2)
	if err != nil {
		t.Error(err)
	}
	if details.Name != "hello2" {
		t.Errorf("Expected 'hello2', got %v", details)
	}

	if details.Parameters[0].Title != "worldParam" {
		t.Errorf("Expected 'worldParam', got %v", details.Parameters[0].Title)
	}
}
