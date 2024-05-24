// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ns1.DataFeedArgs;
import com.pulumi.ns1.Utilities;
import com.pulumi.ns1.inputs.DataFeedState;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a NS1 Data Feed resource. This can be used to create, modify, and delete data feeds.
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
 * import com.pulumi.ns1.DataSource;
 * import com.pulumi.ns1.DataSourceArgs;
 * import com.pulumi.ns1.DataFeed;
 * import com.pulumi.ns1.DataFeedArgs;
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
 *         var example = new DataSource("example", DataSourceArgs.builder()
 *             .name("example")
 *             .sourcetype("nsone_v1")
 *             .build());
 * 
 *         var exampleMonitoring = new DataSource("exampleMonitoring", DataSourceArgs.builder()
 *             .name("example_monitoring")
 *             .sourcetype("nsone_monitoring")
 *             .build());
 * 
 *         var uswestFeed = new DataFeed("uswestFeed", DataFeedArgs.builder()
 *             .name("uswest_feed")
 *             .sourceId(example.id())
 *             .config(Map.of("label", "uswest"))
 *             .build());
 * 
 *         var useastFeed = new DataFeed("useastFeed", DataFeedArgs.builder()
 *             .name("useast_feed")
 *             .sourceId(example.id())
 *             .config(Map.of("label", "useast"))
 *             .build());
 * 
 *         var useastMonitorFeed = new DataFeed("useastMonitorFeed", DataFeedArgs.builder()
 *             .name("useast_monitor_feed")
 *             .sourceId(exampleMonitoring.id())
 *             .config(Map.of("jobid", exampleJob.id()))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## NS1 Documentation
 * 
 * [Datafeed Api Doc](https://ns1.com/api#data-feeds)
 * 
 * ## Import
 * 
 * ```sh
 * $ pulumi import ns1:index/dataFeed:DataFeed &lt;name&gt; &lt;datasource_id&gt;/&lt;datafeed_id&gt;`
 * ```
 * 
 */
@ResourceType(type="ns1:index/dataFeed:DataFeed")
public class DataFeed extends com.pulumi.resources.CustomResource {
    /**
     * The feeds configuration matching the specification in
     * `feed_config` from /data/sourcetypes. `jobid` is required in the `config` for datafeeds connected to NS1 monitoring.
     * 
     */
    @Export(name="config", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> config;

    /**
     * @return The feeds configuration matching the specification in
     * `feed_config` from /data/sourcetypes. `jobid` is required in the `config` for datafeeds connected to NS1 monitoring.
     * 
     */
    public Output<Optional<Map<String,Object>>> config() {
        return Codegen.optional(this.config);
    }
    /**
     * The free form name of the data feed.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The free form name of the data feed.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The data source id that this feed is connected to.
     * 
     */
    @Export(name="sourceId", refs={String.class}, tree="[0]")
    private Output<String> sourceId;

    /**
     * @return The data source id that this feed is connected to.
     * 
     */
    public Output<String> sourceId() {
        return this.sourceId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DataFeed(String name) {
        this(name, DataFeedArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DataFeed(String name, DataFeedArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DataFeed(String name, DataFeedArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ns1:index/dataFeed:DataFeed", name, args == null ? DataFeedArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DataFeed(String name, Output<String> id, @Nullable DataFeedState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ns1:index/dataFeed:DataFeed", name, state, makeResourceOptions(options, id));
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
    public static DataFeed get(String name, Output<String> id, @Nullable DataFeedState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DataFeed(name, id, state, options);
    }
}
