// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ns1
{
    public static class GetZone
    {
        /// <summary>
        /// Provides details about a NS1 Zone. Use this if you would simply like to read
        /// information from NS1 into your configurations. For read/write operations, you
        /// should use a resource.
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
        ///     // Get details about a NS1 Zone.
        ///     var example = Ns1.GetZone.Invoke(new()
        ///     {
        ///         Zone = "terraform.example.io",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetZoneResult> InvokeAsync(GetZoneArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetZoneResult>("ns1:index/getZone:getZone", args ?? new GetZoneArgs(), options.WithDefaults());

        /// <summary>
        /// Provides details about a NS1 Zone. Use this if you would simply like to read
        /// information from NS1 into your configurations. For read/write operations, you
        /// should use a resource.
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
        ///     // Get details about a NS1 Zone.
        ///     var example = Ns1.GetZone.Invoke(new()
        ///     {
        ///         Zone = "terraform.example.io",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetZoneResult> Invoke(GetZoneInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetZoneResult>("ns1:index/getZone:getZone", args ?? new GetZoneInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Provides details about a NS1 Zone. Use this if you would simply like to read
        /// information from NS1 into your configurations. For read/write operations, you
        /// should use a resource.
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
        ///     // Get details about a NS1 Zone.
        ///     var example = Ns1.GetZone.Invoke(new()
        ///     {
        ///         Zone = "terraform.example.io",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetZoneResult> Invoke(GetZoneInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetZoneResult>("ns1:index/getZone:getZone", args ?? new GetZoneInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetZoneArgs : global::Pulumi.InvokeArgs
    {
        [Input("additionalPorts")]
        private List<int>? _additionalPorts;
        public List<int> AdditionalPorts
        {
            get => _additionalPorts ?? (_additionalPorts = new List<int>());
            set => _additionalPorts = value;
        }

        [Input("additionalPrimaries")]
        private List<string>? _additionalPrimaries;

        /// <summary>
        /// List of additional IPv4 addresses for the primary
        /// zone.
        /// </summary>
        public List<string> AdditionalPrimaries
        {
            get => _additionalPrimaries ?? (_additionalPrimaries = new List<string>());
            set => _additionalPrimaries = value;
        }

        [Input("primaryNetwork")]
        public int? PrimaryNetwork { get; set; }

        [Input("primaryPort")]
        public int? PrimaryPort { get; set; }

        /// <summary>
        /// The domain name of the zone.
        /// </summary>
        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetZoneArgs()
        {
        }
        public static new GetZoneArgs Empty => new GetZoneArgs();
    }

    public sealed class GetZoneInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("additionalPorts")]
        private InputList<int>? _additionalPorts;
        public InputList<int> AdditionalPorts
        {
            get => _additionalPorts ?? (_additionalPorts = new InputList<int>());
            set => _additionalPorts = value;
        }

        [Input("additionalPrimaries")]
        private InputList<string>? _additionalPrimaries;

        /// <summary>
        /// List of additional IPv4 addresses for the primary
        /// zone.
        /// </summary>
        public InputList<string> AdditionalPrimaries
        {
            get => _additionalPrimaries ?? (_additionalPrimaries = new InputList<string>());
            set => _additionalPrimaries = value;
        }

        [Input("primaryNetwork")]
        public Input<int>? PrimaryNetwork { get; set; }

        [Input("primaryPort")]
        public Input<int>? PrimaryPort { get; set; }

        /// <summary>
        /// The domain name of the zone.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GetZoneInvokeArgs()
        {
        }
        public static new GetZoneInvokeArgs Empty => new GetZoneInvokeArgs();
    }


    [OutputType]
    public sealed class GetZoneResult
    {
        public readonly ImmutableArray<int> AdditionalPorts;
        /// <summary>
        /// List of additional IPv4 addresses for the primary
        /// zone.
        /// </summary>
        public readonly ImmutableArray<string> AdditionalPrimaries;
        /// <summary>
        /// Authoritative Name Servers.
        /// </summary>
        public readonly string DnsServers;
        /// <summary>
        /// Whether or not DNSSEC is enabled for the zone.
        /// </summary>
        public readonly bool Dnssec;
        /// <summary>
        /// The SOA Expiry.
        /// </summary>
        public readonly int Expiry;
        /// <summary>
        /// The SOA Hostmaster.
        /// </summary>
        public readonly string Hostmaster;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The linked target zone.
        /// </summary>
        public readonly string Link;
        /// <summary>
        /// List of network IDs (`int`) for which the zone should be made
        /// available. Default is network 0, the primary NSONE Global Network.
        /// </summary>
        public readonly ImmutableArray<int> Networks;
        /// <summary>
        /// The SOA NX TTL.
        /// </summary>
        public readonly int NxTtl;
        /// <summary>
        /// The primary zones' IPv4 address.
        /// </summary>
        public readonly string Primary;
        public readonly int? PrimaryNetwork;
        public readonly int? PrimaryPort;
        /// <summary>
        /// The SOA Refresh.
        /// </summary>
        public readonly int Refresh;
        /// <summary>
        /// The SOA Retry.
        /// </summary>
        public readonly int Retry;
        /// <summary>
        /// List of secondary servers. Secondaries is
        /// documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetZoneSecondaryResult> Secondaries;
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The SOA TTL.
        /// </summary>
        public readonly int Ttl;
        public readonly string Zone;

        [OutputConstructor]
        private GetZoneResult(
            ImmutableArray<int> additionalPorts,

            ImmutableArray<string> additionalPrimaries,

            string dnsServers,

            bool dnssec,

            int expiry,

            string hostmaster,

            string id,

            string link,

            ImmutableArray<int> networks,

            int nxTtl,

            string primary,

            int? primaryNetwork,

            int? primaryPort,

            int refresh,

            int retry,

            ImmutableArray<Outputs.GetZoneSecondaryResult> secondaries,

            ImmutableDictionary<string, string> tags,

            int ttl,

            string zone)
        {
            AdditionalPorts = additionalPorts;
            AdditionalPrimaries = additionalPrimaries;
            DnsServers = dnsServers;
            Dnssec = dnssec;
            Expiry = expiry;
            Hostmaster = hostmaster;
            Id = id;
            Link = link;
            Networks = networks;
            NxTtl = nxTtl;
            Primary = primary;
            PrimaryNetwork = primaryNetwork;
            PrimaryPort = primaryPort;
            Refresh = refresh;
            Retry = retry;
            Secondaries = secondaries;
            Tags = tags;
            Ttl = ttl;
            Zone = zone;
        }
    }
}
