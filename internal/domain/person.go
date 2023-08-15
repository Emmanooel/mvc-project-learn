package domain

type Person struct {
	Id    int
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}
