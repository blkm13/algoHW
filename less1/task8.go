package main

import "fmt"

func find(a int, b int, n int, m int){
	var tamin,tamax,tbmin,tbmax int

	if n == 0 && m != 0 {
		fmt.Println("-1")
	}else if n ==0 && m ==0 {
		//какой-то ответ, наверное минимум она могла стоять 0 а максимум а минут
	}else if n !=0 {
		tamin = n +a*(n-1)
		tamax = tamin + 2*a

		tbmin = m + b*(m-1)
		tbmax = tbmin + 2*b

		if tamax <
	}
}

func main ()  {
	var a, b, n, m int
	fmt.Scan(&a,&b,&n, &m)

	if a<b{
		find(a,b,n,m)
	}else {
		find(b,a,m,n)
	}
}
