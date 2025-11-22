// Copyright (c) Robin Mordasiewicz
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/client"
	ds_alert_policy "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/alert_policy"
	ds_api_definition "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/api_definition"
	ds_app_firewall "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/app_firewall"
	ds_aws_vpc_site "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/aws_vpc_site"
	ds_azure_vnet_site "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/azure_vnet_site"
	ds_certificate "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/certificate"
	ds_cloud_credentials "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/cloud_credentials"
	ds_cloud_site "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/cloud_site"
	ds_discovery "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/discovery"
	ds_dns_zone "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/dns_zone"
	ds_gcp_vpc_site "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/gcp_vpc_site"
	ds_healthcheck "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/healthcheck"
	ds_http_loadbalancer "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/http_loadbalancer"
	ds_ip_prefix_set "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/ip_prefix_set"
	ds_namespace "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/namespace"
	ds_network_policy "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/network_policy"
	ds_origin_pool "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/origin_pool"
	ds_rate_limiter "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/rate_limiter"
	ds_role "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/role"
	ds_service_policy "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/service_policy"
	ds_tcp_loadbalancer "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/tcp_loadbalancer"
	ds_tunnel "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/tunnel"
	ds_user "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/user"
	ds_virtual_site "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/virtual_site"
	ds_voltstack_site "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/voltstack_site"
	ds_log_receiver "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/log_receiver"
	ds_route "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/route"
	ds_secret_policy "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/secret_policy"
	ds_service_mesh "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/service_mesh"
	ds_waf_rule "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/waf_rule"
	ds_enhanced_firewall_policy "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/enhanced_firewall_policy"
	ds_forward_proxy_policy "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/forward_proxy_policy"
	ds_known_label "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/known_label"
	ds_site_mesh_group "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/site_mesh_group"
	ds_active_service_policies "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/active_service_policies"
	ds_api_credential "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/api_credential"
	ds_app_type "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/app_type"
	ds_client_connection_limit "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/client_connection_limit"
	ds_dns_lb_pool "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/dns_lb_pool"
	ds_global_log_receiver "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/global_log_receiver"
	ds_malicious_user_mitigation "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/malicious_user_mitigation"
	ds_protocol_policer "github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/datasources/protocol_policer"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/active_service_policies"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/advertise_policy"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/alert_policy"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/alert_receiver"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/bgp"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/bgp_asn_set"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/cdn_cache_rule"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/cdn_loadbalancer"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/api_credential"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/api_definition"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/app_firewall"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/app_type"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/aws_tgw_site"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/aws_vpc_site"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/azure_vnet_site"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/certificate"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/client_connection_limit"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/cloud_credentials"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/cloud_site"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/cluster"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/container_registry"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/discovery"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/dns_domain"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/dns_lb_health_check"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/dns_lb_pool"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/dns_load_balancer"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/dns_zone"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/endpoint"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/enhanced_firewall_policy"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/external_connector"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/fast_acl"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/fleet"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/fast_acl_rule"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/forward_proxy_policy"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/gcp_vpc_site"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/global_log_receiver"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/healthcheck"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/http_loadbalancer"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/ip_prefix_set"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/k8s_cluster"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/k8s_cluster_role"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/k8s_cluster_role_binding"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/known_label"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/log_receiver"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/malicious_user_mitigation"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/namespace"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/nat_policy"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/network_connector"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/network_firewall"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/network_interface"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/network_policy"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/network_policy_rule"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/origin_pool"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/policer"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/protocol_policer"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/rate_limiter"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/role"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/route"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/secret_policy"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/secret_policy_rule"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/securemesh_site"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/service_mesh"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/service_policy"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/service_policy_rule"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/site_mesh_group"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/tcp_loadbalancer"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/tunnel"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/udp_loadbalancer"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/user"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/virtual_host"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/virtual_k8s"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/virtual_network"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/virtual_site"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/voltstack_site"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/subnet"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/user_group"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/waf"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/waf_exclusion_policy"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/waf_rule"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/app_setting"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/policy_based_routing"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/proxy"
	"github.com/robinmordasiewicz/terraform-provider-f5-distributed-cloud/internal/resources/rate_limiter_policy"
)

