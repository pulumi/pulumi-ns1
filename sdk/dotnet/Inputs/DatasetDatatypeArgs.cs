// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ns1.Inputs
{

    public sealed class DatasetDatatypeArgs : global::Pulumi.ResourceArgs
    {
        [Input("data", required: true)]
        private InputMap<object>? _data;
        public InputMap<object> Data
        {
            get => _data ?? (_data = new InputMap<object>());
            set => _data = value;
        }

        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public DatasetDatatypeArgs()
        {
        }
        public static new DatasetDatatypeArgs Empty => new DatasetDatatypeArgs();
    }
}
