package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type thumbnailRequest struct {
	Url string `json:"url"`
}

type ScreenshotApiRequest struct {
	Token string `json:"token"`
	Url string `json:"url"`
	Output string `json:"output"`
	Width int `json:"width"`
	Height int `json:"height"`
	ThumbnailWidth int `json:"thumbnail_width"`
}

func ThumbnailHandler(resWriter http.ResponseWriter, req *http.Request)  {
	var decoded thumbnailRequest
	err := json.NewDecoder(req.Body).Decode(&decoded)
	handleError(err)
	log.Println("decoded search of:", decoded)

	screenRequest := ScreenshotApiRequest{
		Token: "D949342-NCKM4QT-M40JNAC-JG86XAM",
		Url: decoded.Url,
		Output: "json",
		Width: 1920,
		Height: 1080,
		ThumbnailWidth: 300,
	}
	jsonString, err := json.Marshal(screenRequest)
	handleError(err)
	r, err := http.NewRequest("POST", "https://shot.screenshotapi.net/screenshot", bytes.NewBuffer(jsonString))
	handleError(err)
	r.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(r)
	handleError(err)
	defer response.Body.Close()
	type screenshotApiResponse struct{
		Screenshot string `json:"screenshot"`
	}
	var responseScreenshot screenshotApiResponse
	err = json.NewDecoder(response.Body).Decode(&responseScreenshot)
	handleError(err)
	_, err = fmt.Fprintf(resWriter, `{"screenshot": "%s"}`, responseScreenshot.Screenshot)
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
