package hash

import (
	"testing"
)

func TestHash(t *testing.T) {
	hash := New()

	t.Log(hash)

	h := Hash("hellooo")
	t.Log(h)
}
