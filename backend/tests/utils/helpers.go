package utils

import (
	"os"
	"swe-project/backend/datamgr"
	"testing"
)

func Delete_db(path string, t *testing.T) {
	// Close and delete temporary database
	temp, _ := datamgr.DB.DB()
	temp.Close()
	error := os.Remove("temp.db")
	if error != nil {
		t.Error("Failed to remove temporary db")
	}
}
