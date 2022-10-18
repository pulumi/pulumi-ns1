// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RecordFilter {
    /**
     * @return The filters&#39; configuration. Simple key/value pairs
     * determined by the filter type.
     * 
     */
    private @Nullable Map<String,Object> config;
    /**
     * @return Determines whether the filter is applied in the
     * filter chain.
     * 
     */
    private @Nullable Boolean disabled;
    /**
     * @return The type of filter.
     * 
     */
    private String filter;

    private RecordFilter() {}
    /**
     * @return The filters&#39; configuration. Simple key/value pairs
     * determined by the filter type.
     * 
     */
    public Map<String,Object> config() {
        return this.config == null ? Map.of() : this.config;
    }
    /**
     * @return Determines whether the filter is applied in the
     * filter chain.
     * 
     */
    public Optional<Boolean> disabled() {
        return Optional.ofNullable(this.disabled);
    }
    /**
     * @return The type of filter.
     * 
     */
    public String filter() {
        return this.filter;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RecordFilter defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Map<String,Object> config;
        private @Nullable Boolean disabled;
        private String filter;
        public Builder() {}
        public Builder(RecordFilter defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.config = defaults.config;
    	      this.disabled = defaults.disabled;
    	      this.filter = defaults.filter;
        }

        @CustomType.Setter
        public Builder config(@Nullable Map<String,Object> config) {
            this.config = config;
            return this;
        }
        @CustomType.Setter
        public Builder disabled(@Nullable Boolean disabled) {
            this.disabled = disabled;
            return this;
        }
        @CustomType.Setter
        public Builder filter(String filter) {
            this.filter = Objects.requireNonNull(filter);
            return this;
        }
        public RecordFilter build() {
            final var o = new RecordFilter();
            o.config = config;
            o.disabled = disabled;
            o.filter = filter;
            return o;
        }
    }
}
