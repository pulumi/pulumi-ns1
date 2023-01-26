// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ns1.DataSourceArgs;
import com.pulumi.ns1.Utilities;
import com.pulumi.ns1.inputs.DataSourceState;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a NS1 Data Source resource. This can be used to create, modify, and delete data sources.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ns1.DataSource;
 * import com.pulumi.ns1.DataSourceArgs;
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
 *         var example = new DataSource(&#34;example&#34;, DataSourceArgs.builder()        
 *             .sourcetype(&#34;nsone_v1&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## NS1 Documentation
 * 
 * [Datasource Api Doc](https://ns1.com/api#data-sources)
 * 
 * ## Import
 * 
 * ```sh
 *  $ pulumi import ns1:index/dataSource:DataSource &lt;name&gt; &lt;datasource_id&gt;`
 * ```
 * 
 */
@ResourceType(type="ns1:index/dataSource:DataSource")
public class DataSource extends com.pulumi.resources.CustomResource {
    /**
     * The data source configuration, determined by its type,
     * matching the specification in `config` from /data/sourcetypes.
     * 
     */
    @Export(name="config", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> config;

    /**
     * @return The data source configuration, determined by its type,
     * matching the specification in `config` from /data/sourcetypes.
     * 
     */
    public Output<Optional<Map<String,Object>>> config() {
        return Codegen.optional(this.config);
    }
    /**
     * The free form name of the data source.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The free form name of the data source.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The data sources type, listed in API endpoint https://api.nsone.net/v1/data/sourcetypes.
     * 
     */
    @Export(name="sourcetype", type=String.class, parameters={})
    private Output<String> sourcetype;

    /**
     * @return The data sources type, listed in API endpoint https://api.nsone.net/v1/data/sourcetypes.
     * 
     */
    public Output<String> sourcetype() {
        return this.sourcetype;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DataSource(String name) {
        this(name, DataSourceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DataSource(String name, DataSourceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DataSource(String name, DataSourceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ns1:index/dataSource:DataSource", name, args == null ? DataSourceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DataSource(String name, Output<String> id, @Nullable DataSourceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ns1:index/dataSource:DataSource", name, state, makeResourceOptions(options, id));
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
    public static DataSource get(String name, Output<String> id, @Nullable DataSourceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DataSource(name, id, state, options);
    }
}
