package aws

import (
	"strings"

	"github.com/infracost/infracost/api/resources/aws"
	"github.com/infracost/infracost/api/schema"
)

func getDataTransferRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_data_transfer",
		RFunc: newDataTransfer,
	}
}

func newDataTransfer(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	region := strings.ToLower(u.Get("region").String())

	r := &aws.DataTransfer{
		Address: d.Address,
		Region:  region,
	}
	r.PopulateUsage(u)

	return r.BuildResource()
}
