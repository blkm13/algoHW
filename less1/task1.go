package main

import "fmt"

func main (){
	var tRoom, tCond int
	fmt.Scan(&tRoom, &tCond)

	var cMode string
	fmt.Scan(&cMode)

	switch cMode {
	case "freeze":
		if tRoom > tCond {
			tRoom = tCond
		}
	case "heat":
		if tRoom < tCond {
			tRoom = tCond
		}
	case "auto":
		tRoom = tCond
	}

	fmt.Println(tRoom)
	

}
