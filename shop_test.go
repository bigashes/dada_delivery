package dada_delivery

import (
	"testing"
)

func TestCityList(t *testing.T) {
	client := NewClient(Config{
		AppKey:    "",
		AppSecret: "",
		SourceId:  "",
	})

	client.Debug = true

	list, err := client.CityList()
	if err != nil {
		t.Error(err)
		return
	}

	for _, item := range list {
		t.Logf("city %+v \n", item)
	}
	t.Logf("city len %d \n", len(list))
}
