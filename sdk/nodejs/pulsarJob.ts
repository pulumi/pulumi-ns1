// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class PulsarJob extends pulumi.CustomResource {
    /**
     * Get an existing PulsarJob resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PulsarJobState, opts?: pulumi.CustomResourceOptions): PulsarJob {
        return new PulsarJob(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ns1:index/pulsarJob:PulsarJob';

    /**
     * Returns true if the given object is an instance of PulsarJob.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PulsarJob {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PulsarJob.__pulumiType;
    }

    public readonly active!: pulumi.Output<boolean | undefined>;
    public readonly appId!: pulumi.Output<string>;
    public readonly blendMetricWeights!: pulumi.Output<outputs.PulsarJobBlendMetricWeights | undefined>;
    public /*out*/ readonly community!: pulumi.Output<boolean>;
    public readonly config!: pulumi.Output<outputs.PulsarJobConfig | undefined>;
    public /*out*/ readonly customer!: pulumi.Output<number>;
    public /*out*/ readonly jobId!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly shared!: pulumi.Output<boolean | undefined>;
    public readonly typeId!: pulumi.Output<string>;
    public readonly weights!: pulumi.Output<outputs.PulsarJobWeight[] | undefined>;

    /**
     * Create a PulsarJob resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PulsarJobArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PulsarJobArgs | PulsarJobState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PulsarJobState | undefined;
            resourceInputs["active"] = state ? state.active : undefined;
            resourceInputs["appId"] = state ? state.appId : undefined;
            resourceInputs["blendMetricWeights"] = state ? state.blendMetricWeights : undefined;
            resourceInputs["community"] = state ? state.community : undefined;
            resourceInputs["config"] = state ? state.config : undefined;
            resourceInputs["customer"] = state ? state.customer : undefined;
            resourceInputs["jobId"] = state ? state.jobId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["shared"] = state ? state.shared : undefined;
            resourceInputs["typeId"] = state ? state.typeId : undefined;
            resourceInputs["weights"] = state ? state.weights : undefined;
        } else {
            const args = argsOrState as PulsarJobArgs | undefined;
            if ((!args || args.appId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appId'");
            }
            if ((!args || args.typeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'typeId'");
            }
            resourceInputs["active"] = args ? args.active : undefined;
            resourceInputs["appId"] = args ? args.appId : undefined;
            resourceInputs["blendMetricWeights"] = args ? args.blendMetricWeights : undefined;
            resourceInputs["config"] = args ? args.config : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["shared"] = args ? args.shared : undefined;
            resourceInputs["typeId"] = args ? args.typeId : undefined;
            resourceInputs["weights"] = args ? args.weights : undefined;
            resourceInputs["community"] = undefined /*out*/;
            resourceInputs["customer"] = undefined /*out*/;
            resourceInputs["jobId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PulsarJob.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PulsarJob resources.
 */
export interface PulsarJobState {
    active?: pulumi.Input<boolean>;
    appId?: pulumi.Input<string>;
    blendMetricWeights?: pulumi.Input<inputs.PulsarJobBlendMetricWeights>;
    community?: pulumi.Input<boolean>;
    config?: pulumi.Input<inputs.PulsarJobConfig>;
    customer?: pulumi.Input<number>;
    jobId?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    shared?: pulumi.Input<boolean>;
    typeId?: pulumi.Input<string>;
    weights?: pulumi.Input<pulumi.Input<inputs.PulsarJobWeight>[]>;
}

/**
 * The set of arguments for constructing a PulsarJob resource.
 */
export interface PulsarJobArgs {
    active?: pulumi.Input<boolean>;
    appId: pulumi.Input<string>;
    blendMetricWeights?: pulumi.Input<inputs.PulsarJobBlendMetricWeights>;
    config?: pulumi.Input<inputs.PulsarJobConfig>;
    name?: pulumi.Input<string>;
    shared?: pulumi.Input<boolean>;
    typeId: pulumi.Input<string>;
    weights?: pulumi.Input<pulumi.Input<inputs.PulsarJobWeight>[]>;
}
