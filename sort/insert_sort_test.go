package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInserSort(t *testing.T) {
	arr := []int{1, 2, 3, 9, 5, 7, 6}
	ret := insertSort(arr)

	if !reflect.DeepEqual(arr, ret) {
		t.Fail()
	}
	fmt.Println(ret)
}
