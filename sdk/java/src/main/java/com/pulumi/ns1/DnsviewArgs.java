// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DnsviewArgs extends com.pulumi.resources.ResourceArgs {

    public static final DnsviewArgs Empty = new DnsviewArgs();

    @Import(name="name")
    private @Nullable Output<String> name;

    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    @Import(name="networks")
    private @Nullable Output<List<Integer>> networks;

    public Optional<Output<List<Integer>>> networks() {
        return Optional.ofNullable(this.networks);
    }

    @Import(name="preference")
    private @Nullable Output<Integer> preference;

    public Optional<Output<Integer>> preference() {
        return Optional.ofNullable(this.preference);
    }

    @Import(name="readAcls")
    private @Nullable Output<List<String>> readAcls;

    public Optional<Output<List<String>>> readAcls() {
        return Optional.ofNullable(this.readAcls);
    }

    @Import(name="updateAcls")
    private @Nullable Output<List<String>> updateAcls;

    public Optional<Output<List<String>>> updateAcls() {
        return Optional.ofNullable(this.updateAcls);
    }

    @Import(name="zones")
    private @Nullable Output<List<String>> zones;

    public Optional<Output<List<String>>> zones() {
        return Optional.ofNullable(this.zones);
    }

    private DnsviewArgs() {}

    private DnsviewArgs(DnsviewArgs $) {
        this.name = $.name;
        this.networks = $.networks;
        this.preference = $.preference;
        this.readAcls = $.readAcls;
        this.updateAcls = $.updateAcls;
        this.zones = $.zones;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DnsviewArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DnsviewArgs $;

        public Builder() {
            $ = new DnsviewArgs();
        }

        public Builder(DnsviewArgs defaults) {
            $ = new DnsviewArgs(Objects.requireNonNull(defaults));
        }

        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder networks(@Nullable Output<List<Integer>> networks) {
            $.networks = networks;
            return this;
        }

        public Builder networks(List<Integer> networks) {
            return networks(Output.of(networks));
        }

        public Builder networks(Integer... networks) {
            return networks(List.of(networks));
        }

        public Builder preference(@Nullable Output<Integer> preference) {
            $.preference = preference;
            return this;
        }

        public Builder preference(Integer preference) {
            return preference(Output.of(preference));
        }

        public Builder readAcls(@Nullable Output<List<String>> readAcls) {
            $.readAcls = readAcls;
            return this;
        }

        public Builder readAcls(List<String> readAcls) {
            return readAcls(Output.of(readAcls));
        }

        public Builder readAcls(String... readAcls) {
            return readAcls(List.of(readAcls));
        }

        public Builder updateAcls(@Nullable Output<List<String>> updateAcls) {
            $.updateAcls = updateAcls;
            return this;
        }

        public Builder updateAcls(List<String> updateAcls) {
            return updateAcls(Output.of(updateAcls));
        }

        public Builder updateAcls(String... updateAcls) {
            return updateAcls(List.of(updateAcls));
        }

        public Builder zones(@Nullable Output<List<String>> zones) {
            $.zones = zones;
            return this;
        }

        public Builder zones(List<String> zones) {
            return zones(Output.of(zones));
        }

        public Builder zones(String... zones) {
            return zones(List.of(zones));
        }

        public DnsviewArgs build() {
            return $;
        }
    }

}