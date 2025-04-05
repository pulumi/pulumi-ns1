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
    'GetBillingUsageResult',
    'AwaitableGetBillingUsageResult',
    'get_billing_usage',
    'get_billing_usage_output',
]

@pulumi.output_type
class GetBillingUsageResult:
    """
    A collection of values returned by getBillingUsage.
    """
    def __init__(__self__, by_networks=None, china_queries_limit=None, clean_queries=None, ddos_protection_enabled=None, ddos_queries=None, decisions_limit=None, filter_chains_limit=None, from_=None, id=None, include_dedicated_dns_network_in_managed_dns_usage=None, metric_type=None, monitors_limit=None, nxd_protection_enabled=None, nxd_responses=None, queries_limit=None, records_limit=None, to=None, total_usage=None):
        if by_networks and not isinstance(by_networks, list):
            raise TypeError("Expected argument 'by_networks' to be a list")
        pulumi.set(__self__, "by_networks", by_networks)
        if china_queries_limit and not isinstance(china_queries_limit, int):
            raise TypeError("Expected argument 'china_queries_limit' to be a int")
        pulumi.set(__self__, "china_queries_limit", china_queries_limit)
        if clean_queries and not isinstance(clean_queries, int):
            raise TypeError("Expected argument 'clean_queries' to be a int")
        pulumi.set(__self__, "clean_queries", clean_queries)
        if ddos_protection_enabled and not isinstance(ddos_protection_enabled, bool):
            raise TypeError("Expected argument 'ddos_protection_enabled' to be a bool")
        pulumi.set(__self__, "ddos_protection_enabled", ddos_protection_enabled)
        if ddos_queries and not isinstance(ddos_queries, int):
            raise TypeError("Expected argument 'ddos_queries' to be a int")
        pulumi.set(__self__, "ddos_queries", ddos_queries)
        if decisions_limit and not isinstance(decisions_limit, int):
            raise TypeError("Expected argument 'decisions_limit' to be a int")
        pulumi.set(__self__, "decisions_limit", decisions_limit)
        if filter_chains_limit and not isinstance(filter_chains_limit, int):
            raise TypeError("Expected argument 'filter_chains_limit' to be a int")
        pulumi.set(__self__, "filter_chains_limit", filter_chains_limit)
        if from_ and not isinstance(from_, int):
            raise TypeError("Expected argument 'from_' to be a int")
        pulumi.set(__self__, "from_", from_)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if include_dedicated_dns_network_in_managed_dns_usage and not isinstance(include_dedicated_dns_network_in_managed_dns_usage, bool):
            raise TypeError("Expected argument 'include_dedicated_dns_network_in_managed_dns_usage' to be a bool")
        pulumi.set(__self__, "include_dedicated_dns_network_in_managed_dns_usage", include_dedicated_dns_network_in_managed_dns_usage)
        if metric_type and not isinstance(metric_type, str):
            raise TypeError("Expected argument 'metric_type' to be a str")
        pulumi.set(__self__, "metric_type", metric_type)
        if monitors_limit and not isinstance(monitors_limit, int):
            raise TypeError("Expected argument 'monitors_limit' to be a int")
        pulumi.set(__self__, "monitors_limit", monitors_limit)
        if nxd_protection_enabled and not isinstance(nxd_protection_enabled, bool):
            raise TypeError("Expected argument 'nxd_protection_enabled' to be a bool")
        pulumi.set(__self__, "nxd_protection_enabled", nxd_protection_enabled)
        if nxd_responses and not isinstance(nxd_responses, int):
            raise TypeError("Expected argument 'nxd_responses' to be a int")
        pulumi.set(__self__, "nxd_responses", nxd_responses)
        if queries_limit and not isinstance(queries_limit, int):
            raise TypeError("Expected argument 'queries_limit' to be a int")
        pulumi.set(__self__, "queries_limit", queries_limit)
        if records_limit and not isinstance(records_limit, int):
            raise TypeError("Expected argument 'records_limit' to be a int")
        pulumi.set(__self__, "records_limit", records_limit)
        if to and not isinstance(to, int):
            raise TypeError("Expected argument 'to' to be a int")
        pulumi.set(__self__, "to", to)
        if total_usage and not isinstance(total_usage, int):
            raise TypeError("Expected argument 'total_usage' to be a int")
        pulumi.set(__self__, "total_usage", total_usage)

    @property
    @pulumi.getter(name="byNetworks")
    def by_networks(self) -> Sequence['outputs.GetBillingUsageByNetworkResult']:
        """
        (Computed) A list of network-specific query data containing:
        """
        return pulumi.get(self, "by_networks")

    @property
    @pulumi.getter(name="chinaQueriesLimit")
    def china_queries_limit(self) -> int:
        """
        (Computed) The queries limit for the China network.
        """
        return pulumi.get(self, "china_queries_limit")

    @property
    @pulumi.getter(name="cleanQueries")
    def clean_queries(self) -> int:
        """
        Clean queries for this day.
        """
        return pulumi.get(self, "clean_queries")

    @property
    @pulumi.getter(name="ddosProtectionEnabled")
    def ddos_protection_enabled(self) -> bool:
        """
        (Computed) Whether DDoS Protection is enabled.
        """
        return pulumi.get(self, "ddos_protection_enabled")

    @property
    @pulumi.getter(name="ddosQueries")
    def ddos_queries(self) -> int:
        """
        DDoS queries for this day.
        """
        return pulumi.get(self, "ddos_queries")

    @property
    @pulumi.getter(name="decisionsLimit")
    def decisions_limit(self) -> int:
        """
        (Computed) The RUM decisions limit for this billing cycle.
        """
        return pulumi.get(self, "decisions_limit")

    @property
    @pulumi.getter(name="filterChainsLimit")
    def filter_chains_limit(self) -> int:
        """
        (Computed) The filter chains limit for this billing cycle.
        """
        return pulumi.get(self, "filter_chains_limit")

    @property
    @pulumi.getter(name="from")
    def from_(self) -> Optional[int]:
        return pulumi.get(self, "from_")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="includeDedicatedDnsNetworkInManagedDnsUsage")
    def include_dedicated_dns_network_in_managed_dns_usage(self) -> bool:
        """
        (Computed) Whether dedicated DNS usage counts towards managed DNS usage.
        """
        return pulumi.get(self, "include_dedicated_dns_network_in_managed_dns_usage")

    @property
    @pulumi.getter(name="metricType")
    def metric_type(self) -> str:
        return pulumi.get(self, "metric_type")

    @property
    @pulumi.getter(name="monitorsLimit")
    def monitors_limit(self) -> int:
        """
        (Computed) The monitoring jobs limit for this billing cycle.
        """
        return pulumi.get(self, "monitors_limit")

    @property
    @pulumi.getter(name="nxdProtectionEnabled")
    def nxd_protection_enabled(self) -> bool:
        """
        (Computed) Whether NXD Protection is enabled.
        """
        return pulumi.get(self, "nxd_protection_enabled")

    @property
    @pulumi.getter(name="nxdResponses")
    def nxd_responses(self) -> int:
        """
        NXD responses for this day.
        """
        return pulumi.get(self, "nxd_responses")

    @property
    @pulumi.getter(name="queriesLimit")
    def queries_limit(self) -> int:
        """
        (Computed) The queries limit for this billing cycle.
        """
        return pulumi.get(self, "queries_limit")

    @property
    @pulumi.getter(name="recordsLimit")
    def records_limit(self) -> int:
        """
        (Computed) The records limit for this billing cycle.
        """
        return pulumi.get(self, "records_limit")

    @property
    @pulumi.getter
    def to(self) -> Optional[int]:
        return pulumi.get(self, "to")

    @property
    @pulumi.getter(name="totalUsage")
    def total_usage(self) -> int:
        """
        (Computed) The total usage count for the metric. Available for `decisions`, `filter-chains`, `monitors`, and `records` metrics.
        """
        return pulumi.get(self, "total_usage")


