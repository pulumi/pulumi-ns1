// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ns1.Inputs
{

    public sealed class APIKeyDnsRecordsDenyGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        [Input("includeSubdomains", required: true)]
        public Input<bool> IncludeSubdomains { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("zone", required: true)]
        public Input<string> Zone { get; set; } = null!;

        public APIKeyDnsRecordsDenyGetArgs()
        {
        }
        public static new APIKeyDnsRecordsDenyGetArgs Empty => new APIKeyDnsRecordsDenyGetArgs();
    }
}
