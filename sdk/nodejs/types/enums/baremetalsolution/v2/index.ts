// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const NetworkConfigBandwidth = {
    /**
     * Unspecified value.
     */
    BandwidthUnspecified: "BANDWIDTH_UNSPECIFIED",
    /**
     * 1 Gbps.
     */
    Bw1Gbps: "BW_1_GBPS",
    /**
     * 2 Gbps.
     */
    Bw2Gbps: "BW_2_GBPS",
    /**
     * 5 Gbps.
     */
    Bw5Gbps: "BW_5_GBPS",
    /**
     * 10 Gbps.
     */
    Bw10Gbps: "BW_10_GBPS",
} as const;

/**
 * Interconnect bandwidth. Set only when type is CLIENT.
 */
export type NetworkConfigBandwidth = (typeof NetworkConfigBandwidth)[keyof typeof NetworkConfigBandwidth];

export const NetworkConfigServiceCidr = {
    /**
     * Unspecified value.
     */
    ServiceCidrUnspecified: "SERVICE_CIDR_UNSPECIFIED",
    /**
     * Services are disabled for the given network.
     */
    Disabled: "DISABLED",
    /**
     * Use the highest /26 block of the network to host services.
     */
    High26: "HIGH_26",
    /**
     * Use the highest /27 block of the network to host services.
     */
    High27: "HIGH_27",
    /**
     * Use the highest /28 block of the network to host services.
     */
    High28: "HIGH_28",
} as const;

/**
 * Service CIDR, if any.
 */
export type NetworkConfigServiceCidr = (typeof NetworkConfigServiceCidr)[keyof typeof NetworkConfigServiceCidr];

export const NetworkConfigType = {
    /**
     * Unspecified value.
     */
    TypeUnspecified: "TYPE_UNSPECIFIED",
    /**
     * Client network, that is a network peered to a GCP VPC.
     */
    Client: "CLIENT",
    /**
     * Private network, that is a network local to the BMS POD.
     */
    Private: "PRIVATE",
} as const;

/**
 * The type of this network, either Client or Private.
 */
export type NetworkConfigType = (typeof NetworkConfigType)[keyof typeof NetworkConfigType];

export const NfsExportPermissions = {
    /**
     * Unspecified value.
     */
    PermissionsUnspecified: "PERMISSIONS_UNSPECIFIED",
    /**
     * Read-only permission.
     */
    ReadOnly: "READ_ONLY",
    /**
     * Read-write permission.
     */
    ReadWrite: "READ_WRITE",
} as const;

/**
 * Export permissions.
 */
export type NfsExportPermissions = (typeof NfsExportPermissions)[keyof typeof NfsExportPermissions];

export const SnapshotSchedulePolicyState = {
    /**
     * The policy is in an unknown state.
     */
    StateUnspecified: "STATE_UNSPECIFIED",
    /**
     * The policy is been provisioned.
     */
    Provisioned: "PROVISIONED",
} as const;

/**
 * The state of the snapshot schedule policy.
 */
export type SnapshotSchedulePolicyState = (typeof SnapshotSchedulePolicyState)[keyof typeof SnapshotSchedulePolicyState];

export const VolumeConfigProtocol = {
    /**
     * Unspecified value.
     */
    ProtocolUnspecified: "PROTOCOL_UNSPECIFIED",
    /**
     * Fibre channel.
     */
    ProtocolFc: "PROTOCOL_FC",
    /**
     * Network file system.
     */
    ProtocolNfs: "PROTOCOL_NFS",
} as const;

/**
 * Volume protocol.
 */
export type VolumeConfigProtocol = (typeof VolumeConfigProtocol)[keyof typeof VolumeConfigProtocol];

export const VolumeConfigType = {
    /**
     * The unspecified type.
     */
    TypeUnspecified: "TYPE_UNSPECIFIED",
    /**
     * This Volume is on flash.
     */
    Flash: "FLASH",
    /**
     * This Volume is on disk.
     */
    Disk: "DISK",
} as const;

/**
 * The type of this Volume.
 */
export type VolumeConfigType = (typeof VolumeConfigType)[keyof typeof VolumeConfigType];
