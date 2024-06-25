// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ZoneSecondaryArgs extends com.pulumi.resources.ResourceArgs {

    public static final ZoneSecondaryArgs Empty = new ZoneSecondaryArgs();

    /**
     * IPv4 address of the secondary server.
     * 
     */
    @Import(name="ip", required=true)
    private Output<String> ip;

    /**
     * @return IPv4 address of the secondary server.
     * 
     */
    public Output<String> ip() {
        return this.ip;
    }

    /**
     * List of network IDs (`int`) for which the zone
     * should be made available. Default is network 0, the primary NSONE Global
     * Network. Normally, you should not have to worry about this.
     * 
     */
    @Import(name="networks")
    private @Nullable Output<List<Integer>> networks;

    /**
     * @return List of network IDs (`int`) for which the zone
     * should be made available. Default is network 0, the primary NSONE Global
     * Network. Normally, you should not have to worry about this.
     * 
     */
    public Optional<Output<List<Integer>>> networks() {
        return Optional.ofNullable(this.networks);
    }

    /**
     * Whether we send `NOTIFY` messages to the secondary host
     * when the zone changes. Default `false`.
     * 
     */
    @Import(name="notify")
    private @Nullable Output<Boolean> notify;

    /**
     * @return Whether we send `NOTIFY` messages to the secondary host
     * when the zone changes. Default `false`.
     * 
     */
    public Optional<Output<Boolean>> notify_() {
        return Optional.ofNullable(this.notify);
    }

    /**
     * Port of the the secondary server. Default `53`.
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return Port of the the secondary server. Default `53`.
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    private ZoneSecondaryArgs() {}

    private ZoneSecondaryArgs(ZoneSecondaryArgs $) {
        this.ip = $.ip;
        this.networks = $.networks;
        this.notify = $.notify;
        this.port = $.port;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ZoneSecondaryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ZoneSecondaryArgs $;

        public Builder() {
            $ = new ZoneSecondaryArgs();
        }

        public Builder(ZoneSecondaryArgs defaults) {
            $ = new ZoneSecondaryArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ip IPv4 address of the secondary server.
         * 
         * @return builder
         * 
         */
        public Builder ip(Output<String> ip) {
            $.ip = ip;
            return this;
        }

        /**
         * @param ip IPv4 address of the secondary server.
         * 
         * @return builder
         * 
         */
        public Builder ip(String ip) {
            return ip(Output.of(ip));
        }

        /**
         * @param networks List of network IDs (`int`) for which the zone
         * should be made available. Default is network 0, the primary NSONE Global
         * Network. Normally, you should not have to worry about this.
         * 
         * @return builder
         * 
         */
        public Builder networks(@Nullable Output<List<Integer>> networks) {
            $.networks = networks;
            return this;
        }

        /**
         * @param networks List of network IDs (`int`) for which the zone
         * should be made available. Default is network 0, the primary NSONE Global
         * Network. Normally, you should not have to worry about this.
         * 
         * @return builder
         * 
         */
        public Builder networks(List<Integer> networks) {
            return networks(Output.of(networks));
        }

        /**
         * @param networks List of network IDs (`int`) for which the zone
         * should be made available. Default is network 0, the primary NSONE Global
         * Network. Normally, you should not have to worry about this.
         * 
         * @return builder
         * 
         */
        public Builder networks(Integer... networks) {
            return networks(List.of(networks));
        }

        /**
         * @param notify Whether we send `NOTIFY` messages to the secondary host
         * when the zone changes. Default `false`.
         * 
         * @return builder
         * 
         */
        public Builder notify_(@Nullable Output<Boolean> notify) {
            $.notify = notify;
            return this;
        }

        /**
         * @param notify Whether we send `NOTIFY` messages to the secondary host
         * when the zone changes. Default `false`.
         * 
         * @return builder
         * 
         */
        public Builder notify_(Boolean notify) {
            return notify_(Output.of(notify));
        }

        /**
         * @param port Port of the the secondary server. Default `53`.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port Port of the the secondary server. Default `53`.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        public ZoneSecondaryArgs build() {
            if ($.ip == null) {
                throw new MissingRequiredPropertyException("ZoneSecondaryArgs", "ip");
            }
            return $;
        }
    }

}