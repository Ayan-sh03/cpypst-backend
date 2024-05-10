package domain

import (
	"cpypst/internal/models/generated"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type Paste struct {
	Id         int       `json:"id" validate:"gte=0"`
	Title      string    `json:"title" validate:"required,min=1,max=100"`
	Content    string    `json:"content" validate:"required"`
	UserId     int       `json:"user_id" validate:"gte=0"`
	Slug       string    `json:"slug"`
	Syntax     string    `json:"syntax" validate:"required,oneof=none bash c cpp csharp css dockerfile go html java javascript json kotlin markdown mysql nginx php python ruby rust scala swift typescript xml yaml"`
	IsEditable bool      `json:"is_editable"`
	ExpiresAt  time.Time `json:"expires_at" `
	CreatedAt  time.Time `json:"created_at" validate:"lte"`
}

func (p Paste) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

func ToDomainPaste(p generated.Paste) Paste {

	// var paste Paste

	return Paste{
		Id:         int(p.ID),
		Title:      p.Title.String,
		Content:    p.Content,
		UserId:     int(p.UserID.Int32),
		Slug:       p.Slug,
		Syntax:     p.Syntax.String,
		IsEditable: p.Editable.Bool,
		ExpiresAt:  p.ExpirationTime.Time,
		CreatedAt:  p.CreatedAt.Time,
	}

}