// Ensure F5XCProvider satisfies various provider interfaces.
var _ provider.Provider = &F5XCProvider{}

// F5XCProvider defines the provider implementation.
type F5XCProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// F5XCProviderModel describes the provider data model.
type F5XCProviderModel struct {
	APIURL      types.String `tfsdk:"api_url"`
	P12File     types.String `tfsdk:"p12_file"`
	P12Password types.String `tfsdk:"p12_password"`
	APIToken    types.String `tfsdk:"api_token"`
}

// New returns a new provider instance.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &F5XCProvider{
			version: version,
		}
	}
}

// Metadata returns the provider type name.
func (p *F5XCProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "f5xc"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *F5XCProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Terraform provider for F5 Distributed Cloud (F5 XC) services.",
		MarkdownDescription: `
The F5 Distributed Cloud (F5 XC) provider enables infrastructure engineers to manage F5 XC resources using Terraform.

## Authentication

The provider supports two authentication methods:

1. **API Token Authentication** (recommended for automation):
   - Set the ` + "`api_token`" + ` attribute or ` + "`F5XC_API_TOKEN`" + ` environment variable

2. **Certificate Authentication** (P12 file):
   - Set ` + "`p12_file`" + ` and ` + "`p12_password`" + ` attributes or
   - Set ` + "`F5XC_API_P12_FILE`" + ` and ` + "`F5XC_API_P12_PASSWORD`" + ` environment variables

## Example Usage

` + "```hcl" + `
provider "f5xc" {
  api_url   = "https://your-tenant.console.ves.volterra.io/api"
  api_token = var.f5xc_api_token
}
` + "```" + `
`,
		Attributes: map[string]schema.Attribute{
			"api_url": schema.StringAttribute{
				Description: "The URL for the F5 XC API. Can also be set via the F5XC_API_URL environment variable.",
				Optional:    true,
			},
			"p12_file": schema.StringAttribute{
				Description: "Path to the P12 certificate file for authentication. Can also be set via the F5XC_API_P12_FILE environment variable.",
				Optional:    true,
			},
			"p12_password": schema.StringAttribute{
				Description: "Password for the P12 certificate file. Can also be set via the F5XC_API_P12_PASSWORD environment variable.",
				Optional:    true,
				Sensitive:   true,
			},
			"api_token": schema.StringAttribute{
				Description: "API token for authentication. Can also be set via the F5XC_API_TOKEN environment variable.",
				Optional:    true,
				Sensitive:   true,
			},
		},
	}
}

// Configure prepares an F5 XC API client for data sources and resources.
func (p *F5XCProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring F5 XC provider")

	var config F5XCProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get values from environment variables if not set in config
	apiURL := os.Getenv("F5XC_API_URL")
	if !config.APIURL.IsNull() {
		apiURL = config.APIURL.ValueString()
	}

	p12File := os.Getenv("F5XC_API_P12_FILE")
	if !config.P12File.IsNull() {
		p12File = config.P12File.ValueString()
	}

	p12Password := os.Getenv("F5XC_API_P12_PASSWORD")
	if !config.P12Password.IsNull() {
		p12Password = config.P12Password.ValueString()
	}

	apiToken := os.Getenv("F5XC_API_TOKEN")
	if !config.APIToken.IsNull() {
		apiToken = config.APIToken.ValueString()
	}

	// Validate required fields
	if apiURL == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_url"),
			"Missing API URL",
			"The provider requires an API URL to be configured. "+
				"Set the api_url attribute or F5XC_API_URL environment variable.",
		)
	}

	// Validate authentication
	hasToken := apiToken != ""
	hasCert := p12File != ""

	if !hasToken && !hasCert {
		resp.Diagnostics.AddError(
			"Missing Authentication Credentials",
			"The provider requires authentication credentials. "+
				"Configure either api_token (or F5XC_API_TOKEN) for token authentication, "+
				"or p12_file/p12_password (or F5XC_API_P12_FILE/F5XC_API_P12_PASSWORD) for certificate authentication.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Create client configuration
	clientConfig := &client.ClientConfig{
		APIURL: apiURL,
		Auth: &client.AuthConfig{
			P12File:     p12File,
			P12Password: p12Password,
			APIToken:    apiToken,
		},
	}

	// Create the F5 XC client
	f5xcClient, err := client.NewClient(clientConfig)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create F5 XC Client",
			"An unexpected error occurred when creating the F5 XC API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"Error: "+err.Error(),
		)
		return
	}

	tflog.Info(ctx, "F5 XC provider configured successfully")

	// Make the client available during DataSource and Resource type Configure methods.
	resp.DataSourceData = f5xcClient
	resp.ResourceData = f5xcClient
}

