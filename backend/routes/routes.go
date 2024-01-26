package routes

type Route struct {
	Group   string
	Params  string
	Query   string
	Handler string
}

type RouteData struct {
	Routes []Route
}

func GetRoutesData(jsonPath string) *RouteData {
	return nil
}
