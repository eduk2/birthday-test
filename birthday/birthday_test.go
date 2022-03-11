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

}

func TestIsLeapYear(t *testing.T) {
}

func TestJsonToPeople(t *testing.T) {
}
