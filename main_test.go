package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetRandInt(t *testing.T) {
	i := getRandInt()
	if reflect.TypeOf(i).String() != "int" {
		t.Fatal(i)
	}
}

func TestFileExists(t *testing.T) {
	e := fileExists("my-name")
	if !e {
		t.Fatal(e)
	}
}

func TestGetNewName(t *testing.T) {
	myname, err := getNewName()
	if err == nil && reflect.TypeOf(myname).String() != "string" {
		t.Fatal(myname)
	}
}

func TestGetName(t *testing.T) {
	myname, err := getName()
	fmt.Printf("myname: %s\n", myname)
	if err == nil && reflect.TypeOf(myname).String() != "string" {
		t.Fatal(myname)
	}
}

func TestIndex(t *testing.T) {
	t.Skip("Skpping index for now")
}

func TestMain(t *testing.T) {
	t.Skip("Skpping main for now")
}
