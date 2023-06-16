// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ns1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a NS1 User resource. Creating a user sends an invitation email to the
// user's email address. This can be used to create, modify, and delete users.
// The credentials used must have the `manageUsers` permission set.
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
//			exampleTeam, err := ns1.NewTeam(ctx, "exampleTeam", &ns1.TeamArgs{
//				IpWhitelists: ns1.TeamIpWhitelistArray{
//					"1.1.1.1",
//					"2.2.2.2",
//				},
//				DnsViewZones:       pulumi.Bool(false),
//				AccountManageUsers: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ns1.NewUser(ctx, "exampleUser", &ns1.UserArgs{
//				Username: pulumi.String("example_user"),
//				Email:    pulumi.String("user@example.com"),
//				Teams: pulumi.StringArray{
//					exampleTeam.ID(),
//				},
//				Notify: pulumi.AnyMap{
//					"billing": pulumi.Any(false),
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
// ## Permissions
//
// A user will inherit permissions from the teams they are assigned to.
// If a user is assigned to a team and also has individual permissions set on the user, the individual permissions
// will be overridden by the inherited team permissions.
// In a future release, setting permissions on a user that is part of a team will be explicitly disabled.
//
// When a user is removed from all teams completely, they will inherit whatever permissions they had previously.
// If a user is removed from all their teams, it will probably be necessary to run `pulumi up` a second time
// to update the users permissions from their old team permissions to new user-specific permissions.
//
// See [this NS1 Help Center article](https://help.ns1.com/hc/en-us/articles/360024409034-Managing-user-permissions) for an overview of user permission settings.
//
// ## NS1 Documentation
//
// [User Api Docs](https://ns1.com/api#user)
//
// [Managing user permissions](https://help.ns1.com/hc/en-us/articles/360024409034-Managing-user-permissions)
//
// ## Import
//
// ```sh
//
//	$ pulumi import ns1:index/user:User <name> <username>`
//
// ```
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
	// No longer in use.
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
	// Set to true to restrict access to only those IP addresses and networks listed in the **ip_whitelist** field.
	IpWhitelistStrict pulumi.BoolPtrOutput `pulumi:"ipWhitelistStrict"`
	// Array of IP addresses/networks to which to grant the user access.
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
	// No longer in use.
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
	// Set to true to restrict access to only those IP addresses and networks listed in the **ip_whitelist** field.
	IpWhitelistStrict *bool `pulumi:"ipWhitelistStrict"`
	// Array of IP addresses/networks to which to grant the user access.
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
	// No longer in use.
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
	// Set to true to restrict access to only those IP addresses and networks listed in the **ip_whitelist** field.
	IpWhitelistStrict pulumi.BoolPtrInput
	// Array of IP addresses/networks to which to grant the user access.
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
	// No longer in use.
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
	// Set to true to restrict access to only those IP addresses and networks listed in the **ip_whitelist** field.
	IpWhitelistStrict *bool `pulumi:"ipWhitelistStrict"`
	// Array of IP addresses/networks to which to grant the user access.
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
	// No longer in use.
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
	// Set to true to restrict access to only those IP addresses and networks listed in the **ip_whitelist** field.
	IpWhitelistStrict pulumi.BoolPtrInput
	// Array of IP addresses/networks to which to grant the user access.
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
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (i *User) ToUserOutput() UserOutput {
	return i.ToUserOutputWithContext(context.Background())
}

func (i *User) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOutput)
}

// UserArrayInput is an input type that accepts UserArray and UserArrayOutput values.
// You can construct a concrete instance of `UserArrayInput` via:
//
//	UserArray{ UserArgs{...} }
type UserArrayInput interface {
	pulumi.Input

	ToUserArrayOutput() UserArrayOutput
	ToUserArrayOutputWithContext(context.Context) UserArrayOutput
}

type UserArray []UserInput

