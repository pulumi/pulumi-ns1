# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SubnetArgs', 'Subnet']

@pulumi.input_type
class SubnetArgs:
    def __init__(__self__, *,
                 children: Optional[pulumi.Input[int]] = None,
                 desc: Optional[pulumi.Input[str]] = None,
                 dhcp_scoped: Optional[pulumi.Input[bool]] = None,
                 free_addresses: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[int]] = None,
                 parent_id: Optional[pulumi.Input[int]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 total_addresses: Optional[pulumi.Input[str]] = None,
                 used_addresses: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Subnet resource.
        """
        SubnetArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            children=children,
            desc=desc,
            dhcp_scoped=dhcp_scoped,
            free_addresses=free_addresses,
            name=name,
            network_id=network_id,
            parent_id=parent_id,
            prefix=prefix,
            status=status,
            tags=tags,
            total_addresses=total_addresses,
            used_addresses=used_addresses,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             children: Optional[pulumi.Input[int]] = None,
             desc: Optional[pulumi.Input[str]] = None,
             dhcp_scoped: Optional[pulumi.Input[bool]] = None,
             free_addresses: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             network_id: Optional[pulumi.Input[int]] = None,
             parent_id: Optional[pulumi.Input[int]] = None,
             prefix: Optional[pulumi.Input[str]] = None,
             status: Optional[pulumi.Input[str]] = None,
             tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             total_addresses: Optional[pulumi.Input[str]] = None,
             used_addresses: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if dhcp_scoped is None and 'dhcpScoped' in kwargs:
            dhcp_scoped = kwargs['dhcpScoped']
        if free_addresses is None and 'freeAddresses' in kwargs:
            free_addresses = kwargs['freeAddresses']
        if network_id is None and 'networkId' in kwargs:
            network_id = kwargs['networkId']
        if parent_id is None and 'parentId' in kwargs:
            parent_id = kwargs['parentId']
        if total_addresses is None and 'totalAddresses' in kwargs:
            total_addresses = kwargs['totalAddresses']
        if used_addresses is None and 'usedAddresses' in kwargs:
            used_addresses = kwargs['usedAddresses']

        if children is not None:
            _setter("children", children)
        if desc is not None:
            _setter("desc", desc)
        if dhcp_scoped is not None:
            _setter("dhcp_scoped", dhcp_scoped)
        if free_addresses is not None:
            _setter("free_addresses", free_addresses)
        if name is not None:
            _setter("name", name)
        if network_id is not None:
            _setter("network_id", network_id)
        if parent_id is not None:
            _setter("parent_id", parent_id)
        if prefix is not None:
            _setter("prefix", prefix)
        if status is not None:
            _setter("status", status)
        if tags is not None:
            _setter("tags", tags)
        if total_addresses is not None:
            _setter("total_addresses", total_addresses)
        if used_addresses is not None:
            _setter("used_addresses", used_addresses)

    @property
    @pulumi.getter
    def children(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "children")

    @children.setter
    def children(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "children", value)

    @property
    @pulumi.getter
    def desc(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "desc")

    @desc.setter
    def desc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "desc", value)

    @property
    @pulumi.getter(name="dhcpScoped")
    def dhcp_scoped(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "dhcp_scoped")

    @dhcp_scoped.setter
    def dhcp_scoped(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dhcp_scoped", value)

    @property
    @pulumi.getter(name="freeAddresses")
    def free_addresses(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "free_addresses")

    @free_addresses.setter
    def free_addresses(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "free_addresses", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "network_id", value)

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "parent_id")

    @parent_id.setter
    def parent_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "parent_id", value)

    @property
    @pulumi.getter
    def prefix(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "prefix")

    @prefix.setter
    def prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefix", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="totalAddresses")
    def total_addresses(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "total_addresses")

    @total_addresses.setter
    def total_addresses(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "total_addresses", value)

    @property
    @pulumi.getter(name="usedAddresses")
    def used_addresses(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "used_addresses")

    @used_addresses.setter
    def used_addresses(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "used_addresses", value)


@pulumi.input_type
class _SubnetState:
    def __init__(__self__, *,
                 children: Optional[pulumi.Input[int]] = None,
                 desc: Optional[pulumi.Input[str]] = None,
                 dhcp_scoped: Optional[pulumi.Input[bool]] = None,
                 free_addresses: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[int]] = None,
                 parent_id: Optional[pulumi.Input[int]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 total_addresses: Optional[pulumi.Input[str]] = None,
                 used_addresses: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Subnet resources.
        """
        _SubnetState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            children=children,
            desc=desc,
            dhcp_scoped=dhcp_scoped,
            free_addresses=free_addresses,
            name=name,
            network_id=network_id,
            parent_id=parent_id,
            prefix=prefix,
            status=status,
            tags=tags,
            total_addresses=total_addresses,
            used_addresses=used_addresses,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             children: Optional[pulumi.Input[int]] = None,
             desc: Optional[pulumi.Input[str]] = None,
             dhcp_scoped: Optional[pulumi.Input[bool]] = None,
             free_addresses: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             network_id: Optional[pulumi.Input[int]] = None,
             parent_id: Optional[pulumi.Input[int]] = None,
             prefix: Optional[pulumi.Input[str]] = None,
             status: Optional[pulumi.Input[str]] = None,
             tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             total_addresses: Optional[pulumi.Input[str]] = None,
             used_addresses: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if dhcp_scoped is None and 'dhcpScoped' in kwargs:
            dhcp_scoped = kwargs['dhcpScoped']
        if free_addresses is None and 'freeAddresses' in kwargs:
            free_addresses = kwargs['freeAddresses']
        if network_id is None and 'networkId' in kwargs:
            network_id = kwargs['networkId']
        if parent_id is None and 'parentId' in kwargs:
            parent_id = kwargs['parentId']
        if total_addresses is None and 'totalAddresses' in kwargs:
            total_addresses = kwargs['totalAddresses']
        if used_addresses is None and 'usedAddresses' in kwargs:
            used_addresses = kwargs['usedAddresses']

        if children is not None:
            _setter("children", children)
        if desc is not None:
            _setter("desc", desc)
        if dhcp_scoped is not None:
            _setter("dhcp_scoped", dhcp_scoped)
        if free_addresses is not None:
            _setter("free_addresses", free_addresses)
        if name is not None:
            _setter("name", name)
        if network_id is not None:
            _setter("network_id", network_id)
        if parent_id is not None:
            _setter("parent_id", parent_id)
        if prefix is not None:
            _setter("prefix", prefix)
        if status is not None:
            _setter("status", status)
        if tags is not None:
            _setter("tags", tags)
        if total_addresses is not None:
            _setter("total_addresses", total_addresses)
        if used_addresses is not None:
            _setter("used_addresses", used_addresses)

    @property
    @pulumi.getter
    def children(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "children")

    @children.setter
    def children(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "children", value)

    @property
    @pulumi.getter
    def desc(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "desc")

    @desc.setter
    def desc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "desc", value)

    @property
    @pulumi.getter(name="dhcpScoped")
    def dhcp_scoped(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "dhcp_scoped")

    @dhcp_scoped.setter
    def dhcp_scoped(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dhcp_scoped", value)

    @property
    @pulumi.getter(name="freeAddresses")
    def free_addresses(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "free_addresses")

    @free_addresses.setter
    def free_addresses(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "free_addresses", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "network_id", value)

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "parent_id")

    @parent_id.setter
    def parent_id(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "parent_id", value)

    @property
    @pulumi.getter
    def prefix(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "prefix")

    @prefix.setter
    def prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefix", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="totalAddresses")
    def total_addresses(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "total_addresses")

    @total_addresses.setter
    def total_addresses(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "total_addresses", value)

    @property
    @pulumi.getter(name="usedAddresses")
    def used_addresses(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "used_addresses")

    @used_addresses.setter
    def used_addresses(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "used_addresses", value)


