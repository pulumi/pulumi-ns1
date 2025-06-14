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
    /// The provider type for the ns1 package. By default, resources use package-wide configuration
    /// settings, however an explicit `Provider` instance may be created and passed during resource
    /// construction to achieve fine-grained programmatic control over provider settings. See the
    /// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
    /// </summary>
    [Ns1ResourceType("pulumi:providers:ns1")]
    public partial class Provider : global::Pulumi.ProviderResource
    {
        /// <summary>
        /// The ns1 API key (required)
        /// </summary>
        [Output("apikey")]
        public Output<string?> Apikey { get; private set; } = null!;

        /// <summary>
        /// URL prefix (including version) for API calls
        /// </summary>
        [Output("endpoint")]
        public Output<string?> Endpoint { get; private set; } = null!;

        /// <summary>
        /// User-Agent string to use in NS1 API requests
        /// </summary>
        [Output("userAgent")]
        public Output<string?> UserAgent { get; private set; } = null!;


        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs? args = null, CustomResourceOptions? options = null)
            : base("ns1", name, args ?? new ProviderArgs(), MakeResourceOptions(options, ""))
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
        /// This function returns a Terraform config object with terraform-namecased keys,to be used with the Terraform Module Provider.
        /// </summary>
        public global::Pulumi.Output<ProviderTerraformConfigResult> TerraformConfig()
            => global::Pulumi.Deployment.Instance.Call<ProviderTerraformConfigResult>("pulumi:providers:ns1/terraformConfig", CallArgs.Empty, this);
    }

    public sealed class ProviderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ns1 API key (required)
        /// </summary>
        [Input("apikey")]
        public Input<string>? Apikey { get; set; }

        /// <summary>
        /// URL prefix (including version) for API calls
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// Don't validate server SSL/TLS certificate
        /// </summary>
        [Input("ignoreSsl", json: true)]
        public Input<bool>? IgnoreSsl { get; set; }

        /// <summary>
        /// Tune response to rate limits, see docs
        /// </summary>
        [Input("rateLimitParallelism", json: true)]
        public Input<int>? RateLimitParallelism { get; set; }

        /// <summary>
        /// Maximum retries for 50x errors (-1 to disable)
        /// </summary>
        [Input("retryMax", json: true)]
        public Input<int>? RetryMax { get; set; }

        /// <summary>
        /// User-Agent string to use in NS1 API requests
        /// </summary>
        [Input("userAgent")]
        public Input<string>? UserAgent { get; set; }

        public ProviderArgs()
        {
        }
        public static new ProviderArgs Empty => new ProviderArgs();
    }

    /// <summary>
    /// The results of the <see cref="Provider.TerraformConfig"/> method.
    /// </summary>
    [OutputType]
    public sealed class ProviderTerraformConfigResult
    {
        public readonly ImmutableDictionary<string, object> Result;

        [OutputConstructor]
        private ProviderTerraformConfigResult(ImmutableDictionary<string, object> result)
        {
            Result = result;
        }
    }
}
