// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ns1.Inputs
{

    public sealed class RecordRegionArgs : global::Pulumi.ResourceArgs
    {
        [Input("meta")]
        private InputMap<string>? _meta;
        public InputMap<string> Meta
        {
            get => _meta ?? (_meta = new InputMap<string>());
            set => _meta = value;
        }

        /// <summary>
        /// Name of the region (or Answer Group).
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public RecordRegionArgs()
        {
        }
        public static new RecordRegionArgs Empty => new RecordRegionArgs();
    }
}
