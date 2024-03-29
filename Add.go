package koala

import "reflect"

func (c Collection) Add(item ...interface{}) Collection { //add item to source (variadic function)
	decide := len(item)
	if decide == 1 {
		t := Task{
			KindOfTask: "Void", //it means func does not return a value
			Do: func() interface{} { //return interface to return different functions
				return func(col *Collection, itemm interface{}) {
					col.Source = reflect.Append(col.Source, reflect.ValueOf(itemm)) //add item to source

				}

			},
		}
		c.Tasks = append(c.Tasks, t)                                          //update task for the collection
		c.Items = append(c.Items, reflect.ValueOf(item).Index(0).Interface()) //update items for the collection

	} else {

		//we have many item to add to the source!

		t := Task{
			KindOfTask: "Void", //it means func does not return a value
			Do: func() interface{} { //return interface to return different functions
				return func(col *Collection, itemm interface{}) {

					slice := reflect.ValueOf(itemm)
					for i := 0; i < slice.Len(); i++ { //loop over the slice
						col.Source = reflect.Append(col.Source, reflect.ValueOf(slice.Index(i).Interface())) //add item to source
					}

				}
			},
		}

		c.Tasks = append(c.Tasks, t)    //add task
		c.Items = append(c.Items, item) //ad item of task

	}

	return c
}
