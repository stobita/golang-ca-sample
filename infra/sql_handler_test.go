package infra

import "testing"

func TestNewSqlHandler(t *testing.T) {
	sqlHandler := NewSqlHandler()
	if sqlHandler == nil {
		t.Error("NewSqlHandler error")
	}
}

func TestExecute(t *testing.T) {
	sqlHandler := NewSqlHandler()
	_, err := sqlHandler.Execute("CREATE TABLE test (id integer, name varchar(255))")
	_, err = sqlHandler.Execute("INSERT INTO test (id,name) VALUE (?,?)", 1, "sample")
	if err != nil {
		t.Log(err)
		t.Error("Execute error")
	}
}

func TestQuery(t *testing.T) {
	sqlHandler := NewSqlHandler()
	_, err := sqlHandler.Query("SELECT id, name FROM test")
	if err != nil {
		t.Log(err)
		t.Error("Query error")
	}
}
