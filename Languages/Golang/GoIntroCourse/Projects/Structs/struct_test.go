package main

import "testing"

func TestPersonStruct(t *testing.T) {
	newPerson := person{firstName: "trevor", lastName: "fagan"}

	if newPerson.firstName != "trevor" {
		t.Errorf("Expected first name of trevor, got %s", newPerson.firstName)
	}
}