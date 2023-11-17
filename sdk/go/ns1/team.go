// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ns1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ns1/sdk/v3/go/ns1/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a NS1 Team resource. This can be used to create, modify, and delete
// teams. The credentials used must have the `manageTeams` permission set.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-ns1/sdk/v3/go/ns1"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ns1.NewTeam(ctx, "example", &ns1.TeamArgs{
//				AccountManageUsers: pulumi.Bool(false),
//				DnsViewZones:       pulumi.Bool(false),
//				IpWhitelists: ns1.TeamIpWhitelistArray{
//					&ns1.TeamIpWhitelistArgs{
//						Name: pulumi.String("whitelist-1"),
//						Values: pulumi.StringArray{
//							pulumi.String("1.1.1.1"),
//							pulumi.String("2.2.2.2"),
//						},
//					},
//					&ns1.TeamIpWhitelistArgs{
//						Name: pulumi.String("whitelist-2"),
//						Values: pulumi.StringArray{
//							pulumi.String("3.3.3.3"),
//							pulumi.String("4.4.4.4"),
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ns1.NewTeam(ctx, "example2", &ns1.TeamArgs{
//				DataManageDatasources: pulumi.Bool(true),
//				DnsRecordsAllows: ns1.TeamDnsRecordsAllowArray{
//					&ns1.TeamDnsRecordsAllowArgs{
//						Domain:            pulumi.String("terraform.example.io"),
//						IncludeSubdomains: pulumi.Bool(false),
//						Type:              pulumi.String("A"),
//						Zone:              pulumi.String("example.io"),
//					},
//				},
//				DnsViewZones: pulumi.Bool(true),
//				DnsZonesAllows: pulumi.StringArray{
//					pulumi.String("mytest.zone"),
//				},
//				DnsZonesAllowByDefault: pulumi.Bool(true),
//				DnsZonesDenies: pulumi.StringArray{
//					pulumi.String("myother.zone"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## NS1 Documentation
//
// [Team Api Docs](https://ns1.com/api#team)
//
// ## Import
//
// ```sh
//
//	$ pulumi import ns1:index/team:Team <name> <team_id>`
//
// ```
type Team struct {
	pulumi.CustomResourceState

	// Whether the team can modify account settings.
	AccountManageAccountSettings pulumi.BoolPtrOutput `pulumi:"accountManageAccountSettings"`
	// Whether the team can modify account apikeys.
	AccountManageApikeys pulumi.BoolPtrOutput `pulumi:"accountManageApikeys"`
	// Whether the team can manage ip whitelist.
	AccountManageIpWhitelist pulumi.BoolPtrOutput `pulumi:"accountManageIpWhitelist"`
	// Whether the team can modify account payment methods.
	AccountManagePaymentMethods pulumi.BoolPtrOutput `pulumi:"accountManagePaymentMethods"`
	// (Deprecated) No longer in use.
	//
	// Deprecated: obsolete, should no longer be used
	AccountManagePlan pulumi.BoolPtrOutput `pulumi:"accountManagePlan"`
	// Whether the team can modify other teams in the account.
	AccountManageTeams pulumi.BoolPtrOutput `pulumi:"accountManageTeams"`
	// Whether the team can modify account users.
	AccountManageUsers pulumi.BoolPtrOutput `pulumi:"accountManageUsers"`
	// Whether the team can view activity logs.
	AccountViewActivityLog pulumi.BoolPtrOutput `pulumi:"accountViewActivityLog"`
	// Whether the team can view invoices.
	AccountViewInvoices pulumi.BoolPtrOutput `pulumi:"accountViewInvoices"`
	// Whether the team can modify data feeds.
	DataManageDatafeeds pulumi.BoolPtrOutput `pulumi:"dataManageDatafeeds"`
	// Whether the team can modify data sources.
	DataManageDatasources pulumi.BoolPtrOutput `pulumi:"dataManageDatasources"`
	// Whether the team can publish to data feeds.
	DataPushToDatafeeds pulumi.BoolPtrOutput `pulumi:"dataPushToDatafeeds"`
	// Whether the team can manage DHCP.
	// Only relevant for the DDI product.
	DhcpManageDhcp pulumi.BoolPtrOutput `pulumi:"dhcpManageDhcp"`
	// Whether the team can view DHCP.
	// Only relevant for the DDI product.
	DhcpViewDhcp pulumi.BoolPtrOutput `pulumi:"dhcpViewDhcp"`
	// Whether the team can modify the accounts zones.
	DnsManageZones pulumi.BoolPtrOutput `pulumi:"dnsManageZones"`
	// List of records that the team may access.
	DnsRecordsAllows TeamDnsRecordsAllowArrayOutput `pulumi:"dnsRecordsAllows"`
	// List of records that the team may not access.
	DnsRecordsDenies TeamDnsRecordsDenyArrayOutput `pulumi:"dnsRecordsDenies"`
	// Whether the team can view the accounts zones.
	DnsViewZones pulumi.BoolPtrOutput `pulumi:"dnsViewZones"`
	// If true, enable the `dnsZonesAllow` list, otherwise enable the `dnsZonesDeny` list.
	DnsZonesAllowByDefault pulumi.BoolPtrOutput `pulumi:"dnsZonesAllowByDefault"`
	// List of zones that the team may access.
	DnsZonesAllows pulumi.StringArrayOutput `pulumi:"dnsZonesAllows"`
	// List of zones that the team may not access.
	DnsZonesDenies pulumi.StringArrayOutput `pulumi:"dnsZonesDenies"`
	// Array of IP addresses objects to chich to grant the team access. Each object includes a **name** (string), and **values** (array of strings) associated to each "allow" list.
	IpWhitelists TeamIpWhitelistArrayOutput `pulumi:"ipWhitelists"`
	// Whether the team can manage IPAM.
	// Only relevant for the DDI product.
	IpamManageIpam pulumi.BoolPtrOutput `pulumi:"ipamManageIpam"`
	// Whether the team can view IPAM.
	// Only relevant for the DDI product.
	IpamViewIpam pulumi.BoolPtrOutput `pulumi:"ipamViewIpam"`
	// Whether the team can modify monitoring jobs.
	MonitoringManageJobs pulumi.BoolPtrOutput `pulumi:"monitoringManageJobs"`
	// Whether the team can modify notification lists.
	MonitoringManageLists pulumi.BoolPtrOutput `pulumi:"monitoringManageLists"`
	// Whether the team can view monitoring jobs.
	MonitoringViewJobs pulumi.BoolPtrOutput `pulumi:"monitoringViewJobs"`
	// The free form name of the team.
	Name pulumi.StringOutput `pulumi:"name"`
	// Whether the team can manage global active directory.
	// Only relevant for the DDI product.
	SecurityManageActiveDirectory pulumi.BoolPtrOutput `pulumi:"securityManageActiveDirectory"`
	// Whether the team can manage global two factor authentication.
	SecurityManageGlobal2fa pulumi.BoolPtrOutput `pulumi:"securityManageGlobal2fa"`
}

// NewTeam registers a new resource with the given unique name, arguments, and options.
func NewTeam(ctx *pulumi.Context,
	name string, args *TeamArgs, opts ...pulumi.ResourceOption) (*Team, error) {
	if args == nil {
		args = &TeamArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Team
	err := ctx.RegisterResource("ns1:index/team:Team", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTeam gets an existing Team resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTeam(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TeamState, opts ...pulumi.ResourceOption) (*Team, error) {
	var resource Team
	err := ctx.ReadResource("ns1:index/team:Team", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Team resources.
type teamState struct {
	// Whether the team can modify account settings.
	AccountManageAccountSettings *bool `pulumi:"accountManageAccountSettings"`
	// Whether the team can modify account apikeys.
	AccountManageApikeys *bool `pulumi:"accountManageApikeys"`
	// Whether the team can manage ip whitelist.
	AccountManageIpWhitelist *bool `pulumi:"accountManageIpWhitelist"`
	// Whether the team can modify account payment methods.
	AccountManagePaymentMethods *bool `pulumi:"accountManagePaymentMethods"`
	// (Deprecated) No longer in use.
	//
	// Deprecated: obsolete, should no longer be used
	AccountManagePlan *bool `pulumi:"accountManagePlan"`
	// Whether the team can modify other teams in the account.
	AccountManageTeams *bool `pulumi:"accountManageTeams"`
	// Whether the team can modify account users.
	AccountManageUsers *bool `pulumi:"accountManageUsers"`
	// Whether the team can view activity logs.
	AccountViewActivityLog *bool `pulumi:"accountViewActivityLog"`
	// Whether the team can view invoices.
	AccountViewInvoices *bool `pulumi:"accountViewInvoices"`
	// Whether the team can modify data feeds.
	DataManageDatafeeds *bool `pulumi:"dataManageDatafeeds"`
	// Whether the team can modify data sources.
	DataManageDatasources *bool `pulumi:"dataManageDatasources"`
	// Whether the team can publish to data feeds.
	DataPushToDatafeeds *bool `pulumi:"dataPushToDatafeeds"`
	// Whether the team can manage DHCP.
	// Only relevant for the DDI product.
	DhcpManageDhcp *bool `pulumi:"dhcpManageDhcp"`
	// Whether the team can view DHCP.
	// Only relevant for the DDI product.
	DhcpViewDhcp *bool `pulumi:"dhcpViewDhcp"`
	// Whether the team can modify the accounts zones.
	DnsManageZones *bool `pulumi:"dnsManageZones"`
	// List of records that the team may access.
	DnsRecordsAllows []TeamDnsRecordsAllow `pulumi:"dnsRecordsAllows"`
	// List of records that the team may not access.
	DnsRecordsDenies []TeamDnsRecordsDeny `pulumi:"dnsRecordsDenies"`
	// Whether the team can view the accounts zones.
	DnsViewZones *bool `pulumi:"dnsViewZones"`
	// If true, enable the `dnsZonesAllow` list, otherwise enable the `dnsZonesDeny` list.
	DnsZonesAllowByDefault *bool `pulumi:"dnsZonesAllowByDefault"`
	// List of zones that the team may access.
	DnsZonesAllows []string `pulumi:"dnsZonesAllows"`
	// List of zones that the team may not access.
	DnsZonesDenies []string `pulumi:"dnsZonesDenies"`
	// Array of IP addresses objects to chich to grant the team access. Each object includes a **name** (string), and **values** (array of strings) associated to each "allow" list.
	IpWhitelists []TeamIpWhitelist `pulumi:"ipWhitelists"`
	// Whether the team can manage IPAM.
	// Only relevant for the DDI product.
	IpamManageIpam *bool `pulumi:"ipamManageIpam"`
	// Whether the team can view IPAM.
	// Only relevant for the DDI product.
	IpamViewIpam *bool `pulumi:"ipamViewIpam"`
	// Whether the team can modify monitoring jobs.
	MonitoringManageJobs *bool `pulumi:"monitoringManageJobs"`
	// Whether the team can modify notification lists.
	MonitoringManageLists *bool `pulumi:"monitoringManageLists"`
	// Whether the team can view monitoring jobs.
	MonitoringViewJobs *bool `pulumi:"monitoringViewJobs"`
	// The free form name of the team.
	Name *string `pulumi:"name"`
	// Whether the team can manage global active directory.
	// Only relevant for the DDI product.
	SecurityManageActiveDirectory *bool `pulumi:"securityManageActiveDirectory"`
	// Whether the team can manage global two factor authentication.
	SecurityManageGlobal2fa *bool `pulumi:"securityManageGlobal2fa"`
}

type TeamState struct {
	// Whether the team can modify account settings.
	AccountManageAccountSettings pulumi.BoolPtrInput
	// Whether the team can modify account apikeys.
	AccountManageApikeys pulumi.BoolPtrInput
	// Whether the team can manage ip whitelist.
	AccountManageIpWhitelist pulumi.BoolPtrInput
	// Whether the team can modify account payment methods.
	AccountManagePaymentMethods pulumi.BoolPtrInput
	// (Deprecated) No longer in use.
	//
	// Deprecated: obsolete, should no longer be used
	AccountManagePlan pulumi.BoolPtrInput
	// Whether the team can modify other teams in the account.
	AccountManageTeams pulumi.BoolPtrInput
	// Whether the team can modify account users.
	AccountManageUsers pulumi.BoolPtrInput
	// Whether the team can view activity logs.
	AccountViewActivityLog pulumi.BoolPtrInput
	// Whether the team can view invoices.
	AccountViewInvoices pulumi.BoolPtrInput
	// Whether the team can modify data feeds.
	DataManageDatafeeds pulumi.BoolPtrInput
	// Whether the team can modify data sources.
	DataManageDatasources pulumi.BoolPtrInput
	// Whether the team can publish to data feeds.
	DataPushToDatafeeds pulumi.BoolPtrInput
	// Whether the team can manage DHCP.
	// Only relevant for the DDI product.
	DhcpManageDhcp pulumi.BoolPtrInput
	// Whether the team can view DHCP.
	// Only relevant for the DDI product.
	DhcpViewDhcp pulumi.BoolPtrInput
	// Whether the team can modify the accounts zones.
	DnsManageZones pulumi.BoolPtrInput
	// List of records that the team may access.
	DnsRecordsAllows TeamDnsRecordsAllowArrayInput
	// List of records that the team may not access.
	DnsRecordsDenies TeamDnsRecordsDenyArrayInput
	// Whether the team can view the accounts zones.
	DnsViewZones pulumi.BoolPtrInput
	// If true, enable the `dnsZonesAllow` list, otherwise enable the `dnsZonesDeny` list.
	DnsZonesAllowByDefault pulumi.BoolPtrInput
	// List of zones that the team may access.
	DnsZonesAllows pulumi.StringArrayInput
	// List of zones that the team may not access.
	DnsZonesDenies pulumi.StringArrayInput
	// Array of IP addresses objects to chich to grant the team access. Each object includes a **name** (string), and **values** (array of strings) associated to each "allow" list.
	IpWhitelists TeamIpWhitelistArrayInput
	// Whether the team can manage IPAM.
	// Only relevant for the DDI product.
	IpamManageIpam pulumi.BoolPtrInput
	// Whether the team can view IPAM.
	// Only relevant for the DDI product.
	IpamViewIpam pulumi.BoolPtrInput
	// Whether the team can modify monitoring jobs.
	MonitoringManageJobs pulumi.BoolPtrInput
	// Whether the team can modify notification lists.
	MonitoringManageLists pulumi.BoolPtrInput
	// Whether the team can view monitoring jobs.
	MonitoringViewJobs pulumi.BoolPtrInput
	// The free form name of the team.
	Name pulumi.StringPtrInput
	// Whether the team can manage global active directory.
	// Only relevant for the DDI product.
	SecurityManageActiveDirectory pulumi.BoolPtrInput
	// Whether the team can manage global two factor authentication.
	SecurityManageGlobal2fa pulumi.BoolPtrInput
}

func (TeamState) ElementType() reflect.Type {
	return reflect.TypeOf((*teamState)(nil)).Elem()
}

type teamArgs struct {
	// Whether the team can modify account settings.
	AccountManageAccountSettings *bool `pulumi:"accountManageAccountSettings"`
	// Whether the team can modify account apikeys.
	AccountManageApikeys *bool `pulumi:"accountManageApikeys"`
	// Whether the team can manage ip whitelist.
	AccountManageIpWhitelist *bool `pulumi:"accountManageIpWhitelist"`
	// Whether the team can modify account payment methods.
	AccountManagePaymentMethods *bool `pulumi:"accountManagePaymentMethods"`
	// (Deprecated) No longer in use.
	//
	// Deprecated: obsolete, should no longer be used
	AccountManagePlan *bool `pulumi:"accountManagePlan"`
	// Whether the team can modify other teams in the account.
	AccountManageTeams *bool `pulumi:"accountManageTeams"`
	// Whether the team can modify account users.
	AccountManageUsers *bool `pulumi:"accountManageUsers"`
	// Whether the team can view activity logs.
	AccountViewActivityLog *bool `pulumi:"accountViewActivityLog"`
	// Whether the team can view invoices.
	AccountViewInvoices *bool `pulumi:"accountViewInvoices"`
	// Whether the team can modify data feeds.
	DataManageDatafeeds *bool `pulumi:"dataManageDatafeeds"`
	// Whether the team can modify data sources.
	DataManageDatasources *bool `pulumi:"dataManageDatasources"`
	// Whether the team can publish to data feeds.
	DataPushToDatafeeds *bool `pulumi:"dataPushToDatafeeds"`
	// Whether the team can manage DHCP.
	// Only relevant for the DDI product.
	DhcpManageDhcp *bool `pulumi:"dhcpManageDhcp"`
	// Whether the team can view DHCP.
	// Only relevant for the DDI product.
	DhcpViewDhcp *bool `pulumi:"dhcpViewDhcp"`
	// Whether the team can modify the accounts zones.
	DnsManageZones *bool `pulumi:"dnsManageZones"`
	// List of records that the team may access.
	DnsRecordsAllows []TeamDnsRecordsAllow `pulumi:"dnsRecordsAllows"`
	// List of records that the team may not access.
	DnsRecordsDenies []TeamDnsRecordsDeny `pulumi:"dnsRecordsDenies"`
	// Whether the team can view the accounts zones.
	DnsViewZones *bool `pulumi:"dnsViewZones"`
	// If true, enable the `dnsZonesAllow` list, otherwise enable the `dnsZonesDeny` list.
	DnsZonesAllowByDefault *bool `pulumi:"dnsZonesAllowByDefault"`
	// List of zones that the team may access.
	DnsZonesAllows []string `pulumi:"dnsZonesAllows"`
	// List of zones that the team may not access.
	DnsZonesDenies []string `pulumi:"dnsZonesDenies"`
	// Array of IP addresses objects to chich to grant the team access. Each object includes a **name** (string), and **values** (array of strings) associated to each "allow" list.
	IpWhitelists []TeamIpWhitelist `pulumi:"ipWhitelists"`
	// Whether the team can manage IPAM.
	// Only relevant for the DDI product.
	IpamManageIpam *bool `pulumi:"ipamManageIpam"`
	// Whether the team can view IPAM.
	// Only relevant for the DDI product.
	IpamViewIpam *bool `pulumi:"ipamViewIpam"`
	// Whether the team can modify monitoring jobs.
	MonitoringManageJobs *bool `pulumi:"monitoringManageJobs"`
	// Whether the team can modify notification lists.
	MonitoringManageLists *bool `pulumi:"monitoringManageLists"`
	// Whether the team can view monitoring jobs.
	MonitoringViewJobs *bool `pulumi:"monitoringViewJobs"`
	// The free form name of the team.
	Name *string `pulumi:"name"`
	// Whether the team can manage global active directory.
	// Only relevant for the DDI product.
	SecurityManageActiveDirectory *bool `pulumi:"securityManageActiveDirectory"`
	// Whether the team can manage global two factor authentication.
	SecurityManageGlobal2fa *bool `pulumi:"securityManageGlobal2fa"`
}

// The set of arguments for constructing a Team resource.
type TeamArgs struct {
	// Whether the team can modify account settings.
	AccountManageAccountSettings pulumi.BoolPtrInput
	// Whether the team can modify account apikeys.
	AccountManageApikeys pulumi.BoolPtrInput
	// Whether the team can manage ip whitelist.
	AccountManageIpWhitelist pulumi.BoolPtrInput
	// Whether the team can modify account payment methods.
	AccountManagePaymentMethods pulumi.BoolPtrInput
	// (Deprecated) No longer in use.
	//
	// Deprecated: obsolete, should no longer be used
	AccountManagePlan pulumi.BoolPtrInput
	// Whether the team can modify other teams in the account.
	AccountManageTeams pulumi.BoolPtrInput
	// Whether the team can modify account users.
	AccountManageUsers pulumi.BoolPtrInput
	// Whether the team can view activity logs.
	AccountViewActivityLog pulumi.BoolPtrInput
	// Whether the team can view invoices.
	AccountViewInvoices pulumi.BoolPtrInput
	// Whether the team can modify data feeds.
	DataManageDatafeeds pulumi.BoolPtrInput
	// Whether the team can modify data sources.
	DataManageDatasources pulumi.BoolPtrInput
	// Whether the team can publish to data feeds.
	DataPushToDatafeeds pulumi.BoolPtrInput
	// Whether the team can manage DHCP.
	// Only relevant for the DDI product.
	DhcpManageDhcp pulumi.BoolPtrInput
	// Whether the team can view DHCP.
	// Only relevant for the DDI product.
	DhcpViewDhcp pulumi.BoolPtrInput
	// Whether the team can modify the accounts zones.
	DnsManageZones pulumi.BoolPtrInput
	// List of records that the team may access.
	DnsRecordsAllows TeamDnsRecordsAllowArrayInput
	// List of records that the team may not access.
	DnsRecordsDenies TeamDnsRecordsDenyArrayInput
	// Whether the team can view the accounts zones.
	DnsViewZones pulumi.BoolPtrInput
	// If true, enable the `dnsZonesAllow` list, otherwise enable the `dnsZonesDeny` list.
	DnsZonesAllowByDefault pulumi.BoolPtrInput
	// List of zones that the team may access.
	DnsZonesAllows pulumi.StringArrayInput
	// List of zones that the team may not access.
	DnsZonesDenies pulumi.StringArrayInput
	// Array of IP addresses objects to chich to grant the team access. Each object includes a **name** (string), and **values** (array of strings) associated to each "allow" list.
	IpWhitelists TeamIpWhitelistArrayInput
	// Whether the team can manage IPAM.
	// Only relevant for the DDI product.
	IpamManageIpam pulumi.BoolPtrInput
	// Whether the team can view IPAM.
	// Only relevant for the DDI product.
	IpamViewIpam pulumi.BoolPtrInput
	// Whether the team can modify monitoring jobs.
	MonitoringManageJobs pulumi.BoolPtrInput
	// Whether the team can modify notification lists.
	MonitoringManageLists pulumi.BoolPtrInput
	// Whether the team can view monitoring jobs.
	MonitoringViewJobs pulumi.BoolPtrInput
	// The free form name of the team.
	Name pulumi.StringPtrInput
	// Whether the team can manage global active directory.
	// Only relevant for the DDI product.
	SecurityManageActiveDirectory pulumi.BoolPtrInput
	// Whether the team can manage global two factor authentication.
	SecurityManageGlobal2fa pulumi.BoolPtrInput
}

func (TeamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*teamArgs)(nil)).Elem()
}

type TeamInput interface {
	pulumi.Input

	ToTeamOutput() TeamOutput
	ToTeamOutputWithContext(ctx context.Context) TeamOutput
}

func (*Team) ElementType() reflect.Type {
	return reflect.TypeOf((**Team)(nil)).Elem()
}

func (i *Team) ToTeamOutput() TeamOutput {
	return i.ToTeamOutputWithContext(context.Background())
}

func (i *Team) ToTeamOutputWithContext(ctx context.Context) TeamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamOutput)
}

// TeamArrayInput is an input type that accepts TeamArray and TeamArrayOutput values.
// You can construct a concrete instance of `TeamArrayInput` via:
//
//	TeamArray{ TeamArgs{...} }
type TeamArrayInput interface {
	pulumi.Input

	ToTeamArrayOutput() TeamArrayOutput
	ToTeamArrayOutputWithContext(context.Context) TeamArrayOutput
}

type TeamArray []TeamInput

func (TeamArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Team)(nil)).Elem()
}

func (i TeamArray) ToTeamArrayOutput() TeamArrayOutput {
	return i.ToTeamArrayOutputWithContext(context.Background())
}

func (i TeamArray) ToTeamArrayOutputWithContext(ctx context.Context) TeamArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamArrayOutput)
}

