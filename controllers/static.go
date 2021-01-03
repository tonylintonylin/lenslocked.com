package controllers

import "lenslocked.com/views"

// NewStatic returns static views
func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap", "static/home"),
		Contact: views.NewView("bootstrap", "static/contact"),
	}
}

// Static is static views
type Static struct {
	Home    *views.View
	Contact *views.View
}
