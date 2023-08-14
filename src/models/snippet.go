package models


import(
	"time"
)


type Snippet struct {
	Id          uint    `json:"id" gorm:"primaryKey"`
	SnippetName string    `json:"snippetname"`
	Code        string    `json:"code"`
	CreatedAt   time.Time `json:"created_at"`
}