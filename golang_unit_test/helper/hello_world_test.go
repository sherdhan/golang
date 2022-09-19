package helper

import (
	"fmt"
	"runtime"
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

	require.Equal(t, "Hello sherdhan", result, "required: hasil harus sesuai")
	fmt.Println("test with required done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("can't run in mac")
	}

	result := Helloworld("sherdhan")

	require.Equal(t, "Hello sherdhan", result, "skip: hasil harus sesuai")
}

func TestMain(m *testing.M) {
	fmt.Println("before unit test")

	m.Run()

	fmt.Println("after unit test")
}
