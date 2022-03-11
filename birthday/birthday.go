package birthday

// Parse dates from json (year/month/day)
const layoutUK string = "2006/1/2"

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

// Given a JSON list of people with their dates of birth,
// returns people whose birthday is today.
func (b *birthday) PeopleWhoseBirthdayIsToday() (peopleBirthdayTodayList *people) {

	return
}

// Given a JSON list of people with their dates of birth
// and given a specific day (year, month and day)
// Returns people whose birthday is the specific day
func (b *birthday) CkeckBirthday(year, month, day int) *people {

	var peopleBirthdayList = people{}

	return &peopleBirthdayList
}
