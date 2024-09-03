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


public final class DataSourceState extends com.pulumi.resources.ResourceArgs {

    public static final DataSourceState Empty = new DataSourceState();

    /**
     * The data source configuration, determined by its type,
     * matching the specification in `config` from /data/sourcetypes.
     * 
     */
    @Import(name="config")
    private @Nullable Output<Map<String,String>> config;

    /**
     * @return The data source configuration, determined by its type,
     * matching the specification in `config` from /data/sourcetypes.
     * 
     */
    public Optional<Output<Map<String,String>>> config() {
        return Optional.ofNullable(this.config);
    }

    /**
     * The free form name of the data source.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The free form name of the data source.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The data sources type, listed in API endpoint &lt;https://api.nsone.net/v1/data/sourcetypes&gt;.
     * 
     */
    @Import(name="sourcetype")
    private @Nullable Output<String> sourcetype;

    /**
     * @return The data sources type, listed in API endpoint &lt;https://api.nsone.net/v1/data/sourcetypes&gt;.
     * 
     */
    public Optional<Output<String>> sourcetype() {
        return Optional.ofNullable(this.sourcetype);
    }

    private DataSourceState() {}

    private DataSourceState(DataSourceState $) {
        this.config = $.config;
        this.name = $.name;
        this.sourcetype = $.sourcetype;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DataSourceState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DataSourceState $;

        public Builder() {
            $ = new DataSourceState();
        }

        public Builder(DataSourceState defaults) {
            $ = new DataSourceState(Objects.requireNonNull(defaults));
        }

        /**
         * @param config The data source configuration, determined by its type,
         * matching the specification in `config` from /data/sourcetypes.
         * 
         * @return builder
         * 
         */
        public Builder config(@Nullable Output<Map<String,String>> config) {
            $.config = config;
            return this;
        }

        /**
         * @param config The data source configuration, determined by its type,
         * matching the specification in `config` from /data/sourcetypes.
         * 
         * @return builder
         * 
         */
        public Builder config(Map<String,String> config) {
            return config(Output.of(config));
        }

        /**
         * @param name The free form name of the data source.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The free form name of the data source.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param sourcetype The data sources type, listed in API endpoint &lt;https://api.nsone.net/v1/data/sourcetypes&gt;.
         * 
         * @return builder
         * 
         */
        public Builder sourcetype(@Nullable Output<String> sourcetype) {
            $.sourcetype = sourcetype;
            return this;
        }

        /**
         * @param sourcetype The data sources type, listed in API endpoint &lt;https://api.nsone.net/v1/data/sourcetypes&gt;.
         * 
         * @return builder
         * 
         */
        public Builder sourcetype(String sourcetype) {
            return sourcetype(Output.of(sourcetype));
        }

        public DataSourceState build() {
            return $;
        }
    }

}
