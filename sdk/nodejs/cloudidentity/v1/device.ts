// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a device. Only company-owned device may be created. **Note**: This method is available only to customers who have one of the following SKUs: Enterprise Standard, Enterprise Plus, Enterprise for Education, and Cloud Identity Premium
 */
export class Device extends pulumi.CustomResource {
    /**
     * Get an existing Device resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Device {
        return new Device(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:cloudidentity/v1:Device';

    /**
     * Returns true if the given object is an instance of Device.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Device {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Device.__pulumiType;
    }

    /**
     * Attributes specific to Android devices.
     */
    public /*out*/ readonly androidSpecificAttributes!: pulumi.Output<outputs.cloudidentity.v1.GoogleAppsCloudidentityDevicesV1AndroidAttributesResponse>;
    /**
     * Asset tag of the device.
     */
    public readonly assetTag!: pulumi.Output<string>;
    /**
     * Baseband version of the device.
     */
    public /*out*/ readonly basebandVersion!: pulumi.Output<string>;
    /**
     * Device bootloader version. Example: 0.6.7.
     */
    public /*out*/ readonly bootloaderVersion!: pulumi.Output<string>;
    /**
     * Device brand. Example: Samsung.
     */
    public /*out*/ readonly brand!: pulumi.Output<string>;
    /**
     * Build number of the device.
     */
    public /*out*/ readonly buildNumber!: pulumi.Output<string>;
    /**
     * Represents whether the Device is compromised.
     */
    public /*out*/ readonly compromisedState!: pulumi.Output<string>;
    /**
     * When the Company-Owned device was imported. This field is empty for BYOD devices.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Type of device.
     */
    public /*out*/ readonly deviceType!: pulumi.Output<string>;
    /**
     * Whether developer options is enabled on device.
     */
    public /*out*/ readonly enabledDeveloperOptions!: pulumi.Output<boolean>;
    /**
     * Whether USB debugging is enabled on device.
     */
    public /*out*/ readonly enabledUsbDebugging!: pulumi.Output<boolean>;
    /**
     * Device encryption state.
     */
    public /*out*/ readonly encryptionState!: pulumi.Output<string>;
    /**
     * IMEI number of device if GSM device; empty otherwise.
     */
    public /*out*/ readonly imei!: pulumi.Output<string>;
    /**
     * Kernel version of the device.
     */
    public /*out*/ readonly kernelVersion!: pulumi.Output<string>;
    /**
     * Most recent time when device synced with this service.
     */
    public readonly lastSyncTime!: pulumi.Output<string>;
    /**
     * Management state of the device
     */
    public /*out*/ readonly managementState!: pulumi.Output<string>;
    /**
     * Device manufacturer. Example: Motorola.
     */
    public /*out*/ readonly manufacturer!: pulumi.Output<string>;
    /**
     * MEID number of device if CDMA device; empty otherwise.
     */
    public /*out*/ readonly meid!: pulumi.Output<string>;
    /**
     * Model name of device. Example: Pixel 3.
     */
    public /*out*/ readonly model!: pulumi.Output<string>;
    /**
     * [Resource name](https://cloud.google.com/apis/design/resource_names) of the Device in format: `devices/{device_id}`, where device_id is the unique id assigned to the Device.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Mobile or network operator of device, if available.
     */
    public /*out*/ readonly networkOperator!: pulumi.Output<string>;
    /**
     * OS version of the device. Example: Android 8.1.0.
     */
    public /*out*/ readonly osVersion!: pulumi.Output<string>;
    /**
     * Domain name for Google accounts on device. Type for other accounts on device. On Android, will only be populated if |ownership_privilege| is |PROFILE_OWNER| or |DEVICE_OWNER|. Does not include the account signed in to the device policy app if that account's domain has only one account. Examples: "com.example", "xyz.com".
     */
    public /*out*/ readonly otherAccounts!: pulumi.Output<string[]>;
    /**
     * Whether the device is owned by the company or an individual
     */
    public /*out*/ readonly ownerType!: pulumi.Output<string>;
    /**
     * OS release version. Example: 6.0.
     */
    public /*out*/ readonly releaseVersion!: pulumi.Output<string>;
    /**
     * OS security patch update time on device.
     */
    public /*out*/ readonly securityPatchTime!: pulumi.Output<string>;
    /**
     * Serial Number of device. Example: HT82V1A01076.
     */
    public readonly serialNumber!: pulumi.Output<string>;
    /**
     * WiFi MAC addresses of device.
     */
    public readonly wifiMacAddresses!: pulumi.Output<string[]>;

