// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;

namespace Pulumi.Ns1
{
    public static class Config
    {
        private static readonly Pulumi.Config __config = new Pulumi.Config("ns1");
        /// <summary>
        /// The ns1 API key, this is required
        /// </summary>
        public static string? Apikey { get; set; } = __config.Get("apikey") ?? Utilities.GetEnv("NS1_APIKEY");

        public static bool? EnableDdi { get; set; } = __config.GetBoolean("enableDdi");

        public static string? Endpoint { get; set; } = __config.Get("endpoint") ?? Utilities.GetEnv("NS1_ENDPOINT");

        public static bool? IgnoreSsl { get; set; } = __config.GetBoolean("ignoreSsl");

        public static int? RateLimitParallelism { get; set; } = __config.GetInt32("rateLimitParallelism");

    }
}
