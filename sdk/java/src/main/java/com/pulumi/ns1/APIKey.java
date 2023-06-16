// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ns1.APIKeyArgs;
import com.pulumi.ns1.Utilities;
import com.pulumi.ns1.inputs.APIKeyState;
import com.pulumi.ns1.outputs.APIKeyDnsRecordsAllow;
import com.pulumi.ns1.outputs.APIKeyDnsRecordsDeny;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a NS1 Api Key resource. This can be used to create, modify, and delete api keys.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ns1.Team;
 * import com.pulumi.ns1.APIKey;
 * import com.pulumi.ns1.APIKeyArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var exampleTeam = new Team(&#34;exampleTeam&#34;);
 * 
 *         var exampleAPIKey = new APIKey(&#34;exampleAPIKey&#34;, APIKeyArgs.builder()        
 *             .teams(exampleTeam.id())
 *             .ipWhitelists(            
 *                 &#34;1.1.1.1&#34;,
 *                 &#34;2.2.2.2&#34;)
 *             .dnsViewZones(false)
 *             .accountManageUsers(false)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Permissions
 * 
 * An API key will inherit permissions from the teams it is assigned to.
 * If a key is assigned to a team and also has individual permissions set on the key, the individual permissions
 * will be overridden by the inherited team permissions.
 * In a future release, setting permissions on a key that is part of a team will be explicitly disabled.
 * 
 * When a key is removed from all teams completely, it will inherit whatever permissions it had previously.
 * If a key is removed from all it&#39;s teams, it will probably be necessary to run `pulumi up` a second time
 * to update the keys permissions from it&#39;s old team permissions to new key-specific permissions.
 * 
 * See [the NS1 API docs](https://ns1.com/api#getget-all-account-users) for an overview of permission semantics or for [more details](https://help.ns1.com/hc/en-us/articles/360024409034-Managing-user-permissions) about the individual permission flags.
 * 
 * ## NS1 Documentation
 * 
 * [ApiKeys Api Doc](https://ns1.com/api#api-key)
 * 
 * ## Import
 * 
 * ```sh
 *  $ pulumi import ns1:index/aPIKey:APIKey `ns1_apikey`
 * ```
 * 
 *  So for the example above
 * 
 * ```sh
 *  $ pulumi import ns1:index/aPIKey:APIKey example &lt;ID&gt;`
 * ```
 * 
 */
@ResourceType(type="ns1:index/aPIKey:APIKey")
public class APIKey extends com.pulumi.resources.CustomResource {
    /**
     * Whether the apikey can modify account settings.
     * 
     */
    @Export(name="accountManageAccountSettings", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> accountManageAccountSettings;

    /**
     * @return Whether the apikey can modify account settings.
     * 
     */
    public Output<Optional<Boolean>> accountManageAccountSettings() {
        return Codegen.optional(this.accountManageAccountSettings);
    }
    /**
     * Whether the apikey can modify account apikeys.
     * 
     */
    @Export(name="accountManageApikeys", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> accountManageApikeys;

    /**
     * @return Whether the apikey can modify account apikeys.
     * 
     */
    public Output<Optional<Boolean>> accountManageApikeys() {
        return Codegen.optional(this.accountManageApikeys);
    }
    /**
     * Whether the apikey can manage ip whitelist.
     * 
     */
    @Export(name="accountManageIpWhitelist", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> accountManageIpWhitelist;

    /**
     * @return Whether the apikey can manage ip whitelist.
     * 
     */
    public Output<Optional<Boolean>> accountManageIpWhitelist() {
        return Codegen.optional(this.accountManageIpWhitelist);
    }
    /**
     * Whether the apikey can modify account payment methods.
     * 
     */
    @Export(name="accountManagePaymentMethods", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> accountManagePaymentMethods;

    /**
     * @return Whether the apikey can modify account payment methods.
     * 
     */
    public Output<Optional<Boolean>> accountManagePaymentMethods() {
        return Codegen.optional(this.accountManagePaymentMethods);
    }
    /**
     * No longer in use.
     * 
     * @deprecated
     * obsolete, should no longer be used
     * 
     */
    @Deprecated /* obsolete, should no longer be used */
    @Export(name="accountManagePlan", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> accountManagePlan;

    /**
     * @return No longer in use.
     * 
     */
    public Output<Optional<Boolean>> accountManagePlan() {
        return Codegen.optional(this.accountManagePlan);
    }
    /**
     * Whether the apikey can modify other teams in the account.
     * 
     */
    @Export(name="accountManageTeams", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> accountManageTeams;

    /**
     * @return Whether the apikey can modify other teams in the account.
     * 
     */
    public Output<Optional<Boolean>> accountManageTeams() {
        return Codegen.optional(this.accountManageTeams);
    }
    /**
     * Whether the apikey can modify account users.
     * 
     */
    @Export(name="accountManageUsers", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> accountManageUsers;

    /**
     * @return Whether the apikey can modify account users.
     * 
     */
    public Output<Optional<Boolean>> accountManageUsers() {
        return Codegen.optional(this.accountManageUsers);
    }
    /**
     * Whether the apikey can view activity logs.
     * 
     */
    @Export(name="accountViewActivityLog", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> accountViewActivityLog;

    /**
     * @return Whether the apikey can view activity logs.
     * 
     */
    public Output<Optional<Boolean>> accountViewActivityLog() {
        return Codegen.optional(this.accountViewActivityLog);
    }
    /**
     * Whether the apikey can view invoices.
     * 
     */
    @Export(name="accountViewInvoices", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> accountViewInvoices;

    /**
     * @return Whether the apikey can view invoices.
     * 
     */
    public Output<Optional<Boolean>> accountViewInvoices() {
        return Codegen.optional(this.accountViewInvoices);
    }
    /**
     * Whether the apikey can modify data feeds.
     * 
     */
    @Export(name="dataManageDatafeeds", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> dataManageDatafeeds;

    /**
     * @return Whether the apikey can modify data feeds.
     * 
     */
    public Output<Optional<Boolean>> dataManageDatafeeds() {
        return Codegen.optional(this.dataManageDatafeeds);
    }
    /**
     * Whether the apikey can modify data sources.
     * 
     */
    @Export(name="dataManageDatasources", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> dataManageDatasources;

    /**
     * @return Whether the apikey can modify data sources.
     * 
     */
    public Output<Optional<Boolean>> dataManageDatasources() {
        return Codegen.optional(this.dataManageDatasources);
    }
    /**
     * Whether the apikey can publish to data feeds.
     * 
     */
    @Export(name="dataPushToDatafeeds", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> dataPushToDatafeeds;

    /**
     * @return Whether the apikey can publish to data feeds.
     * 
     */
    public Output<Optional<Boolean>> dataPushToDatafeeds() {
        return Codegen.optional(this.dataPushToDatafeeds);
    }
    /**
     * Whether the apikey can manage DHCP.
     * Only relevant for the DDI product.
     * 
     */
    @Export(name="dhcpManageDhcp", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> dhcpManageDhcp;

    /**
     * @return Whether the apikey can manage DHCP.
     * Only relevant for the DDI product.
     * 
     */
    public Output<Optional<Boolean>> dhcpManageDhcp() {
        return Codegen.optional(this.dhcpManageDhcp);
    }
    /**
     * Whether the apikey can view DHCP.
     * Only relevant for the DDI product.
     * 
     */
    @Export(name="dhcpViewDhcp", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> dhcpViewDhcp;

    /**
     * @return Whether the apikey can view DHCP.
     * Only relevant for the DDI product.
     * 
     */
    public Output<Optional<Boolean>> dhcpViewDhcp() {
        return Codegen.optional(this.dhcpViewDhcp);
    }
    /**
     * Whether the apikey can modify the accounts zones.
     * 
     */
    @Export(name="dnsManageZones", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> dnsManageZones;

    /**
     * @return Whether the apikey can modify the accounts zones.
     * 
     */
    public Output<Optional<Boolean>> dnsManageZones() {
        return Codegen.optional(this.dnsManageZones);
    }
    /**
     * List of records that the apikey may access.
     * 
     */
    @Export(name="dnsRecordsAllows", type=List.class, parameters={APIKeyDnsRecordsAllow.class})
    private Output</* @Nullable */ List<APIKeyDnsRecordsAllow>> dnsRecordsAllows;

    /**
     * @return List of records that the apikey may access.
     * 
     */
    public Output<Optional<List<APIKeyDnsRecordsAllow>>> dnsRecordsAllows() {
        return Codegen.optional(this.dnsRecordsAllows);
    }
    /**
     * List of records that the apikey may not access.
     * 
     */
    @Export(name="dnsRecordsDenies", type=List.class, parameters={APIKeyDnsRecordsDeny.class})
    private Output</* @Nullable */ List<APIKeyDnsRecordsDeny>> dnsRecordsDenies;

    /**
     * @return List of records that the apikey may not access.
     * 
     */
    public Output<Optional<List<APIKeyDnsRecordsDeny>>> dnsRecordsDenies() {
        return Codegen.optional(this.dnsRecordsDenies);
    }
    /**
     * Whether the apikey can view the accounts zones.
     * 
     */
    @Export(name="dnsViewZones", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> dnsViewZones;

    /**
     * @return Whether the apikey can view the accounts zones.
     * 
     */
    public Output<Optional<Boolean>> dnsViewZones() {
        return Codegen.optional(this.dnsViewZones);
    }
    /**
     * If true, enable the `dns_zones_allow` list, otherwise enable the `dns_zones_deny` list.
     * 
     */
    @Export(name="dnsZonesAllowByDefault", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> dnsZonesAllowByDefault;

    /**
     * @return If true, enable the `dns_zones_allow` list, otherwise enable the `dns_zones_deny` list.
     * 
     */
    public Output<Optional<Boolean>> dnsZonesAllowByDefault() {
        return Codegen.optional(this.dnsZonesAllowByDefault);
    }
    /**
     * List of zones that the apikey may access.
     * 
     */
    @Export(name="dnsZonesAllows", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> dnsZonesAllows;

    /**
     * @return List of zones that the apikey may access.
     * 
     */
    public Output<Optional<List<String>>> dnsZonesAllows() {
        return Codegen.optional(this.dnsZonesAllows);
    }
    /**
     * List of zones that the apikey may not access.
     * 
     */
    @Export(name="dnsZonesDenies", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> dnsZonesDenies;

    /**
     * @return List of zones that the apikey may not access.
     * 
     */
    public Output<Optional<List<String>>> dnsZonesDenies() {
        return Codegen.optional(this.dnsZonesDenies);
    }
    /**
     * Set to true to restrict access to only those IP addresses and networks listed in the **ip_whitelist** field.
     * 
     */
    @Export(name="ipWhitelistStrict", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> ipWhitelistStrict;

    /**
     * @return Set to true to restrict access to only those IP addresses and networks listed in the **ip_whitelist** field.
     * 
     */
    public Output<Optional<Boolean>> ipWhitelistStrict() {
        return Codegen.optional(this.ipWhitelistStrict);
    }
    /**
     * Array of IP addresses/networks to which to grant the API key access.
     * 
     */
    @Export(name="ipWhitelists", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> ipWhitelists;

    /**
     * @return Array of IP addresses/networks to which to grant the API key access.
     * 
     */
    public Output<Optional<List<String>>> ipWhitelists() {
        return Codegen.optional(this.ipWhitelists);
    }
    /**
     * Whether the apikey can manage IPAM.
     * Only relevant for the DDI product.
     * 
     */
    @Export(name="ipamManageIpam", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> ipamManageIpam;

    /**
     * @return Whether the apikey can manage IPAM.
     * Only relevant for the DDI product.
     * 
     */
    public Output<Optional<Boolean>> ipamManageIpam() {
        return Codegen.optional(this.ipamManageIpam);
    }
    /**
     * Whether the apikey can view IPAM.
     * Only relevant for the DDI product.
     * 
     */
    @Export(name="ipamViewIpam", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> ipamViewIpam;

    /**
     * @return Whether the apikey can view IPAM.
     * Only relevant for the DDI product.
     * 
     */
    public Output<Optional<Boolean>> ipamViewIpam() {
        return Codegen.optional(this.ipamViewIpam);
    }
    /**
     * (Computed) The apikeys authentication token.
     * 
     */
    @Export(name="key", type=String.class, parameters={})
    private Output<String> key;

    /**
     * @return (Computed) The apikeys authentication token.
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * Whether the apikey can modify monitoring jobs.
     * 
     */
    @Export(name="monitoringManageJobs", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> monitoringManageJobs;

    /**
     * @return Whether the apikey can modify monitoring jobs.
     * 
     */
    public Output<Optional<Boolean>> monitoringManageJobs() {
        return Codegen.optional(this.monitoringManageJobs);
    }
    /**
     * Whether the apikey can modify notification lists.
     * 
     */
    @Export(name="monitoringManageLists", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> monitoringManageLists;

    /**
     * @return Whether the apikey can modify notification lists.
     * 
     */
    public Output<Optional<Boolean>> monitoringManageLists() {
        return Codegen.optional(this.monitoringManageLists);
    }
    /**
     * Whether the apikey can view monitoring jobs.
     * 
     */
    @Export(name="monitoringViewJobs", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> monitoringViewJobs;

    /**
     * @return Whether the apikey can view monitoring jobs.
     * 
     */
    public Output<Optional<Boolean>> monitoringViewJobs() {
        return Codegen.optional(this.monitoringViewJobs);
    }
    /**
     * The free form name of the apikey.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The free form name of the apikey.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Whether the apikey can manage global active directory.
     * Only relevant for the DDI product.
     * 
     */
    @Export(name="securityManageActiveDirectory", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> securityManageActiveDirectory;

    /**
     * @return Whether the apikey can manage global active directory.
     * Only relevant for the DDI product.
     * 
     */
    public Output<Optional<Boolean>> securityManageActiveDirectory() {
        return Codegen.optional(this.securityManageActiveDirectory);
    }
    /**
     * Whether the apikey can manage global two factor authentication.
     * 
     */
    @Export(name="securityManageGlobal2fa", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> securityManageGlobal2fa;

    /**
     * @return Whether the apikey can manage global two factor authentication.
     * 
     */
    public Output<Optional<Boolean>> securityManageGlobal2fa() {
        return Codegen.optional(this.securityManageGlobal2fa);
    }
    /**
     * The teams that the apikey belongs to.
     * 
     */
    @Export(name="teams", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> teams;

    /**
     * @return The teams that the apikey belongs to.
     * 
     */
    public Output<Optional<List<String>>> teams() {
        return Codegen.optional(this.teams);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public APIKey(String name) {
        this(name, APIKeyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public APIKey(String name, @Nullable APIKeyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public APIKey(String name, @Nullable APIKeyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ns1:index/aPIKey:APIKey", name, args == null ? APIKeyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private APIKey(String name, Output<String> id, @Nullable APIKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ns1:index/aPIKey:APIKey", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "key"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static APIKey get(String name, Output<String> id, @Nullable APIKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new APIKey(name, id, state, options);
    }
}
