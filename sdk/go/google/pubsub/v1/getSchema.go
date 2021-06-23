// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a schema.
func LookupSchema(ctx *pulumi.Context, args *LookupSchemaArgs, opts ...pulumi.InvokeOption) (*LookupSchemaResult, error) {
	var rv LookupSchemaResult
	err := ctx.Invoke("google-native:pubsub/v1:getSchema", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSchemaArgs struct {
	Project  string  `pulumi:"project"`
	SchemaId string  `pulumi:"schemaId"`
	View     *string `pulumi:"view"`
}

type LookupSchemaResult struct {
	// The definition of the schema. This should contain a string representing the full definition of the schema that is a valid schema definition of the type specified in `type`.
	Definition string `pulumi:"definition"`
	// Required. Name of the schema. Format is `projects/{project}/schemas/{schema}`.
	Name string `pulumi:"name"`
	// The type of the schema definition.
	Type string `pulumi:"type"`
}