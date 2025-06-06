package web_test

import (
	"encoding/json"
	"testing"

	"github.com/fengdotdev/golibs-funcs/v0/web"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestRemoteRes_Json(t *testing.T) {

	type Todo struct {
		UserId    int    `json:"userId"`
		Id        int    `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	url := "https://jsonplaceholder.typicode.com/todos/1"
	data, err := web.GetRemoteResource(url)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(data))

	todo := Todo{}
	err = json.Unmarshal(data, &todo)
	assert.Nil(t, err)
	assert.Equal(t, 1, todo.UserId)
	assert.Equal(t, 1, todo.Id)
}

func TestRemoteRes_Json2(t *testing.T) {

	type Todo struct {
		UserId    int    `json:"userId"`
		Id        int    `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	url := "https://jsonplaceholder.typicode.com/todos/2"
	data, err := web.GetRemoteResource(url)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(data))

	todo := Todo{}
	err = json.Unmarshal(data, &todo)
	assert.Nil(t, err)
	assert.Equal(t, 1, todo.UserId)
	assert.Equal(t, 2, todo.Id)
}

func TestRemoteRes_JPG(t *testing.T) {

	url := "https://picsum.photos/200/300"
	data, err := web.GetRemoteResource(url)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(data))

	//TODO - check if it is a valid image

}

func TestRemoteRes_404(t *testing.T) {

	url := "https://www.google.cl/microsoftpwned.jpg"
	data, err := web.GetRemoteResource(url)

	assert.NotNil(t, err)
	assert.Equal(t, 0, len(data))
}
