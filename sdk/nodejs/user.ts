// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides a NS1 User resource. Creating a user sends an invitation email to the
 * user's email address. This can be used to create, modify, and delete users.
 * The credentials used must have the `manageUsers` permission set.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ns1 from "@pulumi/ns1";
 *
 * const example = new ns1.Team("example", {
 *     name: "Example team",
 *     ipWhitelists: [
 *         "1.1.1.1",
 *         "2.2.2.2",
 *     ],
 *     dnsViewZones: false,
 *     accountManageUsers: false,
 * });
 * const exampleUser = new ns1.User("example", {
 *     name: "Example User",
 *     username: "example_user",
 *     email: "user@example.com",
 *     teams: [example.id],
 *     notify: {
 *         billing: false,
 *     },
 * });
 * ```
 *
 * ## Permissions
 *
 * A user will inherit permissions from the teams they are assigned to.
 * If a user is assigned to a team and also has individual permissions set on the user, the individual permissions
 * will be overridden by the inherited team permissions.
 * In a future release, setting permissions on a user that is part of a team will be explicitly disabled.
 *
 * When a user is removed from all teams completely, they will inherit whatever permissions they had previously.
 * If a user is removed from all their teams, it will probably be necessary to run `pulumi up` a second time
 * to update the users permissions from their old team permissions to new user-specific permissions.
 *
 * See [this NS1 Help Center article](https://help.ns1.com/hc/en-us/articles/360024409034-Managing-user-permissions) for an overview of user permission settings.
 *
 * ## NS1 Documentation
 *
 * [User Api Docs](https://ns1.com/api#user)
 *
 * [Managing user permissions](https://help.ns1.com/hc/en-us/articles/360024409034-Managing-user-permissions)
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import ns1:index/user:User <name> <username>`
 * ```
 */
export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserState, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ns1:index/user:User';

    /**
     * Returns true if the given object is an instance of User.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is User {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === User.__pulumiType;
    }

    /**
     * Whether the user can modify account settings.
     */
    public readonly accountManageAccountSettings!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the user can modify account apikeys.
     */
    public readonly accountManageApikeys!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the user can manage ip whitelist.
     */
    public readonly accountManageIpWhitelist!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the user can modify account payment methods.
     */
    public readonly accountManagePaymentMethods!: pulumi.Output<boolean | undefined>;
    /**
     * No longer in use.
     *
     * @deprecated obsolete, should no longer be used
     */
    public readonly accountManagePlan!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the user can modify other teams in the account.
     */
    public readonly accountManageTeams!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the user can modify account users.
     */
    public readonly accountManageUsers!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the user can view activity logs.
     */
    public readonly accountViewActivityLog!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the user can view invoices.
     */
    public readonly accountViewInvoices!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the user can modify data feeds.
     */
    public readonly dataManageDatafeeds!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the user can modify data sources.
     */
    public readonly dataManageDatasources!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the user can publish to data feeds.
     */
    public readonly dataPushToDatafeeds!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the user can manage DHCP.
     * Only relevant for the DDI product.
     */
    public readonly dhcpManageDhcp!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the user can view DHCP.
     * Only relevant for the DDI product.
     */
    public readonly dhcpViewDhcp!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the user can modify the accounts zones.
     */
    public readonly dnsManageZones!: pulumi.Output<boolean | undefined>;
    public readonly dnsRecordsAllows!: pulumi.Output<outputs.UserDnsRecordsAllow[] | undefined>;
    public readonly dnsRecordsDenies!: pulumi.Output<outputs.UserDnsRecordsDeny[] | undefined>;
    /**
     * Whether the user can view the accounts zones.
     */
    public readonly dnsViewZones!: pulumi.Output<boolean | undefined>;
    /**
     * If true, enable the `dnsZonesAllow` list, otherwise enable the `dnsZonesDeny` list.
     */
    public readonly dnsZonesAllowByDefault!: pulumi.Output<boolean | undefined>;
    /**
     * List of zones that the user may access.
     */
    public readonly dnsZonesAllows!: pulumi.Output<string[] | undefined>;
    /**
     * List of zones that the user may not access.
     */
    public readonly dnsZonesDenies!: pulumi.Output<string[] | undefined>;
    /**
     * The email address of the user.
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * Set to true to restrict access to only those IP addresses and networks listed in the **ip_whitelist** field.
     */
    public readonly ipWhitelistStrict!: pulumi.Output<boolean | undefined>;
    /**
     * Array of IP addresses/networks to which to grant the user access.
     */
    public readonly ipWhitelists!: pulumi.Output<string[] | undefined>;
    /**
     * Whether the user can manage IPAM.
     * Only relevant for the DDI product.
     */
    public readonly ipamManageIpam!: pulumi.Output<boolean | undefined>;
    public readonly ipamViewIpam!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the user can modify monitoring jobs.
     */
    public readonly monitoringManageJobs!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the user can modify notification lists.
     */
    public readonly monitoringManageLists!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the user can view monitoring jobs.
     */
    public readonly monitoringViewJobs!: pulumi.Output<boolean | undefined>;
    /**
     * The free form name of the user.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Whether or not to notify the user of specified events. Only `billing` is available currently.
     */
    public readonly notify!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Whether the user can manage global active directory.
     * Only relevant for the DDI product.
     */
    public readonly securityManageActiveDirectory!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the user can manage global two factor authentication.
     */
    public readonly securityManageGlobal2fa!: pulumi.Output<boolean | undefined>;
    /**
     * The teams that the user belongs to.
     */
    public readonly teams!: pulumi.Output<string[] | undefined>;
    /**
     * The users login name.
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserArgs | UserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserState | undefined;
            resourceInputs["accountManageAccountSettings"] = state ? state.accountManageAccountSettings : undefined;
            resourceInputs["accountManageApikeys"] = state ? state.accountManageApikeys : undefined;
            resourceInputs["accountManageIpWhitelist"] = state ? state.accountManageIpWhitelist : undefined;
            resourceInputs["accountManagePaymentMethods"] = state ? state.accountManagePaymentMethods : undefined;
            resourceInputs["accountManagePlan"] = state ? state.accountManagePlan : undefined;
            resourceInputs["accountManageTeams"] = state ? state.accountManageTeams : undefined;
            resourceInputs["accountManageUsers"] = state ? state.accountManageUsers : undefined;
            resourceInputs["accountViewActivityLog"] = state ? state.accountViewActivityLog : undefined;
            resourceInputs["accountViewInvoices"] = state ? state.accountViewInvoices : undefined;
            resourceInputs["dataManageDatafeeds"] = state ? state.dataManageDatafeeds : undefined;
            resourceInputs["dataManageDatasources"] = state ? state.dataManageDatasources : undefined;
            resourceInputs["dataPushToDatafeeds"] = state ? state.dataPushToDatafeeds : undefined;
            resourceInputs["dhcpManageDhcp"] = state ? state.dhcpManageDhcp : undefined;
            resourceInputs["dhcpViewDhcp"] = state ? state.dhcpViewDhcp : undefined;
            resourceInputs["dnsManageZones"] = state ? state.dnsManageZones : undefined;
            resourceInputs["dnsRecordsAllows"] = state ? state.dnsRecordsAllows : undefined;
            resourceInputs["dnsRecordsDenies"] = state ? state.dnsRecordsDenies : undefined;
            resourceInputs["dnsViewZones"] = state ? state.dnsViewZones : undefined;
            resourceInputs["dnsZonesAllowByDefault"] = state ? state.dnsZonesAllowByDefault : undefined;
            resourceInputs["dnsZonesAllows"] = state ? state.dnsZonesAllows : undefined;
            resourceInputs["dnsZonesDenies"] = state ? state.dnsZonesDenies : undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["ipWhitelistStrict"] = state ? state.ipWhitelistStrict : undefined;
            resourceInputs["ipWhitelists"] = state ? state.ipWhitelists : undefined;
            resourceInputs["ipamManageIpam"] = state ? state.ipamManageIpam : undefined;
            resourceInputs["ipamViewIpam"] = state ? state.ipamViewIpam : undefined;
            resourceInputs["monitoringManageJobs"] = state ? state.monitoringManageJobs : undefined;
            resourceInputs["monitoringManageLists"] = state ? state.monitoringManageLists : undefined;
            resourceInputs["monitoringViewJobs"] = state ? state.monitoringViewJobs : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["notify"] = state ? state.notify : undefined;
            resourceInputs["securityManageActiveDirectory"] = state ? state.securityManageActiveDirectory : undefined;
            resourceInputs["securityManageGlobal2fa"] = state ? state.securityManageGlobal2fa : undefined;
            resourceInputs["teams"] = state ? state.teams : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            if ((!args || args.email === undefined) && !opts.urn) {
                throw new Error("Missing required property 'email'");
            }
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            resourceInputs["accountManageAccountSettings"] = args ? args.accountManageAccountSettings : undefined;
            resourceInputs["accountManageApikeys"] = args ? args.accountManageApikeys : undefined;
            resourceInputs["accountManageIpWhitelist"] = args ? args.accountManageIpWhitelist : undefined;
            resourceInputs["accountManagePaymentMethods"] = args ? args.accountManagePaymentMethods : undefined;
            resourceInputs["accountManagePlan"] = args ? args.accountManagePlan : undefined;
            resourceInputs["accountManageTeams"] = args ? args.accountManageTeams : undefined;
            resourceInputs["accountManageUsers"] = args ? args.accountManageUsers : undefined;
            resourceInputs["accountViewActivityLog"] = args ? args.accountViewActivityLog : undefined;
            resourceInputs["accountViewInvoices"] = args ? args.accountViewInvoices : undefined;
            resourceInputs["dataManageDatafeeds"] = args ? args.dataManageDatafeeds : undefined;
            resourceInputs["dataManageDatasources"] = args ? args.dataManageDatasources : undefined;
            resourceInputs["dataPushToDatafeeds"] = args ? args.dataPushToDatafeeds : undefined;
            resourceInputs["dhcpManageDhcp"] = args ? args.dhcpManageDhcp : undefined;
            resourceInputs["dhcpViewDhcp"] = args ? args.dhcpViewDhcp : undefined;
            resourceInputs["dnsManageZones"] = args ? args.dnsManageZones : undefined;
            resourceInputs["dnsRecordsAllows"] = args ? args.dnsRecordsAllows : undefined;
            resourceInputs["dnsRecordsDenies"] = args ? args.dnsRecordsDenies : undefined;
            resourceInputs["dnsViewZones"] = args ? args.dnsViewZones : undefined;
            resourceInputs["dnsZonesAllowByDefault"] = args ? args.dnsZonesAllowByDefault : undefined;
            resourceInputs["dnsZonesAllows"] = args ? args.dnsZonesAllows : undefined;
            resourceInputs["dnsZonesDenies"] = args ? args.dnsZonesDenies : undefined;
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["ipWhitelistStrict"] = args ? args.ipWhitelistStrict : undefined;
            resourceInputs["ipWhitelists"] = args ? args.ipWhitelists : undefined;
            resourceInputs["ipamManageIpam"] = args ? args.ipamManageIpam : undefined;
            resourceInputs["ipamViewIpam"] = args ? args.ipamViewIpam : undefined;
            resourceInputs["monitoringManageJobs"] = args ? args.monitoringManageJobs : undefined;
            resourceInputs["monitoringManageLists"] = args ? args.monitoringManageLists : undefined;
            resourceInputs["monitoringViewJobs"] = args ? args.monitoringViewJobs : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["notify"] = args ? args.notify : undefined;
            resourceInputs["securityManageActiveDirectory"] = args ? args.securityManageActiveDirectory : undefined;
            resourceInputs["securityManageGlobal2fa"] = args ? args.securityManageGlobal2fa : undefined;
            resourceInputs["teams"] = args ? args.teams : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(User.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * Whether the user can modify account settings.
     */
    accountManageAccountSettings?: pulumi.Input<boolean>;
    /**
     * Whether the user can modify account apikeys.
     */
    accountManageApikeys?: pulumi.Input<boolean>;
    /**
     * Whether the user can manage ip whitelist.
     */
    accountManageIpWhitelist?: pulumi.Input<boolean>;
    /**
     * Whether the user can modify account payment methods.
     */
    accountManagePaymentMethods?: pulumi.Input<boolean>;
    /**
     * No longer in use.
     *
     * @deprecated obsolete, should no longer be used
     */
    accountManagePlan?: pulumi.Input<boolean>;
    /**
     * Whether the user can modify other teams in the account.
     */
    accountManageTeams?: pulumi.Input<boolean>;
    /**
     * Whether the user can modify account users.
     */
    accountManageUsers?: pulumi.Input<boolean>;
    /**
     * Whether the user can view activity logs.
     */
    accountViewActivityLog?: pulumi.Input<boolean>;
    /**
     * Whether the user can view invoices.
     */
    accountViewInvoices?: pulumi.Input<boolean>;
    /**
     * Whether the user can modify data feeds.
     */
    dataManageDatafeeds?: pulumi.Input<boolean>;
    /**
     * Whether the user can modify data sources.
     */
    dataManageDatasources?: pulumi.Input<boolean>;
    /**
     * Whether the user can publish to data feeds.
     */
    dataPushToDatafeeds?: pulumi.Input<boolean>;
    /**
     * Whether the user can manage DHCP.
     * Only relevant for the DDI product.
     */
    dhcpManageDhcp?: pulumi.Input<boolean>;
    /**
     * Whether the user can view DHCP.
     * Only relevant for the DDI product.
     */
    dhcpViewDhcp?: pulumi.Input<boolean>;
    /**
     * Whether the user can modify the accounts zones.
     */
    dnsManageZones?: pulumi.Input<boolean>;
    dnsRecordsAllows?: pulumi.Input<pulumi.Input<inputs.UserDnsRecordsAllow>[]>;
    dnsRecordsDenies?: pulumi.Input<pulumi.Input<inputs.UserDnsRecordsDeny>[]>;
    /**
     * Whether the user can view the accounts zones.
     */
    dnsViewZones?: pulumi.Input<boolean>;
    /**
     * If true, enable the `dnsZonesAllow` list, otherwise enable the `dnsZonesDeny` list.
     */
    dnsZonesAllowByDefault?: pulumi.Input<boolean>;
    /**
     * List of zones that the user may access.
     */
    dnsZonesAllows?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of zones that the user may not access.
     */
    dnsZonesDenies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The email address of the user.
     */
    email?: pulumi.Input<string>;
    /**
     * Set to true to restrict access to only those IP addresses and networks listed in the **ip_whitelist** field.
     */
    ipWhitelistStrict?: pulumi.Input<boolean>;
    /**
     * Array of IP addresses/networks to which to grant the user access.
     */
    ipWhitelists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether the user can manage IPAM.
     * Only relevant for the DDI product.
     */
    ipamManageIpam?: pulumi.Input<boolean>;
    ipamViewIpam?: pulumi.Input<boolean>;
    /**
     * Whether the user can modify monitoring jobs.
     */
    monitoringManageJobs?: pulumi.Input<boolean>;
    /**
     * Whether the user can modify notification lists.
     */
    monitoringManageLists?: pulumi.Input<boolean>;
    /**
     * Whether the user can view monitoring jobs.
     */
    monitoringViewJobs?: pulumi.Input<boolean>;
    /**
     * The free form name of the user.
     */
    name?: pulumi.Input<string>;
    /**
     * Whether or not to notify the user of specified events. Only `billing` is available currently.
     */
    notify?: pulumi.Input<{[key: string]: any}>;
    /**
     * Whether the user can manage global active directory.
     * Only relevant for the DDI product.
     */
    securityManageActiveDirectory?: pulumi.Input<boolean>;
    /**
     * Whether the user can manage global two factor authentication.
     */
    securityManageGlobal2fa?: pulumi.Input<boolean>;
    /**
     * The teams that the user belongs to.
     */
    teams?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The users login name.
     */
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * Whether the user can modify account settings.
     */
    accountManageAccountSettings?: pulumi.Input<boolean>;
    /**
     * Whether the user can modify account apikeys.
     */
    accountManageApikeys?: pulumi.Input<boolean>;
    /**
     * Whether the user can manage ip whitelist.
     */
    accountManageIpWhitelist?: pulumi.Input<boolean>;
    /**
     * Whether the user can modify account payment methods.
     */
    accountManagePaymentMethods?: pulumi.Input<boolean>;
    /**
     * No longer in use.
     *
     * @deprecated obsolete, should no longer be used
     */
    accountManagePlan?: pulumi.Input<boolean>;
    /**
     * Whether the user can modify other teams in the account.
     */
    accountManageTeams?: pulumi.Input<boolean>;
    /**
     * Whether the user can modify account users.
     */
    accountManageUsers?: pulumi.Input<boolean>;
    /**
     * Whether the user can view activity logs.
     */
    accountViewActivityLog?: pulumi.Input<boolean>;
    /**
     * Whether the user can view invoices.
     */
    accountViewInvoices?: pulumi.Input<boolean>;
    /**
     * Whether the user can modify data feeds.
     */
    dataManageDatafeeds?: pulumi.Input<boolean>;
    /**
     * Whether the user can modify data sources.
     */
    dataManageDatasources?: pulumi.Input<boolean>;
    /**
     * Whether the user can publish to data feeds.
     */
    dataPushToDatafeeds?: pulumi.Input<boolean>;
    /**
     * Whether the user can manage DHCP.
     * Only relevant for the DDI product.
     */
    dhcpManageDhcp?: pulumi.Input<boolean>;
    /**
     * Whether the user can view DHCP.
     * Only relevant for the DDI product.
     */
    dhcpViewDhcp?: pulumi.Input<boolean>;
    /**
     * Whether the user can modify the accounts zones.
     */
    dnsManageZones?: pulumi.Input<boolean>;
    dnsRecordsAllows?: pulumi.Input<pulumi.Input<inputs.UserDnsRecordsAllow>[]>;
    dnsRecordsDenies?: pulumi.Input<pulumi.Input<inputs.UserDnsRecordsDeny>[]>;
    /**
     * Whether the user can view the accounts zones.
     */
    dnsViewZones?: pulumi.Input<boolean>;
    /**
     * If true, enable the `dnsZonesAllow` list, otherwise enable the `dnsZonesDeny` list.
     */
    dnsZonesAllowByDefault?: pulumi.Input<boolean>;
    /**
     * List of zones that the user may access.
     */
    dnsZonesAllows?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of zones that the user may not access.
     */
    dnsZonesDenies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The email address of the user.
     */
    email: pulumi.Input<string>;
    /**
     * Set to true to restrict access to only those IP addresses and networks listed in the **ip_whitelist** field.
     */
    ipWhitelistStrict?: pulumi.Input<boolean>;
    /**
     * Array of IP addresses/networks to which to grant the user access.
     */
    ipWhitelists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether the user can manage IPAM.
     * Only relevant for the DDI product.
     */
    ipamManageIpam?: pulumi.Input<boolean>;
    ipamViewIpam?: pulumi.Input<boolean>;
    /**
     * Whether the user can modify monitoring jobs.
     */
    monitoringManageJobs?: pulumi.Input<boolean>;
    /**
     * Whether the user can modify notification lists.
     */
    monitoringManageLists?: pulumi.Input<boolean>;
    /**
     * Whether the user can view monitoring jobs.
     */
    monitoringViewJobs?: pulumi.Input<boolean>;
    /**
     * The free form name of the user.
     */
    name?: pulumi.Input<string>;
    /**
     * Whether or not to notify the user of specified events. Only `billing` is available currently.
     */
    notify?: pulumi.Input<{[key: string]: any}>;
    /**
     * Whether the user can manage global active directory.
     * Only relevant for the DDI product.
     */
    securityManageActiveDirectory?: pulumi.Input<boolean>;
    /**
     * Whether the user can manage global two factor authentication.
     */
    securityManageGlobal2fa?: pulumi.Input<boolean>;
    /**
     * The teams that the user belongs to.
     */
    teams?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The users login name.
     */
    username: pulumi.Input<string>;
}
