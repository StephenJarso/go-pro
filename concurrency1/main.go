package main

import (
	"fmt"
	"time"
)

func main() {
	startNow:=time.Now()
	fmt.Println("main function started")
	CalcStoreTotal(Products)
	fmt.Println("main function complete")
	fmt.Println("This operation took:",time.Since(startNow))
}
