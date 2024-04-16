package utils

import (
	"fmt"
	"testing"
)

func TestArrayToMap(t *testing.T) {
	testArr := []string{"1", "2", "3", "4"}
	testMap := ArrayToMap(testArr, true)
	expectedMap := map[string]bool{"1": true, "2": true, "3": true, "4": true}
	for k, v := range expectedMap {
		if val, exists := testMap[k]; !exists || v != val {
			t.Error(fmt.Sprintf("For key %v: exists: %t, val: %v, expected: %v", k, exists, val, v))
			t.FailNow()
		}
	}
	for k, v := range testMap {
		if val, exists := expectedMap[k]; !exists || v != val {
			t.Error(fmt.Sprintf("For key %v: exists: %t, val: %v, expected: %v", k, exists, v, val))
			t.FailNow()
		}
	}
}
