package main

import "fmt"

// расстановка ноутбуков


func main (){

	var a1, b1, a2 ,b2 int
	fmt.Scan(&a1, &b1, &a2, &b2)



	if a1<b1 {
		c := a1
		a1 = b1
		b1 = c
	}

	if a2 < b2{
		c:= a2
		a2 = b2
		b2 = c
	}



	var mina ,minb int

	if a1 > a2{
		if b1 > a2 {
			mina = a1+b2
			minb = a2
		}else {
			mina = b1 + b2
			minb = a1
		}
	}else if b2 > a1 {
		mina= a2+b1
		minb = a1
	}else {
		mina = b1+b2
		minb = a2
	}




	fmt.Println(mina, minb)

}