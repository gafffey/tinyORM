package session

import (
	"database/sql"
	"fmt"
	"github.com/gafffey/tinyorm/dialect"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"testing"
)

var TestDB *sql.DB

func TestMain(m *testing.M) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", "root", "root", "localhost:3306", "example")
	TestDB, _ = sql.Open("mysql", dsn)
	code := m.Run()
	_ = TestDB.Close()
	os.Exit(code)
}

func NewSession() *Session {
	d, _ := dialect.GetDialect("mysql")
	return New(TestDB, d)
}

func TestSession_Exec(t *testing.T) {
	s := NewSession()
	_, _ = s.Raw(CreateTable).Exec()
	result, _ := s.Raw(InsertUser, "Tom", "Sam").Exec()
	if count, err := result.RowsAffected(); err != nil || count != 2 {
		t.Fatal("expect 2, but got", count)
	}
}

const (
	CreateTable = `
CREATE TABLE IF NOT EXISTS User(
   Id INT UNSIGNED AUTO_INCREMENT,
   Name VARCHAR(100) NOT NULL,
   PRIMARY KEY (id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
`
	InsertUser = "INSERT INTO User(`name`) VALUES (?), (?);"
)
