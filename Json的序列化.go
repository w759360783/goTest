package main

//type Person struct {
//	Name  string
//	Age   int
//	Rmb   float64
//	Sex   bool
//	Hobby []string
//}
//
////结构体转json
//func main001() {
//	person := Person{Name: "于谦", Age: 50, Rmb: 12345.36, Sex: true, Hobby: []string{"抽烟", "喝酒", "烫头"}}
//	bytes, err := json.Marshal(person)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(string(bytes))
//}
//
////映射转json
//func main002() {
//	dataMap := make(map[string]interface{})
//	dataMap["Name"] = "于谦"
//	dataMap["Age"] = 50
//	dataMap["Rmb"] = 12345.36
//	dataMap["Sex"] = true
//	dataMap["Hobby"] = []string{"抽烟", "喝酒", "烫头"}
//
//	bytes, err := json.Marshal(dataMap)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(string(bytes))
//}
//
////切片转json
//func main003(){
//	dataSlice :=make([]map[string] interface{},0)
//	dataMap := make(map[string]interface{})
//	dataMap["Name"] = "于谦"
//	dataMap["Hobby"] = []string{"抽大烟", "喝白酒", "烫黄头"}
//
//	dataMap1 := make(map[string]interface{})
//	dataMap1["Name"] = "于大谦"
//	dataMap1["Hobby"] = []string{"抽小烟", "喝啤酒", "烫给头"}
//
//	dataMap2 := make(map[string]interface{})
//	dataMap2["Name"] = "于小谦"
//	dataMap2["Hobby"] = []string{"抽像烟", "喝黄酒", "烫绿头"}
//
//	dataSlice = append(dataSlice,dataMap,dataMap1,dataMap2)
//
//	bytes, err := json.Marshal(dataSlice)
//
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(string(bytes))
//}
