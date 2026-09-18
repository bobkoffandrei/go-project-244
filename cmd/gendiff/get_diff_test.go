package main

import (
	//"github.com/bobkoffandrei/go-project-244/code"
	"code/internal/parsing"
	"testing"

	//	    "code/formatters"
	//"github.com/stretchr/testify/assert"
	"code"
	"errors"
)

func TestDiffTest(t *testing.T) {
	Tests := []struct {
		name, path1, path2, want string
	}{
		{"noraml1", "../../testdata/fixture/file1.json", "../../testdata/fixture/file2.json", `{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}`},
		{"noraml2", "../../testdata/fixture/file2.json", "../../testdata/fixture/file3.json", `{
  - host: hexlet.io
  + host: hexlet.ru
  - timeout: 20
  + timeout: 30
  - verbose: true
  + verbose: false
}`},
		{"noraml3", "../../testdata/fixture/file4.json", "../../testdata/fixture/file1.json", `{
  + follow: false
    host: hexlet.io
  - port: 8080
  + proxy: 123.234.53.22
  - timeout: 20
  + timeout: 50
  - verbose: false
}`},
		{"oneempty", "../../testdata/fixture/empty.json", "../../testdata/fixture/file1.json", `{
  + follow: false
  + host: hexlet.io
  + proxy: 123.234.53.22
  + timeout: 50
}`}}

	for _, test := range Tests {

		got, err := code.GenDiff(test.path1, test.path2, "")

		if err != nil {
			t.Errorf("%s: Ошибка выполнения GenDiff: %v", test.name, err)
		}

		//got := formatters.FormatJSON(diffTree)

		if got != test.want {
			t.Errorf("%s: got: \n%s, want: \n%s", test.name, got, test.want)
		}
	}

}

func TestPathErrors(t *testing.T) {
	Tests := []struct {
		name, path1, path2 string
		wantErr            error
	}{
		{"wrong path1", "../..//file8.json", "../../testdata/fiadade/file2.json", parsing.ErrFileNotFound},
		{"wrong path2", "../../testdata/fixture/file2.json", "../../testdata/fi123123e/file2.json", parsing.ErrFileNotFound},
		{"wrong both", "../../tesdfgsdfge/file2.json", "../../testdata/fi123123e/file2.json", parsing.ErrFileNotFound},
	}

	for _, test := range Tests {

		t.Run(test.name, func(t *testing.T) {
			_, err := parsing.ParseFile(test.path1)

			if !errors.Is(err, test.wantErr) && err != nil {
				t.Fatalf("%s: ожидали ошибку ErrFileNotFound, получили: %v", test.name, err)
			}

			_, err = parsing.ParseFile(test.path2)

			if !errors.Is(err, test.wantErr) && err != nil {
				t.Errorf("%s: ожидали ошибку ErrFileNotFound, получили: %v", test.name, err)
			}
		})

	}

}

func TestOtherErrors(t *testing.T) {
	Tests := []struct {
		name, path1, path2 string
		wantErr            error
	}{
		{"NotJson", "../../testdata/fixture/file5.json", "../../testdata/fixture/file5.json", parsing.ErrParsingFile},
	}

	for _, test := range Tests {

		t.Run(test.name, func(t *testing.T) {
			_, err := parsing.ParseFile(test.path1)

			if !errors.Is(err, test.wantErr) && err != nil {
				t.Fatalf("%s: ожидали ошибку ErrParsingFile, получили: %v", test.name, err)
			}

			_, err = parsing.ParseFile(test.path2)

			if !errors.Is(err, test.wantErr) && err != nil {
				t.Errorf("%s: ожидали ошибку ErrParsingFile, получили: %v", test.name, err)
			}
		})

	}

}
