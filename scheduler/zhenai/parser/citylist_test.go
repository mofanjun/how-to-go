package parser

import (
	"testing"
	"io/ioutil"
)

func TestParseCityList(t *testing.T) {
	contens, err := ioutil.ReadFile("citylist_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseCityList(contens)

	const resultSize = 470
	expectCities := []string{"阿坝","阿克苏","阿拉善盟ƒ"}

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d",resultSize,
			len(result.Requests))
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d items; but had %d",resultSize,
			len(result.Items))
	}

	for i,city := range expectCities {
		if result.Items[i].(string) != city {
			t.Errorf("excepted city #%d: %s; but was %s",i,city,result.Items[i])
		}
	}
}
