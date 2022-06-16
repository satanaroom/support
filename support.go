package support

import "errors"

type Support struct {
	Id     int    `json:"id" db:"id"`
	Number int    `json:"number" db:"number"`
	Name   string `json:"name" db:"name"`
	Date   string `json:"date" db:"date"`
}

type UpdateSupportItem struct {
	Number *int    `json:"number" db:"number"`
	Name   *string `json:"name" db:"name"`
	Date   *string `json:"date" db:"date"`
}

func (i UpdateSupportItem) Validate() error {
	if i.Number == nil && i.Name == nil && i.Date == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
