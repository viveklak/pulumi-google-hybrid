# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from ._inputs import *

__all__ = ['ServiceServiceLevelObjective']


class ServiceServiceLevelObjective(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 calendar_period: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 goal: Optional[pulumi.Input[float]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rolling_period: Optional[pulumi.Input[str]] = None,
                 service_level_indicator: Optional[pulumi.Input[pulumi.InputType['ServiceLevelIndicatorArgs']]] = None,
                 service_level_objectives_id: Optional[pulumi.Input[str]] = None,
                 services_id: Optional[pulumi.Input[str]] = None,
                 v3_id: Optional[pulumi.Input[str]] = None,
                 v3_id1: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a ServiceLevelObjective for the given Service.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] calendar_period: A calendar period, semantically "since the start of the current ". At this time, only DAY, WEEK, FORTNIGHT, and MONTH are supported.
        :param pulumi.Input[str] display_name: Name used for UI elements listing this SLO.
        :param pulumi.Input[float] goal: The fraction of service that must be good in order for this objective to be met. 0 < goal <= 0.999.
        :param pulumi.Input[str] name: Resource name for this ServiceLevelObjective. The format is: projects/[PROJECT_ID_OR_NUMBER]/services/[SERVICE_ID]/serviceLevelObjectives/[SLO_NAME] 
        :param pulumi.Input[str] rolling_period: A rolling time period, semantically "in the past ". Must be an integer multiple of 1 day no larger than 30 days.
        :param pulumi.Input[pulumi.InputType['ServiceLevelIndicatorArgs']] service_level_indicator: The definition of good service, used to measure and calculate the quality of the Service's performance with respect to a single aspect of service quality.
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

            __props__['calendar_period'] = calendar_period
            __props__['display_name'] = display_name
            __props__['goal'] = goal
            __props__['name'] = name
            __props__['rolling_period'] = rolling_period
            __props__['service_level_indicator'] = service_level_indicator
            if service_level_objectives_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_level_objectives_id'")
            __props__['service_level_objectives_id'] = service_level_objectives_id
            if services_id is None and not opts.urn:
                raise TypeError("Missing required property 'services_id'")
            __props__['services_id'] = services_id
            if v3_id is None and not opts.urn:
                raise TypeError("Missing required property 'v3_id'")
            __props__['v3_id'] = v3_id
            if v3_id1 is None and not opts.urn:
                raise TypeError("Missing required property 'v3_id1'")
            __props__['v3_id1'] = v3_id1
        super(ServiceServiceLevelObjective, __self__).__init__(
            'google-cloud:monitoring/v3:ServiceServiceLevelObjective',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ServiceServiceLevelObjective':
        """
        Get an existing ServiceServiceLevelObjective resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ServiceServiceLevelObjective(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

