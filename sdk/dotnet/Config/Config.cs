// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumi.Ns1
{
    public static class Config
    {
        [global::System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly global::Pulumi.Config __config = new global::Pulumi.Config("ns1");

        private static readonly __Value<string?> _apikey = new __Value<string?>(() => __config.Get("apikey"));
        /// <summary>
        /// The ns1 API key (required)
        /// </summary>
        public static string? Apikey
        {
            get => _apikey.Get();
            set => _apikey.Set(value);
        }

        private static readonly __Value<bool?> _enableDdi = new __Value<bool?>(() => __config.GetBoolean("enableDdi"));
        /// <summary>
        /// Deprecated, no longer in use
        /// </summary>
        public static bool? EnableDdi
        {
            get => _enableDdi.Get();
            set => _enableDdi.Set(value);
        }

        private static readonly __Value<string?> _endpoint = new __Value<string?>(() => __config.Get("endpoint"));
        /// <summary>
        /// URL prefix (including version) for API calls
        /// </summary>
        public static string? Endpoint
        {
            get => _endpoint.Get();
            set => _endpoint.Set(value);
        }

        private static readonly __Value<bool?> _ignoreSsl = new __Value<bool?>(() => __config.GetBoolean("ignoreSsl"));
        /// <summary>
        /// Don't validate server SSL/TLS certificate
        /// </summary>
        public static bool? IgnoreSsl
        {
            get => _ignoreSsl.Get();
            set => _ignoreSsl.Set(value);
        }

        private static readonly __Value<int?> _rateLimitParallelism = new __Value<int?>(() => __config.GetInt32("rateLimitParallelism"));
        /// <summary>
        /// Tune response to rate limits, see docs
        /// </summary>
        public static int? RateLimitParallelism
        {
            get => _rateLimitParallelism.Get();
            set => _rateLimitParallelism.Set(value);
        }

        private static readonly __Value<int?> _retryMax = new __Value<int?>(() => __config.GetInt32("retryMax"));
        /// <summary>
        /// Maximum retries for 50x errors (-1 to disable)
        /// </summary>
        public static int? RetryMax
        {
            get => _retryMax.Get();
            set => _retryMax.Set(value);
        }

        private static readonly __Value<string?> _userAgent = new __Value<string?>(() => __config.Get("userAgent"));
        /// <summary>
        /// User-Agent string to use in NS1 API requests
        /// </summary>
        public static string? UserAgent
        {
            get => _userAgent.Get();
            set => _userAgent.Set(value);
        }

    }
}