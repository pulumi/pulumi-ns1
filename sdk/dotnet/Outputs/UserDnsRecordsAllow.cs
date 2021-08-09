// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ns1.Outputs
{

    [OutputType]
    public sealed class UserDnsRecordsAllow
    {
        public readonly string Domain;
        public readonly bool IncludeSubdomains;
        public readonly string Type;
        public readonly string Zone;

        [OutputConstructor]
        private UserDnsRecordsAllow(
            string domain,

            bool includeSubdomains,

            string type,

            string zone)
        {
            Domain = domain;
            IncludeSubdomains = includeSubdomains;
            Type = type;
            Zone = zone;
        }
    }
}
