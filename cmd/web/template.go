package main

import 	dbmodel "serrver/internal"

type templateData struct {
	Blog *dbmodel.Blog
	Blogs []*dbmodel.Blog
}
