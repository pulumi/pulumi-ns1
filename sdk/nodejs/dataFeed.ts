// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a NS1 Data Feed resource. This can be used to create, modify, and delete data feeds.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ns1 from "@pulumi/ns1";
 *
 * const example = new ns1.DataSource("example", {
 *     sourcetype: "nsone_v1",
 * });
 * const uswestFeed = new ns1.DataFeed("uswest_feed", {
 *     config: {
 *         label: "uswest",
 *     },
 *     sourceId: example.id,
 * });
 * const useastFeed = new ns1.DataFeed("useast_feed", {
 *     config: {
 *         label: "useast",
 *     },
 *     sourceId: example.id,
 * });
 * ```
 * ## NS1 Documentation
 *
 * [Datafeed Api Doc](https://ns1.com/api#data-feeds)
 */
export class DataFeed extends pulumi.CustomResource {
    /**
     * Get an existing DataFeed resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DataFeedState, opts?: pulumi.CustomResourceOptions): DataFeed {
        return new DataFeed(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ns1:index/dataFeed:DataFeed';

    /**
     * Returns true if the given object is an instance of DataFeed.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataFeed {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataFeed.__pulumiType;
    }

    /**
     * The feeds configuration matching the specification in
     * `feedConfig` from /data/sourcetypes.
     */
    public readonly config!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The free form name of the data feed.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The data source id that this feed is connected to.
     */
    public readonly sourceId!: pulumi.Output<string>;

    /**
     * Create a DataFeed resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataFeedArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DataFeedArgs | DataFeedState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DataFeedState | undefined;
            inputs["config"] = state ? state.config : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["sourceId"] = state ? state.sourceId : undefined;
        } else {
            const args = argsOrState as DataFeedArgs | undefined;
            if ((!args || args.sourceId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'sourceId'");
            }
            inputs["config"] = args ? args.config : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["sourceId"] = args ? args.sourceId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DataFeed.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DataFeed resources.
 */
export interface DataFeedState {
    /**
     * The feeds configuration matching the specification in
     * `feedConfig` from /data/sourcetypes.
     */
    readonly config?: pulumi.Input<{[key: string]: any}>;
    /**
     * The free form name of the data feed.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The data source id that this feed is connected to.
     */
    readonly sourceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DataFeed resource.
 */
export interface DataFeedArgs {
    /**
     * The feeds configuration matching the specification in
     * `feedConfig` from /data/sourcetypes.
     */
    readonly config?: pulumi.Input<{[key: string]: any}>;
    /**
     * The free form name of the data feed.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The data source id that this feed is connected to.
     */
    readonly sourceId: pulumi.Input<string>;
}
