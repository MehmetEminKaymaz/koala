package koala

import "reflect"

func (c Collection) RemoveAt(index int) Collection { //it removes element at index
	t := Task{
		KindOfTask: "Void",
		Do: func() interface{} {
			return func(col *Collection, i interface{}) {
				newSlice := reflect.MakeSlice(reflect.SliceOf(col.Source.Index(0).Type()), 0, 1)

				for i := 0; i < col.Source.Len(); i++ {
					if index != i {
						newSlice = reflect.Append(newSlice, col.Source.Index(i))
					}

				}
				col.Source = newSlice
			}
		},
	}

	c.Tasks = append(c.Tasks, t)     //add task
	c.Items = append(c.Items, index) //add item of task

	return c
}
