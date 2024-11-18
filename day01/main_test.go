package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	fmt.Println("Hello, World!")
	assert.Equal(t, "Hello, World!", "Hello, World!")
}
