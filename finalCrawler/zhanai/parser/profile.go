package parser

import (
	"regexp"
	"crawler/finalCrawler/engine"
	"crawler/finalCrawler/model"
	"strings"
	"strconv"
)

var basicRe = regexp.MustCompile(`<div class="des f-cl"[^>]*>([^<]+)</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^<]+)kg</div>`)
var hokouRe = regexp.MustCompile(`<div class="m-btn pink"[^>]*>籍贯:([^<]+)</div>`)
var xinZuoRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([^/(]+)[\S]+-[\S]+</div>`)
var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)


func ParseProfile(contents []byte, url string, name string) engine.ParseResult {
	profile := model.Profile{}

	//
	profile.Name = name
	//基本信息
	str := string(extractString(contents,basicRe))
	str = strings.Replace(str," ","",-1)
	splitStr := strings.Split(str,"|")//[杭州 39岁 中专 离异 161cm 20001-50000元]
	if splitStr == nil || len(splitStr) < 2 {
		result := engine.ParseResult{
			Items: [] engine.Item{
				{
					Url:url,
					Type:"zhenai",
					Id: extractString([]byte(url),idUrlRe),
					PayLoad:profile,
				},
			},
		}
		return  result
	}
	//Age
	strAge := []rune(splitStr[1])
	age,err := strconv.Atoi(string(strAge[:len(strAge) - 1]))
	if err == nil {
		profile.Age = age
	}
	//eduction
	profile.Education = splitStr[2]
	//marriage
	profile.Marriage = splitStr[3]
	if len(splitStr) == 5 {
		profile.Income = splitStr[4]
	} else {
		//height
		height,err := strconv.Atoi(splitStr[4][:len(splitStr[4]) - 2])
		if err == nil {
			profile.Height = height
		}
		//income
		profile.Income = splitStr[5]
	}

	//gender
	if strings.Contains(string(contents),"女士征婚") {
		profile.Gender = "女"
	} else {
		profile.Gender = "男"
	}
	//weight
	weight,err := strconv.Atoi(string(extractString(contents,weightRe)))
	if err == nil {
		profile.Weight = weight
	}
	//hokou
	profile.Hokou = extractString(contents,hokouRe)
	//xinzuo
	profile.Xinzuo = extractString(contents,xinZuoRe)
	//TODO：Occupation后面补上
	profile.Occupation = "未知"
	//house
	if strings.Contains(string(contents),"已购房") {
		profile.House = "已购房"
	} else {
		profile.House = "未购房"
	}
	//car
	if strings.Contains(string(contents),"已买车") {
		profile.Car = "已买车"
	} else {
		profile.Car = "未买车"
	}
	//
	result := engine.ParseResult{
		//Items: [] interface{}{profile},
		Items: [] engine.Item{
			{
				Url:url,
				Type:"zhenai",
				Id: extractString([]byte(url),idUrlRe),
				PayLoad:profile,
			},
		},
	}
	//TODO:完成猜你喜欢的url
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

func ProfilePaser(name string) engine.ParseFunc{
	return func(bytes []byte, url string) engine.ParseResult {
		return ParseProfile(bytes, url, name)
	}
}
