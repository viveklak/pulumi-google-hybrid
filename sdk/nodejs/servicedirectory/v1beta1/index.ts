// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./namespace";
export * from "./namespaceIamPolicy";
export * from "./namespaceService";
export * from "./namespaceServiceEndpoint";
export * from "./namespaceServiceIamPolicy";

// Import resources to register:
import { Namespace } from "./namespace";
import { NamespaceIamPolicy } from "./namespaceIamPolicy";
import { NamespaceService } from "./namespaceService";
import { NamespaceServiceEndpoint } from "./namespaceServiceEndpoint";
import { NamespaceServiceIamPolicy } from "./namespaceServiceIamPolicy";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-cloud:servicedirectory/v1beta1:Namespace":
                return new Namespace(name, <any>undefined, { urn })
            case "google-cloud:servicedirectory/v1beta1:NamespaceIamPolicy":
                return new NamespaceIamPolicy(name, <any>undefined, { urn })
            case "google-cloud:servicedirectory/v1beta1:NamespaceService":
                return new NamespaceService(name, <any>undefined, { urn })
            case "google-cloud:servicedirectory/v1beta1:NamespaceServiceEndpoint":
                return new NamespaceServiceEndpoint(name, <any>undefined, { urn })
            case "google-cloud:servicedirectory/v1beta1:NamespaceServiceIamPolicy":
                return new NamespaceServiceIamPolicy(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-cloud", "servicedirectory/v1beta1", _module)