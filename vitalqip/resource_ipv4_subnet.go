package vitalqip

import (
	"context"
	"fmt"
	"log"

	// "strconv"

	"strings"
	en "terraform-provider-vitalqip/vitalqip/entities"
	cc "terraform-provider-vitalqip/vitalqip/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIPv4Subnet() *schema.Resource {
	return &schema.Resource{
		CreateContext: createIPv4SubnetRecord,
		ReadContext:   getIPv4SubnetRecord,
		UpdateContext: updateIPv4SubnetRecord,
		DeleteContext: deleteIPv4SubnetRecord,

		Schema: map[string]*schema.Schema{
			"org_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Organization Name.",
			},
			"subnet_address": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "IPv4 address of subnet.",
			},
			"subnet_mask": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "Subnet mask of subnet.",
			},
			"network_address": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "IPv4 address of network.",
			},
			"warning_type": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Type of warning when the defined percentage of addresses reached.",
			},
			"warning_percent": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Percentage of managed addresses. Warning if this percentage is reached. The value range is 0 - 100.",
			},
			"subnet_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Name of subnet.",
			},
		},
	}
}

func createIPv4SubnetRecord(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)

	var err error
	var diags diag.Diagnostics
	subnet := getIPv4SubnetFromResourceData(d)

	log.Println("[DEBUG] Subnet post: " + fmt.Sprintf("%v", subnet))

	_, err = objMgr.CreateIPv4Subnet(subnet)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Creation of QIP IPv4 Subnet failed",
			Detail:   fmt.Sprintf("Creation of Subnet (%s) failed: %s", subnet.SubnetAddress, err),
		})
		return diags
	}

	d.SetId(subnet.SubnetAddress)

	return getIPv4SubnetRecord(ctx, d, m)
}

func getIPv4SubnetRecord(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)
	var diags diag.Diagnostics

	orgName := strings.TrimSpace(d.Get("org_name").(string))
	subnetAddress := strings.TrimSpace(d.Get("subnet_address").(string))

	query := map[string]string{
		"orgName":        orgName,
		"subnetAddress":  subnetAddress,
		"addressVersion": "4",
	}

	response, err := objMgr.GetIPv4Subnet(query)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Getting QIP IPv4 Subnet Failed",
			Detail:   fmt.Sprintf("Getting QIP IPv4 Subnet (%s) failed : %s", subnetAddress, err),
		})
		return diags
	}

	flattenIPv4Subnet(d, response)

	return nil
}

func updateIPv4SubnetRecord(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)
	var err error
	var diags diag.Diagnostics
	subnet := getIPv4SubnetFromResourceData(d)

	_, err = objMgr.UpdateIPv4Subnet(subnet)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Updating of QIP IPv4 Subnet failed",
			Detail:   fmt.Sprintf("Updating QIP IPv4 Subnet by Id (%s) failed : %s", d.Id(), err),
		})
		return diags
	}

	return getIPv4SubnetRecord(ctx, d, m)

}

func deleteIPv4SubnetRecord(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	connector := m.(*cc.Connector)
	objMgr := cc.NewObjectManager(connector)
	var diags diag.Diagnostics

	orgName := strings.TrimSpace(d.Get("org_name").(string))
	subnetAddress := strings.TrimSpace(d.Get("subnet_address").(string))

	query := map[string]string{
		"orgName":        orgName,
		"subnetAddress":  subnetAddress,
		"addressVersion": "4",
	}

	log.Println("[DEBUG] Subnet post: " + fmt.Sprintf("%v", query))

	err := objMgr.DeleteIPv4Subnet(query)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Deletion of QIP IPv4 Subnet failed",
			Detail:   fmt.Sprintf("Deleting QIP IPv4 Subnet block by Id (%s) failed : %s", d.Id(), err),
		})
		return diags
	}

	d.SetId(subnetAddress)
	//log.Printf("[DEBUG] %s: Deletion of network block complete", rsSubnetIdString(d))

	return diags
}

func getIPv4SubnetFromResourceData(d *schema.ResourceData) *en.IPv4Subnet {

	orgName := strings.TrimSpace(d.Get("org_name").(string))
	subnetAddress := strings.TrimSpace(d.Get("subnet_address").(string))
	subnetMask := strings.TrimSpace(d.Get("subnet_mask").(string))
	networkAddress := strings.TrimSpace(d.Get("network_address").(string))
	subnetName := strings.TrimSpace(d.Get("subnet_name").(string))
	warningType := d.Get("warning_type").(int)
	warningPercent := d.Get("warning_percent").(int)

	return en.NewIPv4Subnet(en.IPv4Subnet{
		OrgName:           orgName,
		SubnetAddress:     subnetAddress,
		SubnetMask:        subnetMask,
		NetworkAddress:    networkAddress,
		SubnetName:        subnetName,
		WarningType:       warningType,
		WarningPercentage: warningPercent,
		AddressVersion:    4,
	})
}
