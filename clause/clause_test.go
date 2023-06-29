package clause

import (
	"reflect"
	"testing"
)

func TestSelect(t *testing.T) {
	var clause Clause

	clause.Set(LIMIT, 1)
	clause.Set(SELECT, "User", []string{"*"})
	clause.Set(WHERE, "Name = ?", "Tom")
	clause.Set(ORDERBY, "Id DESC")

	sql, vars := clause.Build(SELECT, WHERE, ORDERBY, LIMIT)
	t.Log(sql, vars)

	if sql != "SELECT * FROM User WHERE Name = ? ORDER BY Id DESC LIMIT ?" {
		t.Fatal("failed to build SQL")
	}

	if !reflect.DeepEqual(vars, []interface{}{"Tom", 1}) {
		t.Fatal("failed to build SQLVars")
	}
}
