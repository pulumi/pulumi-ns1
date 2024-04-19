# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['TsigkeyArgs', 'Tsigkey']

@pulumi.input_type
class TsigkeyArgs:
    def __init__(__self__, *,
                 algorithm: pulumi.Input[str],
                 secret: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Tsigkey resource.
        :param pulumi.Input[str] algorithm: The algorithm used to hash the TSIG key's secret.
        :param pulumi.Input[str] secret: The key's secret to be hashed.
        :param pulumi.Input[str] name: The free form name of the tsigkey.
        """
        pulumi.set(__self__, "algorithm", algorithm)
        pulumi.set(__self__, "secret", secret)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def algorithm(self) -> pulumi.Input[str]:
        """
        The algorithm used to hash the TSIG key's secret.
        """
        return pulumi.get(self, "algorithm")

    @algorithm.setter
    def algorithm(self, value: pulumi.Input[str]):
        pulumi.set(self, "algorithm", value)

    @property
    @pulumi.getter
    def secret(self) -> pulumi.Input[str]:
        """
        The key's secret to be hashed.
        """
        return pulumi.get(self, "secret")

    @secret.setter
    def secret(self, value: pulumi.Input[str]):
        pulumi.set(self, "secret", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The free form name of the tsigkey.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _TsigkeyState:
    def __init__(__self__, *,
                 algorithm: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 secret: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Tsigkey resources.
        :param pulumi.Input[str] algorithm: The algorithm used to hash the TSIG key's secret.
        :param pulumi.Input[str] name: The free form name of the tsigkey.
        :param pulumi.Input[str] secret: The key's secret to be hashed.
        """
        if algorithm is not None:
            pulumi.set(__self__, "algorithm", algorithm)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if secret is not None:
            pulumi.set(__self__, "secret", secret)

    @property
    @pulumi.getter
    def algorithm(self) -> Optional[pulumi.Input[str]]:
        """
        The algorithm used to hash the TSIG key's secret.
        """
        return pulumi.get(self, "algorithm")

    @algorithm.setter
    def algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "algorithm", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The free form name of the tsigkey.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def secret(self) -> Optional[pulumi.Input[str]]:
        """
        The key's secret to be hashed.
        """
        return pulumi.get(self, "secret")

    @secret.setter
    def secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret", value)


class Tsigkey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 algorithm: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 secret: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a NS1 TSIG Key resource. This can be used to create, modify, and delete TSIG keys.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_ns1 as ns1

        example = ns1.Tsigkey("example",
            name="ExampleTsigKey",
            algorithm="hmac-sha256",
            secret="Ok1qR5IW1ajVka5cHPEJQIXfLyx5V3PSkFBROAzOn21JumDq6nIpoj6H8rfj5Uo+Ok55ZWQ0Wgrf302fDscHLA==")
        ```
        <!--End PulumiCodeChooser -->
        ## NS1 Documentation

        [TSIG Keys Api Doc](https://ns1.com/api/#tsig)

        ## Import

        ```sh
        $ pulumi import ns1:index/tsigkey:Tsigkey importTest <name>`
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] algorithm: The algorithm used to hash the TSIG key's secret.
        :param pulumi.Input[str] name: The free form name of the tsigkey.
        :param pulumi.Input[str] secret: The key's secret to be hashed.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TsigkeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a NS1 TSIG Key resource. This can be used to create, modify, and delete TSIG keys.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_ns1 as ns1

        example = ns1.Tsigkey("example",
            name="ExampleTsigKey",
            algorithm="hmac-sha256",
            secret="Ok1qR5IW1ajVka5cHPEJQIXfLyx5V3PSkFBROAzOn21JumDq6nIpoj6H8rfj5Uo+Ok55ZWQ0Wgrf302fDscHLA==")
        ```
        <!--End PulumiCodeChooser -->
        ## NS1 Documentation

        [TSIG Keys Api Doc](https://ns1.com/api/#tsig)

        ## Import

        ```sh
        $ pulumi import ns1:index/tsigkey:Tsigkey importTest <name>`
        ```

        :param str resource_name: The name of the resource.
        :param TsigkeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TsigkeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 algorithm: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 secret: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TsigkeyArgs.__new__(TsigkeyArgs)

            if algorithm is None and not opts.urn:
                raise TypeError("Missing required property 'algorithm'")
            __props__.__dict__["algorithm"] = algorithm
            __props__.__dict__["name"] = name
            if secret is None and not opts.urn:
                raise TypeError("Missing required property 'secret'")
            __props__.__dict__["secret"] = secret
        super(Tsigkey, __self__).__init__(
            'ns1:index/tsigkey:Tsigkey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            algorithm: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            secret: Optional[pulumi.Input[str]] = None) -> 'Tsigkey':
        """
        Get an existing Tsigkey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] algorithm: The algorithm used to hash the TSIG key's secret.
        :param pulumi.Input[str] name: The free form name of the tsigkey.
        :param pulumi.Input[str] secret: The key's secret to be hashed.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TsigkeyState.__new__(_TsigkeyState)

        __props__.__dict__["algorithm"] = algorithm
        __props__.__dict__["name"] = name
        __props__.__dict__["secret"] = secret
        return Tsigkey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def algorithm(self) -> pulumi.Output[str]:
        """
        The algorithm used to hash the TSIG key's secret.
        """
        return pulumi.get(self, "algorithm")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The free form name of the tsigkey.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def secret(self) -> pulumi.Output[str]:
        """
        The key's secret to be hashed.
        """
        return pulumi.get(self, "secret")

