package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/yann0917/elastic-demo/client"
)

var (
	subject   Subject
	indexName = "movies"
)

type Subject struct {
	ID     string   `json:"id"`
	Title  string   `json:"title"`
	Genres []string `json:"genre"`
	Year   int      `json:"year"`
}

const mapping = `{
	"mappings": {
		"properties": {
		  "@version": {
			"type": "text",
		  },
		  "genre": {
			"type": "text",
		  },
		  "id": {
			"type": "text",
		  },
		  "title": {
			"type": "text",
		  },
		  "year": {
			"type": "long"
		  }
		}
	  }
  }`

func main() {
	ctx := context.Background()
	exists, err := client.Client.IndexExists(indexName).Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		_, err := client.Client.CreateIndex(indexName).BodyString(mapping).Do(ctx)
		if err != nil {
			panic(err)
		}
	}

	result, err := client.Client.Get().
		Index(indexName).
		Id(strconv.Itoa(1)).
		Do(ctx)
	if err != nil {
		panic(err)
	}
	if result.Found {
		fmt.Printf("Got document %v (version=%d, index=%s, type=%s)\n",
			result.Id, result.Version, result.Index, result.Type)
		err := json.Unmarshal(result.Source, &subject)
		if err != nil {
			panic(err)
		}
		fmt.Println(subject.ID, subject.Title, subject.Genres, subject.Year)
	}
}
