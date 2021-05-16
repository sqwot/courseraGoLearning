package main

import (
	"reflect"
	"testing"
)

type testCase struct {
	Key     string
	User    *User
	IsError bool
}

func TestGetUser(t *testing.T) {
	cases := []testCase{
		testCase{"ok", &User{ID: 27}, false},
		testCase{"fail", nil, true},
		testCase{"not_exist", nil, true},
	}

	for caseNum, item := range cases {
		u, err := GetUser(item.Key)

		if item.IsError && err == nil {
			t.Errorf("[%d] expected error, got nil", caseNum)
		}
		if !item.IsError && err != nil {
			t.Errorf("[%d] expected error: %s", caseNum, err)
		}
		if !reflect.DeepEqual(u, item.User) {
			t.Errorf("[%d] wrong results: got %+v, expect %+v", caseNum, u, item.User)
		}

	}
}
