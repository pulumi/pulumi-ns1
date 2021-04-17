// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ns1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type APIKey struct {
	pulumi.CustomResourceState

	// Whether the apikey can modify account settings.
	AccountManageAccountSettings pulumi.BoolPtrOutput `pulumi:"accountManageAccountSettings"`
	// Whether the apikey can modify account apikeys.
	AccountManageApikeys pulumi.BoolPtrOutput `pulumi:"accountManageApikeys"`
	// Whether the apikey can manage ip whitelist.
	AccountManageIpWhitelist pulumi.BoolPtrOutput `pulumi:"accountManageIpWhitelist"`
	// Whether the apikey can modify account payment methods.
	AccountManagePaymentMethods pulumi.BoolPtrOutput `pulumi:"accountManagePaymentMethods"`
	// Whether the apikey can modify the account plan.
	//
	// Deprecated: obsolete, should no longer be used
	AccountManagePlan pulumi.BoolPtrOutput `pulumi:"accountManagePlan"`
	// Whether the apikey can modify other teams in the account.
	AccountManageTeams pulumi.BoolPtrOutput `pulumi:"accountManageTeams"`
	// Whether the apikey can modify account users.
	AccountManageUsers pulumi.BoolPtrOutput `pulumi:"accountManageUsers"`
	// Whether the apikey can view activity logs.
	AccountViewActivityLog pulumi.BoolPtrOutput `pulumi:"accountViewActivityLog"`
	// Whether the apikey can view invoices.
	AccountViewInvoices pulumi.BoolPtrOutput `pulumi:"accountViewInvoices"`
	// Whether the apikey can modify data feeds.
	DataManageDatafeeds pulumi.BoolPtrOutput `pulumi:"dataManageDatafeeds"`
	// Whether the apikey can modify data sources.
	DataManageDatasources pulumi.BoolPtrOutput `pulumi:"dataManageDatasources"`
	// Whether the apikey can publish to data feeds.
	DataPushToDatafeeds pulumi.BoolPtrOutput `pulumi:"dataPushToDatafeeds"`
	// Whether the apikey can manage DHCP.
	// Only relevant for the DDI product.
	DhcpManageDhcp pulumi.BoolPtrOutput `pulumi:"dhcpManageDhcp"`
	// Whether the apikey can view DHCP.
	// Only relevant for the DDI product.
	DhcpViewDhcp pulumi.BoolPtrOutput `pulumi:"dhcpViewDhcp"`
	// Whether the apikey can modify the accounts zones.
	DnsManageZones pulumi.BoolPtrOutput `pulumi:"dnsManageZones"`
	// Whether the apikey can view the accounts zones.
	DnsViewZones pulumi.BoolPtrOutput `pulumi:"dnsViewZones"`
	// If true, enable the `dnsZonesAllow` list, otherwise enable the `dnsZonesDeny` list.
	DnsZonesAllowByDefault pulumi.BoolPtrOutput `pulumi:"dnsZonesAllowByDefault"`
	// List of zones that the apikey may access.
	DnsZonesAllows pulumi.StringArrayOutput `pulumi:"dnsZonesAllows"`
	// List of zones that the apikey may not access.
	DnsZonesDenies pulumi.StringArrayOutput `pulumi:"dnsZonesDenies"`
	// Sets exclusivity on this IP whitelist.
	IpWhitelistStrict pulumi.BoolPtrOutput `pulumi:"ipWhitelistStrict"`
	// The IP addresses to whitelist for this key.
	IpWhitelists pulumi.StringArrayOutput `pulumi:"ipWhitelists"`
	// Whether the apikey can manage IPAM.
	// Only relevant for the DDI product.
	IpamManageIpam pulumi.BoolPtrOutput `pulumi:"ipamManageIpam"`
	// Whether the apikey can view IPAM.
	// Only relevant for the DDI product.
	IpamViewIpam pulumi.BoolPtrOutput `pulumi:"ipamViewIpam"`
	// (Computed) The apikeys authentication token.
	Key pulumi.StringOutput `pulumi:"key"`
	// Whether the apikey can modify monitoring jobs.
	MonitoringManageJobs pulumi.BoolPtrOutput `pulumi:"monitoringManageJobs"`
	// Whether the apikey can modify notification lists.
	MonitoringManageLists pulumi.BoolPtrOutput `pulumi:"monitoringManageLists"`
	// Whether the apikey can view monitoring jobs.
	MonitoringViewJobs pulumi.BoolPtrOutput `pulumi:"monitoringViewJobs"`
	// The free form name of the apikey.
	Name pulumi.StringOutput `pulumi:"name"`
	// Whether the apikey can manage global active directory.
	// Only relevant for the DDI product.
	SecurityManageActiveDirectory pulumi.BoolPtrOutput `pulumi:"securityManageActiveDirectory"`
	// Whether the apikey can manage global two factor authentication.
	SecurityManageGlobal2fa pulumi.BoolPtrOutput `pulumi:"securityManageGlobal2fa"`
	// The teams that the apikey belongs to.
	Teams pulumi.StringArrayOutput `pulumi:"teams"`
}

