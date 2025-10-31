package main

import (
    "reflect"
    "testing"
)

func TestWalk(t *testing.T) {

    type Profile struct {
        Age  int
        City string
    }

    type Person struct {
        Name    string
        Profile Profile
    }

    type Person2 struct {
        Name    string
        Profile *Profile
    }

    cases := []struct {
        Name          string
        Input         interface{}
        ExpectedCalls []string
    }{
        {
            Name: "struct with one string filed",
            Input: struct {
                Name string
            }{"Chris"},
            ExpectedCalls: []string{"Chris"},
        },
        {
            Name: "struct with two string fileds",
            Input: struct {
                Name string
                City string
            }{"Chris", "London"},
            ExpectedCalls: []string{"Chris", "London"},
        },
        {
            Name: "struct with non string field",
            Input: struct {
                Name string
                Age  int
            }{"Chris", 33},
            ExpectedCalls: []string{"Chris"},
        },
        {
            Name: "struct with nested fields",
            Input: Person{
                "Chris",
                Profile{33, "London"},
            },

            ExpectedCalls: []string{"Chris", "London"},
        },
        {
            Name: "struct with pointer",
            Input: &Person{
                "Chris", Profile{33, "London"},
            },
            ExpectedCalls: []string{"Chris", "London"},
        },
        {
            Name: "struct with nested pointer",
            Input: &Person2{
                "Chris", &Profile{33, "London"},
            },
            ExpectedCalls: []string{"Chris", "London"},
        },
    }

    for _, test := range cases {
        t.Run(test.Name, func(t *testing.T) {
            var got []string
            walk(test.Input, func(input string) {
                got = append(got, input)
            })
            if !reflect.DeepEqual(got, test.ExpectedCalls) {
                t.Errorf("got %+v, want %+v", got, test.ExpectedCalls)
            }
        })
    }
}
