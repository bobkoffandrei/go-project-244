package main

import(
	//"github.com/bobkoffandrei/go-project-244/code"
	"github.com/bobkoffandrei/go-project-244/cmd/parsing"
	"testing"
	//"github.com/stretchr/testify/assert"
	"errors"
)

func TestDiffTest(t *testing.T) {
	Tests := []struct {
		name, path1, path2, want string
	}{
		{"noraml1", "../../testdata/fixture/file1.json", "../../testdata/fixture/file2.json", "- follow: false\n  host: hexlet.io\n- proxy: 123.234.53.22\n- timeout: 50\n+ timeout: 20\n+ verbose: true\n"},
		{"noraml2", "../../testdata/fixture/file2.json", "../../testdata/fixture/file3.json", "- host: hexlet.io\n+ host: hexlet.ru\n- timeout: 20\n+ timeout: 30\n- verbose: true\n+ verbose: false\n"},
		{"noraml3", "../../testdata/fixture/file4.json", "../../testdata/fixture/file1.json", "  host: hexlet.io\n- port: 8080\n- timeout: 20\n+ timeout: 50\n- verbose: false\n+ follow: false\n+ proxy: 123.234.53.22\n"},
		{"oneempty", "../../testdata/fixture/empty.json", "../../testdata/fixture/file1.json", "+ follow: false\n+ host: hexlet.io\n+ proxy: 123.234.53.22\n+ timeout: 50\n"},
	}

	for _, test := range Tests {

	res1, err := parsing.ParseFile(test.path1)

		if err != nil {
		t.Errorf("%s: Ошибка парсинга файлов: %v", test.name, err)
	}


	res2, err := parsing.ParseFile(test.path2)

	if err != nil {
		t.Errorf("%s: Ошибка парсинга файлов: %v", test.name, err)
	}

	got := genDiff(res1, res2)
	if got != test.want {
		t.Errorf("%s: got: \n%s, want: \n%s", test.name, got, test.want)
	}
	}

}

func TestPathErrors(t *testing.T) {
	Tests := []struct {
		name, path1, path2  string
		wantErr error
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

        if !errors.Is(err, test.wantErr)  && err != nil {
            t.Errorf("%s: ожидали ошибку ErrFileNotFound, получили: %v", test.name, err)
        }
})

	}

}

func TestOtherErrors(t *testing.T) {
		Tests := []struct {
		name, path1, path2  string
		wantErr error
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

        if !errors.Is(err, test.wantErr)  && err != nil {
            t.Errorf("%s: ожидали ошибку ErrParsingFile, получили: %v", test.name, err)
        }
})

	}

}
