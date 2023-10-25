# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['DataSourceArgs', 'DataSource']

@pulumi.input_type
class DataSourceArgs:
    def __init__(__self__, *,
                 sourcetype: pulumi.Input[str],
                 config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DataSource resource.
        :param pulumi.Input[str] sourcetype: The data sources type, listed in API endpoint https://api.nsone.net/v1/data/sourcetypes.
        :param pulumi.Input[Mapping[str, Any]] config: The data source configuration, determined by its type,
               matching the specification in `config` from /data/sourcetypes.
        :param pulumi.Input[str] name: The free form name of the data source.
        """
        DataSourceArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            sourcetype=sourcetype,
            config=config,
            name=name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             sourcetype: Optional[pulumi.Input[str]] = None,
             config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if sourcetype is None:
            raise TypeError("Missing 'sourcetype' argument")

        _setter("sourcetype", sourcetype)
        if config is not None:
            _setter("config", config)
        if name is not None:
            _setter("name", name)

    @property
    @pulumi.getter
    def sourcetype(self) -> pulumi.Input[str]:
        """
        The data sources type, listed in API endpoint https://api.nsone.net/v1/data/sourcetypes.
        """
        return pulumi.get(self, "sourcetype")

    @sourcetype.setter
    def sourcetype(self, value: pulumi.Input[str]):
        pulumi.set(self, "sourcetype", value)

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The data source configuration, determined by its type,
        matching the specification in `config` from /data/sourcetypes.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The free form name of the data source.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _DataSourceState:
    def __init__(__self__, *,
                 config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 sourcetype: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DataSource resources.
        :param pulumi.Input[Mapping[str, Any]] config: The data source configuration, determined by its type,
               matching the specification in `config` from /data/sourcetypes.
        :param pulumi.Input[str] name: The free form name of the data source.
        :param pulumi.Input[str] sourcetype: The data sources type, listed in API endpoint https://api.nsone.net/v1/data/sourcetypes.
        """
        _DataSourceState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            config=config,
            name=name,
            sourcetype=sourcetype,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             name: Optional[pulumi.Input[str]] = None,
             sourcetype: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):

        if config is not None:
            _setter("config", config)
        if name is not None:
            _setter("name", name)
        if sourcetype is not None:
            _setter("sourcetype", sourcetype)

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        The data source configuration, determined by its type,
        matching the specification in `config` from /data/sourcetypes.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The free form name of the data source.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def sourcetype(self) -> Optional[pulumi.Input[str]]:
        """
        The data sources type, listed in API endpoint https://api.nsone.net/v1/data/sourcetypes.
        """
        return pulumi.get(self, "sourcetype")

    @sourcetype.setter
    def sourcetype(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sourcetype", value)


class DataSource(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 sourcetype: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a NS1 Data Source resource. This can be used to create, modify, and delete data sources.

        ## NS1 Documentation

        [Datasource Api Doc](https://ns1.com/api#data-sources)

        ## Import

        ```sh
         $ pulumi import ns1:index/dataSource:DataSource <name> <datasource_id>`
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, Any]] config: The data source configuration, determined by its type,
               matching the specification in `config` from /data/sourcetypes.
        :param pulumi.Input[str] name: The free form name of the data source.
        :param pulumi.Input[str] sourcetype: The data sources type, listed in API endpoint https://api.nsone.net/v1/data/sourcetypes.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DataSourceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a NS1 Data Source resource. This can be used to create, modify, and delete data sources.

        ## NS1 Documentation

        [Datasource Api Doc](https://ns1.com/api#data-sources)

        ## Import

        ```sh
         $ pulumi import ns1:index/dataSource:DataSource <name> <datasource_id>`
        ```

        :param str resource_name: The name of the resource.
        :param DataSourceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DataSourceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            DataSourceArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 sourcetype: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DataSourceArgs.__new__(DataSourceArgs)

            __props__.__dict__["config"] = config
            __props__.__dict__["name"] = name
            if sourcetype is None and not opts.urn:
                raise TypeError("Missing required property 'sourcetype'")
            __props__.__dict__["sourcetype"] = sourcetype
        super(DataSource, __self__).__init__(
            'ns1:index/dataSource:DataSource',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            sourcetype: Optional[pulumi.Input[str]] = None) -> 'DataSource':
        """
        Get an existing DataSource resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, Any]] config: The data source configuration, determined by its type,
               matching the specification in `config` from /data/sourcetypes.
        :param pulumi.Input[str] name: The free form name of the data source.
        :param pulumi.Input[str] sourcetype: The data sources type, listed in API endpoint https://api.nsone.net/v1/data/sourcetypes.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DataSourceState.__new__(_DataSourceState)

        __props__.__dict__["config"] = config
        __props__.__dict__["name"] = name
        __props__.__dict__["sourcetype"] = sourcetype
        return DataSource(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        The data source configuration, determined by its type,
        matching the specification in `config` from /data/sourcetypes.
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The free form name of the data source.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def sourcetype(self) -> pulumi.Output[str]:
        """
        The data sources type, listed in API endpoint https://api.nsone.net/v1/data/sourcetypes.
        """
        return pulumi.get(self, "sourcetype")