// TeamMapInput is an input type that accepts TeamMap and TeamMapOutput values.
// You can construct a concrete instance of `TeamMapInput` via:
//
//	TeamMap{ "key": TeamArgs{...} }
type TeamMapInput interface {
	pulumi.Input

	ToTeamMapOutput() TeamMapOutput
	ToTeamMapOutputWithContext(context.Context) TeamMapOutput
}

type TeamMap map[string]TeamInput

func (TeamMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Team)(nil)).Elem()
}

func (i TeamMap) ToTeamMapOutput() TeamMapOutput {
	return i.ToTeamMapOutputWithContext(context.Background())
}

func (i TeamMap) ToTeamMapOutputWithContext(ctx context.Context) TeamMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamMapOutput)
}

type TeamOutput struct{ *pulumi.OutputState }

func (TeamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Team)(nil)).Elem()
}

func (o TeamOutput) ToTeamOutput() TeamOutput {
	return o
}

func (o TeamOutput) ToTeamOutputWithContext(ctx context.Context) TeamOutput {
	return o
}

// Whether the team can modify account settings.
func (o TeamOutput) AccountManageAccountSettings() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolPtrOutput { return v.AccountManageAccountSettings }).(pulumi.BoolPtrOutput)
}

// Whether the team can modify account apikeys.
func (o TeamOutput) AccountManageApikeys() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolPtrOutput { return v.AccountManageApikeys }).(pulumi.BoolPtrOutput)
}

