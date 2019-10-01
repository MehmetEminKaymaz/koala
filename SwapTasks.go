package koala

func (c Collection) SwapTasks(index1, index2 int) Collection { //it swap tasks before execute!
	c.Tasks[index1], c.Tasks[index2] = c.Tasks[index2], c.Tasks[index1] //swap tasks
	c.Items[index1], c.Items[index2] = c.Items[index2], c.Items[index1] //swap items of tasks
	return c
}
