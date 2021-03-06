// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ns1

import (
	"context"
	"reflect"

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
// 	"github.com/pulumi/pulumi-ns1/sdk/v2/go/ns1"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ns1.NewTeam(ctx, "example", &ns1.TeamArgs{
// 			AccountManageUsers: pulumi.Bool(false),
// 			DnsViewZones:       pulumi.Bool(false),
// 			IpWhitelists: ns1.TeamIpWhitelistArray{
// 				&ns1.TeamIpWhitelistArgs{
// 					Name: pulumi.String("whitelist-1"),
// 					Values: pulumi.StringArray{
// 						pulumi.String("1.1.1.1"),
// 						pulumi.String("2.2.2.2"),
// 					},
// 				},
// 				&ns1.TeamIpWhitelistArgs{
// 					Name: pulumi.String("whitelist-2"),
// 					Values: pulumi.StringArray{
// 						pulumi.String("3.3.3.3"),
// 						pulumi.String("4.4.4.4"),
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ns1.NewTeam(ctx, "example2", &ns1.TeamArgs{
// 			DataManageDatasources: pulumi.Bool(true),
// 			DnsViewZones:          pulumi.Bool(true),
// 			DnsZonesAllows: pulumi.StringArray{
// 				pulumi.String("mytest.zone"),
// 			},
// 			DnsZonesAllowByDefault: pulumi.Bool(true),
// 			DnsZonesDenies: pulumi.StringArray{
// 				pulumi.String("myother.zone"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## NS1 Documentation
//
// [Team Api Docs](https://ns1.com/api#team)
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
	// Whether the team can modify the account plan.
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
	// Whether the team can view the accounts zones.
	DnsViewZones pulumi.BoolPtrOutput `pulumi:"dnsViewZones"`
	// If true, enable the `dnsZonesAllow` list, otherwise enable the `dnsZonesDeny` list.
	DnsZonesAllowByDefault pulumi.BoolPtrOutput `pulumi:"dnsZonesAllowByDefault"`
	// List of zones that the team may access.
	DnsZonesAllows pulumi.StringArrayOutput `pulumi:"dnsZonesAllows"`
	// List of zones that the team may not access.
	DnsZonesDenies pulumi.StringArrayOutput `pulumi:"dnsZonesDenies"`
	// The IP addresses to whitelist for this key.
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
	// Whether the team can modify the account plan.
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
	// Whether the team can view the accounts zones.
	DnsViewZones *bool `pulumi:"dnsViewZones"`
	// If true, enable the `dnsZonesAllow` list, otherwise enable the `dnsZonesDeny` list.
	DnsZonesAllowByDefault *bool `pulumi:"dnsZonesAllowByDefault"`
	// List of zones that the team may access.
	DnsZonesAllows []string `pulumi:"dnsZonesAllows"`
	// List of zones that the team may not access.
	DnsZonesDenies []string `pulumi:"dnsZonesDenies"`
	// The IP addresses to whitelist for this key.
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
	// Whether the team can modify the account plan.
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
	// Whether the team can view the accounts zones.
	DnsViewZones pulumi.BoolPtrInput
	// If true, enable the `dnsZonesAllow` list, otherwise enable the `dnsZonesDeny` list.
	DnsZonesAllowByDefault pulumi.BoolPtrInput
	// List of zones that the team may access.
	DnsZonesAllows pulumi.StringArrayInput
	// List of zones that the team may not access.
	DnsZonesDenies pulumi.StringArrayInput
	// The IP addresses to whitelist for this key.
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
	// Whether the team can modify the account plan.
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
	// Whether the team can view the accounts zones.
	DnsViewZones *bool `pulumi:"dnsViewZones"`
	// If true, enable the `dnsZonesAllow` list, otherwise enable the `dnsZonesDeny` list.
	DnsZonesAllowByDefault *bool `pulumi:"dnsZonesAllowByDefault"`
	// List of zones that the team may access.
	DnsZonesAllows []string `pulumi:"dnsZonesAllows"`
	// List of zones that the team may not access.
	DnsZonesDenies []string `pulumi:"dnsZonesDenies"`
	// The IP addresses to whitelist for this key.
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
	// Whether the team can modify the account plan.
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
	// Whether the team can view the accounts zones.
	DnsViewZones pulumi.BoolPtrInput
	// If true, enable the `dnsZonesAllow` list, otherwise enable the `dnsZonesDeny` list.
	DnsZonesAllowByDefault pulumi.BoolPtrInput
	// List of zones that the team may access.
	DnsZonesAllows pulumi.StringArrayInput
	// List of zones that the team may not access.
	DnsZonesDenies pulumi.StringArrayInput
	// The IP addresses to whitelist for this key.
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
	return reflect.TypeOf((*Team)(nil))
}

func (i *Team) ToTeamOutput() TeamOutput {
	return i.ToTeamOutputWithContext(context.Background())
}

func (i *Team) ToTeamOutputWithContext(ctx context.Context) TeamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamOutput)
}

