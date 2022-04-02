package apiserver

import (
	"html/template"
)

var navigationBarHTML string
var homepageTpl *template.Template
var secondViewTpl *template.Template
var thirdViewTpl *template.Template

func init() {
	//navigationBarHTML = assets.MustAssetString("templates/navigation_bar.html")
	//
	//homepageHTML := assets.MustAssetString("templates/index.html")
	//homepageTpl = template.Must(template.New("homepage_view").Parse(homepageHTML))
	//
	//secondViewHTML := assets.MustAssetString("templates/second_view.html")
	//secondViewTpl = template.Must(template.New("second_view").Parse(secondViewHTML))
	//
	//thirdViewFuncMap := ThirdViewFormattingFuncMap()
	//thirdViewHTML := assets.MustAssetString("templates/third_view.html")
	//thirdViewTpl = template.Must(template.New("third_view").Funcs(thirdViewFuncMap).Parse(thirdViewHTML))
}
