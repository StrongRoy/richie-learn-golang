package parser

import (
	"io/ioutil"
	"richie.com/richie/learn_golang/crawler/engine"
	"richie.com/richie/learn_golang/crawler/model"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile(contents, "http://album.zhenai.com/u/1361568533", "Need阳光", "女")

	// verify result

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 "+
			"element; but was %v", result.Items)
	}

	actual := result.Items[0]

	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/1361568533",
		Type: "zhenai",
		Id:   "1361568533",
		Payload: model.Profile{
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

	// Verify result
	if actual != expected {
		t.Errorf("expected %v but was %v",
			expected, actual)
	}
}
