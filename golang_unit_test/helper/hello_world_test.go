package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkSub(b *testing.B) {
	b.Run("sherdhan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Helloworld("sherdhan")
		}
	})

	b.Run("s", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Helloworld("s")
		}
	})
}

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		nama    string
		request string
	}{
		{
			nama:    "sherdhan",
			request: "sherdhan",
		},
		{
			nama:    "syarif",
			request: "syarif",
		},
		{
			nama:    "ahmad",
			request: "ahmad",
		},
	}

	for _, v := range benchmarks {
		b.Run(v.nama, func(b *testing.B) {
			Helloworld(v.request)
		})
	}
}

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

func TestSubtest(t *testing.T) {
	t.Run("sherdhan", func(t *testing.T) {
		result := Helloworld("sherdhan")
		require.Equal(t, "Hello sherdhan", result, "TestSubTest: hasil harus sesuai")
	})
	t.Run("syarif", func(t *testing.T) {
		result := Helloworld("syarif")
		require.Equal(t, "Hello syarif", result, "TestSubTest: hasil harus sesuai")
	})
}

func TestTabletest(t *testing.T) {
	tests := []struct {
		name     string
		req      string
		expected string
	}{
		{
			name:     "test pertama",
			req:      "sherdhan",
			expected: "Hello sherdhan",
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			res := Helloworld(v.req)
			require.Equal(t, v.expected, res)
		})
	}
}
func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	fmt.Println("AFTER UNIT TEST")
}
