// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ns1.Inputs
{

    public sealed class DatasetReportGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("createdAt")]
        public Input<int>? CreatedAt { get; set; }

        [Input("end")]
        public Input<int>? End { get; set; }

        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("start")]
        public Input<int>? Start { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        public DatasetReportGetArgs()
        {
        }
        public static new DatasetReportGetArgs Empty => new DatasetReportGetArgs();
    }
}