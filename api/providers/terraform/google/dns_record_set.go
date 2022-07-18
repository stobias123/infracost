package google

import (
	"github.com/infracost/infracost/api/resources/google"
	"github.com/infracost/infracost/api/schema"
)

func getDNSRecordSetRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "google_dns_record_set",
		RFunc: NewDNSRecordSet,
	}
}

func NewDNSRecordSet(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &google.DNSRecordSet{
		Address: d.Address,
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
