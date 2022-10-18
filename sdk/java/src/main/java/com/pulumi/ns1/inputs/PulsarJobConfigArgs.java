// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PulsarJobConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final PulsarJobConfigArgs Empty = new PulsarJobConfigArgs();

    @Import(name="host")
    private @Nullable Output<String> host;

    public Optional<Output<String>> host() {
        return Optional.ofNullable(this.host);
    }

    @Import(name="http")
    private @Nullable Output<Boolean> http;

    public Optional<Output<Boolean>> http() {
        return Optional.ofNullable(this.http);
    }

    @Import(name="https")
    private @Nullable Output<Boolean> https;

    public Optional<Output<Boolean>> https() {
        return Optional.ofNullable(this.https);
    }

    @Import(name="jobTimeoutMillis")
    private @Nullable Output<Integer> jobTimeoutMillis;

    public Optional<Output<Integer>> jobTimeoutMillis() {
        return Optional.ofNullable(this.jobTimeoutMillis);
    }

    @Import(name="requestTimeoutMillis")
    private @Nullable Output<Integer> requestTimeoutMillis;

    public Optional<Output<Integer>> requestTimeoutMillis() {
        return Optional.ofNullable(this.requestTimeoutMillis);
    }

    @Import(name="staticValues")
    private @Nullable Output<Boolean> staticValues;

    public Optional<Output<Boolean>> staticValues() {
        return Optional.ofNullable(this.staticValues);
    }

    @Import(name="urlPath")
    private @Nullable Output<String> urlPath;

    public Optional<Output<String>> urlPath() {
        return Optional.ofNullable(this.urlPath);
    }

    @Import(name="useXhr")
    private @Nullable Output<Boolean> useXhr;

    public Optional<Output<Boolean>> useXhr() {
        return Optional.ofNullable(this.useXhr);
    }

    private PulsarJobConfigArgs() {}

    private PulsarJobConfigArgs(PulsarJobConfigArgs $) {
        this.host = $.host;
        this.http = $.http;
        this.https = $.https;
        this.jobTimeoutMillis = $.jobTimeoutMillis;
        this.requestTimeoutMillis = $.requestTimeoutMillis;
        this.staticValues = $.staticValues;
        this.urlPath = $.urlPath;
        this.useXhr = $.useXhr;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PulsarJobConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PulsarJobConfigArgs $;

        public Builder() {
            $ = new PulsarJobConfigArgs();
        }

        public Builder(PulsarJobConfigArgs defaults) {
            $ = new PulsarJobConfigArgs(Objects.requireNonNull(defaults));
        }

        public Builder host(@Nullable Output<String> host) {
            $.host = host;
            return this;
        }

        public Builder host(String host) {
            return host(Output.of(host));
        }

        public Builder http(@Nullable Output<Boolean> http) {
            $.http = http;
            return this;
        }

        public Builder http(Boolean http) {
            return http(Output.of(http));
        }

        public Builder https(@Nullable Output<Boolean> https) {
            $.https = https;
            return this;
        }

        public Builder https(Boolean https) {
            return https(Output.of(https));
        }

        public Builder jobTimeoutMillis(@Nullable Output<Integer> jobTimeoutMillis) {
            $.jobTimeoutMillis = jobTimeoutMillis;
            return this;
        }

        public Builder jobTimeoutMillis(Integer jobTimeoutMillis) {
            return jobTimeoutMillis(Output.of(jobTimeoutMillis));
        }

        public Builder requestTimeoutMillis(@Nullable Output<Integer> requestTimeoutMillis) {
            $.requestTimeoutMillis = requestTimeoutMillis;
            return this;
        }

        public Builder requestTimeoutMillis(Integer requestTimeoutMillis) {
            return requestTimeoutMillis(Output.of(requestTimeoutMillis));
        }

        public Builder staticValues(@Nullable Output<Boolean> staticValues) {
            $.staticValues = staticValues;
            return this;
        }

        public Builder staticValues(Boolean staticValues) {
            return staticValues(Output.of(staticValues));
        }

        public Builder urlPath(@Nullable Output<String> urlPath) {
            $.urlPath = urlPath;
            return this;
        }

        public Builder urlPath(String urlPath) {
            return urlPath(Output.of(urlPath));
        }

        public Builder useXhr(@Nullable Output<Boolean> useXhr) {
            $.useXhr = useXhr;
            return this;
        }

        public Builder useXhr(Boolean useXhr) {
            return useXhr(Output.of(useXhr));
        }

        public PulsarJobConfigArgs build() {
            return $;
        }
    }

}
