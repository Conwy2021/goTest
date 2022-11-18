package reflect_test

import (
	"fmt"
	"reflect"
	"testing"
)

type INT int

func Test202210270951(t *testing.T) {
	var a INT
	fmt.Println(reflect.TypeOf(a).Name()) //INT
	fmt.Println(reflect.TypeOf(a).Kind()) //int
}
