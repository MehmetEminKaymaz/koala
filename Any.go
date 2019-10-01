package koala

import "reflect"

func (c Collection) Any(look func(i interface{}) bool) Collection { //it returns a boolean by look function's result
	t := Task{
		KindOfTask: "Bool",
		Do: func() interface{} {
			return func(col *Collection, lookk func(i interface{}) bool) {
				var state bool = false

				for i := 0; i < col.Source.Len(); i++ {
					if lookk(col.Source.Index(i).Interface()) {
						state = true
						col.Source = reflect.ValueOf(state)
						break
					}

				}

				col.Source = reflect.ValueOf(state)
			}
		},
	}

	c.Tasks = append(c.Tasks, t)    //add task
	c.Items = append(c.Items, look) // add item of task

	return c

}
