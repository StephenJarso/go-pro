package main

import(
	"fmt"
	"sort"
	"reflect"
)

func main(){
	products:=map[string]float64{"Kayak":40.9,
	"Lifejacket":345,"Paddle":234,"Hat":2432,}
	fmt.Println(products)
value,ok:=products["Hat"]
if ok{
	fmt.Println("stored value:",value)
} else {
	fmt.Println("no stored value")
}
	p1:=products
	keys:=make([]string,len(products))
	for key,value := range products{
		keys = append(keys,key)
		fmt.Println(key ,":", value)
}
sort.Strings(keys)
for _,value := range keys{
		fmt.Println(value ,":", products[value])
}
fmt.Println("Equal:",reflect.DeepEqual(products,p1))
}