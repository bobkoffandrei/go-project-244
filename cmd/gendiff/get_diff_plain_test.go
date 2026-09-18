package main

import (

	//	"code/parsing"
	//	"github.com/bobkoffandrei/go-project-244/cmd/parsers"
	"testing"
	//	    "code/formatters"
	"code"
	// "errors"
)

func TestDiffPlain(t *testing.T) {
	Tests := []struct {
		name, path1, path2, want string
	}{
		{"Plain1", "../../testdata/fixture/file2.json", "../../testdata/fixture/file1.json", "Property 'follow' was added with value: false\nProperty 'proxy' was added with value: '123.234.53.22'\nProperty 'timeout' was updated. From 20 to 50\nProperty 'verbose' was removed"},
		{"PlainRec1", "../../testdata/fixture/recFile1.json", "../../testdata/fixture/recFile2.json", "Property 'common.follow' was added with value: false\nProperty 'common.setting2' was removed\nProperty 'common.setting3' was updated. From true to null\nProperty 'common.setting4' was added with value: 'blah blah'\nProperty 'common.setting5' was added with value: [complex value]\nProperty 'common.setting6.doge.wow' was updated. From '' to 'so much'\nProperty 'common.setting6.ops' was added with value: 'vops'\nProperty 'group1.baz' was updated. From 'bas' to 'bars'\nProperty 'group1.nest' was updated. From [complex value] to 'str'\nProperty 'group2' was removed\nProperty 'group3' was added with value: [complex value]"},
	}

	for _, test := range Tests {

		got, err := code.GenDiff(test.path1, test.path2, "plain")

		if err != nil {
			t.Errorf("%s: Ошибка выполнения GenDiff: %v", test.name, err)
		}

		//got := formatters.FormatJSON(diffTree)

		if got != test.want {
			t.Errorf("%s: got: \n%s, want: \n%s", test.name, got, test.want)
		}
	}

}
