package main

import (
	"fmt"
)

const (
	KB = 1000
	MB = KB*1000
	GB = MB*1000
	TB = GB*1000
	PB = TB*1000
	EB = PB*1000
	ZB = EB*1000
	YB = ZB*1000
)

func main()  {
	fmt.Println("KB:", KB)
	fmt.Println("MB:",MB)
	fmt.Println("GB:",GB)
	fmt.Println("TB:",TB)
	fmt.Println("EB:",EB)
	//overflow
	//fmt.Println("ZB:",big.NewInt(ZB))
	//fmt.Println("YB:",big.NewInt(YB))
}