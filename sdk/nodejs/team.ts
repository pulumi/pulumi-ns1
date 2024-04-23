// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides a NS1 Team resource. This can be used to create, modify, and delete
 * teams. The credentials used must have the `manageTeams` permission set.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ns1 from "@pulumi/ns1";
 *
 * // Create a new NS1 Team
 * const example = new ns1.Team("example", {
 *     name: "Example team",
 *     ipWhitelists: [
 *         {
 *             name: "whitelist-1",
 *             values: [
 *                 "1.1.1.1",
 *                 "2.2.2.2",
 *             ],
 *         },
 *         {
 *             name: "whitelist-2",
 *             values: [
 *                 "3.3.3.3",
 *                 "4.4.4.4",
 *             ],
 *         },
 *     ],
 *     dnsViewZones: false,
 *     accountManageUsers: false,
 * });
 * // Another team
 * const example2 = new ns1.Team("example2", {
 *     name: "another team",
 *     dnsViewZones: true,
 *     dnsZonesAllowByDefault: true,
 *     dnsZonesAllows: ["mytest.zone"],
 *     dnsZonesDenies: ["myother.zone"],
 *     dnsRecordsAllows: [{
 *         domain: "terraform.example.io",
 *         includeSubdomains: false,
 *         zone: "example.io",
 *         type: "A",
 *     }],
 *     dataManageDatasources: true,
 * });
 * ```
 *
 * ## NS1 Documentation
 *
 * [Team Api Docs](https://ns1.com/api#team)
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import ns1:index/team:Team <name> <team_id>`
 * ```
 */
export class Team extends pulumi.CustomResource {
    /**
     * Get an existing Team resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TeamState, opts?: pulumi.CustomResourceOptions): Team {
        return new Team(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ns1:index/team:Team';

    /**
     * Returns true if the given object is an instance of Team.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Team {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Team.__pulumiType;
    }

    /**
     * Whether the team can modify account settings.
     */
    public readonly accountManageAccountSettings!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the team can modify account apikeys.
     */
    public readonly accountManageApikeys!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the team can manage ip whitelist.
     */
    public readonly accountManageIpWhitelist!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the team can modify account payment methods.
     */
    public readonly accountManagePaymentMethods!: pulumi.Output<boolean | undefined>;
    /**
     * No longer in use.
     *
     * @deprecated obsolete, should no longer be used
     */
    public readonly accountManagePlan!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the team can modify other teams in the account.
     */
    public readonly accountManageTeams!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the team can modify account users.
     */
    public readonly accountManageUsers!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the team can view activity logs.
     */
    public readonly accountViewActivityLog!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the team can view invoices.
     */
    public readonly accountViewInvoices!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the team can modify data feeds.
     */
    public readonly dataManageDatafeeds!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the team can modify data sources.
     */
    public readonly dataManageDatasources!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the team can publish to data feeds.
     */
    public readonly dataPushToDatafeeds!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the team can manage DHCP.
     * Only relevant for the DDI product.
     */
    public readonly dhcpManageDhcp!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the team can view DHCP.
     * Only relevant for the DDI product.
     */
    public readonly dhcpViewDhcp!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the team can modify the accounts zones.
     */
    public readonly dnsManageZones!: pulumi.Output<boolean | undefined>;
    /**
     * List of records that the team may access.
     */
    public readonly dnsRecordsAllows!: pulumi.Output<outputs.TeamDnsRecordsAllow[] | undefined>;
    /**
     * List of records that the team may not access.
     */
    public readonly dnsRecordsDenies!: pulumi.Output<outputs.TeamDnsRecordsDeny[] | undefined>;
    /**
     * Whether the team can view the accounts zones.
     */
    public readonly dnsViewZones!: pulumi.Output<boolean | undefined>;
    /**
     * If true, enable the `dnsZonesAllow` list, otherwise enable the `dnsZonesDeny` list.
     */
    public readonly dnsZonesAllowByDefault!: pulumi.Output<boolean | undefined>;
    /**
     * List of zones that the team may access.
     */
    public readonly dnsZonesAllows!: pulumi.Output<string[] | undefined>;
    /**
     * List of zones that the team may not access.
     */
    public readonly dnsZonesDenies!: pulumi.Output<string[] | undefined>;
    /**
     * Array of IP addresses objects to chich to grant the team access. Each object includes a **name** (string), and **values** (array of strings) associated to each "allow" list.
     */
    public readonly ipWhitelists!: pulumi.Output<outputs.TeamIpWhitelist[] | undefined>;
    /**
     * Whether the team can manage IPAM.
     * Only relevant for the DDI product.
     */
    public readonly ipamManageIpam!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the team can view IPAM.
     * Only relevant for the DDI product.
     */
    public readonly ipamViewIpam!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the team can modify monitoring jobs.
     */
    public readonly monitoringManageJobs!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the team can modify notification lists.
     */
    public readonly monitoringManageLists!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the team can view monitoring jobs.
     */
    public readonly monitoringViewJobs!: pulumi.Output<boolean | undefined>;
    /**
     * The free form name of the team.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Whether the team can manage global active directory.
     * Only relevant for the DDI product.
     */
    public readonly securityManageActiveDirectory!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the team can manage global two factor authentication.
     */
    public readonly securityManageGlobal2fa!: pulumi.Output<boolean | undefined>;

    /**
     * Create a Team resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TeamArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TeamArgs | TeamState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TeamState | undefined;
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
            resourceInputs["ipWhitelists"] = state ? state.ipWhitelists : undefined;
            resourceInputs["ipamManageIpam"] = state ? state.ipamManageIpam : undefined;
            resourceInputs["ipamViewIpam"] = state ? state.ipamViewIpam : undefined;
            resourceInputs["monitoringManageJobs"] = state ? state.monitoringManageJobs : undefined;
            resourceInputs["monitoringManageLists"] = state ? state.monitoringManageLists : undefined;
            resourceInputs["monitoringViewJobs"] = state ? state.monitoringViewJobs : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["securityManageActiveDirectory"] = state ? state.securityManageActiveDirectory : undefined;
            resourceInputs["securityManageGlobal2fa"] = state ? state.securityManageGlobal2fa : undefined;
        } else {
            const args = argsOrState as TeamArgs | undefined;
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
            resourceInputs["ipWhitelists"] = args ? args.ipWhitelists : undefined;
            resourceInputs["ipamManageIpam"] = args ? args.ipamManageIpam : undefined;
            resourceInputs["ipamViewIpam"] = args ? args.ipamViewIpam : undefined;
            resourceInputs["monitoringManageJobs"] = args ? args.monitoringManageJobs : undefined;
            resourceInputs["monitoringManageLists"] = args ? args.monitoringManageLists : undefined;
            resourceInputs["monitoringViewJobs"] = args ? args.monitoringViewJobs : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["securityManageActiveDirectory"] = args ? args.securityManageActiveDirectory : undefined;
            resourceInputs["securityManageGlobal2fa"] = args ? args.securityManageGlobal2fa : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Team.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Team resources.
 */
export interface TeamState {
    /**
     * Whether the team can modify account settings.
     */
    accountManageAccountSettings?: pulumi.Input<boolean>;
    /**
     * Whether the team can modify account apikeys.
     */
    accountManageApikeys?: pulumi.Input<boolean>;
    /**
     * Whether the team can manage ip whitelist.
     */
    accountManageIpWhitelist?: pulumi.Input<boolean>;
    /**
     * Whether the team can modify account payment methods.
     */
    accountManagePaymentMethods?: pulumi.Input<boolean>;
    /**
     * No longer in use.
     *
     * @deprecated obsolete, should no longer be used
     */
    accountManagePlan?: pulumi.Input<boolean>;
    /**
     * Whether the team can modify other teams in the account.
     */
    accountManageTeams?: pulumi.Input<boolean>;
    /**
     * Whether the team can modify account users.
     */
    accountManageUsers?: pulumi.Input<boolean>;
    /**
     * Whether the team can view activity logs.
     */
    accountViewActivityLog?: pulumi.Input<boolean>;
    /**
     * Whether the team can view invoices.
     */
    accountViewInvoices?: pulumi.Input<boolean>;
    /**
     * Whether the team can modify data feeds.
     */
    dataManageDatafeeds?: pulumi.Input<boolean>;
    /**
     * Whether the team can modify data sources.
     */
    dataManageDatasources?: pulumi.Input<boolean>;
    /**
     * Whether the team can publish to data feeds.
     */
    dataPushToDatafeeds?: pulumi.Input<boolean>;
    /**
     * Whether the team can manage DHCP.
     * Only relevant for the DDI product.
     */
    dhcpManageDhcp?: pulumi.Input<boolean>;
    /**
     * Whether the team can view DHCP.
     * Only relevant for the DDI product.
     */
    dhcpViewDhcp?: pulumi.Input<boolean>;
    /**
     * Whether the team can modify the accounts zones.
     */
    dnsManageZones?: pulumi.Input<boolean>;
    /**
     * List of records that the team may access.
     */
    dnsRecordsAllows?: pulumi.Input<pulumi.Input<inputs.TeamDnsRecordsAllow>[]>;
    /**
     * List of records that the team may not access.
     */
    dnsRecordsDenies?: pulumi.Input<pulumi.Input<inputs.TeamDnsRecordsDeny>[]>;
    /**
     * Whether the team can view the accounts zones.
     */
    dnsViewZones?: pulumi.Input<boolean>;
    /**
     * If true, enable the `dnsZonesAllow` list, otherwise enable the `dnsZonesDeny` list.
     */
    dnsZonesAllowByDefault?: pulumi.Input<boolean>;
    /**
     * List of zones that the team may access.
     */
    dnsZonesAllows?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of zones that the team may not access.
     */
    dnsZonesDenies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Array of IP addresses objects to chich to grant the team access. Each object includes a **name** (string), and **values** (array of strings) associated to each "allow" list.
     */
    ipWhitelists?: pulumi.Input<pulumi.Input<inputs.TeamIpWhitelist>[]>;
    /**
     * Whether the team can manage IPAM.
     * Only relevant for the DDI product.
     */
    ipamManageIpam?: pulumi.Input<boolean>;
    /**
     * Whether the team can view IPAM.
     * Only relevant for the DDI product.
     */
    ipamViewIpam?: pulumi.Input<boolean>;
    /**
     * Whether the team can modify monitoring jobs.
     */
    monitoringManageJobs?: pulumi.Input<boolean>;
    /**
     * Whether the team can modify notification lists.
     */
    monitoringManageLists?: pulumi.Input<boolean>;
    /**
     * Whether the team can view monitoring jobs.
     */
    monitoringViewJobs?: pulumi.Input<boolean>;
    /**
     * The free form name of the team.
     */
    name?: pulumi.Input<string>;
    /**
     * Whether the team can manage global active directory.
     * Only relevant for the DDI product.
     */
    securityManageActiveDirectory?: pulumi.Input<boolean>;
    /**
     * Whether the team can manage global two factor authentication.
     */
    securityManageGlobal2fa?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Team resource.
 */
export interface TeamArgs {
    /**
     * Whether the team can modify account settings.
     */
    accountManageAccountSettings?: pulumi.Input<boolean>;
    /**
     * Whether the team can modify account apikeys.
     */
    accountManageApikeys?: pulumi.Input<boolean>;
    /**
     * Whether the team can manage ip whitelist.
     */
    accountManageIpWhitelist?: pulumi.Input<boolean>;
    /**
     * Whether the team can modify account payment methods.
     */
    accountManagePaymentMethods?: pulumi.Input<boolean>;
    /**
     * No longer in use.
     *
     * @deprecated obsolete, should no longer be used
     */
    accountManagePlan?: pulumi.Input<boolean>;
    /**
     * Whether the team can modify other teams in the account.
     */
    accountManageTeams?: pulumi.Input<boolean>;
    /**
     * Whether the team can modify account users.
     */
    accountManageUsers?: pulumi.Input<boolean>;
    /**
     * Whether the team can view activity logs.
     */
    accountViewActivityLog?: pulumi.Input<boolean>;
    /**
     * Whether the team can view invoices.
     */
    accountViewInvoices?: pulumi.Input<boolean>;
    /**
     * Whether the team can modify data feeds.
     */
    dataManageDatafeeds?: pulumi.Input<boolean>;
    /**
     * Whether the team can modify data sources.
     */
    dataManageDatasources?: pulumi.Input<boolean>;
    /**
     * Whether the team can publish to data feeds.
     */
    dataPushToDatafeeds?: pulumi.Input<boolean>;
    /**
     * Whether the team can manage DHCP.
     * Only relevant for the DDI product.
     */
    dhcpManageDhcp?: pulumi.Input<boolean>;
    /**
     * Whether the team can view DHCP.
     * Only relevant for the DDI product.
     */
    dhcpViewDhcp?: pulumi.Input<boolean>;
    /**
     * Whether the team can modify the accounts zones.
     */
    dnsManageZones?: pulumi.Input<boolean>;
    /**
     * List of records that the team may access.
     */
    dnsRecordsAllows?: pulumi.Input<pulumi.Input<inputs.TeamDnsRecordsAllow>[]>;
    /**
     * List of records that the team may not access.
     */
    dnsRecordsDenies?: pulumi.Input<pulumi.Input<inputs.TeamDnsRecordsDeny>[]>;
    /**
     * Whether the team can view the accounts zones.
     */
    dnsViewZones?: pulumi.Input<boolean>;
    /**
     * If true, enable the `dnsZonesAllow` list, otherwise enable the `dnsZonesDeny` list.
     */
    dnsZonesAllowByDefault?: pulumi.Input<boolean>;
    /**
     * List of zones that the team may access.
     */
    dnsZonesAllows?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of zones that the team may not access.
     */
    dnsZonesDenies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Array of IP addresses objects to chich to grant the team access. Each object includes a **name** (string), and **values** (array of strings) associated to each "allow" list.
     */
    ipWhitelists?: pulumi.Input<pulumi.Input<inputs.TeamIpWhitelist>[]>;
    /**
     * Whether the team can manage IPAM.
     * Only relevant for the DDI product.
     */
    ipamManageIpam?: pulumi.Input<boolean>;
    /**
     * Whether the team can view IPAM.
     * Only relevant for the DDI product.
     */
    ipamViewIpam?: pulumi.Input<boolean>;
    /**
     * Whether the team can modify monitoring jobs.
     */
    monitoringManageJobs?: pulumi.Input<boolean>;
    /**
     * Whether the team can modify notification lists.
     */
    monitoringManageLists?: pulumi.Input<boolean>;
    /**
     * Whether the team can view monitoring jobs.
     */
    monitoringViewJobs?: pulumi.Input<boolean>;
    /**
     * The free form name of the team.
     */
    name?: pulumi.Input<string>;
    /**
     * Whether the team can manage global active directory.
     * Only relevant for the DDI product.
     */
    securityManageActiveDirectory?: pulumi.Input<boolean>;
    /**
     * Whether the team can manage global two factor authentication.
     */
    securityManageGlobal2fa?: pulumi.Input<boolean>;
}
