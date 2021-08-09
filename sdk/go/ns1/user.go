// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ns1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type User struct {
	pulumi.CustomResourceState

	// Whether the user can modify account settings.
	AccountManageAccountSettings pulumi.BoolPtrOutput `pulumi:"accountManageAccountSettings"`
	// Whether the user can modify account apikeys.
	AccountManageApikeys pulumi.BoolPtrOutput `pulumi:"accountManageApikeys"`
	// Whether the user can manage ip whitelist.
	AccountManageIpWhitelist pulumi.BoolPtrOutput `pulumi:"accountManageIpWhitelist"`
	// Whether the user can modify account payment methods.
	AccountManagePaymentMethods pulumi.BoolPtrOutput `pulumi:"accountManagePaymentMethods"`
	// **Deprecated** Whether the user can modify the account plan.
	//
	// Deprecated: obsolete, should no longer be used
	AccountManagePlan pulumi.BoolPtrOutput `pulumi:"accountManagePlan"`
	// Whether the user can modify other teams in the account.
	AccountManageTeams pulumi.BoolPtrOutput `pulumi:"accountManageTeams"`
	// Whether the user can modify account users.
	AccountManageUsers pulumi.BoolPtrOutput `pulumi:"accountManageUsers"`
	// Whether the user can view activity logs.
	AccountViewActivityLog pulumi.BoolPtrOutput `pulumi:"accountViewActivityLog"`
	// Whether the user can view invoices.
	AccountViewInvoices pulumi.BoolPtrOutput `pulumi:"accountViewInvoices"`
	// Whether the user can modify data feeds.
	DataManageDatafeeds pulumi.BoolPtrOutput `pulumi:"dataManageDatafeeds"`
	// Whether the user can modify data sources.
	DataManageDatasources pulumi.BoolPtrOutput `pulumi:"dataManageDatasources"`
	// Whether the user can publish to data feeds.
	DataPushToDatafeeds pulumi.BoolPtrOutput `pulumi:"dataPushToDatafeeds"`
	// Whether the user can manage DHCP.
	// Only relevant for the DDI product.
	DhcpManageDhcp pulumi.BoolPtrOutput `pulumi:"dhcpManageDhcp"`
	// Whether the user can view DHCP.
	// Only relevant for the DDI product.
	DhcpViewDhcp pulumi.BoolPtrOutput `pulumi:"dhcpViewDhcp"`
	// Whether the user can modify the accounts zones.
	DnsManageZones   pulumi.BoolPtrOutput           `pulumi:"dnsManageZones"`
	DnsRecordsAllows UserDnsRecordsAllowArrayOutput `pulumi:"dnsRecordsAllows"`
	DnsRecordsDenies UserDnsRecordsDenyArrayOutput  `pulumi:"dnsRecordsDenies"`
	// Whether the user can view the accounts zones.
	DnsViewZones pulumi.BoolPtrOutput `pulumi:"dnsViewZones"`
	// If true, enable the `dnsZonesAllow` list, otherwise enable the `dnsZonesDeny` list.
	DnsZonesAllowByDefault pulumi.BoolPtrOutput `pulumi:"dnsZonesAllowByDefault"`
	// List of zones that the user may access.
	DnsZonesAllows pulumi.StringArrayOutput `pulumi:"dnsZonesAllows"`
	// List of zones that the user may not access.
	DnsZonesDenies pulumi.StringArrayOutput `pulumi:"dnsZonesDenies"`
	// The email address of the user.
	Email pulumi.StringOutput `pulumi:"email"`
	// Sets exclusivity on this IP whitelist.
	IpWhitelistStrict pulumi.BoolPtrOutput `pulumi:"ipWhitelistStrict"`
	// The IP addresses to whitelist for this key.
	IpWhitelists pulumi.StringArrayOutput `pulumi:"ipWhitelists"`
	// Whether the user can manage IPAM.
	// Only relevant for the DDI product.
	IpamManageIpam pulumi.BoolPtrOutput `pulumi:"ipamManageIpam"`
	IpamViewIpam   pulumi.BoolPtrOutput `pulumi:"ipamViewIpam"`
	// Whether the user can modify monitoring jobs.
	MonitoringManageJobs pulumi.BoolPtrOutput `pulumi:"monitoringManageJobs"`
	// Whether the user can modify notification lists.
	MonitoringManageLists pulumi.BoolPtrOutput `pulumi:"monitoringManageLists"`
	// Whether the user can view monitoring jobs.
	MonitoringViewJobs pulumi.BoolPtrOutput `pulumi:"monitoringViewJobs"`
	// The free form name of the user.
	Name pulumi.StringOutput `pulumi:"name"`
	// Whether or not to notify the user of specified events. Only `billing` is available currently.
	Notify pulumi.MapOutput `pulumi:"notify"`
	// Whether the user can manage global active directory.
	// Only relevant for the DDI product.
	SecurityManageActiveDirectory pulumi.BoolPtrOutput `pulumi:"securityManageActiveDirectory"`
	// Whether the user can manage global two factor authentication.
	SecurityManageGlobal2fa pulumi.BoolPtrOutput `pulumi:"securityManageGlobal2fa"`
	// The teams that the user belongs to.
	Teams pulumi.StringArrayOutput `pulumi:"teams"`
	// The users login name.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	var resource User
	err := ctx.RegisterResource("ns1:index/user:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("ns1:index/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// Whether the user can modify account settings.
	AccountManageAccountSettings *bool `pulumi:"accountManageAccountSettings"`
	// Whether the user can modify account apikeys.
	AccountManageApikeys *bool `pulumi:"accountManageApikeys"`
	// Whether the user can manage ip whitelist.
	AccountManageIpWhitelist *bool `pulumi:"accountManageIpWhitelist"`
	// Whether the user can modify account payment methods.
	AccountManagePaymentMethods *bool `pulumi:"accountManagePaymentMethods"`
	// **Deprecated** Whether the user can modify the account plan.
	//
	// Deprecated: obsolete, should no longer be used
	AccountManagePlan *bool `pulumi:"accountManagePlan"`
	// Whether the user can modify other teams in the account.
	AccountManageTeams *bool `pulumi:"accountManageTeams"`
	// Whether the user can modify account users.
	AccountManageUsers *bool `pulumi:"accountManageUsers"`
	// Whether the user can view activity logs.
	AccountViewActivityLog *bool `pulumi:"accountViewActivityLog"`
	// Whether the user can view invoices.
	AccountViewInvoices *bool `pulumi:"accountViewInvoices"`
	// Whether the user can modify data feeds.
	DataManageDatafeeds *bool `pulumi:"dataManageDatafeeds"`
	// Whether the user can modify data sources.
	DataManageDatasources *bool `pulumi:"dataManageDatasources"`
	// Whether the user can publish to data feeds.
	DataPushToDatafeeds *bool `pulumi:"dataPushToDatafeeds"`
	// Whether the user can manage DHCP.
	// Only relevant for the DDI product.
	DhcpManageDhcp *bool `pulumi:"dhcpManageDhcp"`
	// Whether the user can view DHCP.
	// Only relevant for the DDI product.
	DhcpViewDhcp *bool `pulumi:"dhcpViewDhcp"`
	// Whether the user can modify the accounts zones.
	DnsManageZones   *bool                 `pulumi:"dnsManageZones"`
	DnsRecordsAllows []UserDnsRecordsAllow `pulumi:"dnsRecordsAllows"`
	DnsRecordsDenies []UserDnsRecordsDeny  `pulumi:"dnsRecordsDenies"`
	// Whether the user can view the accounts zones.
	DnsViewZones *bool `pulumi:"dnsViewZones"`
	// If true, enable the `dnsZonesAllow` list, otherwise enable the `dnsZonesDeny` list.
	DnsZonesAllowByDefault *bool `pulumi:"dnsZonesAllowByDefault"`
	// List of zones that the user may access.
	DnsZonesAllows []string `pulumi:"dnsZonesAllows"`
	// List of zones that the user may not access.
	DnsZonesDenies []string `pulumi:"dnsZonesDenies"`
	// The email address of the user.
	Email *string `pulumi:"email"`
	// Sets exclusivity on this IP whitelist.
	IpWhitelistStrict *bool `pulumi:"ipWhitelistStrict"`
	// The IP addresses to whitelist for this key.
	IpWhitelists []string `pulumi:"ipWhitelists"`
	// Whether the user can manage IPAM.
	// Only relevant for the DDI product.
	IpamManageIpam *bool `pulumi:"ipamManageIpam"`
	IpamViewIpam   *bool `pulumi:"ipamViewIpam"`
	// Whether the user can modify monitoring jobs.
	MonitoringManageJobs *bool `pulumi:"monitoringManageJobs"`
	// Whether the user can modify notification lists.
	MonitoringManageLists *bool `pulumi:"monitoringManageLists"`
	// Whether the user can view monitoring jobs.
	MonitoringViewJobs *bool `pulumi:"monitoringViewJobs"`
	// The free form name of the user.
	Name *string `pulumi:"name"`
	// Whether or not to notify the user of specified events. Only `billing` is available currently.
	Notify map[string]interface{} `pulumi:"notify"`
	// Whether the user can manage global active directory.
	// Only relevant for the DDI product.
	SecurityManageActiveDirectory *bool `pulumi:"securityManageActiveDirectory"`
	// Whether the user can manage global two factor authentication.
	SecurityManageGlobal2fa *bool `pulumi:"securityManageGlobal2fa"`
	// The teams that the user belongs to.
	Teams []string `pulumi:"teams"`
	// The users login name.
	Username *string `pulumi:"username"`
}