// Whether the team can manage ip whitelist.
func (o TeamOutput) AccountManageIpWhitelist() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolPtrOutput { return v.AccountManageIpWhitelist }).(pulumi.BoolPtrOutput)
}

// Whether the team can modify account payment methods.
func (o TeamOutput) AccountManagePaymentMethods() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolPtrOutput { return v.AccountManagePaymentMethods }).(pulumi.BoolPtrOutput)
}

// (Deprecated) No longer in use.
//
// Deprecated: obsolete, should no longer be used
func (o TeamOutput) AccountManagePlan() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolPtrOutput { return v.AccountManagePlan }).(pulumi.BoolPtrOutput)
}

// Whether the team can modify other teams in the account.
func (o TeamOutput) AccountManageTeams() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolPtrOutput { return v.AccountManageTeams }).(pulumi.BoolPtrOutput)
}

// Whether the team can modify account users.
func (o TeamOutput) AccountManageUsers() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolPtrOutput { return v.AccountManageUsers }).(pulumi.BoolPtrOutput)
}

// Whether the team can view activity logs.
func (o TeamOutput) AccountViewActivityLog() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolPtrOutput { return v.AccountViewActivityLog }).(pulumi.BoolPtrOutput)
}

// Whether the team can view invoices.
func (o TeamOutput) AccountViewInvoices() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolPtrOutput { return v.AccountViewInvoices }).(pulumi.BoolPtrOutput)
}

