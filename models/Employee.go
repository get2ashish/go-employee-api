package models

type Employee struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Salary  int    `json:"salary"`
}
