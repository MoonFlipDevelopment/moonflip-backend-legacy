package main

import "gorm.io/gorm"

func ExistsInDatabase(table string, database, subQuery *gorm.DB) *gorm.DB {
	// GORM doesn't correctly compile Select statements like this, it expects a SELECT * FROM TABLE;
	// and always appends a FROM. So excluding the table results in an invalid syntax.
	// for ease, we just include the table even though it's unnecessary
	return database.Table(table).Select("exists(?) as `exists`", subQuery)
}
