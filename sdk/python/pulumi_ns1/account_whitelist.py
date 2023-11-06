# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['AccountWhitelistArgs', 'AccountWhitelist']

@pulumi.input_type
class AccountWhitelistArgs:
    def __init__(__self__, *,
                 values: pulumi.Input[Sequence[pulumi.Input[str]]],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AccountWhitelist resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] values: Array of IP addresses/networks from which to allow access.
        :param pulumi.Input[str] name: The free form name of the whitelist.
        """
        AccountWhitelistArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            values=values,
            name=name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if values is None:
            raise TypeError("Missing 'values' argument")

        _setter("values", values)
        if name is not None:
            _setter("name", name)

    @property
    @pulumi.getter
    def values(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Array of IP addresses/networks from which to allow access.
        """
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "values", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The free form name of the whitelist.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _AccountWhitelistState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering AccountWhitelist resources.
        :param pulumi.Input[str] name: The free form name of the whitelist.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] values: Array of IP addresses/networks from which to allow access.
        """
        _AccountWhitelistState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            name=name,
            values=values,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             name: Optional[pulumi.Input[str]] = None,
             values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):

        if name is not None:
            _setter("name", name)
        if values is not None:
            _setter("values", values)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The free form name of the whitelist.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def values(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Array of IP addresses/networks from which to allow access.
        """
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "values", value)


class AccountWhitelist(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a NS1 Global IP Whitelist resource.

        This can be used to create, modify, and delete Global IP Whitelists.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ns1 as ns1

        example = ns1.AccountWhitelist("example", values=[
            "1.1.1.1",
            "2.2.2.2",
        ])
        ```

        > You current source IP must be present in one of the whitelists to prevent locking yourself out.
        ## NS1 Documentation

        [Global IP Whitelist Doc](https://ns1.com/api?docId=2282)

        ## Import

        ```sh
         $ pulumi import ns1:index/accountWhitelist:AccountWhitelist example <ID>`
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The free form name of the whitelist.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] values: Array of IP addresses/networks from which to allow access.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AccountWhitelistArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a NS1 Global IP Whitelist resource.

        This can be used to create, modify, and delete Global IP Whitelists.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ns1 as ns1

        example = ns1.AccountWhitelist("example", values=[
            "1.1.1.1",
            "2.2.2.2",
        ])
        ```

        > You current source IP must be present in one of the whitelists to prevent locking yourself out.
        ## NS1 Documentation

        [Global IP Whitelist Doc](https://ns1.com/api?docId=2282)

        ## Import

        ```sh
         $ pulumi import ns1:index/accountWhitelist:AccountWhitelist example <ID>`
        ```

        :param str resource_name: The name of the resource.
        :param AccountWhitelistArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccountWhitelistArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            AccountWhitelistArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AccountWhitelistArgs.__new__(AccountWhitelistArgs)

            __props__.__dict__["name"] = name
            if values is None and not opts.urn:
                raise TypeError("Missing required property 'values'")
            __props__.__dict__["values"] = values
        super(AccountWhitelist, __self__).__init__(
            'ns1:index/accountWhitelist:AccountWhitelist',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            values: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'AccountWhitelist':
        """
        Get an existing AccountWhitelist resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The free form name of the whitelist.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] values: Array of IP addresses/networks from which to allow access.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AccountWhitelistState.__new__(_AccountWhitelistState)

        __props__.__dict__["name"] = name
        __props__.__dict__["values"] = values
        return AccountWhitelist(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The free form name of the whitelist.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def values(self) -> pulumi.Output[Sequence[str]]:
        """
        Array of IP addresses/networks from which to allow access.
        """
        return pulumi.get(self, "values")

