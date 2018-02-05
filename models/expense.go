package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
)

type Expense struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Name      string    `json:"name" db:"name"`
	Value     int       `json:"value" db:"value"`
	Paid      bool      `json:"paid" db:"paid"`
	BoardID   uuid.UUID `json:"board_id" db:"board_id"`
}

// String is not required by pop and may be deleted
func (e Expense) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Expenses is not required by pop and may be deleted
type Expenses []Expense

// String is not required by pop and may be deleted
func (e Expenses) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (e *Expense) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: e.Name, Name: "Name"},
		&validators.IntIsPresent{Field: e.Value, Name: "Value"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (e *Expense) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (e *Expense) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (e Expenses) GetTotal() int {
	value := 0

	for _, element := range e {
		value += element.Value
	}

	return value
}

func (e Expenses) GetTotalPaid() int {
	value := 0
	for _, element := range e {
		if element.Paid {
			value += element.Value
		}
	}

	return value
}
