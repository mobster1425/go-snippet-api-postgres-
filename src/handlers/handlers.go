package handlers

import ("gorm.io/gorm"
"github.com/thedevsaddam/renderer"

)

type Handler struct {
	DB *gorm.DB
	rnd *renderer.Render
}



/*
Handler New Function (New(db *gorm.DB ,rnd *renderer.Render) Handler):

A constructor function for creating a new handler instance.
*/
func New(db *gorm.DB ,rnd *renderer.Render) *Handler {
	return &Handler{db,rnd}
}