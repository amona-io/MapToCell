package cell

import (
	"errors"
	"fmt"
	"handlegeo/naverapi"
	"handlegeo/utils"
	"strconv"
	"strings"
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
	fmt.Println(data)
	return &Cell, nil
}

// NextCell is recursive func and It retrieve Pointer of Cell(Previous Cell) and Array of these Cell.
// Once this func retrieve Pointer of Cell, It finds next cell in the "West Side" of the Previous Cell
// until no more area information are found.
func NextCell(prevCell *cell, cellArray *Array) func() {
	var err error = nil
	centerCoordsArray := strings.Split(prevCell.Center, ",")
	east, north := centerCoordsArray[0], centerCoordsArray[1]
	eastFloat, err := strconv.ParseFloat(east, 64)
	utils.CheckErr(err)
	northFloat, err := strconv.ParseFloat(north, 64)
	utils.CheckErr(err)
	for err == nil {
		nextCell, err := NewCell(eastFloat+100, northFloat)
		if err != nil {
			err = errNoData
			fmt.Println("No more data")
			break
		}
		*cellArray = append(*cellArray, nextCell)
		return NextCell(nextCell, cellArray)
	}
	return nil
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