    /**
     * Create a Device resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DeviceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            inputs["assetTag"] = args ? args.assetTag : undefined;
            inputs["customer"] = args ? args.customer : undefined;
            inputs["lastSyncTime"] = args ? args.lastSyncTime : undefined;
            inputs["serialNumber"] = args ? args.serialNumber : undefined;
            inputs["wifiMacAddresses"] = args ? args.wifiMacAddresses : undefined;
            inputs["androidSpecificAttributes"] = undefined /*out*/;
            inputs["basebandVersion"] = undefined /*out*/;
            inputs["bootloaderVersion"] = undefined /*out*/;
            inputs["brand"] = undefined /*out*/;
            inputs["buildNumber"] = undefined /*out*/;
            inputs["compromisedState"] = undefined /*out*/;
            inputs["createTime"] = undefined /*out*/;
            inputs["deviceType"] = undefined /*out*/;
            inputs["enabledDeveloperOptions"] = undefined /*out*/;
            inputs["enabledUsbDebugging"] = undefined /*out*/;
            inputs["encryptionState"] = undefined /*out*/;
            inputs["imei"] = undefined /*out*/;
            inputs["kernelVersion"] = undefined /*out*/;
            inputs["managementState"] = undefined /*out*/;
            inputs["manufacturer"] = undefined /*out*/;
            inputs["meid"] = undefined /*out*/;
            inputs["model"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["networkOperator"] = undefined /*out*/;
            inputs["osVersion"] = undefined /*out*/;
            inputs["otherAccounts"] = undefined /*out*/;
            inputs["ownerType"] = undefined /*out*/;
            inputs["releaseVersion"] = undefined /*out*/;
            inputs["securityPatchTime"] = undefined /*out*/;
        } else {
            inputs["androidSpecificAttributes"] = undefined /*out*/;
            inputs["assetTag"] = undefined /*out*/;
            inputs["basebandVersion"] = undefined /*out*/;
            inputs["bootloaderVersion"] = undefined /*out*/;
            inputs["brand"] = undefined /*out*/;
            inputs["buildNumber"] = undefined /*out*/;
            inputs["compromisedState"] = undefined /*out*/;
            inputs["createTime"] = undefined /*out*/;
            inputs["deviceType"] = undefined /*out*/;
            inputs["enabledDeveloperOptions"] = undefined /*out*/;
            inputs["enabledUsbDebugging"] = undefined /*out*/;
            inputs["encryptionState"] = undefined /*out*/;
            inputs["imei"] = undefined /*out*/;
            inputs["kernelVersion"] = undefined /*out*/;
            inputs["lastSyncTime"] = undefined /*out*/;
            inputs["managementState"] = undefined /*out*/;
            inputs["manufacturer"] = undefined /*out*/;
            inputs["meid"] = undefined /*out*/;
            inputs["model"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["networkOperator"] = undefined /*out*/;
            inputs["osVersion"] = undefined /*out*/;
            inputs["otherAccounts"] = undefined /*out*/;
            inputs["ownerType"] = undefined /*out*/;
            inputs["releaseVersion"] = undefined /*out*/;
            inputs["securityPatchTime"] = undefined /*out*/;
            inputs["serialNumber"] = undefined /*out*/;
            inputs["wifiMacAddresses"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Device.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Device resource.
 */
export interface DeviceArgs {
    /**
     * Asset tag of the device.
     */
    readonly assetTag?: pulumi.Input<string>;
    readonly customer?: pulumi.Input<string>;
    /**
     * Most recent time when device synced with this service.
     */
    readonly lastSyncTime?: pulumi.Input<string>;
    /**
     * Serial Number of device. Example: HT82V1A01076.
     */
    readonly serialNumber?: pulumi.Input<string>;
    /**
     * WiFi MAC addresses of device.
     */
    readonly wifiMacAddresses?: pulumi.Input<pulumi.Input<string>[]>;
}
