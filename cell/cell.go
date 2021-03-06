package cell

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	database "handlegeo/db"
	"handlegeo/naverapi"
	"handlegeo/utils"
	"strings"
	"time"
)

type Array []*Cell
type Cell struct {
	LeftTop     string
	RightTop    string
	LeftBottom  string
	RightBottom string
	IsInRange   bool
	CenterCity  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CenterX     float64
	CenterY     float64
}

const xMargin = 0.0011317962
const yMargin = 0.0008983153

var errNoData = errors.New("err No data in this Cell")

func AutoMigrate() {
	db, err := database.Conn()
	utils.CheckErr(err)
	err = db.AutoMigrate(&DBCell{})
	utils.CheckErr(err)
}

// NewCell will retrieve UTM-K coords and return "Pointer of Cell(area centered on these coords)"
// If the "Cell" has not any data, It will return Error
// The "Cell" is an area of 100m in width and 100m in height centered on the input coords
func NewCell(east float64, north float64) (*Cell, error) {
	leftTop := fmt.Sprintf("%.10f,%.10f", east-(xMargin/2), north+(yMargin/2))
	rightTop := fmt.Sprintf("%.10f,%.10f", east+(xMargin/2), north+(yMargin/2))
	leftBottom := fmt.Sprintf("%.10f,%.10f", east-(xMargin/2), north-(yMargin/2))
	rightBottom := fmt.Sprintf("%.10f,%.10f", east+(xMargin/2), north-(yMargin/2))
	Cell := Cell{
		LeftTop:     leftTop,
		RightTop:    rightTop,
		LeftBottom:  leftBottom,
		RightBottom: rightBottom,
		CenterX:     east,
		CenterY:     north,
	}
	// Using coords of center, get data of the Cell
	// If there are no data in Cell, it will return error
	data, statusCode := Cell.getCellDataFromAPI()

	if statusCode == 0 {
		Cell.IsInRange = true
		Cell.CenterCity = data
	} else {
		Cell.IsInRange = false
		Cell.CenterCity = "없음"
	}
	return &Cell, nil
}

func goRoCell(east float64, north float64, c chan<- *Cell) {
	leftTop := fmt.Sprintf("%.10f,%.10f", east-(xMargin/2), north+(yMargin/2))
	rightTop := fmt.Sprintf("%.10f,%.10f", east+(xMargin/2), north+(yMargin/2))
	leftBottom := fmt.Sprintf("%.10f,%.10f", east-(xMargin/2), north-(yMargin/2))
	rightBottom := fmt.Sprintf("%.10f,%.10f", east+(xMargin/2), north-(yMargin/2))
	centerX := east
	centerY := north
	Cell := Cell{
		LeftTop:     leftTop,
		RightTop:    rightTop,
		LeftBottom:  leftBottom,
		RightBottom: rightBottom,
		CenterX:     centerX,
		CenterY:     centerY,
	}
	// Using coords of center, get data of the Cell
	// If there are no data in Cell, it will return error
	data, statusCode := Cell.getCellDataFromAPI()
	fmt.Println(data)

	if statusCode == 0 {
		Cell.IsInRange = true
		Cell.CenterCity = data
		c <- &Cell
	} else {
		Cell.IsInRange = false
		Cell.CenterCity = "없음"
		c <- &Cell
	}
}

func UpdateAllCells(cellArray []DBCell, DB *gorm.DB) {
	now := 1
	all := len(cellArray)
	fmt.Println(all)
	for _, dbCell := range cellArray {
		fmt.Println(now)
		Cell := Cell{
			LeftTop:     dbCell.LeftTop,
			LeftBottom:  dbCell.LeftBottom,
			RightTop:    dbCell.RightTop,
			RightBottom: dbCell.RightBottom,
			CenterX:     dbCell.CenterX,
			CenterY:     dbCell.CenterY,
		}
		data, statusCode := Cell.getCellDataFromAPI()
		if statusCode == 0 {
			dbCell.CenterCity = data
			DB.Save(&dbCell)
		}
		fmt.Printf("진행률 %d / %d\n", now, all)
		now += 1
	}
}

