// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ns1
{
    [Ns1ResourceType("ns1:index/aPIKey:APIKey")]
    public partial class APIKey : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether the apikey can modify account settings.
        /// </summary>
        [Output("accountManageAccountSettings")]
        public Output<bool?> AccountManageAccountSettings { get; private set; } = null!;

        /// <summary>
        /// Whether the apikey can modify account apikeys.
        /// </summary>
        [Output("accountManageApikeys")]
        public Output<bool?> AccountManageApikeys { get; private set; } = null!;

        /// <summary>
        /// Whether the apikey can modify account payment methods.
        /// </summary>
        [Output("accountManagePaymentMethods")]
        public Output<bool?> AccountManagePaymentMethods { get; private set; } = null!;

        /// <summary>
        /// Whether the apikey can modify the account plan.
        /// </summary>
        [Output("accountManagePlan")]
        public Output<bool?> AccountManagePlan { get; private set; } = null!;

        /// <summary>
        /// Whether the apikey can modify other teams in the account.
        /// </summary>
        [Output("accountManageTeams")]
        public Output<bool?> AccountManageTeams { get; private set; } = null!;

        /// <summary>
        /// Whether the apikey can modify account users.
        /// </summary>
        [Output("accountManageUsers")]
        public Output<bool?> AccountManageUsers { get; private set; } = null!;

        /// <summary>
        /// Whether the apikey can view activity logs.
        /// </summary>
        [Output("accountViewActivityLog")]
        public Output<bool?> AccountViewActivityLog { get; private set; } = null!;

        /// <summary>
        /// Whether the apikey can view invoices.
        /// </summary>
        [Output("accountViewInvoices")]
        public Output<bool?> AccountViewInvoices { get; private set; } = null!;

        /// <summary>
        /// Whether the apikey can modify data feeds.
        /// </summary>
        [Output("dataManageDatafeeds")]
        public Output<bool?> DataManageDatafeeds { get; private set; } = null!;

        /// <summary>
        /// Whether the apikey can modify data sources.
        /// </summary>
        [Output("dataManageDatasources")]
        public Output<bool?> DataManageDatasources { get; private set; } = null!;

        /// <summary>
        /// Whether the apikey can publish to data feeds.
        /// </summary>
        [Output("dataPushToDatafeeds")]
        public Output<bool?> DataPushToDatafeeds { get; private set; } = null!;

        /// <summary>
        /// Whether the apikey can manage DHCP.
        /// Only relevant for the DDI product.
        /// </summary>
        [Output("dhcpManageDhcp")]
        public Output<bool?> DhcpManageDhcp { get; private set; } = null!;

        /// <summary>
        /// Whether the apikey can view DHCP.
        /// Only relevant for the DDI product.
        /// </summary>
        [Output("dhcpViewDhcp")]
        public Output<bool?> DhcpViewDhcp { get; private set; } = null!;

        /// <summary>
        /// Whether the apikey can modify the accounts zones.
        /// </summary>
        [Output("dnsManageZones")]
        public Output<bool?> DnsManageZones { get; private set; } = null!;

        /// <summary>
        /// Whether the apikey can view the accounts zones.
        /// </summary>
        [Output("dnsViewZones")]
        public Output<bool?> DnsViewZones { get; private set; } = null!;

        /// <summary>
        /// If true, enable the `dns_zones_allow` list, otherwise enable the `dns_zones_deny` list.
        /// </summary>
        [Output("dnsZonesAllowByDefault")]
        public Output<bool?> DnsZonesAllowByDefault { get; private set; } = null!;

        /// <summary>
        /// List of zones that the apikey may access.
        /// </summary>
        [Output("dnsZonesAllows")]
        public Output<ImmutableArray<string>> DnsZonesAllows { get; private set; } = null!;

        /// <summary>
        /// List of zones that the apikey may not access.
        /// </summary>
        [Output("dnsZonesDenies")]
        public Output<ImmutableArray<string>> DnsZonesDenies { get; private set; } = null!;

        /// <summary>
        /// Sets exclusivity on this IP whitelist.
        /// </summary>
        [Output("ipWhitelistStrict")]
        public Output<bool?> IpWhitelistStrict { get; private set; } = null!;

        /// <summary>
        /// The IP addresses to whitelist for this key.
        /// </summary>
        [Output("ipWhitelists")]
        public Output<ImmutableArray<string>> IpWhitelists { get; private set; } = null!;

        /// <summary>
        /// Whether the apikey can manage IPAM.
        /// Only relevant for the DDI product.
        /// </summary>
        [Output("ipamManageIpam")]
        public Output<bool?> IpamManageIpam { get; private set; } = null!;

        /// <summary>
        /// Whether the apikey can view IPAM.
        /// Only relevant for the DDI product.
        /// </summary>
        [Output("ipamViewIpam")]
        public Output<bool?> IpamViewIpam { get; private set; } = null!;

        /// <summary>
        /// The apikeys authentication token.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// Whether the apikey can modify monitoring jobs.
        /// </summary>
        [Output("monitoringManageJobs")]
        public Output<bool?> MonitoringManageJobs { get; private set; } = null!;

        /// <summary>
        /// Whether the apikey can modify notification lists.
        /// </summary>
        [Output("monitoringManageLists")]
        public Output<bool?> MonitoringManageLists { get; private set; } = null!;

        /// <summary>
        /// Whether the apikey can view monitoring jobs.
        /// </summary>
        [Output("monitoringViewJobs")]
        public Output<bool?> MonitoringViewJobs { get; private set; } = null!;

        /// <summary>
        /// The free form name of the apikey.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Whether the apikey can manage global active directory.
        /// Only relevant for the DDI product.
        /// </summary>
        [Output("securityManageActiveDirectory")]
        public Output<bool?> SecurityManageActiveDirectory { get; private set; } = null!;

        /// <summary>
        /// Whether the apikey can manage global two factor authentication.
        /// </summary>
        [Output("securityManageGlobal2fa")]
        public Output<bool?> SecurityManageGlobal2fa { get; private set; } = null!;

        /// <summary>
        /// The teams that the apikey belongs to.
        /// </summary>
        [Output("teams")]
        public Output<ImmutableArray<string>> Teams { get; private set; } = null!;


        /// <summary>
        /// Create a APIKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public APIKey(string name, APIKeyArgs? args = null, CustomResourceOptions? options = null)
            : base("ns1:index/aPIKey:APIKey", name, args ?? new APIKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private APIKey(string name, Input<string> id, APIKeyState? state = null, CustomResourceOptions? options = null)
            : base("ns1:index/aPIKey:APIKey", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing APIKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static APIKey Get(string name, Input<string> id, APIKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new APIKey(name, id, state, options);
        }
    }

    public sealed class APIKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the apikey can modify account settings.
        /// </summary>
        [Input("accountManageAccountSettings")]
        public Input<bool>? AccountManageAccountSettings { get; set; }

        /// <summary>
        /// Whether the apikey can modify account apikeys.
        /// </summary>
        [Input("accountManageApikeys")]
        public Input<bool>? AccountManageApikeys { get; set; }

        /// <summary>
        /// Whether the apikey can modify account payment methods.
        /// </summary>
        [Input("accountManagePaymentMethods")]
        public Input<bool>? AccountManagePaymentMethods { get; set; }

        /// <summary>
        /// Whether the apikey can modify the account plan.
        /// </summary>
        [Input("accountManagePlan")]
        public Input<bool>? AccountManagePlan { get; set; }

        /// <summary>
        /// Whether the apikey can modify other teams in the account.
        /// </summary>
        [Input("accountManageTeams")]
        public Input<bool>? AccountManageTeams { get; set; }

        /// <summary>
        /// Whether the apikey can modify account users.
        /// </summary>
        [Input("accountManageUsers")]
        public Input<bool>? AccountManageUsers { get; set; }

        /// <summary>
        /// Whether the apikey can view activity logs.
        /// </summary>
        [Input("accountViewActivityLog")]
        public Input<bool>? AccountViewActivityLog { get; set; }

        /// <summary>
        /// Whether the apikey can view invoices.
        /// </summary>
        [Input("accountViewInvoices")]
        public Input<bool>? AccountViewInvoices { get; set; }

        /// <summary>
        /// Whether the apikey can modify data feeds.
        /// </summary>
        [Input("dataManageDatafeeds")]
        public Input<bool>? DataManageDatafeeds { get; set; }

        /// <summary>
        /// Whether the apikey can modify data sources.
        /// </summary>
        [Input("dataManageDatasources")]
        public Input<bool>? DataManageDatasources { get; set; }

        /// <summary>
        /// Whether the apikey can publish to data feeds.
        /// </summary>
        [Input("dataPushToDatafeeds")]
        public Input<bool>? DataPushToDatafeeds { get; set; }

        /// <summary>
        /// Whether the apikey can manage DHCP.
        /// Only relevant for the DDI product.
        /// </summary>
        [Input("dhcpManageDhcp")]
        public Input<bool>? DhcpManageDhcp { get; set; }

        /// <summary>
        /// Whether the apikey can view DHCP.
        /// Only relevant for the DDI product.
        /// </summary>
        [Input("dhcpViewDhcp")]
        public Input<bool>? DhcpViewDhcp { get; set; }

        /// <summary>
        /// Whether the apikey can modify the accounts zones.
        /// </summary>
        [Input("dnsManageZones")]
        public Input<bool>? DnsManageZones { get; set; }

        /// <summary>
        /// Whether the apikey can view the accounts zones.
        /// </summary>
        [Input("dnsViewZones")]
        public Input<bool>? DnsViewZones { get; set; }

        /// <summary>
        /// If true, enable the `dns_zones_allow` list, otherwise enable the `dns_zones_deny` list.
        /// </summary>
        [Input("dnsZonesAllowByDefault")]
        public Input<bool>? DnsZonesAllowByDefault { get; set; }

        [Input("dnsZonesAllows")]
        private InputList<string>? _dnsZonesAllows;

        /// <summary>
        /// List of zones that the apikey may access.
        /// </summary>
        public InputList<string> DnsZonesAllows
        {
            get => _dnsZonesAllows ?? (_dnsZonesAllows = new InputList<string>());
            set => _dnsZonesAllows = value;
        }

        [Input("dnsZonesDenies")]
        private InputList<string>? _dnsZonesDenies;

        /// <summary>
        /// List of zones that the apikey may not access.
        /// </summary>
        public InputList<string> DnsZonesDenies
        {
            get => _dnsZonesDenies ?? (_dnsZonesDenies = new InputList<string>());
            set => _dnsZonesDenies = value;
        }

        /// <summary>
        /// Sets exclusivity on this IP whitelist.
        /// </summary>
        [Input("ipWhitelistStrict")]
        public Input<bool>? IpWhitelistStrict { get; set; }

        [Input("ipWhitelists")]
        private InputList<string>? _ipWhitelists;

        /// <summary>
        /// The IP addresses to whitelist for this key.
        /// </summary>
        public InputList<string> IpWhitelists
        {
            get => _ipWhitelists ?? (_ipWhitelists = new InputList<string>());
            set => _ipWhitelists = value;
        }

        /// <summary>
        /// Whether the apikey can manage IPAM.
        /// Only relevant for the DDI product.
        /// </summary>
        [Input("ipamManageIpam")]
        public Input<bool>? IpamManageIpam { get; set; }

        /// <summary>
        /// Whether the apikey can view IPAM.
        /// Only relevant for the DDI product.
        /// </summary>
        [Input("ipamViewIpam")]
        public Input<bool>? IpamViewIpam { get; set; }

        /// <summary>
        /// Whether the apikey can modify monitoring jobs.
        /// </summary>
        [Input("monitoringManageJobs")]
        public Input<bool>? MonitoringManageJobs { get; set; }

        /// <summary>
        /// Whether the apikey can modify notification lists.
        /// </summary>
        [Input("monitoringManageLists")]
        public Input<bool>? MonitoringManageLists { get; set; }

        /// <summary>
        /// Whether the apikey can view monitoring jobs.
        /// </summary>
        [Input("monitoringViewJobs")]
        public Input<bool>? MonitoringViewJobs { get; set; }

        /// <summary>
        /// The free form name of the apikey.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Whether the apikey can manage global active directory.
        /// Only relevant for the DDI product.
        /// </summary>
        [Input("securityManageActiveDirectory")]
        public Input<bool>? SecurityManageActiveDirectory { get; set; }

        /// <summary>
        /// Whether the apikey can manage global two factor authentication.
        /// </summary>
        [Input("securityManageGlobal2fa")]
        public Input<bool>? SecurityManageGlobal2fa { get; set; }

        [Input("teams")]
        private InputList<string>? _teams;

        /// <summary>
        /// The teams that the apikey belongs to.
        /// </summary>
        public InputList<string> Teams
        {
            get => _teams ?? (_teams = new InputList<string>());
            set => _teams = value;
        }

        public APIKeyArgs()
        {
        }
    }

    public sealed class APIKeyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the apikey can modify account settings.
        /// </summary>
        [Input("accountManageAccountSettings")]
        public Input<bool>? AccountManageAccountSettings { get; set; }

        /// <summary>
        /// Whether the apikey can modify account apikeys.
        /// </summary>
        [Input("accountManageApikeys")]
        public Input<bool>? AccountManageApikeys { get; set; }

        /// <summary>
        /// Whether the apikey can modify account payment methods.
        /// </summary>
        [Input("accountManagePaymentMethods")]
        public Input<bool>? AccountManagePaymentMethods { get; set; }

        /// <summary>
        /// Whether the apikey can modify the account plan.
        /// </summary>
        [Input("accountManagePlan")]
        public Input<bool>? AccountManagePlan { get; set; }

        /// <summary>
        /// Whether the apikey can modify other teams in the account.
        /// </summary>
        [Input("accountManageTeams")]
        public Input<bool>? AccountManageTeams { get; set; }

        /// <summary>
        /// Whether the apikey can modify account users.
        /// </summary>
        [Input("accountManageUsers")]
        public Input<bool>? AccountManageUsers { get; set; }

        /// <summary>
        /// Whether the apikey can view activity logs.
        /// </summary>
        [Input("accountViewActivityLog")]
        public Input<bool>? AccountViewActivityLog { get; set; }

        /// <summary>
        /// Whether the apikey can view invoices.
        /// </summary>
        [Input("accountViewInvoices")]
        public Input<bool>? AccountViewInvoices { get; set; }

        /// <summary>
        /// Whether the apikey can modify data feeds.
        /// </summary>
        [Input("dataManageDatafeeds")]
        public Input<bool>? DataManageDatafeeds { get; set; }

        /// <summary>
        /// Whether the apikey can modify data sources.
        /// </summary>
        [Input("dataManageDatasources")]
        public Input<bool>? DataManageDatasources { get; set; }

        /// <summary>
        /// Whether the apikey can publish to data feeds.
        /// </summary>
        [Input("dataPushToDatafeeds")]
        public Input<bool>? DataPushToDatafeeds { get; set; }

        /// <summary>
        /// Whether the apikey can manage DHCP.
        /// Only relevant for the DDI product.
        /// </summary>
        [Input("dhcpManageDhcp")]
        public Input<bool>? DhcpManageDhcp { get; set; }

        /// <summary>
        /// Whether the apikey can view DHCP.
        /// Only relevant for the DDI product.
        /// </summary>
        [Input("dhcpViewDhcp")]
        public Input<bool>? DhcpViewDhcp { get; set; }

        /// <summary>
        /// Whether the apikey can modify the accounts zones.
        /// </summary>
        [Input("dnsManageZones")]
        public Input<bool>? DnsManageZones { get; set; }

        /// <summary>
        /// Whether the apikey can view the accounts zones.
        /// </summary>
        [Input("dnsViewZones")]
        public Input<bool>? DnsViewZones { get; set; }

        /// <summary>
        /// If true, enable the `dns_zones_allow` list, otherwise enable the `dns_zones_deny` list.
        /// </summary>
        [Input("dnsZonesAllowByDefault")]
        public Input<bool>? DnsZonesAllowByDefault { get; set; }

        [Input("dnsZonesAllows")]
        private InputList<string>? _dnsZonesAllows;

        /// <summary>
        /// List of zones that the apikey may access.
        /// </summary>
        public InputList<string> DnsZonesAllows
        {
            get => _dnsZonesAllows ?? (_dnsZonesAllows = new InputList<string>());
            set => _dnsZonesAllows = value;
        }

        [Input("dnsZonesDenies")]
        private InputList<string>? _dnsZonesDenies;

        /// <summary>
        /// List of zones that the apikey may not access.
        /// </summary>
        public InputList<string> DnsZonesDenies
        {
            get => _dnsZonesDenies ?? (_dnsZonesDenies = new InputList<string>());
            set => _dnsZonesDenies = value;
        }

        /// <summary>
        /// Sets exclusivity on this IP whitelist.
        /// </summary>
        [Input("ipWhitelistStrict")]
        public Input<bool>? IpWhitelistStrict { get; set; }

        [Input("ipWhitelists")]
        private InputList<string>? _ipWhitelists;

        /// <summary>
        /// The IP addresses to whitelist for this key.
        /// </summary>
        public InputList<string> IpWhitelists
        {
            get => _ipWhitelists ?? (_ipWhitelists = new InputList<string>());
            set => _ipWhitelists = value;
        }

        /// <summary>
        /// Whether the apikey can manage IPAM.
        /// Only relevant for the DDI product.
        /// </summary>
        [Input("ipamManageIpam")]
        public Input<bool>? IpamManageIpam { get; set; }

        /// <summary>
        /// Whether the apikey can view IPAM.
        /// Only relevant for the DDI product.
        /// </summary>
        [Input("ipamViewIpam")]
        public Input<bool>? IpamViewIpam { get; set; }

        /// <summary>
        /// The apikeys authentication token.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// Whether the apikey can modify monitoring jobs.
        /// </summary>
        [Input("monitoringManageJobs")]
        public Input<bool>? MonitoringManageJobs { get; set; }

        /// <summary>
        /// Whether the apikey can modify notification lists.
        /// </summary>
        [Input("monitoringManageLists")]
        public Input<bool>? MonitoringManageLists { get; set; }

        /// <summary>
        /// Whether the apikey can view monitoring jobs.
        /// </summary>
        [Input("monitoringViewJobs")]
        public Input<bool>? MonitoringViewJobs { get; set; }

        /// <summary>
        /// The free form name of the apikey.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Whether the apikey can manage global active directory.
        /// Only relevant for the DDI product.
        /// </summary>
        [Input("securityManageActiveDirectory")]
        public Input<bool>? SecurityManageActiveDirectory { get; set; }

        /// <summary>
        /// Whether the apikey can manage global two factor authentication.
        /// </summary>
        [Input("securityManageGlobal2fa")]
        public Input<bool>? SecurityManageGlobal2fa { get; set; }

        [Input("teams")]
        private InputList<string>? _teams;

        /// <summary>
        /// The teams that the apikey belongs to.
        /// </summary>
        public InputList<string> Teams
        {
            get => _teams ?? (_teams = new InputList<string>());
            set => _teams = value;
        }

        public APIKeyState()
        {
        }
    }
}
