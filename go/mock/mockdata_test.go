package main

import "testing"

func TestMockData(t *testing.T) {
	t.Run("should return nil error when connect to DB", func(t *testing.T) {
		driverName := "mysql"
		dataSourceName := "bob:101@tcp(localhost:3306)/HelloWorld"

		db, _ := connectDB(driverName, dataSourceName)
		result := db.Ping()

		if result != nil {
			t.Errorf("Ping to server %s should return nil but got %q\n", dataSourceName, result)
		}
	})

	t.Run("should convert text to correct format", func(t *testing.T) {
		text := "4f3c4f5b-c5b4-46f8-ba2b-e7ab0d147aee"
		replaced := "-"
		replacewith := ""
		expect := "4f3c4f5bc5b446f8ba2be7ab0d147aee"

		result := REPLACEALL(text, replaced, replacewith)

		if expect != result {
			t.Errorf("Replace '%s' with '%s', should return %s but got %s\n", replaced, replacewith, expect, result)
		}
	})

	t.Run("should convert text to correct format", func(t *testing.T) {
		text := "4f3c4f5b-c5b4-46f8-ba2b-e7ab0d147aee"
		replaced := "-"
		replacewith := ""
		expect := "4f3c4f5bc5b446f8ba2be7ab0d147aee"

		result := REPLACEALL(text, replaced, replacewith)

		if expect != result {
			t.Errorf("Replace '%s' with '%s', should return %s but got %s\n", replaced, replacewith, expect, result)
		}
	})

	t.Run("should convert num to string with 0 lead", func(t *testing.T) {
		num := 444
		expect := "0444"

		result := formatDigit(num)

		if expect != result {
			t.Errorf("Format num to string, should return %s but got %s\n", expect, result)
		}
	})

	t.Run("should convert num to string with 0 lead 2", func(t *testing.T) {
		num := 1
		expect := "0001"

		result := formatDigit(num)

		if expect != result {
			t.Errorf("Format num to string, should return %s but got %s\n", expect, result)
		}
	})

}
