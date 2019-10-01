package koala

func (c Collection) ClearTasks() Collection {

	c.Tasks = c.Tasks[:0]
	c.Items = c.Items[:0]

	return c
}
