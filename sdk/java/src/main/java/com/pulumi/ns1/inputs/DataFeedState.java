// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DataFeedState extends com.pulumi.resources.ResourceArgs {

    public static final DataFeedState Empty = new DataFeedState();

    /**
     * The feeds configuration matching the specification in
     * `feed_config` from /data/sourcetypes. `jobid` is required in the `config` for datafeeds connected to NS1 monitoring.
     * 
     */
    @Import(name="config")
    private @Nullable Output<Map<String,String>> config;

    /**
     * @return The feeds configuration matching the specification in
     * `feed_config` from /data/sourcetypes. `jobid` is required in the `config` for datafeeds connected to NS1 monitoring.
     * 
     */
    public Optional<Output<Map<String,String>>> config() {
        return Optional.ofNullable(this.config);
    }

    /**
     * The free form name of the data feed.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The free form name of the data feed.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The data source id that this feed is connected to.
     * 
     */
    @Import(name="sourceId")
    private @Nullable Output<String> sourceId;

    /**
     * @return The data source id that this feed is connected to.
     * 
     */
    public Optional<Output<String>> sourceId() {
        return Optional.ofNullable(this.sourceId);
    }

    private DataFeedState() {}

    private DataFeedState(DataFeedState $) {
        this.config = $.config;
        this.name = $.name;
        this.sourceId = $.sourceId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DataFeedState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DataFeedState $;

        public Builder() {
            $ = new DataFeedState();
        }

        public Builder(DataFeedState defaults) {
            $ = new DataFeedState(Objects.requireNonNull(defaults));
        }

        /**
         * @param config The feeds configuration matching the specification in
         * `feed_config` from /data/sourcetypes. `jobid` is required in the `config` for datafeeds connected to NS1 monitoring.
         * 
         * @return builder
         * 
         */
        public Builder config(@Nullable Output<Map<String,String>> config) {
            $.config = config;
            return this;
        }

        /**
         * @param config The feeds configuration matching the specification in
         * `feed_config` from /data/sourcetypes. `jobid` is required in the `config` for datafeeds connected to NS1 monitoring.
         * 
         * @return builder
         * 
         */
        public Builder config(Map<String,String> config) {
            return config(Output.of(config));
        }

        /**
         * @param name The free form name of the data feed.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The free form name of the data feed.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param sourceId The data source id that this feed is connected to.
         * 
         * @return builder
         * 
         */
        public Builder sourceId(@Nullable Output<String> sourceId) {
            $.sourceId = sourceId;
            return this;
        }

        /**
         * @param sourceId The data source id that this feed is connected to.
         * 
         * @return builder
         * 
         */
        public Builder sourceId(String sourceId) {
            return sourceId(Output.of(sourceId));
        }

        public DataFeedState build() {
            return $;
        }
    }

}
