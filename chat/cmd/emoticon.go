package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"chat/app/logic/model"
	"chat/pkg/net/http"
)

var url = "https://raw.githubusercontent.com/zhaoolee/ChineseBQB/master/chinesebqb_github.json"

type result struct {
	Status int        `json:"status"`
	Info   string     `json:"info"`
	Data   []dataList `json:"data"`
}

type dataList struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Url      string `json:"url"`
}

//SyncBQB 同步 https://github.com/zhaoolee/ChineseBQB 表情包入库
func SyncBQB() error {
	client := http.NewRestyClient()
	rsp, err := client.Get(context.Background(), url)
	if err != nil {
		return err
	}
	var rs result
	err = json.Unmarshal(rsp, &rs)
	if err != nil {
		return err
	}
	if rs.Status != 1000 {
		return errors.New(rs.Info)
	}
	var emot []model.EmoticonModel
	for i, datum := range rs.Data {
		nameStart := strings.LastIndex(datum.Name, "-")
		nameEnd := strings.LastIndex(datum.Name, ".")
		if nameEnd == -1 {
			nameEnd = len(datum.Name)
		}
		if nameStart > nameEnd {
			nameStart = 0
		}
		cat := datum.Category[strings.LastIndex(datum.Category, "_")+1:]
		if cat == "BQB" {
			cat = datum.Category[strings.Index(datum.Category, "_")+1:]
		}
		if cat == "BQB" {
			cat = datum.Category
		}
		emot = append(emot, model.EmoticonModel{
			PriID:    model.PriID{ID: uint32(i + 1)},
			Category: cat,
			Name:     datum.Name[nameStart+1 : nameEnd],
			Url:      datum.Url,
		})
	}
	//先清空表
	model.DB.Exec("truncate emoticon")
	return model.DB.CreateInBatches(emot, 200).Error
}