// Resources defines the resources implemented in the provider.
func (p *F5XCProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		active_service_policies.NewActiveServicePoliciesResource,
		advertise_policy.NewAdvertisePolicyResource,
		alert_policy.NewAlertPolicyResource,
		alert_receiver.NewAlertReceiverResource,
		bgp.NewBGPResource,
		bgp_asn_set.NewBGPASNSetResource,
		cdn_cache_rule.NewCDNCacheRuleResource,
		cdn_loadbalancer.NewCDNLoadBalancerResource,
		api_credential.NewAPICredentialResource,
		api_definition.NewAPIDefinitionResource,
		app_firewall.NewAppFirewallResource,
		app_type.NewAppTypeResource,
		aws_tgw_site.NewAWSTGWSiteResource,
		aws_vpc_site.NewAWSVPCSiteResource,
		azure_vnet_site.NewAzureVNETSiteResource,
		certificate.NewCertificateResource,
		client_connection_limit.NewClientConnectionLimitResource,
		cloud_credentials.NewCloudCredentialsResource,
		cloud_site.NewCloudSiteResource,
		cluster.NewClusterResource,
		container_registry.NewContainerRegistryResource,
		discovery.NewDiscoveryResource,
		dns_domain.NewDNSDomainResource,
		dns_lb_health_check.NewDNSLBHealthCheckResource,
		dns_lb_pool.NewDNSLBPoolResource,
		dns_load_balancer.NewDNSLoadBalancerResource,
		dns_zone.NewDNSZoneResource,
		endpoint.NewEndpointResource,
		enhanced_firewall_policy.NewEnhancedFirewallPolicyResource,
		external_connector.NewExternalConnectorResource,
		fast_acl.NewFastACLResource,
		fleet.NewFleetResource,
		fast_acl_rule.NewFastACLRuleResource,
		forward_proxy_policy.NewForwardProxyPolicyResource,
		gcp_vpc_site.NewGCPVPCSiteResource,
		global_log_receiver.NewGlobalLogReceiverResource,
		healthcheck.NewHealthcheckResource,
		http_loadbalancer.NewHTTPLoadBalancerResource,
		ip_prefix_set.NewIPPrefixSetResource,
		k8s_cluster.NewK8sClusterResource,
		k8s_cluster_role.NewK8sClusterRoleResource,
		k8s_cluster_role_binding.NewK8sClusterRoleBindingResource,
		known_label.NewKnownLabelResource,
		log_receiver.NewLogReceiverResource,
		malicious_user_mitigation.NewMaliciousUserMitigationResource,
		namespace.NewNamespaceResource,
		nat_policy.NewNATPolicyResource,
		network_connector.NewNetworkConnectorResource,
		network_firewall.NewNetworkFirewallResource,
		network_interface.NewNetworkInterfaceResource,
		network_policy.NewNetworkPolicyResource,
		network_policy_rule.NewNetworkPolicyRuleResource,
		origin_pool.NewOriginPoolResource,
		policer.NewPolicerResource,
		protocol_policer.NewProtocolPolicerResource,
		rate_limiter.NewRateLimiterResource,
		role.NewRoleResource,
		route.NewRouteResource,
		secret_policy.NewSecretPolicyResource,
		secret_policy_rule.NewSecretPolicyRuleResource,
		securemesh_site.NewSecuremeshSiteResource,
		service_mesh.NewServiceMeshResource,
		service_policy.NewServicePolicyResource,
		service_policy_rule.NewServicePolicyRuleResource,
		site_mesh_group.NewSiteMeshGroupResource,
		tcp_loadbalancer.NewTCPLoadBalancerResource,
		tunnel.NewTunnelResource,
		udp_loadbalancer.NewUDPLoadBalancerResource,
		user.NewUserResource,
		virtual_host.NewVirtualHostResource,
		virtual_k8s.NewVirtualK8sResource,
		virtual_network.NewVirtualNetworkResource,
		virtual_site.NewVirtualSiteResource,
		voltstack_site.NewVoltstackSiteResource,
		subnet.NewSubnetResource,
		user_group.NewUserGroupResource,
		waf.NewWAFResource,
		waf_exclusion_policy.NewWAFExclusionPolicyResource,
		waf_rule.NewWAFRuleResource,
		app_setting.NewAppSettingResource,
		policy_based_routing.NewPolicyBasedRoutingResource,
		proxy.NewProxyResource,
		rate_limiter_policy.NewRateLimiterPolicyResource,
	}
}

