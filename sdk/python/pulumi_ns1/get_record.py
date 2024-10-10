# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
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
from . import outputs

__all__ = [
    'GetRecordResult',
    'AwaitableGetRecordResult',
    'get_record',
    'get_record_output',
]

@pulumi.output_type
class GetRecordResult:
    """
    A collection of values returned by getRecord.
    """
    def __init__(__self__, answers=None, domain=None, filters=None, id=None, link=None, meta=None, override_ttl=None, regions=None, short_answers=None, tags=None, ttl=None, type=None, use_client_subnet=None, zone=None):
        if answers and not isinstance(answers, list):
            raise TypeError("Expected argument 'answers' to be a list")
        pulumi.set(__self__, "answers", answers)
        if domain and not isinstance(domain, str):
            raise TypeError("Expected argument 'domain' to be a str")
        pulumi.set(__self__, "domain", domain)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if link and not isinstance(link, str):
            raise TypeError("Expected argument 'link' to be a str")
        pulumi.set(__self__, "link", link)
        if meta and not isinstance(meta, dict):
            raise TypeError("Expected argument 'meta' to be a dict")
        pulumi.set(__self__, "meta", meta)
        if override_ttl and not isinstance(override_ttl, bool):
            raise TypeError("Expected argument 'override_ttl' to be a bool")
        pulumi.set(__self__, "override_ttl", override_ttl)
        if regions and not isinstance(regions, list):
            raise TypeError("Expected argument 'regions' to be a list")
        pulumi.set(__self__, "regions", regions)
        if short_answers and not isinstance(short_answers, list):
            raise TypeError("Expected argument 'short_answers' to be a list")
        pulumi.set(__self__, "short_answers", short_answers)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if ttl and not isinstance(ttl, int):
            raise TypeError("Expected argument 'ttl' to be a int")
        pulumi.set(__self__, "ttl", ttl)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if use_client_subnet and not isinstance(use_client_subnet, bool):
            raise TypeError("Expected argument 'use_client_subnet' to be a bool")
        pulumi.set(__self__, "use_client_subnet", use_client_subnet)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def answers(self) -> Sequence['outputs.GetRecordAnswerResult']:
        """
        List of NS1 answers.
        """
        return pulumi.get(self, "answers")

    @property
    @pulumi.getter
    def domain(self) -> str:
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def filters(self) -> Sequence['outputs.GetRecordFilterResult']:
        """
        List of NS1 filters.
        """
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def link(self) -> str:
        """
        The target record this links to.
        """
        return pulumi.get(self, "link")

    @property
    @pulumi.getter
    def meta(self) -> Mapping[str, str]:
        """
        Map of metadata
        """
        return pulumi.get(self, "meta")

    @property
    @pulumi.getter(name="overrideTtl")
    def override_ttl(self) -> bool:
        return pulumi.get(self, "override_ttl")

    @property
    @pulumi.getter
    def regions(self) -> Sequence['outputs.GetRecordRegionResult']:
        """
        List of regions.
        """
        return pulumi.get(self, "regions")

    @property
    @pulumi.getter(name="shortAnswers")
    def short_answers(self) -> Sequence[str]:
        return pulumi.get(self, "short_answers")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def ttl(self) -> int:
        """
        The records' time to live (in seconds).
        """
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="useClientSubnet")
    def use_client_subnet(self) -> bool:
        """
        Whether to use EDNS client subnet data when available (in filter chain).
        """
        return pulumi.get(self, "use_client_subnet")

    @property
    @pulumi.getter
    def zone(self) -> str:
        return pulumi.get(self, "zone")


class AwaitableGetRecordResult(GetRecordResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRecordResult(
            answers=self.answers,
            domain=self.domain,
            filters=self.filters,
            id=self.id,
            link=self.link,
            meta=self.meta,
            override_ttl=self.override_ttl,
            regions=self.regions,
            short_answers=self.short_answers,
            tags=self.tags,
            ttl=self.ttl,
            type=self.type,
            use_client_subnet=self.use_client_subnet,
            zone=self.zone)


def get_record(domain: Optional[str] = None,
               type: Optional[str] = None,
               zone: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRecordResult:
    """
    Provides details about a NS1 Record. Use this if you would simply like to read
    information from NS1 into your configurations. For read/write operations, you
    should use a resource.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ns1 as ns1

    # Get details about a NS1 Record.
    example = ns1.get_record(zone="example.io",
        domain="terraform.example.io",
        type="A")
    ```


    :param str domain: The records' domain.
    :param str type: The records' RR type.
    :param str zone: The zone the record belongs to.
    """
    __args__ = dict()
    __args__['domain'] = domain
    __args__['type'] = type
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ns1:index/getRecord:getRecord', __args__, opts=opts, typ=GetRecordResult).value

    return AwaitableGetRecordResult(
        answers=pulumi.get(__ret__, 'answers'),
        domain=pulumi.get(__ret__, 'domain'),
        filters=pulumi.get(__ret__, 'filters'),
        id=pulumi.get(__ret__, 'id'),
        link=pulumi.get(__ret__, 'link'),
        meta=pulumi.get(__ret__, 'meta'),
        override_ttl=pulumi.get(__ret__, 'override_ttl'),
        regions=pulumi.get(__ret__, 'regions'),
        short_answers=pulumi.get(__ret__, 'short_answers'),
        tags=pulumi.get(__ret__, 'tags'),
        ttl=pulumi.get(__ret__, 'ttl'),
        type=pulumi.get(__ret__, 'type'),
        use_client_subnet=pulumi.get(__ret__, 'use_client_subnet'),
        zone=pulumi.get(__ret__, 'zone'))
def get_record_output(domain: Optional[pulumi.Input[str]] = None,
                      type: Optional[pulumi.Input[str]] = None,
                      zone: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRecordResult]:
    """
    Provides details about a NS1 Record. Use this if you would simply like to read
    information from NS1 into your configurations. For read/write operations, you
    should use a resource.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ns1 as ns1

    # Get details about a NS1 Record.
    example = ns1.get_record(zone="example.io",
        domain="terraform.example.io",
        type="A")
    ```


    :param str domain: The records' domain.
    :param str type: The records' RR type.
    :param str zone: The zone the record belongs to.
    """
    __args__ = dict()
    __args__['domain'] = domain
    __args__['type'] = type
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ns1:index/getRecord:getRecord', __args__, opts=opts, typ=GetRecordResult)
    return __ret__.apply(lambda __response__: GetRecordResult(
        answers=pulumi.get(__response__, 'answers'),
        domain=pulumi.get(__response__, 'domain'),
        filters=pulumi.get(__response__, 'filters'),
        id=pulumi.get(__response__, 'id'),
        link=pulumi.get(__response__, 'link'),
        meta=pulumi.get(__response__, 'meta'),
        override_ttl=pulumi.get(__response__, 'override_ttl'),
        regions=pulumi.get(__response__, 'regions'),
        short_answers=pulumi.get(__response__, 'short_answers'),
        tags=pulumi.get(__response__, 'tags'),
        ttl=pulumi.get(__response__, 'ttl'),
        type=pulumi.get(__response__, 'type'),
        use_client_subnet=pulumi.get(__response__, 'use_client_subnet'),
        zone=pulumi.get(__response__, 'zone')))
