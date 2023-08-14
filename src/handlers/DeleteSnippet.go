package handlers

import (
	"feyin/go-restapi/models"
	"net/http"

	//	"fmt"

	"github.com/thedevsaddam/renderer"

	"strings"

	"github.com/go-chi/chi"
)

// needs pointer receiver since we are we are modifying the database
func (h *Handler) DeleteSnippet(w http.ResponseWriter, r *http.Request) {

	// getting the id of the snippet code that wants to deleted
	idstr := strings.TrimSpace(chi.URLParam(r, "id"))

	var snippet models.Snippet

	// find the snippet first
	result := h.DB.First(&snippet, "id = ?", idstr)
	if result.Error != nil {
		//fmt.Println(result.Error)

		h.rnd.JSON(w, http.StatusNotFound, renderer.M{
			"message": "Snippet not found",
			"error":   result.Error,
		})
		return
	}

	// Delete that snippet
	h.DB.Delete(&snippet)

	h.rnd.JSON(w, http.StatusOK, renderer.M{
		"message": "Code Snippet deleted successfully",
	})

}
