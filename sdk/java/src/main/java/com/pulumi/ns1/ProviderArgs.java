// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProviderArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderArgs Empty = new ProviderArgs();

    /**
     * The ns1 API key (required)
     * 
     */
    @Import(name="apikey")
    private @Nullable Output<String> apikey;

    /**
     * @return The ns1 API key (required)
     * 
     */
    public Optional<Output<String>> apikey() {
        return Optional.ofNullable(this.apikey);
    }

    /**
     * URL prefix (including version) for API calls
     * 
     */
    @Import(name="endpoint")
    private @Nullable Output<String> endpoint;

    /**
     * @return URL prefix (including version) for API calls
     * 
     */
    public Optional<Output<String>> endpoint() {
        return Optional.ofNullable(this.endpoint);
    }

    /**
     * Don&#39;t validate server SSL/TLS certificate
     * 
     */
    @Import(name="ignoreSsl", json=true)
    private @Nullable Output<Boolean> ignoreSsl;

    /**
     * @return Don&#39;t validate server SSL/TLS certificate
     * 
     */
    public Optional<Output<Boolean>> ignoreSsl() {
        return Optional.ofNullable(this.ignoreSsl);
    }

    /**
     * Tune response to rate limits, see docs
     * 
     */
    @Import(name="rateLimitParallelism", json=true)
    private @Nullable Output<Integer> rateLimitParallelism;

    /**
     * @return Tune response to rate limits, see docs
     * 
     */
    public Optional<Output<Integer>> rateLimitParallelism() {
        return Optional.ofNullable(this.rateLimitParallelism);
    }

    /**
     * Maximum retries for 50x errors (-1 to disable)
     * 
     */
    @Import(name="retryMax", json=true)
    private @Nullable Output<Integer> retryMax;

    /**
     * @return Maximum retries for 50x errors (-1 to disable)
     * 
     */
    public Optional<Output<Integer>> retryMax() {
        return Optional.ofNullable(this.retryMax);
    }

    /**
     * User-Agent string to use in NS1 API requests
     * 
     */
    @Import(name="userAgent")
    private @Nullable Output<String> userAgent;

    /**
     * @return User-Agent string to use in NS1 API requests
     * 
     */
    public Optional<Output<String>> userAgent() {
        return Optional.ofNullable(this.userAgent);
    }

    private ProviderArgs() {}

    private ProviderArgs(ProviderArgs $) {
        this.apikey = $.apikey;
        this.endpoint = $.endpoint;
        this.ignoreSsl = $.ignoreSsl;
        this.rateLimitParallelism = $.rateLimitParallelism;
        this.retryMax = $.retryMax;
        this.userAgent = $.userAgent;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderArgs $;

        public Builder() {
            $ = new ProviderArgs();
        }

        public Builder(ProviderArgs defaults) {
            $ = new ProviderArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param apikey The ns1 API key (required)
         * 
         * @return builder
         * 
         */
        public Builder apikey(@Nullable Output<String> apikey) {
            $.apikey = apikey;
            return this;
        }

        /**
         * @param apikey The ns1 API key (required)
         * 
         * @return builder
         * 
         */
        public Builder apikey(String apikey) {
            return apikey(Output.of(apikey));
        }

        /**
         * @param endpoint URL prefix (including version) for API calls
         * 
         * @return builder
         * 
         */
        public Builder endpoint(@Nullable Output<String> endpoint) {
            $.endpoint = endpoint;
            return this;
        }

        /**
         * @param endpoint URL prefix (including version) for API calls
         * 
         * @return builder
         * 
         */
        public Builder endpoint(String endpoint) {
            return endpoint(Output.of(endpoint));
        }

        /**
         * @param ignoreSsl Don&#39;t validate server SSL/TLS certificate
         * 
         * @return builder
         * 
         */
        public Builder ignoreSsl(@Nullable Output<Boolean> ignoreSsl) {
            $.ignoreSsl = ignoreSsl;
            return this;
        }

        /**
         * @param ignoreSsl Don&#39;t validate server SSL/TLS certificate
         * 
         * @return builder
         * 
         */
        public Builder ignoreSsl(Boolean ignoreSsl) {
            return ignoreSsl(Output.of(ignoreSsl));
        }

        /**
         * @param rateLimitParallelism Tune response to rate limits, see docs
         * 
         * @return builder
         * 
         */
        public Builder rateLimitParallelism(@Nullable Output<Integer> rateLimitParallelism) {
            $.rateLimitParallelism = rateLimitParallelism;
            return this;
        }

        /**
         * @param rateLimitParallelism Tune response to rate limits, see docs
         * 
         * @return builder
         * 
         */
        public Builder rateLimitParallelism(Integer rateLimitParallelism) {
            return rateLimitParallelism(Output.of(rateLimitParallelism));
        }

        /**
         * @param retryMax Maximum retries for 50x errors (-1 to disable)
         * 
         * @return builder
         * 
         */
        public Builder retryMax(@Nullable Output<Integer> retryMax) {
            $.retryMax = retryMax;
            return this;
        }

        /**
         * @param retryMax Maximum retries for 50x errors (-1 to disable)
         * 
         * @return builder
         * 
         */
        public Builder retryMax(Integer retryMax) {
            return retryMax(Output.of(retryMax));
        }

        /**
         * @param userAgent User-Agent string to use in NS1 API requests
         * 
         * @return builder
         * 
         */
        public Builder userAgent(@Nullable Output<String> userAgent) {
            $.userAgent = userAgent;
            return this;
        }

        /**
         * @param userAgent User-Agent string to use in NS1 API requests
         * 
         * @return builder
         * 
         */
        public Builder userAgent(String userAgent) {
            return userAgent(Output.of(userAgent));
        }

        public ProviderArgs build() {
            return $;
        }
    }

}
