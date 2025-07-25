# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['RedirectCertificateArgs', 'RedirectCertificate']

@pulumi.input_type
class RedirectCertificateArgs:
    def __init__(__self__, *,
                 domain: pulumi.Input[_builtins.str]):
        """
        The set of arguments for constructing a RedirectCertificate resource.
        """
        pulumi.set(__self__, "domain", domain)

    @_builtins.property
    @pulumi.getter
    def domain(self) -> pulumi.Input[_builtins.str]:
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "domain", value)


@pulumi.input_type
class _RedirectCertificateState:
    def __init__(__self__, *,
                 certificate: Optional[pulumi.Input[_builtins.str]] = None,
                 domain: Optional[pulumi.Input[_builtins.str]] = None,
                 errors: Optional[pulumi.Input[_builtins.str]] = None,
                 last_updated: Optional[pulumi.Input[_builtins.int]] = None,
                 valid_from: Optional[pulumi.Input[_builtins.int]] = None,
                 valid_until: Optional[pulumi.Input[_builtins.int]] = None):
        """
        Input properties used for looking up and filtering RedirectCertificate resources.
        """
        if certificate is not None:
            pulumi.set(__self__, "certificate", certificate)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if errors is not None:
            pulumi.set(__self__, "errors", errors)
        if last_updated is not None:
            pulumi.set(__self__, "last_updated", last_updated)
        if valid_from is not None:
            pulumi.set(__self__, "valid_from", valid_from)
        if valid_until is not None:
            pulumi.set(__self__, "valid_until", valid_until)

    @_builtins.property
    @pulumi.getter
    def certificate(self) -> Optional[pulumi.Input[_builtins.str]]:
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "certificate", value)

    @_builtins.property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[_builtins.str]]:
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "domain", value)

    @_builtins.property
    @pulumi.getter
    def errors(self) -> Optional[pulumi.Input[_builtins.str]]:
        return pulumi.get(self, "errors")

    @errors.setter
    def errors(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "errors", value)

    @_builtins.property
    @pulumi.getter(name="lastUpdated")
    def last_updated(self) -> Optional[pulumi.Input[_builtins.int]]:
        return pulumi.get(self, "last_updated")

    @last_updated.setter
    def last_updated(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "last_updated", value)

    @_builtins.property
    @pulumi.getter(name="validFrom")
    def valid_from(self) -> Optional[pulumi.Input[_builtins.int]]:
        return pulumi.get(self, "valid_from")

    @valid_from.setter
    def valid_from(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "valid_from", value)

    @_builtins.property
    @pulumi.getter(name="validUntil")
    def valid_until(self) -> Optional[pulumi.Input[_builtins.int]]:
        return pulumi.get(self, "valid_until")

    @valid_until.setter
    def valid_until(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "valid_until", value)


@pulumi.type_token("ns1:index/redirectCertificate:RedirectCertificate")
class RedirectCertificate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        Create a RedirectCertificate resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RedirectCertificateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a RedirectCertificate resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param RedirectCertificateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RedirectCertificateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RedirectCertificateArgs.__new__(RedirectCertificateArgs)

            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__.__dict__["domain"] = domain
            __props__.__dict__["certificate"] = None
            __props__.__dict__["errors"] = None
            __props__.__dict__["last_updated"] = None
            __props__.__dict__["valid_from"] = None
            __props__.__dict__["valid_until"] = None
        super(RedirectCertificate, __self__).__init__(
            'ns1:index/redirectCertificate:RedirectCertificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            certificate: Optional[pulumi.Input[_builtins.str]] = None,
            domain: Optional[pulumi.Input[_builtins.str]] = None,
            errors: Optional[pulumi.Input[_builtins.str]] = None,
            last_updated: Optional[pulumi.Input[_builtins.int]] = None,
            valid_from: Optional[pulumi.Input[_builtins.int]] = None,
            valid_until: Optional[pulumi.Input[_builtins.int]] = None) -> 'RedirectCertificate':
        """
        Get an existing RedirectCertificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RedirectCertificateState.__new__(_RedirectCertificateState)

        __props__.__dict__["certificate"] = certificate
        __props__.__dict__["domain"] = domain
        __props__.__dict__["errors"] = errors
        __props__.__dict__["last_updated"] = last_updated
        __props__.__dict__["valid_from"] = valid_from
        __props__.__dict__["valid_until"] = valid_until
        return RedirectCertificate(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter
    def certificate(self) -> pulumi.Output[_builtins.str]:
        return pulumi.get(self, "certificate")

    @_builtins.property
    @pulumi.getter
    def domain(self) -> pulumi.Output[_builtins.str]:
        return pulumi.get(self, "domain")

    @_builtins.property
    @pulumi.getter
    def errors(self) -> pulumi.Output[_builtins.str]:
        return pulumi.get(self, "errors")

    @_builtins.property
    @pulumi.getter(name="lastUpdated")
    def last_updated(self) -> pulumi.Output[_builtins.int]:
        return pulumi.get(self, "last_updated")

    @_builtins.property
    @pulumi.getter(name="validFrom")
    def valid_from(self) -> pulumi.Output[_builtins.int]:
        return pulumi.get(self, "valid_from")

    @_builtins.property
    @pulumi.getter(name="validUntil")
    def valid_until(self) -> pulumi.Output[_builtins.int]:
        return pulumi.get(self, "valid_until")