// NewAPIKey registers a new resource with the given unique name, arguments, and options.
func NewAPIKey(ctx *pulumi.Context,
	name string, args *APIKeyArgs, opts ...pulumi.ResourceOption) (*APIKey, error) {
	if args == nil {
		args = &APIKeyArgs{}
	}

	var resource APIKey
	err := ctx.RegisterResource("ns1:index/aPIKey:APIKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAPIKey gets an existing APIKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAPIKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *APIKeyState, opts ...pulumi.ResourceOption) (*APIKey, error) {
	var resource APIKey
	err := ctx.ReadResource("ns1:index/aPIKey:APIKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering APIKey resources.
type apikeyState struct {
	// Whether the apikey can modify account settings.
	AccountManageAccountSettings *bool `pulumi:"accountManageAccountSettings"`
	// Whether the apikey can modify account apikeys.
	AccountManageApikeys *bool `pulumi:"accountManageApikeys"`
	// Whether the apikey can manage ip whitelist.
	AccountManageIpWhitelist *bool `pulumi:"accountManageIpWhitelist"`
	// Whether the apikey can modify account payment methods.
	AccountManagePaymentMethods *bool `pulumi:"accountManagePaymentMethods"`
	// Whether the apikey can modify the account plan.
	//
	// Deprecated: obsolete, should no longer be used
	AccountManagePlan *bool `pulumi:"accountManagePlan"`
	// Whether the apikey can modify other teams in the account.
	AccountManageTeams *bool `pulumi:"accountManageTeams"`
	// Whether the apikey can modify account users.
	AccountManageUsers *bool `pulumi:"accountManageUsers"`
	// Whether the apikey can view activity logs.
	AccountViewActivityLog *bool `pulumi:"accountViewActivityLog"`
	// Whether the apikey can view invoices.
	AccountViewInvoices *bool `pulumi:"accountViewInvoices"`
	// Whether the apikey can modify data feeds.
	DataManageDatafeeds *bool `pulumi:"dataManageDatafeeds"`
	// Whether the apikey can modify data sources.
	DataManageDatasources *bool `pulumi:"dataManageDatasources"`
	// Whether the apikey can publish to data feeds.
	DataPushToDatafeeds *bool `pulumi:"dataPushToDatafeeds"`
	// Whether the apikey can manage DHCP.
	// Only relevant for the DDI product.
	DhcpManageDhcp *bool `pulumi:"dhcpManageDhcp"`
	// Whether the apikey can view DHCP.
	// Only relevant for the DDI product.
	DhcpViewDhcp *bool `pulumi:"dhcpViewDhcp"`
	// Whether the apikey can modify the accounts zones.
	DnsManageZones *bool `pulumi:"dnsManageZones"`
	// Whether the apikey can view the accounts zones.
	DnsViewZones *bool `pulumi:"dnsViewZones"`
	// If true, enable the `dnsZonesAllow` list, otherwise enable the `dnsZonesDeny` list.
	DnsZonesAllowByDefault *bool `pulumi:"dnsZonesAllowByDefault"`
	// List of zones that the apikey may access.
	DnsZonesAllows []string `pulumi:"dnsZonesAllows"`
	// List of zones that the apikey may not access.
	DnsZonesDenies []string `pulumi:"dnsZonesDenies"`
	// Sets exclusivity on this IP whitelist.
	IpWhitelistStrict *bool `pulumi:"ipWhitelistStrict"`
	// The IP addresses to whitelist for this key.
	IpWhitelists []string `pulumi:"ipWhitelists"`
	// Whether the apikey can manage IPAM.
	// Only relevant for the DDI product.
	IpamManageIpam *bool `pulumi:"ipamManageIpam"`
	// Whether the apikey can view IPAM.
	// Only relevant for the DDI product.
	IpamViewIpam *bool `pulumi:"ipamViewIpam"`
	// (Computed) The apikeys authentication token.
	Key *string `pulumi:"key"`
	// Whether the apikey can modify monitoring jobs.
	MonitoringManageJobs *bool `pulumi:"monitoringManageJobs"`
	// Whether the apikey can modify notification lists.
	MonitoringManageLists *bool `pulumi:"monitoringManageLists"`
	// Whether the apikey can view monitoring jobs.
	MonitoringViewJobs *bool `pulumi:"monitoringViewJobs"`
	// The free form name of the apikey.
	Name *string `pulumi:"name"`
	// Whether the apikey can manage global active directory.
	// Only relevant for the DDI product.
	SecurityManageActiveDirectory *bool `pulumi:"securityManageActiveDirectory"`
	// Whether the apikey can manage global two factor authentication.
	SecurityManageGlobal2fa *bool `pulumi:"securityManageGlobal2fa"`
	// The teams that the apikey belongs to.
	Teams []string `pulumi:"teams"`
}

type APIKeyState struct {
	// Whether the apikey can modify account settings.
	AccountManageAccountSettings pulumi.BoolPtrInput
	// Whether the apikey can modify account apikeys.
	AccountManageApikeys pulumi.BoolPtrInput
	// Whether the apikey can manage ip whitelist.
	AccountManageIpWhitelist pulumi.BoolPtrInput
	// Whether the apikey can modify account payment methods.
	AccountManagePaymentMethods pulumi.BoolPtrInput
	// Whether the apikey can modify the account plan.
	//
	// Deprecated: obsolete, should no longer be used
	AccountManagePlan pulumi.BoolPtrInput
	// Whether the apikey can modify other teams in the account.
	AccountManageTeams pulumi.BoolPtrInput
	// Whether the apikey can modify account users.
	AccountManageUsers pulumi.BoolPtrInput
	// Whether the apikey can view activity logs.
	AccountViewActivityLog pulumi.BoolPtrInput
	// Whether the apikey can view invoices.
	AccountViewInvoices pulumi.BoolPtrInput
	// Whether the apikey can modify data feeds.
	DataManageDatafeeds pulumi.BoolPtrInput
	// Whether the apikey can modify data sources.
	DataManageDatasources pulumi.BoolPtrInput
	// Whether the apikey can publish to data feeds.
	DataPushToDatafeeds pulumi.BoolPtrInput
	// Whether the apikey can manage DHCP.
	// Only relevant for the DDI product.
	DhcpManageDhcp pulumi.BoolPtrInput
	// Whether the apikey can view DHCP.
	// Only relevant for the DDI product.
	DhcpViewDhcp pulumi.BoolPtrInput
	// Whether the apikey can modify the accounts zones.
	DnsManageZones pulumi.BoolPtrInput
	// Whether the apikey can view the accounts zones.
	DnsViewZones pulumi.BoolPtrInput
	// If true, enable the `dnsZonesAllow` list, otherwise enable the `dnsZonesDeny` list.
	DnsZonesAllowByDefault pulumi.BoolPtrInput
	// List of zones that the apikey may access.
	DnsZonesAllows pulumi.StringArrayInput
	// List of zones that the apikey may not access.
	DnsZonesDenies pulumi.StringArrayInput
	// Sets exclusivity on this IP whitelist.
	IpWhitelistStrict pulumi.BoolPtrInput
	// The IP addresses to whitelist for this key.
	IpWhitelists pulumi.StringArrayInput
	// Whether the apikey can manage IPAM.
	// Only relevant for the DDI product.
	IpamManageIpam pulumi.BoolPtrInput
	// Whether the apikey can view IPAM.
	// Only relevant for the DDI product.
	IpamViewIpam pulumi.BoolPtrInput
	// (Computed) The apikeys authentication token.
	Key pulumi.StringPtrInput
	// Whether the apikey can modify monitoring jobs.
	MonitoringManageJobs pulumi.BoolPtrInput
	// Whether the apikey can modify notification lists.
	MonitoringManageLists pulumi.BoolPtrInput
	// Whether the apikey can view monitoring jobs.
	MonitoringViewJobs pulumi.BoolPtrInput
	// The free form name of the apikey.
	Name pulumi.StringPtrInput
	// Whether the apikey can manage global active directory.
	// Only relevant for the DDI product.
	SecurityManageActiveDirectory pulumi.BoolPtrInput
	// Whether the apikey can manage global two factor authentication.
	SecurityManageGlobal2fa pulumi.BoolPtrInput
	// The teams that the apikey belongs to.
	Teams pulumi.StringArrayInput
}

func (APIKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*apikeyState)(nil)).Elem()
}

type apikeyArgs struct {
	// Whether the apikey can modify account settings.
	AccountManageAccountSettings *bool `pulumi:"accountManageAccountSettings"`
	// Whether the apikey can modify account apikeys.
	AccountManageApikeys *bool `pulumi:"accountManageApikeys"`
	// Whether the apikey can manage ip whitelist.
	AccountManageIpWhitelist *bool `pulumi:"accountManageIpWhitelist"`
	// Whether the apikey can modify account payment methods.
	AccountManagePaymentMethods *bool `pulumi:"accountManagePaymentMethods"`
	// Whether the apikey can modify the account plan.
	//
	// Deprecated: obsolete, should no longer be used
	AccountManagePlan *bool `pulumi:"accountManagePlan"`
	// Whether the apikey can modify other teams in the account.
	AccountManageTeams *bool `pulumi:"accountManageTeams"`
	// Whether the apikey can modify account users.
	AccountManageUsers *bool `pulumi:"accountManageUsers"`
	// Whether the apikey can view activity logs.
	AccountViewActivityLog *bool `pulumi:"accountViewActivityLog"`
	// Whether the apikey can view invoices.
	AccountViewInvoices *bool `pulumi:"accountViewInvoices"`
	// Whether the apikey can modify data feeds.
	DataManageDatafeeds *bool `pulumi:"dataManageDatafeeds"`
	// Whether the apikey can modify data sources.
	DataManageDatasources *bool `pulumi:"dataManageDatasources"`
	// Whether the apikey can publish to data feeds.
	DataPushToDatafeeds *bool `pulumi:"dataPushToDatafeeds"`
	// Whether the apikey can manage DHCP.
	// Only relevant for the DDI product.
	DhcpManageDhcp *bool `pulumi:"dhcpManageDhcp"`
	// Whether the apikey can view DHCP.
	// Only relevant for the DDI product.
	DhcpViewDhcp *bool `pulumi:"dhcpViewDhcp"`
	// Whether the apikey can modify the accounts zones.
	DnsManageZones *bool `pulumi:"dnsManageZones"`
	// Whether the apikey can view the accounts zones.
	DnsViewZones *bool `pulumi:"dnsViewZones"`
	// If true, enable the `dnsZonesAllow` list, otherwise enable the `dnsZonesDeny` list.
	DnsZonesAllowByDefault *bool `pulumi:"dnsZonesAllowByDefault"`
	// List of zones that the apikey may access.
	DnsZonesAllows []string `pulumi:"dnsZonesAllows"`
	// List of zones that the apikey may not access.
	DnsZonesDenies []string `pulumi:"dnsZonesDenies"`
	// Sets exclusivity on this IP whitelist.
	IpWhitelistStrict *bool `pulumi:"ipWhitelistStrict"`
	// The IP addresses to whitelist for this key.
	IpWhitelists []string `pulumi:"ipWhitelists"`
	// Whether the apikey can manage IPAM.
	// Only relevant for the DDI product.
	IpamManageIpam *bool `pulumi:"ipamManageIpam"`
	// Whether the apikey can view IPAM.
	// Only relevant for the DDI product.
	IpamViewIpam *bool `pulumi:"ipamViewIpam"`
	// Whether the apikey can modify monitoring jobs.
	MonitoringManageJobs *bool `pulumi:"monitoringManageJobs"`
	// Whether the apikey can modify notification lists.
	MonitoringManageLists *bool `pulumi:"monitoringManageLists"`
	// Whether the apikey can view monitoring jobs.
	MonitoringViewJobs *bool `pulumi:"monitoringViewJobs"`
	// The free form name of the apikey.
	Name *string `pulumi:"name"`
	// Whether the apikey can manage global active directory.
	// Only relevant for the DDI product.
	SecurityManageActiveDirectory *bool `pulumi:"securityManageActiveDirectory"`
	// Whether the apikey can manage global two factor authentication.
	SecurityManageGlobal2fa *bool `pulumi:"securityManageGlobal2fa"`
	// The teams that the apikey belongs to.
	Teams []string `pulumi:"teams"`
}

// The set of arguments for constructing a APIKey resource.
type APIKeyArgs struct {
	// Whether the apikey can modify account settings.
	AccountManageAccountSettings pulumi.BoolPtrInput
	// Whether the apikey can modify account apikeys.
	AccountManageApikeys pulumi.BoolPtrInput
	// Whether the apikey can manage ip whitelist.
	AccountManageIpWhitelist pulumi.BoolPtrInput
	// Whether the apikey can modify account payment methods.
	AccountManagePaymentMethods pulumi.BoolPtrInput
	// Whether the apikey can modify the account plan.
	//
	// Deprecated: obsolete, should no longer be used
	AccountManagePlan pulumi.BoolPtrInput
	// Whether the apikey can modify other teams in the account.
	AccountManageTeams pulumi.BoolPtrInput
	// Whether the apikey can modify account users.
	AccountManageUsers pulumi.BoolPtrInput
	// Whether the apikey can view activity logs.
	AccountViewActivityLog pulumi.BoolPtrInput
	// Whether the apikey can view invoices.
	AccountViewInvoices pulumi.BoolPtrInput
	// Whether the apikey can modify data feeds.
	DataManageDatafeeds pulumi.BoolPtrInput
	// Whether the apikey can modify data sources.
	DataManageDatasources pulumi.BoolPtrInput
	// Whether the apikey can publish to data feeds.
	DataPushToDatafeeds pulumi.BoolPtrInput
	// Whether the apikey can manage DHCP.
	// Only relevant for the DDI product.
	DhcpManageDhcp pulumi.BoolPtrInput
	// Whether the apikey can view DHCP.
	// Only relevant for the DDI product.
	DhcpViewDhcp pulumi.BoolPtrInput
	// Whether the apikey can modify the accounts zones.
	DnsManageZones pulumi.BoolPtrInput
	// Whether the apikey can view the accounts zones.
	DnsViewZones pulumi.BoolPtrInput
	// If true, enable the `dnsZonesAllow` list, otherwise enable the `dnsZonesDeny` list.
	DnsZonesAllowByDefault pulumi.BoolPtrInput
	// List of zones that the apikey may access.
	DnsZonesAllows pulumi.StringArrayInput
	// List of zones that the apikey may not access.
	DnsZonesDenies pulumi.StringArrayInput
	// Sets exclusivity on this IP whitelist.
	IpWhitelistStrict pulumi.BoolPtrInput
	// The IP addresses to whitelist for this key.
	IpWhitelists pulumi.StringArrayInput
	// Whether the apikey can manage IPAM.
	// Only relevant for the DDI product.
	IpamManageIpam pulumi.BoolPtrInput
	// Whether the apikey can view IPAM.
	// Only relevant for the DDI product.
	IpamViewIpam pulumi.BoolPtrInput
	// Whether the apikey can modify monitoring jobs.
	MonitoringManageJobs pulumi.BoolPtrInput
	// Whether the apikey can modify notification lists.
	MonitoringManageLists pulumi.BoolPtrInput
	// Whether the apikey can view monitoring jobs.
	MonitoringViewJobs pulumi.BoolPtrInput
	// The free form name of the apikey.
	Name pulumi.StringPtrInput
	// Whether the apikey can manage global active directory.
	// Only relevant for the DDI product.
	SecurityManageActiveDirectory pulumi.BoolPtrInput
	// Whether the apikey can manage global two factor authentication.
	SecurityManageGlobal2fa pulumi.BoolPtrInput
	// The teams that the apikey belongs to.
	Teams pulumi.StringArrayInput
}

func (APIKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apikeyArgs)(nil)).Elem()
}

