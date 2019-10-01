package koala

import "reflect"

func (c Collection) Reverse() Collection { //reverse the collection
	t := Task{
		KindOfTask: "Void2",
		Do: func() interface{} {
			return func(col *Collection) {
				newSlice := reflect.MakeSlice(reflect.SliceOf(col.Source.Index(0).Type()), 0, 1)
				for i := col.Source.Len() - 1; i >= 0; i-- {
					newSlice = reflect.Append(newSlice, col.Source.Index(i))
				}
				col.Source = newSlice
			}
		},
	}

	c.Tasks = append(c.Tasks, t)
	c.Items = append(c.Items, Empty{})

	return c
}
