package azure_test

import (
	"testing"

	"github.com/infracost/infracost/api/providers/terraform/tftest"
)

func TestAzureRMDNStxtRecord(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	tftest.GoldenFileResourceTests(t, "dns_txt_record_test")
}
