// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RecordAnswer {
    /**
     * @return Space delimited string of RDATA fields dependent on the record type.
     * 
     */
    private @Nullable String answer;
    private @Nullable Map<String,Object> meta;
    /**
     * @return The region (Answer Group really) that this answer
     * belongs to. This should be one of the names specified in `regions`. Only a
     * single `region` per answer is currently supported. If you want an answer in
     * multiple regions, duplicating the answer (including metadata) is the correct
     * approach.
     * * `  meta ` - (Optional) meta is supported at the `answer` level. Meta
     *   is documented below.
     * 
     */
    private @Nullable String region;

    private RecordAnswer() {}
    /**
     * @return Space delimited string of RDATA fields dependent on the record type.
     * 
     */
    public Optional<String> answer() {
        return Optional.ofNullable(this.answer);
    }
    public Map<String,Object> meta() {
        return this.meta == null ? Map.of() : this.meta;
    }
    /**
     * @return The region (Answer Group really) that this answer
     * belongs to. This should be one of the names specified in `regions`. Only a
     * single `region` per answer is currently supported. If you want an answer in
     * multiple regions, duplicating the answer (including metadata) is the correct
     * approach.
     * * `  meta ` - (Optional) meta is supported at the `answer` level. Meta
     *   is documented below.
     * 
     */
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RecordAnswer defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String answer;
        private @Nullable Map<String,Object> meta;
        private @Nullable String region;
        public Builder() {}
        public Builder(RecordAnswer defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.answer = defaults.answer;
    	      this.meta = defaults.meta;
    	      this.region = defaults.region;
        }

        @CustomType.Setter
        public Builder answer(@Nullable String answer) {
            this.answer = answer;
            return this;
        }
        @CustomType.Setter
        public Builder meta(@Nullable Map<String,Object> meta) {
            this.meta = meta;
            return this;
        }
        @CustomType.Setter
        public Builder region(@Nullable String region) {
            this.region = region;
            return this;
        }
        public RecordAnswer build() {
            final var o = new RecordAnswer();
            o.answer = answer;
            o.meta = meta;
            o.region = region;
            return o;
        }
    }
}
