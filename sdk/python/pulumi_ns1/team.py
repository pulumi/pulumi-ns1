# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class Team(pulumi.CustomResource):
    account_manage_account_settings: pulumi.Output[bool]
    """
    Whether the team can modify account settings.
    """
    account_manage_apikeys: pulumi.Output[bool]
    """
    Whether the team can modify account apikeys.
    """
    account_manage_payment_methods: pulumi.Output[bool]
    """
    Whether the team can modify account payment methods.
    """
    account_manage_plan: pulumi.Output[bool]
    """
    Whether the team can modify the account plan.
    """
    account_manage_teams: pulumi.Output[bool]
    """
    Whether the team can modify other teams in the account.
    """
    account_manage_users: pulumi.Output[bool]
    """
    Whether the team can modify account users.
    """
    account_view_activity_log: pulumi.Output[bool]
    """
    Whether the team can view activity logs.
    """
    account_view_invoices: pulumi.Output[bool]
    """
    Whether the team can view invoices.
    """
    data_manage_datafeeds: pulumi.Output[bool]
    """
    Whether the team can modify data feeds.
    """
    data_manage_datasources: pulumi.Output[bool]
    """
    Whether the team can modify data sources.
    """
    data_push_to_datafeeds: pulumi.Output[bool]
    """
    Whether the team can publish to data feeds.
    """
    dhcp_manage_dhcp: pulumi.Output[bool]
    """
    Whether the team can manage DHCP.
    Only relevant for the DDI product.
    """
    dhcp_view_dhcp: pulumi.Output[bool]
    """
    Whether the team can view DHCP.
    Only relevant for the DDI product.
    """
    dns_manage_zones: pulumi.Output[bool]
    """
    Whether the team can modify the accounts zones.
    """
    dns_view_zones: pulumi.Output[bool]
    """
    Whether the team can view the accounts zones.
    """
    dns_zones_allow_by_default: pulumi.Output[bool]
    """
    If true, enable the `dns_zones_allow` list, otherwise enable the `dns_zones_deny` list.
    """
    dns_zones_allows: pulumi.Output[list]
    """
    List of zones that the team may access.
    """
    dns_zones_denies: pulumi.Output[list]
    """
    List of zones that the team may not access.
    """
    ip_whitelists: pulumi.Output[list]
    """
    The IP addresses to whitelist for this key.

      * `name` (`str`) - The free form name of the team.
      * `values` (`list`)
    """
    ipam_manage_ipam: pulumi.Output[bool]
    """
    Whether the team can manage IPAM.
    Only relevant for the DDI product.
    """
    ipam_view_ipam: pulumi.Output[bool]
    """
    Whether the team can view IPAM.
    Only relevant for the DDI product.
    """
    monitoring_manage_jobs: pulumi.Output[bool]
    """
    Whether the team can modify monitoring jobs.
    """
    monitoring_manage_lists: pulumi.Output[bool]
    """
    Whether the team can modify notification lists.
    """
    monitoring_view_jobs: pulumi.Output[bool]
    """
    Whether the team can view monitoring jobs.
    """
    name: pulumi.Output[str]
    """
    The free form name of the team.
    """
    security_manage_active_directory: pulumi.Output[bool]
    """
    Whether the team can manage global active directory.
    Only relevant for the DDI product.
    """
    security_manage_global2fa: pulumi.Output[bool]
    """
    Whether the team can manage global two factor authentication.
    """
    def __init__(__self__, resource_name, opts=None, account_manage_account_settings=None, account_manage_apikeys=None, account_manage_payment_methods=None, account_manage_plan=None, account_manage_teams=None, account_manage_users=None, account_view_activity_log=None, account_view_invoices=None, data_manage_datafeeds=None, data_manage_datasources=None, data_push_to_datafeeds=None, dhcp_manage_dhcp=None, dhcp_view_dhcp=None, dns_manage_zones=None, dns_view_zones=None, dns_zones_allow_by_default=None, dns_zones_allows=None, dns_zones_denies=None, ip_whitelists=None, ipam_manage_ipam=None, ipam_view_ipam=None, monitoring_manage_jobs=None, monitoring_manage_lists=None, monitoring_view_jobs=None, name=None, security_manage_active_directory=None, security_manage_global2fa=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a NS1 Team resource. This can be used to create, modify, and delete
        teams. The credentials used must have the `manage_teams` permission set.


        ## NS1 Documentation

        [Team Api Docs](https://ns1.com/api#team)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] account_manage_account_settings: Whether the team can modify account settings.
        :param pulumi.Input[bool] account_manage_apikeys: Whether the team can modify account apikeys.
        :param pulumi.Input[bool] account_manage_payment_methods: Whether the team can modify account payment methods.
        :param pulumi.Input[bool] account_manage_plan: Whether the team can modify the account plan.
        :param pulumi.Input[bool] account_manage_teams: Whether the team can modify other teams in the account.
        :param pulumi.Input[bool] account_manage_users: Whether the team can modify account users.
        :param pulumi.Input[bool] account_view_activity_log: Whether the team can view activity logs.
        :param pulumi.Input[bool] account_view_invoices: Whether the team can view invoices.
        :param pulumi.Input[bool] data_manage_datafeeds: Whether the team can modify data feeds.
        :param pulumi.Input[bool] data_manage_datasources: Whether the team can modify data sources.
        :param pulumi.Input[bool] data_push_to_datafeeds: Whether the team can publish to data feeds.
        :param pulumi.Input[bool] dhcp_manage_dhcp: Whether the team can manage DHCP.
               Only relevant for the DDI product.
        :param pulumi.Input[bool] dhcp_view_dhcp: Whether the team can view DHCP.
               Only relevant for the DDI product.
        :param pulumi.Input[bool] dns_manage_zones: Whether the team can modify the accounts zones.
        :param pulumi.Input[bool] dns_view_zones: Whether the team can view the accounts zones.
        :param pulumi.Input[bool] dns_zones_allow_by_default: If true, enable the `dns_zones_allow` list, otherwise enable the `dns_zones_deny` list.
        :param pulumi.Input[list] dns_zones_allows: List of zones that the team may access.
        :param pulumi.Input[list] dns_zones_denies: List of zones that the team may not access.
        :param pulumi.Input[list] ip_whitelists: The IP addresses to whitelist for this key.
        :param pulumi.Input[bool] ipam_manage_ipam: Whether the team can manage IPAM.
               Only relevant for the DDI product.
        :param pulumi.Input[bool] ipam_view_ipam: Whether the team can view IPAM.
               Only relevant for the DDI product.
        :param pulumi.Input[bool] monitoring_manage_jobs: Whether the team can modify monitoring jobs.
        :param pulumi.Input[bool] monitoring_manage_lists: Whether the team can modify notification lists.
        :param pulumi.Input[bool] monitoring_view_jobs: Whether the team can view monitoring jobs.
        :param pulumi.Input[str] name: The free form name of the team.
        :param pulumi.Input[bool] security_manage_active_directory: Whether the team can manage global active directory.
               Only relevant for the DDI product.
        :param pulumi.Input[bool] security_manage_global2fa: Whether the team can manage global two factor authentication.

        The **ip_whitelists** object supports the following:

          * `name` (`pulumi.Input[str]`) - The free form name of the team.
          * `values` (`pulumi.Input[list]`)
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['account_manage_account_settings'] = account_manage_account_settings
            __props__['account_manage_apikeys'] = account_manage_apikeys
            __props__['account_manage_payment_methods'] = account_manage_payment_methods
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
            __props__['ip_whitelists'] = ip_whitelists
            __props__['ipam_manage_ipam'] = ipam_manage_ipam
            __props__['ipam_view_ipam'] = ipam_view_ipam
            __props__['monitoring_manage_jobs'] = monitoring_manage_jobs
            __props__['monitoring_manage_lists'] = monitoring_manage_lists
            __props__['monitoring_view_jobs'] = monitoring_view_jobs
            __props__['name'] = name
            __props__['security_manage_active_directory'] = security_manage_active_directory
            __props__['security_manage_global2fa'] = security_manage_global2fa
        super(Team, __self__).__init__(
            'ns1:index/team:Team',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, account_manage_account_settings=None, account_manage_apikeys=None, account_manage_payment_methods=None, account_manage_plan=None, account_manage_teams=None, account_manage_users=None, account_view_activity_log=None, account_view_invoices=None, data_manage_datafeeds=None, data_manage_datasources=None, data_push_to_datafeeds=None, dhcp_manage_dhcp=None, dhcp_view_dhcp=None, dns_manage_zones=None, dns_view_zones=None, dns_zones_allow_by_default=None, dns_zones_allows=None, dns_zones_denies=None, ip_whitelists=None, ipam_manage_ipam=None, ipam_view_ipam=None, monitoring_manage_jobs=None, monitoring_manage_lists=None, monitoring_view_jobs=None, name=None, security_manage_active_directory=None, security_manage_global2fa=None):
        """
        Get an existing Team resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] account_manage_account_settings: Whether the team can modify account settings.
        :param pulumi.Input[bool] account_manage_apikeys: Whether the team can modify account apikeys.
        :param pulumi.Input[bool] account_manage_payment_methods: Whether the team can modify account payment methods.
        :param pulumi.Input[bool] account_manage_plan: Whether the team can modify the account plan.
        :param pulumi.Input[bool] account_manage_teams: Whether the team can modify other teams in the account.
        :param pulumi.Input[bool] account_manage_users: Whether the team can modify account users.
        :param pulumi.Input[bool] account_view_activity_log: Whether the team can view activity logs.
        :param pulumi.Input[bool] account_view_invoices: Whether the team can view invoices.
        :param pulumi.Input[bool] data_manage_datafeeds: Whether the team can modify data feeds.
        :param pulumi.Input[bool] data_manage_datasources: Whether the team can modify data sources.
        :param pulumi.Input[bool] data_push_to_datafeeds: Whether the team can publish to data feeds.
        :param pulumi.Input[bool] dhcp_manage_dhcp: Whether the team can manage DHCP.
               Only relevant for the DDI product.
        :param pulumi.Input[bool] dhcp_view_dhcp: Whether the team can view DHCP.
               Only relevant for the DDI product.
        :param pulumi.Input[bool] dns_manage_zones: Whether the team can modify the accounts zones.
        :param pulumi.Input[bool] dns_view_zones: Whether the team can view the accounts zones.
        :param pulumi.Input[bool] dns_zones_allow_by_default: If true, enable the `dns_zones_allow` list, otherwise enable the `dns_zones_deny` list.
        :param pulumi.Input[list] dns_zones_allows: List of zones that the team may access.
        :param pulumi.Input[list] dns_zones_denies: List of zones that the team may not access.
        :param pulumi.Input[list] ip_whitelists: The IP addresses to whitelist for this key.
        :param pulumi.Input[bool] ipam_manage_ipam: Whether the team can manage IPAM.
               Only relevant for the DDI product.
        :param pulumi.Input[bool] ipam_view_ipam: Whether the team can view IPAM.
               Only relevant for the DDI product.
        :param pulumi.Input[bool] monitoring_manage_jobs: Whether the team can modify monitoring jobs.
        :param pulumi.Input[bool] monitoring_manage_lists: Whether the team can modify notification lists.
        :param pulumi.Input[bool] monitoring_view_jobs: Whether the team can view monitoring jobs.
        :param pulumi.Input[str] name: The free form name of the team.
        :param pulumi.Input[bool] security_manage_active_directory: Whether the team can manage global active directory.
               Only relevant for the DDI product.
        :param pulumi.Input[bool] security_manage_global2fa: Whether the team can manage global two factor authentication.

        The **ip_whitelists** object supports the following:

          * `name` (`pulumi.Input[str]`) - The free form name of the team.
          * `values` (`pulumi.Input[list]`)
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
        __props__["ip_whitelists"] = ip_whitelists
        __props__["ipam_manage_ipam"] = ipam_manage_ipam
        __props__["ipam_view_ipam"] = ipam_view_ipam
        __props__["monitoring_manage_jobs"] = monitoring_manage_jobs
        __props__["monitoring_manage_lists"] = monitoring_manage_lists
        __props__["monitoring_view_jobs"] = monitoring_view_jobs
        __props__["name"] = name
        __props__["security_manage_active_directory"] = security_manage_active_directory
        __props__["security_manage_global2fa"] = security_manage_global2fa
        return Team(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

