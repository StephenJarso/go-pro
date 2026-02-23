package main

import(
	"fmt"
)
func calcWithTax(price float64)float64{
	return price+(price*0.2)
}
func calcWithoutTax(price float64)float64{
	return price
}
//using function as arguments
func printPrice(product string,price float64,calculator func(float64) float64){
	fmt.Println("product:",product,"price:",calculator(price))
}
func main(){
	products:=map[string]float64{
		"Kayak":275,
		"Lifejacket":48.95,
	}
	for product,price:=range products{
		var calcFunc func(float64) float64//function data type
		fmt.Println("Function assigned:",calcFunc==nil)//check if a function has been assigned
		if (price>100){
			calcFunc = calcWithTax
			printPrice(product,price,calcWithTax)
		} else{
			calcFunc=calcWithoutTax
			printPrice(product,price,calcWithoutTax)
		}
		totalPrice:=calcFunc(price)
		fmt.Println("function assigned:",calcFunc==nil)//check if a function has been assigned
		fmt.Println("product",product,"Price",totalPrice)
	}
}