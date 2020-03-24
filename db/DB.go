package db

import (
	model "ScreenerDataServer/models"
	util "ScreenerDataServer/utils"

	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func GetScreenerData() (date, timeinminhr string, result []interface{}, erro string) {

	date, timeinminhr = util.GetDateandTime()
	filter := bson.M{"date": date, "timeinmnhr": timeinminhr}
	cursor, err := model.Find(filter, bson.M{}, bson.M{}, 10000000, 0)
	if err != nil {
		erro = err.Error()
		return
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var data = make(map[string]interface{})
		err = cursor.Decode(&data)

		if err != nil {
			erro = err.Error()
			return
		}

		result = append(result, data)
	}
	return
}

func GetScreenerTenMinuteData() (date, timeinminhr string, result []interface{}, erro string) {

	dateunix := util.GetDateandTimeTenMinutes()
	filter := bson.M{"time_pivot_added": dateunix, "screener_name": "tenminutes"}
	cursor, err := model.Find(filter, bson.M{}, bson.M{}, 10000000, 0)
	if err != nil {
		erro = err.Error()
		return
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var data = make(map[string]interface{})
		err = cursor.Decode(&data)

		if err != nil {
			erro = err.Error()
			return
		}

		result = append(result, data)
	}
	return
}

func GetScreenerHourlyData() (date, timeinminhr string, result []interface{}, erro string) {

	dateunix := util.GetDateandTimeHourly()
	if dateunix != 0 {
		filter := bson.M{"time_pivot_added": dateunix, "screener_name": "hourly"}
		cursor, err := model.Find(filter, bson.M{}, bson.M{}, 10000000, 0)
		if err != nil {
			erro = err.Error()
			return
		}
		defer cursor.Close(context.Background())
		for cursor.Next(context.Background()) {
			var data = make(map[string]interface{})
			err = cursor.Decode(&data)

			if err != nil {
				erro = err.Error()
				return
			}

			result = append(result, data)
		}
	}
	return
}

func GetScreenerDataForDateRange(data model.DataPoints) (result map[string][]interface{}, erro string) {
	filter := bson.M{"date": bson.M{"$in": data.Dates}, "Screener_name": data.ScreenerName}
	cursor, err := model.Find(filter, bson.M{}, bson.M{}, 10000000, 0)
	if err != nil {
		erro = err.Error()
		return
	}
	result = make(map[string][]interface{})
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var data = make(map[string]interface{})
		err = cursor.Decode(&data)

		if err != nil {
			erro = err.Error()
			return
		}

		result[data["date"].(string)] = append(result[data["date"].(string)], data)
	}

	return
}
