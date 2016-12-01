package hash

import (
	"testing"
)

func TestHash(t *testing.T) {

	hash := New()

	var keys = []string{"name", "address", "phone", "k101", "k110"}
	var values = []string{"Sourav", "Sinagor", "26300788", "Value1", "Value2"}

	for i := 0; i < 5; i++ {
		hash.Put(keys[i], values[i])
	}

	hash.display()

}
