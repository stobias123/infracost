package aws_test

import (
	"testing"

	"github.com/infracost/infracost/api/providers/terraform/tftest"
)

func TestGlobalAccelerator(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	tftest.GoldenFileResourceTests(t, "global_accelerator_test")
}
