// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides DNSSEC details about a NS1 Zone.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ns1 from "@pulumi/ns1";
 *
 * // Get DNSSEC details about a NS1 Zone.
 * const exampleZone = new ns1.Zone("example", {
 *     zone: "terraform.example.io",
 *     dnssec: true,
 * });
 * const example = ns1.getDNSSecOutput({
 *     zone: exampleZone.zone,
 * });
 * ```
 */
export function getDNSSec(args: GetDNSSecArgs, opts?: pulumi.InvokeOptions): Promise<GetDNSSecResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ns1:index/getDNSSec:getDNSSec", {
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getDNSSec.
 */
export interface GetDNSSecArgs {
    /**
     * The name of the zone to get DNSSEC details for.
     */
    zone: string;
}

/**
 * A collection of values returned by getDNSSec.
 */
export interface GetDNSSecResult {
    /**
     * (Computed) - Delegation field is documented
     * below.
     */
    readonly delegations: outputs.GetDNSSecDelegation[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * (Computed) - Keys field is documented below.
     */
    readonly keys: outputs.GetDNSSecKey[];
    readonly zone: string;
}
/**
 * Provides DNSSEC details about a NS1 Zone.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ns1 from "@pulumi/ns1";
 *
 * // Get DNSSEC details about a NS1 Zone.
 * const exampleZone = new ns1.Zone("example", {
 *     zone: "terraform.example.io",
 *     dnssec: true,
 * });
 * const example = ns1.getDNSSecOutput({
 *     zone: exampleZone.zone,
 * });
 * ```
 */
export function getDNSSecOutput(args: GetDNSSecOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDNSSecResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ns1:index/getDNSSec:getDNSSec", {
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getDNSSec.
 */
export interface GetDNSSecOutputArgs {
    /**
     * The name of the zone to get DNSSEC details for.
     */
    zone: pulumi.Input<string>;
}
