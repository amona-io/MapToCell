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
	// 디비 인스턴스
	DB, err := database.Conn()
	utils.CheckErr(err)

	// [세로운 셀 크롤링]
	// xMargin = 0.0011317962
	// yMargin = 0.0008983153

	// 1 차 크롤링 (수도권)
	//eastStart := 126.541135
	//northStart := 37.104679
	//eastEdge := 127.34131491339947
	//northEdge := 37.76853400669914

	// 2 차 크롤링 (서산 앞바다 ~ 남해)	가로 1067 개  세로 2629 개
	//eastStart := 126.133688368
	//northStart := 34.7430080763
	//eastEdge := 127.34131491339947
	//northEdge := 37.104679

	// 3 차 크롤링 (남해 ~ 제주도)		가로 1067 개  세로 1820 개  // where date(created_at) between '2022-02-01' and '2022-02-25'
	//eastStart := 126.133688368
	//northStart := 33.1080742303
	//eastEdge := 127.34131491339947
	//northEdge := 34.7430080763

	// 4차 크롤링 (가평 ~ 대전)			가로 1987 개 세로 1434 개	//
	eastStart := 127.34131491339947
	northStart := 36.48034986649914
	eastEdge := 129.59019396279947
	northEdge := 37.76853400669914

	cellArray := cell.Array{} // 입력된 셀 담을 배열

	// 첫번째 셀 찾고,
	firstCell, _ := cell.NewCell(eastStart, northStart)
	cellArray = append(cellArray, firstCell)
	err = firstCell.Create(DB)
	utils.CheckErr(err)

	cell.NextCell(firstCell, &cellArray, eastEdge, northEdge, DB)
	fmt.Printf("총 %v 개의 셀이 database에 입력되었습니다.", len(cellArray))

	// ========새로운 셀 크롤링 코드 종료=========

	// [지역 업데이트 하는 경우 -> 지역 범위 minX minY maxX maxY 입력 후 코드 실행하면 해당 지역 업데이트됨 * 행정동 변경등 있을 시 사용]
	//minX := 126.749949
	//minY := 37.421981
	//maxX := 127.245712
	//maxY := 37.708212
	//
	//cells := cell.GetCellsByMinMax(DB, minX, minY, maxX, maxY)
	//cell.UpdateAllCells(cells, DB)

	// =========지역 업데이트 코드 종료=========
}
