package tools

import "fmt"

type Club struct {
	Id               string   `json:"id"`
	CreatedOn        string   `json:"createdOn"`
	UpdatedOn        string   `json:"updatedOn"`
	Name             string   `json:"name"`
	BriefDescription string   `json:"briefDescription"`
	Location         string   `json:"location"`
	WebsiteUrl       string   `json:"websiteUrl"`
	KeyFeatures      string   `json:"keyFeatures"`
	Address          Address  `json:"address"`
	Latitude         float32  `json:"latitude"`
	Longitude        float32  `json:"longitude"`
	AmaChartered     bool     `json:"amaChartered"`
	Driving          bool     `json:"driving"`
	Flying           bool     `json:"flying"`
	Boating          bool     `json:"boating"`
}

// type Location struct {
// 	Latitude float32
// 	Longitude float32
// }

type Address struct {
	Street string  `json:"street"`
	City   string  `json:"city"`
	State  string  `json:"state"`
	Zip    string  `json:"zip"`
}

type DatabaseInterface interface {
	GetClubs() []Club
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{}

	var err error = database.SetupDatabase()
	if (err != nil) {
		fmt.Println(err)
		return nil, err
	}

	return &database, nil

}