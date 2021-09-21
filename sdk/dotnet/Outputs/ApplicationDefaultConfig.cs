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
    public sealed class ApplicationDefaultConfig
    {
        /// <summary>
        /// Indicates whether or not to use HTTP in measurements.
        /// </summary>
        public readonly bool Http;
        /// <summary>
        /// Indicates whether or not to use HTTPS in measurements.
        /// </summary>
        public readonly bool? Https;
        /// <summary>
        /// - Maximum timeout per job
        /// 0, the primary NSONE Global Network. Normally, you should not have to worry about this.
        /// </summary>
        public readonly int? JobTimeoutMillis;
        /// <summary>
        /// Maximum timeout per request.
        /// </summary>
        public readonly int? RequestTimeoutMillis;
        /// <summary>
        /// - Indicates whether or not to skip aggregation for this job's measurements
        /// </summary>
        public readonly bool? StaticValues;
        /// <summary>
        /// - Whether to use XMLHttpRequest (XHR) when taking measurements.
        /// </summary>
        public readonly bool? UseXhr;

        [OutputConstructor]
        private ApplicationDefaultConfig(
            bool http,

            bool? https,

            int? jobTimeoutMillis,

            int? requestTimeoutMillis,

            bool? staticValues,

            bool? useXhr)
        {
            Http = http;
            Https = https;
            JobTimeoutMillis = jobTimeoutMillis;
            RequestTimeoutMillis = requestTimeoutMillis;
            StaticValues = staticValues;
            UseXhr = useXhr;
        }
    }
}
