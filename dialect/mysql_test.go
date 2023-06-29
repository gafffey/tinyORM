package dialect

import (
	"reflect"
	"testing"
)

func TestMysql_DataTypeOf(t *testing.T) {
	dial := &mysql{}
	cases := []struct {
		Value    interface{}
		Expected string
	}{
		{1, "int"},
		{"Tom", "varchar(255)"},
	}

	for _, c := range cases {
		if typ := dial.DataTypeOf(reflect.ValueOf(c.Value)); typ != c.Expected {
			t.Fatal("expect", c.Expected, "but got", typ)
		}
	}
}
