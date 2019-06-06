package util

import (
	"reflect"
	"testing"
)

func TestParseImageStreamImageName(t *testing.T) {
	tests := map[string]struct {
		input        string
		expectedRepo string
		expectedId   string
		expectError  bool
	}{
		"empty string": {
			input:       "",
			expectError: true,
		},
		"one part": {
			input:       "a",
			expectError: true,
		},
		"more than 2 parts": {
			input:       "a@b@c",
			expectError: true,
		},
		"empty name part": {
			input:       "@id",
			expectError: true,
		},
		"empty id part": {
			input:       "name@",
			expectError: true,
		},
		"valid input": {
			input:        "repo@id",
			expectedRepo: "repo",
			expectedId:   "id",
			expectError:  false,
		},
	}

	for name, test := range tests {
		repo, id, err := ParseImageStreamImageName(test.input)
		didError := err != nil
		if e, a := test.expectError, didError; e != a {
			t.Errorf("%s: expected error=%t, got=%t: %s", name, e, a, err)
			continue
		}
		if test.expectError {
			continue
		}
		if e, a := test.expectedRepo, repo; e != a {
			t.Errorf("%s: repo: expected %q, got %q", name, e, a)
			continue
		}
		if e, a := test.expectedId, id; e != a {
			t.Errorf("%s: id: expected %q, got %q", name, e, a)
			continue
		}
	}
}

func TestParseImageStreamTagName(t *testing.T) {
	tests := map[string]struct {
		id           string
		expectedName string
		expectedTag  string
		expectError  bool
	}{
		"empty id": {
			id:          "",
			expectError: true,
		},
		"missing semicolon": {
			id:          "hello",
			expectError: true,
		},
		"too many semicolons": {
			id:          "a:b:c",
			expectError: true,
		},
		"empty name": {
			id:          ":tag",
			expectError: true,
		},
		"empty tag": {
			id:          "name",
			expectError: true,
		},
		"happy path": {
			id:           "name:tag",
			expectError:  false,
			expectedName: "name",
			expectedTag:  "tag",
		},
	}

	for description, testCase := range tests {
		name, tag, err := ParseImageStreamTagName(testCase.id)
		gotError := err != nil
		if e, a := testCase.expectError, gotError; e != a {
			t.Fatalf("%s: expected err: %t, got: %t: %s", description, e, a, err)
		}
		if err != nil {
			continue
		}
		if e, a := testCase.expectedName, name; e != a {
			t.Errorf("%s: name: expected %q, got %q", description, e, a)
		}
		if e, a := testCase.expectedTag, tag; e != a {
			t.Errorf("%s: tag: expected %q, got %q", description, e, a)
		}
	}
}
