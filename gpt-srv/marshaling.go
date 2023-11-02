package proto

import (
	"encoding/json"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func (d *DiplomaFrontResponse) UnmarshalJSON(data []byte) error {
	type Alias DiplomaFrontResponse
	aux := &struct {
		*Alias
		IssueDate string `json:"issue_date"`
	}{
		Alias: (*Alias)(d),
	}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	issueDate, err := time.Parse("2006-01-02", aux.IssueDate)
	if err != nil {
		return err
	}

	d.IssueDate = timestamppb.New(issueDate)
	return nil
}

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

	issueDate, err := time.Parse("2006-01-02", aux.IssueDate)
	if err != nil {
		return err
	}

	d.IssueDate = timestamppb.New(issueDate)
	return nil
}
