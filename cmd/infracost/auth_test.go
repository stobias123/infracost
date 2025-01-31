package main_test

import (
	"github.com/infracost/infracost/api/testutil"
	"testing"
)

func TestAuthNoArgs(t *testing.T) {
	GoldenFileCommandTest(t, testutil.CalcGoldenFileTestdataDirName(), []string{"auth"}, nil)
}

func TestAuthHelpFlag(t *testing.T) {
	GoldenFileCommandTest(t, testutil.CalcGoldenFileTestdataDirName(), []string{"auth", "--help"}, nil)
}
