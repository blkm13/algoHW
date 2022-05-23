package main

import (
	"fmt"
	"strings"
)

func main()  {
	var phoneNum string
	fmt.Scan(&phoneNum)

	phoneNum = strings.Replace(phoneNum, "-", "", -1)
	phoneNum = strings.Replace(phoneNum, "+7","8", 1)
	phoneNum = strings.Replace(phoneNum,"(","",1)
	phoneNum = strings.Replace(phoneNum, ")","",1)

	if len(phoneNum) == 7{
		phoneNum = "8495"+phoneNum
	}
	//fmt.Println(phoneNum)

	for i := 0; i < 3; i++{
		var newNum string
		fmt.Scan(&newNum)

		newNum = strings.Replace(newNum, "-", "", -1)
		newNum = strings.Replace(newNum, "+7","8", 1)
		newNum = strings.Replace(newNum,"(","",1)
		newNum = strings.Replace(newNum, ")","",1)

		if len(newNum) == 7{
			newNum = "8495"+newNum
		}

		//fmt.Println(newNum)

		if newNum == phoneNum {
			fmt.Println("YES")
		}else{
			fmt.Println("NO")
		}
	}

}
