package code

import(

	//"code/parsing"
	//	"github.com/bobkoffandrei/go-project-244/cmd/parsers"
	"testing"
	//    "code/formatters"
//	"errors"
)

func TestDiffPlain(t *testing.T) {
	Tests := []struct {
		name, path1, path2, style, want string
	}{
		{"PlainMain", "./testdata/fixture/file2.json", "./testdata/fixture/file1.json", "plain", "Property 'follow' was added with value: false\nProperty 'proxy' was added with value: '123.234.53.22'\nProperty 'timeout' was updated. From 20 to 50\nProperty 'verbose' was removed"},
		{"StylishMain", "./testdata/fixture/recFile1.json", "./testdata/fixture/recFile2.json", "stylish", `{
    common: {
      + follow: false
        setting1: Value 1
      - setting2: 200
      - setting3: true
      + setting3: null
      + setting4: blah blah
      + setting5: {
            key5: value5
        }
        setting6: {
            doge: {
              - wow: 
              + wow: so much
            }
            key: value
          + ops: vops
        }
    }
    group1: {
      - baz: bas
      + baz: bars
        foo: bar
      - nest: {
            key: value
        }
      + nest: str
    }
  - group2: {
        abc: 12345
        deep: {
            id: 45
        }
    }
  + group3: {
        deep: {
            id: {
                number: 45
            }
        }
        fee: 100500
    }
}`},
	
	}

	for _, test := range Tests {


	got, err := GenDiff(test.path1, test.path2, test.style)

	if err != nil {
		t.Errorf("Ошибка выполнения GenDiff: %v", err)
	}


	if got != test.want {
		t.Errorf("%s: got: \n%s, want: \n%s", test.name, got, test.want)
	}
	}

}


