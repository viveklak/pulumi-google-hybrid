# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from ._enums import *

__all__ = [
    'GoogleCloudChannelV1AssociationInfoArgs',
    'GoogleCloudChannelV1CommitmentSettingsArgs',
    'GoogleCloudChannelV1ContactInfoArgs',
    'GoogleCloudChannelV1ParameterArgs',
    'GoogleCloudChannelV1PeriodArgs',
    'GoogleCloudChannelV1RenewalSettingsArgs',
    'GoogleCloudChannelV1ValueArgs',
    'GoogleTypePostalAddressArgs',
]

@pulumi.input_type
class GoogleCloudChannelV1AssociationInfoArgs:
    def __init__(__self__, *,
                 base_entitlement: Optional[pulumi.Input[str]] = None):
        """
        Association links that an entitlement has to other entitlements.
        :param pulumi.Input[str] base_entitlement: The name of the base entitlement, for which this entitlement is an add-on.
        """
        if base_entitlement is not None:
            pulumi.set(__self__, "base_entitlement", base_entitlement)

    @property
    @pulumi.getter(name="baseEntitlement")
    def base_entitlement(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the base entitlement, for which this entitlement is an add-on.
        """
        return pulumi.get(self, "base_entitlement")

    @base_entitlement.setter
    def base_entitlement(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "base_entitlement", value)


@pulumi.input_type
class GoogleCloudChannelV1CommitmentSettingsArgs:
    def __init__(__self__, *,
                 renewal_settings: Optional[pulumi.Input['GoogleCloudChannelV1RenewalSettingsArgs']] = None):
        """
        Commitment settings for commitment-based offers.
        :param pulumi.Input['GoogleCloudChannelV1RenewalSettingsArgs'] renewal_settings: Optional. Renewal settings applicable for a commitment-based Offer.
        """
        if renewal_settings is not None:
            pulumi.set(__self__, "renewal_settings", renewal_settings)

    @property
    @pulumi.getter(name="renewalSettings")
    def renewal_settings(self) -> Optional[pulumi.Input['GoogleCloudChannelV1RenewalSettingsArgs']]:
        """
        Optional. Renewal settings applicable for a commitment-based Offer.
        """
        return pulumi.get(self, "renewal_settings")

    @renewal_settings.setter
    def renewal_settings(self, value: Optional[pulumi.Input['GoogleCloudChannelV1RenewalSettingsArgs']]):
        pulumi.set(self, "renewal_settings", value)


@pulumi.input_type
class GoogleCloudChannelV1ContactInfoArgs:
    def __init__(__self__, *,
                 email: Optional[pulumi.Input[str]] = None,
                 first_name: Optional[pulumi.Input[str]] = None,
                 last_name: Optional[pulumi.Input[str]] = None,
                 phone: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None):
        """
        Contact information for a customer account.
        :param pulumi.Input[str] email: The customer account's contact email. Required for entitlements that create admin.google.com accounts, and serves as the customer's username for those accounts. Use this email to invite Team customers.
        :param pulumi.Input[str] first_name: The customer account contact's first name. Optional for Team customers.
        :param pulumi.Input[str] last_name: The customer account contact's last name. Optional for Team customers.
        :param pulumi.Input[str] phone: The customer account's contact phone number.
        :param pulumi.Input[str] title: Optional. The customer account contact's job title.
        """
        if email is not None:
            pulumi.set(__self__, "email", email)
        if first_name is not None:
            pulumi.set(__self__, "first_name", first_name)
        if last_name is not None:
            pulumi.set(__self__, "last_name", last_name)
        if phone is not None:
            pulumi.set(__self__, "phone", phone)
        if title is not None:
            pulumi.set(__self__, "title", title)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        The customer account's contact email. Required for entitlements that create admin.google.com accounts, and serves as the customer's username for those accounts. Use this email to invite Team customers.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="firstName")
    def first_name(self) -> Optional[pulumi.Input[str]]:
        """
        The customer account contact's first name. Optional for Team customers.
        """
        return pulumi.get(self, "first_name")

    @first_name.setter
    def first_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "first_name", value)

    @property
    @pulumi.getter(name="lastName")
    def last_name(self) -> Optional[pulumi.Input[str]]:
        """
        The customer account contact's last name. Optional for Team customers.
        """
        return pulumi.get(self, "last_name")

    @last_name.setter
    def last_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_name", value)

    @property
    @pulumi.getter
    def phone(self) -> Optional[pulumi.Input[str]]:
        """
        The customer account's contact phone number.
        """
        return pulumi.get(self, "phone")

    @phone.setter
    def phone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "phone", value)

    @property
    @pulumi.getter
    def title(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. The customer account contact's job title.
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "title", value)


@pulumi.input_type
class GoogleCloudChannelV1ParameterArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input['GoogleCloudChannelV1ValueArgs']] = None):
        """
        Definition for extended entitlement parameters.
        :param pulumi.Input[str] name: Name of the parameter.
        :param pulumi.Input['GoogleCloudChannelV1ValueArgs'] value: Value of the parameter.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the parameter.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input['GoogleCloudChannelV1ValueArgs']]:
        """
        Value of the parameter.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input['GoogleCloudChannelV1ValueArgs']]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class GoogleCloudChannelV1PeriodArgs:
    def __init__(__self__, *,
                 duration: Optional[pulumi.Input[int]] = None,
                 period_type: Optional[pulumi.Input['GoogleCloudChannelV1PeriodPeriodType']] = None):
        """
        Represents period in days/months/years.
        :param pulumi.Input[int] duration: Total duration of Period Type defined.
        :param pulumi.Input['GoogleCloudChannelV1PeriodPeriodType'] period_type: Period Type.
        """
        if duration is not None:
            pulumi.set(__self__, "duration", duration)
        if period_type is not None:
            pulumi.set(__self__, "period_type", period_type)

    @property
    @pulumi.getter
    def duration(self) -> Optional[pulumi.Input[int]]:
        """
        Total duration of Period Type defined.
        """
        return pulumi.get(self, "duration")

    @duration.setter
    def duration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "duration", value)

    @property
    @pulumi.getter(name="periodType")
    def period_type(self) -> Optional[pulumi.Input['GoogleCloudChannelV1PeriodPeriodType']]:
        """
        Period Type.
        """
        return pulumi.get(self, "period_type")

    @period_type.setter
    def period_type(self, value: Optional[pulumi.Input['GoogleCloudChannelV1PeriodPeriodType']]):
        pulumi.set(self, "period_type", value)


@pulumi.input_type
class GoogleCloudChannelV1RenewalSettingsArgs:
    def __init__(__self__, *,
                 enable_renewal: Optional[pulumi.Input[bool]] = None,
                 payment_cycle: Optional[pulumi.Input['GoogleCloudChannelV1PeriodArgs']] = None,
                 payment_plan: Optional[pulumi.Input['GoogleCloudChannelV1RenewalSettingsPaymentPlan']] = None,
                 resize_unit_count: Optional[pulumi.Input[bool]] = None):
        """
        Renewal settings for renewable Offers.
        :param pulumi.Input[bool] enable_renewal: If false, the plan will be completed at the end date.
        :param pulumi.Input['GoogleCloudChannelV1PeriodArgs'] payment_cycle: Describes how frequently the reseller will be billed, such as once per month.
        :param pulumi.Input['GoogleCloudChannelV1RenewalSettingsPaymentPlan'] payment_plan: Describes how a reseller will be billed.
        :param pulumi.Input[bool] resize_unit_count: If true and enable_renewal = true, the unit (for example seats or licenses) will be set to the number of active units at renewal time.
        """
        if enable_renewal is not None:
            pulumi.set(__self__, "enable_renewal", enable_renewal)
        if payment_cycle is not None:
            pulumi.set(__self__, "payment_cycle", payment_cycle)
        if payment_plan is not None:
            pulumi.set(__self__, "payment_plan", payment_plan)
        if resize_unit_count is not None:
            pulumi.set(__self__, "resize_unit_count", resize_unit_count)

    @property
    @pulumi.getter(name="enableRenewal")
    def enable_renewal(self) -> Optional[pulumi.Input[bool]]:
        """
        If false, the plan will be completed at the end date.
        """
        return pulumi.get(self, "enable_renewal")

    @enable_renewal.setter
    def enable_renewal(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_renewal", value)

    @property
    @pulumi.getter(name="paymentCycle")
    def payment_cycle(self) -> Optional[pulumi.Input['GoogleCloudChannelV1PeriodArgs']]:
        """
        Describes how frequently the reseller will be billed, such as once per month.
        """
        return pulumi.get(self, "payment_cycle")

    @payment_cycle.setter
    def payment_cycle(self, value: Optional[pulumi.Input['GoogleCloudChannelV1PeriodArgs']]):
        pulumi.set(self, "payment_cycle", value)

    @property
    @pulumi.getter(name="paymentPlan")
    def payment_plan(self) -> Optional[pulumi.Input['GoogleCloudChannelV1RenewalSettingsPaymentPlan']]:
        """
        Describes how a reseller will be billed.
        """
        return pulumi.get(self, "payment_plan")

    @payment_plan.setter
    def payment_plan(self, value: Optional[pulumi.Input['GoogleCloudChannelV1RenewalSettingsPaymentPlan']]):
        pulumi.set(self, "payment_plan", value)

    @property
    @pulumi.getter(name="resizeUnitCount")
    def resize_unit_count(self) -> Optional[pulumi.Input[bool]]:
        """
        If true and enable_renewal = true, the unit (for example seats or licenses) will be set to the number of active units at renewal time.
        """
        return pulumi.get(self, "resize_unit_count")

    @resize_unit_count.setter
    def resize_unit_count(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "resize_unit_count", value)


@pulumi.input_type
class GoogleCloudChannelV1ValueArgs:
    def __init__(__self__, *,
                 bool_value: Optional[pulumi.Input[bool]] = None,
                 double_value: Optional[pulumi.Input[float]] = None,
                 int64_value: Optional[pulumi.Input[str]] = None,
                 proto_value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 string_value: Optional[pulumi.Input[str]] = None):
        """
        Data type and value of a parameter.
        :param pulumi.Input[bool] bool_value: Represents a boolean value.
        :param pulumi.Input[float] double_value: Represents a double value.
        :param pulumi.Input[str] int64_value: Represents an int64 value.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] proto_value: Represents an 'Any' proto value.
        :param pulumi.Input[str] string_value: Represents a string value.
        """
        if bool_value is not None:
            pulumi.set(__self__, "bool_value", bool_value)
        if double_value is not None:
            pulumi.set(__self__, "double_value", double_value)
        if int64_value is not None:
            pulumi.set(__self__, "int64_value", int64_value)
        if proto_value is not None:
            pulumi.set(__self__, "proto_value", proto_value)
        if string_value is not None:
            pulumi.set(__self__, "string_value", string_value)

    @property
    @pulumi.getter(name="boolValue")
    def bool_value(self) -> Optional[pulumi.Input[bool]]:
        """
        Represents a boolean value.
        """
        return pulumi.get(self, "bool_value")

    @bool_value.setter
    def bool_value(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "bool_value", value)

    @property
    @pulumi.getter(name="doubleValue")
    def double_value(self) -> Optional[pulumi.Input[float]]:
        """
        Represents a double value.
        """
        return pulumi.get(self, "double_value")

    @double_value.setter
    def double_value(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "double_value", value)

    @property
    @pulumi.getter(name="int64Value")
    def int64_value(self) -> Optional[pulumi.Input[str]]:
        """
        Represents an int64 value.
        """
        return pulumi.get(self, "int64_value")

    @int64_value.setter
    def int64_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "int64_value", value)

    @property
    @pulumi.getter(name="protoValue")
    def proto_value(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Represents an 'Any' proto value.
        """
        return pulumi.get(self, "proto_value")

    @proto_value.setter
    def proto_value(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "proto_value", value)

    @property
    @pulumi.getter(name="stringValue")
    def string_value(self) -> Optional[pulumi.Input[str]]:
        """
        Represents a string value.
        """
        return pulumi.get(self, "string_value")

    @string_value.setter
    def string_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "string_value", value)


