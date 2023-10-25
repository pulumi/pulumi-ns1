// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides details about a NS1 Record. Use this if you would simply like to read
 * information from NS1 into your configurations. For read/write operations, you
 * should use a resource.
 */
export function getRecord(args: GetRecordArgs, opts?: pulumi.InvokeOptions): Promise<GetRecordResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ns1:index/getRecord:getRecord", {
        "domain": args.domain,
        "type": args.type,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getRecord.
 */
export interface GetRecordArgs {
    /**
     * The records' domain.
     */
    domain: string;
    /**
     * The records' RR type.
     */
    type: string;
    /**
     * The zone the record belongs to.
     */
    zone: string;
}

/**
 * A collection of values returned by getRecord.
 */
export interface GetRecordResult {
    /**
     * List of NS1 answers.
     */
    readonly answers: outputs.GetRecordAnswer[];
    readonly domain: string;
    /**
     * List of NS1 filters.
     */
    readonly filters: outputs.GetRecordFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The target record this links to.
     */
    readonly link: string;
    /**
     * Map of metadata
     */
    readonly meta: {[key: string]: any};
    readonly overrideTtl: boolean;
    /**
     * List of regions.
     */
    readonly regions: outputs.GetRecordRegion[];
    readonly shortAnswers: string[];
    readonly tags: {[key: string]: any};
    /**
     * The records' time to live (in seconds).
     */
    readonly ttl: number;
    readonly type: string;
    /**
     * Whether to use EDNS client subnet data when available (in filter chain).
     */
    readonly useClientSubnet: boolean;
    readonly zone: string;
}
/**
 * Provides details about a NS1 Record. Use this if you would simply like to read
 * information from NS1 into your configurations. For read/write operations, you
 * should use a resource.
 */
export function getRecordOutput(args: GetRecordOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRecordResult> {
    return pulumi.output(args).apply((a: any) => getRecord(a, opts))
}

/**
 * A collection of arguments for invoking getRecord.
 */
export interface GetRecordOutputArgs {
    /**
     * The records' domain.
     */
    domain: pulumi.Input<string>;
    /**
     * The records' RR type.
     */
    type: pulumi.Input<string>;
    /**
     * The zone the record belongs to.
     */
    zone: pulumi.Input<string>;
}
