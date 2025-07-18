// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ns1.Inputs
{

    public sealed class ZoneSecondaryGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// IPv4 address of the secondary server.
        /// </summary>
        [Input("ip", required: true)]
        public Input<string> Ip { get; set; } = null!;

        [Input("networks")]
        private InputList<int>? _networks;

        /// <summary>
        /// List of network IDs (`int`) for which the zone
        /// should be made available. Default is network 0, the primary NSONE Global
        /// Network. Normally, you should not have to worry about this.
        /// </summary>
        public InputList<int> Networks
        {
            get => _networks ?? (_networks = new InputList<int>());
            set => _networks = value;
        }

        /// <summary>
        /// Whether we send `NOTIFY` messages to the secondary host
        /// when the zone changes. Default `false`.
        /// </summary>
        [Input("notify")]
        public Input<bool>? Notify { get; set; }

        /// <summary>
        /// Port of the the secondary server. Default `53`.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        public ZoneSecondaryGetArgs()
        {
        }
        public static new ZoneSecondaryGetArgs Empty => new ZoneSecondaryGetArgs();
    }
}