class AwaitableGetBillingUsageResult(GetBillingUsageResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBillingUsageResult(
            by_networks=self.by_networks,
            china_queries_limit=self.china_queries_limit,
            clean_queries=self.clean_queries,
            ddos_protection_enabled=self.ddos_protection_enabled,
            ddos_queries=self.ddos_queries,
            decisions_limit=self.decisions_limit,
            filter_chains_limit=self.filter_chains_limit,
            from_=self.from_,
            id=self.id,
            include_dedicated_dns_network_in_managed_dns_usage=self.include_dedicated_dns_network_in_managed_dns_usage,
            metric_type=self.metric_type,
            monitors_limit=self.monitors_limit,
            nxd_protection_enabled=self.nxd_protection_enabled,
            nxd_responses=self.nxd_responses,
            queries_limit=self.queries_limit,
            records_limit=self.records_limit,
            to=self.to,
            total_usage=self.total_usage)


def get_billing_usage(from_: Optional[int] = None,
                      metric_type: Optional[str] = None,
                      to: Optional[int] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBillingUsageResult:
    """
    Provides billing usage details about a NS1 account.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ns1 as ns1

    # Get query usage data for the given timeframe
    queries = ns1.get_billing_usage(metric_type="queries",
        from_=1738368000,
        to=1740787199)
    # Get account limits data for the given timeframe
    limits = ns1.get_billing_usage(metric_type="limits",
        from_=1738368000,
        to=1740787199)
    # Get RUM decisions usage data for the given timeframe
    decisions = ns1.get_billing_usage(metric_type="decisions",
        from_=1738368000,
        to=1740787199)
    # Get filter chains usage data
    filter_chains = ns1.get_billing_usage(metric_type="filter-chains")
    # Get monitoring jobs usage data
    monitors = ns1.get_billing_usage(metric_type="monitors")
    # Get records usage data
    records = ns1.get_billing_usage(metric_type="records")
    pulumi.export("totalQueries", queries.clean_queries)
    pulumi.export("totalDdosQueries", queries.ddos_queries)
    pulumi.export("totalNxdResponses", queries.nxd_responses)
    pulumi.export("queriesLimit", limits.queries_limit)
    pulumi.export("totalDecisions", decisions.total_usage)
    pulumi.export("decisionsLimit", limits.decisions_limit)
    pulumi.export("totalFilterChains", filter_chains.total_usage)
    pulumi.export("filterChainsLimit", limits.filter_chains_limit)
    pulumi.export("totalMonitors", monitors.total_usage)
    pulumi.export("monitorsLimit", limits.monitors_limit)
    pulumi.export("totalRecords", records.total_usage)
    pulumi.export("recordsLimit", limits.records_limit)
    ```


    :param int from_: The start timestamp for the data range in Unix epoch format.
    :param str metric_type: The type of billing metric to retrieve. Must be one of: `queries`, `limits`, `decisions`, `filter-chains`, `monitors`, `records`.
    :param int to: The end timestamp for the data range in Unix epoch format.
    """
    __args__ = dict()
    __args__['from'] = from_
    __args__['metricType'] = metric_type
    __args__['to'] = to
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ns1:index/getBillingUsage:getBillingUsage', __args__, opts=opts, typ=GetBillingUsageResult).value

    return AwaitableGetBillingUsageResult(
        by_networks=pulumi.get(__ret__, 'by_networks'),
        china_queries_limit=pulumi.get(__ret__, 'china_queries_limit'),
        clean_queries=pulumi.get(__ret__, 'clean_queries'),
        ddos_protection_enabled=pulumi.get(__ret__, 'ddos_protection_enabled'),
        ddos_queries=pulumi.get(__ret__, 'ddos_queries'),
        decisions_limit=pulumi.get(__ret__, 'decisions_limit'),
        filter_chains_limit=pulumi.get(__ret__, 'filter_chains_limit'),
        from_=pulumi.get(__ret__, 'from_'),
        id=pulumi.get(__ret__, 'id'),
        include_dedicated_dns_network_in_managed_dns_usage=pulumi.get(__ret__, 'include_dedicated_dns_network_in_managed_dns_usage'),
        metric_type=pulumi.get(__ret__, 'metric_type'),
        monitors_limit=pulumi.get(__ret__, 'monitors_limit'),
        nxd_protection_enabled=pulumi.get(__ret__, 'nxd_protection_enabled'),
        nxd_responses=pulumi.get(__ret__, 'nxd_responses'),
        queries_limit=pulumi.get(__ret__, 'queries_limit'),
        records_limit=pulumi.get(__ret__, 'records_limit'),
        to=pulumi.get(__ret__, 'to'),
        total_usage=pulumi.get(__ret__, 'total_usage'))
def get_billing_usage_output(from_: Optional[pulumi.Input[Optional[int]]] = None,
                             metric_type: Optional[pulumi.Input[str]] = None,
                             to: Optional[pulumi.Input[Optional[int]]] = None,
                             opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetBillingUsageResult]:
    """
    Provides billing usage details about a NS1 account.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ns1 as ns1

    # Get query usage data for the given timeframe
    queries = ns1.get_billing_usage(metric_type="queries",
        from_=1738368000,
        to=1740787199)
    # Get account limits data for the given timeframe
    limits = ns1.get_billing_usage(metric_type="limits",
        from_=1738368000,
        to=1740787199)
    # Get RUM decisions usage data for the given timeframe
    decisions = ns1.get_billing_usage(metric_type="decisions",
        from_=1738368000,
        to=1740787199)
    # Get filter chains usage data
    filter_chains = ns1.get_billing_usage(metric_type="filter-chains")
    # Get monitoring jobs usage data
    monitors = ns1.get_billing_usage(metric_type="monitors")
    # Get records usage data
    records = ns1.get_billing_usage(metric_type="records")
    pulumi.export("totalQueries", queries.clean_queries)
    pulumi.export("totalDdosQueries", queries.ddos_queries)
    pulumi.export("totalNxdResponses", queries.nxd_responses)
    pulumi.export("queriesLimit", limits.queries_limit)
    pulumi.export("totalDecisions", decisions.total_usage)
    pulumi.export("decisionsLimit", limits.decisions_limit)
    pulumi.export("totalFilterChains", filter_chains.total_usage)
    pulumi.export("filterChainsLimit", limits.filter_chains_limit)
    pulumi.export("totalMonitors", monitors.total_usage)
    pulumi.export("monitorsLimit", limits.monitors_limit)
    pulumi.export("totalRecords", records.total_usage)
    pulumi.export("recordsLimit", limits.records_limit)
    ```


    :param int from_: The start timestamp for the data range in Unix epoch format.
    :param str metric_type: The type of billing metric to retrieve. Must be one of: `queries`, `limits`, `decisions`, `filter-chains`, `monitors`, `records`.
    :param int to: The end timestamp for the data range in Unix epoch format.
    """
    __args__ = dict()
    __args__['from'] = from_
    __args__['metricType'] = metric_type
    __args__['to'] = to
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ns1:index/getBillingUsage:getBillingUsage', __args__, opts=opts, typ=GetBillingUsageResult)
    return __ret__.apply(lambda __response__: GetBillingUsageResult(
        by_networks=pulumi.get(__response__, 'by_networks'),
        china_queries_limit=pulumi.get(__response__, 'china_queries_limit'),
        clean_queries=pulumi.get(__response__, 'clean_queries'),
        ddos_protection_enabled=pulumi.get(__response__, 'ddos_protection_enabled'),
        ddos_queries=pulumi.get(__response__, 'ddos_queries'),
        decisions_limit=pulumi.get(__response__, 'decisions_limit'),
        filter_chains_limit=pulumi.get(__response__, 'filter_chains_limit'),
        from_=pulumi.get(__response__, 'from_'),
        id=pulumi.get(__response__, 'id'),
        include_dedicated_dns_network_in_managed_dns_usage=pulumi.get(__response__, 'include_dedicated_dns_network_in_managed_dns_usage'),
        metric_type=pulumi.get(__response__, 'metric_type'),
        monitors_limit=pulumi.get(__response__, 'monitors_limit'),
        nxd_protection_enabled=pulumi.get(__response__, 'nxd_protection_enabled'),
        nxd_responses=pulumi.get(__response__, 'nxd_responses'),
        queries_limit=pulumi.get(__response__, 'queries_limit'),
        records_limit=pulumi.get(__response__, 'records_limit'),
        to=pulumi.get(__response__, 'to'),
        total_usage=pulumi.get(__response__, 'total_usage')))
