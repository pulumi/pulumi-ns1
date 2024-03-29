// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetMonitoringRegionsRegion {
    /**
     * @return 3-letter city code identifying the location of the monitor.
     * 
     */
    private @Nullable String code;
    /**
     * @return City name identifying the location of the monitor.
     * 
     */
    private @Nullable String name;
    /**
     * @return A list of IPv4 and IPv6 subnets the monitor sources requests from.
     * 
     */
    private @Nullable List<String> subnets;

    private GetMonitoringRegionsRegion() {}
    /**
     * @return 3-letter city code identifying the location of the monitor.
     * 
     */
    public Optional<String> code() {
        return Optional.ofNullable(this.code);
    }
    /**
     * @return City name identifying the location of the monitor.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return A list of IPv4 and IPv6 subnets the monitor sources requests from.
     * 
     */
    public List<String> subnets() {
        return this.subnets == null ? List.of() : this.subnets;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetMonitoringRegionsRegion defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String code;
        private @Nullable String name;
        private @Nullable List<String> subnets;
        public Builder() {}
        public Builder(GetMonitoringRegionsRegion defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.code = defaults.code;
    	      this.name = defaults.name;
    	      this.subnets = defaults.subnets;
        }

        @CustomType.Setter
        public Builder code(@Nullable String code) {

            this.code = code;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {

            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder subnets(@Nullable List<String> subnets) {

            this.subnets = subnets;
            return this;
        }
        public Builder subnets(String... subnets) {
            return subnets(List.of(subnets));
        }
        public GetMonitoringRegionsRegion build() {
            final var _resultValue = new GetMonitoringRegionsRegion();
            _resultValue.code = code;
            _resultValue.name = name;
            _resultValue.subnets = subnets;
            return _resultValue;
        }
    }
}
