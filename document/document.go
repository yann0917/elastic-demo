package document

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	. "github.com/yann0917/elastic-demo/client"
	"github.com/yann0917/elastic-demo/utils"
)

type User struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Message  string    `json:"message"`
	PostDate time.Time `json:"post_date"`
}

var (
	index = "users"
	ctx   = context.Background()
)

func Index() {
	user := User{
		ID:       "1",
		Name:     "Mike",
		Message:  "trying out Kibana",
		PostDate: time.Now(),
	}
	doc, err := Client.Index().
		Index(index).
		Id(user.ID).
		BodyJson(user).
		Do(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.PrintIndent(doc)

}

func Get(id string) (user User, err error) {

	result, err := Client.Get().
		Index(index).
		Id(id).
		Do(ctx)
	if err != nil {
		fmt.Println(err)
		// panic(err)
		return
	}

	if result.Found {
		fmt.Printf("Got document %v (version=%d, index=%s, type=%s)\n",
			result.Id, result.Version, result.Index, result.Type)
		utils.PrintIndent(result.Source)

		err = json.Unmarshal(result.Source, &user)
	}

	return
}

func Update(id string) (err error) {
	user := User{
		ID:       "2",
		Name:     "Mike",
		Message:  "trying out elasticSearch",
		PostDate: time.Now(),
	}
	result, err := Client.Update().
		Index(index).
		Id(id).
		Doc(user).
		Do(ctx)
	if err != nil {
		fmt.Println(err)
		// panic(err)
		return
	}

	utils.PrintIndent(result)

	return

}

func Delete(id string) {
	res, err := Client.Delete().
		Index(index).
		Id(id).
		Do(ctx)
	if err != nil {
		panic(err)
	}
	if res.Result == "deleted" {
		fmt.Println("Document " + id + ": deleted")
	}
	utils.PrintIndent(res)

}
