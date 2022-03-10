package main

import (
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

	// ========[세로운 셀 크롤링]========
	// xMargin = 0.0011317962
	// yMargin = 0.0008983153

	// 1 차 크롤링 (수도권)				가로 708 개	세로 740 개
	//eastStart := 126.541135
	//northStart := 37.104679
	//eastEdge := 127.34131491339947
	//northEdge := 37.76853400669914

	// 2 차 크롤링 (서산 앞바다 ~ 남해)	가로 1068 개  세로 2630 개
	//eastStart := 126.133688368
	//northStart := 34.7430080763
	//eastEdge := 127.34131491339947
	//northEdge := 37.104679

	// 3 차 크롤링 (남해 ~ 제주도)		가로 1068 개  세로 1821 개  // where date(created_at) between '2022-02-01' and '2022-02-25' // 여기까지 입력 완료
	//eastStart := 126.133688368
	//northStart := 33.1080742303
	//eastEdge := 127.34131491339947
	//northEdge := 34.7430080763

	// 4차 크롤링 (가평 ~ 대전)			가로 1988 개 세로 1435 개	// 2022-03-02
	//eastStart := 127.34131491339947
	//northStart := 36.48034986649914
	//eastEdge := 129.59019396279947
	//northEdge := 37.76853400669914

	// 5차 크롤링 (대전 ~ 울산)			가로 1988 개 세로 1101	개	//
	//eastStart := 127.34131491339947
	//northStart := 35.49220303649914
	//eastEdge := 129.59019396279947
	//northEdge := 36.48034986649914

	// 6차 크롤링 (남원 ~ 순천 ~ 부산)	가로 1988 개 세로 601 개	// 2022-03-03 01:24:48.272
	//eastStart := 127.34131491339947
	//northStart := 34.95321385649914
	//eastEdge := 129.59019396279947
	//northEdge := 35.49220303649914

	// 7차 크롤링 (순천 ~ 나로도 ~ 거제)	가로 1396 개 세로 577 개
	//eastStart := 127.34131491339947
	//northStart := 34.43578424369914
	//eastEdge := 128.92017061239947
	//northEdge := 34.95321385649914

	// 8차 크롤링 (연평도 ~ 서산)		가로 361 개 세로 740개		// 여기까지 실서버 입력 완료 됨
	//eastStart := 126.133688368
	//northStart := 37.104679
	//eastEdge := 126.541135
	//northEdge := 37.76853400669914

	// TODO 9차 크롤링 (파주, 의정부, 포천, 철원, 가평, 개성) 가로 923 개 세로 656 개
	//eastStart := 126.541135
	//northStart := 37.76853400669914
	//eastEdge := 127.5846510964
	//northEdge := 38.35693052819914

	// TODO 10차 크롤링 (춘천, 화천, 양구, 인제, 고성, 속초, 양양) 가로 1261개 세로 988f개
	//eastStart := 127.5846510964
	//northStart := 37.76853400669914
	//eastEdge := 129.0107143084
	//northEdge := 38.65517120779914

	//cellArray := cell.Array{} // 입력된 셀 담을 배열

	//첫번째 셀 찾고,
	//firstCell, _ := cell.NewCell(eastStart, northStart)
	//cellArray = append(cellArray, firstCell)
	//err = firstCell.Create(DB)
	//utils.CheckErr(err)

	//첫번째 셀 부터 시작하여 범위 검색
	//cell.NextCell(firstCell, &cellArray, eastEdge, northEdge, DB)
	//fmt.Printf("총 %v 개의 셀이 database에 입력되었습니다.", len(cellArray))
	// ========[새로운 셀 크롤링 코드 종료]========

	// ========[지역 업데이트]========
	// [지역 업데이트 하는 경우 -> 지역 범위 minX minY maxX maxY 입력 후 코드 실행하면 해당 지역 업데이트됨 * 행정동 변경등 있을 시 사용]
	// 좌표 범위로 업데이트
	minX := 126.541135
	minY := 37.104679
	maxX := 127.34131491339947
	maxY := 37.76853400669914
	cells := cell.GetCellsByMinMax(DB, minX, minY, maxX, maxY)

	// area_id로 업데이트
	//cells := cell.GetCellsByAreaId(DB, 125)

	// area_ids로 업데이트
	//cells := cell.GetCellsByAreaIds(DB, []int{70, 83, 498})

	cell.UpdateAllCells(cells, DB)
	// ========지역 업데이트 코드 종료========
}
