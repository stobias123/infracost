package azure_test

import (
	"testing"

	"github.com/infracost/infracost/api/providers/terraform/tftest"
)

func TestAzureRMLinuxVirtualMachineScaleSetGoldenFile(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	tftest.GoldenFileResourceTests(t, "linux_virtual_machine_scale_set_test")
}
