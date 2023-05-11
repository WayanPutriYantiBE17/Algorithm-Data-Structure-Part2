package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMindAndMax(t *testing.T){
	t.Run("Test case1",func(t *testing.T){
	expected :="min :-2 index 3 max:8 index:5"
	actual :=FindMinAndMax([]int{5,7,4,-2,-1,8})
	assert.Equal(t,expected,actual,"hasil tidak sama")
})
t.Run("Test case2",func(t *testing.T){
	expected :="min :-5 index 1 max:22 index:3"
	actual :=FindMinAndMax([]int{5,-5,-4,22,7,7})
	assert.Equal(t,expected,actual,"hasil tidak sama")
})
t.Run("Test case3",func(t *testing.T){
	expected :="min :-21 index 4 max:9 index:2"
	actual :=FindMinAndMax([]int{4,3,9,4,-21,7})
	assert.Equal(t,expected,actual,"hasil tidak sama")
})
t.Run("Test case4",func(t *testing.T){
	expected :="min :-1 index 0 max:18 index:5"
	actual :=FindMinAndMax([]int{-1,5,6,4,2,18})
	assert.Equal(t,expected,actual,"hasil tidak sama")
})
t.Run("Test case4",func(t *testing.T){
	expected :="min :-20 index 5 max:7 index:4"
	actual :=FindMinAndMax([]int{-2,5,-7,4,7,-20})
	assert.Equal(t,expected,actual,"hasil tidak sama")
})
}