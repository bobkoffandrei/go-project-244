package main

import(
	//"github.com/bobkoffandrei/go-project-244/code"
	"code/parsers"
	"code/formatters"
	"testing"
	//"github.com/stretchr/testify/assert"
	"errors"
)

func TestDiffTestYaml(t *testing.T) {
	Tests := []struct {
		name, path1, path2, want string
	}{
		{"noraml1", "../../testdata/fixture/file1.yaml", "../../testdata/fixture/file2.yaml", `{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}`},
		{"noraml2", "../../testdata/fixture/file2.yaml", "../../testdata/fixture/file3.yaml", `{
  - host: hexlet.io
  + host: hexlet.ru
  - timeout: 20
  + timeout: 30
  - verbose: true
  + verbose: false
}`},
		{"noraml3", "../../testdata/fixture/file4.yaml", "../../testdata/fixture/file1.yaml", `{
  + follow: false
    host: hexlet.io
  - port: 8080
  + proxy: 123.234.53.22
  - timeout: 20
  + timeout: 50
  - verbose: false
}`},
		{"oneempty", "../../testdata/fixture/empty.yaml", "../../testdata/fixture/file1.yaml", `{
  + follow: false
  + host: hexlet.io
  + proxy: 123.234.53.22
  + timeout: 50
}`},
	}

	for _, test := range Tests {

	res1, err := parsers.ParseFile(test.path1)

		if err != nil {
		t.Errorf("%s: Ошибка парсинга файлов: %v", test.name, err)
	}


	res2, err := parsers.ParseFile(test.path2)

	if err != nil {
		t.Errorf("%s: Ошибка парсинга файлов: %v", test.name, err)
	}

	got := "{\n" +  formatters.FormatStylishWithDepth(genDiff(res1, res2), 0)   + "}"
	if got != test.want {
		t.Errorf("%s: got: \n%s, want: \n%s", test.name, got, test.want)
	}
	}

}

func TestPathErrorsYaml(t *testing.T) {
	Tests := []struct {
		name, path1, path2  string
		wantErr error
	}{
		{"wrong path1", "../..//file8.yaml", "../../testdata/fiadade/file2.yaml", parsers.ErrFileNotFound},
		{"wrong path2", "../../testdata/fixture/file2.yaml", "../../testdata/fi123123e/file2.yaml", parsers.ErrFileNotFound},
		{"wrong both", "../../tesdfgsdfge/file2.yaml", "../../testdata/fi123123e/file2.yaml", parsers.ErrFileNotFound},
	}


	for _, test := range Tests {

		t.Run(test.name, func(t *testing.T) {
	_, err := parsers.ParseFile(test.path1)

	        if !errors.Is(err, test.wantErr) && err != nil {
            t.Fatalf("%s: ожидали ошибку ErrFileNotFound, получили: %v", test.name, err)
        }

	_, err = parsers.ParseFile(test.path2)

        if !errors.Is(err, test.wantErr)  && err != nil {
            t.Errorf("%s: ожидали ошибку ErrFileNotFound, получили: %v", test.name, err)
        }
})

	}

}

func TestOtherErrorsYaml(t *testing.T) {
		Tests := []struct {
		name, path1, path2  string
		wantErr error
	}{
		{"Notyaml", "../../testdata/fixture/file5.yaml", "../../testdata/fixture/file5.yaml", parsers.ErrParsingFile},

	}

	for _, test := range Tests {

			t.Run(test.name, func(t *testing.T) {
	_, err := parsers.ParseFile(test.path1)

	        if !errors.Is(err, test.wantErr) && err != nil {
            t.Fatalf("%s: ожидали ошибку ErrparsersFile, получили: %v", test.name, err)
        }

	_, err = parsers.ParseFile(test.path2)

        if !errors.Is(err, test.wantErr)  && err != nil {
            t.Errorf("%s: ожидали ошибку ErrparsersFile, получили: %v", test.name, err)
        }
})

	}

}
