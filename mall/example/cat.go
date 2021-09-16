package main

import (
	"encoding/json"
	"fmt"
)

// see: https://github.com/ideaviewes/simpleGoTree
type Cat struct {
	Id       int
	Fid      int
	Name     string
	Children []*Cat `orm:"-"`
}

func main() {
	list := []*Cat{
		{
			Id:       1,
			Fid:      0,
			Name:     "test1",
		},{
			Id:       2,
			Fid:      1,
			Name:     "test2",
		},{
			Id:       22,
			Fid:      1,
			Name:     "test2",
		},{
			Id:       3,
			Fid:      2,
			Name:     "test3",
		},{
			Id:       4,
			Fid:      3,
			Name:     "test4",
		},{
			Id:       5,
			Fid:      4,
			Name:     "test5",
		},{
			Id:       6,
			Fid:      5,
			Name:     "test6",
		},
	}
	child := makeChild(5, list)
	fmt.Printf("tree:%+v", child)
}

func tree(list []*Cat) string {
	data := buildData(list)
	result := makeTreeCore(0, data)
	body, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

func buildData(list []*Cat) map[int]map[int]*Cat {
	var data  = make(map[int]map[int]*Cat)
	for _, v := range list {
		id := v.Id
		fid := v.Fid
		if _, ok := data[fid]; !ok {
			data[fid] = make(map[int]*Cat)
		}
		data[fid][id] = v
	}
	return data
}

func makeTreeCore(index int, data map[int]map[int]*Cat) []*Cat {
	tmp := make([]*Cat, 0)
	for id, item := range data[index] {
		if data[id] != nil {
			item.Children = makeTreeCore(id, data)
		}
		tmp = append(tmp, item)
	}
	return tmp
}

func makeChild(id int, list []*Cat) (ids []int) {
	for _, categoryModel := range list {
		if categoryModel.Fid == id {
			ids = append(ids, makeChild(categoryModel.Id, list)...)
		}
	}
	ids = append(ids, id)
	return ids
}