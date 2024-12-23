package main

import (
	"reflect"
	"testing"
)

func Test_Walk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct{ Name string }{"John"},
			[]string{"John"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"John", "London"},
			[]string{"John", "London"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"John", 33},
			[]string{"John"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got: %v, want: %v", got, test.ExpectedCalls)
			}
		})
	}
}