type APIKeyInput interface {
	pulumi.Input

	ToAPIKeyOutput() APIKeyOutput
	ToAPIKeyOutputWithContext(ctx context.Context) APIKeyOutput
}

func (*APIKey) ElementType() reflect.Type {
	return reflect.TypeOf((*APIKey)(nil))
}

func (i *APIKey) ToAPIKeyOutput() APIKeyOutput {
	return i.ToAPIKeyOutputWithContext(context.Background())
}

func (i *APIKey) ToAPIKeyOutputWithContext(ctx context.Context) APIKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APIKeyOutput)
}

func (i *APIKey) ToAPIKeyPtrOutput() APIKeyPtrOutput {
	return i.ToAPIKeyPtrOutputWithContext(context.Background())
}

func (i *APIKey) ToAPIKeyPtrOutputWithContext(ctx context.Context) APIKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APIKeyPtrOutput)
}

type APIKeyPtrInput interface {
	pulumi.Input

	ToAPIKeyPtrOutput() APIKeyPtrOutput
	ToAPIKeyPtrOutputWithContext(ctx context.Context) APIKeyPtrOutput
}

type apikeyPtrType APIKeyArgs

func (*apikeyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**APIKey)(nil))
}

func (i *apikeyPtrType) ToAPIKeyPtrOutput() APIKeyPtrOutput {
	return i.ToAPIKeyPtrOutputWithContext(context.Background())
}

