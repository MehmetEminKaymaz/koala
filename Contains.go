package koala

import "reflect"


//return a boolean value , if exist true else false
//its also variadic function
func(c Collection) Contains(item ...interface{})Collection{

	t:=Task{
		KindOfTask:"Void",
		Do: func() interface{} {
			return func(col *Collection,source interface{}){

				slicev:=reflect.ValueOf(source)
				var state =false
				for i:=0;i<slicev.Len();i++{
					for k:=0; k<col.Source.Len();k++{
						if col.Source.Index(k).Interface()==slicev.Index(i).Interface(){
							state=true
							col.Source=reflect.ValueOf(state)
							return
						}
					}
				}
				col.Source=reflect.ValueOf(state)
			}
		},
	}

	c.Tasks=append(c.Tasks,t)
	c.Items=append(c.Items,item)

	return c
}
