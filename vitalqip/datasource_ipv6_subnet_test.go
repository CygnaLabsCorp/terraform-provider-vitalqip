package vitalqip

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceIPv6Subnet(t *testing.T) {
	dataName := "data.vitalqip_ipv6_subnet.ipv6_subnet_data"
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccConfigWithProvider(
					`
					resource "vitalqip_ipv6_subnet" "ipv6_subnet_resource" {
					org_name= "Demo"
					subnet_address="2000:0:0:10::"
					subnet_name="subnet_name"
					subnet_prefix_length = 60
					block_prefix_length = 48
					block_address="2000::"
					create_reverse_zone=true
					}
					
					data "vitalqip_ipv6_subnet" "ipv6_subnet_data" {
					org_name= "Demo"
					subnet_address="2000:0:0:10::"
					subnet_prefix_length = 60
					depends_on = [vitalqip_ipv6_subnet.ipv6_subnet_resource]
					}
					`,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataName, "org_name", "Demo"),
					resource.TestCheckResourceAttr(dataName, "subnet_address", "2000:0:0:10::"),
					resource.TestCheckResourceAttr(dataName, "subnet_name", "subnet_name"),
					resource.TestCheckResourceAttr(dataName, "subnet_prefix_length", "60"),
				),
			},
		},
	})
}
