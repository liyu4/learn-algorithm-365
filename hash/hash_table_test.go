package hash

import (
	"testing"
)

func TestHash(t *testing.T) {

	hash := New()

	hash.Put("heloo", "word")

	hash.Put("hello", "word2")

	what := hash.Get("hello")

	t.Log(what)
}
