
package main

import (
	"fmt"
	"math/rand"
	"math"
	"gopkg.in/cheggaaa/pb.v1"
)

func main(){
	fmt.Println("Finding value of PI using probability of two numbers being a comprime. running multiple test cases")
	test(10000000,10)
	test(10000000,100)
	test(10000000,1000)
	test(10000000,10000)
	test(10000000,100000)
}
func test(lim int, rnd int){
	var limit int =lim
	var rounds int =rnd
	var count int = 0
	bar := pb.StartNew(rounds)
	for i := 0; i <= rounds; i++ {
		x:= rand.Intn(limit)
		y:= rand.Intn(limit)
		if isComprime(x,y) {
			count++
		}
		bar.Increment()
	}
	bar.FinishPrint("**************************************")
	var pi float64
	// fmt.Println(rounds,count)
	temp:= float64(6*float64(rounds)/float64(count))
	pi = math.Sqrt(temp)
	fmt.Println("value of PI is ",pi)
	fmt.Println("**************************************")
}

func isComprime(x int,y int) bool  {
	a,b := minmax(x,y)
	count:=0
	for i := 1; i <= a; i++ {
		if isFactorOf(i,a) && isFactorOf(i,b) {
			count++
		}
		if count==2 {
			return false
		}
	}
	return true
}
func isFactorOf(i int,a int)  bool{
	if a%i ==0 {
		return true
	}else{
		return false
	}
}
func minmax(x int, y int) (int,int) {
	if x>y {
		return y,x
	}else{
		return x,y
	}
}
