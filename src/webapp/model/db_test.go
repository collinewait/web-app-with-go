package model

type mockDB struct {
	lastQuery   string
	lastArgs    []interface{}
	returnedRow Row
}

func (db *mockDB) QueryRow(query string, args ...interface{}) Row {
	db.lastQuery = query
	db.lastArgs = args
	println("Printing test_db", db.lastArgs[1])
	return db.returnedRow
}

func (db *mockDB) Exec(query string, args ...interface{}) (Result, error) {
	return nil, nil
}

type mockRow struct{}

func (m mockRow) Scan(...interface{}) error {
	return nil
}