func (i *Team) ToTeamPtrOutput() TeamPtrOutput {
	return i.ToTeamPtrOutputWithContext(context.Background())
}

func (i *Team) ToTeamPtrOutputWithContext(ctx context.Context) TeamPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamPtrOutput)
}

type TeamPtrInput interface {
	pulumi.Input

	ToTeamPtrOutput() TeamPtrOutput
	ToTeamPtrOutputWithContext(ctx context.Context) TeamPtrOutput
}

type teamPtrType TeamArgs

func (*teamPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Team)(nil))
}

func (i *teamPtrType) ToTeamPtrOutput() TeamPtrOutput {
	return i.ToTeamPtrOutputWithContext(context.Background())
}

func (i *teamPtrType) ToTeamPtrOutputWithContext(ctx context.Context) TeamPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamPtrOutput)
}

// TeamArrayInput is an input type that accepts TeamArray and TeamArrayOutput values.
// You can construct a concrete instance of `TeamArrayInput` via:
//
//          TeamArray{ TeamArgs{...} }
type TeamArrayInput interface {
	pulumi.Input

	ToTeamArrayOutput() TeamArrayOutput
	ToTeamArrayOutputWithContext(context.Context) TeamArrayOutput
}

type TeamArray []TeamInput

func (TeamArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Team)(nil))
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
//          TeamMap{ "key": TeamArgs{...} }
type TeamMapInput interface {
	pulumi.Input

	ToTeamMapOutput() TeamMapOutput
	ToTeamMapOutputWithContext(context.Context) TeamMapOutput
}

type TeamMap map[string]TeamInput

func (TeamMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Team)(nil))
}

func (i TeamMap) ToTeamMapOutput() TeamMapOutput {
	return i.ToTeamMapOutputWithContext(context.Background())
}

func (i TeamMap) ToTeamMapOutputWithContext(ctx context.Context) TeamMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamMapOutput)
}

type TeamOutput struct {
	*pulumi.OutputState
}

func (TeamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Team)(nil))
}

func (o TeamOutput) ToTeamOutput() TeamOutput {
	return o
}

func (o TeamOutput) ToTeamOutputWithContext(ctx context.Context) TeamOutput {
	return o
}

func (o TeamOutput) ToTeamPtrOutput() TeamPtrOutput {
	return o.ToTeamPtrOutputWithContext(context.Background())
}

func (o TeamOutput) ToTeamPtrOutputWithContext(ctx context.Context) TeamPtrOutput {
	return o.ApplyT(func(v Team) *Team {
		return &v
	}).(TeamPtrOutput)
}

type TeamPtrOutput struct {
	*pulumi.OutputState
}

func (TeamPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Team)(nil))
}

func (o TeamPtrOutput) ToTeamPtrOutput() TeamPtrOutput {
	return o
}

func (o TeamPtrOutput) ToTeamPtrOutputWithContext(ctx context.Context) TeamPtrOutput {
	return o
}

type TeamArrayOutput struct{ *pulumi.OutputState }

func (TeamArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Team)(nil))
}

func (o TeamArrayOutput) ToTeamArrayOutput() TeamArrayOutput {
	return o
}

func (o TeamArrayOutput) ToTeamArrayOutputWithContext(ctx context.Context) TeamArrayOutput {
	return o
}

func (o TeamArrayOutput) Index(i pulumi.IntInput) TeamOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Team {
		return vs[0].([]Team)[vs[1].(int)]
	}).(TeamOutput)
}

type TeamMapOutput struct{ *pulumi.OutputState }

func (TeamMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Team)(nil))
}

func (o TeamMapOutput) ToTeamMapOutput() TeamMapOutput {
	return o
}

func (o TeamMapOutput) ToTeamMapOutputWithContext(ctx context.Context) TeamMapOutput {
	return o
}

func (o TeamMapOutput) MapIndex(k pulumi.StringInput) TeamOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Team {
		return vs[0].(map[string]Team)[vs[1].(string)]
	}).(TeamOutput)
}

func init() {
	pulumi.RegisterOutputType(TeamOutput{})
	pulumi.RegisterOutputType(TeamPtrOutput{})
	pulumi.RegisterOutputType(TeamArrayOutput{})
	pulumi.RegisterOutputType(TeamMapOutput{})
}
