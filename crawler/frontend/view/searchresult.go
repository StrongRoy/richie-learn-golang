package view

import (
	"html/template"
	"io"
	"richie.com/richie/learn_golang/crawler/frontend/model"
)

// ToDo
// fill in query string
// support search button
// rewrite query string
// support paging
// add start page

type SearchResultView struct {
	template *template.Template
}

func CreateSearchResultView(
	filename string) SearchResultView {

	return SearchResultView{
		template: template.Must(
			template.ParseFiles(filename)),
	}
}



func (s SearchResultView) Render(w io.Writer,data model.SearchResult) error {
	return s.template.Execute(w,data)
}