// Whether the team can modify data feeds.
func (o TeamOutput) DataManageDatafeeds() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolPtrOutput { return v.DataManageDatafeeds }).(pulumi.BoolPtrOutput)
}

// Whether the team can modify data sources.
func (o TeamOutput) DataManageDatasources() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolPtrOutput { return v.DataManageDatasources }).(pulumi.BoolPtrOutput)
}

// Whether the team can publish to data feeds.
func (o TeamOutput) DataPushToDatafeeds() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolPtrOutput { return v.DataPushToDatafeeds }).(pulumi.BoolPtrOutput)
}

// Whether the team can manage DHCP.
// Only relevant for the DDI product.
func (o TeamOutput) DhcpManageDhcp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolPtrOutput { return v.DhcpManageDhcp }).(pulumi.BoolPtrOutput)
}

// Whether the team can view DHCP.
// Only relevant for the DDI product.
func (o TeamOutput) DhcpViewDhcp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolPtrOutput { return v.DhcpViewDhcp }).(pulumi.BoolPtrOutput)
}

// Whether the team can modify the accounts zones.
func (o TeamOutput) DnsManageZones() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolPtrOutput { return v.DnsManageZones }).(pulumi.BoolPtrOutput)
}

// List of records that the team may access.
func (o TeamOutput) DnsRecordsAllows() TeamDnsRecordsAllowArrayOutput {
	return o.ApplyT(func(v *Team) TeamDnsRecordsAllowArrayOutput { return v.DnsRecordsAllows }).(TeamDnsRecordsAllowArrayOutput)
}