type UserState struct {
	// Whether the user can modify account settings.
	AccountManageAccountSettings pulumi.BoolPtrInput
	// Whether the user can modify account apikeys.
	AccountManageApikeys pulumi.BoolPtrInput
	// Whether the user can manage ip whitelist.
	AccountManageIpWhitelist pulumi.BoolPtrInput
	// Whether the user can modify account payment methods.
	AccountManagePaymentMethods pulumi.BoolPtrInput
	// **Deprecated** Whether the user can modify the account plan.
	//
	// Deprecated: obsolete, should no longer be used
	AccountManagePlan pulumi.BoolPtrInput
	// Whether the user can modify other teams in the account.
	AccountManageTeams pulumi.BoolPtrInput
	// Whether the user can modify account users.
	AccountManageUsers pulumi.BoolPtrInput
	// Whether the user can view activity logs.
	AccountViewActivityLog pulumi.BoolPtrInput
	// Whether the user can view invoices.
	AccountViewInvoices pulumi.BoolPtrInput
	// Whether the user can modify data feeds.
	DataManageDatafeeds pulumi.BoolPtrInput
	// Whether the user can modify data sources.
	DataManageDatasources pulumi.BoolPtrInput
	// Whether the user can publish to data feeds.
	DataPushToDatafeeds pulumi.BoolPtrInput
	// Whether the user can manage DHCP.
	// Only relevant for the DDI product.
	DhcpManageDhcp pulumi.BoolPtrInput
	// Whether the user can view DHCP.
	// Only relevant for the DDI product.
	DhcpViewDhcp pulumi.BoolPtrInput
	// Whether the user can modify the accounts zones.
	DnsManageZones   pulumi.BoolPtrInput
	DnsRecordsAllows UserDnsRecordsAllowArrayInput
	DnsRecordsDenies UserDnsRecordsDenyArrayInput
	// Whether the user can view the accounts zones.
	DnsViewZones pulumi.BoolPtrInput
	// If true, enable the `dnsZonesAllow` list, otherwise enable the `dnsZonesDeny` list.
	DnsZonesAllowByDefault pulumi.BoolPtrInput
	// List of zones that the user may access.
	DnsZonesAllows pulumi.StringArrayInput
	// List of zones that the user may not access.
	DnsZonesDenies pulumi.StringArrayInput
	// The email address of the user.
	Email pulumi.StringPtrInput
	// Sets exclusivity on this IP whitelist.
	IpWhitelistStrict pulumi.BoolPtrInput
	// The IP addresses to whitelist for this key.
	IpWhitelists pulumi.StringArrayInput
	// Whether the user can manage IPAM.
	// Only relevant for the DDI product.
	IpamManageIpam pulumi.BoolPtrInput
	IpamViewIpam   pulumi.BoolPtrInput
	// Whether the user can modify monitoring jobs.
	MonitoringManageJobs pulumi.BoolPtrInput
	// Whether the user can modify notification lists.
	MonitoringManageLists pulumi.BoolPtrInput
	// Whether the user can view monitoring jobs.
	MonitoringViewJobs pulumi.BoolPtrInput
	// The free form name of the user.
	Name pulumi.StringPtrInput
	// Whether or not to notify the user of specified events. Only `billing` is available currently.
	Notify pulumi.MapInput
	// Whether the user can manage global active directory.
	// Only relevant for the DDI product.
	SecurityManageActiveDirectory pulumi.BoolPtrInput
	// Whether the user can manage global two factor authentication.
	SecurityManageGlobal2fa pulumi.BoolPtrInput
	// The teams that the user belongs to.
	Teams pulumi.StringArrayInput
	// The users login name.
	Username pulumi.StringPtrInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// Whether the user can modify account settings.
	AccountManageAccountSettings *bool `pulumi:"accountManageAccountSettings"`
	// Whether the user can modify account apikeys.
	AccountManageApikeys *bool `pulumi:"accountManageApikeys"`
	// Whether the user can manage ip whitelist.
	AccountManageIpWhitelist *bool `pulumi:"accountManageIpWhitelist"`
	// Whether the user can modify account payment methods.
	AccountManagePaymentMethods *bool `pulumi:"accountManagePaymentMethods"`
	// **Deprecated** Whether the user can modify the account plan.
	//
	// Deprecated: obsolete, should no longer be used
	AccountManagePlan *bool `pulumi:"accountManagePlan"`
	// Whether the user can modify other teams in the account.
	AccountManageTeams *bool `pulumi:"accountManageTeams"`
	// Whether the user can modify account users.
	AccountManageUsers *bool `pulumi:"accountManageUsers"`
	// Whether the user can view activity logs.
	AccountViewActivityLog *bool `pulumi:"accountViewActivityLog"`
	// Whether the user can view invoices.
	AccountViewInvoices *bool `pulumi:"accountViewInvoices"`
	// Whether the user can modify data feeds.
	DataManageDatafeeds *bool `pulumi:"dataManageDatafeeds"`
	// Whether the user can modify data sources.
	DataManageDatasources *bool `pulumi:"dataManageDatasources"`
	// Whether the user can publish to data feeds.
	DataPushToDatafeeds *bool `pulumi:"dataPushToDatafeeds"`
	// Whether the user can manage DHCP.
	// Only relevant for the DDI product.
	DhcpManageDhcp *bool `pulumi:"dhcpManageDhcp"`
	// Whether the user can view DHCP.
	// Only relevant for the DDI product.
	DhcpViewDhcp *bool `pulumi:"dhcpViewDhcp"`
	// Whether the user can modify the accounts zones.
	DnsManageZones   *bool                 `pulumi:"dnsManageZones"`
	DnsRecordsAllows []UserDnsRecordsAllow `pulumi:"dnsRecordsAllows"`
	DnsRecordsDenies []UserDnsRecordsDeny  `pulumi:"dnsRecordsDenies"`
	// Whether the user can view the accounts zones.
	DnsViewZones *bool `pulumi:"dnsViewZones"`
	// If true, enable the `dnsZonesAllow` list, otherwise enable the `dnsZonesDeny` list.
	DnsZonesAllowByDefault *bool `pulumi:"dnsZonesAllowByDefault"`
	// List of zones that the user may access.
	DnsZonesAllows []string `pulumi:"dnsZonesAllows"`
	// List of zones that the user may not access.
	DnsZonesDenies []string `pulumi:"dnsZonesDenies"`
	// The email address of the user.
	Email string `pulumi:"email"`
	// Sets exclusivity on this IP whitelist.
	IpWhitelistStrict *bool `pulumi:"ipWhitelistStrict"`
	// The IP addresses to whitelist for this key.
	IpWhitelists []string `pulumi:"ipWhitelists"`
	// Whether the user can manage IPAM.
	// Only relevant for the DDI product.
	IpamManageIpam *bool `pulumi:"ipamManageIpam"`
	IpamViewIpam   *bool `pulumi:"ipamViewIpam"`
	// Whether the user can modify monitoring jobs.
	MonitoringManageJobs *bool `pulumi:"monitoringManageJobs"`
	// Whether the user can modify notification lists.
	MonitoringManageLists *bool `pulumi:"monitoringManageLists"`
	// Whether the user can view monitoring jobs.
	MonitoringViewJobs *bool `pulumi:"monitoringViewJobs"`
	// The free form name of the user.
	Name *string `pulumi:"name"`
	// Whether or not to notify the user of specified events. Only `billing` is available currently.
	Notify map[string]interface{} `pulumi:"notify"`
	// Whether the user can manage global active directory.
	// Only relevant for the DDI product.
	SecurityManageActiveDirectory *bool `pulumi:"securityManageActiveDirectory"`
	// Whether the user can manage global two factor authentication.
	SecurityManageGlobal2fa *bool `pulumi:"securityManageGlobal2fa"`
	// The teams that the user belongs to.
	Teams []string `pulumi:"teams"`
	// The users login name.
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// Whether the user can modify account settings.
	AccountManageAccountSettings pulumi.BoolPtrInput
	// Whether the user can modify account apikeys.
	AccountManageApikeys pulumi.BoolPtrInput
	// Whether the user can manage ip whitelist.
	AccountManageIpWhitelist pulumi.BoolPtrInput
	// Whether the user can modify account payment methods.
	AccountManagePaymentMethods pulumi.BoolPtrInput
	// **Deprecated** Whether the user can modify the account plan.
	//
	// Deprecated: obsolete, should no longer be used
	AccountManagePlan pulumi.BoolPtrInput
	// Whether the user can modify other teams in the account.
	AccountManageTeams pulumi.BoolPtrInput
	// Whether the user can modify account users.
	AccountManageUsers pulumi.BoolPtrInput
	// Whether the user can view activity logs.
	AccountViewActivityLog pulumi.BoolPtrInput
	// Whether the user can view invoices.
	AccountViewInvoices pulumi.BoolPtrInput
	// Whether the user can modify data feeds.
	DataManageDatafeeds pulumi.BoolPtrInput
	// Whether the user can modify data sources.
	DataManageDatasources pulumi.BoolPtrInput
	// Whether the user can publish to data feeds.
	DataPushToDatafeeds pulumi.BoolPtrInput
	// Whether the user can manage DHCP.
	// Only relevant for the DDI product.
	DhcpManageDhcp pulumi.BoolPtrInput
	// Whether the user can view DHCP.
	// Only relevant for the DDI product.
	DhcpViewDhcp pulumi.BoolPtrInput
	// Whether the user can modify the accounts zones.
	DnsManageZones   pulumi.BoolPtrInput
	DnsRecordsAllows UserDnsRecordsAllowArrayInput
	DnsRecordsDenies UserDnsRecordsDenyArrayInput
	// Whether the user can view the accounts zones.
	DnsViewZones pulumi.BoolPtrInput
	// If true, enable the `dnsZonesAllow` list, otherwise enable the `dnsZonesDeny` list.
	DnsZonesAllowByDefault pulumi.BoolPtrInput
	// List of zones that the user may access.
	DnsZonesAllows pulumi.StringArrayInput
	// List of zones that the user may not access.
	DnsZonesDenies pulumi.StringArrayInput
	// The email address of the user.
	Email pulumi.StringInput
	// Sets exclusivity on this IP whitelist.
	IpWhitelistStrict pulumi.BoolPtrInput
	// The IP addresses to whitelist for this key.
	IpWhitelists pulumi.StringArrayInput
	// Whether the user can manage IPAM.
	// Only relevant for the DDI product.
	IpamManageIpam pulumi.BoolPtrInput
	IpamViewIpam   pulumi.BoolPtrInput
	// Whether the user can modify monitoring jobs.
	MonitoringManageJobs pulumi.BoolPtrInput
	// Whether the user can modify notification lists.
	MonitoringManageLists pulumi.BoolPtrInput
	// Whether the user can view monitoring jobs.
	MonitoringViewJobs pulumi.BoolPtrInput
	// The free form name of the user.
	Name pulumi.StringPtrInput
	// Whether or not to notify the user of specified events. Only `billing` is available currently.
	Notify pulumi.MapInput
	// Whether the user can manage global active directory.
	// Only relevant for the DDI product.
	SecurityManageActiveDirectory pulumi.BoolPtrInput
	// Whether the user can manage global two factor authentication.
	SecurityManageGlobal2fa pulumi.BoolPtrInput
	// The teams that the user belongs to.
	Teams pulumi.StringArrayInput
	// The users login name.
	Username pulumi.StringInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}

