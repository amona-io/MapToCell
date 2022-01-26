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

	// 서울 전역 *식신 입력 데이터*//
	//eastStart := 126.541135
	//northStart := 37.104679
	//eastEdge := 127.3410408
	//northEdge := 37.7682567

	// 좁은지역 (테스트용)
	eastStart := 126.883335
	northStart := 37.492318
	eastEdge := 126.888742
	northEdge := 37.496591

	cellArray := cell.Array{} // 입력된 셀 담을 배열

	// 첫번째 셀 찾고,
	firstCell, _ := cell.NewCell(eastStart, northStart)
	cellArray = append(cellArray, firstCell)
	err = firstCell.Create(DB)
	utils.CheckErr(err)

	cell.NextCell(firstCell, &cellArray, eastEdge, northEdge, DB)
	fmt.Printf("총 %v 개의 셀이 database에 입력되었습니다.", len(cellArray))
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details

}
