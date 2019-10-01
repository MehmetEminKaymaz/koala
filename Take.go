package koala

import "reflect"

func (c Collection) Take(num int) Collection { //it select records
	t := Task{
		KindOfTask: "Void",
		Do: func() interface{} {
			return func(col *Collection, g interface{}) {
				tk := g.(int)
				newSlice := reflect.MakeSlice(reflect.SliceOf(col.Source.Index(0).Type()), 0, 1)
				if tk > col.Source.Len() {
					tk = col.Source.Len()
				}
				for i := 0; i < tk; i++ {
					newSlice = reflect.Append(newSlice, col.Source.Index(i))
				}
				col.Source = newSlice
			}
		},
	}

	c.Tasks = append(c.Tasks, t)
	c.Items = append(c.Items, num)

	return c

}
