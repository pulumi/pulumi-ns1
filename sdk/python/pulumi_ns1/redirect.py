# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['RedirectArgs', 'Redirect']

@pulumi.input_type
class RedirectArgs:
    def __init__(__self__, *,
                 domain: pulumi.Input[str],
                 path: pulumi.Input[str],
                 target: pulumi.Input[str],
                 certificate_id: Optional[pulumi.Input[str]] = None,
                 forwarding_mode: Optional[pulumi.Input[str]] = None,
                 forwarding_type: Optional[pulumi.Input[str]] = None,
                 https_forced: Optional[pulumi.Input[bool]] = None,
                 query_forwarding: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Redirect resource.
        :param pulumi.Input[str] domain: The domain the redirect refers to.
        :param pulumi.Input[str] path: The path on the domain to redirect from.
        :param pulumi.Input[str] target: The URL to redirect to.
        :param pulumi.Input[str] certificate_id: The certificate redirect id.
        :param pulumi.Input[str] forwarding_mode: How the target is interpreted:
               * __all__       appends the entire incoming path to the target destination;
               * __capture__   appends only the part of the incoming path corresponding to the wildcard (*);
               * __none__      does not append any part of the incoming path.
        :param pulumi.Input[str] forwarding_type: How the redirect is executed:
               * __permanent__ (HTTP 301) indicates to search engines that they should remove the old page from
               their database and replace it with the new target page (this is recommended for SEO);
               * __temporary__ (HTTP 302) less common, indicates that search engines should keep the old domain or
               page indexed as the redirect is only temporary (while both pages might appear in the
               search results, a temporary redirect suggests to the search engine that it should
               prefer the new target page);
               * __masking__   preserves the redirected domain in the browser's address bar (this lets users see the
               address they entered, even though the displayed content comes from a different web page).
        :param pulumi.Input[bool] https_forced: Forces redirect for users that try to visit HTTP domain to HTTPS instead.
        :param pulumi.Input[bool] query_forwarding: Enables the query string of a URL to be applied directly to the new target URL.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Tags associated with the configuration.
        """
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "path", path)
        pulumi.set(__self__, "target", target)
        if certificate_id is not None:
            pulumi.set(__self__, "certificate_id", certificate_id)
        if forwarding_mode is not None:
            pulumi.set(__self__, "forwarding_mode", forwarding_mode)
        if forwarding_type is not None:
            pulumi.set(__self__, "forwarding_type", forwarding_type)
        if https_forced is not None:
            pulumi.set(__self__, "https_forced", https_forced)
        if query_forwarding is not None:
            pulumi.set(__self__, "query_forwarding", query_forwarding)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        """
        The domain the redirect refers to.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def path(self) -> pulumi.Input[str]:
        """
        The path on the domain to redirect from.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: pulumi.Input[str]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter
    def target(self) -> pulumi.Input[str]:
        """
        The URL to redirect to.
        """
        return pulumi.get(self, "target")

    @target.setter
    def target(self, value: pulumi.Input[str]):
        pulumi.set(self, "target", value)

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> Optional[pulumi.Input[str]]:
        """
        The certificate redirect id.
        """
        return pulumi.get(self, "certificate_id")

    @certificate_id.setter
    def certificate_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_id", value)

    @property
    @pulumi.getter(name="forwardingMode")
    def forwarding_mode(self) -> Optional[pulumi.Input[str]]:
        """
        How the target is interpreted:
        * __all__       appends the entire incoming path to the target destination;
        * __capture__   appends only the part of the incoming path corresponding to the wildcard (*);
        * __none__      does not append any part of the incoming path.
        """
        return pulumi.get(self, "forwarding_mode")

    @forwarding_mode.setter
    def forwarding_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "forwarding_mode", value)

    @property
    @pulumi.getter(name="forwardingType")
    def forwarding_type(self) -> Optional[pulumi.Input[str]]:
        """
        How the redirect is executed:
        * __permanent__ (HTTP 301) indicates to search engines that they should remove the old page from
        their database and replace it with the new target page (this is recommended for SEO);
        * __temporary__ (HTTP 302) less common, indicates that search engines should keep the old domain or
        page indexed as the redirect is only temporary (while both pages might appear in the
        search results, a temporary redirect suggests to the search engine that it should
        prefer the new target page);
        * __masking__   preserves the redirected domain in the browser's address bar (this lets users see the
        address they entered, even though the displayed content comes from a different web page).
        """
        return pulumi.get(self, "forwarding_type")

    @forwarding_type.setter
    def forwarding_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "forwarding_type", value)

    @property
    @pulumi.getter(name="httpsForced")
    def https_forced(self) -> Optional[pulumi.Input[bool]]:
        """
        Forces redirect for users that try to visit HTTP domain to HTTPS instead.
        """
        return pulumi.get(self, "https_forced")

    @https_forced.setter
    def https_forced(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "https_forced", value)

    @property
    @pulumi.getter(name="queryForwarding")
    def query_forwarding(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables the query string of a URL to be applied directly to the new target URL.
        """
        return pulumi.get(self, "query_forwarding")

    @query_forwarding.setter
    def query_forwarding(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "query_forwarding", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Tags associated with the configuration.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _RedirectState:
    def __init__(__self__, *,
                 certificate_id: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 forwarding_mode: Optional[pulumi.Input[str]] = None,
                 forwarding_type: Optional[pulumi.Input[str]] = None,
                 https_enabled: Optional[pulumi.Input[bool]] = None,
                 https_forced: Optional[pulumi.Input[bool]] = None,
                 last_updated: Optional[pulumi.Input[int]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 query_forwarding: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 target: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Redirect resources.
        :param pulumi.Input[str] certificate_id: The certificate redirect id.
        :param pulumi.Input[str] domain: The domain the redirect refers to.
        :param pulumi.Input[str] forwarding_mode: How the target is interpreted:
               * __all__       appends the entire incoming path to the target destination;
               * __capture__   appends only the part of the incoming path corresponding to the wildcard (*);
               * __none__      does not append any part of the incoming path.
        :param pulumi.Input[str] forwarding_type: How the redirect is executed:
               * __permanent__ (HTTP 301) indicates to search engines that they should remove the old page from
               their database and replace it with the new target page (this is recommended for SEO);
               * __temporary__ (HTTP 302) less common, indicates that search engines should keep the old domain or
               page indexed as the redirect is only temporary (while both pages might appear in the
               search results, a temporary redirect suggests to the search engine that it should
               prefer the new target page);
               * __masking__   preserves the redirected domain in the browser's address bar (this lets users see the
               address they entered, even though the displayed content comes from a different web page).
        :param pulumi.Input[bool] https_enabled: True if HTTPS is supported on the source domain by using Let's Encrypt certificates.
        :param pulumi.Input[bool] https_forced: Forces redirect for users that try to visit HTTP domain to HTTPS instead.
        :param pulumi.Input[int] last_updated: The Unix timestamp representing when the certificate was last signed.
        :param pulumi.Input[str] path: The path on the domain to redirect from.
        :param pulumi.Input[bool] query_forwarding: Enables the query string of a URL to be applied directly to the new target URL.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Tags associated with the configuration.
        :param pulumi.Input[str] target: The URL to redirect to.
        """
        if certificate_id is not None:
            pulumi.set(__self__, "certificate_id", certificate_id)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if forwarding_mode is not None:
            pulumi.set(__self__, "forwarding_mode", forwarding_mode)
        if forwarding_type is not None:
            pulumi.set(__self__, "forwarding_type", forwarding_type)
        if https_enabled is not None:
            pulumi.set(__self__, "https_enabled", https_enabled)
        if https_forced is not None:
            pulumi.set(__self__, "https_forced", https_forced)
        if last_updated is not None:
            pulumi.set(__self__, "last_updated", last_updated)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if query_forwarding is not None:
            pulumi.set(__self__, "query_forwarding", query_forwarding)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if target is not None:
            pulumi.set(__self__, "target", target)

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> Optional[pulumi.Input[str]]:
        """
        The certificate redirect id.
        """
        return pulumi.get(self, "certificate_id")

    @certificate_id.setter
    def certificate_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_id", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        The domain the redirect refers to.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="forwardingMode")
    def forwarding_mode(self) -> Optional[pulumi.Input[str]]:
        """
        How the target is interpreted:
        * __all__       appends the entire incoming path to the target destination;
        * __capture__   appends only the part of the incoming path corresponding to the wildcard (*);
        * __none__      does not append any part of the incoming path.
        """
        return pulumi.get(self, "forwarding_mode")

    @forwarding_mode.setter
    def forwarding_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "forwarding_mode", value)

    @property
    @pulumi.getter(name="forwardingType")
    def forwarding_type(self) -> Optional[pulumi.Input[str]]:
        """
        How the redirect is executed:
        * __permanent__ (HTTP 301) indicates to search engines that they should remove the old page from
        their database and replace it with the new target page (this is recommended for SEO);
        * __temporary__ (HTTP 302) less common, indicates that search engines should keep the old domain or
        page indexed as the redirect is only temporary (while both pages might appear in the
        search results, a temporary redirect suggests to the search engine that it should
        prefer the new target page);
        * __masking__   preserves the redirected domain in the browser's address bar (this lets users see the
        address they entered, even though the displayed content comes from a different web page).
        """
        return pulumi.get(self, "forwarding_type")

    @forwarding_type.setter
    def forwarding_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "forwarding_type", value)

    @property
    @pulumi.getter(name="httpsEnabled")
    def https_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        True if HTTPS is supported on the source domain by using Let's Encrypt certificates.
        """
        return pulumi.get(self, "https_enabled")

    @https_enabled.setter
    def https_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "https_enabled", value)

    @property
    @pulumi.getter(name="httpsForced")
    def https_forced(self) -> Optional[pulumi.Input[bool]]:
        """
        Forces redirect for users that try to visit HTTP domain to HTTPS instead.
        """
        return pulumi.get(self, "https_forced")

    @https_forced.setter
    def https_forced(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "https_forced", value)

    @property
    @pulumi.getter(name="lastUpdated")
    def last_updated(self) -> Optional[pulumi.Input[int]]:
        """
        The Unix timestamp representing when the certificate was last signed.
        """
        return pulumi.get(self, "last_updated")

    @last_updated.setter
    def last_updated(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "last_updated", value)

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[str]]:
        """
        The path on the domain to redirect from.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter(name="queryForwarding")
    def query_forwarding(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables the query string of a URL to be applied directly to the new target URL.
        """
        return pulumi.get(self, "query_forwarding")

    @query_forwarding.setter
    def query_forwarding(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "query_forwarding", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Tags associated with the configuration.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def target(self) -> Optional[pulumi.Input[str]]:
        """
        The URL to redirect to.
        """
        return pulumi.get(self, "target")

    @target.setter
    def target(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target", value)


class Redirect(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate_id: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 forwarding_mode: Optional[pulumi.Input[str]] = None,
                 forwarding_type: Optional[pulumi.Input[str]] = None,
                 https_forced: Optional[pulumi.Input[bool]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 query_forwarding: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a NS1 Redirect resource. This can be used to create, modify, and delete redirects.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ns1 as ns1

        example = ns1.Redirect("example",
            domain="www.example.com",
            path="/from/path",
            target="https://url.com/target/path")
        ```

        ### Additional Examples

        ```python
        import pulumi
        import pulumi_ns1 as ns1

        example = ns1.RedirectCertificate("example", domain="www.example.com")
        ```

        ## NS1 Documentation

        [Redirect Api Doc](https://ns1.com/api#redirect)

        # ns1\\_redirect\\_certificate

        Provides a NS1 Redirect Certificate resource. This can be used to create, modify, and delete redirect certificates.

        ## NS1 Documentation

        [Redirect Api Doc](https://ns1.com/api#redirect)

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_id: The certificate redirect id.
        :param pulumi.Input[str] domain: The domain the redirect refers to.
        :param pulumi.Input[str] forwarding_mode: How the target is interpreted:
               * __all__       appends the entire incoming path to the target destination;
               * __capture__   appends only the part of the incoming path corresponding to the wildcard (*);
               * __none__      does not append any part of the incoming path.
        :param pulumi.Input[str] forwarding_type: How the redirect is executed:
               * __permanent__ (HTTP 301) indicates to search engines that they should remove the old page from
               their database and replace it with the new target page (this is recommended for SEO);
               * __temporary__ (HTTP 302) less common, indicates that search engines should keep the old domain or
               page indexed as the redirect is only temporary (while both pages might appear in the
               search results, a temporary redirect suggests to the search engine that it should
               prefer the new target page);
               * __masking__   preserves the redirected domain in the browser's address bar (this lets users see the
               address they entered, even though the displayed content comes from a different web page).
        :param pulumi.Input[bool] https_forced: Forces redirect for users that try to visit HTTP domain to HTTPS instead.
        :param pulumi.Input[str] path: The path on the domain to redirect from.
        :param pulumi.Input[bool] query_forwarding: Enables the query string of a URL to be applied directly to the new target URL.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Tags associated with the configuration.
        :param pulumi.Input[str] target: The URL to redirect to.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RedirectArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a NS1 Redirect resource. This can be used to create, modify, and delete redirects.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ns1 as ns1

        example = ns1.Redirect("example",
            domain="www.example.com",
            path="/from/path",
            target="https://url.com/target/path")
        ```

        ### Additional Examples

        ```python
        import pulumi
        import pulumi_ns1 as ns1

        example = ns1.RedirectCertificate("example", domain="www.example.com")
        ```

        ## NS1 Documentation

        [Redirect Api Doc](https://ns1.com/api#redirect)

        # ns1\\_redirect\\_certificate

        Provides a NS1 Redirect Certificate resource. This can be used to create, modify, and delete redirect certificates.

        ## NS1 Documentation

        [Redirect Api Doc](https://ns1.com/api#redirect)

        :param str resource_name: The name of the resource.
        :param RedirectArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RedirectArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate_id: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 forwarding_mode: Optional[pulumi.Input[str]] = None,
                 forwarding_type: Optional[pulumi.Input[str]] = None,
                 https_forced: Optional[pulumi.Input[bool]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 query_forwarding: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 target: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RedirectArgs.__new__(RedirectArgs)

            __props__.__dict__["certificate_id"] = certificate_id
            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__.__dict__["domain"] = domain
            __props__.__dict__["forwarding_mode"] = forwarding_mode
            __props__.__dict__["forwarding_type"] = forwarding_type
            __props__.__dict__["https_forced"] = https_forced
            if path is None and not opts.urn:
                raise TypeError("Missing required property 'path'")
            __props__.__dict__["path"] = path
            __props__.__dict__["query_forwarding"] = query_forwarding
            __props__.__dict__["tags"] = tags
            if target is None and not opts.urn:
                raise TypeError("Missing required property 'target'")
            __props__.__dict__["target"] = target
            __props__.__dict__["https_enabled"] = None
            __props__.__dict__["last_updated"] = None
        super(Redirect, __self__).__init__(
            'ns1:index/redirect:Redirect',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            certificate_id: Optional[pulumi.Input[str]] = None,
            domain: Optional[pulumi.Input[str]] = None,
            forwarding_mode: Optional[pulumi.Input[str]] = None,
            forwarding_type: Optional[pulumi.Input[str]] = None,
            https_enabled: Optional[pulumi.Input[bool]] = None,
            https_forced: Optional[pulumi.Input[bool]] = None,
            last_updated: Optional[pulumi.Input[int]] = None,
            path: Optional[pulumi.Input[str]] = None,
            query_forwarding: Optional[pulumi.Input[bool]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            target: Optional[pulumi.Input[str]] = None) -> 'Redirect':
        """
        Get an existing Redirect resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_id: The certificate redirect id.
        :param pulumi.Input[str] domain: The domain the redirect refers to.
        :param pulumi.Input[str] forwarding_mode: How the target is interpreted:
               * __all__       appends the entire incoming path to the target destination;
               * __capture__   appends only the part of the incoming path corresponding to the wildcard (*);
               * __none__      does not append any part of the incoming path.
        :param pulumi.Input[str] forwarding_type: How the redirect is executed:
               * __permanent__ (HTTP 301) indicates to search engines that they should remove the old page from
               their database and replace it with the new target page (this is recommended for SEO);
               * __temporary__ (HTTP 302) less common, indicates that search engines should keep the old domain or
               page indexed as the redirect is only temporary (while both pages might appear in the
               search results, a temporary redirect suggests to the search engine that it should
               prefer the new target page);
               * __masking__   preserves the redirected domain in the browser's address bar (this lets users see the
               address they entered, even though the displayed content comes from a different web page).
        :param pulumi.Input[bool] https_enabled: True if HTTPS is supported on the source domain by using Let's Encrypt certificates.
        :param pulumi.Input[bool] https_forced: Forces redirect for users that try to visit HTTP domain to HTTPS instead.
        :param pulumi.Input[int] last_updated: The Unix timestamp representing when the certificate was last signed.
        :param pulumi.Input[str] path: The path on the domain to redirect from.
        :param pulumi.Input[bool] query_forwarding: Enables the query string of a URL to be applied directly to the new target URL.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: Tags associated with the configuration.
        :param pulumi.Input[str] target: The URL to redirect to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RedirectState.__new__(_RedirectState)

        __props__.__dict__["certificate_id"] = certificate_id
        __props__.__dict__["domain"] = domain
        __props__.__dict__["forwarding_mode"] = forwarding_mode
        __props__.__dict__["forwarding_type"] = forwarding_type
        __props__.__dict__["https_enabled"] = https_enabled
        __props__.__dict__["https_forced"] = https_forced
        __props__.__dict__["last_updated"] = last_updated
        __props__.__dict__["path"] = path
        __props__.__dict__["query_forwarding"] = query_forwarding
        __props__.__dict__["tags"] = tags
        __props__.__dict__["target"] = target
        return Redirect(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> pulumi.Output[str]:
        """
        The certificate redirect id.
        """
        return pulumi.get(self, "certificate_id")

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[str]:
        """
        The domain the redirect refers to.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="forwardingMode")
    def forwarding_mode(self) -> pulumi.Output[Optional[str]]:
        """
        How the target is interpreted:
        * __all__       appends the entire incoming path to the target destination;
        * __capture__   appends only the part of the incoming path corresponding to the wildcard (*);
        * __none__      does not append any part of the incoming path.
        """
        return pulumi.get(self, "forwarding_mode")

    @property
    @pulumi.getter(name="forwardingType")
    def forwarding_type(self) -> pulumi.Output[Optional[str]]:
        """
        How the redirect is executed:
        * __permanent__ (HTTP 301) indicates to search engines that they should remove the old page from
        their database and replace it with the new target page (this is recommended for SEO);
        * __temporary__ (HTTP 302) less common, indicates that search engines should keep the old domain or
        page indexed as the redirect is only temporary (while both pages might appear in the
        search results, a temporary redirect suggests to the search engine that it should
        prefer the new target page);
        * __masking__   preserves the redirected domain in the browser's address bar (this lets users see the
        address they entered, even though the displayed content comes from a different web page).
        """
        return pulumi.get(self, "forwarding_type")

    @property
    @pulumi.getter(name="httpsEnabled")
    def https_enabled(self) -> pulumi.Output[bool]:
        """
        True if HTTPS is supported on the source domain by using Let's Encrypt certificates.
        """
        return pulumi.get(self, "https_enabled")

    @property
    @pulumi.getter(name="httpsForced")
    def https_forced(self) -> pulumi.Output[bool]:
        """
        Forces redirect for users that try to visit HTTP domain to HTTPS instead.
        """
        return pulumi.get(self, "https_forced")

    @property
    @pulumi.getter(name="lastUpdated")
    def last_updated(self) -> pulumi.Output[int]:
        """
        The Unix timestamp representing when the certificate was last signed.
        """
        return pulumi.get(self, "last_updated")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[str]:
        """
        The path on the domain to redirect from.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="queryForwarding")
    def query_forwarding(self) -> pulumi.Output[Optional[bool]]:
        """
        Enables the query string of a URL to be applied directly to the new target URL.
        """
        return pulumi.get(self, "query_forwarding")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Tags associated with the configuration.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def target(self) -> pulumi.Output[str]:
        """
        The URL to redirect to.
        """
        return pulumi.get(self, "target")

