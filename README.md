# koala
Immutable and Lazy loading Golang library to deal with slices


### Installation

```
 
  go get github.com/MehmetEminKaymaz/koala
 
 
```

### Documentation


#### koala.Add

It adds items to the source , also it is variadic function.

```Go

list = list.Add(24)//add 24 to the source in one task or

list = list.Add(24,25,26,27) //add 24,25,26 and 27 to the source in one task

```

### koala.Where

It returns data according to a given condition

```Go

list=list.Where(func(x interface{}) bool { 
     return x.(int)>=10
    })

```


### koala.Foreach

```Go

list=list.Foreach(func(x interface{}) (y interface{}) {
      y=x.(int)*x.(int)
      return
    })

```

### koala.Contains

It returns a boolean value , if exist true else false

```Go

list=list.Contains(1) 
//or
list = list.Contains(2,3,4,5,6)//variadic function
//you can also pass a slice

```

### koala.SwapTask

It swaps task before execution 

```Go

list=list.Add(24)
list=list.Add(25)
//so output should be [24,25]

//but

list = list.SwapTask(0,1) // it swaps Add(24) and Add(25) so output should be [25,24]



```

### koala.Do

Tasks can not be execute before calling Do , so do function executes all task.

```Go

list=list.Add(24)
list=list.Add(25)

list = list.SwapTask(0,1) 

result = list.Do().Value() // result = [25,24]



```

### koala.Value

It returns the result


### koala.Cancel

It deletes task before execute

```Go

list=list.Add(24)
list=list.Add(25)

list = list.SwapTask(0,1) 

list=list.Cancel(2) // SwapTask is cancelled

//so output [24,25]



```

### koala.LazyCancel

It deletes task in execution

### koala.ClearTasks

It clear all tasks before execution

### koala.Lenght

It returns lenght of the source //source must be slice or array


### koala.Reverse

It reverse the source

### koala.RemoveAt

Remove item from source using index

### koala.ElementAt

Get element from source using index

### koala.Skip

Skip items in source 






