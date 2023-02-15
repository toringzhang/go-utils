package basics

import (
	"reflect"
	"testing"
)

type Love struct {
	Name string `json:"name"`
	Long int    `json:"long"`
}

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Connect struct {
		Phone string `json:"phone"`
		Email string `json:"email"`
	} `json:"connect"`
	Love *Love `json:"love"`
}

func Test_TravelStruct(t *testing.T) {
	var user = &User{
		Name: "me",
		Age:  10,
		Connect: struct {
			Phone string "json:\"phone\""
			Email string "json:\"email\""
		}{Phone: "1244555", Email: "aaa@email.com"},
		Love: &Love{
			Name: "ball",
			Long: 3,
		},
	}
	value := reflect.ValueOf(user)
	TravelStruct(value, "user")
}