type UserInput interface {
	pulumi.Input

	ToUserOutput() UserOutput
	ToUserOutputWithContext(ctx context.Context) UserOutput
}

func (*User) ElementType() reflect.Type {
	return reflect.TypeOf((*User)(nil))
}

func (i *User) ToUserOutput() UserOutput {
	return i.ToUserOutputWithContext(context.Background())
}

func (i *User) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOutput)
}

func (i *User) ToUserPtrOutput() UserPtrOutput {
	return i.ToUserPtrOutputWithContext(context.Background())
}

func (i *User) ToUserPtrOutputWithContext(ctx context.Context) UserPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPtrOutput)
}

type UserPtrInput interface {
	pulumi.Input

	ToUserPtrOutput() UserPtrOutput
	ToUserPtrOutputWithContext(ctx context.Context) UserPtrOutput
}

type userPtrType UserArgs

func (*userPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil))
}

func (i *userPtrType) ToUserPtrOutput() UserPtrOutput {
	return i.ToUserPtrOutputWithContext(context.Background())
}

func (i *userPtrType) ToUserPtrOutputWithContext(ctx context.Context) UserPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPtrOutput)
}

// UserArrayInput is an input type that accepts UserArray and UserArrayOutput values.
// You can construct a concrete instance of `UserArrayInput` via:
//
//          UserArray{ UserArgs{...} }
type UserArrayInput interface {
	pulumi.Input

	ToUserArrayOutput() UserArrayOutput
	ToUserArrayOutputWithContext(context.Context) UserArrayOutput
}

