// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ns1
{
    /// <summary>
    /// Provides a NS1 User resource. Creating a user sends an invitation email to the
    /// user's email address. This can be used to create, modify, and delete users.
    /// The credentials used must have the `manage_users` permission set.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Ns1 = Pulumi.Ns1;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleTeam = new Ns1.Team("exampleTeam", new()
    ///     {
    ///         IpWhitelists = new[]
    ///         {
    ///             "1.1.1.1",
    ///             "2.2.2.2",
    ///         },
    ///         DnsViewZones = false,
    ///         AccountManageUsers = false,
    ///     });
    /// 
    ///     var exampleUser = new Ns1.User("exampleUser", new()
    ///     {
    ///         Username = "example_user",
    ///         Email = "user@example.com",
    ///         Teams = new[]
    ///         {
    ///             exampleTeam.Id,
    ///         },
    ///         Notify = 
    ///         {
    ///             { "billing", false },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Permissions
    /// 
    /// A user will inherit permissions from the teams they are assigned to.
    /// If a user is assigned to a team and also has individual permissions set on the user, the individual permissions
    /// will be overridden by the inherited team permissions.
    /// In a future release, setting permissions on a user that is part of a team will be explicitly disabled.
    /// 
    /// When a user is removed from all teams completely, they will inherit whatever permissions they had previously.
    /// If a user is removed from all their teams, it will probably be necessary to run `pulumi up` a second time
    /// to update the users permissions from their old team permissions to new user-specific permissions.
    /// 
    /// See [this NS1 Help Center article](https://help.ns1.com/hc/en-us/articles/360024409034-Managing-user-permissions) for an overview of user permission settings.
    /// 
    /// ## NS1 Documentation
    /// 
    /// [User Api Docs](https://ns1.com/api#user)
    /// 
    /// [Managing user permissions](https://help.ns1.com/hc/en-us/articles/360024409034-Managing-user-permissions)
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import ns1:index/user:User &lt;name&gt; &lt;username&gt;`
    /// ```
    /// </summary>
    [Ns1ResourceType("ns1:index/user:User")]
    public partial class User : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether the user can modify account settings.
        /// </summary>
        [Output("accountManageAccountSettings")]
        public Output<bool?> AccountManageAccountSettings { get; private set; } = null!;

        /// <summary>
        /// Whether the user can modify account apikeys.
        /// </summary>
        [Output("accountManageApikeys")]
        public Output<bool?> AccountManageApikeys { get; private set; } = null!;

        /// <summary>
        /// Whether the user can manage ip whitelist.
        /// </summary>
        [Output("accountManageIpWhitelist")]
        public Output<bool?> AccountManageIpWhitelist { get; private set; } = null!;

        /// <summary>
        /// Whether the user can modify account payment methods.
        /// </summary>
        [Output("accountManagePaymentMethods")]
        public Output<bool?> AccountManagePaymentMethods { get; private set; } = null!;

        /// <summary>
        /// No longer in use.
        /// </summary>
        [Output("accountManagePlan")]
        public Output<bool?> AccountManagePlan { get; private set; } = null!;

        /// <summary>
        /// Whether the user can modify other teams in the account.
        /// </summary>
        [Output("accountManageTeams")]
        public Output<bool?> AccountManageTeams { get; private set; } = null!;

        /// <summary>
        /// Whether the user can modify account users.
        /// </summary>
        [Output("accountManageUsers")]
        public Output<bool?> AccountManageUsers { get; private set; } = null!;

        /// <summary>
        /// Whether the user can view activity logs.
        /// </summary>
        [Output("accountViewActivityLog")]
        public Output<bool?> AccountViewActivityLog { get; private set; } = null!;

        /// <summary>
        /// Whether the user can view invoices.
        /// </summary>
        [Output("accountViewInvoices")]
        public Output<bool?> AccountViewInvoices { get; private set; } = null!;

        /// <summary>
        /// Whether the user can modify data feeds.
        /// </summary>
        [Output("dataManageDatafeeds")]
        public Output<bool?> DataManageDatafeeds { get; private set; } = null!;

        /// <summary>
        /// Whether the user can modify data sources.
        /// </summary>
        [Output("dataManageDatasources")]
        public Output<bool?> DataManageDatasources { get; private set; } = null!;

        /// <summary>
        /// Whether the user can publish to data feeds.
        /// </summary>
        [Output("dataPushToDatafeeds")]
        public Output<bool?> DataPushToDatafeeds { get; private set; } = null!;

        /// <summary>
        /// Whether the user can manage DHCP.
        /// Only relevant for the DDI product.
        /// </summary>
        [Output("dhcpManageDhcp")]
        public Output<bool?> DhcpManageDhcp { get; private set; } = null!;

        /// <summary>
        /// Whether the user can view DHCP.
        /// Only relevant for the DDI product.
        /// </summary>
        [Output("dhcpViewDhcp")]
        public Output<bool?> DhcpViewDhcp { get; private set; } = null!;

        /// <summary>
        /// Whether the user can modify the accounts zones.
        /// </summary>
        [Output("dnsManageZones")]
        public Output<bool?> DnsManageZones { get; private set; } = null!;

        [Output("dnsRecordsAllows")]
        public Output<ImmutableArray<Outputs.UserDnsRecordsAllow>> DnsRecordsAllows { get; private set; } = null!;

        [Output("dnsRecordsDenies")]
        public Output<ImmutableArray<Outputs.UserDnsRecordsDeny>> DnsRecordsDenies { get; private set; } = null!;

        /// <summary>
        /// Whether the user can view the accounts zones.
        /// </summary>
        [Output("dnsViewZones")]
        public Output<bool?> DnsViewZones { get; private set; } = null!;

        /// <summary>
        /// If true, enable the `dns_zones_allow` list, otherwise enable the `dns_zones_deny` list.
        /// </summary>
        [Output("dnsZonesAllowByDefault")]
        public Output<bool?> DnsZonesAllowByDefault { get; private set; } = null!;

        /// <summary>
        /// List of zones that the user may access.
        /// </summary>
        [Output("dnsZonesAllows")]
        public Output<ImmutableArray<string>> DnsZonesAllows { get; private set; } = null!;

        /// <summary>
        /// List of zones that the user may not access.
        /// </summary>
        [Output("dnsZonesDenies")]
        public Output<ImmutableArray<string>> DnsZonesDenies { get; private set; } = null!;

        /// <summary>
        /// The email address of the user.
        /// </summary>
        [Output("email")]
        public Output<string> Email { get; private set; } = null!;

        /// <summary>
        /// Set to true to restrict access to only those IP addresses and networks listed in the **ip_whitelist** field.
        /// </summary>
        [Output("ipWhitelistStrict")]
        public Output<bool?> IpWhitelistStrict { get; private set; } = null!;

        /// <summary>
        /// Array of IP addresses/networks to which to grant the user access.
        /// </summary>
        [Output("ipWhitelists")]
        public Output<ImmutableArray<string>> IpWhitelists { get; private set; } = null!;

        /// <summary>
        /// Whether the user can manage IPAM.
        /// Only relevant for the DDI product.
        /// </summary>
        [Output("ipamManageIpam")]
        public Output<bool?> IpamManageIpam { get; private set; } = null!;

        [Output("ipamViewIpam")]
        public Output<bool?> IpamViewIpam { get; private set; } = null!;

        /// <summary>
        /// Whether the user can modify monitoring jobs.
        /// </summary>
        [Output("monitoringManageJobs")]
        public Output<bool?> MonitoringManageJobs { get; private set; } = null!;

        /// <summary>
        /// Whether the user can modify notification lists.
        /// </summary>
        [Output("monitoringManageLists")]
        public Output<bool?> MonitoringManageLists { get; private set; } = null!;

        /// <summary>
        /// Whether the user can view monitoring jobs.
        /// </summary>
        [Output("monitoringViewJobs")]
        public Output<bool?> MonitoringViewJobs { get; private set; } = null!;

        /// <summary>
        /// The free form name of the user.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Whether or not to notify the user of specified events. Only `billing` is available currently.
        /// </summary>
        [Output("notify")]
        public Output<ImmutableDictionary<string, object>?> Notify { get; private set; } = null!;

        /// <summary>
        /// Whether the user can manage global active directory.
        /// Only relevant for the DDI product.
        /// </summary>
        [Output("securityManageActiveDirectory")]
        public Output<bool?> SecurityManageActiveDirectory { get; private set; } = null!;

        /// <summary>
        /// Whether the user can manage global two factor authentication.
        /// </summary>
        [Output("securityManageGlobal2fa")]
        public Output<bool?> SecurityManageGlobal2fa { get; private set; } = null!;

        /// <summary>
        /// The teams that the user belongs to.
        /// </summary>
        [Output("teams")]
        public Output<ImmutableArray<string>> Teams { get; private set; } = null!;

        /// <summary>
        /// The users login name.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a User resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public User(string name, UserArgs args, CustomResourceOptions? options = null)
            : base("ns1:index/user:User", name, args ?? new UserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private User(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
            : base("ns1:index/user:User", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing User resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static User Get(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
        {
            return new User(name, id, state, options);
        }
    }

    public sealed class UserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the user can modify account settings.
        /// </summary>
        [Input("accountManageAccountSettings")]
        public Input<bool>? AccountManageAccountSettings { get; set; }

        /// <summary>
        /// Whether the user can modify account apikeys.
        /// </summary>
        [Input("accountManageApikeys")]
        public Input<bool>? AccountManageApikeys { get; set; }

        /// <summary>
        /// Whether the user can manage ip whitelist.
        /// </summary>
        [Input("accountManageIpWhitelist")]
        public Input<bool>? AccountManageIpWhitelist { get; set; }

        /// <summary>
        /// Whether the user can modify account payment methods.
        /// </summary>
        [Input("accountManagePaymentMethods")]
        public Input<bool>? AccountManagePaymentMethods { get; set; }

        /// <summary>
        /// No longer in use.
        /// </summary>
        [Input("accountManagePlan")]
        public Input<bool>? AccountManagePlan { get; set; }

        /// <summary>
        /// Whether the user can modify other teams in the account.
        /// </summary>
        [Input("accountManageTeams")]
        public Input<bool>? AccountManageTeams { get; set; }

        /// <summary>
        /// Whether the user can modify account users.
        /// </summary>
        [Input("accountManageUsers")]
        public Input<bool>? AccountManageUsers { get; set; }

        /// <summary>
        /// Whether the user can view activity logs.
        /// </summary>
        [Input("accountViewActivityLog")]
        public Input<bool>? AccountViewActivityLog { get; set; }

        /// <summary>
        /// Whether the user can view invoices.
        /// </summary>
        [Input("accountViewInvoices")]
        public Input<bool>? AccountViewInvoices { get; set; }

        /// <summary>
        /// Whether the user can modify data feeds.
        /// </summary>
        [Input("dataManageDatafeeds")]
        public Input<bool>? DataManageDatafeeds { get; set; }

        /// <summary>
        /// Whether the user can modify data sources.
        /// </summary>
        [Input("dataManageDatasources")]
        public Input<bool>? DataManageDatasources { get; set; }

        /// <summary>
        /// Whether the user can publish to data feeds.
        /// </summary>
        [Input("dataPushToDatafeeds")]
        public Input<bool>? DataPushToDatafeeds { get; set; }

        /// <summary>
        /// Whether the user can manage DHCP.
        /// Only relevant for the DDI product.
        /// </summary>
        [Input("dhcpManageDhcp")]
        public Input<bool>? DhcpManageDhcp { get; set; }

        /// <summary>
        /// Whether the user can view DHCP.
        /// Only relevant for the DDI product.
        /// </summary>
        [Input("dhcpViewDhcp")]
        public Input<bool>? DhcpViewDhcp { get; set; }

        /// <summary>
        /// Whether the user can modify the accounts zones.
        /// </summary>
        [Input("dnsManageZones")]
        public Input<bool>? DnsManageZones { get; set; }

        [Input("dnsRecordsAllows")]
        private InputList<Inputs.UserDnsRecordsAllowArgs>? _dnsRecordsAllows;
        public InputList<Inputs.UserDnsRecordsAllowArgs> DnsRecordsAllows
        {
            get => _dnsRecordsAllows ?? (_dnsRecordsAllows = new InputList<Inputs.UserDnsRecordsAllowArgs>());
            set => _dnsRecordsAllows = value;
        }

        [Input("dnsRecordsDenies")]
        private InputList<Inputs.UserDnsRecordsDenyArgs>? _dnsRecordsDenies;
        public InputList<Inputs.UserDnsRecordsDenyArgs> DnsRecordsDenies
        {
            get => _dnsRecordsDenies ?? (_dnsRecordsDenies = new InputList<Inputs.UserDnsRecordsDenyArgs>());
            set => _dnsRecordsDenies = value;
        }

        /// <summary>
        /// Whether the user can view the accounts zones.
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
        /// List of zones that the user may access.
        /// </summary>
        public InputList<string> DnsZonesAllows
        {
            get => _dnsZonesAllows ?? (_dnsZonesAllows = new InputList<string>());
            set => _dnsZonesAllows = value;
        }

        [Input("dnsZonesDenies")]
        private InputList<string>? _dnsZonesDenies;

        /// <summary>
        /// List of zones that the user may not access.
        /// </summary>
        public InputList<string> DnsZonesDenies
        {
            get => _dnsZonesDenies ?? (_dnsZonesDenies = new InputList<string>());
            set => _dnsZonesDenies = value;
        }

        /// <summary>
        /// The email address of the user.
        /// </summary>
        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        /// <summary>
        /// Set to true to restrict access to only those IP addresses and networks listed in the **ip_whitelist** field.
        /// </summary>
        [Input("ipWhitelistStrict")]
        public Input<bool>? IpWhitelistStrict { get; set; }

        [Input("ipWhitelists")]
        private InputList<string>? _ipWhitelists;

        /// <summary>
        /// Array of IP addresses/networks to which to grant the user access.
        /// </summary>
        public InputList<string> IpWhitelists
        {
            get => _ipWhitelists ?? (_ipWhitelists = new InputList<string>());
            set => _ipWhitelists = value;
        }

        /// <summary>
        /// Whether the user can manage IPAM.
        /// Only relevant for the DDI product.
        /// </summary>
        [Input("ipamManageIpam")]
        public Input<bool>? IpamManageIpam { get; set; }

        [Input("ipamViewIpam")]
        public Input<bool>? IpamViewIpam { get; set; }

        /// <summary>
        /// Whether the user can modify monitoring jobs.
        /// </summary>
        [Input("monitoringManageJobs")]
        public Input<bool>? MonitoringManageJobs { get; set; }

        /// <summary>
        /// Whether the user can modify notification lists.
        /// </summary>
        [Input("monitoringManageLists")]
        public Input<bool>? MonitoringManageLists { get; set; }

        /// <summary>
        /// Whether the user can view monitoring jobs.
        /// </summary>
        [Input("monitoringViewJobs")]
        public Input<bool>? MonitoringViewJobs { get; set; }

        /// <summary>
        /// The free form name of the user.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notify")]
        private InputMap<object>? _notify;

        /// <summary>
        /// Whether or not to notify the user of specified events. Only `billing` is available currently.
        /// </summary>
        public InputMap<object> Notify
        {
            get => _notify ?? (_notify = new InputMap<object>());
            set => _notify = value;
        }

        /// <summary>
        /// Whether the user can manage global active directory.
        /// Only relevant for the DDI product.
        /// </summary>
        [Input("securityManageActiveDirectory")]
        public Input<bool>? SecurityManageActiveDirectory { get; set; }

        /// <summary>
        /// Whether the user can manage global two factor authentication.
        /// </summary>
        [Input("securityManageGlobal2fa")]
        public Input<bool>? SecurityManageGlobal2fa { get; set; }

        [Input("teams")]
        private InputList<string>? _teams;

        /// <summary>
        /// The teams that the user belongs to.
        /// </summary>
        public InputList<string> Teams
        {
            get => _teams ?? (_teams = new InputList<string>());
            set => _teams = value;
        }

        /// <summary>
        /// The users login name.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public UserArgs()
        {
        }
        public static new UserArgs Empty => new UserArgs();
    }

    public sealed class UserState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the user can modify account settings.
        /// </summary>
        [Input("accountManageAccountSettings")]
        public Input<bool>? AccountManageAccountSettings { get; set; }

        /// <summary>
        /// Whether the user can modify account apikeys.
        /// </summary>
        [Input("accountManageApikeys")]
        public Input<bool>? AccountManageApikeys { get; set; }

        /// <summary>
        /// Whether the user can manage ip whitelist.
        /// </summary>
        [Input("accountManageIpWhitelist")]
        public Input<bool>? AccountManageIpWhitelist { get; set; }

        /// <summary>
        /// Whether the user can modify account payment methods.
        /// </summary>
        [Input("accountManagePaymentMethods")]
        public Input<bool>? AccountManagePaymentMethods { get; set; }

        /// <summary>
        /// No longer in use.
        /// </summary>
        [Input("accountManagePlan")]
        public Input<bool>? AccountManagePlan { get; set; }

        /// <summary>
        /// Whether the user can modify other teams in the account.
        /// </summary>
        [Input("accountManageTeams")]
        public Input<bool>? AccountManageTeams { get; set; }

        /// <summary>
        /// Whether the user can modify account users.
        /// </summary>
        [Input("accountManageUsers")]
        public Input<bool>? AccountManageUsers { get; set; }

        /// <summary>
        /// Whether the user can view activity logs.
        /// </summary>
        [Input("accountViewActivityLog")]
        public Input<bool>? AccountViewActivityLog { get; set; }

        /// <summary>
        /// Whether the user can view invoices.
        /// </summary>
        [Input("accountViewInvoices")]
        public Input<bool>? AccountViewInvoices { get; set; }

        /// <summary>
        /// Whether the user can modify data feeds.
        /// </summary>
        [Input("dataManageDatafeeds")]
        public Input<bool>? DataManageDatafeeds { get; set; }

        /// <summary>
        /// Whether the user can modify data sources.
        /// </summary>
        [Input("dataManageDatasources")]
        public Input<bool>? DataManageDatasources { get; set; }

        /// <summary>
        /// Whether the user can publish to data feeds.
        /// </summary>
        [Input("dataPushToDatafeeds")]
        public Input<bool>? DataPushToDatafeeds { get; set; }

        /// <summary>
        /// Whether the user can manage DHCP.
        /// Only relevant for the DDI product.
        /// </summary>
        [Input("dhcpManageDhcp")]
        public Input<bool>? DhcpManageDhcp { get; set; }

        /// <summary>
        /// Whether the user can view DHCP.
        /// Only relevant for the DDI product.
        /// </summary>
        [Input("dhcpViewDhcp")]
        public Input<bool>? DhcpViewDhcp { get; set; }

        /// <summary>
        /// Whether the user can modify the accounts zones.
        /// </summary>
        [Input("dnsManageZones")]
        public Input<bool>? DnsManageZones { get; set; }

        [Input("dnsRecordsAllows")]
        private InputList<Inputs.UserDnsRecordsAllowGetArgs>? _dnsRecordsAllows;
        public InputList<Inputs.UserDnsRecordsAllowGetArgs> DnsRecordsAllows
        {
            get => _dnsRecordsAllows ?? (_dnsRecordsAllows = new InputList<Inputs.UserDnsRecordsAllowGetArgs>());
            set => _dnsRecordsAllows = value;
        }

        [Input("dnsRecordsDenies")]
        private InputList<Inputs.UserDnsRecordsDenyGetArgs>? _dnsRecordsDenies;
        public InputList<Inputs.UserDnsRecordsDenyGetArgs> DnsRecordsDenies
        {
            get => _dnsRecordsDenies ?? (_dnsRecordsDenies = new InputList<Inputs.UserDnsRecordsDenyGetArgs>());
            set => _dnsRecordsDenies = value;
        }

        /// <summary>
        /// Whether the user can view the accounts zones.
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
        /// List of zones that the user may access.
        /// </summary>
        public InputList<string> DnsZonesAllows
        {
            get => _dnsZonesAllows ?? (_dnsZonesAllows = new InputList<string>());
            set => _dnsZonesAllows = value;
        }

        [Input("dnsZonesDenies")]
        private InputList<string>? _dnsZonesDenies;

        /// <summary>
        /// List of zones that the user may not access.
        /// </summary>
        public InputList<string> DnsZonesDenies
        {
            get => _dnsZonesDenies ?? (_dnsZonesDenies = new InputList<string>());
            set => _dnsZonesDenies = value;
        }

        /// <summary>
        /// The email address of the user.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// Set to true to restrict access to only those IP addresses and networks listed in the **ip_whitelist** field.
        /// </summary>
        [Input("ipWhitelistStrict")]
        public Input<bool>? IpWhitelistStrict { get; set; }

        [Input("ipWhitelists")]
        private InputList<string>? _ipWhitelists;

        /// <summary>
        /// Array of IP addresses/networks to which to grant the user access.
        /// </summary>
        public InputList<string> IpWhitelists
        {
            get => _ipWhitelists ?? (_ipWhitelists = new InputList<string>());
            set => _ipWhitelists = value;
        }

        /// <summary>
        /// Whether the user can manage IPAM.
        /// Only relevant for the DDI product.
        /// </summary>
        [Input("ipamManageIpam")]
        public Input<bool>? IpamManageIpam { get; set; }

        [Input("ipamViewIpam")]
        public Input<bool>? IpamViewIpam { get; set; }

        /// <summary>
        /// Whether the user can modify monitoring jobs.
        /// </summary>
        [Input("monitoringManageJobs")]
        public Input<bool>? MonitoringManageJobs { get; set; }

        /// <summary>
        /// Whether the user can modify notification lists.
        /// </summary>
        [Input("monitoringManageLists")]
        public Input<bool>? MonitoringManageLists { get; set; }

        /// <summary>
        /// Whether the user can view monitoring jobs.
        /// </summary>
        [Input("monitoringViewJobs")]
        public Input<bool>? MonitoringViewJobs { get; set; }

        /// <summary>
        /// The free form name of the user.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notify")]
        private InputMap<object>? _notify;

        /// <summary>
        /// Whether or not to notify the user of specified events. Only `billing` is available currently.
        /// </summary>
        public InputMap<object> Notify
        {
            get => _notify ?? (_notify = new InputMap<object>());
            set => _notify = value;
        }

        /// <summary>
        /// Whether the user can manage global active directory.
        /// Only relevant for the DDI product.
        /// </summary>
        [Input("securityManageActiveDirectory")]
        public Input<bool>? SecurityManageActiveDirectory { get; set; }

        /// <summary>
        /// Whether the user can manage global two factor authentication.
        /// </summary>
        [Input("securityManageGlobal2fa")]
        public Input<bool>? SecurityManageGlobal2fa { get; set; }

        [Input("teams")]
        private InputList<string>? _teams;

        /// <summary>
        /// The teams that the user belongs to.
        /// </summary>
        public InputList<string> Teams
        {
            get => _teams ?? (_teams = new InputList<string>());
            set => _teams = value;
        }

        /// <summary>
        /// The users login name.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public UserState()
        {
        }
        public static new UserState Empty => new UserState();
    }
}
