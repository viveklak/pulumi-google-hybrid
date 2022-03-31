// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./getRelease";
export * from "./getRuleset";
export * from "./release";
export * from "./ruleset";

// Import resources to register:
import { Release } from "./release";
import { Ruleset } from "./ruleset";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-native:firebaserules/v1:Release":
                return new Release(name, <any>undefined, { urn })
            case "google-native:firebaserules/v1:Ruleset":
                return new Ruleset(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-native", "firebaserules/v1", _module)