type UserArray []UserInput

func (UserArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*User)(nil))
}

func (i UserArray) ToUserArrayOutput() UserArrayOutput {
	return i.ToUserArrayOutputWithContext(context.Background())
}

func (i UserArray) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserArrayOutput)
}

// UserMapInput is an input type that accepts UserMap and UserMapOutput values.
// You can construct a concrete instance of `UserMapInput` via:
//
//          UserMap{ "key": UserArgs{...} }
type UserMapInput interface {
	pulumi.Input

	ToUserMapOutput() UserMapOutput
	ToUserMapOutputWithContext(context.Context) UserMapOutput
}

type UserMap map[string]UserInput

func (UserMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*User)(nil))
}

func (i UserMap) ToUserMapOutput() UserMapOutput {
	return i.ToUserMapOutputWithContext(context.Background())
}

func (i UserMap) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserMapOutput)
}

type UserOutput struct {
	*pulumi.OutputState
}

func (UserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*User)(nil))
}

func (o UserOutput) ToUserOutput() UserOutput {
	return o
}

func (o UserOutput) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return o
}

func (o UserOutput) ToUserPtrOutput() UserPtrOutput {
	return o.ToUserPtrOutputWithContext(context.Background())
}

func (o UserOutput) ToUserPtrOutputWithContext(ctx context.Context) UserPtrOutput {
	return o.ApplyT(func(v User) *User {
		return &v
	}).(UserPtrOutput)
}

