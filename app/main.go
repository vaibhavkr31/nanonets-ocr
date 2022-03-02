package main

import (
	"github.com/dripcapital/nanonets-ocr/app/nanonets"
	"github.com/dripcapital/nanonets-ocr/app/services"
	st "github.com/dripcapital/nanonets-ocr/app/structs"
	"github.com/dripcapital/nanonets-ocr/app/utils"
)

func main() {

	filePath := utils.SaveImageLocally("FileUrl")
	resp, err := nanonets.PredictViaImageFile(filePath)
	if err != nil {
		return
	}
	bldocResp := st.BlDocResponse{RawResponse: resp}
	services.StoreNanonetResponse(bldocResp)
}
