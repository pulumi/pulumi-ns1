// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a NS1 TSIG Key resource. This can be used to create, modify, and delete TSIG keys.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ns1 from "@pulumi/ns1";
 *
 * const example = new ns1.Tsigkey("example", {
 *     name: "ExampleTsigKey",
 *     algorithm: "hmac-sha256",
 *     secret: "Ok1qR5IW1ajVka5cHPEJQIXfLyx5V3PSkFBROAzOn21JumDq6nIpoj6H8rfj5Uo+Ok55ZWQ0Wgrf302fDscHLA==",
 * });
 * ```
 * ## NS1 Documentation
 *
 * [TSIG Keys Api Doc](https://ns1.com/api/#tsig)
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import ns1:index/tsigkey:Tsigkey importTest <name>`
 * ```
 */
export class Tsigkey extends pulumi.CustomResource {
    /**
     * Get an existing Tsigkey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TsigkeyState, opts?: pulumi.CustomResourceOptions): Tsigkey {
        return new Tsigkey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ns1:index/tsigkey:Tsigkey';

    /**
     * Returns true if the given object is an instance of Tsigkey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Tsigkey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Tsigkey.__pulumiType;
    }

    /**
     * The algorithm used to hash the TSIG key's secret.
     */
    public readonly algorithm!: pulumi.Output<string>;
    /**
     * The free form name of the tsigkey.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The key's secret to be hashed.
     */
    public readonly secret!: pulumi.Output<string>;

    /**
     * Create a Tsigkey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TsigkeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TsigkeyArgs | TsigkeyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TsigkeyState | undefined;
            resourceInputs["algorithm"] = state ? state.algorithm : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["secret"] = state ? state.secret : undefined;
        } else {
            const args = argsOrState as TsigkeyArgs | undefined;
            if ((!args || args.algorithm === undefined) && !opts.urn) {
                throw new Error("Missing required property 'algorithm'");
            }
            if ((!args || args.secret === undefined) && !opts.urn) {
                throw new Error("Missing required property 'secret'");
            }
            resourceInputs["algorithm"] = args ? args.algorithm : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["secret"] = args ? args.secret : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Tsigkey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Tsigkey resources.
 */
export interface TsigkeyState {
    /**
     * The algorithm used to hash the TSIG key's secret.
     */
    algorithm?: pulumi.Input<string>;
    /**
     * The free form name of the tsigkey.
     */
    name?: pulumi.Input<string>;
    /**
     * The key's secret to be hashed.
     */
    secret?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Tsigkey resource.
 */
export interface TsigkeyArgs {
    /**
     * The algorithm used to hash the TSIG key's secret.
     */
    algorithm: pulumi.Input<string>;
    /**
     * The free form name of the tsigkey.
     */
    name?: pulumi.Input<string>;
    /**
     * The key's secret to be hashed.
     */
    secret: pulumi.Input<string>;
}
