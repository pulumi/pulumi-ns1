// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ns1.Inputs
{

    public sealed class NotifyListNotificationArgs : Pulumi.ResourceArgs
    {
        [Input("config", required: true)]
        private InputMap<object>? _config;

        /// <summary>
        /// Configuration details for the given notifier type.
        /// </summary>
        public InputMap<object> Config
        {
            get => _config ?? (_config = new InputMap<object>());
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
    }
}
