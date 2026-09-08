package main

import(

	"github.com/bobkoffandrei/go-project-244/code/parsing"
	//	"github.com/bobkoffandrei/go-project-244/cmd/parsers"
	"testing"
	    "github.com/bobkoffandrei/go-project-244/code/formatters"
//	"errors"
)

func TestDiffPlain(t *testing.T) {
	Tests := []struct {
		name, path1, path2, want string
	}{
		{"Plain1", "../../testdata/fixture/file2.json", "../../testdata/fixture/file1.json", "Property 'follow' was added with value: false\nProperty 'proxy' was added with value: '123.234.53.22'\nProperty 'timeout' was updated. From 20 to 50\nProperty 'verbose' was removed\n"},
		{"PlainRec1", "../../testdata/fixture/recFile1.json", "../../testdata/fixture/recFile2.json", "Property 'common.follow' was added with value: false\nProperty 'common.setting2' was removed\nProperty 'common.setting3' was updated. From true to null\nProperty 'common.setting4' was added with value: 'blah blah'\nProperty 'common.setting5' was added with value: [complex value]\nProperty 'common.setting6.doge.wow' was updated. From '' to 'so much'\nProperty 'common.setting6.ops' was added with value: 'vops'\nProperty 'group1.baz' was updated. From 'bas' to 'bars'\nProperty 'group1.nest' was updated. From [complex value] to 'str'\nProperty 'group2' was removed\nProperty 'group3' was added with value: [complex value]\n"},
	
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

	diffTree := genDiff(res1, res2)

		got := formatters.FormatPlain(diffTree)

	if got != test.want {
		t.Errorf("%s: got: \n%s, want: \n%s", test.name, got, test.want)
	}
	}

}


