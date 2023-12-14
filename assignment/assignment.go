package assignment

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"myproject/models"
	"myproject/service"
)

const url = "https://jsonplaceholder.typicode.com/todos/"

func FetchData(httpClient service.HttpInterface, id int) error {
	todo := models.Todo{}
	resp, err := httpClient.Get(fmt.Sprintf("%v%v", url, id))
	if err != nil {
		log.Printf("error %v", err)
		return err
	}
	//var result map[string]interface{}
	//_ = json.NewDecoder(resp.Body).Decode(&result)
	//fmt.Printf("For Id => %v, Title => %v\n", id, result["title"])
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error %v", err)
		return err
	}
	_ = json.Unmarshal(data, &todo)
	fmt.Printf("For Id => %v, Title => %v\n", id, todo.Title)
	return nil
}
