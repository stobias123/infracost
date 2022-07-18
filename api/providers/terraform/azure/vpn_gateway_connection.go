package azure

import (
	"github.com/infracost/infracost/api/resources/azure"
	"github.com/infracost/infracost/api/schema"
)

func getAzureRMVPNGatewayConnectionRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "azurerm_vpn_gateway_connection",
		RFunc: newVPNGatewayConnection,
	}
}

func newVPNGatewayConnection(d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	v := &azure.VPNGatewayConnection{
		Address: d.Address,
		Region:  d.Get("region").String(),
	}

	return v.BuildResource()
}
