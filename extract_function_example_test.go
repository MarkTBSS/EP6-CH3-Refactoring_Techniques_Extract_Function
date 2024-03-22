package main

import (
	"testing"
)

type Case struct {
	name string
	age  int
	want string
}

func TestGetPrice(t *testing.T) {
	testcase_0 := Case{name: "Free Ticket when age is 0", age: 0, want: "free"}
	testcase_1 := Case{name: "Free Ticket when age is lower 12", age: 11, want: "free"}
	testcase_2 := Case{name: "Ticket $15 when age is equal 12", age: 12, want: "$15"}
	testcase_3 := Case{name: "Ticket $15 when age is over 12", age: 13, want: "$15"}
	tests := []Case{testcase_0, testcase_1, testcase_2, testcase_3}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			person := Person{Age: testcase.age}
			got := person.getPrice()
			if got != testcase.want {
				t.Errorf("Price(%d) = %s; want %s", testcase.age, got, testcase.want)
			}
		})
	}
}
