package main

import (
	"fmt"
	"handlegeo/cell"
	database "handlegeo/db"
	"handlegeo/utils"
)

func migrateDB() {
	cell.AutoMigrate()
}

func main() {
	// 싱글톤 디비 인스턴스
	DB, err := database.Conn()
	utils.CheckErr(err)
	// 바다와 육지 섞인경우	-> 총 쿼리 수 3599 개 (가로 59 / 세로 61)
	// * 바다의 경우도 해안가인 경우 행정구역 있음
	//eastStart := 888800.00
	//northStart := 1961100.00
	//eastEdge := 894600.00
	//northEdge := 1967100.00

	// 육지의 경우			-> 총 쿼리 수 399 개 (가로 19 / 세로 21)
	//eastStart := 947600.00
	//northStart := 1943400.00
	//eastEdge := 949400.00
	//northEdge := 1945400.00

	eastStart := 935600.00
	northStart := 1936700.00
	eastEdge := 971400.00
	northEdge := 1967900.00

	cellArray := cell.Array{}		// 입력된 셀 담을 배열

	// 첫번째 셀 찾고,
	firstCell, _ := cell.NewCell(eastStart, northStart)
	cellArray = append(cellArray, firstCell)
	err = firstCell.Create(DB)
	utils.CheckErr(err)

	cell.NextCell(firstCell, &cellArray, eastEdge, northEdge)
	fmt.Printf("총 %v 개의 셀이 database에 입력되었습니다.", len(cellArray))
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details

	//result := cell.GetCellsByRange(DB, 947600.00, 1943400.00, 200)
	//fmt.Println(result)
}
