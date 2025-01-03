# VitalQIP Provider for Terraform

The following is an example contents of a provider configuration file named main.tf:

```
provider "vitalqip" {
  server = "127.0.0.1"
  port = "1880"
  context = "workflow"
  username = "qipman"
  password = "qipman"
}
```

Where the fields represent the following:
- **server**: The IP address of the CAA server.
- **port**: The port used to access the CAA server.
- **context**: This is the URL root context of the CAA server. Default value is workflow.
- **username**: Username to authenticate with QIP.
- **password**: Password to authenticate with QIP.

## Resources

Below are the available resources for the following objectTypes:

-   QIP IPv4 Subnet (vitalqip_ipv4_subnet)
-   QIP IPv6 Subnet (vitalqip_ipv6_subnet)

## Data Sources

Below are the available VitalQIP data sources:

-   QIP IPv4 Subnet (vitalqip_ipv4_subnet)
-   QIP IPv6 Subnet (vitalqip_ipv6_subnet)