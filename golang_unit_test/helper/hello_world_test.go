package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloworld(t *testing.T) {
	result := Helloworld("sherdhan")

	if result != "Hello sherdhan" {
		// panic("salah")
		// t.Fail()
		t.Error("error: hasil tidak sesuai")
	}

	fmt.Println("func pertama done")
}

func TestHelloworld2(t *testing.T) {
	result := Helloworld("sherdhan")

	if result != "Hello sherdhan" {
		// panic("salah")
		// t.FailNow()
		t.Fatal("fatal: hasil tidak sesuai")
	}

	fmt.Println("func kedua done")
}

func TestHelloworldassert(t *testing.T) {
	result := Helloworld("sherdhan")

	assert.Equal(t, "Hello sherdhan", result, "assert: hasil harus sesuai")
	fmt.Println("test with assert done")
}

func TestHelloworldrequire(t *testing.T) {
	result := Helloworld("sherdhan")

	require.Equal(t, "Hello sherdhan", result, "assert: hasil harus sesuai")
	fmt.Println("test with required done")
}
