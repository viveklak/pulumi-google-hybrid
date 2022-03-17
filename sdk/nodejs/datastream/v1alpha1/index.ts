// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./connectionProfile";
export * from "./getConnectionProfile";
export * from "./getPrivateConnection";
export * from "./getRoute";
export * from "./getStream";
export * from "./privateConnection";
export * from "./route";
export * from "./stream";

// Export enums:
export * from "../../types/enums/datastream/v1alpha1";

// Import resources to register:
import { ConnectionProfile } from "./connectionProfile";
import { PrivateConnection } from "./privateConnection";
import { Route } from "./route";
import { Stream } from "./stream";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-native:datastream/v1alpha1:ConnectionProfile":
                return new ConnectionProfile(name, <any>undefined, { urn })
            case "google-native:datastream/v1alpha1:PrivateConnection":
                return new PrivateConnection(name, <any>undefined, { urn })
            case "google-native:datastream/v1alpha1:Route":
                return new Route(name, <any>undefined, { urn })
            case "google-native:datastream/v1alpha1:Stream":
                return new Stream(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-native", "datastream/v1alpha1", _module)
