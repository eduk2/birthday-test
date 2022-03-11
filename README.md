### Coding Challenge Guidelines

Given a JSON list of people and their dates of birth, write a function to list out the people whose birthday is today.

If a person was born on Feb 29th, then their birthday during a non-leap year should be considered to be Feb 28th. For simplicity, you can consider a leap year to be any year divisble by 4.

Input sample:

```
[
	["Doe", "John", "1982/10/08"],
	["Wayne", "Bruce", "1965/01/30"],
	["Gaga", "Lady", "1986/03/28"],
	["Mark", "Curry", "1988/02/29"]
]
```

You can use whichever programming language you like. The assignment should take no more than an hour - if it's taking you longer than that, consider whether you're overcomplicating things.

### Evaluation Criteria

Must have:

* Tests to check if your code is correct, robust and complete
* Iterative development, backed by commit history
* Testing more scenarios than the input sample provided above. Think about edge cases.

Nice to have:

* Instructions on how to run the code
* TDD
* Modular, cohesive code with sensible separation of concerns
* No over-engineering (you don't need a web framework, database etc). Keep it as simple as possible, but no simpler :)


# SOLUTION

## About the project

A package was created and given a json file with last names, names and birthday date, return a list of people whose birthday is today.

## Layout

```tree
├── birthday
│   ├── birthday_test.go
│   ├── birthday.go
│   └── sample.json
├── go.mod
├── main.go
└── README.md
```

A brief description of the layout:

* `birthday` contains the main packages
* `birthday_test.go` tests of the main file of birthday package
* `birthday.go` the main file of birthday package
* `sample.json` json file example to test the package funcionality 
* `go.mod` the root of dependency management
* `main.go` the main file of the project where we can see the output checking the sample file in the birthday package
* `README.mod` this file, with Coding Challenge Guidelines, solution, instructions and about the project

## Getting started

### Instructions on how to run the code
 
Run the test:

	go test ./birthday/

Run the main file

	go run main.go

OutPut if you run the main file on `2021/11/14` with the current `sample.json`

	File opened successfully: ./birthday/sample.json
	Today is: 2021/11/14
	&birthday.people{birthday.person{LastName:"Muller", Name:"David", Birthday:"2000/11/14"}, birthday.person{LastName:"Smith", Name:"Will", Birthday:"2021/11/14"}}

### birthday package

- This has a public constructor called `NewBirthday`. You need to define a json file path. For example:

	NewBirthday("./birthday/sample.json")

This returns a struck birthday with the jsonFile used. And with peopleList, this is a []person and person is a struck defining the people inside the json file.

- This have other public function called `CkeckBirthday`. With year, month and day, it returns people (from json file defined in the constructor) whose birthday is the specificed day.

- And by other hand it has a function called `PeopleWhoseBirthdayIsToday`. It uses the funtion called `CkeckBirthday` to check the people whose birthday is today.

- Finally we have private functions:
	- `jsonToPeople` which converts from a json file to people
	- `isLeapYear` which says if a year is a leap-year
	- `check` to check errors and run panic function with the description