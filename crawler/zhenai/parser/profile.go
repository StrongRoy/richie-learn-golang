package parser

import (
	"regexp"
	"richie.com/richie/learn_golang/crawler/engine"
	"richie.com/richie/learn_golang/crawler/model"
	"richie.com/richie/learn_golang/crawler_distributed/config"
	"strconv"
)

var ageMarriageXingzuoRe = regexp.MustCompile(`<div.class="m-btn.purple"[^>]+>([^<]+)</div><div.class="m-btn.purple"[^>]+>(\d+)岁</div><div.class="m-btn.purple"[^>]+>([^<]+)</div>`)
var heightWeightRe = regexp.MustCompile(`<div.class="m-btn.purple"[^>]+>(\d+)cm</div><div.class="m-btn.purple"[^>]+>(\d+)kg</div><div.class="m-btn.purple"[^>]+>工作地:([^<]+)</div>`)

var inComeRe = regexp.MustCompile(`<div.class="m-btn.purple"[^>]+>月收入:([^<]+)</div><div.class="m-btn.purple"[^>]+>([^<]+)</div><div.class="m-btn.purple"[^>]+>([^<]+)</div>`)
var guessRe = regexp.MustCompile(`<a class="exp-user-name"[^>]*href="(http://album.zhenai.com/u/[\d]+)">([^<])`)
var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func parseProfile(contents []byte, url, name, gender string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	profile.Gender = gender

	match := ageMarriageXingzuoRe.FindSubmatch(contents)
	if len(match) >= 4 {
		profile.Marriage = string(match[1])
		age, err := strconv.Atoi(string(match[2]))
		if err == nil {
			profile.Age = age
		}
		profile.Xinzuo = string(match[3])
	}

	match = heightWeightRe.FindSubmatch(contents)
	if len(match) >= 4 {
		height, err := strconv.Atoi(string(match[1]))
		if err == nil {
			profile.Height = height
		}
		weight, err := strconv.Atoi(string(match[2]))
		if err == nil {
			profile.Weight = weight
		}
		profile.WorkPlace = string(match[3])
	}
	match = inComeRe.FindSubmatch(contents)
	if len(match) >= 4 {
		profile.Income = string(match[1])
		profile.Occupation = string(match[2])
		profile.Education = string(match[3])
	}

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      extractString([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}
	//猜你喜欢的人
	matchs := guessRe.FindAllSubmatch(contents, -1) //-1表示匹配所有
	for _, m := range matchs {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: NewProfileParser(string(m[2]), string(m[3])),
		})
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

type ParseProfile struct {
	userName   string
	userGender string
}

func (p *ParseProfile) Parser(
	contents []byte, url string) engine.ParseResult {
	return parseProfile(contents, url, p.userName, p.userGender)
}

func (p *ParseProfile) Serialize() (name string, args interface{}) {
	return config.ParseProfile, p.userName
}

func NewProfileParser(name, gender string) *ParseProfile {
	return &ParseProfile{
		userName:   name,
		userGender: gender,
	}
}
