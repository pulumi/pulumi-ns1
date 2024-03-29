// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ApplicationDefaultConfig {
    /**
     * @return Indicates whether or not to use HTTP in measurements.
     * 
     */
    private Boolean http;
    /**
     * @return Indicates whether or not to use HTTPS in measurements.
     * 
     */
    private @Nullable Boolean https;
    /**
     * @return Maximum timeout per job
     * 0, the primary NSONE Global Network. Normally, you should not have to worry about this.
     * 
     */
    private @Nullable Integer jobTimeoutMillis;
    /**
     * @return Maximum timeout per request.
     * 
     */
    private @Nullable Integer requestTimeoutMillis;
    /**
     * @return Indicates whether or not to skip aggregation for this job&#39;s measurements
     * 
     */
    private @Nullable Boolean staticValues;
    /**
     * @return Whether to use XMLHttpRequest (XHR) when taking measurements.
     * 
     */
    private @Nullable Boolean useXhr;

    private ApplicationDefaultConfig() {}
    /**
     * @return Indicates whether or not to use HTTP in measurements.
     * 
     */
    public Boolean http() {
        return this.http;
    }
    /**
     * @return Indicates whether or not to use HTTPS in measurements.
     * 
     */
    public Optional<Boolean> https() {
        return Optional.ofNullable(this.https);
    }
    /**
     * @return Maximum timeout per job
     * 0, the primary NSONE Global Network. Normally, you should not have to worry about this.
     * 
     */
    public Optional<Integer> jobTimeoutMillis() {
        return Optional.ofNullable(this.jobTimeoutMillis);
    }
    /**
     * @return Maximum timeout per request.
     * 
     */
    public Optional<Integer> requestTimeoutMillis() {
        return Optional.ofNullable(this.requestTimeoutMillis);
    }
    /**
     * @return Indicates whether or not to skip aggregation for this job&#39;s measurements
     * 
     */
    public Optional<Boolean> staticValues() {
        return Optional.ofNullable(this.staticValues);
    }
    /**
     * @return Whether to use XMLHttpRequest (XHR) when taking measurements.
     * 
     */
    public Optional<Boolean> useXhr() {
        return Optional.ofNullable(this.useXhr);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ApplicationDefaultConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean http;
        private @Nullable Boolean https;
        private @Nullable Integer jobTimeoutMillis;
        private @Nullable Integer requestTimeoutMillis;
        private @Nullable Boolean staticValues;
        private @Nullable Boolean useXhr;
        public Builder() {}
        public Builder(ApplicationDefaultConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.http = defaults.http;
    	      this.https = defaults.https;
    	      this.jobTimeoutMillis = defaults.jobTimeoutMillis;
    	      this.requestTimeoutMillis = defaults.requestTimeoutMillis;
    	      this.staticValues = defaults.staticValues;
    	      this.useXhr = defaults.useXhr;
        }

        @CustomType.Setter
        public Builder http(Boolean http) {
            if (http == null) {
              throw new MissingRequiredPropertyException("ApplicationDefaultConfig", "http");
            }
            this.http = http;
            return this;
        }
        @CustomType.Setter
        public Builder https(@Nullable Boolean https) {

            this.https = https;
            return this;
        }
        @CustomType.Setter
        public Builder jobTimeoutMillis(@Nullable Integer jobTimeoutMillis) {

            this.jobTimeoutMillis = jobTimeoutMillis;
            return this;
        }
        @CustomType.Setter
        public Builder requestTimeoutMillis(@Nullable Integer requestTimeoutMillis) {

            this.requestTimeoutMillis = requestTimeoutMillis;
            return this;
        }
        @CustomType.Setter
        public Builder staticValues(@Nullable Boolean staticValues) {

            this.staticValues = staticValues;
            return this;
        }
        @CustomType.Setter
        public Builder useXhr(@Nullable Boolean useXhr) {

            this.useXhr = useXhr;
            return this;
        }
        public ApplicationDefaultConfig build() {
            final var _resultValue = new ApplicationDefaultConfig();
            _resultValue.http = http;
            _resultValue.https = https;
            _resultValue.jobTimeoutMillis = jobTimeoutMillis;
            _resultValue.requestTimeoutMillis = requestTimeoutMillis;
            _resultValue.staticValues = staticValues;
            _resultValue.useXhr = useXhr;
            return _resultValue;
        }
    }
}
