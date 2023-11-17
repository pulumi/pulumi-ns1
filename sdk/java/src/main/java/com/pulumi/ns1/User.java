// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ns1;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ns1.UserArgs;
import com.pulumi.ns1.Utilities;
import com.pulumi.ns1.inputs.UserState;
import com.pulumi.ns1.outputs.UserDnsRecordsAllow;
import com.pulumi.ns1.outputs.UserDnsRecordsDeny;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a NS1 User resource. Creating a user sends an invitation email to the
 * user&#39;s email address. This can be used to create, modify, and delete users.
 * The credentials used must have the `manage_users` permission set.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ns1.Team;
 * import com.pulumi.ns1.TeamArgs;
 * import com.pulumi.ns1.User;
 * import com.pulumi.ns1.UserArgs;
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
 *         var exampleTeam = new Team(&#34;exampleTeam&#34;, TeamArgs.builder()        
 *             .ipWhitelists(            
 *                 &#34;1.1.1.1&#34;,
 *                 &#34;2.2.2.2&#34;)
 *             .dnsViewZones(false)
 *             .accountManageUsers(false)
 *             .build());
 * 
 *         var exampleUser = new User(&#34;exampleUser&#34;, UserArgs.builder()        
 *             .username(&#34;example_user&#34;)
 *             .email(&#34;user@example.com&#34;)
 *             .teams(exampleTeam.id())
 *             .notify(Map.of(&#34;billing&#34;, false))
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Permissions
 * 
 * A user will inherit permissions from the teams they are assigned to.
 * If a user is assigned to a team and also has individual permissions set on the user, the individual permissions
 * will be overridden by the inherited team permissions.
 * In a future release, setting permissions on a user that is part of a team will be explicitly disabled.
 * 
 * When a user is removed from all teams completely, they will inherit whatever permissions they had previously.
 * If a user is removed from all their teams, it will probably be necessary to run `pulumi up` a second time
 * to update the users permissions from their old team permissions to new user-specific permissions.
 * 
 * See [this NS1 Help Center article](https://help.ns1.com/hc/en-us/articles/360024409034-Managing-user-permissions) for an overview of user permission settings.
 * 
 * ## NS1 Documentation
 * 
 * [User Api Docs](https://ns1.com/api#user)
 * 
 * [Managing user permissions](https://help.ns1.com/hc/en-us/articles/360024409034-Managing-user-permissions)
 * 
 * ## Import
 * 
 * ```sh
 *  $ pulumi import ns1:index/user:User &lt;name&gt; &lt;username&gt;`
 * ```
 * 
 */
@ResourceType(type="ns1:index/user:User")
public class User extends com.pulumi.resources.CustomResource {
    /**
     * Whether the user can modify account settings.
     * 
     */
    @Export(name="accountManageAccountSettings", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> accountManageAccountSettings;

    /**
     * @return Whether the user can modify account settings.
     * 
     */
    public Output<Optional<Boolean>> accountManageAccountSettings() {
        return Codegen.optional(this.accountManageAccountSettings);
    }
    /**
     * Whether the user can modify account apikeys.
     * 
     */
    @Export(name="accountManageApikeys", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> accountManageApikeys;

    /**
     * @return Whether the user can modify account apikeys.
     * 
     */
    public Output<Optional<Boolean>> accountManageApikeys() {
        return Codegen.optional(this.accountManageApikeys);
    }
    /**
     * Whether the user can manage ip whitelist.
     * 
     */
    @Export(name="accountManageIpWhitelist", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> accountManageIpWhitelist;

    /**
     * @return Whether the user can manage ip whitelist.
     * 
     */
    public Output<Optional<Boolean>> accountManageIpWhitelist() {
        return Codegen.optional(this.accountManageIpWhitelist);
    }
    /**
     * Whether the user can modify account payment methods.
     * 
     */
    @Export(name="accountManagePaymentMethods", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> accountManagePaymentMethods;

    /**
     * @return Whether the user can modify account payment methods.
     * 
     */
    public Output<Optional<Boolean>> accountManagePaymentMethods() {
        return Codegen.optional(this.accountManagePaymentMethods);
    }
    /**
     * (Deprecated) No longer in use.
     * 
     * @deprecated
     * obsolete, should no longer be used
     * 
     */
    @Deprecated /* obsolete, should no longer be used */
    @Export(name="accountManagePlan", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> accountManagePlan;

    /**
     * @return (Deprecated) No longer in use.
     * 
     */
    public Output<Optional<Boolean>> accountManagePlan() {
        return Codegen.optional(this.accountManagePlan);
    }
    /**
     * Whether the user can modify other teams in the account.
     * 
     */
    @Export(name="accountManageTeams", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> accountManageTeams;

    /**
     * @return Whether the user can modify other teams in the account.
     * 
     */
    public Output<Optional<Boolean>> accountManageTeams() {
        return Codegen.optional(this.accountManageTeams);
    }
    /**
     * Whether the user can modify account users.
     * 
     */
    @Export(name="accountManageUsers", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> accountManageUsers;

    /**
     * @return Whether the user can modify account users.
     * 
     */
    public Output<Optional<Boolean>> accountManageUsers() {
        return Codegen.optional(this.accountManageUsers);
    }
    /**
     * Whether the user can view activity logs.
     * 
     */
    @Export(name="accountViewActivityLog", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> accountViewActivityLog;

    /**
     * @return Whether the user can view activity logs.
     * 
     */
    public Output<Optional<Boolean>> accountViewActivityLog() {
        return Codegen.optional(this.accountViewActivityLog);
    }
    /**
     * Whether the user can view invoices.
     * 
     */
    @Export(name="accountViewInvoices", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> accountViewInvoices;

    /**
     * @return Whether the user can view invoices.
     * 
     */
    public Output<Optional<Boolean>> accountViewInvoices() {
        return Codegen.optional(this.accountViewInvoices);
    }
    /**
     * Whether the user can modify data feeds.
     * 
     */
    @Export(name="dataManageDatafeeds", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> dataManageDatafeeds;

    /**
     * @return Whether the user can modify data feeds.
     * 
     */
    public Output<Optional<Boolean>> dataManageDatafeeds() {
        return Codegen.optional(this.dataManageDatafeeds);
    }
    /**
     * Whether the user can modify data sources.
     * 
     */
    @Export(name="dataManageDatasources", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> dataManageDatasources;

    /**
     * @return Whether the user can modify data sources.
     * 
     */
    public Output<Optional<Boolean>> dataManageDatasources() {
        return Codegen.optional(this.dataManageDatasources);
    }
    /**
     * Whether the user can publish to data feeds.
     * 
     */
    @Export(name="dataPushToDatafeeds", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> dataPushToDatafeeds;

    /**
     * @return Whether the user can publish to data feeds.
     * 
     */
    public Output<Optional<Boolean>> dataPushToDatafeeds() {
        return Codegen.optional(this.dataPushToDatafeeds);
    }
    /**
     * Whether the user can manage DHCP.
     * Only relevant for the DDI product.
     * 
     */
    @Export(name="dhcpManageDhcp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> dhcpManageDhcp;

    /**
     * @return Whether the user can manage DHCP.
     * Only relevant for the DDI product.
     * 
     */
    public Output<Optional<Boolean>> dhcpManageDhcp() {
        return Codegen.optional(this.dhcpManageDhcp);
    }
    /**
     * Whether the user can view DHCP.
     * Only relevant for the DDI product.
     * 
     */
    @Export(name="dhcpViewDhcp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> dhcpViewDhcp;

    /**
     * @return Whether the user can view DHCP.
     * Only relevant for the DDI product.
     * 
     */
    public Output<Optional<Boolean>> dhcpViewDhcp() {
        return Codegen.optional(this.dhcpViewDhcp);
    }
    /**
     * Whether the user can modify the accounts zones.
     * 
     */
    @Export(name="dnsManageZones", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> dnsManageZones;

    /**
     * @return Whether the user can modify the accounts zones.
     * 
     */
    public Output<Optional<Boolean>> dnsManageZones() {
        return Codegen.optional(this.dnsManageZones);
    }
    @Export(name="dnsRecordsAllows", refs={List.class,UserDnsRecordsAllow.class}, tree="[0,1]")
    private Output</* @Nullable */ List<UserDnsRecordsAllow>> dnsRecordsAllows;

    public Output<Optional<List<UserDnsRecordsAllow>>> dnsRecordsAllows() {
        return Codegen.optional(this.dnsRecordsAllows);
    }
    @Export(name="dnsRecordsDenies", refs={List.class,UserDnsRecordsDeny.class}, tree="[0,1]")
    private Output</* @Nullable */ List<UserDnsRecordsDeny>> dnsRecordsDenies;

    public Output<Optional<List<UserDnsRecordsDeny>>> dnsRecordsDenies() {
        return Codegen.optional(this.dnsRecordsDenies);
    }
    /**
     * Whether the user can view the accounts zones.
     * 
     */
    @Export(name="dnsViewZones", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> dnsViewZones;

    /**
     * @return Whether the user can view the accounts zones.
     * 
     */
    public Output<Optional<Boolean>> dnsViewZones() {
        return Codegen.optional(this.dnsViewZones);
    }
    /**
     * If true, enable the `dns_zones_allow` list, otherwise enable the `dns_zones_deny` list.
     * 
     */
    @Export(name="dnsZonesAllowByDefault", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> dnsZonesAllowByDefault;

    /**
     * @return If true, enable the `dns_zones_allow` list, otherwise enable the `dns_zones_deny` list.
     * 
     */
    public Output<Optional<Boolean>> dnsZonesAllowByDefault() {
        return Codegen.optional(this.dnsZonesAllowByDefault);
    }
    /**
     * List of zones that the user may access.
     * 
     */
    @Export(name="dnsZonesAllows", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> dnsZonesAllows;

    /**
     * @return List of zones that the user may access.
     * 
     */
    public Output<Optional<List<String>>> dnsZonesAllows() {
        return Codegen.optional(this.dnsZonesAllows);
    }
    /**
     * List of zones that the user may not access.
     * 
     */
    @Export(name="dnsZonesDenies", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> dnsZonesDenies;

    /**
     * @return List of zones that the user may not access.
     * 
     */
    public Output<Optional<List<String>>> dnsZonesDenies() {
        return Codegen.optional(this.dnsZonesDenies);
    }
    /**
     * The email address of the user.
     * 
     */
    @Export(name="email", refs={String.class}, tree="[0]")
    private Output<String> email;

    /**
     * @return The email address of the user.
     * 
     */
    public Output<String> email() {
        return this.email;
    }
    /**
     * Set to true to restrict access to only those IP addresses and networks listed in the **ip_whitelist** field.
     * 
     */
    @Export(name="ipWhitelistStrict", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> ipWhitelistStrict;

    /**
     * @return Set to true to restrict access to only those IP addresses and networks listed in the **ip_whitelist** field.
     * 
     */
    public Output<Optional<Boolean>> ipWhitelistStrict() {
        return Codegen.optional(this.ipWhitelistStrict);
    }
    /**
     * Array of IP addresses/networks to which to grant the user access.
     * 
     */
    @Export(name="ipWhitelists", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> ipWhitelists;

    /**
     * @return Array of IP addresses/networks to which to grant the user access.
     * 
     */
    public Output<Optional<List<String>>> ipWhitelists() {
        return Codegen.optional(this.ipWhitelists);
    }
    /**
     * Whether the user can manage IPAM.
     * Only relevant for the DDI product.
     * 
     */
    @Export(name="ipamManageIpam", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> ipamManageIpam;

    /**
     * @return Whether the user can manage IPAM.
     * Only relevant for the DDI product.
     * 
     */
    public Output<Optional<Boolean>> ipamManageIpam() {
        return Codegen.optional(this.ipamManageIpam);
    }
    @Export(name="ipamViewIpam", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> ipamViewIpam;

    public Output<Optional<Boolean>> ipamViewIpam() {
        return Codegen.optional(this.ipamViewIpam);
    }
    /**
     * Whether the user can modify monitoring jobs.
     * 
     */
    @Export(name="monitoringManageJobs", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> monitoringManageJobs;

    /**
     * @return Whether the user can modify monitoring jobs.
     * 
     */
    public Output<Optional<Boolean>> monitoringManageJobs() {
        return Codegen.optional(this.monitoringManageJobs);
    }
    /**
     * Whether the user can modify notification lists.
     * 
     */
    @Export(name="monitoringManageLists", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> monitoringManageLists;

    /**
     * @return Whether the user can modify notification lists.
     * 
     */
    public Output<Optional<Boolean>> monitoringManageLists() {
        return Codegen.optional(this.monitoringManageLists);
    }
    /**
     * Whether the user can view monitoring jobs.
     * 
     */
    @Export(name="monitoringViewJobs", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> monitoringViewJobs;

    /**
     * @return Whether the user can view monitoring jobs.
     * 
     */
    public Output<Optional<Boolean>> monitoringViewJobs() {
        return Codegen.optional(this.monitoringViewJobs);
    }
    /**
     * The free form name of the user.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The free form name of the user.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Whether or not to notify the user of specified events. Only `billing` is available currently.
     * 
     */
    @Export(name="notify", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> notify;

    /**
     * @return Whether or not to notify the user of specified events. Only `billing` is available currently.
     * 
     */
    public Output<Optional<Map<String,Object>>> notify_() {
        return Codegen.optional(this.notify);
    }
    /**
     * Whether the user can manage global active directory.
     * Only relevant for the DDI product.
     * 
     */
    @Export(name="securityManageActiveDirectory", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> securityManageActiveDirectory;

    /**
     * @return Whether the user can manage global active directory.
     * Only relevant for the DDI product.
     * 
     */
    public Output<Optional<Boolean>> securityManageActiveDirectory() {
        return Codegen.optional(this.securityManageActiveDirectory);
    }
    /**
     * Whether the user can manage global two factor authentication.
     * 
     */
    @Export(name="securityManageGlobal2fa", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> securityManageGlobal2fa;

    /**
     * @return Whether the user can manage global two factor authentication.
     * 
     */
    public Output<Optional<Boolean>> securityManageGlobal2fa() {
        return Codegen.optional(this.securityManageGlobal2fa);
    }
    /**
     * The teams that the user belongs to.
     * 
     */
    @Export(name="teams", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> teams;

    /**
     * @return The teams that the user belongs to.
     * 
     */
    public Output<Optional<List<String>>> teams() {
        return Codegen.optional(this.teams);
    }
    /**
     * The users login name.
     * 
     */
    @Export(name="username", refs={String.class}, tree="[0]")
    private Output<String> username;

    /**
     * @return The users login name.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public User(String name) {
        this(name, UserArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public User(String name, UserArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public User(String name, UserArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ns1:index/user:User", name, args == null ? UserArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private User(String name, Output<String> id, @Nullable UserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ns1:index/user:User", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
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
    public static User get(String name, Output<String> id, @Nullable UserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new User(name, id, state, options);
    }
}
