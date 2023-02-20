package model

type Product struct {
	Name        string `json:"name"`
	Price       string `json:"price"`
	ImageURL    string `json:"image_url"`
	ImgAltText  string `json:"image_description"`
}

type Hero struct {
	Headline    string `json:"headline"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	ButtonLabel string `json:"btn_label"`
}

type Feature struct {
	Headline    string `json:"headline"`
	Description string `json:"description"`
}

type LandingPage struct {
	PageTitle string `json:"page_title"`
	Hero Hero `json:"hero"`
	Features []Feature `json:"features"`
	Products []Product `json:"products"`
}