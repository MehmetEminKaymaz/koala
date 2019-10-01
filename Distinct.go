package koala

import "reflect"

func (c Collection) Distinct() Collection { //get distinct items from source
	t := Task{
		KindOfTask: "Void2",
		Do: func() interface{} {
			return func(col *Collection) {

				slice := reflect.MakeSlice(reflect.SliceOf(col.Source.Index(0).Type()), 0, 1) //create a new slice
				IsHere := func(item interface{}) (state bool) {
					state = false
					for k := 0; k < slice.Len(); k++ {
						if slice.Index(k).Interface() == item {
							state = true
						}
					}
					return
				}
				for i := 0; i < col.Source.Len(); i++ {
					if !IsHere(col.Source.Index(i).Interface()) {
						slice = reflect.Append(slice, col.Source.Index(i))
					}
				}

				col.Source = slice

			}
		},
	}

	c.Tasks = append(c.Tasks, t)       //add task
	c.Items = append(c.Items, Empty{}) //add item of task

	return c
}
