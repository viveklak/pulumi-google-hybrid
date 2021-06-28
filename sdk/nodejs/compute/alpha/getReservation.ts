// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Retrieves information about the specified reservation.
 */
export function getReservation(args: GetReservationArgs, opts?: pulumi.InvokeOptions): Promise<GetReservationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:compute/alpha:getReservation", {
        "project": args.project,
        "reservation": args.reservation,
        "zone": args.zone,
    }, opts);
}

export interface GetReservationArgs {
    project: string;
    reservation: string;
    zone: string;
}

export interface GetReservationResult {
    /**
     * Full or partial URL to a parent commitment. This field displays for reservations that are tied to a commitment.
     */
    readonly commitment: string;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * Type of the resource. Always compute#reservations for reservations.
     */
    readonly kind: string;
    /**
     * The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * Reserved for future use.
     */
    readonly satisfiesPzs: boolean;
    /**
     * Server-defined fully-qualified URL for this resource.
     */
    readonly selfLink: string;
    /**
     * Server-defined URL for this resource with the resource id.
     */
    readonly selfLinkWithId: string;
    /**
     * Share-settings for shared-reservation
     */
    readonly shareSettings: outputs.compute.alpha.AllocationShareSettingsResponse;
    /**
     * Reservation for instances with specific machine shapes.
     */
    readonly specificReservation: outputs.compute.alpha.AllocationSpecificSKUReservationResponse;
    /**
     * Indicates whether the reservation can be consumed by VMs with affinity for "any" reservation. If the field is set, then only VMs that target the reservation by name can consume from this reservation.
     */
    readonly specificReservationRequired: boolean;
    /**
     * The status of the reservation.
     */
    readonly status: string;
    /**
     * Zone in which the reservation resides. A zone must be provided if the reservation is created within a commitment.
     */
    readonly zone: string;
}