func (i *apikeyPtrType) ToAPIKeyPtrOutputWithContext(ctx context.Context) APIKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APIKeyPtrOutput)
}

// APIKeyArrayInput is an input type that accepts APIKeyArray and APIKeyArrayOutput values.
// You can construct a concrete instance of `APIKeyArrayInput` via:
//
//          APIKeyArray{ APIKeyArgs{...} }
type APIKeyArrayInput interface {
	pulumi.Input

	ToAPIKeyArrayOutput() APIKeyArrayOutput
	ToAPIKeyArrayOutputWithContext(context.Context) APIKeyArrayOutput
}

type APIKeyArray []APIKeyInput

func (APIKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*APIKey)(nil))
}

func (i APIKeyArray) ToAPIKeyArrayOutput() APIKeyArrayOutput {
	return i.ToAPIKeyArrayOutputWithContext(context.Background())
}

func (i APIKeyArray) ToAPIKeyArrayOutputWithContext(ctx context.Context) APIKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APIKeyArrayOutput)
}

// APIKeyMapInput is an input type that accepts APIKeyMap and APIKeyMapOutput values.
// You can construct a concrete instance of `APIKeyMapInput` via:
//
//          APIKeyMap{ "key": APIKeyArgs{...} }
type APIKeyMapInput interface {
	pulumi.Input

	ToAPIKeyMapOutput() APIKeyMapOutput
	ToAPIKeyMapOutputWithContext(context.Context) APIKeyMapOutput
}

