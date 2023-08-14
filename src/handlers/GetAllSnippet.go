package handlers

import (
	"feyin/go-restapi/models"
	"fmt"
	"net/http"

	//	"fmt"
	//"encoding/json"

	"github.com/thedevsaddam/renderer"
)

// doesnt need pointer receiver since we are just using the find method here , we are not
// modifying the database
func (h Handler) GetAllSnippet(w http.ResponseWriter, r *http.Request) {
	fmt.Print("beginning")
	var Snippets []models.Snippet

	result := h.DB.Find(&Snippets)
	if result.Error != nil {
		//fmt.Println(result.Error)

		h.rnd.JSON(w, http.StatusNotFound, renderer.M{
			"message": "Snippets not found",
			"error":   result.Error,
		})
		return
	}


	// h.rnd.JSON(w, http.StatusOK, jsonData)
// no need for marshalling the renderer marshalls into json automatically for us
	h.rnd.JSON(w, http.StatusOK, renderer.M{
		"data": Snippets,
	})

}
