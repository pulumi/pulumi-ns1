// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides details about a NS1 Zone. Use this if you would simply like to read
 * information from NS1 into your configurations. For read/write operations, you
 * should use a resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ns1 from "@pulumi/ns1";
 *
 * // Get details about a NS1 Zone.
 * const example = ns1.getZone({
 *     zone: "terraform.example.io",
 * });
 * ```
 */
export function getZone(args: GetZoneArgs, opts?: pulumi.InvokeOptions): Promise<GetZoneResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ns1:index/getZone:getZone", {
        "additionalPorts": args.additionalPorts,
        "additionalPrimaries": args.additionalPrimaries,
        "primaryNetwork": args.primaryNetwork,
        "primaryPort": args.primaryPort,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getZone.
 */
export interface GetZoneArgs {
    additionalPorts?: number[];
    /**
     * List of additional IPv4 addresses for the primary
     * zone.
     */
    additionalPrimaries?: string[];
    primaryNetwork?: number;
    primaryPort?: number;
    /**
     * The domain name of the zone.
     */
    zone: string;
}

/**
 * A collection of values returned by getZone.
 */
export interface GetZoneResult {
    readonly additionalPorts?: number[];
    /**
     * List of additional IPv4 addresses for the primary
     * zone.
     */
    readonly additionalPrimaries?: string[];
    /**
     * Authoritative Name Servers.
     */
    readonly dnsServers: string;
    /**
     * Whether or not DNSSEC is enabled for the zone.
     */
    readonly dnssec: boolean;
    /**
     * The SOA Expiry.
     */
    readonly expiry: number;
    /**
     * The SOA Hostmaster.
     */
    readonly hostmaster: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The linked target zone.
     */
    readonly link: string;
    /**
     * List of network IDs (`int`) for which the zone should be made
     * available. Default is network 0, the primary NSONE Global Network.
     */
    readonly networks: number[];
    /**
     * The SOA NX TTL.
     */
    readonly nxTtl: number;
    /**
     * The primary zones' IPv4 address.
     */
    readonly primary: string;
    readonly primaryNetwork?: number;
    readonly primaryPort?: number;
    /**
     * The SOA Refresh.
     */
    readonly refresh: number;
    /**
     * The SOA Retry.
     */
    readonly retry: number;
    /**
     * List of secondary servers. Secondaries is
     * documented below.
     */
    readonly secondaries: outputs.GetZoneSecondary[];
    readonly tags: {[key: string]: string};
    /**
     * The SOA TTL.
     */
    readonly ttl: number;
    readonly zone: string;
}
/**
 * Provides details about a NS1 Zone. Use this if you would simply like to read
 * information from NS1 into your configurations. For read/write operations, you
 * should use a resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ns1 from "@pulumi/ns1";
 *
 * // Get details about a NS1 Zone.
 * const example = ns1.getZone({
 *     zone: "terraform.example.io",
 * });
 * ```
 */
export function getZoneOutput(args: GetZoneOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetZoneResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ns1:index/getZone:getZone", {
        "additionalPorts": args.additionalPorts,
        "additionalPrimaries": args.additionalPrimaries,
        "primaryNetwork": args.primaryNetwork,
        "primaryPort": args.primaryPort,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getZone.
 */
export interface GetZoneOutputArgs {
    additionalPorts?: pulumi.Input<pulumi.Input<number>[]>;
    /**
     * List of additional IPv4 addresses for the primary
     * zone.
     */
    additionalPrimaries?: pulumi.Input<pulumi.Input<string>[]>;
    primaryNetwork?: pulumi.Input<number>;
    primaryPort?: pulumi.Input<number>;
    /**
     * The domain name of the zone.
     */
    zone: pulumi.Input<string>;
}
