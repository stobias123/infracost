package aws

import (
	"github.com/infracost/infracost/api/resources/aws"
	"github.com/infracost/infracost/api/schema"
)

func getNewKMSExternalKeyRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "aws_kms_external_key",
		RFunc: NewKMSExternalKey,
	}
}

func NewKMSExternalKey(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &aws.KMSExternalKey{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
