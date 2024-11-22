package utils

import (
	"fmt"
	"log"

	en "terraform-provider-vitalqip/vitalqip/entities"
)

/* CreateSubnet */
func (objMgr *ObjectManager) CreateIpv6Subnet(subnet *en.IPv6Subnet) error {
	_, err := objMgr.connector.CreateObject(subnet, "qipaddsubnet")
	if err != nil {
		return err
	}

	return nil
}

/* get Subnet by Id ref */
func (objMgr *ObjectManager) GetIPv6Subnet(query map[string]string) (*en.IPv6SubnetGet, error) {
	subnet := &en.IPv6SubnetGet{}
	queryParams := en.NewQueryParams(query)
	err := objMgr.connector.GetObject(en.NewIPv6Subnet(en.IPv6Subnet{}), "qipgetsubnet", &subnet, queryParams)
	log.Printf("[DEBUG] Get QIP Ipv6 Subnet: %s \n", subnet)
	return subnet, err
}

/* delete Subnet by Id ref */
func (objMgr *ObjectManager) DeleteIPv6Subnet(query map[string]string) error {
	queryParams := en.NewQueryParams(query)
	log.Println("[DEBUG] Subnet delete: " + fmt.Sprintf("%v", queryParams))
	_, err := objMgr.connector.DeleteObject(en.NewIPv6Subnet(en.IPv6Subnet{}), "qipdeletesubnet", queryParams)
	return err
}

/* UpdateSubnet */
func (objMgr *ObjectManager) UpdateIPv6Subnet(subnet *en.IPv6SubnetModify) (*en.IPv6SubnetModify, error) {

	_, err := objMgr.connector.UpdateObject(subnet, "qipmodifysubnet")
	if err != nil {
		return nil, err
	}

	return subnet, nil
}
