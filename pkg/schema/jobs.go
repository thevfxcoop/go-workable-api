package schema

import (
	"encoding/json"
	"time"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Job struct {
	Id             string `json:"id"`
	Title          string `json:"title,omitempty"`
	FullTitle      string `json:"full_title,omitempty"`
	Shortcode      string `json:"shortcode,omitempty"`
	Code           string `json:"code,omitempty"`
	State          string `json:"state,omitempty"`
	Department     string `json:"department,omitempty"`
	Url            string `json:"url,omitempty"`
	ApplicationUrl string `json:"application_url,omitempty"`
	ShortLink      string `json:"shortlink,omitempty"`
	Location       struct {
		LocationStr   string `json:"location_str,omitempty"`
		Country       string `json:"country,omitempty"`
		CountryCode   string `json:"country_code,omitempty"`
		Region        string `json:"region,omitempty"`
		RegionCode    string `json:"region_code,omitempty"`
		City          string `json:"city,omitempty"`
		ZipCode       string `json:"zip_code,omitempty"`
		Telecommuting bool   `json:"telecommuting,omitempty"`
	} `json:"location,omitempty"`
	Description     string    `json:"description,omitempty"`
	FullDescription string    `json:"full_description,omitempty"`
	Industry        string    `json:"industry,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *Job) String() string {
	if data, err := json.MarshalIndent(this, "", "  "); err != nil {
		panic(err)
	} else {
		return "<workable.job " + string(data) + ">"
	}
}
