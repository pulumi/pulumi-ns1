// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.ns1.outputs.GetMonitoringRegionsRegion;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetMonitoringRegionsResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return A set of the available monitoring regions. Regions is
     * documented below.
     * 
     */
    private @Nullable List<GetMonitoringRegionsRegion> regions;

    private GetMonitoringRegionsResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return A set of the available monitoring regions. Regions is
     * documented below.
     * 
     */
    public List<GetMonitoringRegionsRegion> regions() {
        return this.regions == null ? List.of() : this.regions;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetMonitoringRegionsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private @Nullable List<GetMonitoringRegionsRegion> regions;
        public Builder() {}
        public Builder(GetMonitoringRegionsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.regions = defaults.regions;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetMonitoringRegionsResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder regions(@Nullable List<GetMonitoringRegionsRegion> regions) {

            this.regions = regions;
            return this;
        }
        public Builder regions(GetMonitoringRegionsRegion... regions) {
            return regions(List.of(regions));
        }
        public GetMonitoringRegionsResult build() {
            final var _resultValue = new GetMonitoringRegionsResult();
            _resultValue.id = id;
            _resultValue.regions = regions;
            return _resultValue;
        }
    }
}