// DataSources defines the data sources implemented in the provider.
func (p *F5XCProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		ds_alert_policy.NewAlertPolicyDataSource,
		ds_api_definition.NewAPIDefinitionDataSource,
		ds_app_firewall.NewAppFirewallDataSource,
		ds_aws_vpc_site.NewAWSVPCSiteDataSource,
		ds_azure_vnet_site.NewAzureVNETSiteDataSource,
		ds_certificate.NewCertificateDataSource,
		ds_cloud_credentials.NewCloudCredentialsDataSource,
		ds_cloud_site.NewCloudSiteDataSource,
		ds_discovery.NewDiscoveryDataSource,
		ds_dns_zone.NewDNSZoneDataSource,
		ds_gcp_vpc_site.NewGCPVPCSiteDataSource,
		ds_healthcheck.NewHealthcheckDataSource,
		ds_http_loadbalancer.NewHTTPLoadBalancerDataSource,
		ds_ip_prefix_set.NewIPPrefixSetDataSource,
		ds_namespace.NewNamespaceDataSource,
		ds_network_policy.NewNetworkPolicyDataSource,
		ds_origin_pool.NewOriginPoolDataSource,
		ds_rate_limiter.NewRateLimiterDataSource,
		ds_role.NewRoleDataSource,
		ds_service_policy.NewServicePolicyDataSource,
		ds_tcp_loadbalancer.NewTCPLoadBalancerDataSource,
		ds_tunnel.NewTunnelDataSource,
		ds_user.NewUserDataSource,
		ds_virtual_site.NewVirtualSiteDataSource,
		ds_voltstack_site.NewVoltstackSiteDataSource,
		ds_waf_rule.NewWAFRuleDataSource,
		ds_route.NewRouteDataSource,
		ds_log_receiver.NewLogReceiverDataSource,
		ds_secret_policy.NewSecretPolicyDataSource,
		ds_service_mesh.NewServiceMeshDataSource,
		ds_site_mesh_group.NewSiteMeshGroupDataSource,
		ds_forward_proxy_policy.NewForwardProxyPolicyDataSource,
		ds_enhanced_firewall_policy.NewEnhancedFirewallPolicyDataSource,
		ds_known_label.NewKnownLabelDataSource,
		ds_global_log_receiver.NewGlobalLogReceiverDataSource,
		ds_malicious_user_mitigation.NewMaliciousUserMitigationDataSource,
		ds_protocol_policer.NewProtocolPolicerDataSource,
		ds_app_type.NewAppTypeDataSource,
		ds_dns_lb_pool.NewDNSLBPoolDataSource,
		ds_api_credential.NewAPICredentialDataSource,
		ds_client_connection_limit.NewClientConnectionLimitDataSource,
		ds_active_service_policies.NewActiveServicePoliciesDataSource,
	}
}
