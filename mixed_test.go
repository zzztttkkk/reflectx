package reflectx_test

import (
	"fmt"
	"testing"

	"github.com/zzztttkkk/reflectx"
)

func init() {
	reflectx.RegisterOf[reflectx.EmptyMeta]().TagNames("json").Unexposed()
}

type _CommonA struct {
	A string `json:"a"`
	B string `json:"b"`
}

func init() {
	ptr := reflectx.Ptr[_CommonA]()
	reflectx.FieldOf[_CommonA, reflectx.EmptyMeta](&ptr.A).SetName("common_a")
}

type _X struct {
	_CommonA
}

type UserA struct {
	A string `json:"a1"`
	_CommonA
	_X
}

func TestMixed(t *testing.T) {
	ptr := reflectx.Ptr[UserA]()

	field1 := reflectx.FieldOf[UserA, reflectx.EmptyMeta](&ptr._CommonA.A)
	fmt.Println(field1.StructField(), field1.Offset())
	field2 := reflectx.FieldOf[UserA, reflectx.EmptyMeta](&(ptr._X._CommonA.A))
	fmt.Println(field2.StructField(), field2.Offset())
	fmt.Println(field1.StructField() == field2.StructField())

	for _, v := range reflectx.TypeInfoOf[UserA, reflectx.EmptyMeta]().Fields {
		fmt.Println(v.String())
	}
}
