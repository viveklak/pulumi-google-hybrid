// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./domain";
export * from "./domainIamPolicy";
export * from "./peeringIamPolicy";

// Import resources to register:
import { Domain } from "./domain";
import { DomainIamPolicy } from "./domainIamPolicy";
import { PeeringIamPolicy } from "./peeringIamPolicy";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-cloud:managedidentities/v1beta1:Domain":
                return new Domain(name, <any>undefined, { urn })
            case "google-cloud:managedidentities/v1beta1:DomainIamPolicy":
                return new DomainIamPolicy(name, <any>undefined, { urn })
            case "google-cloud:managedidentities/v1beta1:PeeringIamPolicy":
                return new PeeringIamPolicy(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-cloud", "managedidentities/v1beta1", _module)
