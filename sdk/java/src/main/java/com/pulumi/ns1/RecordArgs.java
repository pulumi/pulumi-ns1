// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.ns1.inputs.RecordAnswerArgs;
import com.pulumi.ns1.inputs.RecordFilterArgs;
import com.pulumi.ns1.inputs.RecordRegionArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RecordArgs extends com.pulumi.resources.ResourceArgs {

    public static final RecordArgs Empty = new RecordArgs();

    /**
     * One or more NS1 answers for the records&#39; specified type.
     * Answers are documented below.
     * 
     */
    @Import(name="answers")
    private @Nullable Output<List<RecordAnswerArgs>> answers;

    /**
     * @return One or more NS1 answers for the records&#39; specified type.
     * Answers are documented below.
     * 
     */
    public Optional<Output<List<RecordAnswerArgs>>> answers() {
        return Optional.ofNullable(this.answers);
    }

    /**
     * The records&#39; domain. Cannot have leading or trailing
     * dots - see the example above and `FQDN formatting` below.
     * 
     */
    @Import(name="domain", required=true)
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
    @Import(name="filters")
    private @Nullable Output<List<RecordFilterArgs>> filters;

    /**
     * @return One or more NS1 filters for the record(order matters).
     * Filters are documented below.
     * 
     */
    public Optional<Output<List<RecordFilterArgs>>> filters() {
        return Optional.ofNullable(this.filters);
    }

    /**
     * The target record to link to. This means this record is a
     * &#39;linked&#39; record, and it inherits all properties from its target.
     * 
     */
    @Import(name="link")
    private @Nullable Output<String> link;

    /**
     * @return The target record to link to. This means this record is a
     * &#39;linked&#39; record, and it inherits all properties from its target.
     * 
     */
    public Optional<Output<String>> link() {
        return Optional.ofNullable(this.link);
    }

    @Import(name="meta")
    private @Nullable Output<Map<String,Object>> meta;

    public Optional<Output<Map<String,Object>>> meta() {
        return Optional.ofNullable(this.meta);
    }

    @Import(name="overrideTtl")
    private @Nullable Output<Boolean> overrideTtl;

    public Optional<Output<Boolean>> overrideTtl() {
        return Optional.ofNullable(this.overrideTtl);
    }

    /**
     * One or more &#34;regions&#34; for the record. These are really
     * just groupings based on metadata, and are called &#34;Answer Groups&#34; in the NS1 UI,
     * but remain `regions` here for legacy reasons. Regions are
     * documented below. Please note the ordering requirement!
     * 
     */
    @Import(name="regions")
    private @Nullable Output<List<RecordRegionArgs>> regions;

    /**
     * @return One or more &#34;regions&#34; for the record. These are really
     * just groupings based on metadata, and are called &#34;Answer Groups&#34; in the NS1 UI,
     * but remain `regions` here for legacy reasons. Regions are
     * documented below. Please note the ordering requirement!
     * 
     */
    public Optional<Output<List<RecordRegionArgs>>> regions() {
        return Optional.ofNullable(this.regions);
    }

    /**
     * @deprecated
     * short_answers will be deprecated in a future release. It is suggested to migrate to a regular &#34;answers&#34; block.
     * 
     */
    @Deprecated /* short_answers will be deprecated in a future release. It is suggested to migrate to a regular ""answers"" block. */
    @Import(name="shortAnswers")
    private @Nullable Output<List<String>> shortAnswers;

    /**
     * @deprecated
     * short_answers will be deprecated in a future release. It is suggested to migrate to a regular &#34;answers&#34; block.
     * 
     */
    @Deprecated /* short_answers will be deprecated in a future release. It is suggested to migrate to a regular ""answers"" block. */
    public Optional<Output<List<String>>> shortAnswers() {
        return Optional.ofNullable(this.shortAnswers);
    }

    /**
     * The records&#39; time to live (in seconds).
     * 
     */
    @Import(name="ttl")
    private @Nullable Output<Integer> ttl;

    /**
     * @return The records&#39; time to live (in seconds).
     * 
     */
    public Optional<Output<Integer>> ttl() {
        return Optional.ofNullable(this.ttl);
    }

    /**
     * The records&#39; RR type.
     * 
     */
    @Import(name="type", required=true)
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
    @Import(name="useClientSubnet")
    private @Nullable Output<Boolean> useClientSubnet;

    /**
     * @return Whether to use EDNS client subnet data when
     * available(in filter chain).
     * * `  meta ` - (Optional) meta is supported at the `record` level. Meta
     *   is documented below.
     * 
     */
    public Optional<Output<Boolean>> useClientSubnet() {
        return Optional.ofNullable(this.useClientSubnet);
    }

    /**
     * The zone the record belongs to. Cannot have leading or
     * trailing dots (&#34;.&#34;) - see the example above and `FQDN formatting` below.
     * 
     */
    @Import(name="zone", required=true)
    private Output<String> zone;

    /**
     * @return The zone the record belongs to. Cannot have leading or
     * trailing dots (&#34;.&#34;) - see the example above and `FQDN formatting` below.
     * 
     */
    public Output<String> zone() {
        return this.zone;
    }

    private RecordArgs() {}

    private RecordArgs(RecordArgs $) {
        this.answers = $.answers;
        this.domain = $.domain;
        this.filters = $.filters;
        this.link = $.link;
        this.meta = $.meta;
        this.overrideTtl = $.overrideTtl;
        this.regions = $.regions;
        this.shortAnswers = $.shortAnswers;
        this.ttl = $.ttl;
        this.type = $.type;
        this.useClientSubnet = $.useClientSubnet;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RecordArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RecordArgs $;

        public Builder() {
            $ = new RecordArgs();
        }

        public Builder(RecordArgs defaults) {
            $ = new RecordArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param answers One or more NS1 answers for the records&#39; specified type.
         * Answers are documented below.
         * 
         * @return builder
         * 
         */
        public Builder answers(@Nullable Output<List<RecordAnswerArgs>> answers) {
            $.answers = answers;
            return this;
        }

        /**
         * @param answers One or more NS1 answers for the records&#39; specified type.
         * Answers are documented below.
         * 
         * @return builder
         * 
         */
        public Builder answers(List<RecordAnswerArgs> answers) {
            return answers(Output.of(answers));
        }

        /**
         * @param answers One or more NS1 answers for the records&#39; specified type.
         * Answers are documented below.
         * 
         * @return builder
         * 
         */
        public Builder answers(RecordAnswerArgs... answers) {
            return answers(List.of(answers));
        }

        /**
         * @param domain The records&#39; domain. Cannot have leading or trailing
         * dots - see the example above and `FQDN formatting` below.
         * 
         * @return builder
         * 
         */
        public Builder domain(Output<String> domain) {
            $.domain = domain;
            return this;
        }

        /**
         * @param domain The records&#39; domain. Cannot have leading or trailing
         * dots - see the example above and `FQDN formatting` below.
         * 
         * @return builder
         * 
         */
        public Builder domain(String domain) {
            return domain(Output.of(domain));
        }

        /**
         * @param filters One or more NS1 filters for the record(order matters).
         * Filters are documented below.
         * 
         * @return builder
         * 
         */
        public Builder filters(@Nullable Output<List<RecordFilterArgs>> filters) {
            $.filters = filters;
            return this;
        }

        /**
         * @param filters One or more NS1 filters for the record(order matters).
         * Filters are documented below.
         * 
         * @return builder
         * 
         */
        public Builder filters(List<RecordFilterArgs> filters) {
            return filters(Output.of(filters));
        }

        /**
         * @param filters One or more NS1 filters for the record(order matters).
         * Filters are documented below.
         * 
         * @return builder
         * 
         */
        public Builder filters(RecordFilterArgs... filters) {
            return filters(List.of(filters));
        }

        /**
         * @param link The target record to link to. This means this record is a
         * &#39;linked&#39; record, and it inherits all properties from its target.
         * 
         * @return builder
         * 
         */
        public Builder link(@Nullable Output<String> link) {
            $.link = link;
            return this;
        }

        /**
         * @param link The target record to link to. This means this record is a
         * &#39;linked&#39; record, and it inherits all properties from its target.
         * 
         * @return builder
         * 
         */
        public Builder link(String link) {
            return link(Output.of(link));
        }

        public Builder meta(@Nullable Output<Map<String,Object>> meta) {
            $.meta = meta;
            return this;
        }

        public Builder meta(Map<String,Object> meta) {
            return meta(Output.of(meta));
        }

        public Builder overrideTtl(@Nullable Output<Boolean> overrideTtl) {
            $.overrideTtl = overrideTtl;
            return this;
        }

        public Builder overrideTtl(Boolean overrideTtl) {
            return overrideTtl(Output.of(overrideTtl));
        }

        /**
         * @param regions One or more &#34;regions&#34; for the record. These are really
         * just groupings based on metadata, and are called &#34;Answer Groups&#34; in the NS1 UI,
         * but remain `regions` here for legacy reasons. Regions are
         * documented below. Please note the ordering requirement!
         * 
         * @return builder
         * 
         */
        public Builder regions(@Nullable Output<List<RecordRegionArgs>> regions) {
            $.regions = regions;
            return this;
        }

        /**
         * @param regions One or more &#34;regions&#34; for the record. These are really
         * just groupings based on metadata, and are called &#34;Answer Groups&#34; in the NS1 UI,
         * but remain `regions` here for legacy reasons. Regions are
         * documented below. Please note the ordering requirement!
         * 
         * @return builder
         * 
         */
        public Builder regions(List<RecordRegionArgs> regions) {
            return regions(Output.of(regions));
        }

        /**
         * @param regions One or more &#34;regions&#34; for the record. These are really
         * just groupings based on metadata, and are called &#34;Answer Groups&#34; in the NS1 UI,
         * but remain `regions` here for legacy reasons. Regions are
         * documented below. Please note the ordering requirement!
         * 
         * @return builder
         * 
         */
        public Builder regions(RecordRegionArgs... regions) {
            return regions(List.of(regions));
        }

        /**
         * @return builder
         * 
         * @deprecated
         * short_answers will be deprecated in a future release. It is suggested to migrate to a regular &#34;answers&#34; block.
         * 
         */
        @Deprecated /* short_answers will be deprecated in a future release. It is suggested to migrate to a regular ""answers"" block. */
        public Builder shortAnswers(@Nullable Output<List<String>> shortAnswers) {
            $.shortAnswers = shortAnswers;
            return this;
        }

        /**
         * @return builder
         * 
         * @deprecated
         * short_answers will be deprecated in a future release. It is suggested to migrate to a regular &#34;answers&#34; block.
         * 
         */
        @Deprecated /* short_answers will be deprecated in a future release. It is suggested to migrate to a regular ""answers"" block. */
        public Builder shortAnswers(List<String> shortAnswers) {
            return shortAnswers(Output.of(shortAnswers));
        }

        /**
         * @return builder
         * 
         * @deprecated
         * short_answers will be deprecated in a future release. It is suggested to migrate to a regular &#34;answers&#34; block.
         * 
         */
        @Deprecated /* short_answers will be deprecated in a future release. It is suggested to migrate to a regular ""answers"" block. */
        public Builder shortAnswers(String... shortAnswers) {
            return shortAnswers(List.of(shortAnswers));
        }

        /**
         * @param ttl The records&#39; time to live (in seconds).
         * 
         * @return builder
         * 
         */
        public Builder ttl(@Nullable Output<Integer> ttl) {
            $.ttl = ttl;
            return this;
        }

        /**
         * @param ttl The records&#39; time to live (in seconds).
         * 
         * @return builder
         * 
         */
        public Builder ttl(Integer ttl) {
            return ttl(Output.of(ttl));
        }

        /**
         * @param type The records&#39; RR type.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The records&#39; RR type.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param useClientSubnet Whether to use EDNS client subnet data when
         * available(in filter chain).
         * * `  meta ` - (Optional) meta is supported at the `record` level. Meta
         *   is documented below.
         * 
         * @return builder
         * 
         */
        public Builder useClientSubnet(@Nullable Output<Boolean> useClientSubnet) {
            $.useClientSubnet = useClientSubnet;
            return this;
        }

        /**
         * @param useClientSubnet Whether to use EDNS client subnet data when
         * available(in filter chain).
         * * `  meta ` - (Optional) meta is supported at the `record` level. Meta
         *   is documented below.
         * 
         * @return builder
         * 
         */
        public Builder useClientSubnet(Boolean useClientSubnet) {
            return useClientSubnet(Output.of(useClientSubnet));
        }

        /**
         * @param zone The zone the record belongs to. Cannot have leading or
         * trailing dots (&#34;.&#34;) - see the example above and `FQDN formatting` below.
         * 
         * @return builder
         * 
         */
        public Builder zone(Output<String> zone) {
            $.zone = zone;
            return this;
        }

        /**
         * @param zone The zone the record belongs to. Cannot have leading or
         * trailing dots (&#34;.&#34;) - see the example above and `FQDN formatting` below.
         * 
         * @return builder
         * 
         */
        public Builder zone(String zone) {
            return zone(Output.of(zone));
        }

        public RecordArgs build() {
            $.domain = Objects.requireNonNull($.domain, "expected parameter 'domain' to be non-null");
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            $.zone = Objects.requireNonNull($.zone, "expected parameter 'zone' to be non-null");
            return $;
        }
    }

}
