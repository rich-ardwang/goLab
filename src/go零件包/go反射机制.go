package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age int
	Sex string
}

func (user User) Say (msg string) string {
	fmt.Println("say",msg)
	return "It's OK."
}

func main() {
	fmt.Println("main start...")
	//反射定律
	//1.反射可以将“接口类型变量”转换为“反射类型对象”
	//2.反射可以将“反射类型对象”转换为“接口类型变量”
	//3.如果要修改“反射类型对象”，其值必须是“可写的”
	//所谓的反射类型，就是reflect.Type和reflect.Value

	//1.“接口类型变量”=>“反射类型对象”
	var a int = 30
	v := reflect.ValueOf(a) //返回reflect.Value类型对象，值为30
	t := reflect.TypeOf(a) //返回reflect.Type类型对象，值为int
	fmt.Println(v)
	fmt.Println(t)
	v = reflect.ValueOf(&a) //返回reflect.Value类型对象，值为&a，变量a的地址
	t = reflect.TypeOf(&a) //返回reflect.Type类型对象，值为*int
	fmt.Println(v)
	fmt.Println(t)
	fmt.Println("---------------------------------------------------------------")

	//2.“反射类型对象”=>“接口类型变量”
	//最关键的两步
	//a.v1 := value.Interface()  // 返回的是一个接口变量
	//b.v2 := v1.(float64)       // 再判断这个接口变量是否能转化为某类型变量
	var b int = 30
	value := reflect.ValueOf(&b) //返回reflect.Value类型对象，值为&b，变量b的地址
	t1 := value.Interface().(*int) //类型断言，断定value中type=*int
	fmt.Printf("%T %v\n", t1, t1) // *int 0xc420086008
	fmt.Println(*t1) // 30

	cir := 6.28
	value2 := reflect.ValueOf(cir)
	t2 := value2.Interface().(float64)
	fmt.Printf("%T %v\n", t2, t2) // float64 6.28
	fmt.Println(t2) // 6.28

	/*
	t3 := value2.Interface().(int) //panic: interface conversion: interface {} is float64, not int
	fmt.Println(t3)
	*/
	fmt.Println("---------------------------------------------------------------")

	//3.修改“反射类型对象”
	var circle float64 = 6.28
	vCir := reflect.ValueOf(circle)
	fmt.Println(vCir.CanSet())           // false，这是reflect.value类型，它的值是circle变量值的副本，即使能修改，也是没有意义的
	fmt.Println(vCir)                    // 6.28

	vCir2 := reflect.ValueOf(&circle)
	fmt.Println(vCir2.CanSet())          // false，这是reflect.value类型，它的值是circle变量的地址(指针)，修改也是没有意义的
	fmt.Println(vCir2)                   // 0xc420086008

	vCir3 := reflect.ValueOf(&circle).Elem() //这步操作相当于把指针取消引用，可以直接修改它指向的变量的值
	//vCir3 := reflect.ValueOf(circle).Elem()  //panic: reflect: call of reflect.Value.Elem on float64 Value
	fmt.Println(vCir3.CanSet())          // true
	fmt.Println(vCir3)                   // 6.28

	vCir3.SetFloat(3.14)              //必须使用相应类型的函数，这里为Float，所以用SetFloat
	fmt.Println(circle)                  // 3.14
	fmt.Println("---------------------------------------------------------------")


	//结构体类型
	user := User{"Wang",36,"mile"}
	rType := reflect.TypeOf(user)
	fmt.Println(rType) //main.User
	fmt.Printf("%T,%T,%v,%v\n",user,rType,user,rType) //main.User,*reflect.rtype,{Wang 36 mile},main.User
	fmt.Println("********************************************")

	// rType的value
	fmt.Println(rType)           // main.User
	// rType的value所指向对象的类型
	fmt.Println(rType.Kind())    // struct
	// rType的value所指向对象的名称
	fmt.Println(rType.Name())    // User
	// rType的value所指向对象的全名(包含包)
	fmt.Println(rType.String())  // main.User
	// rType的value所指向对象的字段个数
	fmt.Println(rType.NumField())  // 3
	// rType的value所指向对象对象的方法个数
	fmt.Println(rType.NumMethod())  // 1
	fmt.Println("********************************************")

	rValue := reflect.ValueOf(user)
	fmt.Println(rValue)              // {Wang 36 mile}
	fmt.Println(rValue.Kind())       // struct
	fmt.Println(rValue.String())     // <main.User Value>
	fmt.Println(rValue.NumMethod())  // 1
	fmt.Println(rValue.NumField())   // 3
	fmt.Println("********************************************")
	fmt.Println("---------------------------------------------------------------")

	//修改结构体Field的值
	//1.修改变量一定需要通过Elem()
	//2.可使用FieldByName(fieldName)或者Field(index)
	user1 := User{"小小", 23,"Man"}
	rvalue1 := reflect.ValueOf(&user1) //一定是取得变量user1的指针
	rvalue1.Elem().FieldByName("Sex").SetString("Woman") //Elem()将user1的指针取消引用，取得对应field，使用对应类型函数修改其值
	fmt.Println(user1)                      // {小小 23 Woman}
	rvalue1.Elem().Field(1).SetInt(10)
	fmt.Println(user1)                      // {小小 10 Woman}
	fmt.Println("---------------------------------------------------------------")

	//结构体的Method
	//函数的调用使用`MethodByName()`或者`Method(index)`的call方法
	//返回值是一个数组
	rMsg := reflect.ValueOf(user)
	args := []reflect.Value{reflect.ValueOf("hello 1")}

	method := rMsg.Method(0)
	res := method.Call(args)       // say hello 1

	method1 := rMsg.MethodByName("Say")
	res1 := method1.Call(args)      // say hello 1
	fmt.Println(res)                  // [It's OK.]
	fmt.Println(res[0])               // It's OK.
	fmt.Println(res1)                  // [It's OK.]
	fmt.Println(res1[0])               // It's OK.
	fmt.Println("---------------------------------------------------------------")

	fmt.Println("main end.")
}
