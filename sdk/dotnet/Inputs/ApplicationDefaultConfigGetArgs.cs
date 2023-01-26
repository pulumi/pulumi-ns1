// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ns1.Inputs
{

    public sealed class ApplicationDefaultConfigGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether or not to use HTTP in measurements.
        /// </summary>
        [Input("http", required: true)]
        public Input<bool> Http { get; set; } = null!;

        /// <summary>
        /// Indicates whether or not to use HTTPS in measurements.
        /// </summary>
        [Input("https")]
        public Input<bool>? Https { get; set; }

        /// <summary>
        /// Maximum timeout per job
        /// 0, the primary NSONE Global Network. Normally, you should not have to worry about this.
        /// </summary>
        [Input("jobTimeoutMillis")]
        public Input<int>? JobTimeoutMillis { get; set; }

        /// <summary>
        /// Maximum timeout per request.
        /// </summary>
        [Input("requestTimeoutMillis")]
        public Input<int>? RequestTimeoutMillis { get; set; }

        /// <summary>
        /// Indicates whether or not to skip aggregation for this job's measurements
        /// </summary>
        [Input("staticValues")]
        public Input<bool>? StaticValues { get; set; }

        /// <summary>
        /// Whether to use XMLHttpRequest (XHR) when taking measurements.
        /// </summary>
        [Input("useXhr")]
        public Input<bool>? UseXhr { get; set; }

        public ApplicationDefaultConfigGetArgs()
        {
        }
        public static new ApplicationDefaultConfigGetArgs Empty => new ApplicationDefaultConfigGetArgs();
    }
}
