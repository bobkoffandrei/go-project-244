package main

import(

	///"code/parsing"
	//	"github.com/bobkoffandrei/go-project-244/cmd/parsers"
	"testing"
	//    "code/formatters"
//	"errors"
"code"
)

func TestDiffRecTest(t *testing.T) {
	Tests := []struct {
		name, path1, path2, want string
	}{
		{"noraml1", "../../testdata/fixture/recFile1.json", "../../testdata/fixture/recFile2.json", `{
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



