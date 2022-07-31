package main

//
//import "fmt"
//
//func main() {
//	m := 10
//	defer fmt.Println("first defer", m)
//	m = 100
//	defer func() {
//		fmt.Println("second defer", m)
//	}()
//	m *= 10
//	defer fmt.Println("third defer", m)
//	funcVal := func1()
//	funcVal()
//	m *= 10
//}
//
//func func1() func() {
//	defer fmt.Println("111")
//	return func() {
//		defer fmt.Println("222")
//	}
//}
