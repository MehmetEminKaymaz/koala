package koala


func(c Collection) LazyCancel(count int) Collection{//it cancels the last (count) tasks!

t:=Task{
	KindOfTask:"Void2",
	Do: func() interface{} {
		return func(col *Collection){
			c.Tasks=c.Tasks[:len(c.Tasks)-count] //clear tasks
			c.Items=c.Items[:len(c.Items)-count]//clear item of tasks
		}
	},
}

c.Tasks=append(c.Tasks,t)
c.Items=append(c.Items,Empty{})



	return c
}