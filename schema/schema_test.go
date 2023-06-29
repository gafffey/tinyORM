package schema

import (
	"github.com/gafffey/tinyorm/dialect"
	"testing"
)

type User struct {
	Id   int    `tinyorm:"id"`
	Name string `tinyorm:"name"`
}

var TestDial, _ = dialect.GetDialect("mysql")

func TestParser(t *testing.T) {
	schema := Parse(&User{}, TestDial)
	if schema.Name != "User" || len(schema.Fields) != 2 {
		t.Fatal("failed to parse User struct")
	}

	if schema.GetField("Id").Name != "Id" ||
		schema.GetField("Id").Type != "int" ||
		schema.GetField("Id").Tag != "id" {
		t.Fatal("failed to parse Id field")
	}

	if schema.GetField("Name").Name != "Name" ||
		schema.GetField("Name").Type != "varchar(255)" ||
		schema.GetField("Name").Tag != "name" {
		t.Fatal("failed to parse Name field")
	}
}

func TestSchema_RecordValues(t *testing.T) {
	schema := Parse(&User{}, TestDial)
	values := schema.RecordValues(&User{1, "Tom"})

	id := values[0].(int)
	name := values[1].(string)

	if id != 1 || name != "Tom" {
		t.Fatal("failed to get values")
	}
}
