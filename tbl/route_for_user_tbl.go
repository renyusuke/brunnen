package tbl

type RouteForUser struct {
	id       int64
	title    string
	path     string
	children []RouteForAdmin
}
