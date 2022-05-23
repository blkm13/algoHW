package main

import (
	"fmt"
	"sort"
)

func main(){
	var d, e int
	a:= make([]int,3)

	fmt.Scan(&a[0], &a[1], &a[2])
	//fmt.Println(a)
	fmt.Scan(&d, &e)
	sort.Ints(a)
	if d < e{
		k := d
		d = e
		e = k
	}
	if a[0] <= e && a[1] <= d{
		fmt.Println("YES")
	}else {
		fmt.Println("NO")
	}


}
