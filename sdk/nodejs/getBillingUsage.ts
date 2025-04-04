// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides billing usage details about a NS1 account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ns1 from "@pulumi/ns1";
 *
 * // Get query usage data for the given timeframe
 * const queries = new ns1.index.BillingUsage("queries", {
 *     metricType: "queries",
 *     from: 1731605824,
 *     to: 1734197824,
 * });
 * // Get account limits data for the given timeframe
 * const limits = new ns1.index.BillingUsage("limits", {
 *     metricType: "limits",
 *     from: 1731605824,
 *     to: 1734197824,
 * });
 * // Get RUM decisions usage data for the given timeframe
 * const decisions = new ns1.index.BillingUsage("decisions", {
 *     metricType: "decisions",
 *     from: 1731605824,
 *     to: 1734197824,
 * });
 * // Get filter chains usage data
 * const filterChains = new ns1.index.BillingUsage("filter_chains", {metricType: "filter-chains"});
 * // Get monitoring jobs usage data
 * const monitors = new ns1.index.BillingUsage("monitors", {metricType: "monitors"});
 * // Get records usage data
 * const records = new ns1.index.BillingUsage("records", {metricType: "records"});
 * export const totalQueries = queries.cleanQueries;
 * export const totalDdosQueries = queries.ddosQueries;
 * export const totalNxdResponses = queries.nxdResponses;
 * export const queriesLimit = limits.queriesLimit;
 * export const totalDecisions = decisions.totalUsage;
 * export const decisionsLimit = limits.decisionsLimit;
 * export const totalFilterChains = filterChains.totalUsage;
 * export const filterChainsLimit = limits.filterChainsLimit;
 * export const totalMonitors = monitors.totalUsage;
 * export const monitorsLimit = limits.monitorsLimit;
 * export const totalRecords = records.totalUsage;
 * export const recordsLimit = limits.recordsLimit;
 * ```
 */
export function getBillingUsage(args: GetBillingUsageArgs, opts?: pulumi.InvokeOptions): Promise<GetBillingUsageResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ns1:index/getBillingUsage:getBillingUsage", {
        "from": args.from,
        "metricType": args.metricType,
        "to": args.to,
    }, opts);
}

/**
 * A collection of arguments for invoking getBillingUsage.
 */
export interface GetBillingUsageArgs {
    /**
     * The start timestamp for the data range in Unix epoch format.
     */
    from: number;
    /**
     * The type of billing metric to retrieve. Must be one of: `queries`, `limits`, `decisions`, `filter-chains`, `monitors`, `records`.
     */
    metricType: string;
    /**
     * The end timestamp for the data range in Unix epoch format.
     */
    to: number;
}

/**
 * A collection of values returned by getBillingUsage.
 */
export interface GetBillingUsageResult {
    /**
     * (Computed) A list of network-specific query data containing:
     */
    readonly byNetworks: outputs.GetBillingUsageByNetwork[];
    /**
     * (Computed) The queries limit for the China network.
     */
    readonly chinaQueriesLimit: number;
    /**
     * Clean queries for this day.
     */
    readonly cleanQueries: number;
    /**
     * (Computed) Whether DDoS Protection is enabled.
     */
    readonly ddosProtectionEnabled: boolean;
    /**
     * DDoS queries for this day.
     */
    readonly ddosQueries: number;
    /**
     * (Computed) The RUM decisions limit for this billing cycle.
     */
    readonly decisionsLimit: number;
    /**
     * (Computed) The filter chains limit for this billing cycle.
     */
    readonly filterChainsLimit: number;
    readonly from: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * (Computed) Whether dedicated DNS usage counts towards managed DNS usage.
     */
    readonly includeDedicatedDnsNetworkInManagedDnsUsage: boolean;
    readonly metricType: string;
    /**
     * (Computed) The monitoring jobs limit for this billing cycle.
     */
    readonly monitorsLimit: number;
    /**
     * (Computed) Whether NXD Protection is enabled.
     */
    readonly nxdProtectionEnabled: boolean;
    /**
     * NXD responses for this day.
     */
    readonly nxdResponses: number;
    /**
     * (Computed) The queries limit for this billing cycle.
     */
    readonly queriesLimit: number;
    /**
     * (Computed) The records limit for this billing cycle.
     */
    readonly recordsLimit: number;
    readonly to: number;
    /**
     * (Computed) The total usage count for the metric. Available for `decisions`, `filter-chains`, `monitors`, and `records` metrics.
     */
    readonly totalUsage: number;
}
/**
 * Provides billing usage details about a NS1 account.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ns1 from "@pulumi/ns1";
 *
 * // Get query usage data for the given timeframe
 * const queries = new ns1.index.BillingUsage("queries", {
 *     metricType: "queries",
 *     from: 1731605824,
 *     to: 1734197824,
 * });
 * // Get account limits data for the given timeframe
 * const limits = new ns1.index.BillingUsage("limits", {
 *     metricType: "limits",
 *     from: 1731605824,
 *     to: 1734197824,
 * });
 * // Get RUM decisions usage data for the given timeframe
 * const decisions = new ns1.index.BillingUsage("decisions", {
 *     metricType: "decisions",
 *     from: 1731605824,
 *     to: 1734197824,
 * });
 * // Get filter chains usage data
 * const filterChains = new ns1.index.BillingUsage("filter_chains", {metricType: "filter-chains"});
 * // Get monitoring jobs usage data
 * const monitors = new ns1.index.BillingUsage("monitors", {metricType: "monitors"});
 * // Get records usage data
 * const records = new ns1.index.BillingUsage("records", {metricType: "records"});
 * export const totalQueries = queries.cleanQueries;
 * export const totalDdosQueries = queries.ddosQueries;
 * export const totalNxdResponses = queries.nxdResponses;
 * export const queriesLimit = limits.queriesLimit;
 * export const totalDecisions = decisions.totalUsage;
 * export const decisionsLimit = limits.decisionsLimit;
 * export const totalFilterChains = filterChains.totalUsage;
 * export const filterChainsLimit = limits.filterChainsLimit;
 * export const totalMonitors = monitors.totalUsage;
 * export const monitorsLimit = limits.monitorsLimit;
 * export const totalRecords = records.totalUsage;
 * export const recordsLimit = limits.recordsLimit;
 * ```
 */
export function getBillingUsageOutput(args: GetBillingUsageOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetBillingUsageResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ns1:index/getBillingUsage:getBillingUsage", {
        "from": args.from,
        "metricType": args.metricType,
        "to": args.to,
    }, opts);
}

/**
 * A collection of arguments for invoking getBillingUsage.
 */
export interface GetBillingUsageOutputArgs {
    /**
     * The start timestamp for the data range in Unix epoch format.
     */
    from: pulumi.Input<number>;
    /**
     * The type of billing metric to retrieve. Must be one of: `queries`, `limits`, `decisions`, `filter-chains`, `monitors`, `records`.
     */
    metricType: pulumi.Input<string>;
    /**
     * The end timestamp for the data range in Unix epoch format.
     */
    to: pulumi.Input<number>;
}
