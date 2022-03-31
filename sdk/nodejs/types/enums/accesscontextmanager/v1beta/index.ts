// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const BasicLevelCombiningFunction = {
    /**
     * All `Conditions` must be true for the `BasicLevel` to be true.
     */
    And: "AND",
    /**
     * If at least one `Condition` is true, then the `BasicLevel` is true.
     */
    Or: "OR",
} as const;

/**
 * How the `conditions` list should be combined to determine if a request is granted this `AccessLevel`. If AND is used, each `Condition` in `conditions` must be satisfied for the `AccessLevel` to be applied. If OR is used, at least one `Condition` in `conditions` must be satisfied for the `AccessLevel` to be applied. Default behavior is AND.
 */
export type BasicLevelCombiningFunction = (typeof BasicLevelCombiningFunction)[keyof typeof BasicLevelCombiningFunction];

export const DevicePolicyAllowedDeviceManagementLevelsItem = {
    /**
     * The device's management level is not specified or not known.
     */
    ManagementUnspecified: "MANAGEMENT_UNSPECIFIED",
    /**
     * The device is not managed.
     */
    None: "NONE",
    /**
     * Basic management is enabled, which is generally limited to monitoring and wiping the corporate account.
     */
    Basic: "BASIC",
    /**
     * Complete device management. This includes more thorough monitoring and the ability to directly manage the device (such as remote wiping). This can be enabled through the Android Enterprise Platform.
     */
    Complete: "COMPLETE",
} as const;

export type DevicePolicyAllowedDeviceManagementLevelsItem = (typeof DevicePolicyAllowedDeviceManagementLevelsItem)[keyof typeof DevicePolicyAllowedDeviceManagementLevelsItem];

export const DevicePolicyAllowedEncryptionStatusesItem = {
    /**
     * The encryption status of the device is not specified or not known.
     */
    EncryptionUnspecified: "ENCRYPTION_UNSPECIFIED",
    /**
     * The device does not support encryption.
     */
    EncryptionUnsupported: "ENCRYPTION_UNSUPPORTED",
    /**
     * The device supports encryption, but is currently unencrypted.
     */
    Unencrypted: "UNENCRYPTED",
    /**
     * The device is encrypted.
     */
    Encrypted: "ENCRYPTED",
} as const;

export type DevicePolicyAllowedEncryptionStatusesItem = (typeof DevicePolicyAllowedEncryptionStatusesItem)[keyof typeof DevicePolicyAllowedEncryptionStatusesItem];

export const OsConstraintOsType = {
    /**
     * The operating system of the device is not specified or not known.
     */
    OsUnspecified: "OS_UNSPECIFIED",
    /**
     * A desktop Mac operating system.
     */
    DesktopMac: "DESKTOP_MAC",
    /**
     * A desktop Windows operating system.
     */
    DesktopWindows: "DESKTOP_WINDOWS",
    /**
     * A desktop Linux operating system.
     */
    DesktopLinux: "DESKTOP_LINUX",
    /**
     * A desktop ChromeOS operating system.
     */
    DesktopChromeOs: "DESKTOP_CHROME_OS",
    /**
     * An Android operating system.
     */
    Android: "ANDROID",
    /**
     * An iOS operating system.
     */
    Ios: "IOS",
} as const;

/**
 * Required. The allowed OS type.
 */
export type OsConstraintOsType = (typeof OsConstraintOsType)[keyof typeof OsConstraintOsType];

export const ServicePerimeterPerimeterType = {
    /**
     * Regular Perimeter.
     */
    PerimeterTypeRegular: "PERIMETER_TYPE_REGULAR",
    /**
     * Perimeter Bridge.
     */
    PerimeterTypeBridge: "PERIMETER_TYPE_BRIDGE",
} as const;

/**
 * Perimeter type indicator. A single project is allowed to be a member of single regular perimeter, but multiple service perimeter bridges. A project cannot be a included in a perimeter bridge without being included in regular perimeter. For perimeter bridges, restricted/unrestricted service lists as well as access lists must be empty.
 */
export type ServicePerimeterPerimeterType = (typeof ServicePerimeterPerimeterType)[keyof typeof ServicePerimeterPerimeterType];
