package main

import (
	"fmt"
	"handlegeo/cell"
)

func main() {
	cellArray := cell.Array{}
	a, _ := cell.NewCell(950000, 1950000)  //950000, 1950000
	cellArray = append(cellArray, a)

	cell.NextCell(a, &cellArray)
	fmt.Println(len(cellArray))
}
