package cell

import (
	"errors"
	"fmt"
	"handlegeo/naverapi"
	"handlegeo/utils"
	"strconv"
	"strings"
	"time"
)

type Array []*cell
type cell struct {
	CellNum		int
	LeftTop		string
	RightTop	string
	LeftBottom	string
	RightBottom	string
	Center		string
	IsInRange	bool
	CenterCity	string
}
var errNoData = errors.New("err No data in this cell")

// NewCell will retrieve UTM-K coords and return "Pointer of Cell(area centered on these coords)"
// If the "Cell" has not any data, It will return Error
// The "Cell" is an area of 100m in width and 100m in height centered on the input coords
func NewCell(east float64, north float64) (*cell, error) {
	center := fmt.Sprintf("%.2f,%.2f",east, north)
	leftTop := fmt.Sprintf("%.2f,%.2f",east-50.00,north+50.00)
	rightTop := fmt.Sprintf("%.2f,%.2f",east+50.00,north+50.00)
	leftBottom := fmt.Sprintf("%.2f,%.2f",east-50.00,north-50.00)
	rightBottom := fmt.Sprintf("%.2f,%.2f",east+50.00,north-50.00)
	Cell := cell{
		Center		: center,
		LeftTop		: leftTop,
		RightTop	: rightTop,
		LeftBottom	: leftBottom,
		RightBottom	: rightBottom,
	}
	// Using coords of center, get data of the Cell
	// If there are no data in Cell, it will return error
	data, statusCode := Cell.getCellData()

	if statusCode != 0 {
		return nil, errNoData
	}
	Cell.IsInRange = true
	Cell.CenterCity = data
	return &Cell, nil
}

func GoRoCell(east float64, north float64, c chan <- *cell)  {
	center := fmt.Sprintf("%.2f,%.2f",east, north)
	leftTop := fmt.Sprintf("%.2f,%.2f",east-50.00,north+50.00)
	rightTop := fmt.Sprintf("%.2f,%.2f",east+50.00,north+50.00)
	leftBottom := fmt.Sprintf("%.2f,%.2f",east-50.00,north-50.00)
	rightBottom := fmt.Sprintf("%.2f,%.2f",east+50.00,north-50.00)
	Cell := cell{
		Center		: center,
		LeftTop		: leftTop,
		RightTop	: rightTop,
		LeftBottom	: leftBottom,
		RightBottom	: rightBottom,
	}
	// Using coords of center, get data of the Cell
	// If there are no data in Cell, it will return error
	data, statusCode := Cell.getCellData()

	if statusCode == 0 {
		Cell.IsInRange = true
		Cell.CenterCity = data
		c <- &Cell
	}
}


// NextCell is recursive func and It retrieve Pointer of Cell(Previous Cell) and Array of these Cell.
// Once this func retrieve Pointer of Cell, It finds next cell in the "West Side" of the Previous Cell
// until no more area information are found.
func NextCell(prevCell *cell, cellArray *Array) {
	//var err error = nil
	centerCoordsArray := strings.Split(prevCell.Center, ",")

	east, north := centerCoordsArray[0], centerCoordsArray[1]
	eastFloat, err := strconv.ParseFloat(east, 64)
	utils.CheckErr(err)
	northFloat, err := strconv.ParseFloat(north, 64)
	utils.CheckErr(err)

	eastEdge := 960000.00
	northEdge := 1960000.00
	margin := 100.00

	ch := make(chan *cell)
	length := 0

	// 재귀 > 범위 지정하고, 그 범위를 도는 걸로 함수 변경  (테스트 : 950000, 1950000 ~ 900000, 1900000)
	// Next Cell 함수 내에서 채널 만들고, (함수명 알맞게 변경)
	// 채널에 쿼리 결과를 담는 식으로 New Cell 과 유사한 함수 생성
	// 본 함수에서 만든 채널에 > 셀 정보 불러오는 함수로 채널을 인자로 넘겨줘서 그 채널에 데이터 담고

	for eastCoord, northCoord := eastFloat, northFloat; eastCoord <= eastEdge && northCoord <= northEdge; {
		time.Sleep(time.Microsecond*350)
		fmt.Printf("가로좌표: %.2f // 세로좌표: %.2f\n", eastCoord, northCoord)
		length += 1
		if eastCoord == eastEdge {	// eastCoord : 초기화 // northCoord : + 100
			if northCoord == northEdge {
				break
			}
			go GoRoCell(eastCoord, northCoord+margin, ch)
			eastCoord = 950000
			northCoord += margin

		} else {
			go GoRoCell(eastCoord+margin, northCoord, ch)
			eastCoord += margin
		}
	}

	for i:=1 ; i < length; i++ {
		a := <- ch
		*cellArray = append(*cellArray, a)
	}
}

// getCellData is private func in package "cell" and it retrieve cell struct as argument
// It will return area info and status code.
// If there are any Information of the Cell Area, It returns "Legal Name" of the area in which the Cell is located
// And It returns "Status Code : 0"
// But If there are no data in the Cell Area, It returns empty string as Cell Data
func (c cell) getCellData() (string, int) {
	coords := c.Center
	cellData := naverapi.RequestAPI(coords)
	statusCode := utils.GetStatusCode(cellData)
	if statusCode != 0 {
		return "No Data", statusCode
	}
	jsonCellData := utils.StringToJSON(cellData)
	areaInfo := getAreaInfo(jsonCellData)
	return areaInfo, 0
}

// getAreaInfo is private func in package "cell" and it retrieve "JSON" data (in Go we mapping it with "map[string]interface{}")
// This function extract the "Legal Name" from the JSON data as type of string
func getAreaInfo(jsonData map[string]interface{}) string {
	queryData := jsonData["results"].([]interface{})
	rawAreaData := queryData[0].(map[string]interface{})["region"]	// 법정동
	areaInfoFirst := rawAreaData.(map[string]interface{})["area1"].(map[string]interface{})["name"]
	areaInfoSecond := rawAreaData.(map[string]interface{})["area2"].(map[string]interface{})["name"]
	areaInfoThird := rawAreaData.(map[string]interface{})["area3"].(map[string]interface{})["name"]
	areaInfoFourth := rawAreaData.(map[string]interface{})["area4"].(map[string]interface{})["name"]
	areaInfo := fmt.Sprintf("%s %s %s %s", areaInfoFirst, areaInfoSecond, areaInfoThird, areaInfoFourth)
	return areaInfo
}