// NextCell is recursive func and It retrieve Pointer of Cell(Previous Cell) and Array of these Cell.
// Once this func retrieve Pointer of Cell, It finds next Cell in the "West Side" of the Previous Cell
// until no more area information are found.
func NextCell(prevCell *Cell, cellArray *Array, eastEdge float64, northEdge float64, DB *gorm.DB) {
	totalTimeout := time.After(18 * 60 * 60 * time.Second)
	east, north := prevCell.CenterX, prevCell.CenterY

	//xMargin := 0.0011317962
	//yMargin := 0.0008983153

	ch := make(chan *Cell)

	//for eastCoord, northCoord := east, north; eastCoord <= eastEdge && northCoord <= northEdge; {
	for eastCoord, northCoord := east, north; true; {
		time.Sleep(time.Microsecond * 10000)
		fmt.Printf("가로좌표: %.10f // 세로좌표: %.10f\n", eastCoord, northCoord)
		if eastCoord >= eastEdge { // eastCoord : 초기화 // northCoord : + 100
			fmt.Println("there")
			if northCoord >= northEdge {
				fmt.Println("종료!!")
				break
			} else {
				eastCoord = east
				go goRoCell(eastCoord, northCoord+yMargin, ch)
				northCoord += yMargin
			}
		} else {
			go goRoCell(eastCoord+xMargin, northCoord, ch)
			eastCoord += xMargin
			fmt.Println(eastCoord, eastEdge)
		}
	}
	turn := 1
	for {
		select {
		case Cell := <-ch:
			*cellArray = append(*cellArray, Cell)
			turn += 1
			err := Cell.Create(DB)
			utils.CheckErr(err)
			fmt.Printf("%v 번째 : %v\n", turn, fmt.Sprintf("%.18f,%.18f", Cell.CenterX, Cell.CenterY))
		case <-totalTimeout: // 채널이 비면 함수 종료
			return
		}
	}
}

// getCellDataFromAPI is private func in package "Cell" and it retrieve Cell struct as argument
// It will return area info and status code.
// If there are any Information of the Cell Area, It returns "Legal Name" of the area in which the Cell is located
// And It returns "Status Code : 0"
// But If there are no data in the Cell Area, It returns empty string as Cell Data
func (c Cell) getCellDataFromAPI() (string, int) {
	coords := utils.JoinFloatToCoord(c.CenterX, c.CenterY)
	cellData := naverapi.RequestAPI(coords)
	statusCode := utils.GetStatusCode(cellData)
	if statusCode != 0 {
		return "No Data", statusCode
	}
	jsonCellData := utils.StringToJSON(cellData)
	areaInfo := getAreaInfo(jsonCellData)
	return areaInfo, 0
}

// getAreaInfo is private func in package "Cell" and it retrieve "JSON" data (in Go we mapping it with "map[string]interface{}")
// This function extract the "Legal Name" from the JSON data as type of string
func getAreaInfo(jsonData map[string]interface{}) string {
	queryData := jsonData["results"].([]interface{})
	rawAreaData := queryData[0].(map[string]interface{})["region"] // 법정동
	areaInfoFirst := rawAreaData.(map[string]interface{})["area1"].(map[string]interface{})["name"]
	areaInfoSecond := rawAreaData.(map[string]interface{})["area2"].(map[string]interface{})["name"]
	areaInfoThird := rawAreaData.(map[string]interface{})["area3"].(map[string]interface{})["name"]
	areaInfoFourth := rawAreaData.(map[string]interface{})["area4"].(map[string]interface{})["name"]
	rawAreaInfo := fmt.Sprintf("%s %s %s %s", areaInfoFirst, areaInfoSecond, areaInfoThird, areaInfoFourth)
	areaInfo := fmt.Sprintf("%s", strings.TrimRight(rawAreaInfo, " "))
	return areaInfo
}
