// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetRecordRegion {
    /**
     * @return Map of metadata
     * 
     */
    private Map<String,Object> meta;
    private String name;

    private GetRecordRegion() {}
    /**
     * @return Map of metadata
     * 
     */
    public Map<String,Object> meta() {
        return this.meta;
    }
    public String name() {
        return this.name;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRecordRegion defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Map<String,Object> meta;
        private String name;
        public Builder() {}
        public Builder(GetRecordRegion defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.meta = defaults.meta;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder meta(Map<String,Object> meta) {
            this.meta = Objects.requireNonNull(meta);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public GetRecordRegion build() {
            final var o = new GetRecordRegion();
            o.meta = meta;
            o.name = name;
            return o;
        }
    }
}
