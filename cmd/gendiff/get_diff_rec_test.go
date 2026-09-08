package main

import(

	"github.com/bobkoffandrei/go-project-244/code/parsing"
	//	"github.com/bobkoffandrei/go-project-244/cmd/parsers"
	"testing"
	    "github.com/bobkoffandrei/go-project-244/code/formatters"
//	"errors"
)

func TestDiffRecTest(t *testing.T) {
	Tests := []struct {
		name, path1, path2, want string
	}{
		{"noraml1", "../../testdata/fixture/recFile1.json", "../../testdata/fixture/recFile2.json", "{\n  common: {\n    + follow: false\n      setting1: Value 1\n    - setting2: 200\n    - setting3: true\n    + setting3: <nil>\n    + setting4: blah blah\n    + setting5: {\n        key5: value5\n      }\n      setting6: {\n          doge: {\n            - wow: \n            + wow: so much\n          }\n          key: value\n        + ops: vops\n      }\n  }\n  group1: {\n    - baz: bas\n    + baz: bars\n      foo: bar\n    - nest: {\n        key: value\n      }\n    + nest: str\n  }\n- group2: {\n    abc: 12345\n    deep: {\n      id: 45\n    }\n  }\n+ group3: {\n    deep: {\n      id: {\n        number: 45\n      }\n    }\n    fee: 100500\n  }\n}"},
	
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

		got := "{\n" +  formatters.FormatStylishWithDepth(genDiff(res1, res2), 0)   + "}"

	if got != test.want {
		t.Errorf("%s: got: \n%s, want: \n%s", test.name, got, test.want)
	}
	}

}


