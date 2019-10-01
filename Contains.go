package koala

import "reflect"

//Contains returns a boolean value , if exist true else false
//its also variadic function

func(c Collection) Contains(item ...interface{})Collection{

	t:=Task{ //create task
		KindOfTask:"Void", //type of task
		Do: func() interface{} { //the function return the task
			return func(col *Collection,source interface{}){

				slicev:=reflect.ValueOf(source)
				var state =false
				for i:=0;i<slicev.Len();i++{ //loop over slice
					for k:=0; k<col.Source.Len();k++{
						if col.Source.Index(k).Interface()==slicev.Index(i).Interface(){
							state=true
							col.Source=reflect.ValueOf(state)
							return
						}
					}
				}
				col.Source=reflect.ValueOf(state)//change source
			}
		},
	}

	c.Tasks=append(c.Tasks,t) //update collection's tasks
	c.Items=append(c.Items,item) //update collection's items

	return c //return the collection
}
