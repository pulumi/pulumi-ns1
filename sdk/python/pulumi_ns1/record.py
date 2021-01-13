# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Record']


class Record(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 answers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RecordAnswerArgs']]]]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RecordFilterArgs']]]]] = None,
                 link: Optional[pulumi.Input[str]] = None,
                 meta: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 regions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RecordRegionArgs']]]]] = None,
                 short_answers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 use_client_subnet: Optional[pulumi.Input[bool]] = None,
                 zone: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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
               * ` meta` - (Optional) meta is supported at the `record` level. Meta
               is documented below.
        :param pulumi.Input[str] zone: The zone the record belongs to. Cannot have leading or
               trailing dots (".") - see the example above and `FQDN formatting` below.
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

            __props__['answers'] = answers
            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__['domain'] = domain
            __props__['filters'] = filters
            __props__['link'] = link
            __props__['meta'] = meta
            __props__['regions'] = regions
            if short_answers is not None and not opts.urn:
                warnings.warn("""short_answers will be deprecated in a future release. It is suggested to migrate to a regular \"answers\" block.""", DeprecationWarning)
                pulumi.log.warn("short_answers is deprecated: short_answers will be deprecated in a future release. It is suggested to migrate to a regular \"answers\" block.")
            __props__['short_answers'] = short_answers
            __props__['ttl'] = ttl
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
            __props__['use_client_subnet'] = use_client_subnet
            if zone is None and not opts.urn:
                raise TypeError("Missing required property 'zone'")
            __props__['zone'] = zone
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
            domain: Optional[pulumi.Input[str]] = None,
            filters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RecordFilterArgs']]]]] = None,
            link: Optional[pulumi.Input[str]] = None,
            meta: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            regions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RecordRegionArgs']]]]] = None,
            short_answers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
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
               * ` meta` - (Optional) meta is supported at the `record` level. Meta
               is documented below.
        :param pulumi.Input[str] zone: The zone the record belongs to. Cannot have leading or
               trailing dots (".") - see the example above and `FQDN formatting` below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["answers"] = answers
        __props__["domain"] = domain
        __props__["filters"] = filters
        __props__["link"] = link
        __props__["meta"] = meta
        __props__["regions"] = regions
        __props__["short_answers"] = short_answers
        __props__["ttl"] = ttl
        __props__["type"] = type
        __props__["use_client_subnet"] = use_client_subnet
        __props__["zone"] = zone
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
        return pulumi.get(self, "short_answers")

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
        * ` meta` - (Optional) meta is supported at the `record` level. Meta
        is documented below.
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

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

