package example

import (
	"github.com/rocna/go_tools"
	"testing"
)

func TestRandom(t *testing.T) {

	println(go_tools.Random.RandString(10))

	println(go_tools.Random.RandNumber(10))
}
