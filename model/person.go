package model

type Community struct {
	Name string `json:"name"`
}

type Communities []Community

type Person struct {
	Name        string      `json:"name"`
	Age         uint8       `json:"age"`
	Communities Communities `json:"communities"`
}

type Persons []Person
