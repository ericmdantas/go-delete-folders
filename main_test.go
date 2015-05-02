package main

import (
	_ "github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveThisAndChildren(t *testing.T) {
	f := Folder{"a"}

	RemoveThisAndChildren(f)
}

func TestInit(t *testing.T) {
	Init()
}

func BenchmarkInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Init()
	}
}

func BenchmarkRemoveThisAndChildren(b *testing.B) {
	f := Folder{"a"}

	for i := 0; i < b.N; i++ {
		RemoveThisAndChildren(f)
	}
}
