package handler

import (
	db "ScreenerDataServer/db"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetScreenerData(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	var err2 string
	date, timeinhrmin, datasint, err2 := db.GetScreenerData()

	if err2 != "" {
		json.NewEncoder(w).Encode(bson.M{"success": false, "error": err2})
		return
	}

	if len(datasint) == 0 {
		var arr = make([]string, 0)
		json.NewEncoder(w).Encode(bson.M{"success": true, "error": err2, "data": arr})

	} else {
		json.NewEncoder(w).Encode(bson.M{"success": true, "error": err2, "data": datasint, "date": date, "time": timeinhrmin})
	}
}

func GetScreenerHourlyData(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	var err2 string
	date, timeinhrmin, datasint, err2 := db.GetScreenerHourlyData()

	if err2 != "" {
		json.NewEncoder(w).Encode(bson.M{"success": false, "error": err2})
		return
	}

	if len(datasint) == 0 {
		var arr = make([]string, 0)
		json.NewEncoder(w).Encode(bson.M{"success": true, "error": err2, "data": arr})

	} else {
		json.NewEncoder(w).Encode(bson.M{"success": true, "error": err2, "data": datasint, "date": date, "time": timeinhrmin})
	}
}
func GetScreenerTenMinuteData(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	var err2 string
	date, timeinhrmin, datasint, err2 := db.GetScreenerTenMinuteData()

	if err2 != "" {
		json.NewEncoder(w).Encode(bson.M{"success": false, "error": err2})
		return
	}

	if len(datasint) == 0 {
		var arr = make([]string, 0)
		json.NewEncoder(w).Encode(bson.M{"success": true, "error": err2, "data": arr})

	} else {
		json.NewEncoder(w).Encode(bson.M{"success": true, "error": err2, "data": datasint, "date": date, "time": timeinhrmin})
	}
}
