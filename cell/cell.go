package cell

import (
	"fmt"
	"handlegeo/naverapi"
)

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

type cellArray []cell

func NewCell(east float64, north float64) *cell {
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
	Cell.getCellData()
	return &Cell
}

func NextCell(prevCell cell) {
	return
}

func (c cell) getCellData() {
	coords := c.Center
	cellData := naverapi.RequestAPI(coords)
	fmt.Println(cellData)
	naverapi.GetStatusCode(cellData)
}