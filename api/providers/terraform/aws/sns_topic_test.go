package aws_test

import (
	"testing"

	"github.com/infracost/infracost/api/providers/terraform/tftest"
)

func TestSNSTopicGoldenFile(t *testing.T) {
	t.Parallel()
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	tftest.GoldenFileResourceTests(t, "sns_topic_test")
}
