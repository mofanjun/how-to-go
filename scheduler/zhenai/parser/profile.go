package parser

import (
	"crawler/scheduler/engine"
	"regexp"
	"crawler/scheduler/model"
	"strings"
	"strconv"
	"fmt"
)

//var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
//var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
//var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
//var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d]+)KG</span></td>`)
//var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
//var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
//var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
//var occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
//var hokouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
//var xinzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span>([^<]+)</td>`)
//var houseRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
//var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)

var basicRe = regexp.MustCompile(`<div class="des f-cl"[^>]*>([^<]+)</div>`)


//func ParseProfile(contents []byte, name string) engine.ParseResult {
//	profile := model.Profile{}
//	//name
//	profile.Name = name
//	//age
//	age,err := strconv.Atoi(string(extractString(contents,ageRe)))
//	if err == nil {
//		profile.Age = age
//	}
//	//height
//	height,err := strconv.Atoi(string(extractString(contents,heightRe)))
//	if err == nil {
//		profile.Height = height
//	}
//	//weight
//	weight,err := strconv.Atoi(string(extractString(contents,weightRe)))
//	if err == nil {
//		profile.Weight = weight
//	}
//	//gender
//	profile.Gender = extractString(contents,genderRe)
//	//income
//	profile.Income = extractString(contents,incomeRe)
//	//marriage
//	profile.Marriage = extractString(contents,marriageRe)
//	//education
//	profile.Education = extractString(contents,educationRe)
//	//occupation
//	profile.Occupation = extractString(contents,occupationRe)
//	//hoKou
//	profile.Hokou = extractString(contents,hokouRe)
//	//xinzuo
//	profile.Xinzuo = extractString(contents,xinzuoRe)
//	//house
//	profile.House = extractString(contents,houseRe)
//	//car
//	profile.Car = extractString(contents,carRe)
//
//	result := engine.ParseResult{
//		Items: [] interface{}{profile},
//	}
//
//	return result
//}

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}

	//
	profile.Name = name
	//基本信息
	str := string(extractString(contents,basicRe))
	str = strings.Replace(str," ","",-1)
	splitStr := strings.Split(str,"|")//[杭州 39岁 中专 离异 161cm 20001-50000元]
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
	//height
	height,err := strconv.Atoi(splitStr[4][:len(splitStr[4]) - 2])
	if err == nil {
		profile.Height = height
	}
	//income
	profile.Income = splitStr[5]
	//gender
	//weight
	//hokou
	//xinzuo
	//TODO：Occupation后面补上
	profile.Occupation = "未知"
	//house
	if strings.Contains(string(contents),"已购房") {
		profile.House = "已购房"
	} else {
		profile.House = "未购房"
	}
	//car
	if strings.Contains(string(contents),"已购车") {
		profile.House = "已购车"
	} else {
		profile.House = "未购车"
	}
	//
	fmt.Println("Now profile >",profile)

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