func (UserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
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
//	UserMap{ "key": UserArgs{...} }
type UserMapInput interface {
	pulumi.Input

	ToUserMapOutput() UserMapOutput
	ToUserMapOutputWithContext(context.Context) UserMapOutput
}

type UserMap map[string]UserInput

func (UserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (i UserMap) ToUserMapOutput() UserMapOutput {
	return i.ToUserMapOutputWithContext(context.Background())
}

func (i UserMap) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserMapOutput)
}

type UserOutput struct{ *pulumi.OutputState }

func (UserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (o UserOutput) ToUserOutput() UserOutput {
	return o
}

func (o UserOutput) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return o
}

// Whether the user can modify account settings.
func (o UserOutput) AccountManageAccountSettings() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.AccountManageAccountSettings }).(pulumi.BoolPtrOutput)
}

// Whether the user can modify account apikeys.
func (o UserOutput) AccountManageApikeys() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.AccountManageApikeys }).(pulumi.BoolPtrOutput)
}

// Whether the user can manage ip whitelist.
func (o UserOutput) AccountManageIpWhitelist() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.AccountManageIpWhitelist }).(pulumi.BoolPtrOutput)
}

// Whether the user can modify account payment methods.
func (o UserOutput) AccountManagePaymentMethods() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.AccountManagePaymentMethods }).(pulumi.BoolPtrOutput)
}

// No longer in use.
//
// Deprecated: obsolete, should no longer be used
func (o UserOutput) AccountManagePlan() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.AccountManagePlan }).(pulumi.BoolPtrOutput)
}

// Whether the user can modify other teams in the account.
func (o UserOutput) AccountManageTeams() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.AccountManageTeams }).(pulumi.BoolPtrOutput)
}

// Whether the user can modify account users.
func (o UserOutput) AccountManageUsers() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.AccountManageUsers }).(pulumi.BoolPtrOutput)
}

// Whether the user can view activity logs.
func (o UserOutput) AccountViewActivityLog() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.AccountViewActivityLog }).(pulumi.BoolPtrOutput)
}

// Whether the user can view invoices.
func (o UserOutput) AccountViewInvoices() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.AccountViewInvoices }).(pulumi.BoolPtrOutput)
}

// Whether the user can modify data feeds.
func (o UserOutput) DataManageDatafeeds() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.DataManageDatafeeds }).(pulumi.BoolPtrOutput)
}

// Whether the user can modify data sources.
func (o UserOutput) DataManageDatasources() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.DataManageDatasources }).(pulumi.BoolPtrOutput)
}

// Whether the user can publish to data feeds.
func (o UserOutput) DataPushToDatafeeds() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.DataPushToDatafeeds }).(pulumi.BoolPtrOutput)
}

// Whether the user can manage DHCP.
// Only relevant for the DDI product.
func (o UserOutput) DhcpManageDhcp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.DhcpManageDhcp }).(pulumi.BoolPtrOutput)
}

// Whether the user can view DHCP.
// Only relevant for the DDI product.
func (o UserOutput) DhcpViewDhcp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.DhcpViewDhcp }).(pulumi.BoolPtrOutput)
}

// Whether the user can modify the accounts zones.
func (o UserOutput) DnsManageZones() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.DnsManageZones }).(pulumi.BoolPtrOutput)
}

func (o UserOutput) DnsRecordsAllows() UserDnsRecordsAllowArrayOutput {
	return o.ApplyT(func(v *User) UserDnsRecordsAllowArrayOutput { return v.DnsRecordsAllows }).(UserDnsRecordsAllowArrayOutput)
}

func (o UserOutput) DnsRecordsDenies() UserDnsRecordsDenyArrayOutput {
	return o.ApplyT(func(v *User) UserDnsRecordsDenyArrayOutput { return v.DnsRecordsDenies }).(UserDnsRecordsDenyArrayOutput)
}

// Whether the user can view the accounts zones.
func (o UserOutput) DnsViewZones() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.DnsViewZones }).(pulumi.BoolPtrOutput)
}

