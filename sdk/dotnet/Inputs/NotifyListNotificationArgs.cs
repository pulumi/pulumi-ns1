// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ns1.Inputs
{

    public sealed class NotifyListNotificationArgs : global::Pulumi.ResourceArgs
    {
        [Input("config", required: true)]
        private InputMap<string>? _config;

        /// <summary>
        /// Configuration details for the given notifier type.
        /// </summary>
        public InputMap<string> Config
        {
            get => _config ?? (_config = new InputMap<string>());
            set => _config = value;
        }

        /// <summary>
        /// The type of notifier. Available notifiers are indicated in /notifytypes endpoint.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public NotifyListNotificationArgs()
        {
        }
        public static new NotifyListNotificationArgs Empty => new NotifyListNotificationArgs();
    }
}
