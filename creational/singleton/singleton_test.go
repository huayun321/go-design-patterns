package singleton

import (
	"reflect"
	"testing"
)

func TestGetInstance(t *testing.T) {
	w := new(singleton)
	tests := []struct {
		name string
		want *singleton
	}{
		{name: "getInstance", want: w},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInstance(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInstance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_singleton_AddOne(t *testing.T) {
	tests := []struct {
		name string
		s    *singleton
		want int
	}{
		{name: "addOne1", s: GetInstance(), want: 1},
		{name: "addOne2", s: GetInstance(), want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.AddOne(); got != tt.want {
				t.Errorf("singleton.AddOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
