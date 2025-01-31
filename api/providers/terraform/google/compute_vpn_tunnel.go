package google

import (
	"github.com/infracost/infracost/api/resources/google"
	"github.com/infracost/infracost/api/schema"
)

func getComputeVPNTunnelRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "google_compute_vpn_tunnel",
		RFunc: NewComputeVPNTunnel,
	}
}

func NewComputeVPNTunnel(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	r := &google.ComputeVPNTunnel{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}

	r.PopulateUsage(u)
	return r.BuildResource()
}
