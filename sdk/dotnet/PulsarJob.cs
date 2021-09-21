// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ns1
{
    [Ns1ResourceType("ns1:index/pulsarJob:PulsarJob")]
    public partial class PulsarJob : Pulumi.CustomResource
    {
        [Output("active")]
        public Output<bool?> Active { get; private set; } = null!;

        [Output("appId")]
        public Output<string> AppId { get; private set; } = null!;

        [Output("blendMetricWeights")]
        public Output<Outputs.PulsarJobBlendMetricWeights?> BlendMetricWeights { get; private set; } = null!;

        [Output("community")]
        public Output<bool> Community { get; private set; } = null!;

        [Output("config")]
        public Output<Outputs.PulsarJobConfig?> Config { get; private set; } = null!;

        [Output("customer")]
        public Output<int> Customer { get; private set; } = null!;

        [Output("jobId")]
        public Output<string> JobId { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("shared")]
        public Output<bool?> Shared { get; private set; } = null!;

        [Output("typeId")]
        public Output<string> TypeId { get; private set; } = null!;

        [Output("weights")]
        public Output<ImmutableArray<Outputs.PulsarJobWeight>> Weights { get; private set; } = null!;


        /// <summary>
        /// Create a PulsarJob resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PulsarJob(string name, PulsarJobArgs args, CustomResourceOptions? options = null)
            : base("ns1:index/pulsarJob:PulsarJob", name, args ?? new PulsarJobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PulsarJob(string name, Input<string> id, PulsarJobState? state = null, CustomResourceOptions? options = null)
            : base("ns1:index/pulsarJob:PulsarJob", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PulsarJob resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PulsarJob Get(string name, Input<string> id, PulsarJobState? state = null, CustomResourceOptions? options = null)
        {
            return new PulsarJob(name, id, state, options);
        }
    }

    public sealed class PulsarJobArgs : Pulumi.ResourceArgs
    {
        [Input("active")]
        public Input<bool>? Active { get; set; }

        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

        [Input("blendMetricWeights")]
        public Input<Inputs.PulsarJobBlendMetricWeightsArgs>? BlendMetricWeights { get; set; }

        [Input("config")]
        public Input<Inputs.PulsarJobConfigArgs>? Config { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("shared")]
        public Input<bool>? Shared { get; set; }

        [Input("typeId", required: true)]
        public Input<string> TypeId { get; set; } = null!;

        [Input("weights")]
        private InputList<Inputs.PulsarJobWeightArgs>? _weights;
        public InputList<Inputs.PulsarJobWeightArgs> Weights
        {
            get => _weights ?? (_weights = new InputList<Inputs.PulsarJobWeightArgs>());
            set => _weights = value;
        }

        public PulsarJobArgs()
        {
        }
    }

    public sealed class PulsarJobState : Pulumi.ResourceArgs
    {
        [Input("active")]
        public Input<bool>? Active { get; set; }

        [Input("appId")]
        public Input<string>? AppId { get; set; }

        [Input("blendMetricWeights")]
        public Input<Inputs.PulsarJobBlendMetricWeightsGetArgs>? BlendMetricWeights { get; set; }

        [Input("community")]
        public Input<bool>? Community { get; set; }

        [Input("config")]
        public Input<Inputs.PulsarJobConfigGetArgs>? Config { get; set; }

        [Input("customer")]
        public Input<int>? Customer { get; set; }

        [Input("jobId")]
        public Input<string>? JobId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("shared")]
        public Input<bool>? Shared { get; set; }

        [Input("typeId")]
        public Input<string>? TypeId { get; set; }

        [Input("weights")]
        private InputList<Inputs.PulsarJobWeightGetArgs>? _weights;
        public InputList<Inputs.PulsarJobWeightGetArgs> Weights
        {
            get => _weights ?? (_weights = new InputList<Inputs.PulsarJobWeightGetArgs>());
            set => _weights = value;
        }

        public PulsarJobState()
        {
        }
    }
}
