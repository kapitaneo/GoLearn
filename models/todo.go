package models

//Todo model for return
type Todo struct {
	Name string
	Done bool
}

//Get return Todo
func Get() []Todo {
	todos := []Todo{
		{"Learn Go", true},
		{"Create site", true},
		{"Profit", false},
	}
	return todos
}
