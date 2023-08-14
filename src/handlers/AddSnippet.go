package handlers

import (
	"encoding/json"
	"feyin/go-restapi/models"
	"fmt"
	"net/http"
	"time"

	"github.com/thedevsaddam/renderer"
)

// needs pointer receiver since we are we are modifying the database
func (h Handler) AddSnippet(w http.ResponseWriter, r *http.Request) {

	// creating an instance of Snippet that would be use for decoding into struct
	//which will also be used to store the data into the database
	var snippet models.Snippet

	if err := json.NewDecoder(r.Body).Decode(&snippet); err != nil {
		// remember rnd pckg is used to send json response and it collects 3 values
		// the writer intance, status code, message to be sent , the mssg is a map data structure
		// here we are sending the err
		h.rnd.JSON(w, http.StatusProcessing, err)
		// we are returning from the function since we are getting an err while decoding
		//so theres no need to conti ue the execution of the function
		return
	}

	// Set CreatedAt to the current time
	snippet.CreatedAt = time.Now()

	// validating input

	if snippet.Code == "" && snippet.SnippetName == "" {
		h.rnd.JSON(w, http.StatusBadRequest, renderer.M{
			"message1": "the code input field is requested",
			"message2": "the snippet Name field is requested",
		})
		// return from func no need to continue execution of func
		return
	}
	// Append to the Snippet table
	result := h.DB.Create(&snippet)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	// returning the inserted id  as json response

	h.rnd.JSON(w, http.StatusCreated, renderer.M{
		"message": "Snippet created successfully",
		//"snippet_id": cm.ID.Hex(),
		"snippet Added": result.RowsAffected,
	})

}
