// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a NS1 Global IP Whitelist resource.
 *
 * This can be used to create, modify, and delete Global IP Whitelists.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ns1 from "@pulumi/ns1";
 *
 * const example = new ns1.AccountWhitelist("example", {
 *     name: "Example Whitelist",
 *     values: [
 *         "1.1.1.1",
 *         "2.2.2.2",
 *     ],
 * });
 * ```
 *
 * > You current source IP must be present in one of the whitelists to prevent locking yourself out.
 *
 * ## NS1 Documentation
 *
 * [Global IP Whitelist Doc](https://ns1.com/api?docId=2282)
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import ns1:index/accountWhitelist:AccountWhitelist example <ID>`
 * ```
 */
export class AccountWhitelist extends pulumi.CustomResource {
    /**
     * Get an existing AccountWhitelist resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AccountWhitelistState, opts?: pulumi.CustomResourceOptions): AccountWhitelist {
        return new AccountWhitelist(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ns1:index/accountWhitelist:AccountWhitelist';

    /**
     * Returns true if the given object is an instance of AccountWhitelist.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AccountWhitelist {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AccountWhitelist.__pulumiType;
    }

    /**
     * The free form name of the whitelist.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Array of IP addresses/networks from which to allow access.
     */
    public readonly values!: pulumi.Output<string[]>;

    /**
     * Create a AccountWhitelist resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AccountWhitelistArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AccountWhitelistArgs | AccountWhitelistState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AccountWhitelistState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["values"] = state ? state.values : undefined;
        } else {
            const args = argsOrState as AccountWhitelistArgs | undefined;
            if ((!args || args.values === undefined) && !opts.urn) {
                throw new Error("Missing required property 'values'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["values"] = args ? args.values : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AccountWhitelist.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AccountWhitelist resources.
 */
export interface AccountWhitelistState {
    /**
     * The free form name of the whitelist.
     */
    name?: pulumi.Input<string>;
    /**
     * Array of IP addresses/networks from which to allow access.
     */
    values?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a AccountWhitelist resource.
 */
export interface AccountWhitelistArgs {
    /**
     * The free form name of the whitelist.
     */
    name?: pulumi.Input<string>;
    /**
     * Array of IP addresses/networks from which to allow access.
     */
    values: pulumi.Input<pulumi.Input<string>[]>;
}
