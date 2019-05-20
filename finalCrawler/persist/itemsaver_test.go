package persist

import "testing"
import "crawler/finalCrawler/model"

func TestSave(t *testing.T) {
	profile := model.Profile{
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

	save(profile)
}
