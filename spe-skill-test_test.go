package main

import (
	"reflect"
	"testing"
)

func TestSpeSkillTest_narcissistic(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		r    *SpeSkillTest
		args args
		want bool
	}{
		{
			name: "narcissistic(153)",
			r:    &SpeSkillTest{},
			args: args{
				number: 153,
			},
			want: true,
		},
		{
			name: "narcissistic(111)",
			r:    &SpeSkillTest{},
			args: args{
				number: 111,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SpeSkillTest{}
			if got := r.narcissistic(tt.args.number); got != tt.want {
				t.Errorf("SpeSkillTest.narcissistic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpeSkillTest_findOutlier(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		r    *SpeSkillTest
		args args
		want interface{}
	}{
		{
			name: "[2, 4, 0, 100, 4, 11, 2602, 36]",
			r:    &SpeSkillTest{},
			args: args{
				arr: []int{2, 4, 0, 100, 4, 11, 2602, 36},
			},
			want: int(11),
		},
		{
			name: "[160, 3, 1719, 19, 11, 13, -21]",
			r:    &SpeSkillTest{},
			args: args{
				arr: []int{160, 3, 1719, 19, 11, 13, -21},
			},
			want: int(160),
		},
		{
			name: "[11, 13, 15, 19, 9, 13, -21]",
			r:    &SpeSkillTest{},
			args: args{
				arr: []int{11, 13, 15, 19, 9, 13, -21},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SpeSkillTest{}
			if got := r.findOutlier(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SpeSkillTest.findOutlier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpeSkillTest_findNeedle(t *testing.T) {
	type args struct {
		stack  []string
		needle string
	}
	tests := []struct {
		name string
		r    *SpeSkillTest
		args args
		want int
	}{
		{
			name: `findNeedle(["red", "blue", "yellow", "black", "grey"], "blue")`,
			r:    &SpeSkillTest{},
			args: args{
				stack:  []string{"red", "blue", "yellow", "black", "grey"},
				needle: "blue",
			},
			want: 1,
		},
		{
			name: `findNeedle(["red", "blue", "yellow", "black", "grey"], "silver")`,
			r:    &SpeSkillTest{},
			args: args{
				stack:  []string{"red", "blue", "yellow", "black", "grey"},
				needle: "silver",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SpeSkillTest{}
			if got := r.findNeedle(tt.args.stack, tt.args.needle); got != tt.want {
				t.Errorf("SpeSkillTest.findNeedle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpeSkillTest_blueOcean(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		r    *SpeSkillTest
		args args
		want []int
	}{
		{
			name: "blueOcean([1,2,3,4,6,10], [1]): ",
			r:    &SpeSkillTest{},
			args: args{
				arr1: []int{1, 2, 3, 4, 6, 10},
				arr2: []int{1},
			},
			want: []int{2, 3, 4, 6, 10},
		},
		{
			name: "blueOcean([1,5,5,5,5,3], [5]): ",
			r:    &SpeSkillTest{},
			args: args{
				arr1: []int{1, 5, 5, 5, 5, 3},
				arr2: []int{5},
			},
			want: []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SpeSkillTest{}
			if got := r.blueOcean(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SpeSkillTest.blueOcean() = %v, want %v", got, tt.want)
			}
		})
	}
}
