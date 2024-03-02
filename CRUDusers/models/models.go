package models

type User struct {
	Firstname  string `json:"firstname" bson:"firstname"`
	Lastname   string `json:"lastname" bson:"lastname"`
	Department string `json:"department" bson:"department"`
	Age        string `json:"age" bson:"age"`
	Tel        string `json:"tel" bson:"tel"`
}
