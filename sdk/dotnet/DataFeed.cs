// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ns1
{
    /// <summary>
    /// Provides a NS1 Data Feed resource. This can be used to create, modify, and delete data feeds.
    /// 
    /// 
    /// ## NS1 Documentation
    /// 
    /// [Datafeed Api Doc](https://ns1.com/api#data-feeds)
    /// </summary>
    public partial class DataFeed : Pulumi.CustomResource
    {
        /// <summary>
        /// The feeds configuration matching the specification in
        /// `feed_config` from /data/sourcetypes.
        /// </summary>
        [Output("config")]
        public Output<ImmutableDictionary<string, object>?> Config { get; private set; } = null!;

        /// <summary>
        /// The free form name of the data feed.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The data source id that this feed is connected to.
        /// </summary>
        [Output("sourceId")]
        public Output<string> SourceId { get; private set; } = null!;


        /// <summary>
        /// Create a DataFeed resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DataFeed(string name, DataFeedArgs args, CustomResourceOptions? options = null)
            : base("ns1:index/dataFeed:DataFeed", name, args ?? new DataFeedArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DataFeed(string name, Input<string> id, DataFeedState? state = null, CustomResourceOptions? options = null)
            : base("ns1:index/dataFeed:DataFeed", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DataFeed resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DataFeed Get(string name, Input<string> id, DataFeedState? state = null, CustomResourceOptions? options = null)
        {
            return new DataFeed(name, id, state, options);
        }
    }

    public sealed class DataFeedArgs : Pulumi.ResourceArgs
    {
        [Input("config")]
        private InputMap<object>? _config;

        /// <summary>
        /// The feeds configuration matching the specification in
        /// `feed_config` from /data/sourcetypes.
        /// </summary>
        public InputMap<object> Config
        {
            get => _config ?? (_config = new InputMap<object>());
            set => _config = value;
        }

        /// <summary>
        /// The free form name of the data feed.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The data source id that this feed is connected to.
        /// </summary>
        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        public DataFeedArgs()
        {
        }
    }

    public sealed class DataFeedState : Pulumi.ResourceArgs
    {
        [Input("config")]
        private InputMap<object>? _config;

        /// <summary>
        /// The feeds configuration matching the specification in
        /// `feed_config` from /data/sourcetypes.
        /// </summary>
        public InputMap<object> Config
        {
            get => _config ?? (_config = new InputMap<object>());
            set => _config = value;
        }

        /// <summary>
        /// The free form name of the data feed.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The data source id that this feed is connected to.
        /// </summary>
        [Input("sourceId")]
        public Input<string>? SourceId { get; set; }

        public DataFeedState()
        {
        }
    }
}
