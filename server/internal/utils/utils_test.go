package utils

import (
	"testing"

	"github.com/ButterCMS/go-landing-page/server/internal/model"
)

func TestProcessFeaturesSection(t *testing.T) {
	testSection := map[string]interface{}{
		"features": []interface{}{
			map[string]interface{}{
				"headline":    "Feature 1",
				"description": "Description for Feature 1",
			},
			map[string]interface{}{
				"headline":    "Feature 2",
				"description": "Description for Feature 2",
			},
		},
	}

	expectedFeatures := []model.Feature{
		{
			Headline:    "Feature 1",
			Description: "Description for Feature 1",
		},
		{
			Headline:    "Feature 2",
			Description: "Description for Feature 2",
		},
	}

	resultFeatures := ProcessFeaturesSection(testSection)

	for i := 0; i < len(resultFeatures); i++ {
		if resultFeatures[i].Headline != expectedFeatures[i].Headline ||
			resultFeatures[i].Description != expectedFeatures[i].Description {
			t.Errorf("Unexpected feature result at index %d, got: %v, want: %v", i, resultFeatures[i], expectedFeatures[i])
		}
	}
}

func TestProcessProductsSection(t *testing.T) {
	testSection := map[string]interface{}{
		"products": []interface{}{
			map[string]interface{}{
				"name":              "Product 1",
				"price":             "$10.00",
				"image_url":         "https://example.com/product1.png",
				"image_description": "Image description for Product 1",
			},
			map[string]interface{}{
				"name":              "Product 2",
				"price":             "$20.00",
				"image_url":         "https://example.com/product2.png",
				"image_description": "Image description for Product 2",
			},
		},
	}

	expectedProducts := []model.Product{
		{
			Name:       "Product 1",
			Price:      "$10.00",
			ImageURL:   "https://example.com/product1.png",
			ImgAltText: "Image description for Product 1",
		},
		{
			Name:       "Product 2",
			Price:      "$20.00",
			ImageURL:   "https://example.com/product2.png",
			ImgAltText: "Image description for Product 2",
		},
	}

	resultProducts := ProcessProductsSection(testSection)

	for i := 0; i < len(resultProducts); i++ {
		if resultProducts[i].Name != expectedProducts[i].Name ||
			resultProducts[i].Price != expectedProducts[i].Price ||
			resultProducts[i].ImageURL != expectedProducts[i].ImageURL ||
			resultProducts[i].ImgAltText != expectedProducts[i].ImgAltText {
			t.Errorf("Unexpected product result at index %d, got: %v, want: %v", i, resultProducts[i], expectedProducts[i])
		}
	}
}

func TestProcessHeroSection(t *testing.T) {
	testSection := map[string]interface{}{
			"hero": map[string]interface{}{
					"headline":    "Hero headline",
					"description": "Hero description",
					"image_url":   "https://example.com/hero.png",
					"btn_label":   "Hero button label",
			},
	}

	expectedHero := model.Hero{
			Headline:    "Hero headline",
			Description: "Hero description",
			ImageURL:    "https://example.com/hero.png",
			ButtonLabel: "Hero button label",
	}

	hero := ProcessHeroSection(testSection)
	if hero != expectedHero {
			t.Errorf("Unexpected result. Got %v, expected %v", hero, expectedHero)
	}
}

func TestGetValue(t *testing.T) {
	testMap := map[string]interface{}{
			"name": "John",
			"age":  30,
	}

	// Test getting a string value
	name, err := GetValue[string](testMap, "name")
	if err != nil {
			t.Errorf("Unexpected error: %v", err)
	}
	if name != "John" {
			t.Errorf("Unexpected value. Got %v, expected 'John'", name)
	}

	// Test getting an int value
	age, err := GetValue[int](testMap, "age")
	if err != nil {
			t.Errorf("Unexpected error: %v", err)
	}
	if age != 30 {
			t.Errorf("Unexpected value. Got %v, expected 30", age)
	}

	// Test getting a value that doesn't exist
	_, err = GetValue[string](testMap, "address")
	if err == nil {
			t.Errorf("Expected an error, but got nil")
	}
}
