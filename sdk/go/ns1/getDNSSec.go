// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ns1

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides DNSSEC details about a NS1 Zone.
func GetDNSSec(ctx *pulumi.Context, args *GetDNSSecArgs, opts ...pulumi.InvokeOption) (*GetDNSSecResult, error) {
	var rv GetDNSSecResult
	err := ctx.Invoke("ns1:index/getDNSSec:getDNSSec", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDNSSec.
type GetDNSSecArgs struct {
	// The name of the zone to get DNSSEC details for.
	Zone string `pulumi:"zone"`
}

// A collection of values returned by getDNSSec.
type GetDNSSecResult struct {
	// (Computed) - Delegation field is documented
	// below.
	Delegation GetDNSSecDelegation `pulumi:"delegation"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// (Computed) - Keys field is documented below.
	Keys GetDNSSecKeys `pulumi:"keys"`
	Zone string        `pulumi:"zone"`
}
