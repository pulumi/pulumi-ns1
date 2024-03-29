// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ns1.Inputs
{

    public sealed class RecordRegionGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("meta")]
        private InputMap<object>? _meta;
        public InputMap<object> Meta
        {
            get => _meta ?? (_meta = new InputMap<object>());
            set => _meta = value;
        }

        /// <summary>
        /// Name of the region (or Answer Group).
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public RecordRegionGetArgs()
        {
        }
        public static new RecordRegionGetArgs Empty => new RecordRegionGetArgs();
    }
}
