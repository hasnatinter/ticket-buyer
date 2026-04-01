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

	testHost := os.Getenv("DB_TEST_HOST")
	if testHost == "" {
		testHost = "test_db"
	}
	os.Setenv("DB_HOST", testHost)

	migrateBin := os.Getenv("MIGRATE_BIN")
	if migrateBin == "" {
		migrateBin = "/usr/src/app/bin/migrate_test"
	}
	err := exec.Command(migrateBin, "up").Run()

	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		err = exec.Command(migrateBin, "reset").Run()
	})

	db := conn.ConnectDb()
	return db
}
