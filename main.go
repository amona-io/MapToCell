package main

import (
	"fmt"
	"handlegeo/cell"
	database "handlegeo/db"
)

func main() {
	DB := database.DBConn()
	database.MigrateDB(DB)

	// 바다와 육지 섞인경우
	//eastStart := 888800.00
	//northStart := 1961100.00
	//eastEdge := 894600.00
	//northEdge := 1967100.00

	eastStart := 947600.00
	northStart := 1943400.00
	eastEdge := 949400.00
	northEdge := 1945400.00

	cellArray := cell.Array{}

	//a, _ := cell.NewCell(950000, 1950000)  //950000, 1950000
	firstCell, _ := cell.NewCell(eastStart, northStart)  //950000, 1950000
	cellArray = append(cellArray, firstCell)

	cell.NextCell(firstCell, &cellArray, eastEdge, northEdge)
	fmt.Printf("총 %v 개의 셀이 database에 입력되었습니다.", len(cellArray))
}
