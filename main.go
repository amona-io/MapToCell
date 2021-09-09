package main

import "handlegeo/cell"

func main() {
	//coords := "127.105399,37.3595704"
	//coords := "1000000,1700000"				// UTM-K (동-서,북-남)
	//apiCall := naverapi.RequestAPI(coords)
	//fmt.Println(apiCall)

	cell.NewCell(1000000.00, 2000000.00)
	cell.NewCell(000000.00, 2000000.00)
}
