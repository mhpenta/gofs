package gofs_test

import (
	"testing"

	"github.com/samjtro/gofs"
)

func hello(world string) {}

func Test(t *testing.T) {
	_ = gofs.Get(hello)
}
