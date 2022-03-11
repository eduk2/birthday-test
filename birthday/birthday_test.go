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

func TestPeopleWhoseBirthdayIs(t *testing.T) {
}

func TestIsLeapYear(t *testing.T) {
}

func TestJsonToPeople(t *testing.T) {
}
