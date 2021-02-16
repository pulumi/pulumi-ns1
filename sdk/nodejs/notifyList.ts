// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Provides a NS1 Notify List resource. This can be used to create, modify, and delete notify lists.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ns1 from "@pulumi/ns1";
 *
 * const nl = new ns1.NotifyList("nl", {
 *     notifications: [
 *         {
 *             config: {
 *                 url: "http://www.mywebhook.com",
 *             },
 *             type: "webhook",
 *         },
 *         {
 *             config: {
 *                 email: "test@test.com",
 *             },
 *             type: "email",
 *         },
 *     ],
 * });
 * ```
 * ## NS1 Documentation
 *
 * [NotifyList Api Doc](https://ns1.com/api#notification-lists)
 */
export class NotifyList extends pulumi.CustomResource {
    /**
     * Get an existing NotifyList resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NotifyListState, opts?: pulumi.CustomResourceOptions): NotifyList {
        return new NotifyList(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ns1:index/notifyList:NotifyList';

    /**
     * Returns true if the given object is an instance of NotifyList.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NotifyList {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NotifyList.__pulumiType;
    }

    /**
     * The free-form display name for the notify list.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of notifiers. All notifiers in a notification list will receive notifications whenever an event is send to the list (e.g., when a monitoring job fails). Notifiers are documented below.
     */
    public readonly notifications!: pulumi.Output<outputs.NotifyListNotification[] | undefined>;

    /**
     * Create a NotifyList resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: NotifyListArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NotifyListArgs | NotifyListState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NotifyListState | undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["notifications"] = state ? state.notifications : undefined;
        } else {
            const args = argsOrState as NotifyListArgs | undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["notifications"] = args ? args.notifications : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(NotifyList.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NotifyList resources.
 */
export interface NotifyListState {
    /**
     * The free-form display name for the notify list.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of notifiers. All notifiers in a notification list will receive notifications whenever an event is send to the list (e.g., when a monitoring job fails). Notifiers are documented below.
     */
    readonly notifications?: pulumi.Input<pulumi.Input<inputs.NotifyListNotification>[]>;
}

/**
 * The set of arguments for constructing a NotifyList resource.
 */
export interface NotifyListArgs {
    /**
     * The free-form display name for the notify list.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of notifiers. All notifiers in a notification list will receive notifications whenever an event is send to the list (e.g., when a monitoring job fails). Notifiers are documented below.
     */
    readonly notifications?: pulumi.Input<pulumi.Input<inputs.NotifyListNotification>[]>;
}
