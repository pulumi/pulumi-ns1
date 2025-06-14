// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides a NS1 Pulsar application resource. This can be used to create, modify, and delete applications.
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import ns1:index/application:Application `ns1_application` 
 * ```
 *
 * So for the example above:
 *
 * ```sh
 * $ pulumi import ns1:index/application:Application example terraform.example.io`
 * ```
 */
export class Application extends pulumi.CustomResource {
    /**
     * Get an existing Application resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApplicationState, opts?: pulumi.CustomResourceOptions): Application {
        return new Application(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ns1:index/application:Application';

    /**
     * Returns true if the given object is an instance of Application.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Application {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Application.__pulumiType;
    }

    /**
     * Indicates whether or not this application is currently active and usable for traffic
     * steering.
     */
    public readonly active!: pulumi.Output<boolean | undefined>;
    /**
     * The amount of time (in milliseconds) the browser should wait before running
     * measurements.
     */
    public readonly browserWaitMillis!: pulumi.Output<number | undefined>;
    /**
     * Default job configuration. If a field is present here and not on a specific job
     * associated with this application, the default value specified here is used..
     */
    public readonly defaultConfig!: pulumi.Output<outputs.ApplicationDefaultConfig>;
    /**
     * Number of jobs to measure per user impression.
     */
    public readonly jobsPerTransaction!: pulumi.Output<number | undefined>;
    /**
     * Descriptive name for this Pulsar app.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a Application resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ApplicationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApplicationArgs | ApplicationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApplicationState | undefined;
            resourceInputs["active"] = state ? state.active : undefined;
            resourceInputs["browserWaitMillis"] = state ? state.browserWaitMillis : undefined;
            resourceInputs["defaultConfig"] = state ? state.defaultConfig : undefined;
            resourceInputs["jobsPerTransaction"] = state ? state.jobsPerTransaction : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as ApplicationArgs | undefined;
            resourceInputs["active"] = args ? args.active : undefined;
            resourceInputs["browserWaitMillis"] = args ? args.browserWaitMillis : undefined;
            resourceInputs["defaultConfig"] = args ? args.defaultConfig : undefined;
            resourceInputs["jobsPerTransaction"] = args ? args.jobsPerTransaction : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Application.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Application resources.
 */
export interface ApplicationState {
    /**
     * Indicates whether or not this application is currently active and usable for traffic
     * steering.
     */
    active?: pulumi.Input<boolean>;
    /**
     * The amount of time (in milliseconds) the browser should wait before running
     * measurements.
     */
    browserWaitMillis?: pulumi.Input<number>;
    /**
     * Default job configuration. If a field is present here and not on a specific job
     * associated with this application, the default value specified here is used..
     */
    defaultConfig?: pulumi.Input<inputs.ApplicationDefaultConfig>;
    /**
     * Number of jobs to measure per user impression.
     */
    jobsPerTransaction?: pulumi.Input<number>;
    /**
     * Descriptive name for this Pulsar app.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Application resource.
 */
export interface ApplicationArgs {
    /**
     * Indicates whether or not this application is currently active and usable for traffic
     * steering.
     */
    active?: pulumi.Input<boolean>;
    /**
     * The amount of time (in milliseconds) the browser should wait before running
     * measurements.
     */
    browserWaitMillis?: pulumi.Input<number>;
    /**
     * Default job configuration. If a field is present here and not on a specific job
     * associated with this application, the default value specified here is used..
     */
    defaultConfig?: pulumi.Input<inputs.ApplicationDefaultConfig>;
    /**
     * Number of jobs to measure per user impression.
     */
    jobsPerTransaction?: pulumi.Input<number>;
    /**
     * Descriptive name for this Pulsar app.
     */
    name?: pulumi.Input<string>;
}
