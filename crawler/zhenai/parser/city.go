package parser

import (
	"regexp"
	"richie.com/richie/learn_golang/crawler/engine"
	"richie.com/richie/learn_golang/crawler_distributed/config"
)

var (
	profileRe = regexp.MustCompile(
		`<th><a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a></th>.*?><span[^>]*>[^>]*>([^<]+)</td>`)
	cityUrlRe = regexp.MustCompile(
		`href="(http://www.zhenai.com/zhenghun/[^"]+)`)
)

func ParseCity(
	contents []byte,_ string) engine.ParseResult {
	matches := profileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(m[1]),
				Parser: NewProfileParser(string(m[2]),string(m[3])),
			})
	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url:        string(m[1]),
				Parser: engine.NewFuncParser(ParseCity,config.ParseCity),
			})
	}
	return result

}
