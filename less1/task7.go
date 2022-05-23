package main

import "fmt"

func main()  {
	var n,k,m int

	fmt.Scan(&n,&k, &m)

	var mCount = 0

	if k>m {

		for n >= k {

			n -= k
			mCount += k / m
			n += k % m
		}
	}


	fmt.Println(mCount)
}
