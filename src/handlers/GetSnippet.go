package handlers

import (
	"feyin/go-restapi/models"
	"net/http"

	//"fmt"
	"github.com/go-chi/chi"
	//"encoding/json"
	"github.com/thedevsaddam/renderer"
)

// doesnt need pointer receiver since we are just using the find method here , we are not
// modifying the database
func (h Handler) GetSnippet(w http.ResponseWriter, r *http.Request) {

	// Get the snippet name from the URL parameter
	snippetName := chi.URLParam(r, "snippetName")

	var snippet models.Snippet

	//result := h.DB.First(&snippet, snippetName);
	// result := h.DB.First(&snippet, "snippetname = ?", snippetName)
	result := h.DB.Where("snippet_name = ?", snippetName).First(&snippet)

	if result.Error != nil {
		h.rnd.JSON(w, http.StatusNotFound, renderer.M{
			"message": "Snippet not found",
			"error":   result.Error,
		})
		return
		//fmt.Println(result.Error)
	}

	// h.rnd.JSON(w, http.StatusOK, jsonData)

	// no need for marshalling the renderer marshalls into json automatically for us
	h.rnd.JSON(w, http.StatusOK, renderer.M{
		"data": snippet,
	})

}
