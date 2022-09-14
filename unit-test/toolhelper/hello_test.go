package toolhelper

import (
	"testing"
)

func TestTester_helper(t *testing.T) {
	res := Tester_helper("sherdhan")

	if res != "Hello sherdhan" {
		panic("result error")
	}
}
