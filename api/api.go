package api

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// ResponseData response data of the nasa api
type ResponseData struct {
	Copyright   string `json:"copyright"`
	Date        string `json:"date"`
	Explanation string `json:"explanation"`
	Title       string `json:"title"`
	URL         string `json:"url"`
}

// GetAPOD gets the Astronomy Picture of the Day (APOD)
func GetAPOD(date string) ResponseData {
	var apiKey = os.Getenv("API_KEY")
	url := "https://api.nasa.gov/planetary/apod?api_key=" + apiKey
	if date != "" {
		url += "&date=" + date
	}

	client := new(http.Client)

	req, err := http.NewRequest("GET", url, nil)
	checkError(err)
	resp, err := client.Do(req)
	checkError(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	var response ResponseData

	json.Unmarshal(body, &response)
	return response
}

// DownloadImage download a image given a image url
func DownloadImage(url string, imgName string) {
	resp, err := http.Get(url)
	checkError(err)
	defer resp.Body.Close()
	file, _ := os.Create(imgName + ".jpg")
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Fatal("error")
	}
}
