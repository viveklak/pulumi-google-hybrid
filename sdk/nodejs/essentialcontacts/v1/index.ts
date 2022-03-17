// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./contact";
export * from "./folderContact";
export * from "./getContact";
export * from "./getFolderContact";
export * from "./getOrganizationContact";
export * from "./organizationContact";

// Export enums:
export * from "../../types/enums/essentialcontacts/v1";

// Import resources to register:
import { Contact } from "./contact";
import { FolderContact } from "./folderContact";
import { OrganizationContact } from "./organizationContact";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-native:essentialcontacts/v1:Contact":
                return new Contact(name, <any>undefined, { urn })
            case "google-native:essentialcontacts/v1:FolderContact":
                return new FolderContact(name, <any>undefined, { urn })
            case "google-native:essentialcontacts/v1:OrganizationContact":
                return new OrganizationContact(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-native", "essentialcontacts/v1", _module)