type UserPtrOutput struct {
	*pulumi.OutputState
}

func (UserPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil))
}

func (o UserPtrOutput) ToUserPtrOutput() UserPtrOutput {
	return o
}

func (o UserPtrOutput) ToUserPtrOutputWithContext(ctx context.Context) UserPtrOutput {
	return o
}

type UserArrayOutput struct{ *pulumi.OutputState }

func (UserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]User)(nil))
}

func (o UserArrayOutput) ToUserArrayOutput() UserArrayOutput {
	return o
}

func (o UserArrayOutput) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return o
}

func (o UserArrayOutput) Index(i pulumi.IntInput) UserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) User {
		return vs[0].([]User)[vs[1].(int)]
	}).(UserOutput)
}

type UserMapOutput struct{ *pulumi.OutputState }

func (UserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]User)(nil))
}

func (o UserMapOutput) ToUserMapOutput() UserMapOutput {
	return o
}

func (o UserMapOutput) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return o
}

func (o UserMapOutput) MapIndex(k pulumi.StringInput) UserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) User {
		return vs[0].(map[string]User)[vs[1].(string)]
	}).(UserOutput)
}

func init() {
	pulumi.RegisterOutputType(UserOutput{})
	pulumi.RegisterOutputType(UserPtrOutput{})
	pulumi.RegisterOutputType(UserArrayOutput{})
	pulumi.RegisterOutputType(UserMapOutput{})
}
