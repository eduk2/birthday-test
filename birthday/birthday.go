package birthday

type person struct {
	LastName string `json:"lastName"`
	Name     string `json:"firstName"`
	Birthday string `json:"Birthday"`
}
type people []person

type birthday struct {
	jsonFile   string
	peopleList people
}