type APIKeyMap map[string]APIKeyInput

func (APIKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*APIKey)(nil))
}

func (i APIKeyMap) ToAPIKeyMapOutput() APIKeyMapOutput {
	return i.ToAPIKeyMapOutputWithContext(context.Background())
}

func (i APIKeyMap) ToAPIKeyMapOutputWithContext(ctx context.Context) APIKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APIKeyMapOutput)
}

type APIKeyOutput struct {
	*pulumi.OutputState
}

func (APIKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*APIKey)(nil))
}

func (o APIKeyOutput) ToAPIKeyOutput() APIKeyOutput {
	return o
}

func (o APIKeyOutput) ToAPIKeyOutputWithContext(ctx context.Context) APIKeyOutput {
	return o
}

func (o APIKeyOutput) ToAPIKeyPtrOutput() APIKeyPtrOutput {
	return o.ToAPIKeyPtrOutputWithContext(context.Background())
}

func (o APIKeyOutput) ToAPIKeyPtrOutputWithContext(ctx context.Context) APIKeyPtrOutput {
	return o.ApplyT(func(v APIKey) *APIKey {
		return &v
	}).(APIKeyPtrOutput)
}

type APIKeyPtrOutput struct {
	*pulumi.OutputState
}

func (APIKeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**APIKey)(nil))
}

