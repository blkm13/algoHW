package main

import "fmt"

func main()  {
	var a,b, c = 4, 5, 7
	fmt.Scan(&a,&b,&c)


	if c < 0 {
		fmt.Println("NO SOLUTION")
	}else {
		if a == 0{
			if c*c - b == 0 {
				fmt.Println("MANY ёSOLUTION")
			}else {
				fmt.Println("NO SOLUTION")
			}
		}else if (c*c-b)%a == 0 {
			fmt.Println((c*c-b)/a)
		}else {
			fmt.Println("NO SOLUTION")
		}
	}


}
