package birthday

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestPeopleWhoseBirthdayIsToday(t *testing.T) {

	date := time.Now().Format(layoutUK)
	today := strings.Split(date, "/")
	monthToday := today[1]
	dayToday := today[2]

	var birthdaysSample = &birthday{
		peopleList: people{
			person{
				"Today", "Always", "2001/" + monthToday + "/" + dayToday,
			},
			person{
				"Today 2", "Always", "2018/" + monthToday + "/" + dayToday,
			},
			person{
				"Lee", "David", "1990/05/2",
			},
			person{
				"Lopez", "Jennifer", "1970/10/29",
			},
			person{
				"Muller", "David", "1995/5/2",
			},
		},
	}

	birthdaysExpected := &people{
		person{
			"Today", "Always", "2001/" + monthToday + "/" + dayToday,
		},
		person{
			"Today 2", "Always", "2018/" + monthToday + "/" + dayToday,
		},
	}

	peopleChecked := birthdaysSample.PeopleWhoseBirthdayIsToday()

	if !reflect.DeepEqual(peopleChecked, birthdaysExpected) {
		t.Errorf("%+v did not equal to %+v", peopleChecked, birthdaysExpected)
	}

}

func TestCkeckBirthday(t *testing.T) {

	var birthdaysSample = &birthday{
		peopleList: people{
			person{
				"Smith", "George", "2001/12/24",
			},
			person{
				"Lady", "Gaga", "2000/05/02",
			},
			person{
				"Lee", "David", "1990/05/2",
			},
			person{
				"Lopez", "Jennifer", "1970/10/29",
			},
			person{
				"Muller", "David", "1995/5/2",
			},
			person{
				"Leal", "Robert", "1995/02/27",
			},
			person{
				"Robles", "Rosa", "1960/02/28",
			},
			person{
				"Andersen", "Jack", "2000/02/29",
			},
		},
	}

	birthdaysExpected := &people{
		person{
			"Lady", "Gaga", "2000/05/02",
		},
		person{
			"Lee", "David", "1990/05/2",
		},
		person{
			"Muller", "David", "1995/5/2",
		},
	}

	peopleChecked := birthdaysSample.CkeckBirthday(2021, 5, 2)
	if !reflect.DeepEqual(peopleChecked, birthdaysExpected) {
		t.Errorf("%+v did not equal to %+v", peopleChecked, birthdaysExpected)
	}

	//More cases with the critical day in not leap year
	notLeapYearBirthdaysExpected := &people{
		person{
			"Robles", "Rosa", "1960/02/28",
		},
		person{
			"Andersen", "Jack", "2000/02/29",
		},
	}
	leapYearChecked := birthdaysSample.CkeckBirthday(2001, 2, 28)
	if !reflect.DeepEqual(leapYearChecked, notLeapYearBirthdaysExpected) {
		t.Errorf("%+v did not equal to %+v", leapYearChecked, notLeapYearBirthdaysExpected)
	}

	//More cases with the critical day in leap year
	leapYearBirthdaysExpected := &people{
		person{
			"Robles", "Rosa", "1960/02/28",
		},
	}
	leapYearChecked = birthdaysSample.CkeckBirthday(2012, 2, 28)
	if !reflect.DeepEqual(leapYearChecked, leapYearBirthdaysExpected) {
		t.Errorf("%+v did not equal to %+v", leapYearChecked, leapYearBirthdaysExpected)
	}

}

func TestIsLeapYear(t *testing.T) {
	yearsListToTest := []struct {
		searching int
		expected  bool
	}{
		{2000, true},
		{2010, false},
		{2020, true},
		{2021, false},
		{2024, true},
	}

	for _, year := range yearsListToTest {

		yearSearchedIsLeapYear := isLeapYear(year.searching)
		if yearSearchedIsLeapYear != year.expected {
			t.Errorf("With %v: %+v and it was expecting %+v", year.searching, yearSearchedIsLeapYear, year.expected)
		}

	}

}

func TestJsonToPeople(t *testing.T) {

	peopleList := jsonToPeople("sample.json")

	if len(*peopleList) == 0 {
		t.Fatal("Sample file doesn't have any person")
	}

}
