// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ns1.RecordArgs;
import com.pulumi.ns1.Utilities;
import com.pulumi.ns1.inputs.RecordState;
import com.pulumi.ns1.outputs.RecordAnswer;
import com.pulumi.ns1.outputs.RecordFilter;
import com.pulumi.ns1.outputs.RecordRegion;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a NS1 Record resource. This can be used to create, modify, and delete records.
 * 
 * ## NS1 Documentation
 * 
 * [Record Api Doc](https://ns1.com/api#records)
 * 
 * ## Import
 * 
 * ```sh
 *  $ pulumi import ns1:index/record:Record &lt;name&gt; &lt;zone&gt;/&lt;domain&gt;/&lt;type&gt;`
 * ```
 * 
 *  So for the example above
 * 
 * ```sh
 *  $ pulumi import ns1:index/record:Record www terraform.example.io/www.terraform.example.io/CNAME`
 * ```
 * 
 */
@ResourceType(type="ns1:index/record:Record")
public class Record extends com.pulumi.resources.CustomResource {
    /**
     * One or more NS1 answers for the records&#39; specified type.
     * Answers are documented below.
     * 
     */
    @Export(name="answers", type=List.class, parameters={RecordAnswer.class})
    private Output</* @Nullable */ List<RecordAnswer>> answers;

    /**
     * @return One or more NS1 answers for the records&#39; specified type.
     * Answers are documented below.
     * 
     */
    public Output<Optional<List<RecordAnswer>>> answers() {
        return Codegen.optional(this.answers);
    }
    /**
     * The records&#39; domain. Cannot have leading or trailing
     * dots - see the example above and `FQDN formatting` below.
     * 
     */
    @Export(name="domain", type=String.class, parameters={})
    private Output<String> domain;

    /**
     * @return The records&#39; domain. Cannot have leading or trailing
     * dots - see the example above and `FQDN formatting` below.
     * 
     */
    public Output<String> domain() {
        return this.domain;
    }
    /**
     * One or more NS1 filters for the record(order matters).
     * Filters are documented below.
     * 
     */
    @Export(name="filters", type=List.class, parameters={RecordFilter.class})
    private Output</* @Nullable */ List<RecordFilter>> filters;

    /**
     * @return One or more NS1 filters for the record(order matters).
     * Filters are documented below.
     * 
     */
    public Output<Optional<List<RecordFilter>>> filters() {
        return Codegen.optional(this.filters);
    }
    /**
     * The target record to link to. This means this record is a
     * &#39;linked&#39; record, and it inherits all properties from its target.
     * 
     */
    @Export(name="link", type=String.class, parameters={})
    private Output</* @Nullable */ String> link;

    /**
     * @return The target record to link to. This means this record is a
     * &#39;linked&#39; record, and it inherits all properties from its target.
     * 
     */
    public Output<Optional<String>> link() {
        return Codegen.optional(this.link);
    }
    @Export(name="meta", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> meta;

    public Output<Optional<Map<String,Object>>> meta() {
        return Codegen.optional(this.meta);
    }
    @Export(name="overrideTtl", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> overrideTtl;

    public Output<Optional<Boolean>> overrideTtl() {
        return Codegen.optional(this.overrideTtl);
    }
    /**
     * One or more &#34;regions&#34; for the record. These are really
     * just groupings based on metadata, and are called &#34;Answer Groups&#34; in the NS1 UI,
     * but remain `regions` here for legacy reasons. Regions are
     * documented below. Please note the ordering requirement!
     * 
     */
    @Export(name="regions", type=List.class, parameters={RecordRegion.class})
    private Output</* @Nullable */ List<RecordRegion>> regions;

    /**
     * @return One or more &#34;regions&#34; for the record. These are really
     * just groupings based on metadata, and are called &#34;Answer Groups&#34; in the NS1 UI,
     * but remain `regions` here for legacy reasons. Regions are
     * documented below. Please note the ordering requirement!
     * 
     */
    public Output<Optional<List<RecordRegion>>> regions() {
        return Codegen.optional(this.regions);
    }
    /**
     * @deprecated
     * short_answers will be deprecated in a future release. It is suggested to migrate to a regular &#34;answers&#34; block.
     * 
     */
    @Deprecated /* short_answers will be deprecated in a future release. It is suggested to migrate to a regular ""answers"" block. */
    @Export(name="shortAnswers", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> shortAnswers;

    public Output<Optional<List<String>>> shortAnswers() {
        return Codegen.optional(this.shortAnswers);
    }
    /**
     * The records&#39; time to live (in seconds).
     * 
     */
    @Export(name="ttl", type=Integer.class, parameters={})
    private Output<Integer> ttl;

    /**
     * @return The records&#39; time to live (in seconds).
     * 
     */
    public Output<Integer> ttl() {
        return this.ttl;
    }
    /**
     * The records&#39; RR type.
     * 
     */
    @Export(name="type", type=String.class, parameters={})
    private Output<String> type;

    /**
     * @return The records&#39; RR type.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * Whether to use EDNS client subnet data when
     * available(in filter chain).
     * * `  meta ` - (Optional) meta is supported at the `record` level. Meta
     *   is documented below.
     * 
     */
    @Export(name="useClientSubnet", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> useClientSubnet;

    /**
     * @return Whether to use EDNS client subnet data when
     * available(in filter chain).
     * * `  meta ` - (Optional) meta is supported at the `record` level. Meta
     *   is documented below.
     * 
     */
    public Output<Optional<Boolean>> useClientSubnet() {
        return Codegen.optional(this.useClientSubnet);
    }
    /**
     * The zone the record belongs to. Cannot have leading or
     * trailing dots (&#34;.&#34;) - see the example above and `FQDN formatting` below.
     * 
     */
    @Export(name="zone", type=String.class, parameters={})
    private Output<String> zone;

    /**
     * @return The zone the record belongs to. Cannot have leading or
     * trailing dots (&#34;.&#34;) - see the example above and `FQDN formatting` below.
     * 
     */
    public Output<String> zone() {
        return this.zone;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Record(String name) {
        this(name, RecordArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Record(String name, RecordArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Record(String name, RecordArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ns1:index/record:Record", name, args == null ? RecordArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Record(String name, Output<String> id, @Nullable RecordState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ns1:index/record:Record", name, state, makeResourceOptions(options, id));
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
    public static Record get(String name, Output<String> id, @Nullable RecordState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Record(name, id, state, options);
    }
}