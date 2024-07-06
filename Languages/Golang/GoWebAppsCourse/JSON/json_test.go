package main

import (
	"testing"
)

func Test_JSON(t *testing.T) {
	daddyLongLegs := Spider{10, "daddy long legs", "gray"}
	marshallJSON(daddyLongLegs)
}