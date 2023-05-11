package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumBuyProduct(t *testing.T){
	t.Run("Test case1",func(t *testing.T){
	expected :=3
	actual :=MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000})
	assert.Equal(t,expected,actual,"hasil tidak sama")
})
t.Run("Test case2",func(t *testing.T){
	expected :=4
	actual :=MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000})
	assert.Equal(t,expected,actual,"hasil tidak sama")
})
t.Run("Test case3",func(t *testing.T){
	expected :=4
	actual :=MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000})
	assert.Equal(t,expected,actual,"hasil tidak sama")
})
t.Run("Test case4",func(t *testing.T){
	expected :=1
	MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000})
	assert.Equal(t,expected,actual,"hasil tidak sama")
})
t.Run("Test case4",func(t *testing.T){
	expected :=0
	actual :=MaximumBuyProduct(0, []int{10000, 30000})
	assert.Equal(t,expected,actual,"hasil tidak sama")
})
}