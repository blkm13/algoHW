package main

import "fmt"

func main(){
	var k1, k2, n2, p2, m int
	fmt.Scan(&k1, &m, &k2, &p2, &n2 )

	var kCount int //кол-во квартир на этаже

	if k2 % (m*(p2-1)+n2) ==0{
		kCount = k2 /( m*(p2-1)+ n2)
	} else {
		kCount = k2 /( m*(p2-1)+ n2) +1
	}

	var k, n1, p1 int

	if k2 > m*(p2-1){
		for kCount*m*(p2-1)+n2-1 < k2{
			if k1 == kCount{
				break
			}
			k+=1

			p1 = k1/m*kCount +1
			n1 = (k1 - (p1-1)*m*kCount)/kCount +1

			kCount +=1

			if k == 1{
				resN1 := n1
				resP1 := p1
			} else{
				if resN1 != 0{

				}
			}


		}
	}

}