// List of records that the team may not access.
func (o TeamOutput) DnsRecordsDenies() TeamDnsRecordsDenyArrayOutput {
	return o.ApplyT(func(v *Team) TeamDnsRecordsDenyArrayOutput { return v.DnsRecordsDenies }).(TeamDnsRecordsDenyArrayOutput)
}

// Whether the team can view the accounts zones.
func (o TeamOutput) DnsViewZones() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolPtrOutput { return v.DnsViewZones }).(pulumi.BoolPtrOutput)
}

// If true, enable the `dnsZonesAllow` list, otherwise enable the `dnsZonesDeny` list.
func (o TeamOutput) DnsZonesAllowByDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolPtrOutput { return v.DnsZonesAllowByDefault }).(pulumi.BoolPtrOutput)
}

// List of zones that the team may access.
func (o TeamOutput) DnsZonesAllows() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Team) pulumi.StringArrayOutput { return v.DnsZonesAllows }).(pulumi.StringArrayOutput)
}

// List of zones that the team may not access.
func (o TeamOutput) DnsZonesDenies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Team) pulumi.StringArrayOutput { return v.DnsZonesDenies }).(pulumi.StringArrayOutput)
}

// Array of IP addresses objects to chich to grant the team access. Each object includes a **name** (string), and **values** (array of strings) associated to each "allow" list.
func (o TeamOutput) IpWhitelists() TeamIpWhitelistArrayOutput {
	return o.ApplyT(func(v *Team) TeamIpWhitelistArrayOutput { return v.IpWhitelists }).(TeamIpWhitelistArrayOutput)
}

