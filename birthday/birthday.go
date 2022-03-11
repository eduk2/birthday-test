package birthday

// Parse dates from json (year/month/day)
const layoutUK string = "2006/1/2"

type person struct {
	LastName   string `json:"lastName"`
	Name       string `json:"firstName"`
	jsonFile   string
	peopleList people
}

// Given a JSON list of people with their dates of birth,
// returns people whose birthday is today.
func (b *birthday) PeopleWhoseBirthdayIsToday() (peopleBirthdayTodayList *people) {

	return
}
