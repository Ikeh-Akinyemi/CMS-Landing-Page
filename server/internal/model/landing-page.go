package model

type Product struct {
	// Name of the product
	Name        string `json:"name"`
	// Price of the product
	Price       string `json:"price"`
	// URL for the product's image
	ImageURL    string `json:"image_url"`
	// Alternative text for the product's image
	ImgAltText  string `json:"image_description"`
}

type Hero struct {
	// Headline for the hero section
	Headline    string `json:"headline"`
	// Description for the hero section
	Description string `json:"description"`
	// URL for the hero section's image
	ImageURL    string `json:"image_url"`
	// Label for the hero section's button
	ButtonLabel string `json:"btn_label"`
}

type Feature struct {
	// Headline for the feature
	Headline    string `json:"headline"`
	// Description for the feature
	Description string `json:"description"`
}

type LandingPage struct {
	// Title of the landing page
	PageTitle string `json:"page_title"`
	// Data for the hero section
	Hero Hero `json:"hero"`
	// List of features for the landing page
	Features []Feature `json:"features"`
	// List of products for the landing page
	Products []Product `json:"products"`
}