func (o APIKeyPtrOutput) ToAPIKeyPtrOutput() APIKeyPtrOutput {
	return o
}

func (o APIKeyPtrOutput) ToAPIKeyPtrOutputWithContext(ctx context.Context) APIKeyPtrOutput {
	return o
}

type APIKeyArrayOutput struct{ *pulumi.OutputState }

func (APIKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]APIKey)(nil))
}

func (o APIKeyArrayOutput) ToAPIKeyArrayOutput() APIKeyArrayOutput {
	return o
}

func (o APIKeyArrayOutput) ToAPIKeyArrayOutputWithContext(ctx context.Context) APIKeyArrayOutput {
	return o
}

func (o APIKeyArrayOutput) Index(i pulumi.IntInput) APIKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) APIKey {
		return vs[0].([]APIKey)[vs[1].(int)]
	}).(APIKeyOutput)
}

type APIKeyMapOutput struct{ *pulumi.OutputState }

func (APIKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]APIKey)(nil))
}

func (o APIKeyMapOutput) ToAPIKeyMapOutput() APIKeyMapOutput {
	return o
}

func (o APIKeyMapOutput) ToAPIKeyMapOutputWithContext(ctx context.Context) APIKeyMapOutput {
	return o
}

func (o APIKeyMapOutput) MapIndex(k pulumi.StringInput) APIKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) APIKey {
		return vs[0].(map[string]APIKey)[vs[1].(string)]
	}).(APIKeyOutput)
}

func init() {
	pulumi.RegisterOutputType(APIKeyOutput{})
	pulumi.RegisterOutputType(APIKeyPtrOutput{})
	pulumi.RegisterOutputType(APIKeyArrayOutput{})
	pulumi.RegisterOutputType(APIKeyMapOutput{})
}
