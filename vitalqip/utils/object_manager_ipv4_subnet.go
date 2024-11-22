package utils

import (
	"fmt"
	"log"

	en "terraform-provider-vitalqip/vitalqip/entities"
)

/* CreateSubnet */
func (objMgr *ObjectManager) CreateIPv4Subnet(subnet *en.IPv4Subnet) (*en.IPv4Subnet, error) {
	_, err := objMgr.connector.CreateObject(subnet, "qipaddsubnet")
	if err != nil {
		return nil, err
	}

	return subnet, err
}

/* get Subnet by Id ref */
func (objMgr *ObjectManager) GetIPv4Subnet(query map[string]string) (*en.IPv4Subnet, error) {
	subnet := &en.IPv4Subnet{}
	queryParams := en.NewQueryParams(query)
	err := objMgr.connector.GetObject(en.NewIPv4Subnet(en.IPv4Subnet{}), "qipgetsubnet", &subnet, queryParams)
	log.Printf("[DEBUG] Get QIP Ipv4 Subnet: %s \n", subnet)
	return subnet, err
}

/* delete Subnet by Id ref */
func (objMgr *ObjectManager) DeleteIPv4Subnet(query map[string]string) error {
	queryParams := en.NewQueryParams(query)
	log.Println("[DEBUG] Subnet post: " + fmt.Sprintf("%v", queryParams))
	_, err := objMgr.connector.DeleteObject(en.NewIPv4Subnet(en.IPv4Subnet{}), "qipdeletesubnet", queryParams)
	log.Printf("delete subnet %s", query["subnetAddress"])
	return err
}

/* UpdateSubnet */
func (objMgr *ObjectManager) UpdateIPv4Subnet(subnet *en.IPv4Subnet) (*en.IPv4Subnet, error) {

	_, err := objMgr.connector.UpdateObject(subnet, "qipmodifysubnet")
	if err != nil {
		return nil, err
	}

	return subnet, nil
}
