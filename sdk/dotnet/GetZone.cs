// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Ns1 = Pulumi.Ns1;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Ns1.GetZone.InvokeAsync(new Ns1.GetZoneArgs
        ///         {
        ///             Zone = "terraform.example.io",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetZoneResult> InvokeAsync(GetZoneArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetZoneResult>("ns1:index/getZone:getZone", args ?? new GetZoneArgs(), options.WithDefaults());

        /// <summary>
        /// Provides details about a NS1 Zone. Use this if you would simply like to read
        /// information from NS1 into your configurations. For read/write operations, you
        /// should use a resource.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Ns1 = Pulumi.Ns1;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Ns1.GetZone.InvokeAsync(new Ns1.GetZoneArgs
        ///         {
        ///             Zone = "terraform.example.io",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetZoneResult> Invoke(GetZoneInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetZoneResult>("ns1:index/getZone:getZone", args ?? new GetZoneInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetZoneArgs : Pulumi.InvokeArgs
    {
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

        /// <summary>
        /// The domain name of the zone.
        /// </summary>
        [Input("zone", required: true)]
        public string Zone { get; set; } = null!;

        public GetZoneArgs()
        {
        }
    }

    public sealed class GetZoneInvokeArgs : Pulumi.InvokeArgs
    {
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

        /// <summary>
        /// The domain name of the zone.
        /// </summary>
        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public GetZoneInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetZoneResult
    {
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
        /// <summary>
        /// The SOA TTL.
        /// </summary>
        public readonly int Ttl;
        public readonly string Zone;

        [OutputConstructor]
        private GetZoneResult(
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

            int refresh,

            int retry,

            ImmutableArray<Outputs.GetZoneSecondaryResult> secondaries,

            int ttl,

            string zone)
        {
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
            Refresh = refresh;
            Retry = retry;
            Secondaries = secondaries;
            Ttl = ttl;
            Zone = zone;
        }
    }
}
