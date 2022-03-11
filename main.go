package main

import (
	sb "bdtest/birthday"
	"fmt"
)

func main() {

	birthdays := sb.NewBirthday("./birthday/sample.json")
	fmt.Printf("%#v\n", birthdays.PeopleWhoseBirthdayIsToday())

}
