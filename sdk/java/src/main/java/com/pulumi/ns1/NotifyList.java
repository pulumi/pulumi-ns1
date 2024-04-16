// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ns1.NotifyListArgs;
import com.pulumi.ns1.Utilities;
import com.pulumi.ns1.inputs.NotifyListState;
import com.pulumi.ns1.outputs.NotifyListNotification;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a NS1 Notify List resource. This can be used to create, modify, and delete notify lists.
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
 * import com.pulumi.ns1.NotifyList;
 * import com.pulumi.ns1.NotifyListArgs;
 * import com.pulumi.ns1.inputs.NotifyListNotificationArgs;
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
 *         var nl = new NotifyList(&#34;nl&#34;, NotifyListArgs.builder()        
 *             .name(&#34;my notify list&#34;)
 *             .notifications(            
 *                 NotifyListNotificationArgs.builder()
 *                     .type(&#34;webhook&#34;)
 *                     .config(Map.of(&#34;url&#34;, &#34;http://www.mywebhook.com&#34;))
 *                     .build(),
 *                 NotifyListNotificationArgs.builder()
 *                     .type(&#34;email&#34;)
 *                     .config(Map.of(&#34;email&#34;, &#34;test@test.com&#34;))
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## NS1 Documentation
 * 
 * [NotifyList Api Doc](https://ns1.com/api#notification-lists)
 * 
 * ## Import
 * 
 * ```sh
 * $ pulumi import ns1:index/notifyList:NotifyList &lt;name&gt; &lt;notifylist_id&gt;`
 * ```
 * 
 */
@ResourceType(type="ns1:index/notifyList:NotifyList")
public class NotifyList extends com.pulumi.resources.CustomResource {
    /**
     * The free-form display name for the notify list.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The free-form display name for the notify list.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * A list of notifiers. All notifiers in a notification list will receive notifications whenever an event is send to the list (e.g., when a monitoring job fails). Notifiers are documented below.
     * 
     */
    @Export(name="notifications", refs={List.class,NotifyListNotification.class}, tree="[0,1]")
    private Output</* @Nullable */ List<NotifyListNotification>> notifications;

    /**
     * @return A list of notifiers. All notifiers in a notification list will receive notifications whenever an event is send to the list (e.g., when a monitoring job fails). Notifiers are documented below.
     * 
     */
    public Output<Optional<List<NotifyListNotification>>> notifications() {
        return Codegen.optional(this.notifications);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NotifyList(String name) {
        this(name, NotifyListArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NotifyList(String name, @Nullable NotifyListArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NotifyList(String name, @Nullable NotifyListArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ns1:index/notifyList:NotifyList", name, args == null ? NotifyListArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private NotifyList(String name, Output<String> id, @Nullable NotifyListState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ns1:index/notifyList:NotifyList", name, state, makeResourceOptions(options, id));
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
    public static NotifyList get(String name, Output<String> id, @Nullable NotifyListState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NotifyList(name, id, state, options);
    }
}
