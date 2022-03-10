// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./device";
export * from "./getDevice";
export * from "./getGroup";
export * from "./getMembership";
export * from "./group";
export * from "./membership";

// Export enums:
export * from "../../types/enums/cloudidentity/v1";

// Import resources to register:
import { Device } from "./device";
import { Group } from "./group";
import { Membership } from "./membership";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-native:cloudidentity/v1:Device":
                return new Device(name, <any>undefined, { urn })
            case "google-native:cloudidentity/v1:Group":
                return new Group(name, <any>undefined, { urn })
            case "google-native:cloudidentity/v1:Membership":
                return new Membership(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-native", "cloudidentity/v1", _module)