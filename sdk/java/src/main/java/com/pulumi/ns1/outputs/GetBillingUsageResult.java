// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.ns1.outputs.GetBillingUsageByNetwork;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetBillingUsageResult {
    /**
     * @return (Computed) A list of network-specific query data containing:
     * 
     */
    private List<GetBillingUsageByNetwork> byNetworks;
    /**
     * @return (Computed) The queries limit for the China network.
     * 
     */
    private Integer chinaQueriesLimit;
    /**
     * @return Clean queries for this day.
     * 
     */
    private Integer cleanQueries;
    /**
     * @return (Computed) Whether DDoS Protection is enabled.
     * 
     */
    private Boolean ddosProtectionEnabled;
    /**
     * @return DDoS queries for this day.
     * 
     */
    private Integer ddosQueries;
    /**
     * @return (Computed) The RUM decisions limit for this billing cycle.
     * 
     */
    private Integer decisionsLimit;
    /**
     * @return (Computed) The filter chains limit for this billing cycle.
     * 
     */
    private Integer filterChainsLimit;
    private Integer from;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return (Computed) Whether dedicated DNS usage counts towards managed DNS usage.
     * 
     */
    private Boolean includeDedicatedDnsNetworkInManagedDnsUsage;
    private String metricType;
    /**
     * @return (Computed) The monitoring jobs limit for this billing cycle.
     * 
     */
    private Integer monitorsLimit;
    /**
     * @return (Computed) Whether NXD Protection is enabled.
     * 
     */
    private Boolean nxdProtectionEnabled;
    /**
     * @return NXD responses for this day.
     * 
     */
    private Integer nxdResponses;
    /**
     * @return (Computed) The queries limit for this billing cycle.
     * 
     */
    private Integer queriesLimit;
    /**
     * @return (Computed) The records limit for this billing cycle.
     * 
     */
    private Integer recordsLimit;
    private Integer to;
    /**
     * @return (Computed) The total usage count for the metric. Available for `decisions`, `filter-chains`, `monitors`, and `records` metrics.
     * 
     */
    private Integer totalUsage;

    private GetBillingUsageResult() {}
    /**
     * @return (Computed) A list of network-specific query data containing:
     * 
     */
    public List<GetBillingUsageByNetwork> byNetworks() {
        return this.byNetworks;
    }
    /**
     * @return (Computed) The queries limit for the China network.
     * 
     */
    public Integer chinaQueriesLimit() {
        return this.chinaQueriesLimit;
    }
    /**
     * @return Clean queries for this day.
     * 
     */
    public Integer cleanQueries() {
        return this.cleanQueries;
    }
    /**
     * @return (Computed) Whether DDoS Protection is enabled.
     * 
     */
    public Boolean ddosProtectionEnabled() {
        return this.ddosProtectionEnabled;
    }
    /**
     * @return DDoS queries for this day.
     * 
     */
    public Integer ddosQueries() {
        return this.ddosQueries;
    }
    /**
     * @return (Computed) The RUM decisions limit for this billing cycle.
     * 
     */
    public Integer decisionsLimit() {
        return this.decisionsLimit;
    }
    /**
     * @return (Computed) The filter chains limit for this billing cycle.
     * 
     */
    public Integer filterChainsLimit() {
        return this.filterChainsLimit;
    }
    public Integer from() {
        return this.from;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return (Computed) Whether dedicated DNS usage counts towards managed DNS usage.
     * 
     */
    public Boolean includeDedicatedDnsNetworkInManagedDnsUsage() {
        return this.includeDedicatedDnsNetworkInManagedDnsUsage;
    }
    public String metricType() {
        return this.metricType;
    }
    /**
     * @return (Computed) The monitoring jobs limit for this billing cycle.
     * 
     */
    public Integer monitorsLimit() {
        return this.monitorsLimit;
    }
    /**
     * @return (Computed) Whether NXD Protection is enabled.
     * 
     */
    public Boolean nxdProtectionEnabled() {
        return this.nxdProtectionEnabled;
    }
    /**
     * @return NXD responses for this day.
     * 
     */
    public Integer nxdResponses() {
        return this.nxdResponses;
    }
    /**
     * @return (Computed) The queries limit for this billing cycle.
     * 
     */
    public Integer queriesLimit() {
        return this.queriesLimit;
    }
    /**
     * @return (Computed) The records limit for this billing cycle.
     * 
     */
    public Integer recordsLimit() {
        return this.recordsLimit;
    }
    public Integer to() {
        return this.to;
    }
    /**
     * @return (Computed) The total usage count for the metric. Available for `decisions`, `filter-chains`, `monitors`, and `records` metrics.
     * 
     */
    public Integer totalUsage() {
        return this.totalUsage;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBillingUsageResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetBillingUsageByNetwork> byNetworks;
        private Integer chinaQueriesLimit;
        private Integer cleanQueries;
        private Boolean ddosProtectionEnabled;
        private Integer ddosQueries;
        private Integer decisionsLimit;
        private Integer filterChainsLimit;
        private Integer from;
        private String id;
        private Boolean includeDedicatedDnsNetworkInManagedDnsUsage;
        private String metricType;
        private Integer monitorsLimit;
        private Boolean nxdProtectionEnabled;
        private Integer nxdResponses;
        private Integer queriesLimit;
        private Integer recordsLimit;
        private Integer to;
        private Integer totalUsage;
        public Builder() {}
        public Builder(GetBillingUsageResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.byNetworks = defaults.byNetworks;
    	      this.chinaQueriesLimit = defaults.chinaQueriesLimit;
    	      this.cleanQueries = defaults.cleanQueries;
    	      this.ddosProtectionEnabled = defaults.ddosProtectionEnabled;
    	      this.ddosQueries = defaults.ddosQueries;
    	      this.decisionsLimit = defaults.decisionsLimit;
    	      this.filterChainsLimit = defaults.filterChainsLimit;
    	      this.from = defaults.from;
    	      this.id = defaults.id;
    	      this.includeDedicatedDnsNetworkInManagedDnsUsage = defaults.includeDedicatedDnsNetworkInManagedDnsUsage;
    	      this.metricType = defaults.metricType;
    	      this.monitorsLimit = defaults.monitorsLimit;
    	      this.nxdProtectionEnabled = defaults.nxdProtectionEnabled;
    	      this.nxdResponses = defaults.nxdResponses;
    	      this.queriesLimit = defaults.queriesLimit;
    	      this.recordsLimit = defaults.recordsLimit;
    	      this.to = defaults.to;
    	      this.totalUsage = defaults.totalUsage;
        }

        @CustomType.Setter
        public Builder byNetworks(List<GetBillingUsageByNetwork> byNetworks) {
            if (byNetworks == null) {
              throw new MissingRequiredPropertyException("GetBillingUsageResult", "byNetworks");
            }
            this.byNetworks = byNetworks;
            return this;
        }
        public Builder byNetworks(GetBillingUsageByNetwork... byNetworks) {
            return byNetworks(List.of(byNetworks));
        }
        @CustomType.Setter
        public Builder chinaQueriesLimit(Integer chinaQueriesLimit) {
            if (chinaQueriesLimit == null) {
              throw new MissingRequiredPropertyException("GetBillingUsageResult", "chinaQueriesLimit");
            }
            this.chinaQueriesLimit = chinaQueriesLimit;
            return this;
        }
        @CustomType.Setter
        public Builder cleanQueries(Integer cleanQueries) {
            if (cleanQueries == null) {
              throw new MissingRequiredPropertyException("GetBillingUsageResult", "cleanQueries");
            }
            this.cleanQueries = cleanQueries;
            return this;
        }
        @CustomType.Setter
        public Builder ddosProtectionEnabled(Boolean ddosProtectionEnabled) {
            if (ddosProtectionEnabled == null) {
              throw new MissingRequiredPropertyException("GetBillingUsageResult", "ddosProtectionEnabled");
            }
            this.ddosProtectionEnabled = ddosProtectionEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder ddosQueries(Integer ddosQueries) {
            if (ddosQueries == null) {
              throw new MissingRequiredPropertyException("GetBillingUsageResult", "ddosQueries");
            }
            this.ddosQueries = ddosQueries;
            return this;
        }
        @CustomType.Setter
        public Builder decisionsLimit(Integer decisionsLimit) {
            if (decisionsLimit == null) {
              throw new MissingRequiredPropertyException("GetBillingUsageResult", "decisionsLimit");
            }
            this.decisionsLimit = decisionsLimit;
            return this;
        }
        @CustomType.Setter
        public Builder filterChainsLimit(Integer filterChainsLimit) {
            if (filterChainsLimit == null) {
              throw new MissingRequiredPropertyException("GetBillingUsageResult", "filterChainsLimit");
            }
            this.filterChainsLimit = filterChainsLimit;
            return this;
        }
        @CustomType.Setter
        public Builder from(Integer from) {
            if (from == null) {
              throw new MissingRequiredPropertyException("GetBillingUsageResult", "from");
            }
            this.from = from;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetBillingUsageResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder includeDedicatedDnsNetworkInManagedDnsUsage(Boolean includeDedicatedDnsNetworkInManagedDnsUsage) {
            if (includeDedicatedDnsNetworkInManagedDnsUsage == null) {
              throw new MissingRequiredPropertyException("GetBillingUsageResult", "includeDedicatedDnsNetworkInManagedDnsUsage");
            }
            this.includeDedicatedDnsNetworkInManagedDnsUsage = includeDedicatedDnsNetworkInManagedDnsUsage;
            return this;
        }
        @CustomType.Setter
        public Builder metricType(String metricType) {
            if (metricType == null) {
              throw new MissingRequiredPropertyException("GetBillingUsageResult", "metricType");
            }
            this.metricType = metricType;
            return this;
        }
        @CustomType.Setter
        public Builder monitorsLimit(Integer monitorsLimit) {
            if (monitorsLimit == null) {
              throw new MissingRequiredPropertyException("GetBillingUsageResult", "monitorsLimit");
            }
            this.monitorsLimit = monitorsLimit;
            return this;
        }
        @CustomType.Setter
        public Builder nxdProtectionEnabled(Boolean nxdProtectionEnabled) {
            if (nxdProtectionEnabled == null) {
              throw new MissingRequiredPropertyException("GetBillingUsageResult", "nxdProtectionEnabled");
            }
            this.nxdProtectionEnabled = nxdProtectionEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder nxdResponses(Integer nxdResponses) {
            if (nxdResponses == null) {
              throw new MissingRequiredPropertyException("GetBillingUsageResult", "nxdResponses");
            }
            this.nxdResponses = nxdResponses;
            return this;
        }
        @CustomType.Setter
        public Builder queriesLimit(Integer queriesLimit) {
            if (queriesLimit == null) {
              throw new MissingRequiredPropertyException("GetBillingUsageResult", "queriesLimit");
            }
            this.queriesLimit = queriesLimit;
            return this;
        }
        @CustomType.Setter
        public Builder recordsLimit(Integer recordsLimit) {
            if (recordsLimit == null) {
              throw new MissingRequiredPropertyException("GetBillingUsageResult", "recordsLimit");
            }
            this.recordsLimit = recordsLimit;
            return this;
        }
        @CustomType.Setter
        public Builder to(Integer to) {
            if (to == null) {
              throw new MissingRequiredPropertyException("GetBillingUsageResult", "to");
            }
            this.to = to;
            return this;
        }
        @CustomType.Setter
        public Builder totalUsage(Integer totalUsage) {
            if (totalUsage == null) {
              throw new MissingRequiredPropertyException("GetBillingUsageResult", "totalUsage");
            }
            this.totalUsage = totalUsage;
            return this;
        }
        public GetBillingUsageResult build() {
            final var _resultValue = new GetBillingUsageResult();
            _resultValue.byNetworks = byNetworks;
            _resultValue.chinaQueriesLimit = chinaQueriesLimit;
            _resultValue.cleanQueries = cleanQueries;
            _resultValue.ddosProtectionEnabled = ddosProtectionEnabled;
            _resultValue.ddosQueries = ddosQueries;
            _resultValue.decisionsLimit = decisionsLimit;
            _resultValue.filterChainsLimit = filterChainsLimit;
            _resultValue.from = from;
            _resultValue.id = id;
            _resultValue.includeDedicatedDnsNetworkInManagedDnsUsage = includeDedicatedDnsNetworkInManagedDnsUsage;
            _resultValue.metricType = metricType;
            _resultValue.monitorsLimit = monitorsLimit;
            _resultValue.nxdProtectionEnabled = nxdProtectionEnabled;
            _resultValue.nxdResponses = nxdResponses;
            _resultValue.queriesLimit = queriesLimit;
            _resultValue.recordsLimit = recordsLimit;
            _resultValue.to = to;
            _resultValue.totalUsage = totalUsage;
            return _resultValue;
        }
    }
}
