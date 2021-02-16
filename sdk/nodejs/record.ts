// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Provides a NS1 Record resource. This can be used to create, modify, and delete records.
 *
 * ## NS1 Documentation
 *
 * [Record Api Doc](https://ns1.com/api#records)
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import ns1:index/record:Record <name> <zone>/<domain>/<type>`
 * ```
 *
 *  So for the example above
 *
 * ```sh
 *  $ pulumi import ns1:index/record:Record www terraform.example.io/www.terraform.example.io/CNAME`
 * ```
 */
export class Record extends pulumi.CustomResource {
    /**
     * Get an existing Record resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RecordState, opts?: pulumi.CustomResourceOptions): Record {
        return new Record(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ns1:index/record:Record';

    /**
     * Returns true if the given object is an instance of Record.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Record {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Record.__pulumiType;
    }

    /**
     * One or more NS1 answers for the records' specified type.
     * Answers are documented below.
     */
    public readonly answers!: pulumi.Output<outputs.RecordAnswer[] | undefined>;
    /**
     * The records' domain. Cannot have leading or trailing
     * dots - see the example above and `FQDN formatting` below.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * One or more NS1 filters for the record(order matters).
     * Filters are documented below.
     */
    public readonly filters!: pulumi.Output<outputs.RecordFilter[] | undefined>;
    /**
     * The target record to link to. This means this record is a
     * 'linked' record, and it inherits all properties from its target.
     */
    public readonly link!: pulumi.Output<string | undefined>;
    public readonly meta!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * One or more "regions" for the record. These are really
     * just groupings based on metadata, and are called "Answer Groups" in the NS1 UI,
     * but remain `regions` here for legacy reasons. Regions are
     * documented below. Please note the ordering requirement!
     */
    public readonly regions!: pulumi.Output<outputs.RecordRegion[] | undefined>;
    /**
     * @deprecated short_answers will be deprecated in a future release. It is suggested to migrate to a regular "answers" block.
     */
    public readonly shortAnswers!: pulumi.Output<string[] | undefined>;
    /**
     * The records' time to live (in seconds).
     */
    public readonly ttl!: pulumi.Output<number>;
    /**
     * The records' RR type.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Whether to use EDNS client subnet data when
     * available(in filter chain).
     * * ` meta` - (Optional) meta is supported at the `record` level. Meta
     * is documented below.
     */
    public readonly useClientSubnet!: pulumi.Output<boolean | undefined>;
    /**
     * The zone the record belongs to. Cannot have leading or
     * trailing dots (".") - see the example above and `FQDN formatting` below.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a Record resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RecordArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RecordArgs | RecordState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RecordState | undefined;
            inputs["answers"] = state ? state.answers : undefined;
            inputs["domain"] = state ? state.domain : undefined;
            inputs["filters"] = state ? state.filters : undefined;
            inputs["link"] = state ? state.link : undefined;
            inputs["meta"] = state ? state.meta : undefined;
            inputs["regions"] = state ? state.regions : undefined;
            inputs["shortAnswers"] = state ? state.shortAnswers : undefined;
            inputs["ttl"] = state ? state.ttl : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["useClientSubnet"] = state ? state.useClientSubnet : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as RecordArgs | undefined;
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            if ((!args || args.zone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zone'");
            }
            inputs["answers"] = args ? args.answers : undefined;
            inputs["domain"] = args ? args.domain : undefined;
            inputs["filters"] = args ? args.filters : undefined;
            inputs["link"] = args ? args.link : undefined;
            inputs["meta"] = args ? args.meta : undefined;
            inputs["regions"] = args ? args.regions : undefined;
            inputs["shortAnswers"] = args ? args.shortAnswers : undefined;
            inputs["ttl"] = args ? args.ttl : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["useClientSubnet"] = args ? args.useClientSubnet : undefined;
            inputs["zone"] = args ? args.zone : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Record.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Record resources.
 */
export interface RecordState {
    /**
     * One or more NS1 answers for the records' specified type.
     * Answers are documented below.
     */
    readonly answers?: pulumi.Input<pulumi.Input<inputs.RecordAnswer>[]>;
    /**
     * The records' domain. Cannot have leading or trailing
     * dots - see the example above and `FQDN formatting` below.
     */
    readonly domain?: pulumi.Input<string>;
    /**
     * One or more NS1 filters for the record(order matters).
     * Filters are documented below.
     */
    readonly filters?: pulumi.Input<pulumi.Input<inputs.RecordFilter>[]>;
    /**
     * The target record to link to. This means this record is a
     * 'linked' record, and it inherits all properties from its target.
     */
    readonly link?: pulumi.Input<string>;
    readonly meta?: pulumi.Input<{[key: string]: any}>;
    /**
     * One or more "regions" for the record. These are really
     * just groupings based on metadata, and are called "Answer Groups" in the NS1 UI,
     * but remain `regions` here for legacy reasons. Regions are
     * documented below. Please note the ordering requirement!
     */
    readonly regions?: pulumi.Input<pulumi.Input<inputs.RecordRegion>[]>;
    /**
     * @deprecated short_answers will be deprecated in a future release. It is suggested to migrate to a regular "answers" block.
     */
    readonly shortAnswers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The records' time to live (in seconds).
     */
    readonly ttl?: pulumi.Input<number>;
    /**
     * The records' RR type.
     */
    readonly type?: pulumi.Input<string>;
    /**
     * Whether to use EDNS client subnet data when
     * available(in filter chain).
     * * ` meta` - (Optional) meta is supported at the `record` level. Meta
     * is documented below.
     */
    readonly useClientSubnet?: pulumi.Input<boolean>;
    /**
     * The zone the record belongs to. Cannot have leading or
     * trailing dots (".") - see the example above and `FQDN formatting` below.
     */
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Record resource.
 */
export interface RecordArgs {
    /**
     * One or more NS1 answers for the records' specified type.
     * Answers are documented below.
     */
    readonly answers?: pulumi.Input<pulumi.Input<inputs.RecordAnswer>[]>;
    /**
     * The records' domain. Cannot have leading or trailing
     * dots - see the example above and `FQDN formatting` below.
     */
    readonly domain: pulumi.Input<string>;
    /**
     * One or more NS1 filters for the record(order matters).
     * Filters are documented below.
     */
    readonly filters?: pulumi.Input<pulumi.Input<inputs.RecordFilter>[]>;
    /**
     * The target record to link to. This means this record is a
     * 'linked' record, and it inherits all properties from its target.
     */
    readonly link?: pulumi.Input<string>;
    readonly meta?: pulumi.Input<{[key: string]: any}>;
    /**
     * One or more "regions" for the record. These are really
     * just groupings based on metadata, and are called "Answer Groups" in the NS1 UI,
     * but remain `regions` here for legacy reasons. Regions are
     * documented below. Please note the ordering requirement!
     */
    readonly regions?: pulumi.Input<pulumi.Input<inputs.RecordRegion>[]>;
    /**
     * @deprecated short_answers will be deprecated in a future release. It is suggested to migrate to a regular "answers" block.
     */
    readonly shortAnswers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The records' time to live (in seconds).
     */
    readonly ttl?: pulumi.Input<number>;
    /**
     * The records' RR type.
     */
    readonly type: pulumi.Input<string>;
    /**
     * Whether to use EDNS client subnet data when
     * available(in filter chain).
     * * ` meta` - (Optional) meta is supported at the `record` level. Meta
     * is documented below.
     */
    readonly useClientSubnet?: pulumi.Input<boolean>;
    /**
     * The zone the record belongs to. Cannot have leading or
     * trailing dots (".") - see the example above and `FQDN formatting` below.
     */
    readonly zone: pulumi.Input<string>;
}
