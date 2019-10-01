package koala

import (
	"errors"
	"reflect"
)





type Task struct {
	KindOfTask string
	Do func() interface{}
}

type Empty struct {

}


type Collection struct {
	Source reflect.Value
	Tasks []Task
	Items []interface{}

}


func From(source interface{}) (Collection,error){

	val:=reflect.ValueOf(source)
	switch val.Kind() {

	case reflect.Slice,reflect.Array:

		myMap:=make( []Task,0,1)
		myItems:=make([]interface{},0,1)


		return Collection{
			Source:reflect.ValueOf(source),
			Tasks:myMap,
			Items:myItems,

		},nil

	default:

		return Collection{},errors.New("It is not slice or array")

	}

}

