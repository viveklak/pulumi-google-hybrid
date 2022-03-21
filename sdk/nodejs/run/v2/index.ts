// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./getService";
export * from "./getServiceIamPolicy";
export * from "./service";
export * from "./serviceIamPolicy";

// Export enums:
export * from "../../types/enums/run/v2";

// Import resources to register:
import { Service } from "./service";
import { ServiceIamPolicy } from "./serviceIamPolicy";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-native:run/v2:Service":
                return new Service(name, <any>undefined, { urn })
            case "google-native:run/v2:ServiceIamPolicy":
                return new ServiceIamPolicy(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-native", "run/v2", _module)
