package parser

import (
	"testing"
	"io/ioutil"
	"crawler/singleCrawler/model"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents,"小顺儿")

	if len(result.Items) != 1 {
		t.Errorf("result Item length must 1,but get %v",result)
	}

	exceped := model.Profile{
		Name:"小顺儿",
		Gender:"女",
		Age:29,
		Height:169,
		Weight:52,
		Income:"3001-5000元",
		Marriage:"未婚",
		Education:"大学本科",
		Occupation:"会计",
		Hokou:"四川阿坝",
		Xinzuo:"魔羯座",
		House:"和家人同住",
		Car:"未购车",
	}

	profile := result.Items[0].(model.Profile)

	if exceped != profile {
		t.Errorf("Exceped %v but get %v",exceped,profile)
	}
}
