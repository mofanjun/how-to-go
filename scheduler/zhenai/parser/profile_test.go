package parser

import (
	"testing"
	"io/ioutil"
	"crawler/scheduler/model"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents,"灰兔儿")

	if len(result.Items) != 1 {
		t.Errorf("result Item length must 1,but get %v",result)
	}

	exceped := model.Profile{
		Name:"灰兔儿",
		Gender:"女",
		Age:39,
		Height:161,
		Weight:52,
		Income:"20001-50000元",
		Marriage:"离异",
		Education:"中专",
		Occupation:"未知",
		Hokou:"浙江杭州",
		Xinzuo:"天秤座",
		House:"已购房",
		Car:"已买车",
	}

	profile := result.Items[0].(model.Profile)

	if exceped != profile {
		t.Errorf("Exceped %v but get %v",exceped,profile)
	}
}
