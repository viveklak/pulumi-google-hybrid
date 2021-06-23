// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the contents of a FHIR resource. Implements the FHIR standard read interaction ([DSTU2](https://hl7.org/implement/standards/fhir/DSTU2/http.html#read), [STU3](https://hl7.org/implement/standards/fhir/STU3/http.html#read), [R4](https://hl7.org/implement/standards/fhir/R4/http.html#read)). Also supports the FHIR standard conditional read interaction ([DSTU2](https://hl7.org/implement/standards/fhir/DSTU2/http.html#cread), [STU3](https://hl7.org/implement/standards/fhir/STU3/http.html#cread), [R4](https://hl7.org/implement/standards/fhir/R4/http.html#cread)) specified by supplying an `If-Modified-Since` header with a date/time value or an `If-None-Match` header with an ETag value. On success, the response body contains a JSON-encoded representation of the resource. Errors generated by the FHIR store contain a JSON-encoded `OperationOutcome` resource describing the reason for the error. If the request cannot be mapped to a valid API method on a FHIR store, a generic GCP error might be returned instead. For samples that show how to call `read`, see [Getting a FHIR resource](/healthcare/docs/how-tos/fhir-resources#getting_a_fhir_resource).
func GetFhir(ctx *pulumi.Context, args *GetFhirArgs, opts ...pulumi.InvokeOption) (*GetFhirResult, error) {
	var rv GetFhirResult
	err := ctx.Invoke("google-native:healthcare/v1beta1:getFhir", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetFhirArgs struct {
	DatasetId   string `pulumi:"datasetId"`
	FhirId      string `pulumi:"fhirId"`
	FhirId1     string `pulumi:"fhirId1"`
	FhirStoreId string `pulumi:"fhirStoreId"`
	Location    string `pulumi:"location"`
	Project     string `pulumi:"project"`
}

type GetFhirResult struct {
	// The HTTP Content-Type header value specifying the content type of the body.
	ContentType string `pulumi:"contentType"`
	// The HTTP request/response body as raw binary.
	Data string `pulumi:"data"`
	// Application specific response metadata. Must be set in the first response for streaming APIs.
	Extensions []map[string]string `pulumi:"extensions"`
}