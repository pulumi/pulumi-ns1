// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ns1
{
    /// <summary>
    /// Provides a NS1 Data Source resource. This can be used to create, modify, and delete data sources.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Ns1 = Pulumi.Ns1;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Ns1.DataSource("example", new()
    ///     {
    ///         Name = "example",
    ///         Sourcetype = "nsone_v1",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## NS1 Documentation
    /// 
    /// [Datasource Api Doc](https://ns1.com/api#data-sources)
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import ns1:index/dataSource:DataSource &lt;name&gt; &lt;datasource_id&gt;`
    /// ```
    /// </summary>
    [Ns1ResourceType("ns1:index/dataSource:DataSource")]
    public partial class DataSource : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The data source configuration, determined by its type,
        /// matching the specification in `config` from /data/sourcetypes.
        /// </summary>
        [Output("config")]
        public Output<ImmutableDictionary<string, string>?> Config { get; private set; } = null!;

        /// <summary>
        /// The free form name of the data source.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The data sources type, listed in API endpoint https://api.nsone.net/v1/data/sourcetypes.
        /// </summary>
        [Output("sourcetype")]
        public Output<string> Sourcetype { get; private set; } = null!;


        /// <summary>
        /// Create a DataSource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DataSource(string name, DataSourceArgs args, CustomResourceOptions? options = null)
            : base("ns1:index/dataSource:DataSource", name, args ?? new DataSourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DataSource(string name, Input<string> id, DataSourceState? state = null, CustomResourceOptions? options = null)
            : base("ns1:index/dataSource:DataSource", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DataSource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DataSource Get(string name, Input<string> id, DataSourceState? state = null, CustomResourceOptions? options = null)
        {
            return new DataSource(name, id, state, options);
        }
    }

    public sealed class DataSourceArgs : global::Pulumi.ResourceArgs
    {
        [Input("config")]
        private InputMap<string>? _config;

        /// <summary>
        /// The data source configuration, determined by its type,
        /// matching the specification in `config` from /data/sourcetypes.
        /// </summary>
        public InputMap<string> Config
        {
            get => _config ?? (_config = new InputMap<string>());
            set => _config = value;
        }

        /// <summary>
        /// The free form name of the data source.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The data sources type, listed in API endpoint https://api.nsone.net/v1/data/sourcetypes.
        /// </summary>
        [Input("sourcetype", required: true)]
        public Input<string> Sourcetype { get; set; } = null!;

        public DataSourceArgs()
        {
        }
        public static new DataSourceArgs Empty => new DataSourceArgs();
    }

    public sealed class DataSourceState : global::Pulumi.ResourceArgs
    {
        [Input("config")]
        private InputMap<string>? _config;

        /// <summary>
        /// The data source configuration, determined by its type,
        /// matching the specification in `config` from /data/sourcetypes.
        /// </summary>
        public InputMap<string> Config
        {
            get => _config ?? (_config = new InputMap<string>());
            set => _config = value;
        }

        /// <summary>
        /// The free form name of the data source.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The data sources type, listed in API endpoint https://api.nsone.net/v1/data/sourcetypes.
        /// </summary>
        [Input("sourcetype")]
        public Input<string>? Sourcetype { get; set; }

        public DataSourceState()
        {
        }
        public static new DataSourceState Empty => new DataSourceState();
    }
}
