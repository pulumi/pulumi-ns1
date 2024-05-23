// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ns1.AccountWhitelistArgs;
import com.pulumi.ns1.Utilities;
import com.pulumi.ns1.inputs.AccountWhitelistState;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Provides a NS1 Global IP Whitelist resource.
 * 
 * This can be used to create, modify, and delete Global IP Whitelists.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ns1.AccountWhitelist;
 * import com.pulumi.ns1.AccountWhitelistArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var example = new AccountWhitelist("example", AccountWhitelistArgs.builder()
 *             .name("Example Whitelist")
 *             .values(            
 *                 "1.1.1.1",
 *                 "2.2.2.2")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * &gt; You current source IP must be present in one of the whitelists to prevent locking yourself out.
 * 
 * ## NS1 Documentation
 * 
 * [Global IP Whitelist Doc](https://ns1.com/api?docId=2282)
 * 
 * ## Import
 * 
 * ```sh
 * $ pulumi import ns1:index/accountWhitelist:AccountWhitelist example &lt;ID&gt;`
 * ```
 * 
 */
@ResourceType(type="ns1:index/accountWhitelist:AccountWhitelist")
public class AccountWhitelist extends com.pulumi.resources.CustomResource {
    /**
     * The free form name of the whitelist.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The free form name of the whitelist.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Array of IP addresses/networks from which to allow access.
     * 
     */
    @Export(name="values", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> values;

    /**
     * @return Array of IP addresses/networks from which to allow access.
     * 
     */
    public Output<List<String>> values() {
        return this.values;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AccountWhitelist(String name) {
        this(name, AccountWhitelistArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AccountWhitelist(String name, AccountWhitelistArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AccountWhitelist(String name, AccountWhitelistArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ns1:index/accountWhitelist:AccountWhitelist", name, args == null ? AccountWhitelistArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AccountWhitelist(String name, Output<String> id, @Nullable AccountWhitelistState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ns1:index/accountWhitelist:AccountWhitelist", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static AccountWhitelist get(String name, Output<String> id, @Nullable AccountWhitelistState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AccountWhitelist(name, id, state, options);
    }
}
