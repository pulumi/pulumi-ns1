// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.ns1.inputs.APIKeyDnsRecordsAllowArgs;
import com.pulumi.ns1.inputs.APIKeyDnsRecordsDenyArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class APIKeyArgs extends com.pulumi.resources.ResourceArgs {

    public static final APIKeyArgs Empty = new APIKeyArgs();

    /**
     * Whether the apikey can modify account settings.
     * 
     */
    @Import(name="accountManageAccountSettings")
    private @Nullable Output<Boolean> accountManageAccountSettings;

    /**
     * @return Whether the apikey can modify account settings.
     * 
     */
    public Optional<Output<Boolean>> accountManageAccountSettings() {
        return Optional.ofNullable(this.accountManageAccountSettings);
    }

    /**
     * Whether the apikey can modify account apikeys.
     * 
     */
    @Import(name="accountManageApikeys")
    private @Nullable Output<Boolean> accountManageApikeys;

    /**
     * @return Whether the apikey can modify account apikeys.
     * 
     */
    public Optional<Output<Boolean>> accountManageApikeys() {
        return Optional.ofNullable(this.accountManageApikeys);
    }

    /**
     * Whether the apikey can manage ip whitelist.
     * 
     */
    @Import(name="accountManageIpWhitelist")
    private @Nullable Output<Boolean> accountManageIpWhitelist;

    /**
     * @return Whether the apikey can manage ip whitelist.
     * 
     */
    public Optional<Output<Boolean>> accountManageIpWhitelist() {
        return Optional.ofNullable(this.accountManageIpWhitelist);
    }

    /**
     * Whether the apikey can modify account payment methods.
     * 
     */
    @Import(name="accountManagePaymentMethods")
    private @Nullable Output<Boolean> accountManagePaymentMethods;

    /**
     * @return Whether the apikey can modify account payment methods.
     * 
     */
    public Optional<Output<Boolean>> accountManagePaymentMethods() {
        return Optional.ofNullable(this.accountManagePaymentMethods);
    }

    /**
     * No longer in use.
     * 
     * @deprecated
     * obsolete, should no longer be used
     * 
     */
    @Deprecated /* obsolete, should no longer be used */
    @Import(name="accountManagePlan")
    private @Nullable Output<Boolean> accountManagePlan;

    /**
     * @return No longer in use.
     * 
     * @deprecated
     * obsolete, should no longer be used
     * 
     */
    @Deprecated /* obsolete, should no longer be used */
    public Optional<Output<Boolean>> accountManagePlan() {
        return Optional.ofNullable(this.accountManagePlan);
    }

    /**
     * Whether the apikey can modify other teams in the account.
     * 
     */
    @Import(name="accountManageTeams")
    private @Nullable Output<Boolean> accountManageTeams;

    /**
     * @return Whether the apikey can modify other teams in the account.
     * 
     */
    public Optional<Output<Boolean>> accountManageTeams() {
        return Optional.ofNullable(this.accountManageTeams);
    }

    /**
     * Whether the apikey can modify account users.
     * 
     */
    @Import(name="accountManageUsers")
    private @Nullable Output<Boolean> accountManageUsers;

    /**
     * @return Whether the apikey can modify account users.
     * 
     */
    public Optional<Output<Boolean>> accountManageUsers() {
        return Optional.ofNullable(this.accountManageUsers);
    }

    /**
     * Whether the apikey can view activity logs.
     * 
     */
    @Import(name="accountViewActivityLog")
    private @Nullable Output<Boolean> accountViewActivityLog;

    /**
     * @return Whether the apikey can view activity logs.
     * 
     */
    public Optional<Output<Boolean>> accountViewActivityLog() {
        return Optional.ofNullable(this.accountViewActivityLog);
    }

    /**
     * Whether the apikey can view invoices.
     * 
     */
    @Import(name="accountViewInvoices")
    private @Nullable Output<Boolean> accountViewInvoices;

    /**
     * @return Whether the apikey can view invoices.
     * 
     */
    public Optional<Output<Boolean>> accountViewInvoices() {
        return Optional.ofNullable(this.accountViewInvoices);
    }

    /**
     * Whether the apikey can modify data feeds.
     * 
     */
    @Import(name="dataManageDatafeeds")
    private @Nullable Output<Boolean> dataManageDatafeeds;

    /**
     * @return Whether the apikey can modify data feeds.
     * 
     */
    public Optional<Output<Boolean>> dataManageDatafeeds() {
        return Optional.ofNullable(this.dataManageDatafeeds);
    }

    /**
     * Whether the apikey can modify data sources.
     * 
     */
    @Import(name="dataManageDatasources")
    private @Nullable Output<Boolean> dataManageDatasources;

    /**
     * @return Whether the apikey can modify data sources.
     * 
     */
    public Optional<Output<Boolean>> dataManageDatasources() {
        return Optional.ofNullable(this.dataManageDatasources);
    }

    /**
     * Whether the apikey can publish to data feeds.
     * 
     */
    @Import(name="dataPushToDatafeeds")
    private @Nullable Output<Boolean> dataPushToDatafeeds;

    /**
     * @return Whether the apikey can publish to data feeds.
     * 
     */
    public Optional<Output<Boolean>> dataPushToDatafeeds() {
        return Optional.ofNullable(this.dataPushToDatafeeds);
    }

    /**
     * Whether the apikey can manage DHCP.
     * Only relevant for the DDI product.
     * 
     */
    @Import(name="dhcpManageDhcp")
    private @Nullable Output<Boolean> dhcpManageDhcp;

    /**
     * @return Whether the apikey can manage DHCP.
     * Only relevant for the DDI product.
     * 
     */
    public Optional<Output<Boolean>> dhcpManageDhcp() {
        return Optional.ofNullable(this.dhcpManageDhcp);
    }

    /**
     * Whether the apikey can view DHCP.
     * Only relevant for the DDI product.
     * 
     */
    @Import(name="dhcpViewDhcp")
    private @Nullable Output<Boolean> dhcpViewDhcp;

    /**
     * @return Whether the apikey can view DHCP.
     * Only relevant for the DDI product.
     * 
     */
    public Optional<Output<Boolean>> dhcpViewDhcp() {
        return Optional.ofNullable(this.dhcpViewDhcp);
    }

    /**
     * Whether the apikey can modify the accounts zones.
     * 
     */
    @Import(name="dnsManageZones")
    private @Nullable Output<Boolean> dnsManageZones;

    /**
     * @return Whether the apikey can modify the accounts zones.
     * 
     */
    public Optional<Output<Boolean>> dnsManageZones() {
        return Optional.ofNullable(this.dnsManageZones);
    }

    /**
     * List of records that the apikey may access.
     * 
     */
    @Import(name="dnsRecordsAllows")
    private @Nullable Output<List<APIKeyDnsRecordsAllowArgs>> dnsRecordsAllows;

    /**
     * @return List of records that the apikey may access.
     * 
     */
    public Optional<Output<List<APIKeyDnsRecordsAllowArgs>>> dnsRecordsAllows() {
        return Optional.ofNullable(this.dnsRecordsAllows);
    }

    /**
     * List of records that the apikey may not access.
     * 
     */
    @Import(name="dnsRecordsDenies")
    private @Nullable Output<List<APIKeyDnsRecordsDenyArgs>> dnsRecordsDenies;

    /**
     * @return List of records that the apikey may not access.
     * 
     */
    public Optional<Output<List<APIKeyDnsRecordsDenyArgs>>> dnsRecordsDenies() {
        return Optional.ofNullable(this.dnsRecordsDenies);
    }

    /**
     * Whether the apikey can view the accounts zones.
     * 
     */
    @Import(name="dnsViewZones")
    private @Nullable Output<Boolean> dnsViewZones;

    /**
     * @return Whether the apikey can view the accounts zones.
     * 
     */
    public Optional<Output<Boolean>> dnsViewZones() {
        return Optional.ofNullable(this.dnsViewZones);
    }

    /**
     * If true, enable the `dns_zones_allow` list, otherwise enable the `dns_zones_deny` list.
     * 
     */
    @Import(name="dnsZonesAllowByDefault")
    private @Nullable Output<Boolean> dnsZonesAllowByDefault;

    /**
     * @return If true, enable the `dns_zones_allow` list, otherwise enable the `dns_zones_deny` list.
     * 
     */
    public Optional<Output<Boolean>> dnsZonesAllowByDefault() {
        return Optional.ofNullable(this.dnsZonesAllowByDefault);
    }

    /**
     * List of zones that the apikey may access.
     * 
     */
    @Import(name="dnsZonesAllows")
    private @Nullable Output<List<String>> dnsZonesAllows;

    /**
     * @return List of zones that the apikey may access.
     * 
     */
    public Optional<Output<List<String>>> dnsZonesAllows() {
        return Optional.ofNullable(this.dnsZonesAllows);
    }

    /**
     * List of zones that the apikey may not access.
     * 
     */
    @Import(name="dnsZonesDenies")
    private @Nullable Output<List<String>> dnsZonesDenies;

    /**
     * @return List of zones that the apikey may not access.
     * 
     */
    public Optional<Output<List<String>>> dnsZonesDenies() {
        return Optional.ofNullable(this.dnsZonesDenies);
    }

    /**
     * Set to true to restrict access to only those IP addresses and networks listed in the **ip_whitelist** field.
     * 
     */
    @Import(name="ipWhitelistStrict")
    private @Nullable Output<Boolean> ipWhitelistStrict;

    /**
     * @return Set to true to restrict access to only those IP addresses and networks listed in the **ip_whitelist** field.
     * 
     */
    public Optional<Output<Boolean>> ipWhitelistStrict() {
        return Optional.ofNullable(this.ipWhitelistStrict);
    }

    /**
     * Array of IP addresses/networks to which to grant the API key access.
     * 
     */
    @Import(name="ipWhitelists")
    private @Nullable Output<List<String>> ipWhitelists;

    /**
     * @return Array of IP addresses/networks to which to grant the API key access.
     * 
     */
    public Optional<Output<List<String>>> ipWhitelists() {
        return Optional.ofNullable(this.ipWhitelists);
    }

    /**
     * Whether the apikey can manage IPAM.
     * Only relevant for the DDI product.
     * 
     */
    @Import(name="ipamManageIpam")
    private @Nullable Output<Boolean> ipamManageIpam;

    /**
     * @return Whether the apikey can manage IPAM.
     * Only relevant for the DDI product.
     * 
     */
    public Optional<Output<Boolean>> ipamManageIpam() {
        return Optional.ofNullable(this.ipamManageIpam);
    }

    /**
     * Whether the apikey can view IPAM.
     * Only relevant for the DDI product.
     * 
     */
    @Import(name="ipamViewIpam")
    private @Nullable Output<Boolean> ipamViewIpam;

    /**
     * @return Whether the apikey can view IPAM.
     * Only relevant for the DDI product.
     * 
     */
    public Optional<Output<Boolean>> ipamViewIpam() {
        return Optional.ofNullable(this.ipamViewIpam);
    }

    /**
     * Whether the apikey can modify monitoring jobs.
     * 
     */
    @Import(name="monitoringManageJobs")
    private @Nullable Output<Boolean> monitoringManageJobs;

    /**
     * @return Whether the apikey can modify monitoring jobs.
     * 
     */
    public Optional<Output<Boolean>> monitoringManageJobs() {
        return Optional.ofNullable(this.monitoringManageJobs);
    }

    /**
     * Whether the apikey can modify notification lists.
     * 
     */
    @Import(name="monitoringManageLists")
    private @Nullable Output<Boolean> monitoringManageLists;

    /**
     * @return Whether the apikey can modify notification lists.
     * 
     */
    public Optional<Output<Boolean>> monitoringManageLists() {
        return Optional.ofNullable(this.monitoringManageLists);
    }

    /**
     * Whether the apikey can view monitoring jobs.
     * 
     */
    @Import(name="monitoringViewJobs")
    private @Nullable Output<Boolean> monitoringViewJobs;

    /**
     * @return Whether the apikey can view monitoring jobs.
     * 
     */
    public Optional<Output<Boolean>> monitoringViewJobs() {
        return Optional.ofNullable(this.monitoringViewJobs);
    }

    /**
     * The free form name of the apikey.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The free form name of the apikey.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Whether the apikey can manage global active directory.
     * Only relevant for the DDI product.
     * 
     */
    @Import(name="securityManageActiveDirectory")
    private @Nullable Output<Boolean> securityManageActiveDirectory;

    /**
     * @return Whether the apikey can manage global active directory.
     * Only relevant for the DDI product.
     * 
     */
    public Optional<Output<Boolean>> securityManageActiveDirectory() {
        return Optional.ofNullable(this.securityManageActiveDirectory);
    }

    /**
     * Whether the apikey can manage global two factor authentication.
     * 
     */
    @Import(name="securityManageGlobal2fa")
    private @Nullable Output<Boolean> securityManageGlobal2fa;

    /**
     * @return Whether the apikey can manage global two factor authentication.
     * 
     */
    public Optional<Output<Boolean>> securityManageGlobal2fa() {
        return Optional.ofNullable(this.securityManageGlobal2fa);
    }

    /**
     * The teams that the apikey belongs to.
     * 
     */
    @Import(name="teams")
    private @Nullable Output<List<String>> teams;

    /**
     * @return The teams that the apikey belongs to.
     * 
     */
    public Optional<Output<List<String>>> teams() {
        return Optional.ofNullable(this.teams);
    }

    private APIKeyArgs() {}

    private APIKeyArgs(APIKeyArgs $) {
        this.accountManageAccountSettings = $.accountManageAccountSettings;
        this.accountManageApikeys = $.accountManageApikeys;
        this.accountManageIpWhitelist = $.accountManageIpWhitelist;
        this.accountManagePaymentMethods = $.accountManagePaymentMethods;
        this.accountManagePlan = $.accountManagePlan;
        this.accountManageTeams = $.accountManageTeams;
        this.accountManageUsers = $.accountManageUsers;
        this.accountViewActivityLog = $.accountViewActivityLog;
        this.accountViewInvoices = $.accountViewInvoices;
        this.dataManageDatafeeds = $.dataManageDatafeeds;
        this.dataManageDatasources = $.dataManageDatasources;
        this.dataPushToDatafeeds = $.dataPushToDatafeeds;
        this.dhcpManageDhcp = $.dhcpManageDhcp;
        this.dhcpViewDhcp = $.dhcpViewDhcp;
        this.dnsManageZones = $.dnsManageZones;
        this.dnsRecordsAllows = $.dnsRecordsAllows;
        this.dnsRecordsDenies = $.dnsRecordsDenies;
        this.dnsViewZones = $.dnsViewZones;
        this.dnsZonesAllowByDefault = $.dnsZonesAllowByDefault;
        this.dnsZonesAllows = $.dnsZonesAllows;
        this.dnsZonesDenies = $.dnsZonesDenies;
        this.ipWhitelistStrict = $.ipWhitelistStrict;
        this.ipWhitelists = $.ipWhitelists;
        this.ipamManageIpam = $.ipamManageIpam;
        this.ipamViewIpam = $.ipamViewIpam;
        this.monitoringManageJobs = $.monitoringManageJobs;
        this.monitoringManageLists = $.monitoringManageLists;
        this.monitoringViewJobs = $.monitoringViewJobs;
        this.name = $.name;
        this.securityManageActiveDirectory = $.securityManageActiveDirectory;
        this.securityManageGlobal2fa = $.securityManageGlobal2fa;
        this.teams = $.teams;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(APIKeyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private APIKeyArgs $;

        public Builder() {
            $ = new APIKeyArgs();
        }

        public Builder(APIKeyArgs defaults) {
            $ = new APIKeyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accountManageAccountSettings Whether the apikey can modify account settings.
         * 
         * @return builder
         * 
         */
        public Builder accountManageAccountSettings(@Nullable Output<Boolean> accountManageAccountSettings) {
            $.accountManageAccountSettings = accountManageAccountSettings;
            return this;
        }

        /**
         * @param accountManageAccountSettings Whether the apikey can modify account settings.
         * 
         * @return builder
         * 
         */
        public Builder accountManageAccountSettings(Boolean accountManageAccountSettings) {
            return accountManageAccountSettings(Output.of(accountManageAccountSettings));
        }

        /**
         * @param accountManageApikeys Whether the apikey can modify account apikeys.
         * 
         * @return builder
         * 
         */
        public Builder accountManageApikeys(@Nullable Output<Boolean> accountManageApikeys) {
            $.accountManageApikeys = accountManageApikeys;
            return this;
        }

        /**
         * @param accountManageApikeys Whether the apikey can modify account apikeys.
         * 
         * @return builder
         * 
         */
        public Builder accountManageApikeys(Boolean accountManageApikeys) {
            return accountManageApikeys(Output.of(accountManageApikeys));
        }

        /**
         * @param accountManageIpWhitelist Whether the apikey can manage ip whitelist.
         * 
         * @return builder
         * 
         */
        public Builder accountManageIpWhitelist(@Nullable Output<Boolean> accountManageIpWhitelist) {
            $.accountManageIpWhitelist = accountManageIpWhitelist;
            return this;
        }

        /**
         * @param accountManageIpWhitelist Whether the apikey can manage ip whitelist.
         * 
         * @return builder
         * 
         */
        public Builder accountManageIpWhitelist(Boolean accountManageIpWhitelist) {
            return accountManageIpWhitelist(Output.of(accountManageIpWhitelist));
        }

        /**
         * @param accountManagePaymentMethods Whether the apikey can modify account payment methods.
         * 
         * @return builder
         * 
         */
        public Builder accountManagePaymentMethods(@Nullable Output<Boolean> accountManagePaymentMethods) {
            $.accountManagePaymentMethods = accountManagePaymentMethods;
            return this;
        }

        /**
         * @param accountManagePaymentMethods Whether the apikey can modify account payment methods.
         * 
         * @return builder
         * 
         */
        public Builder accountManagePaymentMethods(Boolean accountManagePaymentMethods) {
            return accountManagePaymentMethods(Output.of(accountManagePaymentMethods));
        }

        /**
         * @param accountManagePlan No longer in use.
         * 
         * @return builder
         * 
         * @deprecated
         * obsolete, should no longer be used
         * 
         */
        @Deprecated /* obsolete, should no longer be used */
        public Builder accountManagePlan(@Nullable Output<Boolean> accountManagePlan) {
            $.accountManagePlan = accountManagePlan;
            return this;
        }

        /**
         * @param accountManagePlan No longer in use.
         * 
         * @return builder
         * 
         * @deprecated
         * obsolete, should no longer be used
         * 
         */
        @Deprecated /* obsolete, should no longer be used */
        public Builder accountManagePlan(Boolean accountManagePlan) {
            return accountManagePlan(Output.of(accountManagePlan));
        }

        /**
         * @param accountManageTeams Whether the apikey can modify other teams in the account.
         * 
         * @return builder
         * 
         */
        public Builder accountManageTeams(@Nullable Output<Boolean> accountManageTeams) {
            $.accountManageTeams = accountManageTeams;
            return this;
        }

        /**
         * @param accountManageTeams Whether the apikey can modify other teams in the account.
         * 
         * @return builder
         * 
         */
        public Builder accountManageTeams(Boolean accountManageTeams) {
            return accountManageTeams(Output.of(accountManageTeams));
        }

        /**
         * @param accountManageUsers Whether the apikey can modify account users.
         * 
         * @return builder
         * 
         */
        public Builder accountManageUsers(@Nullable Output<Boolean> accountManageUsers) {
            $.accountManageUsers = accountManageUsers;
            return this;
        }

        /**
         * @param accountManageUsers Whether the apikey can modify account users.
         * 
         * @return builder
         * 
         */
        public Builder accountManageUsers(Boolean accountManageUsers) {
            return accountManageUsers(Output.of(accountManageUsers));
        }

        /**
         * @param accountViewActivityLog Whether the apikey can view activity logs.
         * 
         * @return builder
         * 
         */
        public Builder accountViewActivityLog(@Nullable Output<Boolean> accountViewActivityLog) {
            $.accountViewActivityLog = accountViewActivityLog;
            return this;
        }

        /**
         * @param accountViewActivityLog Whether the apikey can view activity logs.
         * 
         * @return builder
         * 
         */
        public Builder accountViewActivityLog(Boolean accountViewActivityLog) {
            return accountViewActivityLog(Output.of(accountViewActivityLog));
        }

        /**
         * @param accountViewInvoices Whether the apikey can view invoices.
         * 
         * @return builder
         * 
         */
        public Builder accountViewInvoices(@Nullable Output<Boolean> accountViewInvoices) {
            $.accountViewInvoices = accountViewInvoices;
            return this;
        }

        /**
         * @param accountViewInvoices Whether the apikey can view invoices.
         * 
         * @return builder
         * 
         */
        public Builder accountViewInvoices(Boolean accountViewInvoices) {
            return accountViewInvoices(Output.of(accountViewInvoices));
        }

        /**
         * @param dataManageDatafeeds Whether the apikey can modify data feeds.
         * 
         * @return builder
         * 
         */
        public Builder dataManageDatafeeds(@Nullable Output<Boolean> dataManageDatafeeds) {
            $.dataManageDatafeeds = dataManageDatafeeds;
            return this;
        }

        /**
         * @param dataManageDatafeeds Whether the apikey can modify data feeds.
         * 
         * @return builder
         * 
         */
        public Builder dataManageDatafeeds(Boolean dataManageDatafeeds) {
            return dataManageDatafeeds(Output.of(dataManageDatafeeds));
        }

        /**
         * @param dataManageDatasources Whether the apikey can modify data sources.
         * 
         * @return builder
         * 
         */
        public Builder dataManageDatasources(@Nullable Output<Boolean> dataManageDatasources) {
            $.dataManageDatasources = dataManageDatasources;
            return this;
        }

        /**
         * @param dataManageDatasources Whether the apikey can modify data sources.
         * 
         * @return builder
         * 
         */
        public Builder dataManageDatasources(Boolean dataManageDatasources) {
            return dataManageDatasources(Output.of(dataManageDatasources));
        }

        /**
         * @param dataPushToDatafeeds Whether the apikey can publish to data feeds.
         * 
         * @return builder
         * 
         */
        public Builder dataPushToDatafeeds(@Nullable Output<Boolean> dataPushToDatafeeds) {
            $.dataPushToDatafeeds = dataPushToDatafeeds;
            return this;
        }

        /**
         * @param dataPushToDatafeeds Whether the apikey can publish to data feeds.
         * 
         * @return builder
         * 
         */
        public Builder dataPushToDatafeeds(Boolean dataPushToDatafeeds) {
            return dataPushToDatafeeds(Output.of(dataPushToDatafeeds));
        }

        /**
         * @param dhcpManageDhcp Whether the apikey can manage DHCP.
         * Only relevant for the DDI product.
         * 
         * @return builder
         * 
         */
        public Builder dhcpManageDhcp(@Nullable Output<Boolean> dhcpManageDhcp) {
            $.dhcpManageDhcp = dhcpManageDhcp;
            return this;
        }

        /**
         * @param dhcpManageDhcp Whether the apikey can manage DHCP.
         * Only relevant for the DDI product.
         * 
         * @return builder
         * 
         */
        public Builder dhcpManageDhcp(Boolean dhcpManageDhcp) {
            return dhcpManageDhcp(Output.of(dhcpManageDhcp));
        }

        /**
         * @param dhcpViewDhcp Whether the apikey can view DHCP.
         * Only relevant for the DDI product.
         * 
         * @return builder
         * 
         */
        public Builder dhcpViewDhcp(@Nullable Output<Boolean> dhcpViewDhcp) {
            $.dhcpViewDhcp = dhcpViewDhcp;
            return this;
        }

        /**
         * @param dhcpViewDhcp Whether the apikey can view DHCP.
         * Only relevant for the DDI product.
         * 
         * @return builder
         * 
         */
        public Builder dhcpViewDhcp(Boolean dhcpViewDhcp) {
            return dhcpViewDhcp(Output.of(dhcpViewDhcp));
        }

        /**
         * @param dnsManageZones Whether the apikey can modify the accounts zones.
         * 
         * @return builder
         * 
         */
        public Builder dnsManageZones(@Nullable Output<Boolean> dnsManageZones) {
            $.dnsManageZones = dnsManageZones;
            return this;
        }

        /**
         * @param dnsManageZones Whether the apikey can modify the accounts zones.
         * 
         * @return builder
         * 
         */
        public Builder dnsManageZones(Boolean dnsManageZones) {
            return dnsManageZones(Output.of(dnsManageZones));
        }

        /**
         * @param dnsRecordsAllows List of records that the apikey may access.
         * 
         * @return builder
         * 
         */
        public Builder dnsRecordsAllows(@Nullable Output<List<APIKeyDnsRecordsAllowArgs>> dnsRecordsAllows) {
            $.dnsRecordsAllows = dnsRecordsAllows;
            return this;
        }

        /**
         * @param dnsRecordsAllows List of records that the apikey may access.
         * 
         * @return builder
         * 
         */
        public Builder dnsRecordsAllows(List<APIKeyDnsRecordsAllowArgs> dnsRecordsAllows) {
            return dnsRecordsAllows(Output.of(dnsRecordsAllows));
        }

        /**
         * @param dnsRecordsAllows List of records that the apikey may access.
         * 
         * @return builder
         * 
         */
        public Builder dnsRecordsAllows(APIKeyDnsRecordsAllowArgs... dnsRecordsAllows) {
            return dnsRecordsAllows(List.of(dnsRecordsAllows));
        }

        /**
         * @param dnsRecordsDenies List of records that the apikey may not access.
         * 
         * @return builder
         * 
         */
        public Builder dnsRecordsDenies(@Nullable Output<List<APIKeyDnsRecordsDenyArgs>> dnsRecordsDenies) {
            $.dnsRecordsDenies = dnsRecordsDenies;
            return this;
        }

        /**
         * @param dnsRecordsDenies List of records that the apikey may not access.
         * 
         * @return builder
         * 
         */
        public Builder dnsRecordsDenies(List<APIKeyDnsRecordsDenyArgs> dnsRecordsDenies) {
            return dnsRecordsDenies(Output.of(dnsRecordsDenies));
        }

        /**
         * @param dnsRecordsDenies List of records that the apikey may not access.
         * 
         * @return builder
         * 
         */
        public Builder dnsRecordsDenies(APIKeyDnsRecordsDenyArgs... dnsRecordsDenies) {
            return dnsRecordsDenies(List.of(dnsRecordsDenies));
        }

        /**
         * @param dnsViewZones Whether the apikey can view the accounts zones.
         * 
         * @return builder
         * 
         */
        public Builder dnsViewZones(@Nullable Output<Boolean> dnsViewZones) {
            $.dnsViewZones = dnsViewZones;
            return this;
        }

        /**
         * @param dnsViewZones Whether the apikey can view the accounts zones.
         * 
         * @return builder
         * 
         */
        public Builder dnsViewZones(Boolean dnsViewZones) {
            return dnsViewZones(Output.of(dnsViewZones));
        }

        /**
         * @param dnsZonesAllowByDefault If true, enable the `dns_zones_allow` list, otherwise enable the `dns_zones_deny` list.
         * 
         * @return builder
         * 
         */
        public Builder dnsZonesAllowByDefault(@Nullable Output<Boolean> dnsZonesAllowByDefault) {
            $.dnsZonesAllowByDefault = dnsZonesAllowByDefault;
            return this;
        }

        /**
         * @param dnsZonesAllowByDefault If true, enable the `dns_zones_allow` list, otherwise enable the `dns_zones_deny` list.
         * 
         * @return builder
         * 
         */
        public Builder dnsZonesAllowByDefault(Boolean dnsZonesAllowByDefault) {
            return dnsZonesAllowByDefault(Output.of(dnsZonesAllowByDefault));
        }

        /**
         * @param dnsZonesAllows List of zones that the apikey may access.
         * 
         * @return builder
         * 
         */
        public Builder dnsZonesAllows(@Nullable Output<List<String>> dnsZonesAllows) {
            $.dnsZonesAllows = dnsZonesAllows;
            return this;
        }

        /**
         * @param dnsZonesAllows List of zones that the apikey may access.
         * 
         * @return builder
         * 
         */
        public Builder dnsZonesAllows(List<String> dnsZonesAllows) {
            return dnsZonesAllows(Output.of(dnsZonesAllows));
        }

        /**
         * @param dnsZonesAllows List of zones that the apikey may access.
         * 
         * @return builder
         * 
         */
        public Builder dnsZonesAllows(String... dnsZonesAllows) {
            return dnsZonesAllows(List.of(dnsZonesAllows));
        }

        /**
         * @param dnsZonesDenies List of zones that the apikey may not access.
         * 
         * @return builder
         * 
         */
        public Builder dnsZonesDenies(@Nullable Output<List<String>> dnsZonesDenies) {
            $.dnsZonesDenies = dnsZonesDenies;
            return this;
        }

        /**
         * @param dnsZonesDenies List of zones that the apikey may not access.
         * 
         * @return builder
         * 
         */
        public Builder dnsZonesDenies(List<String> dnsZonesDenies) {
            return dnsZonesDenies(Output.of(dnsZonesDenies));
        }

        /**
         * @param dnsZonesDenies List of zones that the apikey may not access.
         * 
         * @return builder
         * 
         */
        public Builder dnsZonesDenies(String... dnsZonesDenies) {
            return dnsZonesDenies(List.of(dnsZonesDenies));
        }

        /**
         * @param ipWhitelistStrict Set to true to restrict access to only those IP addresses and networks listed in the **ip_whitelist** field.
         * 
         * @return builder
         * 
         */
        public Builder ipWhitelistStrict(@Nullable Output<Boolean> ipWhitelistStrict) {
            $.ipWhitelistStrict = ipWhitelistStrict;
            return this;
        }

        /**
         * @param ipWhitelistStrict Set to true to restrict access to only those IP addresses and networks listed in the **ip_whitelist** field.
         * 
         * @return builder
         * 
         */
        public Builder ipWhitelistStrict(Boolean ipWhitelistStrict) {
            return ipWhitelistStrict(Output.of(ipWhitelistStrict));
        }

        /**
         * @param ipWhitelists Array of IP addresses/networks to which to grant the API key access.
         * 
         * @return builder
         * 
         */
        public Builder ipWhitelists(@Nullable Output<List<String>> ipWhitelists) {
            $.ipWhitelists = ipWhitelists;
            return this;
        }

        /**
         * @param ipWhitelists Array of IP addresses/networks to which to grant the API key access.
         * 
         * @return builder
         * 
         */
        public Builder ipWhitelists(List<String> ipWhitelists) {
            return ipWhitelists(Output.of(ipWhitelists));
        }

        /**
         * @param ipWhitelists Array of IP addresses/networks to which to grant the API key access.
         * 
         * @return builder
         * 
         */
        public Builder ipWhitelists(String... ipWhitelists) {
            return ipWhitelists(List.of(ipWhitelists));
        }

        /**
         * @param ipamManageIpam Whether the apikey can manage IPAM.
         * Only relevant for the DDI product.
         * 
         * @return builder
         * 
         */
        public Builder ipamManageIpam(@Nullable Output<Boolean> ipamManageIpam) {
            $.ipamManageIpam = ipamManageIpam;
            return this;
        }

        /**
         * @param ipamManageIpam Whether the apikey can manage IPAM.
         * Only relevant for the DDI product.
         * 
         * @return builder
         * 
         */
        public Builder ipamManageIpam(Boolean ipamManageIpam) {
            return ipamManageIpam(Output.of(ipamManageIpam));
        }

        /**
         * @param ipamViewIpam Whether the apikey can view IPAM.
         * Only relevant for the DDI product.
         * 
         * @return builder
         * 
         */
        public Builder ipamViewIpam(@Nullable Output<Boolean> ipamViewIpam) {
            $.ipamViewIpam = ipamViewIpam;
            return this;
        }

        /**
         * @param ipamViewIpam Whether the apikey can view IPAM.
         * Only relevant for the DDI product.
         * 
         * @return builder
         * 
         */
        public Builder ipamViewIpam(Boolean ipamViewIpam) {
            return ipamViewIpam(Output.of(ipamViewIpam));
        }

        /**
         * @param monitoringManageJobs Whether the apikey can modify monitoring jobs.
         * 
         * @return builder
         * 
         */
        public Builder monitoringManageJobs(@Nullable Output<Boolean> monitoringManageJobs) {
            $.monitoringManageJobs = monitoringManageJobs;
            return this;
        }

        /**
         * @param monitoringManageJobs Whether the apikey can modify monitoring jobs.
         * 
         * @return builder
         * 
         */
        public Builder monitoringManageJobs(Boolean monitoringManageJobs) {
            return monitoringManageJobs(Output.of(monitoringManageJobs));
        }

        /**
         * @param monitoringManageLists Whether the apikey can modify notification lists.
         * 
         * @return builder
         * 
         */
        public Builder monitoringManageLists(@Nullable Output<Boolean> monitoringManageLists) {
            $.monitoringManageLists = monitoringManageLists;
            return this;
        }

        /**
         * @param monitoringManageLists Whether the apikey can modify notification lists.
         * 
         * @return builder
         * 
         */
        public Builder monitoringManageLists(Boolean monitoringManageLists) {
            return monitoringManageLists(Output.of(monitoringManageLists));
        }

        /**
         * @param monitoringViewJobs Whether the apikey can view monitoring jobs.
         * 
         * @return builder
         * 
         */
        public Builder monitoringViewJobs(@Nullable Output<Boolean> monitoringViewJobs) {
            $.monitoringViewJobs = monitoringViewJobs;
            return this;
        }

        /**
         * @param monitoringViewJobs Whether the apikey can view monitoring jobs.
         * 
         * @return builder
         * 
         */
        public Builder monitoringViewJobs(Boolean monitoringViewJobs) {
            return monitoringViewJobs(Output.of(monitoringViewJobs));
        }

        /**
         * @param name The free form name of the apikey.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The free form name of the apikey.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param securityManageActiveDirectory Whether the apikey can manage global active directory.
         * Only relevant for the DDI product.
         * 
         * @return builder
         * 
         */
        public Builder securityManageActiveDirectory(@Nullable Output<Boolean> securityManageActiveDirectory) {
            $.securityManageActiveDirectory = securityManageActiveDirectory;
            return this;
        }

        /**
         * @param securityManageActiveDirectory Whether the apikey can manage global active directory.
         * Only relevant for the DDI product.
         * 
         * @return builder
         * 
         */
        public Builder securityManageActiveDirectory(Boolean securityManageActiveDirectory) {
            return securityManageActiveDirectory(Output.of(securityManageActiveDirectory));
        }

        /**
         * @param securityManageGlobal2fa Whether the apikey can manage global two factor authentication.
         * 
         * @return builder
         * 
         */
        public Builder securityManageGlobal2fa(@Nullable Output<Boolean> securityManageGlobal2fa) {
            $.securityManageGlobal2fa = securityManageGlobal2fa;
            return this;
        }

        /**
         * @param securityManageGlobal2fa Whether the apikey can manage global two factor authentication.
         * 
         * @return builder
         * 
         */
        public Builder securityManageGlobal2fa(Boolean securityManageGlobal2fa) {
            return securityManageGlobal2fa(Output.of(securityManageGlobal2fa));
        }

        /**
         * @param teams The teams that the apikey belongs to.
         * 
         * @return builder
         * 
         */
        public Builder teams(@Nullable Output<List<String>> teams) {
            $.teams = teams;
            return this;
        }

        /**
         * @param teams The teams that the apikey belongs to.
         * 
         * @return builder
         * 
         */
        public Builder teams(List<String> teams) {
            return teams(Output.of(teams));
        }

        /**
         * @param teams The teams that the apikey belongs to.
         * 
         * @return builder
         * 
         */
        public Builder teams(String... teams) {
            return teams(List.of(teams));
        }

        public APIKeyArgs build() {
            return $;
        }
    }

}