@pulumi.input_type
class GoogleTypePostalAddressArgs:
    def __init__(__self__, *,
                 region_code: pulumi.Input[str],
                 address_lines: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 administrative_area: Optional[pulumi.Input[str]] = None,
                 language_code: Optional[pulumi.Input[str]] = None,
                 locality: Optional[pulumi.Input[str]] = None,
                 organization: Optional[pulumi.Input[str]] = None,
                 postal_code: Optional[pulumi.Input[str]] = None,
                 recipients: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 revision: Optional[pulumi.Input[int]] = None,
                 sorting_code: Optional[pulumi.Input[str]] = None,
                 sublocality: Optional[pulumi.Input[str]] = None):
        """
        Represents a postal address, e.g. for postal delivery or payments addresses. Given a postal address, a postal service can deliver items to a premise, P.O. Box or similar. It is not intended to model geographical locations (roads, towns, mountains). In typical usage an address would be created via user input or from importing existing data, depending on the type of process. Advice on address input / editing: - Use an i18n-ready address widget such as https://github.com/google/libaddressinput) - Users should not be presented with UI elements for input or editing of fields outside countries where that field is used. For more guidance on how to use this schema, please see: https://support.google.com/business/answer/6397478
        :param pulumi.Input[str] region_code: CLDR region code of the country/region of the address. This is never inferred and it is up to the user to ensure the value is correct. See https://cldr.unicode.org/ and https://www.unicode.org/cldr/charts/30/supplemental/territory_information.html for details. Example: "CH" for Switzerland.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] address_lines: Unstructured address lines describing the lower levels of an address. Because values in address_lines do not have type information and may sometimes contain multiple values in a single field (e.g. "Austin, TX"), it is important that the line order is clear. The order of address lines should be "envelope order" for the country/region of the address. In places where this can vary (e.g. Japan), address_language is used to make it explicit (e.g. "ja" for large-to-small ordering and "ja-Latn" or "en" for small-to-large). This way, the most specific line of an address can be selected based on the language. The minimum permitted structural representation of an address consists of a region_code with all remaining information placed in the address_lines. It would be possible to format such an address very approximately without geocoding, but no semantic reasoning could be made about any of the address components until it was at least partially resolved. Creating an address only containing a region_code and address_lines, and then geocoding is the recommended way to handle completely unstructured addresses (as opposed to guessing which parts of the address should be localities or administrative areas).
        :param pulumi.Input[str] administrative_area: Optional. Highest administrative subdivision which is used for postal addresses of a country or region. For example, this can be a state, a province, an oblast, or a prefecture. Specifically, for Spain this is the province and not the autonomous community (e.g. "Barcelona" and not "Catalonia"). Many countries don't use an administrative area in postal addresses. E.g. in Switzerland this should be left unpopulated.
        :param pulumi.Input[str] language_code: Optional. BCP-47 language code of the contents of this address (if known). This is often the UI language of the input form or is expected to match one of the languages used in the address' country/region, or their transliterated equivalents. This can affect formatting in certain countries, but is not critical to the correctness of the data and will never affect any validation or other non-formatting related operations. If this value is not known, it should be omitted (rather than specifying a possibly incorrect default). Examples: "zh-Hant", "ja", "ja-Latn", "en".
        :param pulumi.Input[str] locality: Optional. Generally refers to the city/town portion of the address. Examples: US city, IT comune, UK post town. In regions of the world where localities are not well defined or do not fit into this structure well, leave locality empty and use address_lines.
        :param pulumi.Input[str] organization: Optional. The name of the organization at the address.
        :param pulumi.Input[str] postal_code: Optional. Postal code of the address. Not all countries use or require postal codes to be present, but where they are used, they may trigger additional validation with other parts of the address (e.g. state/zip validation in the U.S.A.).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] recipients: Optional. The recipient at the address. This field may, under certain circumstances, contain multiline information. For example, it might contain "care of" information.
        :param pulumi.Input[int] revision: The schema revision of the `PostalAddress`. This must be set to 0, which is the latest revision. All new revisions **must** be backward compatible with old revisions.
        :param pulumi.Input[str] sorting_code: Optional. Additional, country-specific, sorting code. This is not used in most regions. Where it is used, the value is either a string like "CEDEX", optionally followed by a number (e.g. "CEDEX 7"), or just a number alone, representing the "sector code" (Jamaica), "delivery area indicator" (Malawi) or "post office indicator" (e.g. Côte d'Ivoire).
        :param pulumi.Input[str] sublocality: Optional. Sublocality of the address. For example, this can be neighborhoods, boroughs, districts.
        """
        pulumi.set(__self__, "region_code", region_code)
        if address_lines is not None:
            pulumi.set(__self__, "address_lines", address_lines)
        if administrative_area is not None:
            pulumi.set(__self__, "administrative_area", administrative_area)
        if language_code is not None:
            pulumi.set(__self__, "language_code", language_code)
        if locality is not None:
            pulumi.set(__self__, "locality", locality)
        if organization is not None:
            pulumi.set(__self__, "organization", organization)
        if postal_code is not None:
            pulumi.set(__self__, "postal_code", postal_code)
        if recipients is not None:
            pulumi.set(__self__, "recipients", recipients)
        if revision is not None:
            pulumi.set(__self__, "revision", revision)
        if sorting_code is not None:
            pulumi.set(__self__, "sorting_code", sorting_code)
        if sublocality is not None:
            pulumi.set(__self__, "sublocality", sublocality)

    @property
    @pulumi.getter(name="regionCode")
    def region_code(self) -> pulumi.Input[str]:
        """
        CLDR region code of the country/region of the address. This is never inferred and it is up to the user to ensure the value is correct. See https://cldr.unicode.org/ and https://www.unicode.org/cldr/charts/30/supplemental/territory_information.html for details. Example: "CH" for Switzerland.
        """
        return pulumi.get(self, "region_code")

    @region_code.setter
    def region_code(self, value: pulumi.Input[str]):
        pulumi.set(self, "region_code", value)

    @property
    @pulumi.getter(name="addressLines")
    def address_lines(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Unstructured address lines describing the lower levels of an address. Because values in address_lines do not have type information and may sometimes contain multiple values in a single field (e.g. "Austin, TX"), it is important that the line order is clear. The order of address lines should be "envelope order" for the country/region of the address. In places where this can vary (e.g. Japan), address_language is used to make it explicit (e.g. "ja" for large-to-small ordering and "ja-Latn" or "en" for small-to-large). This way, the most specific line of an address can be selected based on the language. The minimum permitted structural representation of an address consists of a region_code with all remaining information placed in the address_lines. It would be possible to format such an address very approximately without geocoding, but no semantic reasoning could be made about any of the address components until it was at least partially resolved. Creating an address only containing a region_code and address_lines, and then geocoding is the recommended way to handle completely unstructured addresses (as opposed to guessing which parts of the address should be localities or administrative areas).
        """
        return pulumi.get(self, "address_lines")

    @address_lines.setter
    def address_lines(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "address_lines", value)

    @property
    @pulumi.getter(name="administrativeArea")
    def administrative_area(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Highest administrative subdivision which is used for postal addresses of a country or region. For example, this can be a state, a province, an oblast, or a prefecture. Specifically, for Spain this is the province and not the autonomous community (e.g. "Barcelona" and not "Catalonia"). Many countries don't use an administrative area in postal addresses. E.g. in Switzerland this should be left unpopulated.
        """
        return pulumi.get(self, "administrative_area")

    @administrative_area.setter
    def administrative_area(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "administrative_area", value)

    @property
    @pulumi.getter(name="languageCode")
    def language_code(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. BCP-47 language code of the contents of this address (if known). This is often the UI language of the input form or is expected to match one of the languages used in the address' country/region, or their transliterated equivalents. This can affect formatting in certain countries, but is not critical to the correctness of the data and will never affect any validation or other non-formatting related operations. If this value is not known, it should be omitted (rather than specifying a possibly incorrect default). Examples: "zh-Hant", "ja", "ja-Latn", "en".
        """
        return pulumi.get(self, "language_code")

    @language_code.setter
    def language_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "language_code", value)

    @property
    @pulumi.getter
    def locality(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Generally refers to the city/town portion of the address. Examples: US city, IT comune, UK post town. In regions of the world where localities are not well defined or do not fit into this structure well, leave locality empty and use address_lines.
        """
        return pulumi.get(self, "locality")

    @locality.setter
    def locality(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "locality", value)

    @property
    @pulumi.getter
    def organization(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. The name of the organization at the address.
        """
        return pulumi.get(self, "organization")

    @organization.setter
    def organization(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization", value)

    @property
    @pulumi.getter(name="postalCode")
    def postal_code(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Postal code of the address. Not all countries use or require postal codes to be present, but where they are used, they may trigger additional validation with other parts of the address (e.g. state/zip validation in the U.S.A.).
        """
        return pulumi.get(self, "postal_code")

    @postal_code.setter
    def postal_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "postal_code", value)

    @property
    @pulumi.getter
    def recipients(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Optional. The recipient at the address. This field may, under certain circumstances, contain multiline information. For example, it might contain "care of" information.
        """
        return pulumi.get(self, "recipients")

    @recipients.setter
    def recipients(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "recipients", value)

    @property
    @pulumi.getter
    def revision(self) -> Optional[pulumi.Input[int]]:
        """
        The schema revision of the `PostalAddress`. This must be set to 0, which is the latest revision. All new revisions **must** be backward compatible with old revisions.
        """
        return pulumi.get(self, "revision")

    @revision.setter
    def revision(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "revision", value)

    @property
    @pulumi.getter(name="sortingCode")
    def sorting_code(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Additional, country-specific, sorting code. This is not used in most regions. Where it is used, the value is either a string like "CEDEX", optionally followed by a number (e.g. "CEDEX 7"), or just a number alone, representing the "sector code" (Jamaica), "delivery area indicator" (Malawi) or "post office indicator" (e.g. Côte d'Ivoire).
        """
        return pulumi.get(self, "sorting_code")

    @sorting_code.setter
    def sorting_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sorting_code", value)

    @property
    @pulumi.getter
    def sublocality(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Sublocality of the address. For example, this can be neighborhoods, boroughs, districts.
        """
        return pulumi.get(self, "sublocality")

    @sublocality.setter
    def sublocality(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sublocality", value)


