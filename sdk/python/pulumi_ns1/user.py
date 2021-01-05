# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['User']


class User(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_manage_account_settings: Optional[pulumi.Input[bool]] = None,
                 account_manage_apikeys: Optional[pulumi.Input[bool]] = None,
                 account_manage_payment_methods: Optional[pulumi.Input[bool]] = None,
                 account_manage_plan: Optional[pulumi.Input[bool]] = None,
                 account_manage_teams: Optional[pulumi.Input[bool]] = None,
                 account_manage_users: Optional[pulumi.Input[bool]] = None,
                 account_view_activity_log: Optional[pulumi.Input[bool]] = None,
                 account_view_invoices: Optional[pulumi.Input[bool]] = None,
                 data_manage_datafeeds: Optional[pulumi.Input[bool]] = None,
                 data_manage_datasources: Optional[pulumi.Input[bool]] = None,
                 data_push_to_datafeeds: Optional[pulumi.Input[bool]] = None,
                 dhcp_manage_dhcp: Optional[pulumi.Input[bool]] = None,
                 dhcp_view_dhcp: Optional[pulumi.Input[bool]] = None,
                 dns_manage_zones: Optional[pulumi.Input[bool]] = None,
                 dns_view_zones: Optional[pulumi.Input[bool]] = None,
                 dns_zones_allow_by_default: Optional[pulumi.Input[bool]] = None,
                 dns_zones_allows: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 dns_zones_denies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 ip_whitelist_strict: Optional[pulumi.Input[bool]] = None,
                 ip_whitelists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ipam_manage_ipam: Optional[pulumi.Input[bool]] = None,
                 ipam_view_ipam: Optional[pulumi.Input[bool]] = None,
                 monitoring_manage_jobs: Optional[pulumi.Input[bool]] = None,
                 monitoring_manage_lists: Optional[pulumi.Input[bool]] = None,
                 monitoring_view_jobs: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notify: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 security_manage_active_directory: Optional[pulumi.Input[bool]] = None,
                 security_manage_global2fa: Optional[pulumi.Input[bool]] = None,
                 teams: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a User resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] account_manage_account_settings: Whether the user can modify account settings.
        :param pulumi.Input[bool] account_manage_apikeys: Whether the user can modify account apikeys.
        :param pulumi.Input[bool] account_manage_payment_methods: Whether the user can modify account payment methods.
        :param pulumi.Input[bool] account_manage_plan: **Deprecated** Whether the user can modify the account plan.
        :param pulumi.Input[bool] account_manage_teams: Whether the user can modify other teams in the account.
        :param pulumi.Input[bool] account_manage_users: Whether the user can modify account users.
        :param pulumi.Input[bool] account_view_activity_log: Whether the user can view activity logs.
        :param pulumi.Input[bool] account_view_invoices: Whether the user can view invoices.
        :param pulumi.Input[bool] data_manage_datafeeds: Whether the user can modify data feeds.
        :param pulumi.Input[bool] data_manage_datasources: Whether the user can modify data sources.
        :param pulumi.Input[bool] data_push_to_datafeeds: Whether the user can publish to data feeds.
        :param pulumi.Input[bool] dhcp_manage_dhcp: Whether the user can manage DHCP.
               Only relevant for the DDI product.
        :param pulumi.Input[bool] dhcp_view_dhcp: Whether the user can view DHCP.
               Only relevant for the DDI product.
        :param pulumi.Input[bool] dns_manage_zones: Whether the user can modify the accounts zones.
        :param pulumi.Input[bool] dns_view_zones: Whether the user can view the accounts zones.
        :param pulumi.Input[bool] dns_zones_allow_by_default: If true, enable the `dns_zones_allow` list, otherwise enable the `dns_zones_deny` list.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_zones_allows: List of zones that the user may access.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_zones_denies: List of zones that the user may not access.
        :param pulumi.Input[str] email: The email address of the user.
        :param pulumi.Input[bool] ip_whitelist_strict: Sets exclusivity on this IP whitelist.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_whitelists: The IP addresses to whitelist for this key.
        :param pulumi.Input[bool] ipam_manage_ipam: Whether the user can manage IPAM.
               Only relevant for the DDI product.
        :param pulumi.Input[bool] monitoring_manage_jobs: Whether the user can modify monitoring jobs.
        :param pulumi.Input[bool] monitoring_manage_lists: Whether the user can modify notification lists.
        :param pulumi.Input[bool] monitoring_view_jobs: Whether the user can view monitoring jobs.
        :param pulumi.Input[str] name: The free form name of the user.
        :param pulumi.Input[Mapping[str, Any]] notify: Whether or not to notify the user of specified events. Only `billing` is available currently.
        :param pulumi.Input[bool] security_manage_active_directory: Whether the user can manage global active directory.
               Only relevant for the DDI product.
        :param pulumi.Input[bool] security_manage_global2fa: Whether the user can manage global two factor authentication.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] teams: The teams that the user belongs to.
        :param pulumi.Input[str] username: The users login name.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['account_manage_account_settings'] = account_manage_account_settings
            __props__['account_manage_apikeys'] = account_manage_apikeys
            __props__['account_manage_payment_methods'] = account_manage_payment_methods
            if account_manage_plan is not None:
                warnings.warn("""obsolete, should no longer be used""", DeprecationWarning)
                pulumi.log.warn("account_manage_plan is deprecated: obsolete, should no longer be used")
            __props__['account_manage_plan'] = account_manage_plan
            __props__['account_manage_teams'] = account_manage_teams
            __props__['account_manage_users'] = account_manage_users
            __props__['account_view_activity_log'] = account_view_activity_log
            __props__['account_view_invoices'] = account_view_invoices
            __props__['data_manage_datafeeds'] = data_manage_datafeeds
            __props__['data_manage_datasources'] = data_manage_datasources
            __props__['data_push_to_datafeeds'] = data_push_to_datafeeds
            __props__['dhcp_manage_dhcp'] = dhcp_manage_dhcp
            __props__['dhcp_view_dhcp'] = dhcp_view_dhcp
            __props__['dns_manage_zones'] = dns_manage_zones
            __props__['dns_view_zones'] = dns_view_zones
            __props__['dns_zones_allow_by_default'] = dns_zones_allow_by_default
            __props__['dns_zones_allows'] = dns_zones_allows
            __props__['dns_zones_denies'] = dns_zones_denies
            if email is None:
                raise TypeError("Missing required property 'email'")
            __props__['email'] = email
            __props__['ip_whitelist_strict'] = ip_whitelist_strict
            __props__['ip_whitelists'] = ip_whitelists
            __props__['ipam_manage_ipam'] = ipam_manage_ipam
            __props__['ipam_view_ipam'] = ipam_view_ipam
            __props__['monitoring_manage_jobs'] = monitoring_manage_jobs
            __props__['monitoring_manage_lists'] = monitoring_manage_lists
            __props__['monitoring_view_jobs'] = monitoring_view_jobs
            __props__['name'] = name
            __props__['notify'] = notify
            __props__['security_manage_active_directory'] = security_manage_active_directory
            __props__['security_manage_global2fa'] = security_manage_global2fa
            __props__['teams'] = teams
            if username is None:
                raise TypeError("Missing required property 'username'")
            __props__['username'] = username
        super(User, __self__).__init__(
            'ns1:index/user:User',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_manage_account_settings: Optional[pulumi.Input[bool]] = None,
            account_manage_apikeys: Optional[pulumi.Input[bool]] = None,
            account_manage_payment_methods: Optional[pulumi.Input[bool]] = None,
            account_manage_plan: Optional[pulumi.Input[bool]] = None,
            account_manage_teams: Optional[pulumi.Input[bool]] = None,
            account_manage_users: Optional[pulumi.Input[bool]] = None,
            account_view_activity_log: Optional[pulumi.Input[bool]] = None,
            account_view_invoices: Optional[pulumi.Input[bool]] = None,
            data_manage_datafeeds: Optional[pulumi.Input[bool]] = None,
            data_manage_datasources: Optional[pulumi.Input[bool]] = None,
            data_push_to_datafeeds: Optional[pulumi.Input[bool]] = None,
            dhcp_manage_dhcp: Optional[pulumi.Input[bool]] = None,
            dhcp_view_dhcp: Optional[pulumi.Input[bool]] = None,
            dns_manage_zones: Optional[pulumi.Input[bool]] = None,
            dns_view_zones: Optional[pulumi.Input[bool]] = None,
            dns_zones_allow_by_default: Optional[pulumi.Input[bool]] = None,
            dns_zones_allows: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            dns_zones_denies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            email: Optional[pulumi.Input[str]] = None,
            ip_whitelist_strict: Optional[pulumi.Input[bool]] = None,
            ip_whitelists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            ipam_manage_ipam: Optional[pulumi.Input[bool]] = None,
            ipam_view_ipam: Optional[pulumi.Input[bool]] = None,
            monitoring_manage_jobs: Optional[pulumi.Input[bool]] = None,
            monitoring_manage_lists: Optional[pulumi.Input[bool]] = None,
            monitoring_view_jobs: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            notify: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            security_manage_active_directory: Optional[pulumi.Input[bool]] = None,
            security_manage_global2fa: Optional[pulumi.Input[bool]] = None,
            teams: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'User':
        """
        Get an existing User resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] account_manage_account_settings: Whether the user can modify account settings.
        :param pulumi.Input[bool] account_manage_apikeys: Whether the user can modify account apikeys.
        :param pulumi.Input[bool] account_manage_payment_methods: Whether the user can modify account payment methods.
        :param pulumi.Input[bool] account_manage_plan: **Deprecated** Whether the user can modify the account plan.
        :param pulumi.Input[bool] account_manage_teams: Whether the user can modify other teams in the account.
        :param pulumi.Input[bool] account_manage_users: Whether the user can modify account users.
        :param pulumi.Input[bool] account_view_activity_log: Whether the user can view activity logs.
        :param pulumi.Input[bool] account_view_invoices: Whether the user can view invoices.
        :param pulumi.Input[bool] data_manage_datafeeds: Whether the user can modify data feeds.
        :param pulumi.Input[bool] data_manage_datasources: Whether the user can modify data sources.
        :param pulumi.Input[bool] data_push_to_datafeeds: Whether the user can publish to data feeds.
        :param pulumi.Input[bool] dhcp_manage_dhcp: Whether the user can manage DHCP.
               Only relevant for the DDI product.
        :param pulumi.Input[bool] dhcp_view_dhcp: Whether the user can view DHCP.
               Only relevant for the DDI product.
        :param pulumi.Input[bool] dns_manage_zones: Whether the user can modify the accounts zones.
        :param pulumi.Input[bool] dns_view_zones: Whether the user can view the accounts zones.
        :param pulumi.Input[bool] dns_zones_allow_by_default: If true, enable the `dns_zones_allow` list, otherwise enable the `dns_zones_deny` list.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_zones_allows: List of zones that the user may access.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] dns_zones_denies: List of zones that the user may not access.
        :param pulumi.Input[str] email: The email address of the user.
        :param pulumi.Input[bool] ip_whitelist_strict: Sets exclusivity on this IP whitelist.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_whitelists: The IP addresses to whitelist for this key.
        :param pulumi.Input[bool] ipam_manage_ipam: Whether the user can manage IPAM.
               Only relevant for the DDI product.
        :param pulumi.Input[bool] monitoring_manage_jobs: Whether the user can modify monitoring jobs.
        :param pulumi.Input[bool] monitoring_manage_lists: Whether the user can modify notification lists.
        :param pulumi.Input[bool] monitoring_view_jobs: Whether the user can view monitoring jobs.
        :param pulumi.Input[str] name: The free form name of the user.
        :param pulumi.Input[Mapping[str, Any]] notify: Whether or not to notify the user of specified events. Only `billing` is available currently.
        :param pulumi.Input[bool] security_manage_active_directory: Whether the user can manage global active directory.
               Only relevant for the DDI product.
        :param pulumi.Input[bool] security_manage_global2fa: Whether the user can manage global two factor authentication.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] teams: The teams that the user belongs to.
        :param pulumi.Input[str] username: The users login name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["account_manage_account_settings"] = account_manage_account_settings
        __props__["account_manage_apikeys"] = account_manage_apikeys
        __props__["account_manage_payment_methods"] = account_manage_payment_methods
        __props__["account_manage_plan"] = account_manage_plan
        __props__["account_manage_teams"] = account_manage_teams
        __props__["account_manage_users"] = account_manage_users
        __props__["account_view_activity_log"] = account_view_activity_log
        __props__["account_view_invoices"] = account_view_invoices
        __props__["data_manage_datafeeds"] = data_manage_datafeeds
        __props__["data_manage_datasources"] = data_manage_datasources
        __props__["data_push_to_datafeeds"] = data_push_to_datafeeds
        __props__["dhcp_manage_dhcp"] = dhcp_manage_dhcp
        __props__["dhcp_view_dhcp"] = dhcp_view_dhcp
        __props__["dns_manage_zones"] = dns_manage_zones
        __props__["dns_view_zones"] = dns_view_zones
        __props__["dns_zones_allow_by_default"] = dns_zones_allow_by_default
        __props__["dns_zones_allows"] = dns_zones_allows
        __props__["dns_zones_denies"] = dns_zones_denies
        __props__["email"] = email
        __props__["ip_whitelist_strict"] = ip_whitelist_strict
        __props__["ip_whitelists"] = ip_whitelists
        __props__["ipam_manage_ipam"] = ipam_manage_ipam
        __props__["ipam_view_ipam"] = ipam_view_ipam
        __props__["monitoring_manage_jobs"] = monitoring_manage_jobs
        __props__["monitoring_manage_lists"] = monitoring_manage_lists
        __props__["monitoring_view_jobs"] = monitoring_view_jobs
        __props__["name"] = name
        __props__["notify"] = notify
        __props__["security_manage_active_directory"] = security_manage_active_directory
        __props__["security_manage_global2fa"] = security_manage_global2fa
        __props__["teams"] = teams
        __props__["username"] = username
        return User(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountManageAccountSettings")
    def account_manage_account_settings(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the user can modify account settings.
        """
        return pulumi.get(self, "account_manage_account_settings")

    @property
    @pulumi.getter(name="accountManageApikeys")
    def account_manage_apikeys(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the user can modify account apikeys.
        """
        return pulumi.get(self, "account_manage_apikeys")

    @property
    @pulumi.getter(name="accountManagePaymentMethods")
    def account_manage_payment_methods(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the user can modify account payment methods.
        """
        return pulumi.get(self, "account_manage_payment_methods")

    @property
    @pulumi.getter(name="accountManagePlan")
    def account_manage_plan(self) -> pulumi.Output[Optional[bool]]:
        """
        **Deprecated** Whether the user can modify the account plan.
        """
        return pulumi.get(self, "account_manage_plan")

    @property
    @pulumi.getter(name="accountManageTeams")
    def account_manage_teams(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the user can modify other teams in the account.
        """
        return pulumi.get(self, "account_manage_teams")

    @property
    @pulumi.getter(name="accountManageUsers")
    def account_manage_users(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the user can modify account users.
        """
        return pulumi.get(self, "account_manage_users")

    @property
    @pulumi.getter(name="accountViewActivityLog")
    def account_view_activity_log(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the user can view activity logs.
        """
        return pulumi.get(self, "account_view_activity_log")

    @property
    @pulumi.getter(name="accountViewInvoices")
    def account_view_invoices(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the user can view invoices.
        """
        return pulumi.get(self, "account_view_invoices")

    @property
    @pulumi.getter(name="dataManageDatafeeds")
    def data_manage_datafeeds(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the user can modify data feeds.
        """
        return pulumi.get(self, "data_manage_datafeeds")

    @property
    @pulumi.getter(name="dataManageDatasources")
    def data_manage_datasources(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the user can modify data sources.
        """
        return pulumi.get(self, "data_manage_datasources")

    @property
    @pulumi.getter(name="dataPushToDatafeeds")
    def data_push_to_datafeeds(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the user can publish to data feeds.
        """
        return pulumi.get(self, "data_push_to_datafeeds")

    @property
    @pulumi.getter(name="dhcpManageDhcp")
    def dhcp_manage_dhcp(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the user can manage DHCP.
        Only relevant for the DDI product.
        """
        return pulumi.get(self, "dhcp_manage_dhcp")

    @property
    @pulumi.getter(name="dhcpViewDhcp")
    def dhcp_view_dhcp(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the user can view DHCP.
        Only relevant for the DDI product.
        """
        return pulumi.get(self, "dhcp_view_dhcp")

    @property
    @pulumi.getter(name="dnsManageZones")
    def dns_manage_zones(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the user can modify the accounts zones.
        """
        return pulumi.get(self, "dns_manage_zones")

    @property
    @pulumi.getter(name="dnsViewZones")
    def dns_view_zones(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the user can view the accounts zones.
        """
        return pulumi.get(self, "dns_view_zones")

    @property
    @pulumi.getter(name="dnsZonesAllowByDefault")
    def dns_zones_allow_by_default(self) -> pulumi.Output[Optional[bool]]:
        """
        If true, enable the `dns_zones_allow` list, otherwise enable the `dns_zones_deny` list.
        """
        return pulumi.get(self, "dns_zones_allow_by_default")

    @property
    @pulumi.getter(name="dnsZonesAllows")
    def dns_zones_allows(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of zones that the user may access.
        """
        return pulumi.get(self, "dns_zones_allows")

    @property
    @pulumi.getter(name="dnsZonesDenies")
    def dns_zones_denies(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of zones that the user may not access.
        """
        return pulumi.get(self, "dns_zones_denies")

    @property
    @pulumi.getter
    def email(self) -> pulumi.Output[str]:
        """
        The email address of the user.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="ipWhitelistStrict")
    def ip_whitelist_strict(self) -> pulumi.Output[Optional[bool]]:
        """
        Sets exclusivity on this IP whitelist.
        """
        return pulumi.get(self, "ip_whitelist_strict")

    @property
    @pulumi.getter(name="ipWhitelists")
    def ip_whitelists(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The IP addresses to whitelist for this key.
        """
        return pulumi.get(self, "ip_whitelists")

    @property
    @pulumi.getter(name="ipamManageIpam")
    def ipam_manage_ipam(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the user can manage IPAM.
        Only relevant for the DDI product.
        """
        return pulumi.get(self, "ipam_manage_ipam")

    @property
    @pulumi.getter(name="ipamViewIpam")
    def ipam_view_ipam(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "ipam_view_ipam")

    @property
    @pulumi.getter(name="monitoringManageJobs")
    def monitoring_manage_jobs(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the user can modify monitoring jobs.
        """
        return pulumi.get(self, "monitoring_manage_jobs")

    @property
    @pulumi.getter(name="monitoringManageLists")
    def monitoring_manage_lists(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the user can modify notification lists.
        """
        return pulumi.get(self, "monitoring_manage_lists")

    @property
    @pulumi.getter(name="monitoringViewJobs")
    def monitoring_view_jobs(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the user can view monitoring jobs.
        """
        return pulumi.get(self, "monitoring_view_jobs")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The free form name of the user.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def notify(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Whether or not to notify the user of specified events. Only `billing` is available currently.
        """
        return pulumi.get(self, "notify")

    @property
    @pulumi.getter(name="securityManageActiveDirectory")
    def security_manage_active_directory(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the user can manage global active directory.
        Only relevant for the DDI product.
        """
        return pulumi.get(self, "security_manage_active_directory")

    @property
    @pulumi.getter(name="securityManageGlobal2fa")
    def security_manage_global2fa(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the user can manage global two factor authentication.
        """
        return pulumi.get(self, "security_manage_global2fa")

    @property
    @pulumi.getter
    def teams(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The teams that the user belongs to.
        """
        return pulumi.get(self, "teams")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        """
        The users login name.
        """
        return pulumi.get(self, "username")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

