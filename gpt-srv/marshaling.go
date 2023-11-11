package proto

import (
	"encoding/json"
)

func (d *DiplomaBackResponse) UnmarshalJSON(data []byte) error {
	type Alias DiplomaBackResponse
	aux := &struct {
		*Alias
		IssueDate string `json:"issue_date"`
	}{
		Alias: (*Alias)(d),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	return nil
}
