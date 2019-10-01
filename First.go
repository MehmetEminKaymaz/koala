package koala

func (c Collection) First() Collection { //it returns the first element of the collection
	t := Task{
		KindOfTask: "Void2",
		Do: func() interface{} {
			return func(col *Collection) {
				col.Source = col.Source.Index(0)
			}
		},
	}

	c.Tasks = append(c.Tasks, t)
	c.Items = append(c.Items, Empty{})
	return c
}
