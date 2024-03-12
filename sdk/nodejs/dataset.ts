// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class Dataset extends pulumi.CustomResource {
    /**
     * Get an existing Dataset resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatasetState, opts?: pulumi.CustomResourceOptions): Dataset {
        return new Dataset(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ns1:index/dataset:Dataset';

    /**
     * Returns true if the given object is an instance of Dataset.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Dataset {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Dataset.__pulumiType;
    }

    public readonly datatype!: pulumi.Output<outputs.DatasetDatatype>;
    public readonly exportType!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly recipientEmails!: pulumi.Output<string[] | undefined>;
    public readonly repeat!: pulumi.Output<outputs.DatasetRepeat | undefined>;
    public /*out*/ readonly reports!: pulumi.Output<outputs.DatasetReport[]>;
    public readonly timeframe!: pulumi.Output<outputs.DatasetTimeframe>;

    /**
     * Create a Dataset resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatasetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatasetArgs | DatasetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatasetState | undefined;
            resourceInputs["datatype"] = state ? state.datatype : undefined;
            resourceInputs["exportType"] = state ? state.exportType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["recipientEmails"] = state ? state.recipientEmails : undefined;
            resourceInputs["repeat"] = state ? state.repeat : undefined;
            resourceInputs["reports"] = state ? state.reports : undefined;
            resourceInputs["timeframe"] = state ? state.timeframe : undefined;
        } else {
            const args = argsOrState as DatasetArgs | undefined;
            if ((!args || args.datatype === undefined) && !opts.urn) {
                throw new Error("Missing required property 'datatype'");
            }
            if ((!args || args.exportType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'exportType'");
            }
            if ((!args || args.timeframe === undefined) && !opts.urn) {
                throw new Error("Missing required property 'timeframe'");
            }
            resourceInputs["datatype"] = args ? args.datatype : undefined;
            resourceInputs["exportType"] = args ? args.exportType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["recipientEmails"] = args ? args.recipientEmails : undefined;
            resourceInputs["repeat"] = args ? args.repeat : undefined;
            resourceInputs["timeframe"] = args ? args.timeframe : undefined;
            resourceInputs["reports"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Dataset.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Dataset resources.
 */
export interface DatasetState {
    datatype?: pulumi.Input<inputs.DatasetDatatype>;
    exportType?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    recipientEmails?: pulumi.Input<pulumi.Input<string>[]>;
    repeat?: pulumi.Input<inputs.DatasetRepeat>;
    reports?: pulumi.Input<pulumi.Input<inputs.DatasetReport>[]>;
    timeframe?: pulumi.Input<inputs.DatasetTimeframe>;
}

/**
 * The set of arguments for constructing a Dataset resource.
 */
export interface DatasetArgs {
    datatype: pulumi.Input<inputs.DatasetDatatype>;
    exportType: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    recipientEmails?: pulumi.Input<pulumi.Input<string>[]>;
    repeat?: pulumi.Input<inputs.DatasetRepeat>;
    timeframe: pulumi.Input<inputs.DatasetTimeframe>;
}
