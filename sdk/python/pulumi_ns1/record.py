# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['RecordArgs', 'Record']

@pulumi.input_type
class RecordArgs:
    def __init__(__self__, *,
                 domain: pulumi.Input[str],
                 type: pulumi.Input[str],
                 zone: pulumi.Input[str],
                 answers: Optional[pulumi.Input[Sequence[pulumi.Input['RecordAnswerArgs']]]] = None,
                 blocked_tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 filters: Optional[pulumi.Input[Sequence[pulumi.Input['RecordFilterArgs']]]] = None,
                 link: Optional[pulumi.Input[str]] = None,
                 meta: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 override_ttl: Optional[pulumi.Input[bool]] = None,
                 regions: Optional[pulumi.Input[Sequence[pulumi.Input['RecordRegionArgs']]]] = None,
                 short_answers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 use_client_subnet: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Record resource.
        :param pulumi.Input[str] domain: The records' domain. Cannot have leading or trailing
               dots - see the example above and `FQDN formatting` below.
        :param pulumi.Input[str] type: The records' RR type.
        :param pulumi.Input[str] zone: The zone the record belongs to. Cannot have leading or
               trailing dots (".") - see the example above and `FQDN formatting` below.
        :param pulumi.Input[Sequence[pulumi.Input['RecordAnswerArgs']]] answers: One or more NS1 answers for the records' specified type.
               Answers are documented below.
        :param pulumi.Input[Sequence[pulumi.Input['RecordFilterArgs']]] filters: One or more NS1 filters for the record(order matters).
               Filters are documented below.
        :param pulumi.Input[str] link: The target record to link to. This means this record is a
               'linked' record, and it inherits all properties from its target.
        :param pulumi.Input[Sequence[pulumi.Input['RecordRegionArgs']]] regions: One or more "regions" for the record. These are really
               just groupings based on metadata, and are called "Answer Groups" in the NS1 UI,
               but remain `regions` here for legacy reasons. Regions are
               documented below. Please note the ordering requirement!
        :param pulumi.Input[int] ttl: The records' time to live (in seconds).
        :param pulumi.Input[bool] use_client_subnet: Whether to use EDNS client subnet data when
               available(in filter chain).
        """
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "zone", zone)
        if answers is not None:
            pulumi.set(__self__, "answers", answers)
        if blocked_tags is not None:
            pulumi.set(__self__, "blocked_tags", blocked_tags)
        if filters is not None:
            pulumi.set(__self__, "filters", filters)
        if link is not None:
            pulumi.set(__self__, "link", link)
        if meta is not None:
            pulumi.set(__self__, "meta", meta)
        if override_ttl is not None:
            pulumi.set(__self__, "override_ttl", override_ttl)
        if regions is not None:
            pulumi.set(__self__, "regions", regions)
        if short_answers is not None:
            warnings.warn("""short_answers will be deprecated in a future release. It is suggested to migrate to a regular \"answers\" block.""", DeprecationWarning)
            pulumi.log.warn("""short_answers is deprecated: short_answers will be deprecated in a future release. It is suggested to migrate to a regular \"answers\" block.""")
        if short_answers is not None:
            pulumi.set(__self__, "short_answers", short_answers)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)
        if use_client_subnet is not None:
            pulumi.set(__self__, "use_client_subnet", use_client_subnet)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        """
        The records' domain. Cannot have leading or trailing
        dots - see the example above and `FQDN formatting` below.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The records' RR type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Input[str]:
        """
        The zone the record belongs to. Cannot have leading or
        trailing dots (".") - see the example above and `FQDN formatting` below.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone", value)

    @property
    @pulumi.getter
    def answers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RecordAnswerArgs']]]]:
        """
        One or more NS1 answers for the records' specified type.
        Answers are documented below.
        """
        return pulumi.get(self, "answers")

    @answers.setter
    def answers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RecordAnswerArgs']]]]):
        pulumi.set(self, "answers", value)

    @property
    @pulumi.getter(name="blockedTags")
    def blocked_tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "blocked_tags")

    @blocked_tags.setter
    def blocked_tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "blocked_tags", value)

    @property
    @pulumi.getter
    def filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RecordFilterArgs']]]]:
        """
        One or more NS1 filters for the record(order matters).
        Filters are documented below.
        """
        return pulumi.get(self, "filters")

    @filters.setter
    def filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RecordFilterArgs']]]]):
        pulumi.set(self, "filters", value)

    @property
    @pulumi.getter
    def link(self) -> Optional[pulumi.Input[str]]:
        """
        The target record to link to. This means this record is a
        'linked' record, and it inherits all properties from its target.
        """
        return pulumi.get(self, "link")

    @link.setter
    def link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "link", value)

    @property
    @pulumi.getter
    def meta(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "meta")

    @meta.setter
    def meta(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "meta", value)

    @property
    @pulumi.getter(name="overrideTtl")
    def override_ttl(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "override_ttl")

    @override_ttl.setter
    def override_ttl(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "override_ttl", value)

    @property
    @pulumi.getter
    def regions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RecordRegionArgs']]]]:
        """
        One or more "regions" for the record. These are really
        just groupings based on metadata, and are called "Answer Groups" in the NS1 UI,
        but remain `regions` here for legacy reasons. Regions are
        documented below. Please note the ordering requirement!
        """
        return pulumi.get(self, "regions")

    @regions.setter
    def regions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RecordRegionArgs']]]]):
        pulumi.set(self, "regions", value)

    @property
    @pulumi.getter(name="shortAnswers")
    def short_answers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        warnings.warn("""short_answers will be deprecated in a future release. It is suggested to migrate to a regular \"answers\" block.""", DeprecationWarning)
        pulumi.log.warn("""short_answers is deprecated: short_answers will be deprecated in a future release. It is suggested to migrate to a regular \"answers\" block.""")

        return pulumi.get(self, "short_answers")

    @short_answers.setter
    def short_answers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "short_answers", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[int]]:
        """
        The records' time to live (in seconds).
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ttl", value)

    @property
    @pulumi.getter(name="useClientSubnet")
    def use_client_subnet(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to use EDNS client subnet data when
        available(in filter chain).
        """
        return pulumi.get(self, "use_client_subnet")

    @use_client_subnet.setter
    def use_client_subnet(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_client_subnet", value)


@pulumi.input_type
class _RecordState:
    def __init__(__self__, *,
                 answers: Optional[pulumi.Input[Sequence[pulumi.Input['RecordAnswerArgs']]]] = None,
                 blocked_tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 filters: Optional[pulumi.Input[Sequence[pulumi.Input['RecordFilterArgs']]]] = None,
                 link: Optional[pulumi.Input[str]] = None,
                 meta: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 override_ttl: Optional[pulumi.Input[bool]] = None,
                 regions: Optional[pulumi.Input[Sequence[pulumi.Input['RecordRegionArgs']]]] = None,
                 short_answers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 use_client_subnet: Optional[pulumi.Input[bool]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Record resources.
        :param pulumi.Input[Sequence[pulumi.Input['RecordAnswerArgs']]] answers: One or more NS1 answers for the records' specified type.
               Answers are documented below.
        :param pulumi.Input[str] domain: The records' domain. Cannot have leading or trailing
               dots - see the example above and `FQDN formatting` below.
        :param pulumi.Input[Sequence[pulumi.Input['RecordFilterArgs']]] filters: One or more NS1 filters for the record(order matters).
               Filters are documented below.
        :param pulumi.Input[str] link: The target record to link to. This means this record is a
               'linked' record, and it inherits all properties from its target.
        :param pulumi.Input[Sequence[pulumi.Input['RecordRegionArgs']]] regions: One or more "regions" for the record. These are really
               just groupings based on metadata, and are called "Answer Groups" in the NS1 UI,
               but remain `regions` here for legacy reasons. Regions are
               documented below. Please note the ordering requirement!
        :param pulumi.Input[int] ttl: The records' time to live (in seconds).
        :param pulumi.Input[str] type: The records' RR type.
        :param pulumi.Input[bool] use_client_subnet: Whether to use EDNS client subnet data when
               available(in filter chain).
        :param pulumi.Input[str] zone: The zone the record belongs to. Cannot have leading or
               trailing dots (".") - see the example above and `FQDN formatting` below.
        """
        if answers is not None:
            pulumi.set(__self__, "answers", answers)
        if blocked_tags is not None:
            pulumi.set(__self__, "blocked_tags", blocked_tags)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if filters is not None:
            pulumi.set(__self__, "filters", filters)
        if link is not None:
            pulumi.set(__self__, "link", link)
        if meta is not None:
            pulumi.set(__self__, "meta", meta)
        if override_ttl is not None:
            pulumi.set(__self__, "override_ttl", override_ttl)
        if regions is not None:
            pulumi.set(__self__, "regions", regions)
        if short_answers is not None:
            warnings.warn("""short_answers will be deprecated in a future release. It is suggested to migrate to a regular \"answers\" block.""", DeprecationWarning)
            pulumi.log.warn("""short_answers is deprecated: short_answers will be deprecated in a future release. It is suggested to migrate to a regular \"answers\" block.""")
        if short_answers is not None:
            pulumi.set(__self__, "short_answers", short_answers)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if use_client_subnet is not None:
            pulumi.set(__self__, "use_client_subnet", use_client_subnet)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def answers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RecordAnswerArgs']]]]:
        """
        One or more NS1 answers for the records' specified type.
        Answers are documented below.
        """
        return pulumi.get(self, "answers")

    @answers.setter
    def answers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RecordAnswerArgs']]]]):
        pulumi.set(self, "answers", value)

    @property
    @pulumi.getter(name="blockedTags")
    def blocked_tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "blocked_tags")

    @blocked_tags.setter
    def blocked_tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "blocked_tags", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[str]]:
        """
        The records' domain. Cannot have leading or trailing
        dots - see the example above and `FQDN formatting` below.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter
    def filters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RecordFilterArgs']]]]:
        """
        One or more NS1 filters for the record(order matters).
        Filters are documented below.
        """
        return pulumi.get(self, "filters")

    @filters.setter
    def filters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RecordFilterArgs']]]]):
        pulumi.set(self, "filters", value)

    @property
    @pulumi.getter
    def link(self) -> Optional[pulumi.Input[str]]:
        """
        The target record to link to. This means this record is a
        'linked' record, and it inherits all properties from its target.
        """
        return pulumi.get(self, "link")

    @link.setter
    def link(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "link", value)

    @property
    @pulumi.getter
    def meta(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "meta")

    @meta.setter
    def meta(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "meta", value)

    @property
    @pulumi.getter(name="overrideTtl")
    def override_ttl(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "override_ttl")

    @override_ttl.setter
    def override_ttl(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "override_ttl", value)

    @property
    @pulumi.getter
    def regions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RecordRegionArgs']]]]:
        """
        One or more "regions" for the record. These are really
        just groupings based on metadata, and are called "Answer Groups" in the NS1 UI,
        but remain `regions` here for legacy reasons. Regions are
        documented below. Please note the ordering requirement!
        """
        return pulumi.get(self, "regions")

    @regions.setter
    def regions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RecordRegionArgs']]]]):
        pulumi.set(self, "regions", value)

    @property
    @pulumi.getter(name="shortAnswers")
    def short_answers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        warnings.warn("""short_answers will be deprecated in a future release. It is suggested to migrate to a regular \"answers\" block.""", DeprecationWarning)
        pulumi.log.warn("""short_answers is deprecated: short_answers will be deprecated in a future release. It is suggested to migrate to a regular \"answers\" block.""")

        return pulumi.get(self, "short_answers")

    @short_answers.setter
    def short_answers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "short_answers", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[int]]:
        """
        The records' time to live (in seconds).
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ttl", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The records' RR type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="useClientSubnet")
    def use_client_subnet(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to use EDNS client subnet data when
        available(in filter chain).
        """
        return pulumi.get(self, "use_client_subnet")

    @use_client_subnet.setter
    def use_client_subnet(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_client_subnet", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        The zone the record belongs to. Cannot have leading or
        trailing dots (".") - see the example above and `FQDN formatting` below.
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class Record(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 answers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RecordAnswerArgs']]]]] = None,
                 blocked_tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RecordFilterArgs']]]]] = None,
                 link: Optional[pulumi.Input[str]] = None,
                 meta: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 override_ttl: Optional[pulumi.Input[bool]] = None,
                 regions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RecordRegionArgs']]]]] = None,
                 short_answers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 use_client_subnet: Optional[pulumi.Input[bool]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a NS1 Record resource. This can be used to create, modify, and delete records.

        ## NS1 Documentation

        [Record Api Doc](https://ns1.com/api#records)

        ## Import

        ```sh
         $ pulumi import ns1:index/record:Record <name> <zone>/<domain>/<type>`
        ```

         So for the example above

        ```sh
         $ pulumi import ns1:index/record:Record www terraform.example.io/www.terraform.example.io/CNAME`
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RecordAnswerArgs']]]] answers: One or more NS1 answers for the records' specified type.
               Answers are documented below.
        :param pulumi.Input[str] domain: The records' domain. Cannot have leading or trailing
               dots - see the example above and `FQDN formatting` below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RecordFilterArgs']]]] filters: One or more NS1 filters for the record(order matters).
               Filters are documented below.
        :param pulumi.Input[str] link: The target record to link to. This means this record is a
               'linked' record, and it inherits all properties from its target.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RecordRegionArgs']]]] regions: One or more "regions" for the record. These are really
               just groupings based on metadata, and are called "Answer Groups" in the NS1 UI,
               but remain `regions` here for legacy reasons. Regions are
               documented below. Please note the ordering requirement!
        :param pulumi.Input[int] ttl: The records' time to live (in seconds).
        :param pulumi.Input[str] type: The records' RR type.
        :param pulumi.Input[bool] use_client_subnet: Whether to use EDNS client subnet data when
               available(in filter chain).
        :param pulumi.Input[str] zone: The zone the record belongs to. Cannot have leading or
               trailing dots (".") - see the example above and `FQDN formatting` below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RecordArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a NS1 Record resource. This can be used to create, modify, and delete records.

        ## NS1 Documentation

        [Record Api Doc](https://ns1.com/api#records)

        ## Import

        ```sh
         $ pulumi import ns1:index/record:Record <name> <zone>/<domain>/<type>`
        ```

         So for the example above

        ```sh
         $ pulumi import ns1:index/record:Record www terraform.example.io/www.terraform.example.io/CNAME`
        ```

        :param str resource_name: The name of the resource.
        :param RecordArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RecordArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 answers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RecordAnswerArgs']]]]] = None,
                 blocked_tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RecordFilterArgs']]]]] = None,
                 link: Optional[pulumi.Input[str]] = None,
                 meta: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 override_ttl: Optional[pulumi.Input[bool]] = None,
                 regions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RecordRegionArgs']]]]] = None,
                 short_answers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 use_client_subnet: Optional[pulumi.Input[bool]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RecordArgs.__new__(RecordArgs)

            __props__.__dict__["answers"] = answers
            __props__.__dict__["blocked_tags"] = blocked_tags
            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__.__dict__["domain"] = domain
            __props__.__dict__["filters"] = filters
            __props__.__dict__["link"] = link
            __props__.__dict__["meta"] = meta
            __props__.__dict__["override_ttl"] = override_ttl
            __props__.__dict__["regions"] = regions
            __props__.__dict__["short_answers"] = short_answers
            __props__.__dict__["tags"] = tags
            __props__.__dict__["ttl"] = ttl
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["use_client_subnet"] = use_client_subnet
            if zone is None and not opts.urn:
                raise TypeError("Missing required property 'zone'")
            __props__.__dict__["zone"] = zone
        super(Record, __self__).__init__(
            'ns1:index/record:Record',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            answers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RecordAnswerArgs']]]]] = None,
            blocked_tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            domain: Optional[pulumi.Input[str]] = None,
            filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RecordFilterArgs']]]]] = None,
            link: Optional[pulumi.Input[str]] = None,
            meta: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            override_ttl: Optional[pulumi.Input[bool]] = None,
            regions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RecordRegionArgs']]]]] = None,
            short_answers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            ttl: Optional[pulumi.Input[int]] = None,
            type: Optional[pulumi.Input[str]] = None,
            use_client_subnet: Optional[pulumi.Input[bool]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'Record':
        """
        Get an existing Record resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RecordAnswerArgs']]]] answers: One or more NS1 answers for the records' specified type.
               Answers are documented below.
        :param pulumi.Input[str] domain: The records' domain. Cannot have leading or trailing
               dots - see the example above and `FQDN formatting` below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RecordFilterArgs']]]] filters: One or more NS1 filters for the record(order matters).
               Filters are documented below.
        :param pulumi.Input[str] link: The target record to link to. This means this record is a
               'linked' record, and it inherits all properties from its target.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RecordRegionArgs']]]] regions: One or more "regions" for the record. These are really
               just groupings based on metadata, and are called "Answer Groups" in the NS1 UI,
               but remain `regions` here for legacy reasons. Regions are
               documented below. Please note the ordering requirement!
        :param pulumi.Input[int] ttl: The records' time to live (in seconds).
        :param pulumi.Input[str] type: The records' RR type.
        :param pulumi.Input[bool] use_client_subnet: Whether to use EDNS client subnet data when
               available(in filter chain).
        :param pulumi.Input[str] zone: The zone the record belongs to. Cannot have leading or
               trailing dots (".") - see the example above and `FQDN formatting` below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RecordState.__new__(_RecordState)

        __props__.__dict__["answers"] = answers
        __props__.__dict__["blocked_tags"] = blocked_tags
        __props__.__dict__["domain"] = domain
        __props__.__dict__["filters"] = filters
        __props__.__dict__["link"] = link
        __props__.__dict__["meta"] = meta
        __props__.__dict__["override_ttl"] = override_ttl
        __props__.__dict__["regions"] = regions
        __props__.__dict__["short_answers"] = short_answers
        __props__.__dict__["tags"] = tags
        __props__.__dict__["ttl"] = ttl
        __props__.__dict__["type"] = type
        __props__.__dict__["use_client_subnet"] = use_client_subnet
        __props__.__dict__["zone"] = zone
        return Record(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def answers(self) -> pulumi.Output[Optional[Sequence['outputs.RecordAnswer']]]:
        """
        One or more NS1 answers for the records' specified type.
        Answers are documented below.
        """
        return pulumi.get(self, "answers")

    @property
    @pulumi.getter(name="blockedTags")
    def blocked_tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "blocked_tags")

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[str]:
        """
        The records' domain. Cannot have leading or trailing
        dots - see the example above and `FQDN formatting` below.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def filters(self) -> pulumi.Output[Optional[Sequence['outputs.RecordFilter']]]:
        """
        One or more NS1 filters for the record(order matters).
        Filters are documented below.
        """
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def link(self) -> pulumi.Output[Optional[str]]:
        """
        The target record to link to. This means this record is a
        'linked' record, and it inherits all properties from its target.
        """
        return pulumi.get(self, "link")

    @property
    @pulumi.getter
    def meta(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        return pulumi.get(self, "meta")

    @property
    @pulumi.getter(name="overrideTtl")
    def override_ttl(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "override_ttl")

    @property
    @pulumi.getter
    def regions(self) -> pulumi.Output[Optional[Sequence['outputs.RecordRegion']]]:
        """
        One or more "regions" for the record. These are really
        just groupings based on metadata, and are called "Answer Groups" in the NS1 UI,
        but remain `regions` here for legacy reasons. Regions are
        documented below. Please note the ordering requirement!
        """
        return pulumi.get(self, "regions")

    @property
    @pulumi.getter(name="shortAnswers")
    def short_answers(self) -> pulumi.Output[Optional[Sequence[str]]]:
        warnings.warn("""short_answers will be deprecated in a future release. It is suggested to migrate to a regular \"answers\" block.""", DeprecationWarning)
        pulumi.log.warn("""short_answers is deprecated: short_answers will be deprecated in a future release. It is suggested to migrate to a regular \"answers\" block.""")

        return pulumi.get(self, "short_answers")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[int]:
        """
        The records' time to live (in seconds).
        """
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The records' RR type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="useClientSubnet")
    def use_client_subnet(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether to use EDNS client subnet data when
        available(in filter chain).
        """
        return pulumi.get(self, "use_client_subnet")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        The zone the record belongs to. Cannot have leading or
        trailing dots (".") - see the example above and `FQDN formatting` below.
        """
        return pulumi.get(self, "zone")

