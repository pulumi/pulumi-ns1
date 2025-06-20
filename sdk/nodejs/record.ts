// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides a NS1 Record resource. This can be used to create, modify, and delete records.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as external from "@pulumi/external";
 * import * as ns1 from "@pulumi/ns1";
 * import * as std from "@pulumi/std";
 *
 * const example = new ns1.Zone("example", {zone: "terraform.example.io"});
 * const ns1 = new ns1.DataSource("ns1", {
 *     name: "ns1_source",
 *     sourcetype: "nsone_v1",
 * });
 * const foo = new ns1.DataFeed("foo", {
 *     name: "foo_feed",
 *     sourceId: ns1.id,
 *     config: {
 *         label: "foo",
 *     },
 * });
 * const bar = new ns1.DataFeed("bar", {
 *     name: "bar_feed",
 *     sourceId: ns1.id,
 *     config: {
 *         label: "bar",
 *     },
 * });
 * const www = new ns1.Record("www", {
 *     zone: tld.zone,
 *     domain: `www.${tld.zone}`,
 *     type: "CNAME",
 *     ttl: 60,
 *     meta: {
 *         up: "true",
 *     },
 *     regions: [
 *         {
 *             name: "east",
 *             meta: {
 *                 georegion: "US-EAST",
 *             },
 *         },
 *         {
 *             name: "usa",
 *             meta: {
 *                 country: "US",
 *             },
 *         },
 *     ],
 *     answers: [
 *         {
 *             answer: `sub1.${tld.zone}`,
 *             region: "east",
 *             meta: {
 *                 up: pulumi.interpolate`{"feed":"${foo.id}"}`,
 *             },
 *         },
 *         {
 *             answer: `sub2.${tld.zone}`,
 *             meta: {
 *                 up: pulumi.interpolate`{"feed":"${bar.id}"}`,
 *                 connections: "3",
 *             },
 *         },
 *         {
 *             answer: `sub3.${tld.zone}`,
 *             meta: {
 *                 pulsar: JSON.stringify([{
 *                     job_id: "abcdef",
 *                     bias: "*0.55",
 *                     a5m_cutoff: 0.9,
 *                 }]),
 *                 subdivisions: JSON.stringify({
 *                     BR: [
 *                         "SP",
 *                         "SC",
 *                     ],
 *                     DZ: [
 *                         "01",
 *                         "02",
 *                         "03",
 *                     ],
 *                 }),
 *             },
 *         },
 *     ],
 *     filters: [{
 *         filter: "select_first_n",
 *         config: {
 *             N: "1",
 *         },
 *     }],
 * });
 * // Some other non-NS1 provider that returns a zone with a trailing dot and a domain with a leading dot.
 * const baz = new external.index.Source("baz", {
 *     zone: "terraform.example.io.",
 *     domain: ".www.terraform.example.io",
 * });
 * // Basic record showing how to clean a zone or domain field that comes from
 * // another non-NS1 resource. DNS names often end in '.' characters to signify
 * // the root of the DNS tree, but the NS1 provider does not support this.
 * //
 * // In other cases, a domain or zone may be passed in with a preceding dot ('.')
 * // character which would likewise lead the system to fail.
 * const external = new ns1.Record("external", {
 *     zone: std.replace({
 *         text: zone,
 *         search: "/(^\\.)|(\\.$)/",
 *         replace: "",
 *     }).then(invoke => invoke.result),
 *     domain: std.replace({
 *         text: domain,
 *         search: "/(^\\.)|(\\.$)/",
 *         replace: "",
 *     }).then(invoke => invoke.result),
 *     type: "CNAME",
 * });
 * ```
 *
 * ## NS1 Documentation
 *
 * [Record Api Doc](https://ns1.com/api#records)
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import ns1:index/record:Record <name> <zone>/<domain>/<type>`
 * ```
 *
 * So for the example above:
 *
 * ```sh
 * $ pulumi import ns1:index/record:Record www terraform.example.io/www.terraform.example.io/CNAME`
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
    public readonly blockedTags!: pulumi.Output<string[] | undefined>;
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
    public readonly meta!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly overrideAddressRecords!: pulumi.Output<boolean | undefined>;
    public readonly overrideTtl!: pulumi.Output<boolean | undefined>;
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
     * map of tags in the form of `"key" = "value"` where both key and value are strings
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RecordState | undefined;
            resourceInputs["answers"] = state ? state.answers : undefined;
            resourceInputs["blockedTags"] = state ? state.blockedTags : undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["filters"] = state ? state.filters : undefined;
            resourceInputs["link"] = state ? state.link : undefined;
            resourceInputs["meta"] = state ? state.meta : undefined;
            resourceInputs["overrideAddressRecords"] = state ? state.overrideAddressRecords : undefined;
            resourceInputs["overrideTtl"] = state ? state.overrideTtl : undefined;
            resourceInputs["regions"] = state ? state.regions : undefined;
            resourceInputs["shortAnswers"] = state ? state.shortAnswers : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["ttl"] = state ? state.ttl : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["useClientSubnet"] = state ? state.useClientSubnet : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
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
            resourceInputs["answers"] = args ? args.answers : undefined;
            resourceInputs["blockedTags"] = args ? args.blockedTags : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["filters"] = args ? args.filters : undefined;
            resourceInputs["link"] = args ? args.link : undefined;
            resourceInputs["meta"] = args ? args.meta : undefined;
            resourceInputs["overrideAddressRecords"] = args ? args.overrideAddressRecords : undefined;
            resourceInputs["overrideTtl"] = args ? args.overrideTtl : undefined;
            resourceInputs["regions"] = args ? args.regions : undefined;
            resourceInputs["shortAnswers"] = args ? args.shortAnswers : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["useClientSubnet"] = args ? args.useClientSubnet : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Record.__pulumiType, name, resourceInputs, opts);
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
    answers?: pulumi.Input<pulumi.Input<inputs.RecordAnswer>[]>;
    blockedTags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The records' domain. Cannot have leading or trailing
     * dots - see the example above and `FQDN formatting` below.
     */
    domain?: pulumi.Input<string>;
    /**
     * One or more NS1 filters for the record(order matters).
     * Filters are documented below.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.RecordFilter>[]>;
    /**
     * The target record to link to. This means this record is a
     * 'linked' record, and it inherits all properties from its target.
     */
    link?: pulumi.Input<string>;
    meta?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    overrideAddressRecords?: pulumi.Input<boolean>;
    overrideTtl?: pulumi.Input<boolean>;
    /**
     * One or more "regions" for the record. These are really
     * just groupings based on metadata, and are called "Answer Groups" in the NS1 UI,
     * but remain `regions` here for legacy reasons. Regions are
     * documented below. Please note the ordering requirement!
     */
    regions?: pulumi.Input<pulumi.Input<inputs.RecordRegion>[]>;
    /**
     * @deprecated short_answers will be deprecated in a future release. It is suggested to migrate to a regular "answers" block.
     */
    shortAnswers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * map of tags in the form of `"key" = "value"` where both key and value are strings
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The records' time to live (in seconds).
     */
    ttl?: pulumi.Input<number>;
    /**
     * The records' RR type.
     */
    type?: pulumi.Input<string>;
    /**
     * Whether to use EDNS client subnet data when
     * available(in filter chain).
     * * ` meta` - (Optional) meta is supported at the `record` level. Meta
     * is documented below.
     */
    useClientSubnet?: pulumi.Input<boolean>;
    /**
     * The zone the record belongs to. Cannot have leading or
     * trailing dots (".") - see the example above and `FQDN formatting` below.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Record resource.
 */
export interface RecordArgs {
    /**
     * One or more NS1 answers for the records' specified type.
     * Answers are documented below.
     */
    answers?: pulumi.Input<pulumi.Input<inputs.RecordAnswer>[]>;
    blockedTags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The records' domain. Cannot have leading or trailing
     * dots - see the example above and `FQDN formatting` below.
     */
    domain: pulumi.Input<string>;
    /**
     * One or more NS1 filters for the record(order matters).
     * Filters are documented below.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.RecordFilter>[]>;
    /**
     * The target record to link to. This means this record is a
     * 'linked' record, and it inherits all properties from its target.
     */
    link?: pulumi.Input<string>;
    meta?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    overrideAddressRecords?: pulumi.Input<boolean>;
    overrideTtl?: pulumi.Input<boolean>;
    /**
     * One or more "regions" for the record. These are really
     * just groupings based on metadata, and are called "Answer Groups" in the NS1 UI,
     * but remain `regions` here for legacy reasons. Regions are
     * documented below. Please note the ordering requirement!
     */
    regions?: pulumi.Input<pulumi.Input<inputs.RecordRegion>[]>;
    /**
     * @deprecated short_answers will be deprecated in a future release. It is suggested to migrate to a regular "answers" block.
     */
    shortAnswers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * map of tags in the form of `"key" = "value"` where both key and value are strings
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The records' time to live (in seconds).
     */
    ttl?: pulumi.Input<number>;
    /**
     * The records' RR type.
     */
    type: pulumi.Input<string>;
    /**
     * Whether to use EDNS client subnet data when
     * available(in filter chain).
     * * ` meta` - (Optional) meta is supported at the `record` level. Meta
     * is documented below.
     */
    useClientSubnet?: pulumi.Input<boolean>;
    /**
     * The zone the record belongs to. Cannot have leading or
     * trailing dots (".") - see the example above and `FQDN formatting` below.
     */
    zone: pulumi.Input<string>;
}
