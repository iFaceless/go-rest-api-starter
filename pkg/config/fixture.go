package config

import (
	"fmt"
	"path"
	"runtime"

	"github.com/iFaceless/fixture"
)

func NewTestFixture() *fixture.TestFixture {
	return fixture.New(
		fixture.Database(fmt.Sprintf("mysql://%s", MySQLConfig.Master)),
		fixture.SchemaFilepath(path.Join(projectDir(), "testdata", "schema.sql")),
		fixture.DataDir(path.Join(projectDir(), "testdata", "fixtures")),
	)
}

func projectDir() string {
	_, f, _, _ := runtime.Caller(1)
	curDir := path.Dir(f)
	return path.Dir(path.Dir(curDir))
}
