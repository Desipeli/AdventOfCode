package main

import (
	"log"
	"reflect"
	"testing"
)

func TestIsSafe(t *testing.T) {
	type test struct {
		input []string
		want  bool
	}

	tests := []test{
		{
			input: []string{"1", "2", "4", "6", "9"},
			want:  true,
		},
		{
			input: []string{"4", "3", "2", "1"},
			want:  true,
		},
		{
			input: []string{"1", "2", "2", "3"},
			want:  false,
		},
		{
			input: []string{"1", "2", "6"},
			want:  false,
		},
		{
			input: []string{"6", "2", "1"},
			want:  false,
		},
		{
			input: []string{"3", "2", "1", "2"},
			want:  false,
		},
		{
			input: []string{"1", "2", "3", "1"},
			want:  false,
		},
	}

	for _, tc := range tests {
		got := IsSafe(tc.input)
		if !reflect.DeepEqual(got, tc.want) {
			log.Fatalf("expected %v got %v with input %v", tc.want, got, tc.input)
		}
	}
}
