// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ns1.Inputs
{

    public sealed class RecordAnswerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Space delimited string of RDATA fields dependent on the record type.
        /// </summary>
        [Input("answer")]
        public Input<string>? Answer { get; set; }

        [Input("meta")]
        private InputMap<object>? _meta;
        public InputMap<object> Meta
        {
            get => _meta ?? (_meta = new InputMap<object>());
            set => _meta = value;
        }

        /// <summary>
        /// The region (Answer Group really) that this answer
        /// belongs to. This should be one of the names specified in `regions`. Only a
        /// single `region` per answer is currently supported. If you want an answer in
        /// multiple regions, duplicating the answer (including metadata) is the correct
        /// approach.
        /// * ` meta` - (Optional) meta is supported at the `answer` level. Meta
        /// is documented below.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public RecordAnswerArgs()
        {
        }
    }
}