// Whether the team can manage IPAM.
// Only relevant for the DDI product.
func (o TeamOutput) IpamManageIpam() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolPtrOutput { return v.IpamManageIpam }).(pulumi.BoolPtrOutput)
}

// Whether the team can view IPAM.
// Only relevant for the DDI product.
func (o TeamOutput) IpamViewIpam() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolPtrOutput { return v.IpamViewIpam }).(pulumi.BoolPtrOutput)
}

// Whether the team can modify monitoring jobs.
func (o TeamOutput) MonitoringManageJobs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolPtrOutput { return v.MonitoringManageJobs }).(pulumi.BoolPtrOutput)
}

// Whether the team can modify notification lists.
func (o TeamOutput) MonitoringManageLists() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolPtrOutput { return v.MonitoringManageLists }).(pulumi.BoolPtrOutput)
}

// Whether the team can view monitoring jobs.
func (o TeamOutput) MonitoringViewJobs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolPtrOutput { return v.MonitoringViewJobs }).(pulumi.BoolPtrOutput)
}

// The free form name of the team.
func (o TeamOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Team) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Whether the team can manage global active directory.
// Only relevant for the DDI product.
func (o TeamOutput) SecurityManageActiveDirectory() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolPtrOutput { return v.SecurityManageActiveDirectory }).(pulumi.BoolPtrOutput)
}

