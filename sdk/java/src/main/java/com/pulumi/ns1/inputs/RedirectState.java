// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RedirectState extends com.pulumi.resources.ResourceArgs {

    public static final RedirectState Empty = new RedirectState();

    @Import(name="certificateId")
    private @Nullable Output<String> certificateId;

    public Optional<Output<String>> certificateId() {
        return Optional.ofNullable(this.certificateId);
    }

    /**
     * The domain the redirect refers to.
     * 
     */
    @Import(name="domain")
    private @Nullable Output<String> domain;

    /**
     * @return The domain the redirect refers to.
     * 
     */
    public Optional<Output<String>> domain() {
        return Optional.ofNullable(this.domain);
    }

    /**
     * How the target is interpreted:
     * * __all__       appends the entire incoming path to the target destination;
     * * __capture__   appends only the part of the incoming path corresponding to the wildcard (*);
     * * __none__      does not append any part of the incoming path.
     * 
     */
    @Import(name="forwardingMode")
    private @Nullable Output<String> forwardingMode;

    /**
     * @return How the target is interpreted:
     * * __all__       appends the entire incoming path to the target destination;
     * * __capture__   appends only the part of the incoming path corresponding to the wildcard (*);
     * * __none__      does not append any part of the incoming path.
     * 
     */
    public Optional<Output<String>> forwardingMode() {
        return Optional.ofNullable(this.forwardingMode);
    }

    /**
     * How the redirect is executed:
     * * __permanent__ (HTTP 301) indicates to search engines that they should remove the old page from
     *   their database and replace it with the new target page (this is recommended for SEO);
     * * __temporary__ (HTTP 302) less common, indicates that search engines should keep the old domain or
     *   page indexed as the redirect is only temporary (while both pages might appear in the
     *   search results, a temporary redirect suggests to the search engine that it should
     *   prefer the new target page);
     * * __masking__   preserves the redirected domain in the browser&#39;s address bar (this lets users see the
     *   address they entered, even though the displayed content comes from a different web page).
     * 
     */
    @Import(name="forwardingType")
    private @Nullable Output<String> forwardingType;

    /**
     * @return How the redirect is executed:
     * * __permanent__ (HTTP 301) indicates to search engines that they should remove the old page from
     *   their database and replace it with the new target page (this is recommended for SEO);
     * * __temporary__ (HTTP 302) less common, indicates that search engines should keep the old domain or
     *   page indexed as the redirect is only temporary (while both pages might appear in the
     *   search results, a temporary redirect suggests to the search engine that it should
     *   prefer the new target page);
     * * __masking__   preserves the redirected domain in the browser&#39;s address bar (this lets users see the
     *   address they entered, even though the displayed content comes from a different web page).
     * 
     */
    public Optional<Output<String>> forwardingType() {
        return Optional.ofNullable(this.forwardingType);
    }

    /**
     * True if HTTPS is supported on the source domain by using Let&#39;s Encrypt certificates.
     * 
     */
    @Import(name="httpsEnabled")
    private @Nullable Output<Boolean> httpsEnabled;

    /**
     * @return True if HTTPS is supported on the source domain by using Let&#39;s Encrypt certificates.
     * 
     */
    public Optional<Output<Boolean>> httpsEnabled() {
        return Optional.ofNullable(this.httpsEnabled);
    }

    /**
     * Forces redirect for users that try to visit HTTP domain to HTTPS instead.
     * 
     */
    @Import(name="httpsForced")
    private @Nullable Output<Boolean> httpsForced;

    /**
     * @return Forces redirect for users that try to visit HTTP domain to HTTPS instead.
     * 
     */
    public Optional<Output<Boolean>> httpsForced() {
        return Optional.ofNullable(this.httpsForced);
    }

    /**
     * The Unix timestamp representing when the certificate was last signed.
     * 
     */
    @Import(name="lastUpdated")
    private @Nullable Output<Integer> lastUpdated;

    /**
     * @return The Unix timestamp representing when the certificate was last signed.
     * 
     */
    public Optional<Output<Integer>> lastUpdated() {
        return Optional.ofNullable(this.lastUpdated);
    }

    /**
     * The path on the domain to redirect from.
     * 
     */
    @Import(name="path")
    private @Nullable Output<String> path;

    /**
     * @return The path on the domain to redirect from.
     * 
     */
    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
    }

    /**
     * Enables the query string of a URL to be applied directly to the new target URL.
     * 
     */
    @Import(name="queryForwarding")
    private @Nullable Output<Boolean> queryForwarding;

    /**
     * @return Enables the query string of a URL to be applied directly to the new target URL.
     * 
     */
    public Optional<Output<Boolean>> queryForwarding() {
        return Optional.ofNullable(this.queryForwarding);
    }

    /**
     * Tags associated with the configuration.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return Tags associated with the configuration.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * The URL to redirect to.
     * 
     */
    @Import(name="target")
    private @Nullable Output<String> target;

    /**
     * @return The URL to redirect to.
     * 
     */
    public Optional<Output<String>> target() {
        return Optional.ofNullable(this.target);
    }

    private RedirectState() {}

    private RedirectState(RedirectState $) {
        this.certificateId = $.certificateId;
        this.domain = $.domain;
        this.forwardingMode = $.forwardingMode;
        this.forwardingType = $.forwardingType;
        this.httpsEnabled = $.httpsEnabled;
        this.httpsForced = $.httpsForced;
        this.lastUpdated = $.lastUpdated;
        this.path = $.path;
        this.queryForwarding = $.queryForwarding;
        this.tags = $.tags;
        this.target = $.target;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RedirectState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RedirectState $;

        public Builder() {
            $ = new RedirectState();
        }

        public Builder(RedirectState defaults) {
            $ = new RedirectState(Objects.requireNonNull(defaults));
        }

        public Builder certificateId(@Nullable Output<String> certificateId) {
            $.certificateId = certificateId;
            return this;
        }

        public Builder certificateId(String certificateId) {
            return certificateId(Output.of(certificateId));
        }

        /**
         * @param domain The domain the redirect refers to.
         * 
         * @return builder
         * 
         */
        public Builder domain(@Nullable Output<String> domain) {
            $.domain = domain;
            return this;
        }

        /**
         * @param domain The domain the redirect refers to.
         * 
         * @return builder
         * 
         */
        public Builder domain(String domain) {
            return domain(Output.of(domain));
        }

        /**
         * @param forwardingMode How the target is interpreted:
         * * __all__       appends the entire incoming path to the target destination;
         * * __capture__   appends only the part of the incoming path corresponding to the wildcard (*);
         * * __none__      does not append any part of the incoming path.
         * 
         * @return builder
         * 
         */
        public Builder forwardingMode(@Nullable Output<String> forwardingMode) {
            $.forwardingMode = forwardingMode;
            return this;
        }

        /**
         * @param forwardingMode How the target is interpreted:
         * * __all__       appends the entire incoming path to the target destination;
         * * __capture__   appends only the part of the incoming path corresponding to the wildcard (*);
         * * __none__      does not append any part of the incoming path.
         * 
         * @return builder
         * 
         */
        public Builder forwardingMode(String forwardingMode) {
            return forwardingMode(Output.of(forwardingMode));
        }

        /**
         * @param forwardingType How the redirect is executed:
         * * __permanent__ (HTTP 301) indicates to search engines that they should remove the old page from
         *   their database and replace it with the new target page (this is recommended for SEO);
         * * __temporary__ (HTTP 302) less common, indicates that search engines should keep the old domain or
         *   page indexed as the redirect is only temporary (while both pages might appear in the
         *   search results, a temporary redirect suggests to the search engine that it should
         *   prefer the new target page);
         * * __masking__   preserves the redirected domain in the browser&#39;s address bar (this lets users see the
         *   address they entered, even though the displayed content comes from a different web page).
         * 
         * @return builder
         * 
         */
        public Builder forwardingType(@Nullable Output<String> forwardingType) {
            $.forwardingType = forwardingType;
            return this;
        }

        /**
         * @param forwardingType How the redirect is executed:
         * * __permanent__ (HTTP 301) indicates to search engines that they should remove the old page from
         *   their database and replace it with the new target page (this is recommended for SEO);
         * * __temporary__ (HTTP 302) less common, indicates that search engines should keep the old domain or
         *   page indexed as the redirect is only temporary (while both pages might appear in the
         *   search results, a temporary redirect suggests to the search engine that it should
         *   prefer the new target page);
         * * __masking__   preserves the redirected domain in the browser&#39;s address bar (this lets users see the
         *   address they entered, even though the displayed content comes from a different web page).
         * 
         * @return builder
         * 
         */
        public Builder forwardingType(String forwardingType) {
            return forwardingType(Output.of(forwardingType));
        }

        /**
         * @param httpsEnabled True if HTTPS is supported on the source domain by using Let&#39;s Encrypt certificates.
         * 
         * @return builder
         * 
         */
        public Builder httpsEnabled(@Nullable Output<Boolean> httpsEnabled) {
            $.httpsEnabled = httpsEnabled;
            return this;
        }

        /**
         * @param httpsEnabled True if HTTPS is supported on the source domain by using Let&#39;s Encrypt certificates.
         * 
         * @return builder
         * 
         */
        public Builder httpsEnabled(Boolean httpsEnabled) {
            return httpsEnabled(Output.of(httpsEnabled));
        }

        /**
         * @param httpsForced Forces redirect for users that try to visit HTTP domain to HTTPS instead.
         * 
         * @return builder
         * 
         */
        public Builder httpsForced(@Nullable Output<Boolean> httpsForced) {
            $.httpsForced = httpsForced;
            return this;
        }

        /**
         * @param httpsForced Forces redirect for users that try to visit HTTP domain to HTTPS instead.
         * 
         * @return builder
         * 
         */
        public Builder httpsForced(Boolean httpsForced) {
            return httpsForced(Output.of(httpsForced));
        }

        /**
         * @param lastUpdated The Unix timestamp representing when the certificate was last signed.
         * 
         * @return builder
         * 
         */
        public Builder lastUpdated(@Nullable Output<Integer> lastUpdated) {
            $.lastUpdated = lastUpdated;
            return this;
        }

        /**
         * @param lastUpdated The Unix timestamp representing when the certificate was last signed.
         * 
         * @return builder
         * 
         */
        public Builder lastUpdated(Integer lastUpdated) {
            return lastUpdated(Output.of(lastUpdated));
        }

        /**
         * @param path The path on the domain to redirect from.
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path The path on the domain to redirect from.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        /**
         * @param queryForwarding Enables the query string of a URL to be applied directly to the new target URL.
         * 
         * @return builder
         * 
         */
        public Builder queryForwarding(@Nullable Output<Boolean> queryForwarding) {
            $.queryForwarding = queryForwarding;
            return this;
        }

        /**
         * @param queryForwarding Enables the query string of a URL to be applied directly to the new target URL.
         * 
         * @return builder
         * 
         */
        public Builder queryForwarding(Boolean queryForwarding) {
            return queryForwarding(Output.of(queryForwarding));
        }

        /**
         * @param tags Tags associated with the configuration.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Tags associated with the configuration.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags Tags associated with the configuration.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param target The URL to redirect to.
         * 
         * @return builder
         * 
         */
        public Builder target(@Nullable Output<String> target) {
            $.target = target;
            return this;
        }

        /**
         * @param target The URL to redirect to.
         * 
         * @return builder
         * 
         */
        public Builder target(String target) {
            return target(Output.of(target));
        }

        public RedirectState build() {
            return $;
        }
    }

}