class Subnet(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 children: Optional[pulumi.Input[int]] = None,
                 desc: Optional[pulumi.Input[str]] = None,
                 dhcp_scoped: Optional[pulumi.Input[bool]] = None,
                 free_addresses: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[int]] = None,
                 parent_id: Optional[pulumi.Input[int]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 total_addresses: Optional[pulumi.Input[str]] = None,
                 used_addresses: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a Subnet resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SubnetArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Subnet resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SubnetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SubnetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            SubnetArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 children: Optional[pulumi.Input[int]] = None,
                 desc: Optional[pulumi.Input[str]] = None,
                 dhcp_scoped: Optional[pulumi.Input[bool]] = None,
                 free_addresses: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_id: Optional[pulumi.Input[int]] = None,
                 parent_id: Optional[pulumi.Input[int]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 total_addresses: Optional[pulumi.Input[str]] = None,
                 used_addresses: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SubnetArgs.__new__(SubnetArgs)

            __props__.__dict__["children"] = children
            __props__.__dict__["desc"] = desc
            __props__.__dict__["dhcp_scoped"] = dhcp_scoped
            __props__.__dict__["free_addresses"] = free_addresses
            __props__.__dict__["name"] = name
            __props__.__dict__["network_id"] = network_id
            __props__.__dict__["parent_id"] = parent_id
            __props__.__dict__["prefix"] = prefix
            __props__.__dict__["status"] = status
            __props__.__dict__["tags"] = tags
            __props__.__dict__["total_addresses"] = total_addresses
            __props__.__dict__["used_addresses"] = used_addresses
        super(Subnet, __self__).__init__(
            'ns1:index/subnet:Subnet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            children: Optional[pulumi.Input[int]] = None,
            desc: Optional[pulumi.Input[str]] = None,
            dhcp_scoped: Optional[pulumi.Input[bool]] = None,
            free_addresses: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_id: Optional[pulumi.Input[int]] = None,
            parent_id: Optional[pulumi.Input[int]] = None,
            prefix: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            total_addresses: Optional[pulumi.Input[str]] = None,
            used_addresses: Optional[pulumi.Input[str]] = None) -> 'Subnet':
        """
        Get an existing Subnet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SubnetState.__new__(_SubnetState)

        __props__.__dict__["children"] = children
        __props__.__dict__["desc"] = desc
        __props__.__dict__["dhcp_scoped"] = dhcp_scoped
        __props__.__dict__["free_addresses"] = free_addresses
        __props__.__dict__["name"] = name
        __props__.__dict__["network_id"] = network_id
        __props__.__dict__["parent_id"] = parent_id
        __props__.__dict__["prefix"] = prefix
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["total_addresses"] = total_addresses
        __props__.__dict__["used_addresses"] = used_addresses
        return Subnet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def children(self) -> pulumi.Output[int]:
        return pulumi.get(self, "children")

    @property
    @pulumi.getter
    def desc(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "desc")

    @property
    @pulumi.getter(name="dhcpScoped")
    def dhcp_scoped(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "dhcp_scoped")

    @property
    @pulumi.getter(name="freeAddresses")
    def free_addresses(self) -> pulumi.Output[str]:
        return pulumi.get(self, "free_addresses")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> pulumi.Output[Optional[int]]:
        return pulumi.get(self, "network_id")

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> pulumi.Output[int]:
        return pulumi.get(self, "parent_id")

    @property
    @pulumi.getter
    def prefix(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "prefix")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Mapping[str, Any]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="totalAddresses")
    def total_addresses(self) -> pulumi.Output[str]:
        return pulumi.get(self, "total_addresses")

    @property
    @pulumi.getter(name="usedAddresses")
    def used_addresses(self) -> pulumi.Output[str]:
        return pulumi.get(self, "used_addresses")

