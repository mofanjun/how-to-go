package parser

import (
	"crawler/singleCrawler/engine"
	"regexp"
	"crawler/singleCrawler/model"
	"strconv"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d]+)KG</span></td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var hokouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var xinzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	//name
	profile.Name = name
	//age
	age,err := strconv.Atoi(string(extractString(contents,ageRe)))
	if err == nil {
		profile.Age = age
	}
	//height
	height,err := strconv.Atoi(string(extractString(contents,heightRe)))
	if err == nil {
		profile.Height = height
	}
	//weight
	weight,err := strconv.Atoi(string(extractString(contents,weightRe)))
	if err == nil {
		profile.Weight = weight
	}
	//gender
	profile.Gender = extractString(contents,genderRe)
	//income
	profile.Income = extractString(contents,incomeRe)
	//marriage
	profile.Marriage = extractString(contents,marriageRe)
	//education
	profile.Education = extractString(contents,educationRe)
	//occupation
	profile.Occupation = extractString(contents,occupationRe)
	//hoKou
	profile.Hokou = extractString(contents,hokouRe)
	//xinzuo
	profile.Xinzuo = extractString(contents,xinzuoRe)
	//house
	profile.House = extractString(contents,houseRe)
	//car
	profile.Car = extractString(contents,carRe)

	result := engine.ParseResult{
		Items: [] interface{}{profile},
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
