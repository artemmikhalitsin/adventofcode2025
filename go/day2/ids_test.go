package main

import (
	"testing"
)

func TestFindInvalidIds(t *testing.T) {
	tt := []struct {
		Range      Range
		InvalidIds []int
	}{
		{Range{10, 22}, []int{11, 22}},
		{Range{95, 115}, []int{99}},
		{Range{998, 1012}, []int{1010}},
		{Range{1188511880, 1188511890}, []int{1188511885}},
		{Range{222220, 222224}, []int{222222}},
		{Range{1698522, 1698528}, []int{}},
		{Range{446443, 446449}, []int{446446}},
		{Range{38593856, 38593862}, []int{38593859}},
		{Range{2121212118, 2121212124}, []int{}},
	}
	for _, tc := range tt {
		result := findInvalid(tc.Range, isInvalidV1)
		if !equalSlices(result, tc.InvalidIds) {
			t.Errorf("For range %v, expected %v, got %v", tc.Range, tc.InvalidIds, result)
		}
	}
}

func TestFindInvalidIdsV2(t *testing.T) {
	tt := []struct {
		Range      Range
		InvalidIds []int
	}{
		{Range{10, 22}, []int{11, 22}},
		{Range{95, 115}, []int{99, 111}},
		{Range{998, 1012}, []int{999, 1010}},
		{Range{1188511880, 1188511890}, []int{1188511885}},
		{Range{222220, 222224}, []int{222222}},
		{Range{1698522, 1698528}, []int{}},
		{Range{446443, 446449}, []int{446446}},
		{Range{565653, 565659}, []int{565656}},
		{Range{38593856, 38593862}, []int{38593859}},
		{Range{824824821, 824824827}, []int{824824824}},
		{Range{2121212118, 2121212124}, []int{2121212121}},
	}
	for _, tc := range tt {
		result := findInvalid(tc.Range, isInvalidV2)
		if !equalSlices(result, tc.InvalidIds) {
			t.Errorf("For range %v, expected %v, got %v", tc.Range, tc.InvalidIds, result)
		}
	}
}

func TestGetSumOfInvalidV1(t *testing.T) {
	ranges := []Range{
		{10, 22},
		{95, 115},
		{998, 1012},
		{1188511880, 1188511890},
		{222220, 222224},
		{1698522, 1698528},
		{446443, 446449},
		{38593856, 38593862},
		{2121212118, 2121212124},
	}

	result := GetSumOfInvalidV1(ranges)
	if result != 1227775554 {
		t.Errorf("Expected sum of invalid IDs to be 1227775554, got %d", result)
	}
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
