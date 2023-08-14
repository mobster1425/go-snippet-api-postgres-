package handlers

import (
	"encoding/json"
	"feyin/go-restapi/models"
	"net/http"

	//"fmt"

	"strings"

	"github.com/go-chi/chi"
	"github.com/thedevsaddam/renderer"
)

// needs pointer receiver since we are we are modifying the database
func (h *Handler) UpdateSnippet(w http.ResponseWriter, r *http.Request) {
	idstr := strings.TrimSpace(chi.URLParam(r, "codeid"))

	var updatedSnippets models.Snippet
	// storing the data recideved that would be used for uopdates into this var
	if err := json.NewDecoder(r.Body).Decode(&updatedSnippets); err != nil {
		h.rnd.JSON(w, http.StatusProcessing, err)
		return
	}

	// validating input

	if updatedSnippets.Code == "" && updatedSnippets.SnippetName == "" {
		h.rnd.JSON(w, http.StatusBadRequest, renderer.M{
			"message1": "the code input field is requested",
			"message2": "the snippet Name field is requested",
		})
		// return from func no need to continue execution of func
		return
	}

	var snippet models.Snippet
	// storing the data received from the database here
	// this is the data that will be updated
	if result := h.DB.First(&snippet, "id = ?", idstr); result.Error != nil {
		//fmt.Println(result.Error)

		h.rnd.JSON(w, http.StatusNotFound, renderer.M{
			"message": "Snippet not found",
			"error":   result.Error,
		})
		return
	}

	// updating the snippet found with the data from the frontend
	snippet.Code = updatedSnippets.Code
	snippet.SnippetName = updatedSnippets.SnippetName

	//saving the changes to the database
	h.DB.Save(&snippet)

	// returning data to the frontend
	h.rnd.JSON(w, http.StatusOK, renderer.M{
		"message": "Snippet updated successfully",
	})

}
