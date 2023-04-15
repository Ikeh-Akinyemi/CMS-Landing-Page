package utils

import (
	"fmt"

	"github.com/ButterCMS/go-landing-page/server/internal/model"
)

func ProcessFeaturesSection(section map[string]interface{}) []model.Feature {
	unparsedFeatures, _ := GetValue[[]interface{}](section, "features")

	features := []model.Feature{}
	for _, unparsedFeature := range unparsedFeatures {
		typedFeature := unparsedFeature.(map[string]interface{})

		headline, _ := GetValue[string](typedFeature, "headline")
		description, _ := GetValue[string](typedFeature, "description")

		feature := model.Feature{
			Headline: headline,
			Description: description,
		}

		features = append(features, feature)
	}

	return features
}

func ProcessProductsSection(section map[string]interface{}) []model.Product {
	unparsedProducts, _ := GetValue[[]interface{}](section, "products")

	products := []model.Product{}
	for _, unparsedProduct := range unparsedProducts {
		typedProduct := unparsedProduct.(map[string]interface{})

		name, _ := GetValue[string](typedProduct, "name")
		price, _ := GetValue[string](typedProduct, "price")
		imageURL, _ := GetValue[string](typedProduct, "image_url")
		imgAltText, _ := GetValue[string](typedProduct, "image_description")

		product := model.Product{
			Name: name,
			Price: price,
			ImageURL: imageURL,
			ImgAltText: imgAltText,
		}

		products = append(products, product)
	}

	return products
}

func ProcessHeroSection(section map[string]interface{}) model.Hero {
	unparsedHero, _ := GetValue[map[string]interface{}](section, "hero")

	headline, _ := GetValue[string](unparsedHero, "headline")
	description, _ := GetValue[string](unparsedHero, "description")
	imageURL, _ := GetValue[string](unparsedHero, "image_url")
	buttonLabel, _ := GetValue[string](unparsedHero, "btn_label")

	return model.Hero{
		Headline: headline,
		Description: description,
		ImageURL: imageURL,
		ButtonLabel: buttonLabel,
	}
}

func GetValue[T any](input map[string]interface{}, name string) (T, error) {
	field, ok := input[name].(T)

	if !ok {
		var null T
		return null, fmt.Errorf("unable to get property %s from map", name)
	}

	return field, nil
}