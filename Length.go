package koala

import "reflect"

func(c Collection) Length()int {//len(collection)
	return reflect.Indirect(c.Source).Len()
}


func (c Collection) ToSlice() interface{}{ //source to slice

	return c.Source
}