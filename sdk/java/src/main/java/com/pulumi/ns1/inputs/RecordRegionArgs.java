// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RecordRegionArgs extends com.pulumi.resources.ResourceArgs {

    public static final RecordRegionArgs Empty = new RecordRegionArgs();

    @Import(name="meta")
    private @Nullable Output<Map<String,String>> meta;

    public Optional<Output<Map<String,String>>> meta() {
        return Optional.ofNullable(this.meta);
    }

    /**
     * Name of the region (or Answer Group).
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Name of the region (or Answer Group).
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    private RecordRegionArgs() {}

    private RecordRegionArgs(RecordRegionArgs $) {
        this.meta = $.meta;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RecordRegionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RecordRegionArgs $;

        public Builder() {
            $ = new RecordRegionArgs();
        }

        public Builder(RecordRegionArgs defaults) {
            $ = new RecordRegionArgs(Objects.requireNonNull(defaults));
        }

        public Builder meta(@Nullable Output<Map<String,String>> meta) {
            $.meta = meta;
            return this;
        }

        public Builder meta(Map<String,String> meta) {
            return meta(Output.of(meta));
        }

        /**
         * @param name Name of the region (or Answer Group).
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the region (or Answer Group).
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public RecordRegionArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("RecordRegionArgs", "name");
            }
            return $;
        }
    }

}
