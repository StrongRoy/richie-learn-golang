package persist

import (
	"context"
	"encoding/json"
	"gopkg.in/olivere/elastic.v5"
	"richie.com/richie/learn_golang/crawler/engine"
	"richie.com/richie/learn_golang/crawler/model"
	"testing"
)

func TestSaver(t *testing.T) {

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


	// TODO: try to start up elastic search
	// here using docker go client.
	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	const index = "dating_test"


	// Save expected item
	err = Save(client,index,expected)
	if err != nil {
		panic(err)
	}

	// Fetch saved item
	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s", *resp.Source)

	var actual engine.Item

	json.Unmarshal(*resp.Source, &actual)

	actualProfile , _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile
	// Verify result
	if actual != expected {
		t.Errorf("expected %v but was %v",
			expected, actual)
	}
}
