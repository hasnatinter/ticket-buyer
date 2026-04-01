package helpers

import (
	"app/internal/conn"
	"os"
	"os/exec"
	"testing"

	"gorm.io/gorm"
)

func SetUp(t *testing.T) *gorm.DB {
	t.Helper()

	os.Setenv("DB_HOST", "test_db")
	err := exec.Command("/usr/src/app/bin/migrate_test", "up").Run()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		err = exec.Command("/usr/src/app/bin/migrate_test", "reset").Run()
	})

	db := conn.ConnectDb()
	return db
}
