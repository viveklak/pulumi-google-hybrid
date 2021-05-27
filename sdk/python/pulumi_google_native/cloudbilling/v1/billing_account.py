# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['BillingAccountArgs', 'BillingAccount']

@pulumi.input_type
class BillingAccountArgs:
    def __init__(__self__, *,
                 display_name: Optional[pulumi.Input[str]] = None,
                 master_billing_account: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a BillingAccount resource.
        :param pulumi.Input[str] display_name: The display name given to the billing account, such as `My Billing Account`. This name is displayed in the Google Cloud Console.
        :param pulumi.Input[str] master_billing_account: If this account is a [subaccount](https://cloud.google.com/billing/docs/concepts), then this will be the resource name of the parent billing account that it is being resold through. Otherwise this will be empty.
        """
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if master_billing_account is not None:
            pulumi.set(__self__, "master_billing_account", master_billing_account)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name given to the billing account, such as `My Billing Account`. This name is displayed in the Google Cloud Console.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="masterBillingAccount")
    def master_billing_account(self) -> Optional[pulumi.Input[str]]:
        """
        If this account is a [subaccount](https://cloud.google.com/billing/docs/concepts), then this will be the resource name of the parent billing account that it is being resold through. Otherwise this will be empty.
        """
        return pulumi.get(self, "master_billing_account")

    @master_billing_account.setter
    def master_billing_account(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "master_billing_account", value)


class BillingAccount(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 master_billing_account: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This method creates [billing subaccounts](https://cloud.google.com/billing/docs/concepts#subaccounts). Google Cloud resellers should use the Channel Services APIs, [accounts.customers.create](https://cloud.google.com/channel/docs/reference/rest/v1/accounts.customers/create) and [accounts.customers.entitlements.create](https://cloud.google.com/channel/docs/reference/rest/v1/accounts.customers.entitlements/create). When creating a subaccount, the current authenticated user must have the `billing.accounts.update` IAM permission on the parent account, which is typically given to billing account [administrators](https://cloud.google.com/billing/docs/how-to/billing-access). This method will return an error if the parent account has not been provisioned as a reseller account.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: The display name given to the billing account, such as `My Billing Account`. This name is displayed in the Google Cloud Console.
        :param pulumi.Input[str] master_billing_account: If this account is a [subaccount](https://cloud.google.com/billing/docs/concepts), then this will be the resource name of the parent billing account that it is being resold through. Otherwise this will be empty.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[BillingAccountArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This method creates [billing subaccounts](https://cloud.google.com/billing/docs/concepts#subaccounts). Google Cloud resellers should use the Channel Services APIs, [accounts.customers.create](https://cloud.google.com/channel/docs/reference/rest/v1/accounts.customers/create) and [accounts.customers.entitlements.create](https://cloud.google.com/channel/docs/reference/rest/v1/accounts.customers.entitlements/create). When creating a subaccount, the current authenticated user must have the `billing.accounts.update` IAM permission on the parent account, which is typically given to billing account [administrators](https://cloud.google.com/billing/docs/how-to/billing-access). This method will return an error if the parent account has not been provisioned as a reseller account.

        :param str resource_name: The name of the resource.
        :param BillingAccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BillingAccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 master_billing_account: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BillingAccountArgs.__new__(BillingAccountArgs)

            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["master_billing_account"] = master_billing_account
            __props__.__dict__["name"] = None
            __props__.__dict__["open"] = None
        super(BillingAccount, __self__).__init__(
            'google-native:cloudbilling/v1:BillingAccount',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'BillingAccount':
        """
        Get an existing BillingAccount resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = BillingAccountArgs.__new__(BillingAccountArgs)

        __props__.__dict__["display_name"] = None
        __props__.__dict__["master_billing_account"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["open"] = None
        return BillingAccount(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The display name given to the billing account, such as `My Billing Account`. This name is displayed in the Google Cloud Console.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="masterBillingAccount")
    def master_billing_account(self) -> pulumi.Output[str]:
        """
        If this account is a [subaccount](https://cloud.google.com/billing/docs/concepts), then this will be the resource name of the parent billing account that it is being resold through. Otherwise this will be empty.
        """
        return pulumi.get(self, "master_billing_account")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name of the billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, `billingAccounts/012345-567890-ABCDEF` would be the resource name for billing account `012345-567890-ABCDEF`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def open(self) -> pulumi.Output[bool]:
        """
        True if the billing account is open, and will therefore be charged for any usage on associated projects. False if the billing account is closed, and therefore projects associated with it will be unable to use paid services.
        """
        return pulumi.get(self, "open")

