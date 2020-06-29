package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Life int
}

type People struct {
	Animal
	Name string //"姓名"
	Age  int    // 年龄
}

func (p People) GetName() string {
	fmt.Println("GetName called:", p.Name)
	name := p.Name
	return name
}

func (p *People) SetName(name string) {
	fmt.Println("SetName called:", name)
	p.Name = name
}

func (p *People) Eat(food string, grams int) (power int) {
	fmt.Println(p.Name, "Eat called", food, grams)
	return 5 * grams
}

func main() {
	p := People{Name: "张三", Age: 20, Animal: Animal{Life: 100}}
	//valueApi(p)
	typeApi(p)
}

/**
访问和修改对象任意属性的值
访问对象的任意方法
 */
func valueApi(o People) {
	oValue := reflect.ValueOf(o)
	fmt.Println(oValue) // {{100}，张三，20}
	fmt.Println("查看张三的原始类型")
	fmt.Println(oValue.Kind())

	fmt.Println("查看p的所有属性值")
	for i := 0; i < oValue.NumField(); i++ {
		fValue := oValue.Field(i)
		fmt.Println(fValue.Interface())
	}
	fmt.Println("获取父类的属性值")
	fatherFieldValue := oValue.FieldByIndex([]int{0, 0})
	fmt.Println(fatherFieldValue.Interface())
	fmt.Println("获得指针values的内容进而获得成员的值")
	opValue := reflect.ValueOf(&o)
	//opValue是一个People指针的Value，oValue是一个People值的Value
	oValue = opValue.Elem()
	fmt.Println(oValue.Field(1).Interface())
	//修改对象的属性
	oValue.FieldByName("Age").SetInt(60)
	fmt.Println(oValue)
	//调用对象的方法
	mValue := opValue.MethodByName("SetName")
	mValue.Call([]reflect.Value{reflect.ValueOf("张三疯")})
	fmt.Println(oValue)

	mValue = opValue.MethodByName("Eat")
	ret := mValue.Call([]reflect.Value{reflect.ValueOf("猫屎咖啡"), reflect.ValueOf(100)})
	fmt.Println("吃完的热量是", ret[0].Interface())
}

//对任意对象细致入微的检测
func typeApi(obj interface{}) {
	oType := reflect.TypeOf(obj)
	fmt.Println(oType)
	//原始类型
	oKind := oType.Kind()
	fmt.Println(oKind)
	//类型名称
	fmt.Println(oType.Name())
	//属性和方法的个数（属于指针的方法不属于实例，属于实例的方法 一定属于指针类型的对象）
	fmt.Println("属性个数：", oType.NumField())
	fmt.Println("方法个数：", oType.NumMethod())

	fmt.Println("全部属性：")
	for i := 0; i < oType.NumField(); i++ {
		structField := oType.Field(i)
		fmt.Println(structField.Name, structField.Type)
	}
	fmt.Println("全部方法：")
	for i := 0; i < oType.NumMethod(); i++ {
		structMethod := oType.Method(i)
		fmt.Println(structMethod.Type,structMethod.Name)
	}

	//获得父类的属性
	fmt.Println(oType.FieldByIndex([]int{0,0}).Name)
}