// Whether the team can manage global two factor authentication.
func (o TeamOutput) SecurityManageGlobal2fa() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolPtrOutput { return v.SecurityManageGlobal2fa }).(pulumi.BoolPtrOutput)
}

type TeamArrayOutput struct{ *pulumi.OutputState }

func (TeamArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Team)(nil)).Elem()
}

func (o TeamArrayOutput) ToTeamArrayOutput() TeamArrayOutput {
	return o
}

func (o TeamArrayOutput) ToTeamArrayOutputWithContext(ctx context.Context) TeamArrayOutput {
	return o
}

func (o TeamArrayOutput) Index(i pulumi.IntInput) TeamOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Team {
		return vs[0].([]*Team)[vs[1].(int)]
	}).(TeamOutput)
}

type TeamMapOutput struct{ *pulumi.OutputState }

func (TeamMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Team)(nil)).Elem()
}

func (o TeamMapOutput) ToTeamMapOutput() TeamMapOutput {
	return o
}

func (o TeamMapOutput) ToTeamMapOutputWithContext(ctx context.Context) TeamMapOutput {
	return o
}

func (o TeamMapOutput) MapIndex(k pulumi.StringInput) TeamOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Team {
		return vs[0].(map[string]*Team)[vs[1].(string)]
	}).(TeamOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TeamInput)(nil)).Elem(), &Team{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamArrayInput)(nil)).Elem(), TeamArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamMapInput)(nil)).Elem(), TeamMap{})
	pulumi.RegisterOutputType(TeamOutput{})
	pulumi.RegisterOutputType(TeamArrayOutput{})
	pulumi.RegisterOutputType(TeamMapOutput{})
}
