package router

import (
	"net/http"

	handler "ScreenerDataServer/handlers"
)

// Route type description
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes contains all routes
type Routes []Route

var routes = Routes{

	Route{
		"getscreenerdata",
		"POST",
		"/getscreenerdata",
		handler.GetScreenerData,
	},
	Route{
		"getscreenertenminutedata",
		"POST",
		"/getscreenertenminutedata",
		handler.GetScreenerTenMinuteData,
	},
	Route{
		"getscreenerhourlydata",
		"POST",
		"/getscreenerhourlydata",
		handler.GetScreenerHourlyData,
	},
}
