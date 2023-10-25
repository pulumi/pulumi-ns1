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
    /// Provides a NS1 Notify List resource. This can be used to create, modify, and delete notify lists.
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
    ///     var nl = new Ns1.NotifyList("nl", new()
    ///     {
    ///         Notifications = new[]
    ///         {
    ///             new Ns1.Inputs.NotifyListNotificationArgs
    ///             {
    ///                 Config = 
    ///                 {
    ///                     { "url", "http://www.mywebhook.com" },
    ///                 },
    ///                 Type = "webhook",
    ///             },
    ///             new Ns1.Inputs.NotifyListNotificationArgs
    ///             {
    ///                 Config = 
    ///                 {
    ///                     { "email", "test@test.com" },
    ///                 },
    ///                 Type = "email",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ## NS1 Documentation
    /// 
    /// [NotifyList Api Doc](https://ns1.com/api#notification-lists)
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import ns1:index/notifyList:NotifyList &lt;name&gt; &lt;notifylist_id&gt;`
    /// ```
    /// </summary>
    [Ns1ResourceType("ns1:index/notifyList:NotifyList")]
    public partial class NotifyList : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The free-form display name for the notify list.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of notifiers. All notifiers in a notification list will receive notifications whenever an event is send to the list (e.g., when a monitoring job fails). Notifiers are documented below.
        /// </summary>
        [Output("notifications")]
        public Output<ImmutableArray<Outputs.NotifyListNotification>> Notifications { get; private set; } = null!;


        /// <summary>
        /// Create a NotifyList resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NotifyList(string name, NotifyListArgs? args = null, CustomResourceOptions? options = null)
            : base("ns1:index/notifyList:NotifyList", name, args ?? new NotifyListArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NotifyList(string name, Input<string> id, NotifyListState? state = null, CustomResourceOptions? options = null)
            : base("ns1:index/notifyList:NotifyList", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NotifyList resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NotifyList Get(string name, Input<string> id, NotifyListState? state = null, CustomResourceOptions? options = null)
        {
            return new NotifyList(name, id, state, options);
        }
    }

    public sealed class NotifyListArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The free-form display name for the notify list.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notifications")]
        private InputList<Inputs.NotifyListNotificationArgs>? _notifications;

        /// <summary>
        /// A list of notifiers. All notifiers in a notification list will receive notifications whenever an event is send to the list (e.g., when a monitoring job fails). Notifiers are documented below.
        /// </summary>
        public InputList<Inputs.NotifyListNotificationArgs> Notifications
        {
            get => _notifications ?? (_notifications = new InputList<Inputs.NotifyListNotificationArgs>());
            set => _notifications = value;
        }

        public NotifyListArgs()
        {
        }
        public static new NotifyListArgs Empty => new NotifyListArgs();
    }

    public sealed class NotifyListState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The free-form display name for the notify list.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("notifications")]
        private InputList<Inputs.NotifyListNotificationGetArgs>? _notifications;

        /// <summary>
        /// A list of notifiers. All notifiers in a notification list will receive notifications whenever an event is send to the list (e.g., when a monitoring job fails). Notifiers are documented below.
        /// </summary>
        public InputList<Inputs.NotifyListNotificationGetArgs> Notifications
        {
            get => _notifications ?? (_notifications = new InputList<Inputs.NotifyListNotificationGetArgs>());
            set => _notifications = value;
        }

        public NotifyListState()
        {
        }
        public static new NotifyListState Empty => new NotifyListState();
    }
}
