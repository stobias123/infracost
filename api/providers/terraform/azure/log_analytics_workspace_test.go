package azure_test

import (
	"testing"

	"github.com/infracost/infracost/api/providers/terraform/tftest"
)

func TestLogAnalyticsWorkspaceGoldenFile(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	tftest.GoldenFileResourceTestsWithOpts(t, "log_analytics_workspace_test", &tftest.GoldenFileOptions{
		CaptureLogs: true,
	})
}
