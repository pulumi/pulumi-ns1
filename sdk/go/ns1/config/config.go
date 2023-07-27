// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi-ns1/sdk/v3/go/ns1/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

var _ = internal.GetEnvOrDefault

// The ns1 API key (required)
func GetApikey(ctx *pulumi.Context) string {
	return config.Get(ctx, "ns1:apikey")
}

// Deprecated, no longer in use
func GetEnableDdi(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "ns1:enableDdi")
}

// URL prefix (including version) for API calls
func GetEndpoint(ctx *pulumi.Context) string {
	return config.Get(ctx, "ns1:endpoint")
}

// Don't validate server SSL/TLS certificate
func GetIgnoreSsl(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "ns1:ignoreSsl")
}

// Tune response to rate limits, see docs
func GetRateLimitParallelism(ctx *pulumi.Context) int {
	return config.GetInt(ctx, "ns1:rateLimitParallelism")
}

// Maximum retries for 50x errors (-1 to disable)
func GetRetryMax(ctx *pulumi.Context) int {
	return config.GetInt(ctx, "ns1:retryMax")
}

// User-Agent string to use in NS1 API requests
func GetUserAgent(ctx *pulumi.Context) string {
	return config.Get(ctx, "ns1:userAgent")
}
