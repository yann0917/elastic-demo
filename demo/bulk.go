package demo

import (
	"fmt"
	"reflect"
	"time"

	"github.com/olivere/elastic/v7"
	. "github.com/yann0917/elastic-demo/client"
)

func Bulk() {
	users := []User{
		{
			ID:       "1",
			Name:     "Mike",
			Message:  "trying out Kibana",
			PostDate: time.Now(),
		},
		{
			ID:       "2",
			Name:     "Jack",
			Message:  "trying out ElasticSearch",
			PostDate: time.Now(),
		},
	}

	user3 := User{
		ID:       "3",
		Name:     "Tom",
		Message:  "I am Tom",
		PostDate: time.Now(),
	}

	user4 := User{
		ID:       "4",
		Name:     "Jerry",
		Message:  "I am Jerry",
		PostDate: time.Now(),
	}

	user5 := user4
	user5.Name = "Lawson"

	req := Client.Bulk()
	for _, user := range users {
		doc := elastic.NewBulkIndexRequest().Index(index).Id(user.ID).Doc(user)
		req = req.Add(doc)
	}

	req3 := elastic.NewBulkIndexRequest().Index(index).Id(user3.ID).Doc(user3)
	req4 := elastic.NewBulkIndexRequest().OpType("create").Index(index).Id(user4.ID).Doc(user4)
	req5 := elastic.NewBulkDeleteRequest().Index(index).Id("1")
	req6 := elastic.NewBulkUpdateRequest().Index(index).Id("2").Doc(user5)

	req = req.Add(req3).Add(req4).Add(req5).Add(req6)

	if req.NumberOfActions() == 0 {
		fmt.Println("Actions all clear!")
	}
	resp, err := req.Do(ctx)
	if err != nil {
		fmt.Println(err)
	}
	// utils.PrintIndent(resp)
	failed := resp.Failed()
	l := len(failed)
	if l > 0 {
		fmt.Printf("Error(%d), %#v\n", l, resp.Errors)

	}

	searchResult, err := Client.Search().
		Index(index).
		Sort("id", false). // 按id升序排序
		Pretty(true).
		Do(ctx) // 执行
	if err != nil {
		panic(err)
	}

	var u User
	for _, item := range searchResult.Each(reflect.TypeOf(u)) {
		if t, ok := item.(User); ok {
			fmt.Printf("Found: User(id=%s, Name=%s, Message=%s)\n", t.ID, t.Name, t.Message)
		}
	}
}
