// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ns1.TsigkeyArgs;
import com.pulumi.ns1.Utilities;
import com.pulumi.ns1.inputs.TsigkeyState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a NS1 TSIG Key resource. This can be used to create, modify, and delete TSIG keys.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ns1.Tsigkey;
 * import com.pulumi.ns1.TsigkeyArgs;
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
 *         var example = new Tsigkey(&#34;example&#34;, TsigkeyArgs.builder()        
 *             .name(&#34;ExampleTsigKey&#34;)
 *             .algorithm(&#34;hmac-sha256&#34;)
 *             .secret(&#34;Ok1qR5IW1ajVka5cHPEJQIXfLyx5V3PSkFBROAzOn21JumDq6nIpoj6H8rfj5Uo+Ok55ZWQ0Wgrf302fDscHLA==&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * ## NS1 Documentation
 * 
 * [TSIG Keys Api Doc](https://ns1.com/api/#tsig)
 * 
 * ## Import
 * 
 * ```sh
 * $ pulumi import ns1:index/tsigkey:Tsigkey importTest &lt;name&gt;`
 * ```
 * 
 */
@ResourceType(type="ns1:index/tsigkey:Tsigkey")
public class Tsigkey extends com.pulumi.resources.CustomResource {
    /**
     * The algorithm used to hash the TSIG key&#39;s secret.
     * 
     */
    @Export(name="algorithm", refs={String.class}, tree="[0]")
    private Output<String> algorithm;

    /**
     * @return The algorithm used to hash the TSIG key&#39;s secret.
     * 
     */
    public Output<String> algorithm() {
        return this.algorithm;
    }
    /**
     * The free form name of the tsigkey.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The free form name of the tsigkey.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The key&#39;s secret to be hashed.
     * 
     */
    @Export(name="secret", refs={String.class}, tree="[0]")
    private Output<String> secret;

    /**
     * @return The key&#39;s secret to be hashed.
     * 
     */
    public Output<String> secret() {
        return this.secret;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Tsigkey(String name) {
        this(name, TsigkeyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Tsigkey(String name, TsigkeyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Tsigkey(String name, TsigkeyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ns1:index/tsigkey:Tsigkey", name, args == null ? TsigkeyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Tsigkey(String name, Output<String> id, @Nullable TsigkeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ns1:index/tsigkey:Tsigkey", name, state, makeResourceOptions(options, id));
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
    public static Tsigkey get(String name, Output<String> id, @Nullable TsigkeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Tsigkey(name, id, state, options);
    }
}
