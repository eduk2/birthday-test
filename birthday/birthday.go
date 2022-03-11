package birthday

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const (
	// Special day
	feb29thMonth, feb29thDay int = 2, 29
	// Parse dates from json (year/month/day)
	layoutUK string = "2006/1/2"
)

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

	date := time.Now().Format(layoutUK)
	fmt.Println("Today is:", date)
	today, _ := time.Parse(layoutUK, date)

	monthToday := int(today.Month())
	dayToday := today.Day()
	yearToday := today.Year()

	peopleBirthdayTodayList = b.CkeckBirthday(yearToday, monthToday, dayToday)

	return
}

// Given a JSON list of people with their dates of birth
// and given a specific day (year, month and day)
// Returns people whose birthday is the specific day
func (b *birthday) CkeckBirthday(year, month, day int) *people {

	var peopleBirthdayList = people{}
	var isCheckingALeapYear bool = isLeapYear(year)

	for _, person := range b.peopleList {

		birthday, err := time.Parse(layoutUK, person.Birthday)
		check(err)

		monthPerson := int(birthday.Month())
		dayPerson := birthday.Day()

		//if it is not checking a leap year and the person was born on XXXX/02/29
		// then their birthday is a day before it: XXXX/02/28
		if !isCheckingALeapYear && monthPerson == feb29thMonth && dayPerson == feb29thDay {
			dayPerson--
		}

		if month == monthPerson && day == dayPerson {
			peopleBirthdayList = append(peopleBirthdayList, person)
		}

	}

	return &peopleBirthdayList
}

// Check if a year is a leap-year.
// For simplicity, it was considered a leap-year is any year divisible by 4.
func isLeapYear(year int) bool {
	var leapYear bool = false
	if year%4 == 0 {
		leapYear = true
	}
	return leapYear
}

// Convert from a json file to people
func jsonToPeople(jsonFile string) *people {

	var peopleList people
	peopleJson, err := os.ReadFile(jsonFile)
	check(err)
	fmt.Println("File opened successfully:", jsonFile)

	err = json.Unmarshal(peopleJson, &peopleList)
	check(err)

	return &peopleList
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
