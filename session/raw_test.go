package session

import (
	"database/sql"
	"fmt"
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
	return New(TestDB)
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
   id INT UNSIGNED AUTO_INCREMENT,
   name VARCHAR(100) NOT NULL,
   PRIMARY KEY (id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
`
	InsertUser = "INSERT INTO User(`name`) VALUES (?), (?);"
)
