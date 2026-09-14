package main

import(

	"code/parsing"
	//	"github.com/bobkoffandrei/go-project-244/cmd/parsers"
	"testing"
	    "code/formatters"
//	"errors" 
)

func TestDiffJson(t *testing.T) {
	Tests := []struct {
		name, path1, path2, want string
	}{
		{"Plain1", "../../testdata/fixture/file2.json", "../../testdata/fixture/file1.json",`{
  "diff": [
    {
      "type": "added",
      "key": "follow",
      "value": false
    },
    {
      "type": "unchanged",
      "key": "host",
      "value": "hexlet.io"
    },
    {
      "type": "added",
      "key": "proxy",
      "value": "123.234.53.22"
    },
    {
      "type": "changed",
      "key": "timeout",
      "value": 50,
      "oldValue": 20
    },
    {
      "type": "removed",
      "key": "verbose",
      "oldValue": true
    }
  ]
}`},
{"Plain1", "../../testdata/fixture/recFile1.json", "../../testdata/fixture/recFile2.json", `{
  "diff": [
    {
      "type": "nested",
      "key": "common",
      "children": [
        {
          "type": "added",
          "key": "follow",
          "value": false
        },
        {
          "type": "unchanged",
          "key": "setting1",
          "value": "Value 1"
        },
        {
          "type": "removed",
          "key": "setting2",
          "oldValue": 200
        },
        {
          "type": "changed",
          "key": "setting3",
          "oldValue": true
        },
        {
          "type": "added",
          "key": "setting4",
          "value": "blah blah"
        },
        {
          "type": "added",
          "key": "setting5",
          "value": {
            "key5": "value5"
          }
        },
        {
          "type": "nested",
          "key": "setting6",
          "children": [
            {
              "type": "nested",
              "key": "doge",
              "children": [
                {
                  "type": "changed",
                  "key": "wow",
                  "value": "so much",
                  "oldValue": ""
                }
              ]
            },
            {
              "type": "unchanged",
              "key": "key",
              "value": "value"
            },
            {
              "type": "added",
              "key": "ops",
              "value": "vops"
            }
          ]
        }
      ]
    },
    {
      "type": "nested",
      "key": "group1",
      "children": [
        {
          "type": "changed",
          "key": "baz",
          "value": "bars",
          "oldValue": "bas"
        },
        {
          "type": "unchanged",
          "key": "foo",
          "value": "bar"
        },
        {
          "type": "changed",
          "key": "nest",
          "value": "str",
          "oldValue": {
            "key": "value"
          }
        }
      ]
    },
    {
      "type": "removed",
      "key": "group2",
      "oldValue": {
        "abc": 12345,
        "deep": {
          "id": 45
        }
      }
    },
    {
      "type": "added",
      "key": "group3",
      "value": {
        "deep": {
          "id": {
            "number": 45
          }
        },
        "fee": 100500
      }
    }
  ]
}`},
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

		got := formatters.FormatJSON(diffTree)

	if got != test.want {
		t.Errorf("%s: got: \n%s, want: \n%s", test.name, got, test.want)
	}
	}

}


