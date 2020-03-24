package utils

import (
	model "ScreenerDataServer/models"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func PostFile(filedir, filename, userID string) (version string, err error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// this step is very important
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		log.Println("error writing to buffer")
		return "", err
	}

	// open file handle
	fh, err := os.Open(filedir)
	if err != nil {
		log.Println("error opening file")
		return "", err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return "", err
	}
	fw, err := bodyWriter.CreateFormField("userID")
	if err != nil {
		log.Println("error opening file")
		return "", err
	}
	_, err = io.Copy(fw, strings.NewReader(userID))
	if err != nil {
		return "", err
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	targetURL := "http://localhost:6002/v1/upload"
	resp, err := http.Post(targetURL, contentType, bodyBuf)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	log.Println(resp.Status)
	log.Println(string(respBody))
	var response map[string]string
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		panic(err)
	}
	return response["versionid"], err
}

func GetDateandTime() (date, timeinhrmin string) {
	dateandtimemap := make(map[string]interface{})
	model.FindOne(bson.M{}, bson.M{"date": 1, "timeinmnhr": 1, "timeadded": 1}, bson.M{"timeadded": -1}).Decode(&dateandtimemap)
	date = dateandtimemap["date"].(string)
	timeinhrmin = dateandtimemap["timeinmnhr"].(string)
	return
}

func GetDateandTimeHourly() (unixtime int64) {
	dateandtimemap := make(map[string]interface{})
	model.FindOne(bson.M{"screener_name": "hourly"}, bson.M{"time_pivot_added": 1}, bson.M{"timeadded": -1}).Decode(&dateandtimemap)
	if len(dateandtimemap) > 0 {
		unixtime = dateandtimemap["time_pivot_added"].(int64)
	}

	return
}
func GetDateandTimeTenMinutes() (unixtime int64) {
	dateandtimemap := make(map[string]interface{})
	model.FindOne(bson.M{"screener_name": "tenminutes"}, bson.M{"time_pivot_added": 1}, bson.M{"time_pivot_added": -1}).Decode(&dateandtimemap)
	unixtime = dateandtimemap["time_pivot_added"].(int64)

	return
}
