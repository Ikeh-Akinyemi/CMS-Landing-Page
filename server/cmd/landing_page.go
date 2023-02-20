package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"os"

	ButterCMS "github.com/ButterCMS/buttercms-go"
	"github.com/ButterCMS/go-landing-page/server/internal/model"
	"github.com/ButterCMS/go-landing-page/server/internal/utils"
)

func landingPage(w http.ResponseWriter, r *http.Request) {
	ButterCMS.SetAuthToken(os.Getenv("BUTTERCMS_API_TOKEN"))
	data := map[string]string{}
	res, err := ButterCMS.GetPage("*", "fruit-veggie-landing-page", data)
	if err != nil {
		log.Fatal(err)
	}

	templateData := model.LandingPage{}
	templateData.PageTitle, _ = utils.GetValue[string](res.Page.Fields, "page_title")
	templateData.Features = utils.ProcessFeaturesSection(res.Page.Fields)
	templateData.Products = utils.ProcessProductsSection(res.Page.Fields)
	templateData.Hero = utils.ProcessHeroSection(res.Page.Fields)

	render(w, r, "./ui/html/template.html", &templateData)
}

// render retrieves and render the appropriate template set from the cache based on the page name
func render(w http.ResponseWriter, r *http.Request, name string, td *model.LandingPage) {
	t, err := template.ParseFiles(name)
	if err != nil {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)

	err = t.Execute(buf, td)
	if err != nil {
		log.Fatal(err)
	}

	buf.WriteTo(w)
}