// If true, enable the `dnsZonesAllow` list, otherwise enable the `dnsZonesDeny` list.
func (o UserOutput) DnsZonesAllowByDefault() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.DnsZonesAllowByDefault }).(pulumi.BoolPtrOutput)
}

// List of zones that the user may access.
func (o UserOutput) DnsZonesAllows() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *User) pulumi.StringArrayOutput { return v.DnsZonesAllows }).(pulumi.StringArrayOutput)
}

// List of zones that the user may not access.
func (o UserOutput) DnsZonesDenies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *User) pulumi.StringArrayOutput { return v.DnsZonesDenies }).(pulumi.StringArrayOutput)
}

// The email address of the user.
func (o UserOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// Set to true to restrict access to only those IP addresses and networks listed in the **ip_whitelist** field.
func (o UserOutput) IpWhitelistStrict() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.IpWhitelistStrict }).(pulumi.BoolPtrOutput)
}

// Array of IP addresses/networks to which to grant the user access.
func (o UserOutput) IpWhitelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *User) pulumi.StringArrayOutput { return v.IpWhitelists }).(pulumi.StringArrayOutput)
}

// Whether the user can manage IPAM.
// Only relevant for the DDI product.
func (o UserOutput) IpamManageIpam() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.IpamManageIpam }).(pulumi.BoolPtrOutput)
}

func (o UserOutput) IpamViewIpam() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.IpamViewIpam }).(pulumi.BoolPtrOutput)
}

// Whether the user can modify monitoring jobs.
func (o UserOutput) MonitoringManageJobs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.MonitoringManageJobs }).(pulumi.BoolPtrOutput)
}

// Whether the user can modify notification lists.
func (o UserOutput) MonitoringManageLists() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.MonitoringManageLists }).(pulumi.BoolPtrOutput)
}

// Whether the user can view monitoring jobs.
func (o UserOutput) MonitoringViewJobs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.MonitoringViewJobs }).(pulumi.BoolPtrOutput)
}

// The free form name of the user.
func (o UserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Whether or not to notify the user of specified events. Only `billing` is available currently.
func (o UserOutput) Notify() pulumi.MapOutput {
	return o.ApplyT(func(v *User) pulumi.MapOutput { return v.Notify }).(pulumi.MapOutput)
}

// Whether the user can manage global active directory.
// Only relevant for the DDI product.
func (o UserOutput) SecurityManageActiveDirectory() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.SecurityManageActiveDirectory }).(pulumi.BoolPtrOutput)
}

// Whether the user can manage global two factor authentication.
func (o UserOutput) SecurityManageGlobal2fa() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *User) pulumi.BoolPtrOutput { return v.SecurityManageGlobal2fa }).(pulumi.BoolPtrOutput)
}

// The teams that the user belongs to.
func (o UserOutput) Teams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *User) pulumi.StringArrayOutput { return v.Teams }).(pulumi.StringArrayOutput)
}

// The users login name.
func (o UserOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type UserArrayOutput struct{ *pulumi.OutputState }

func (UserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (o UserArrayOutput) ToUserArrayOutput() UserArrayOutput {
	return o
}

func (o UserArrayOutput) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return o
}

func (o UserArrayOutput) Index(i pulumi.IntInput) UserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *User {
		return vs[0].([]*User)[vs[1].(int)]
	}).(UserOutput)
}

type UserMapOutput struct{ *pulumi.OutputState }

func (UserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (o UserMapOutput) ToUserMapOutput() UserMapOutput {
	return o
}

func (o UserMapOutput) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return o
}

func (o UserMapOutput) MapIndex(k pulumi.StringInput) UserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *User {
		return vs[0].(map[string]*User)[vs[1].(string)]
	}).(UserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserInput)(nil)).Elem(), &User{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserArrayInput)(nil)).Elem(), UserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserMapInput)(nil)).Elem(), UserMap{})
	pulumi.RegisterOutputType(UserOutput{})
	pulumi.RegisterOutputType(UserArrayOutput{})
	pulumi.RegisterOutputType(UserMapOutput{})
}
