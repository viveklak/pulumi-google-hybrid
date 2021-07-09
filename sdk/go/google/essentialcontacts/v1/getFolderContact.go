// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a single contact.
func LookupFolderContact(ctx *pulumi.Context, args *LookupFolderContactArgs, opts ...pulumi.InvokeOption) (*LookupFolderContactResult, error) {
	var rv LookupFolderContactResult
	err := ctx.Invoke("google-native:essentialcontacts/v1:getFolderContact", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFolderContactArgs struct {
	ContactId string `pulumi:"contactId"`
	FolderId  string `pulumi:"folderId"`
}

type LookupFolderContactResult struct {
	// The email address to send notifications to. This does not need to be a Google account.
	Email string `pulumi:"email"`
	// The preferred language for notifications, as a ISO 639-1 language code. See [Supported languages](https://cloud.google.com/resource-manager/docs/managing-notification-contacts#supported-languages) for a list of supported languages.
	LanguageTag string `pulumi:"languageTag"`
	// The identifier for the contact. Format: {resource_type}/{resource_id}/contacts/{contact_id}
	Name string `pulumi:"name"`
	// The categories of notifications that the contact will receive communications for.
	NotificationCategorySubscriptions []string `pulumi:"notificationCategorySubscriptions"`
	// The last time the validation_state was updated, either manually or automatically. A contact is considered stale if its validation state was updated more than 1 year ago.
	ValidateTime string `pulumi:"validateTime"`
	// The validity of the contact. A contact is considered valid if it is the correct recipient for notifications for a particular resource.
	ValidationState string `pulumi:"validationState"`
}
