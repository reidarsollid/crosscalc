package crisscross

import (
	"github.com/magiconair/properties/assert"
	"math"
	"testing"
)

var diagonal = [][]int{
	{1, 1, 1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1, 1, 1},
	{1, 1, 1, 1, 1, 1, 1, 1},
	{1, 2, 1, 1, 1, 1, 1, 1},
	{1, 2, 1, 1, 1, 1, 1, 1},
	{1, 2, 1, 1, 1, 1, 1, 1},
	{1, 2, 1, 1, 1, 1, 1, 1},
}

var rightCross = [][]int{
	{2, 2, 2, 2, 2, 2, 2},
	{2, 2, 2, 2, 2, 2, 2},
	{2, 2, 2, 7, 2, 2, 2},
	{2, 2, 2, 2, 7, 2, 2},
	{2, 6, 2, 2, 2, 7, 2},
	{2, 2, 6, 2, 2, 2, 7},
	{2, 2, 2, 6, 2, 2, 2},
}

var leftCross = [][]int{
	{2, 2, 2, 2, 2, 2, 2},
	{2, 2, 2, 2, 2, 2, 2},
	{2, 2, 2, 8, 2, 2, 2},
	{2, 2, 8, 2, 2, 7, 2},
	{2, 8, 2, 2, 7, 2, 2},
	{8, 2, 2, 7, 2, 2, 2},
	{2, 2, 7, 2, 2, 2, 2},
}

var row = [][]int{
	{0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0},
	{5, 5, 5, 5, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0},
}

var all = [][]int{
	{0, 0, 0, 0, 4, 4, 4, 4},
	{0, 0, 0, 0, 0, 4, 4, 0},
	{0, 0, 0, 3, 0, 0, 0, 0},
	{0, 0, 3, 6, 0, 0, 0, 8},
	{0, 3, 0, 0, 6, 0, 0, 8},
	{3, 0, 0, 0, 0, 6, 0, 8},
	{0, 0, 0, 0, 0, 0, 6, 8},
}

func TestProductDiagonal(t *testing.T) {
	expected := int(math.Pow(2, 4))
	actual := ProductDiagonal(diagonal...)
	assert.Equal(t, actual, expected)
}

func TestProductRightCross(t *testing.T) {
	expected := int(math.Pow(7, 4))
	actual := ProductRightCross(rightCross...)
	assert.Equal(t, actual, expected)
}

func TestProductLeftCross(t *testing.T) {
	expected := int(math.Pow(8, 4))
	actual := ProductLeftCross(leftCross...)
	assert.Equal(t, actual, expected)
}

func TestProductRow(t *testing.T) {
	expected := int(math.Pow(5, 4))
	actual := ProductRow(row...)
	assert.Equal(t, actual, expected)
}

func TestCrissCrossProducts(t *testing.T) {
	max := 0
	product := 0
	max, product = CrissCrossProducts(ProductRow, max, all...)
	assert.Equal(t, product, int(math.Pow(4, 4)))
	max, product = CrissCrossProducts(ProductLeftCross, max, all...)
	assert.Equal(t, product, int(math.Pow(3, 4)))
	max, product = CrissCrossProducts(ProductRightCross, max, all...)
	assert.Equal(t, product, int(math.Pow(6, 4)))
	max, product = CrissCrossProducts(ProductDiagonal, max, all...)
	assert.Equal(t, product, int(math.Pow(8, 4)))

}
