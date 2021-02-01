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
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import ns1:index/zone:Zone &lt;name&gt; &lt;zone&gt;`
    /// ```
    /// 
    ///  So for the example above
    /// 
    /// ```sh
    ///  $ pulumi import ns1:index/zone:Zone example terraform.example.io`
    /// ```
    /// </summary>
    [Ns1ResourceType("ns1:index/zone:Zone")]
    public partial class Zone : Pulumi.CustomResource
    {
        /// <summary>
        /// List of additional IPv4 addresses for the primary
        /// zone. Conflicts with `secondaries`.
        /// </summary>
        [Output("additionalPrimaries")]
        public Output<ImmutableArray<string>> AdditionalPrimaries { get; private set; } = null!;

        [Output("autogenerateNsRecord")]
        public Output<bool?> AutogenerateNsRecord { get; private set; } = null!;

        /// <summary>
        /// (Computed) Authoritative Name Servers.
        /// </summary>
        [Output("dnsServers")]
        public Output<string> DnsServers { get; private set; } = null!;

        /// <summary>
        /// Whether or not DNSSEC is enabled for the zone.
        /// Note that DNSSEC must be enabled on the account by support for this to be set
        /// to `true`.
        /// </summary>
        [Output("dnssec")]
        public Output<bool> Dnssec { get; private set; } = null!;

        /// <summary>
        /// The SOA Expiry. Conflicts with `primary` and
        /// `additional_primaries` (default must be accepted).
        /// </summary>
        [Output("expiry")]
        public Output<int> Expiry { get; private set; } = null!;

        /// <summary>
        /// (Computed) The SOA Hostmaster.
        /// </summary>
        [Output("hostmaster")]
        public Output<string> Hostmaster { get; private set; } = null!;

        /// <summary>
        /// The target zone(domain name) to link to.
        /// </summary>
        [Output("link")]
        public Output<string?> Link { get; private set; } = null!;

        /// <summary>
        /// - List of network IDs (`int`) for which the zone
        /// should be made available. Default is network 0, the primary NSONE Global
        /// Network. Normally, you should not have to worry about this.
        /// </summary>
        [Output("networks")]
        public Output<ImmutableArray<int>> Networks { get; private set; } = null!;

        /// <summary>
        /// The SOA NX TTL. Conflicts with `primary` and
        /// `additional_primaries` (default must be accepted).
        /// </summary>
        [Output("nxTtl")]
        public Output<int> NxTtl { get; private set; } = null!;

        /// <summary>
        /// The primary zones' IPv4 address. This makes the zone a
        /// secondary. Conflicts with `secondaries`.
        /// </summary>
        [Output("primary")]
        public Output<string?> Primary { get; private set; } = null!;

        /// <summary>
        /// The SOA Refresh. Conflicts with `primary` and
        /// `additional_primaries` (default must be accepted).
        /// </summary>
        [Output("refresh")]
        public Output<int> Refresh { get; private set; } = null!;

        /// <summary>
        /// The SOA Retry. Conflicts with `primary` and
        /// `additional_primaries` (default must be accepted).
        /// </summary>
        [Output("retry")]
        public Output<int> Retry { get; private set; } = null!;

        /// <summary>
        /// List of secondary servers. This makes the zone a
        /// primary. Conflicts with `primary` and `additional_primaries`.
        /// Secondaries is documented below.
        /// </summary>
        [Output("secondaries")]
        public Output<ImmutableArray<Outputs.ZoneSecondary>> Secondaries { get; private set; } = null!;

        /// <summary>
        /// The SOA TTL.
        /// </summary>
        [Output("ttl")]
        public Output<int> Ttl { get; private set; } = null!;

        /// <summary>
        /// The domain name of the zone.
        /// </summary>
        [Output("zone")]
        public Output<string> ZoneName { get; private set; } = null!;


        /// <summary>
        /// Create a Zone resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Zone(string name, ZoneArgs args, CustomResourceOptions? options = null)
            : base("ns1:index/zone:Zone", name, args ?? new ZoneArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Zone(string name, Input<string> id, ZoneState? state = null, CustomResourceOptions? options = null)
            : base("ns1:index/zone:Zone", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Zone resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Zone Get(string name, Input<string> id, ZoneState? state = null, CustomResourceOptions? options = null)
        {
            return new Zone(name, id, state, options);
        }
    }

    public sealed class ZoneArgs : Pulumi.ResourceArgs
    {
        [Input("additionalPrimaries")]
        private InputList<string>? _additionalPrimaries;

        /// <summary>
        /// List of additional IPv4 addresses for the primary
        /// zone. Conflicts with `secondaries`.
        /// </summary>
        public InputList<string> AdditionalPrimaries
        {
            get => _additionalPrimaries ?? (_additionalPrimaries = new InputList<string>());
            set => _additionalPrimaries = value;
        }

        [Input("autogenerateNsRecord")]
        public Input<bool>? AutogenerateNsRecord { get; set; }

        /// <summary>
        /// Whether or not DNSSEC is enabled for the zone.
        /// Note that DNSSEC must be enabled on the account by support for this to be set
        /// to `true`.
        /// </summary>
        [Input("dnssec")]
        public Input<bool>? Dnssec { get; set; }

        /// <summary>
        /// The SOA Expiry. Conflicts with `primary` and
        /// `additional_primaries` (default must be accepted).
        /// </summary>
        [Input("expiry")]
        public Input<int>? Expiry { get; set; }

        /// <summary>
        /// The target zone(domain name) to link to.
        /// </summary>
        [Input("link")]
        public Input<string>? Link { get; set; }

        [Input("networks")]
        private InputList<int>? _networks;

        /// <summary>
        /// - List of network IDs (`int`) for which the zone
        /// should be made available. Default is network 0, the primary NSONE Global
        /// Network. Normally, you should not have to worry about this.
        /// </summary>
        public InputList<int> Networks
        {
            get => _networks ?? (_networks = new InputList<int>());
            set => _networks = value;
        }

        /// <summary>
        /// The SOA NX TTL. Conflicts with `primary` and
        /// `additional_primaries` (default must be accepted).
        /// </summary>
        [Input("nxTtl")]
        public Input<int>? NxTtl { get; set; }

        /// <summary>
        /// The primary zones' IPv4 address. This makes the zone a
        /// secondary. Conflicts with `secondaries`.
        /// </summary>
        [Input("primary")]
        public Input<string>? Primary { get; set; }

        /// <summary>
        /// The SOA Refresh. Conflicts with `primary` and
        /// `additional_primaries` (default must be accepted).
        /// </summary>
        [Input("refresh")]
        public Input<int>? Refresh { get; set; }

        /// <summary>
        /// The SOA Retry. Conflicts with `primary` and
        /// `additional_primaries` (default must be accepted).
        /// </summary>
        [Input("retry")]
        public Input<int>? Retry { get; set; }

        [Input("secondaries")]
        private InputList<Inputs.ZoneSecondaryArgs>? _secondaries;

        /// <summary>
        /// List of secondary servers. This makes the zone a
        /// primary. Conflicts with `primary` and `additional_primaries`.
        /// Secondaries is documented below.
        /// </summary>
        public InputList<Inputs.ZoneSecondaryArgs> Secondaries
        {
            get => _secondaries ?? (_secondaries = new InputList<Inputs.ZoneSecondaryArgs>());
            set => _secondaries = value;
        }

        /// <summary>
        /// The SOA TTL.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// The domain name of the zone.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> ZoneName { get; set; } = null!;

        public ZoneArgs()
        {
        }
    }

    public sealed class ZoneState : Pulumi.ResourceArgs
    {
        [Input("additionalPrimaries")]
        private InputList<string>? _additionalPrimaries;

        /// <summary>
        /// List of additional IPv4 addresses for the primary
        /// zone. Conflicts with `secondaries`.
        /// </summary>
        public InputList<string> AdditionalPrimaries
        {
            get => _additionalPrimaries ?? (_additionalPrimaries = new InputList<string>());
            set => _additionalPrimaries = value;
        }

        [Input("autogenerateNsRecord")]
        public Input<bool>? AutogenerateNsRecord { get; set; }

        /// <summary>
        /// (Computed) Authoritative Name Servers.
        /// </summary>
        [Input("dnsServers")]
        public Input<string>? DnsServers { get; set; }

        /// <summary>
        /// Whether or not DNSSEC is enabled for the zone.
        /// Note that DNSSEC must be enabled on the account by support for this to be set
        /// to `true`.
        /// </summary>
        [Input("dnssec")]
        public Input<bool>? Dnssec { get; set; }

        /// <summary>
        /// The SOA Expiry. Conflicts with `primary` and
        /// `additional_primaries` (default must be accepted).
        /// </summary>
        [Input("expiry")]
        public Input<int>? Expiry { get; set; }

        /// <summary>
        /// (Computed) The SOA Hostmaster.
        /// </summary>
        [Input("hostmaster")]
        public Input<string>? Hostmaster { get; set; }

        /// <summary>
        /// The target zone(domain name) to link to.
        /// </summary>
        [Input("link")]
        public Input<string>? Link { get; set; }

        [Input("networks")]
        private InputList<int>? _networks;

        /// <summary>
        /// - List of network IDs (`int`) for which the zone
        /// should be made available. Default is network 0, the primary NSONE Global
        /// Network. Normally, you should not have to worry about this.
        /// </summary>
        public InputList<int> Networks
        {
            get => _networks ?? (_networks = new InputList<int>());
            set => _networks = value;
        }

        /// <summary>
        /// The SOA NX TTL. Conflicts with `primary` and
        /// `additional_primaries` (default must be accepted).
        /// </summary>
        [Input("nxTtl")]
        public Input<int>? NxTtl { get; set; }

        /// <summary>
        /// The primary zones' IPv4 address. This makes the zone a
        /// secondary. Conflicts with `secondaries`.
        /// </summary>
        [Input("primary")]
        public Input<string>? Primary { get; set; }

        /// <summary>
        /// The SOA Refresh. Conflicts with `primary` and
        /// `additional_primaries` (default must be accepted).
        /// </summary>
        [Input("refresh")]
        public Input<int>? Refresh { get; set; }

        /// <summary>
        /// The SOA Retry. Conflicts with `primary` and
        /// `additional_primaries` (default must be accepted).
        /// </summary>
        [Input("retry")]
        public Input<int>? Retry { get; set; }

        [Input("secondaries")]
        private InputList<Inputs.ZoneSecondaryGetArgs>? _secondaries;

        /// <summary>
        /// List of secondary servers. This makes the zone a
        /// primary. Conflicts with `primary` and `additional_primaries`.
        /// Secondaries is documented below.
        /// </summary>
        public InputList<Inputs.ZoneSecondaryGetArgs> Secondaries
        {
            get => _secondaries ?? (_secondaries = new InputList<Inputs.ZoneSecondaryGetArgs>());
            set => _secondaries = value;
        }

        /// <summary>
        /// The SOA TTL.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        /// <summary>
        /// The domain name of the zone.
        /// </summary>
        [Input("zone")]
        public Input<string>? ZoneName { get; set; }

        public ZoneState()
        {
        }
    }
}
