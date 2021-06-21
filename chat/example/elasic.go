package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"reflect"
)

//see: https://www.bookstack.cn/read/topgoer/82d8230568a4be5b.md

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

var host = "http://192.168.8.76:9200"

func example() {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(host))
	if err != nil {
		fmt.Printf("new client err:%v\n", err)
		return
	}
	p := Person{
		Name:    "zhangsan",
		Age:     10,
		Married: false,
	}
	put, err := client.Index().Index("user").BodyJson(p).Do(context.Background())
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put.Id, put.Index, put.Type)
}

var client *elastic.Client

func Init() {
	var err error
	client, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(host))
	if err != nil {
		panic(err)
	}
	info, code, err := client.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esversion, err := client.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)
}

func main() {
	Init()
	create()
	delete()
	update()
	gets()
	query()
	list(3, 1)
}

type Employee struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	About     string   `json:"about"`
	Interests []string `json:"interests"`
}

func create() {
	//使用结构体
	e1 := Employee{"Jane", "Smith", 32, "I like to collect rock albums", []string{"music"}}
	put1, err := client.Index().Index("megacorp").
		Id("1").BodyJson(e1).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put1.Id, put1.Index, put1.Type)

	e2 := `{"first_name":"John","last_name":"Smith","age":25,"about":"I love to go rock climbing","interests":["sports","music"]}`
	put2, err := client.Index().Index("megacorp").Id("2").
		BodyJson(e2).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put2.Id, put2.Index, put2.Type)

	e3 := `{"first_name":"Douglas","last_name":"Fir","age":35,"about":"I like to build cabinets","interests":["forestry"]}`
	put3, err := client.Index().Index("megacorp").Id("3").
		BodyJson(e3).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put3.Id, put3.Index, put3.Type)
}

func delete() {
	res, err := client.Delete().Index("megacorp").Id("1").Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("delete result %v\n", res.Result)
}

func update() {
	res, err := client.Update().Index("megacorp").Id("2").
		Doc(map[string]interface{}{"age": 21}).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("update age %s\n", res.Result)
}

func gets() {
	get, err := client.Get().Index("megacorp").Id("2").Do(context.Background())
	if err != nil {
		panic(err)
	}
	if get.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get.Id, get.Version, get.Index, get.Type)
	}
}

func query() {
	var res *elastic.SearchResult
	var err error

	res, err = client.Search("megacorp").Do(context.Background())
	printEmployee(res, err)

	q := elastic.NewQueryStringQuery("last_name:smith")
	res, err = client.Search("megacorp").Query(q).Do(context.Background())
	printEmployee(res, err)

	//条件查询
	//年龄大于30岁的
	boolQ := elastic.NewBoolQuery()
	boolQ.Must(elastic.NewMatchQuery("last_name", "smith"))
	boolQ.Filter(elastic.NewRangeQuery("age").Gt(30))
	res, err = client.Search("megacorp").Query(boolQ).Do(context.Background())
	printEmployee(res, err)

	//短语搜索 搜索about字段中有 rock climbing
	matchPhraseQuery := elastic.NewMatchPhraseQuery("about", "rock climbing")
	res, err = client.Search("megacorp").Query(matchPhraseQuery).Do(context.Background())
	printEmployee(res, err)

	//分析 interests
	aggs := elastic.NewTermsAggregation().Field("interests.keyword")
	res, err = client.Search("megacorp").Aggregation("all_interests", aggs).Do(context.Background())
	printEmployee(res, err)
}

func list(size, page int) {
	res, err := client.Search("megacorp").Size(size).From((page - 1) * size).Do(context.Background())
	printEmployee(res, err)
}

//打印查询到的Employee
func printEmployee(res *elastic.SearchResult, err error) {
	if err != nil {
		panic(err)
	}
	var typ Employee
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(Employee)
		fmt.Printf("%#v\n", t)
	}
	fmt.Println("-----------")
}
