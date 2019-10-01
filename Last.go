package koala


func(c Collection) Last() Collection{//it returns the last element of the collection
	t:=Task{
		KindOfTask:"Void2",
		Do: func() interface{} {
			return func(col *Collection){
				col.Source=col.Source.Index(col.Source.Len()-1)
			}
		},
	}

	c.Tasks=append(c.Tasks,t)
	c.Items=append(c.Items,Empty{})

	return c
}
