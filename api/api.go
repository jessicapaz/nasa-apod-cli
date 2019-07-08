package api

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// ResponseData response data of the nasa api
type ResponseData struct {
	Copyright   string `json:"copyright"`
	Date        string `json:"date"`
	Explanation string `json:"explanation"`
	Title       string `json:"title"`
	URL         string `json:"url"`
}

// DateRange range of two dates
type DateRange struct {
	Start string
	End   string
}

var apiKey = os.Getenv("API_KEY")

// GetAPOD gets the Astronomy Picture of the Day (APOD)
func GetAPOD(date string) ResponseData {
	url := "https://api.nasa.gov/planetary/apod?api_key=" + apiKey
	if date != "" {
		url += "&date=" + date
	}
	resp, err := http.Get(url)
	checkError(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	var res ResponseData
	err = json.Unmarshal(body, &res)
	checkError(err)
	return res
}

// GetAPODs gets the Astronomy Picture of the Day (APOD) for a range of dates
func GetAPODs(dt DateRange) []ResponseData {

	startDt := stringToDate(dt.Start)
	endDt := stringToDate(dt.End)

	var r []ResponseData
	for startDt.Before(endDt) == true || startDt == endDt {
		tmp := startDt.Format("2006-01-02")
		url := "https://api.nasa.gov/planetary/apod?api_key=" + apiKey
		url += "&date=" + tmp
		ch := make(chan []byte)
		go makeRequest(url, ch)
		var res ResponseData
		err := json.Unmarshal(<-ch, &res)
		checkError(err)
		r = append(r, res)
		startDt = startDt.AddDate(0, 0, 1)
	}
	return r
}

func makeRequest(url string, ch chan<- []byte) {
	resp, err := http.Get(url)
	checkError(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	ch <- body
}

// DownloadImage download an image given an url
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

func stringToDate(s string) time.Time {
	layout := "2006-01-02"
	t, _ := time.Parse(layout, s)
	return t
}
