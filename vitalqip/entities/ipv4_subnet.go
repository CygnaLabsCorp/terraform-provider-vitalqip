package entities

import (
	"fmt"
	"strings"
)

type IPv4Subnet struct {
	ObjBase           `json:"-"`
	OrgName           string `json:"orgName,omitempty"`
	SubnetAddress     string `json:"subnetAddress"`
	SubnetMask        string `json:"subnetMask"`
	NetworkAddress    string `json:"networkAddress"`
	SubnetName        string `json:"subnetName"`
	WarningType       int    `json:"warningType"`
	WarningPercentage int    `json:"warningPercentage"`
	AddressVersion    int    `json:"addressVersion"`
}

func (subnet IPv4Subnet) String() string {
	var sb strings.Builder

	sb.WriteString("IPv4Subnet {\n")
	sb.WriteString(fmt.Sprintf("  OrgName: %s\n", subnet.OrgName))
	sb.WriteString(fmt.Sprintf("  SubnetAddress: %s\n", subnet.SubnetAddress))
	sb.WriteString(fmt.Sprintf("  SubnetMask: %s\n", subnet.SubnetMask))
	sb.WriteString(fmt.Sprintf("  NetworkAddress: %s\n", subnet.NetworkAddress))
	sb.WriteString(fmt.Sprintf("  SubnetName: %s\n", subnet.SubnetName))
	sb.WriteString(fmt.Sprintf("  WarningType: %d\n", subnet.WarningType))
	sb.WriteString(fmt.Sprintf("  WarningPercentage: %d\n", subnet.WarningPercentage))
	sb.WriteString(fmt.Sprintf("  AddressVersion: %d\n", subnet.AddressVersion))
	sb.WriteString("}\n")

	return sb.String()
}

func NewIPv4Subnet(sb IPv4Subnet) *IPv4Subnet {
	res := sb
	res.objectType = "ipv4_subnet"
	return &res
}
