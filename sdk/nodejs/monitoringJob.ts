// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Provides a NS1 Monitoring Job resource. This can be used to create, modify, and delete monitoring jobs.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ns1 from "@pulumi/ns1";
 *
 * const uswestMonitor = new ns1.MonitoringJob("uswest_monitor", {
 *     active: true,
 *     config: {
 *         host: "example-elb-uswest.aws.amazon.com",
 *         port: 443,
 *         send: `HEAD / HTTP/1.0
 * 
 * `,
 *         ssl: 1,
 *     },
 *     frequency: 60,
 *     jobType: "tcp",
 *     policy: "quorum",
 *     rapidRecheck: true,
 *     regions: [
 *         "sjc",
 *         "sin",
 *         "lga",
 *     ],
 *     rules: [{
 *         comparison: "contains",
 *         key: "output",
 *         value: "200 OK",
 *     }],
 * });
 * ```
 * ## NS1 Documentation
 *
 * [MonitoringJob Api Doc](https://ns1.com/api#monitoring-jobs)
 */
export class MonitoringJob extends pulumi.CustomResource {
    /**
     * Get an existing MonitoringJob resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MonitoringJobState, opts?: pulumi.CustomResourceOptions): MonitoringJob {
        return new MonitoringJob(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ns1:index/monitoringJob:MonitoringJob';

    /**
     * Returns true if the given object is an instance of MonitoringJob.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MonitoringJob {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MonitoringJob.__pulumiType;
    }

    /**
     * Indicates if the job is active or temporarily disabled.
     */
    public readonly active!: pulumi.Output<boolean | undefined>;
    /**
     * A configuration dictionary with keys and values depending on the job_type. Configuration details for each jobType are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
     */
    public readonly config!: pulumi.Output<{[key: string]: any}>;
    /**
     * The frequency, in seconds, at which to run the monitoring job in each region.
     */
    public readonly frequency!: pulumi.Output<number>;
    /**
     * The type of monitoring job to be run. Refer to the NS1 API documentation (https://ns1.com/api#monitoring-jobs) for supported values which include ping, tcp, dns, http.
     */
    public readonly jobType!: pulumi.Output<string>;
    /**
     * The free-form display name for the monitoring job.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Freeform notes to be included in any notifications about this job.
     */
    public readonly notes!: pulumi.Output<string | undefined>;
    /**
     * The time in seconds after a failure to wait before sending a notification.
     */
    public readonly notifyDelay!: pulumi.Output<number | undefined>;
    /**
     * If true, a notification is sent when a job returns to an "up" state.
     */
    public readonly notifyFailback!: pulumi.Output<boolean | undefined>;
    public readonly notifyList!: pulumi.Output<string | undefined>;
    /**
     * If true, notifications are sent for any regional failure (and failback if desired), in addition to global state notifications.
     */
    public readonly notifyRegional!: pulumi.Output<boolean | undefined>;
    /**
     * The time in seconds between repeat notifications of a failed job.
     */
    public readonly notifyRepeat!: pulumi.Output<number | undefined>;
    /**
     * The policy for determining the monitor's global status
     * based on the status of the job in all regions. See NS1 API docs for supported values.
     */
    public readonly policy!: pulumi.Output<string | undefined>;
    /**
     * If true, on any apparent state change, the job is quickly re-run after one second to confirm the state change before notification.
     */
    public readonly rapidRecheck!: pulumi.Output<boolean | undefined>;
    /**
     * The list of region codes in which to run the monitoring
     * job. See NS1 API docs for supported values.
     */
    public readonly regions!: pulumi.Output<string[]>;
    /**
     * A list of rules for determining failure conditions. Each rule acts on one of the outputs from the monitoring job. You must specify key (the output key); comparison (a comparison to perform on the the output); and value (the value to compare to). For example, {"key":"rtt", "comparison":"<", "value":100} is a rule requiring the rtt from a job to be under 100ms, or the job will be marked failed. Available output keys, comparators, and value types are are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
     */
    public readonly rules!: pulumi.Output<outputs.MonitoringJobRule[] | undefined>;

    /**
     * Create a MonitoringJob resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MonitoringJobArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MonitoringJobArgs | MonitoringJobState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MonitoringJobState | undefined;
            inputs["active"] = state ? state.active : undefined;
            inputs["config"] = state ? state.config : undefined;
            inputs["frequency"] = state ? state.frequency : undefined;
            inputs["jobType"] = state ? state.jobType : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["notes"] = state ? state.notes : undefined;
            inputs["notifyDelay"] = state ? state.notifyDelay : undefined;
            inputs["notifyFailback"] = state ? state.notifyFailback : undefined;
            inputs["notifyList"] = state ? state.notifyList : undefined;
            inputs["notifyRegional"] = state ? state.notifyRegional : undefined;
            inputs["notifyRepeat"] = state ? state.notifyRepeat : undefined;
            inputs["policy"] = state ? state.policy : undefined;
            inputs["rapidRecheck"] = state ? state.rapidRecheck : undefined;
            inputs["regions"] = state ? state.regions : undefined;
            inputs["rules"] = state ? state.rules : undefined;
        } else {
            const args = argsOrState as MonitoringJobArgs | undefined;
            if ((!args || args.config === undefined) && !opts.urn) {
                throw new Error("Missing required property 'config'");
            }
            if ((!args || args.frequency === undefined) && !opts.urn) {
                throw new Error("Missing required property 'frequency'");
            }
            if ((!args || args.jobType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'jobType'");
            }
            if ((!args || args.regions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'regions'");
            }
            inputs["active"] = args ? args.active : undefined;
            inputs["config"] = args ? args.config : undefined;
            inputs["frequency"] = args ? args.frequency : undefined;
            inputs["jobType"] = args ? args.jobType : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["notes"] = args ? args.notes : undefined;
            inputs["notifyDelay"] = args ? args.notifyDelay : undefined;
            inputs["notifyFailback"] = args ? args.notifyFailback : undefined;
            inputs["notifyList"] = args ? args.notifyList : undefined;
            inputs["notifyRegional"] = args ? args.notifyRegional : undefined;
            inputs["notifyRepeat"] = args ? args.notifyRepeat : undefined;
            inputs["policy"] = args ? args.policy : undefined;
            inputs["rapidRecheck"] = args ? args.rapidRecheck : undefined;
            inputs["regions"] = args ? args.regions : undefined;
            inputs["rules"] = args ? args.rules : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(MonitoringJob.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MonitoringJob resources.
 */
export interface MonitoringJobState {
    /**
     * Indicates if the job is active or temporarily disabled.
     */
    readonly active?: pulumi.Input<boolean>;
    /**
     * A configuration dictionary with keys and values depending on the job_type. Configuration details for each jobType are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
     */
    readonly config?: pulumi.Input<{[key: string]: any}>;
    /**
     * The frequency, in seconds, at which to run the monitoring job in each region.
     */
    readonly frequency?: pulumi.Input<number>;
    /**
     * The type of monitoring job to be run. Refer to the NS1 API documentation (https://ns1.com/api#monitoring-jobs) for supported values which include ping, tcp, dns, http.
     */
    readonly jobType?: pulumi.Input<string>;
    /**
     * The free-form display name for the monitoring job.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Freeform notes to be included in any notifications about this job.
     */
    readonly notes?: pulumi.Input<string>;
    /**
     * The time in seconds after a failure to wait before sending a notification.
     */
    readonly notifyDelay?: pulumi.Input<number>;
    /**
     * If true, a notification is sent when a job returns to an "up" state.
     */
    readonly notifyFailback?: pulumi.Input<boolean>;
    readonly notifyList?: pulumi.Input<string>;
    /**
     * If true, notifications are sent for any regional failure (and failback if desired), in addition to global state notifications.
     */
    readonly notifyRegional?: pulumi.Input<boolean>;
    /**
     * The time in seconds between repeat notifications of a failed job.
     */
    readonly notifyRepeat?: pulumi.Input<number>;
    /**
     * The policy for determining the monitor's global status
     * based on the status of the job in all regions. See NS1 API docs for supported values.
     */
    readonly policy?: pulumi.Input<string>;
    /**
     * If true, on any apparent state change, the job is quickly re-run after one second to confirm the state change before notification.
     */
    readonly rapidRecheck?: pulumi.Input<boolean>;
    /**
     * The list of region codes in which to run the monitoring
     * job. See NS1 API docs for supported values.
     */
    readonly regions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of rules for determining failure conditions. Each rule acts on one of the outputs from the monitoring job. You must specify key (the output key); comparison (a comparison to perform on the the output); and value (the value to compare to). For example, {"key":"rtt", "comparison":"<", "value":100} is a rule requiring the rtt from a job to be under 100ms, or the job will be marked failed. Available output keys, comparators, and value types are are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
     */
    readonly rules?: pulumi.Input<pulumi.Input<inputs.MonitoringJobRule>[]>;
}

/**
 * The set of arguments for constructing a MonitoringJob resource.
 */
export interface MonitoringJobArgs {
    /**
     * Indicates if the job is active or temporarily disabled.
     */
    readonly active?: pulumi.Input<boolean>;
    /**
     * A configuration dictionary with keys and values depending on the job_type. Configuration details for each jobType are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
     */
    readonly config: pulumi.Input<{[key: string]: any}>;
    /**
     * The frequency, in seconds, at which to run the monitoring job in each region.
     */
    readonly frequency: pulumi.Input<number>;
    /**
     * The type of monitoring job to be run. Refer to the NS1 API documentation (https://ns1.com/api#monitoring-jobs) for supported values which include ping, tcp, dns, http.
     */
    readonly jobType: pulumi.Input<string>;
    /**
     * The free-form display name for the monitoring job.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Freeform notes to be included in any notifications about this job.
     */
    readonly notes?: pulumi.Input<string>;
    /**
     * The time in seconds after a failure to wait before sending a notification.
     */
    readonly notifyDelay?: pulumi.Input<number>;
    /**
     * If true, a notification is sent when a job returns to an "up" state.
     */
    readonly notifyFailback?: pulumi.Input<boolean>;
    readonly notifyList?: pulumi.Input<string>;
    /**
     * If true, notifications are sent for any regional failure (and failback if desired), in addition to global state notifications.
     */
    readonly notifyRegional?: pulumi.Input<boolean>;
    /**
     * The time in seconds between repeat notifications of a failed job.
     */
    readonly notifyRepeat?: pulumi.Input<number>;
    /**
     * The policy for determining the monitor's global status
     * based on the status of the job in all regions. See NS1 API docs for supported values.
     */
    readonly policy?: pulumi.Input<string>;
    /**
     * If true, on any apparent state change, the job is quickly re-run after one second to confirm the state change before notification.
     */
    readonly rapidRecheck?: pulumi.Input<boolean>;
    /**
     * The list of region codes in which to run the monitoring
     * job. See NS1 API docs for supported values.
     */
    readonly regions: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of rules for determining failure conditions. Each rule acts on one of the outputs from the monitoring job. You must specify key (the output key); comparison (a comparison to perform on the the output); and value (the value to compare to). For example, {"key":"rtt", "comparison":"<", "value":100} is a rule requiring the rtt from a job to be under 100ms, or the job will be marked failed. Available output keys, comparators, and value types are are found by submitting a GET request to https://api.nsone.net/v1/monitoring/jobtypes.
     */
    readonly rules?: pulumi.Input<pulumi.Input<inputs.MonitoringJobRule>[]>;
}
