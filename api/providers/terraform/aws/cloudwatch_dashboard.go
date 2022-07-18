package aws

import (
	"github.com/infracost/infracost/api/resources/aws"
	"github.com/infracost/infracost/api/schema"
)

func getCloudwatchDashboardRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_cloudwatch_dashboard",
		RFunc: NewCloudwatchDashboard,
	}
}
func NewCloudwatchDashboard(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &aws.CloudwatchDashboard{
		Address: d.Address,
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
