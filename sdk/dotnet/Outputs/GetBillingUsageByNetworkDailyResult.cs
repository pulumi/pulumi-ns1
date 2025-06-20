// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ns1.Outputs
{

    [OutputType]
    public sealed class GetBillingUsageByNetworkDailyResult
    {
        /// <summary>
        /// Clean queries for this day.
        /// </summary>
        public readonly int CleanQueries;
        /// <summary>
        /// DDoS queries for this day.
        /// </summary>
        public readonly int DdosQueries;
        /// <summary>
        /// NXD responses for this day.
        /// </summary>
        public readonly int NxdResponses;
        /// <summary>
        /// The timestamp for the day.
        /// </summary>
        public readonly int Timestamp;

        [OutputConstructor]
        private GetBillingUsageByNetworkDailyResult(
            int cleanQueries,

            int ddosQueries,

            int nxdResponses,

            int timestamp)
        {
            CleanQueries = cleanQueries;
            DdosQueries = ddosQueries;
            NxdResponses = nxdResponses;
            Timestamp = timestamp;
        }
    }
}
