// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.ns1.inputs.RecordAnswerArgs;
import com.pulumi.ns1.inputs.RecordFilterArgs;
import com.pulumi.ns1.inputs.RecordRegionArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RecordState extends com.pulumi.resources.ResourceArgs {

    public static final RecordState Empty = new RecordState();

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

    @Import(name="blockedTags")
    private @Nullable Output<List<String>> blockedTags;

    public Optional<Output<List<String>>> blockedTags() {
        return Optional.ofNullable(this.blockedTags);
    }

    /**
     * The records&#39; domain. Cannot have leading or trailing
     * dots - see the example above and `FQDN formatting` below.
     * 
     */
    @Import(name="domain")
    private @Nullable Output<String> domain;

    /**
     * @return The records&#39; domain. Cannot have leading or trailing
     * dots - see the example above and `FQDN formatting` below.
     * 
     */
    public Optional<Output<String>> domain() {
        return Optional.ofNullable(this.domain);
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
    private @Nullable Output<Map<String,String>> meta;

    public Optional<Output<Map<String,String>>> meta() {
        return Optional.ofNullable(this.meta);
    }

    @Import(name="overrideAddressRecords")
    private @Nullable Output<Boolean> overrideAddressRecords;

    public Optional<Output<Boolean>> overrideAddressRecords() {
        return Optional.ofNullable(this.overrideAddressRecords);
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
     * map of tags in the form of `&#34;key&#34; = &#34;value&#34;` where both key and value are strings
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return map of tags in the form of `&#34;key&#34; = &#34;value&#34;` where both key and value are strings
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
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
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return The records&#39; RR type.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
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
    @Import(name="zone")
    private @Nullable Output<String> zone;

    /**
     * @return The zone the record belongs to. Cannot have leading or
     * trailing dots (&#34;.&#34;) - see the example above and `FQDN formatting` below.
     * 
     */
    public Optional<Output<String>> zone() {
        return Optional.ofNullable(this.zone);
    }

    private RecordState() {}

    private RecordState(RecordState $) {
        this.answers = $.answers;
        this.blockedTags = $.blockedTags;
        this.domain = $.domain;
        this.filters = $.filters;
        this.link = $.link;
        this.meta = $.meta;
        this.overrideAddressRecords = $.overrideAddressRecords;
        this.overrideTtl = $.overrideTtl;
        this.regions = $.regions;
        this.shortAnswers = $.shortAnswers;
        this.tags = $.tags;
        this.ttl = $.ttl;
        this.type = $.type;
        this.useClientSubnet = $.useClientSubnet;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RecordState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RecordState $;

        public Builder() {
            $ = new RecordState();
        }

        public Builder(RecordState defaults) {
            $ = new RecordState(Objects.requireNonNull(defaults));
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

        public Builder blockedTags(@Nullable Output<List<String>> blockedTags) {
            $.blockedTags = blockedTags;
            return this;
        }

        public Builder blockedTags(List<String> blockedTags) {
            return blockedTags(Output.of(blockedTags));
        }

        public Builder blockedTags(String... blockedTags) {
            return blockedTags(List.of(blockedTags));
        }

        /**
         * @param domain The records&#39; domain. Cannot have leading or trailing
         * dots - see the example above and `FQDN formatting` below.
         * 
         * @return builder
         * 
         */
        public Builder domain(@Nullable Output<String> domain) {
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

        public Builder meta(@Nullable Output<Map<String,String>> meta) {
            $.meta = meta;
            return this;
        }

        public Builder meta(Map<String,String> meta) {
            return meta(Output.of(meta));
        }

        public Builder overrideAddressRecords(@Nullable Output<Boolean> overrideAddressRecords) {
            $.overrideAddressRecords = overrideAddressRecords;
            return this;
        }

        public Builder overrideAddressRecords(Boolean overrideAddressRecords) {
            return overrideAddressRecords(Output.of(overrideAddressRecords));
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
         * @param tags map of tags in the form of `&#34;key&#34; = &#34;value&#34;` where both key and value are strings
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags map of tags in the form of `&#34;key&#34; = &#34;value&#34;` where both key and value are strings
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
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
        public Builder type(@Nullable Output<String> type) {
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
        public Builder zone(@Nullable Output<String> zone) {
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

        public RecordState build() {
            return $;
        }
    }

}
