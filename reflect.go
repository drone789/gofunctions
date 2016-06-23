package main

import (
	"fmt"
	"reflect"
)

func index(slice interface{}, v interface{}) int {
	if slice := reflect.ValueOf(slice); slice.Kind() == reflect.Slice {
		for i := 0; i < slice.Len(); i++ {
			if reflect.DeepEqual(v, slice.Index(i).Interface()) {
				return i
			}
		}
	}
	return -1
}

func main() {
	slice1 :=[]int{1,2,3,5,6,7}
	fmt.Println(index(slice1,3))

	slice2 :=[]float64{2.3,4.5,5.6,2.7,6.9}
	fmt.Println(index(slice2, 6.9))

	slice3 :=[]interface{}{'a',2,4,	6.6,'c',"abc"}
	fmt.Println(index(slice3,6.6))
}
