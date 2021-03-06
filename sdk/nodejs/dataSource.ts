// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a NS1 Data Source resource. This can be used to create, modify, and delete data sources.
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
 * ```
 * ## NS1 Documentation
 *
 * [Datasource Api Doc](https://ns1.com/api#data-sources)
 */
export class DataSource extends pulumi.CustomResource {
    /**
     * Get an existing DataSource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DataSourceState, opts?: pulumi.CustomResourceOptions): DataSource {
        return new DataSource(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ns1:index/dataSource:DataSource';

    /**
     * Returns true if the given object is an instance of DataSource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataSource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataSource.__pulumiType;
    }

    /**
     * The data source configuration, determined by its type,
     * matching the specification in `config` from /data/sourcetypes.
     */
    public readonly config!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The free form name of the data source.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The data sources type, listed in API endpoint https://api.nsone.net/v1/data/sourcetypes.
     */
    public readonly sourcetype!: pulumi.Output<string>;

    /**
     * Create a DataSource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataSourceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DataSourceArgs | DataSourceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DataSourceState | undefined;
            inputs["config"] = state ? state.config : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["sourcetype"] = state ? state.sourcetype : undefined;
        } else {
            const args = argsOrState as DataSourceArgs | undefined;
            if ((!args || args.sourcetype === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourcetype'");
            }
            inputs["config"] = args ? args.config : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["sourcetype"] = args ? args.sourcetype : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DataSource.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DataSource resources.
 */
export interface DataSourceState {
    /**
     * The data source configuration, determined by its type,
     * matching the specification in `config` from /data/sourcetypes.
     */
    readonly config?: pulumi.Input<{[key: string]: any}>;
    /**
     * The free form name of the data source.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The data sources type, listed in API endpoint https://api.nsone.net/v1/data/sourcetypes.
     */
    readonly sourcetype?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DataSource resource.
 */
export interface DataSourceArgs {
    /**
     * The data source configuration, determined by its type,
     * matching the specification in `config` from /data/sourcetypes.
     */
    readonly config?: pulumi.Input<{[key: string]: any}>;
    /**
     * The free form name of the data source.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The data sources type, listed in API endpoint https://api.nsone.net/v1/data/sourcetypes.
     */
    readonly sourcetype: pulumi.Input<string>;
}
