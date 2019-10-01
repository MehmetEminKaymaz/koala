package koala

func (c Collection) ElementAt(index int) Collection { //it returns element by index value
	t := Task{
		KindOfTask: "Void",
		Do: func() interface{} {
			return func(col *Collection, i interface{}) {
				col.Source = col.Source.Index(i.(int))
			}
		},
	}

	c.Tasks = append(c.Tasks, t)
	c.Items = append(c.Items, index)

	return c

}
