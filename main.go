package main

import (
	"fmt"
	"handlegeo/cell"
)

func main() {
	cellArray := cell.Array{}
	a, _ := cell.NewCell(1000000.00, 2000000.00)
	cellArray = append(cellArray, a)
	fmt.Println(*cellArray[0])

	cell.NextCell(a, &cellArray)
	fmt.Println(len(cellArray))
}
