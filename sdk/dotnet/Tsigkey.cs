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
    /// Provides a NS1 TSIG Key resource. This can be used to create, modify, and delete TSIG keys.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Ns1 = Pulumi.Ns1;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Ns1.Tsigkey("example", new()
    ///     {
    ///         Algorithm = "hmac-sha256",
    ///         Secret = "Ok1qR5IW1ajVka5cHPEJQIXfLyx5V3PSkFBROAzOn21JumDq6nIpoj6H8rfj5Uo+Ok55ZWQ0Wgrf302fDscHLA==",
    ///     });
    /// 
    /// });
    /// ```
    /// ## NS1 Documentation
    /// 
    /// [TSIG Keys Api Doc](https://ns1.com/api/#tsig)
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import ns1:index/tsigkey:Tsigkey importTest &lt;name&gt;`
    /// ```
    /// </summary>
    [Ns1ResourceType("ns1:index/tsigkey:Tsigkey")]
    public partial class Tsigkey : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The algorithm used to hash the TSIG key's secret.
        /// </summary>
        [Output("algorithm")]
        public Output<string> Algorithm { get; private set; } = null!;

        /// <summary>
        /// The free form name of the tsigkey.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The key's secret to be hashed.
        /// </summary>
        [Output("secret")]
        public Output<string> Secret { get; private set; } = null!;


        /// <summary>
        /// Create a Tsigkey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Tsigkey(string name, TsigkeyArgs args, CustomResourceOptions? options = null)
            : base("ns1:index/tsigkey:Tsigkey", name, args ?? new TsigkeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Tsigkey(string name, Input<string> id, TsigkeyState? state = null, CustomResourceOptions? options = null)
            : base("ns1:index/tsigkey:Tsigkey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Tsigkey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Tsigkey Get(string name, Input<string> id, TsigkeyState? state = null, CustomResourceOptions? options = null)
        {
            return new Tsigkey(name, id, state, options);
        }
    }

    public sealed class TsigkeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The algorithm used to hash the TSIG key's secret.
        /// </summary>
        [Input("algorithm", required: true)]
        public Input<string> Algorithm { get; set; } = null!;

        /// <summary>
        /// The free form name of the tsigkey.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The key's secret to be hashed.
        /// </summary>
        [Input("secret", required: true)]
        public Input<string> Secret { get; set; } = null!;

        public TsigkeyArgs()
        {
        }
        public static new TsigkeyArgs Empty => new TsigkeyArgs();
    }

    public sealed class TsigkeyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The algorithm used to hash the TSIG key's secret.
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        /// <summary>
        /// The free form name of the tsigkey.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The key's secret to be hashed.
        /// </summary>
        [Input("secret")]
        public Input<string>? Secret { get; set; }

        public TsigkeyState()
        {
        }
        public static new TsigkeyState Empty => new TsigkeyState();
    }
}
