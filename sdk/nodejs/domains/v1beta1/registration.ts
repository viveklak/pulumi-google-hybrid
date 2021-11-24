// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Registers a new domain name and creates a corresponding `Registration` resource. Call `RetrieveRegisterParameters` first to check availability of the domain name and determine parameters like price that are needed to build a call to this method. A successful call creates a `Registration` resource in state `REGISTRATION_PENDING`, which resolves to `ACTIVE` within 1-2 minutes, indicating that the domain was successfully registered. If the resource ends up in state `REGISTRATION_FAILED`, it indicates that the domain was not registered successfully, and you can safely delete the resource and retry registration.
 * Auto-naming is currently not supported for this resource.
 */
export class Registration extends pulumi.CustomResource {
    /**
     * Get an existing Registration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Registration {
        return new Registration(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:domains/v1beta1:Registration';

    /**
     * Returns true if the given object is an instance of Registration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Registration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Registration.__pulumiType;
    }

    /**
     * Settings for contact information linked to the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureContactSettings` method.
     */
    public readonly contactSettings!: pulumi.Output<outputs.domains.v1beta1.ContactSettingsResponse>;
    /**
     * The creation timestamp of the `Registration` resource.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Settings controlling the DNS configuration of the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureDnsSettings` method.
     */
    public readonly dnsSettings!: pulumi.Output<outputs.domains.v1beta1.DnsSettingsResponse>;
    /**
     * Immutable. The domain name. Unicode domain names must be expressed in Punycode format.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * The expiration timestamp of the `Registration`.
     */
    public /*out*/ readonly expireTime!: pulumi.Output<string>;
    /**
     * The set of issues with the `Registration` that require attention.
     */
    public /*out*/ readonly issues!: pulumi.Output<string[]>;
    /**
     * Set of labels associated with the `Registration`.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Settings for management of the `Registration`, including renewal, billing, and transfer. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureManagementSettings` method.
     */
    public readonly managementSettings!: pulumi.Output<outputs.domains.v1beta1.ManagementSettingsResponse>;
    /**
     * Name of the `Registration` resource, in the format `projects/*&#47;locations/*&#47;registrations/`.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Pending contact settings for the `Registration`. Updates to the `contact_settings` field that change its `registrant_contact` or `privacy` fields require email confirmation by the `registrant_contact` before taking effect. This field is set only if there are pending updates to the `contact_settings` that have not been confirmed. To confirm the changes, the `registrant_contact` must follow the instructions in the email they receive.
     */
    public /*out*/ readonly pendingContactSettings!: pulumi.Output<outputs.domains.v1beta1.ContactSettingsResponse>;
    /**
     * The state of the `Registration`
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Set of options for the `contact_settings.privacy` field that this `Registration` supports.
     */
    public /*out*/ readonly supportedPrivacy!: pulumi.Output<string[]>;

    /**
     * Create a Registration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegistrationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.contactSettings === undefined) && !opts.urn) {
                throw new Error("Missing required property 'contactSettings'");
            }
            if ((!args || args.domainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainName'");
            }
            if ((!args || args.yearlyPrice === undefined) && !opts.urn) {
                throw new Error("Missing required property 'yearlyPrice'");
            }
            resourceInputs["contactNotices"] = args ? args.contactNotices : undefined;
            resourceInputs["contactSettings"] = args ? args.contactSettings : undefined;
            resourceInputs["dnsSettings"] = args ? args.dnsSettings : undefined;
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["domainNotices"] = args ? args.domainNotices : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["managementSettings"] = args ? args.managementSettings : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["validateOnly"] = args ? args.validateOnly : undefined;
            resourceInputs["yearlyPrice"] = args ? args.yearlyPrice : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["expireTime"] = undefined /*out*/;
            resourceInputs["issues"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["pendingContactSettings"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["supportedPrivacy"] = undefined /*out*/;
        } else {
            resourceInputs["contactSettings"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["dnsSettings"] = undefined /*out*/;
            resourceInputs["domainName"] = undefined /*out*/;
            resourceInputs["expireTime"] = undefined /*out*/;
            resourceInputs["issues"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["managementSettings"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["pendingContactSettings"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["supportedPrivacy"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Registration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Registration resource.
 */
export interface RegistrationArgs {
    /**
     * The list of contact notices that the caller acknowledges. The notices needed here depend on the values specified in `registration.contact_settings`.
     */
    contactNotices?: pulumi.Input<pulumi.Input<enums.domains.v1beta1.RegistrationContactNoticesItem>[]>;
    /**
     * Settings for contact information linked to the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureContactSettings` method.
     */
    contactSettings: pulumi.Input<inputs.domains.v1beta1.ContactSettingsArgs>;
    /**
     * Settings controlling the DNS configuration of the `Registration`. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureDnsSettings` method.
     */
    dnsSettings?: pulumi.Input<inputs.domains.v1beta1.DnsSettingsArgs>;
    /**
     * Immutable. The domain name. Unicode domain names must be expressed in Punycode format.
     */
    domainName: pulumi.Input<string>;
    /**
     * The list of domain notices that you acknowledge. Call `RetrieveRegisterParameters` to see the notices that need acknowledgement.
     */
    domainNotices?: pulumi.Input<pulumi.Input<enums.domains.v1beta1.RegistrationDomainNoticesItem>[]>;
    /**
     * Set of labels associated with the `Registration`.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * Settings for management of the `Registration`, including renewal, billing, and transfer. You cannot update these with the `UpdateRegistration` method. To update these settings, use the `ConfigureManagementSettings` method.
     */
    managementSettings?: pulumi.Input<inputs.domains.v1beta1.ManagementSettingsArgs>;
    project?: pulumi.Input<string>;
    /**
     * When true, only validation is performed, without actually registering the domain. Follows: https://cloud.google.com/apis/design/design_patterns#request_validation
     */
    validateOnly?: pulumi.Input<boolean>;
    /**
     * Yearly price to register or renew the domain. The value that should be put here can be obtained from RetrieveRegisterParameters or SearchDomains calls.
     */
    yearlyPrice: pulumi.Input<inputs.domains.v1beta1.MoneyArgs>;
}
