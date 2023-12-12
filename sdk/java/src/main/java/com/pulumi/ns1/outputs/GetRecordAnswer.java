// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetRecordAnswer {
    private String answer;
    /**
     * @return Map of metadata
     * 
     */
    private Map<String,Object> meta;
    private String region;

    private GetRecordAnswer() {}
    public String answer() {
        return this.answer;
    }
    /**
     * @return Map of metadata
     * 
     */
    public Map<String,Object> meta() {
        return this.meta;
    }
    public String region() {
        return this.region;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRecordAnswer defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String answer;
        private Map<String,Object> meta;
        private String region;
        public Builder() {}
        public Builder(GetRecordAnswer defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.answer = defaults.answer;
    	      this.meta = defaults.meta;
    	      this.region = defaults.region;
        }

        @CustomType.Setter
        public Builder answer(String answer) {
            this.answer = Objects.requireNonNull(answer);
            return this;
        }
        @CustomType.Setter
        public Builder meta(Map<String,Object> meta) {
            this.meta = Objects.requireNonNull(meta);
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            this.region = Objects.requireNonNull(region);
            return this;
        }
        public GetRecordAnswer build() {
            final var _resultValue = new GetRecordAnswer();
            _resultValue.answer = answer;
            _resultValue.meta = meta;
            _resultValue.region = region;
            return _resultValue;
        }
    }
}
