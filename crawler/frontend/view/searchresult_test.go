package view

import (
	"os"
	"richie.com/richie/learn_golang/crawler/engine"
	"richie.com/richie/learn_golang/crawler/frontend/model"
	comment "richie.com/richie/learn_golang/crawler/model"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView("template.html")
	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1361568533",
		Type: "zhenai",
		Id:   "1361568533",
		Payload: comment.Profile{
			Age:        29,
			Height:     163,
			Weight:     60,
			Name:       "Need阳光",
			Gender:     "女",
			Income:     "3-5千",
			Marriage:   "未婚",
			Education:  "大学本科",
			WorkPlace:  "阿坝汶川",
			Occupation: "政府机构",
			Xinzuo:     "魔羯座(12.22-01.19)",
		},
	}

	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	out, err := os.Create("template.test.html")
	err = view.Render(out,page)

	if err != nil {
		panic(err)
	}

}
