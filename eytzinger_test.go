package eytzinger

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestEytzinger(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	want := []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15}
	got := Eytzinger(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestIsEytzinger(t *testing.T) {
	// Test cases
	testCases := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "Valid input",
			input:    []int{8, 4, 12, 2, 6, 10, 14, 1, 3, 5, 7, 9, 11, 13, 15},
			expected: true,
		},
		{
			name:     "Invalid input",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			expected: false,
		},
	}

	// Test loop
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the function
			result := IsEytzinger(tc.input)

			// Check the result
			if result != tc.expected {
				t.Errorf("Expected %v but got %v", tc.expected, result)
			}
		})
	}
}

func TestSearch(t *testing.T) {
	str1 := []string{"foo"}
	str2 := []string{"ab", "ca"}
	str3 := []string{"mo", "qo", "vo"}
	str4 := []string{"ab", "ad", "ca", "xy"}

	// slice with repeating elements
	strRepeats := []string{"ba", "ca", "da", "da", "da", "ka", "ma", "ma", "ta"}

	// slice with all element equal
	strSame := []string{"xx", "xx", "xx"}

	tests := []struct {
		data      []string
		target    string
		wantPos   int
		wantFound bool
	}{
		{[]string{}, "foo", 0, false},
		{[]string{}, "", 0, false},

		{str1, "foo", 0, true},
		{str1, "bar", 0, false},
		{str1, "zx", 1, false},

		{str2, "aa", 0, false},
		{str2, "ab", 0, true},
		{str2, "ad", 1, false},
		{str2, "ca", 1, true},
		{str2, "ra", 2, false},

		{str3, "bb", 0, false},
		{str3, "mo", 0, true},
		{str3, "nb", 1, false},
		{str3, "qo", 1, true},
		{str3, "tr", 2, false},
		{str3, "vo", 2, true},
		{str3, "xr", 3, false},

		{str4, "aa", 0, false},
		{str4, "ab", 0, true},
		{str4, "ac", 1, false},
		{str4, "ad", 1, true},
		{str4, "ax", 2, false},
		{str4, "ca", 2, true},
		{str4, "cc", 3, false},
		{str4, "dd", 3, false},
		{str4, "xy", 3, true},
		{str4, "zz", 4, false},

		{strRepeats, "da", 2, true},
		{strRepeats, "db", 5, false},
		{strRepeats, "ma", 6, true},
		{strRepeats, "mb", 8, false},

		{strSame, "xx", 0, true},
		{strSame, "ab", 0, false},
		{strSame, "zz", 3, false},
	}
	for _, tt := range tests {
		t.Run(tt.target, func(t *testing.T) {
			data := Eytzinger(tt.data)
			index := genIndex(len(tt.data))
			index = append(index, len(tt.data))

			// Test Search
			{
				pos, found := Search(data, tt.target)
				if index[pos] != tt.wantPos || found != tt.wantFound {
					t.Errorf("BinarySearch got (%v, %v), want (%v, %v)", index[pos], found, tt.wantPos, tt.wantFound)
				}
			}

			// Test SearchFunc
			{
				pos, found := SearchFunc(data, tt.target, strings.Compare)
				if index[pos] != tt.wantPos || found != tt.wantFound {
					t.Errorf("BinarySearch got (%v, %v), want (%v, %v)", index[pos], found, tt.wantPos, tt.wantFound)
				}
			}
		})
	}
}

func genIndex(n int) []int {
	index := make([]int, n)
	for i := range index {
		index[i] = i
	}
	return Eytzinger(index)
}

func TestSearchInts(t *testing.T) {
	data := []int{20, 30, 40, 50, 60, 70, 80, 90}
	index := genIndex(len(data))
	eytzinger := Eytzinger(data)
	tests := []struct {
		target    int
		wantPos   int
		wantFound bool
	}{
		{20, 0, true},
		{30, 1, true},
		{40, 2, true},
		{50, 3, true},
		{60, 4, true},
		{70, 5, true},
		{80, 6, true},
		{90, 7, true},

		{23, 1, false},
		{43, 3, false},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.target), func(t *testing.T) {
			{
				pos, found := Search(eytzinger, tt.target)
				if index[pos] != tt.wantPos || found != tt.wantFound {
					t.Errorf("BinarySearch got (%v, %v), want (%v, %v)", index[pos], found, tt.wantPos, tt.wantFound)
				}
			}
		})
	}
}
