package vitalqip

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"vitalqip": testAccProvider,
	}
}

func testProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}

}

var server = fmt.Sprintf(
	`provider "vitalqip" {
	server = "127.0.0.1"
	port = "1880"
	password = "qipman"
	username = "qipman"
	}`)
