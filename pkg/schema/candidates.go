package schema

import (
	"encoding/json"
	"time"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Candidate struct {
	Id          string `json:"id"`
	Name        string `json:"name,omitempty"`
	Firstname   string `json:"firstname,omitempty"`
	Lastname    string `json:"lastname,omitempty"`
	Headline    string `json:"headline,omitempty"`
	Summary     string `json:"summary,omitempty"`
	Address     string `json:"address,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Email       string `json:"email,omitempty"`
	CoverLetter string `json:"cover_letter,omitempty"`
	Education   []struct {
		Degree       string `json:"degree,omitempty"`
		School       string `json:"school,omitempty"`
		FieldOfStudy string `json:"field_of_study,omitempty"`
		StartDate    string `json:"start_date,omitempty"`
		EndDate      string `json:"end_date,omitempty"`
	} `json:"education_entries,omitempty"`
	Experience []struct {
		Title     string `json:"title,omitempty"`
		Summary   string `json:"summary,omitempty"`
		StartDate string `json:"start_date,omitempty"`
		EndDate   string `json:"end_date,omitempty"`
		Current   bool   `json:"current,omitempty"`
		Company   string `json:"company,omitempty"`
		Industry  string `json:"industry,omitempty"`
	} `json:"experience_entries,omitempty"`
	Account struct {
		Subdomain string `json:"subdomain,omitempty"`
		Name      string `json:"name,omitempty"`
	} `json:"account,omitempty"`
	Job struct {
		Shortcode string `json:"shortcode,omitempty"`
		Title     string `json:"title,omitempty"`
	} `json:"job,omitempty"`
	Stage              string    `json:"stage,omitempty"`
	Disqualified       bool      `json:"disqualified,omitempty"`
	DisqualifiedReason string    `json:"disqualification_reason,omitempty"`
	Sourced            bool      `json:"sourced,omitempty"`
	ProfileUrl         string    `json:"profile_url,omitempty"`
	Domain             string    `json:"domain,omitempty"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
	UpdatedAt          time.Time `json:"updated_at,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *Candidate) String() string {
	if data, err := json.MarshalIndent(this, "", "  "); err != nil {
		panic(err)
	} else {
		return "<workable.candidate " + string(data) + ">"
	}
}
