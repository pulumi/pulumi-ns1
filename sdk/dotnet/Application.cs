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
    /// Provides a NS1 Pulsar application resource. This can be used to create, modify, and delete applications.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Ns1 = Pulumi.Ns1;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         // Create a new pulsar application with default config
    ///         var ns1App = new Ns1.Application("ns1App", new Ns1.ApplicationArgs
    ///         {
    ///             DefaultConfig = new Ns1.Inputs.ApplicationDefaultConfigArgs
    ///             {
    ///                 Http = true,
    ///                 Https = false,
    ///                 Job_timeout_millis = 100,
    ///                 Request_timeout_millis = 100,
    ///                 Static_values = true,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## NS1 Documentation
    /// 
    /// [Application Api Docs](https://ns1.com/api#get-list-pulsar-applications)
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import ns1:index/application:Application `ns1_application`
    /// ```
    /// 
    ///  So for the example above
    /// 
    /// ```sh
    ///  $ pulumi import ns1:index/application:Application example terraform.example.io`
    /// ```
    /// </summary>
    [Ns1ResourceType("ns1:index/application:Application")]
    public partial class Application : Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates whether or not this application is currently active and usable for traffic
        /// steering.
        /// </summary>
        [Output("active")]
        public Output<bool?> Active { get; private set; } = null!;

        /// <summary>
        /// The amount of time (in milliseconds) the browser should wait before running
        /// measurements.
        /// </summary>
        [Output("browserWaitMillis")]
        public Output<int?> BrowserWaitMillis { get; private set; } = null!;

        /// <summary>
        /// -(Optional) Default job configuration. If a field is present here and not on a specific job
        /// associated with this application, the default value specified here is used..
        /// </summary>
        [Output("defaultConfig")]
        public Output<Outputs.ApplicationDefaultConfig?> DefaultConfig { get; private set; } = null!;

        /// <summary>
        /// -(Optional) Number of jobs to measure per user impression.
        /// </summary>
        [Output("jobsPerTransaction")]
        public Output<int?> JobsPerTransaction { get; private set; } = null!;

        /// <summary>
        /// Descriptive name for this Pulsar app.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a Application resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Application(string name, ApplicationArgs? args = null, CustomResourceOptions? options = null)
            : base("ns1:index/application:Application", name, args ?? new ApplicationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Application(string name, Input<string> id, ApplicationState? state = null, CustomResourceOptions? options = null)
            : base("ns1:index/application:Application", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Application resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Application Get(string name, Input<string> id, ApplicationState? state = null, CustomResourceOptions? options = null)
        {
            return new Application(name, id, state, options);
        }
    }

    public sealed class ApplicationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether or not this application is currently active and usable for traffic
        /// steering.
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// The amount of time (in milliseconds) the browser should wait before running
        /// measurements.
        /// </summary>
        [Input("browserWaitMillis")]
        public Input<int>? BrowserWaitMillis { get; set; }

        /// <summary>
        /// -(Optional) Default job configuration. If a field is present here and not on a specific job
        /// associated with this application, the default value specified here is used..
        /// </summary>
        [Input("defaultConfig")]
        public Input<Inputs.ApplicationDefaultConfigArgs>? DefaultConfig { get; set; }

        /// <summary>
        /// -(Optional) Number of jobs to measure per user impression.
        /// </summary>
        [Input("jobsPerTransaction")]
        public Input<int>? JobsPerTransaction { get; set; }

        /// <summary>
        /// Descriptive name for this Pulsar app.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ApplicationArgs()
        {
        }
    }

    public sealed class ApplicationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether or not this application is currently active and usable for traffic
        /// steering.
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// The amount of time (in milliseconds) the browser should wait before running
        /// measurements.
        /// </summary>
        [Input("browserWaitMillis")]
        public Input<int>? BrowserWaitMillis { get; set; }

        /// <summary>
        /// -(Optional) Default job configuration. If a field is present here and not on a specific job
        /// associated with this application, the default value specified here is used..
        /// </summary>
        [Input("defaultConfig")]
        public Input<Inputs.ApplicationDefaultConfigGetArgs>? DefaultConfig { get; set; }

        /// <summary>
        /// -(Optional) Number of jobs to measure per user impression.
        /// </summary>
        [Input("jobsPerTransaction")]
        public Input<int>? JobsPerTransaction { get; set; }

        /// <summary>
        /// Descriptive name for this Pulsar app.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ApplicationState()
        {
        }
    }
}
