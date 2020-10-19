# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['Provider']


class Provider(pulumi.ProviderResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 apikey: Optional[pulumi.Input[str]] = None,
                 enable_ddi: Optional[pulumi.Input[bool]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 ignore_ssl: Optional[pulumi.Input[bool]] = None,
                 rate_limit_parallelism: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The provider type for the ns1 package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] apikey: The ns1 API key, this is required
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

            if apikey is None:
                apikey = _utilities.get_env('NS1_APIKEY')
            __props__['apikey'] = apikey
            __props__['enable_ddi'] = pulumi.Output.from_input(enable_ddi).apply(pulumi.runtime.to_json) if enable_ddi is not None else None
            if endpoint is None:
                endpoint = _utilities.get_env('NS1_ENDPOINT')
            __props__['endpoint'] = endpoint
            __props__['ignore_ssl'] = pulumi.Output.from_input(ignore_ssl).apply(pulumi.runtime.to_json) if ignore_ssl is not None else None
            __props__['rate_limit_parallelism'] = pulumi.Output.from_input(rate_limit_parallelism).apply(pulumi.runtime.to_json) if rate_limit_parallelism is not None else None
        super(Provider, __self__).__init__(
            'ns1',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

