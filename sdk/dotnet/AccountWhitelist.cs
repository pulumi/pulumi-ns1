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
    /// Provides a NS1 Global IP Whitelist resource.
    /// 
    /// This can be used to create, modify, and delete Global IP Whitelists.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Ns1 = Pulumi.Ns1;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Ns1.AccountWhitelist("example", new()
    ///     {
    ///         Name = "Example Whitelist",
    ///         Values = new[]
    ///         {
    ///             "1.1.1.1",
    ///             "2.2.2.2",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// &gt; You current source IP must be present in one of the whitelists to prevent locking yourself out.
    /// 
    /// ## NS1 Documentation
    /// 
    /// [Global IP Whitelist Doc](https://ns1.com/api?docId=2282)
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import ns1:index/accountWhitelist:AccountWhitelist example &lt;ID&gt;`
    /// ```
    /// </summary>
    [Ns1ResourceType("ns1:index/accountWhitelist:AccountWhitelist")]
    public partial class AccountWhitelist : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The free form name of the whitelist.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Array of IP addresses/networks from which to allow access.
        /// </summary>
        [Output("values")]
        public Output<ImmutableArray<string>> Values { get; private set; } = null!;


        /// <summary>
        /// Create a AccountWhitelist resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccountWhitelist(string name, AccountWhitelistArgs args, CustomResourceOptions? options = null)
            : base("ns1:index/accountWhitelist:AccountWhitelist", name, args ?? new AccountWhitelistArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccountWhitelist(string name, Input<string> id, AccountWhitelistState? state = null, CustomResourceOptions? options = null)
            : base("ns1:index/accountWhitelist:AccountWhitelist", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AccountWhitelist resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccountWhitelist Get(string name, Input<string> id, AccountWhitelistState? state = null, CustomResourceOptions? options = null)
        {
            return new AccountWhitelist(name, id, state, options);
        }
    }

    public sealed class AccountWhitelistArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The free form name of the whitelist.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("values", required: true)]
        private InputList<string>? _values;

        /// <summary>
        /// Array of IP addresses/networks from which to allow access.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public AccountWhitelistArgs()
        {
        }
        public static new AccountWhitelistArgs Empty => new AccountWhitelistArgs();
    }

    public sealed class AccountWhitelistState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The free form name of the whitelist.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("values")]
        private InputList<string>? _values;

        /// <summary>
        /// Array of IP addresses/networks from which to allow access.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public AccountWhitelistState()
        {
        }
        public static new AccountWhitelistState Empty => new AccountWhitelistState();
    }
}
