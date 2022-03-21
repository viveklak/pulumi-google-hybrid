// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information associated with a Product. Possible errors: * Returns NOT_FOUND if the Product does not exist.
func LookupProduct(ctx *pulumi.Context, args *LookupProductArgs, opts ...pulumi.InvokeOption) (*LookupProductResult, error) {
	var rv LookupProductResult
	err := ctx.Invoke("google-native:vision/v1:getProduct", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProductArgs struct {
	Location  string  `pulumi:"location"`
	ProductId string  `pulumi:"productId"`
	Project   *string `pulumi:"project"`
}

type LookupProductResult struct {
	// User-provided metadata to be stored with this product. Must be at most 4096 characters long.
	Description string `pulumi:"description"`
	// The user-provided name for this Product. Must not be empty. Must be at most 4096 characters long.
	DisplayName string `pulumi:"displayName"`
	// The resource name of the product. Format is: `projects/PROJECT_ID/locations/LOC_ID/products/PRODUCT_ID`. This field is ignored when creating a product.
	Name string `pulumi:"name"`
	// Immutable. The category for the product identified by the reference image. This should be one of "homegoods-v2", "apparel-v2", "toys-v2", "packagedgoods-v1" or "general-v1". The legacy categories "homegoods", "apparel", and "toys" are still supported, but these should not be used for new products.
	ProductCategory string `pulumi:"productCategory"`
	// Key-value pairs that can be attached to a product. At query time, constraints can be specified based on the product_labels. Note that integer values can be provided as strings, e.g. "1199". Only strings with integer values can match a range-based restriction which is to be supported soon. Multiple values can be assigned to the same key. One product may have up to 500 product_labels. Notice that the total number of distinct product_labels over all products in one ProductSet cannot exceed 1M, otherwise the product search pipeline will refuse to work for that ProductSet.
	ProductLabels []KeyValueResponse `pulumi:"productLabels"`
}

func LookupProductOutput(ctx *pulumi.Context, args LookupProductOutputArgs, opts ...pulumi.InvokeOption) LookupProductResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProductResult, error) {
			args := v.(LookupProductArgs)
			r, err := LookupProduct(ctx, &args, opts...)
			return *r, err
		}).(LookupProductResultOutput)
}

type LookupProductOutputArgs struct {
	Location  pulumi.StringInput    `pulumi:"location"`
	ProductId pulumi.StringInput    `pulumi:"productId"`
	Project   pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupProductOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProductArgs)(nil)).Elem()
}

type LookupProductResultOutput struct{ *pulumi.OutputState }

func (LookupProductResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProductResult)(nil)).Elem()
}

func (o LookupProductResultOutput) ToLookupProductResultOutput() LookupProductResultOutput {
	return o
}

func (o LookupProductResultOutput) ToLookupProductResultOutputWithContext(ctx context.Context) LookupProductResultOutput {
	return o
}

// User-provided metadata to be stored with this product. Must be at most 4096 characters long.
func (o LookupProductResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProductResult) string { return v.Description }).(pulumi.StringOutput)
}

// The user-provided name for this Product. Must not be empty. Must be at most 4096 characters long.
func (o LookupProductResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProductResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The resource name of the product. Format is: `projects/PROJECT_ID/locations/LOC_ID/products/PRODUCT_ID`. This field is ignored when creating a product.
func (o LookupProductResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProductResult) string { return v.Name }).(pulumi.StringOutput)
}

// Immutable. The category for the product identified by the reference image. This should be one of "homegoods-v2", "apparel-v2", "toys-v2", "packagedgoods-v1" or "general-v1". The legacy categories "homegoods", "apparel", and "toys" are still supported, but these should not be used for new products.
func (o LookupProductResultOutput) ProductCategory() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProductResult) string { return v.ProductCategory }).(pulumi.StringOutput)
}

// Key-value pairs that can be attached to a product. At query time, constraints can be specified based on the product_labels. Note that integer values can be provided as strings, e.g. "1199". Only strings with integer values can match a range-based restriction which is to be supported soon. Multiple values can be assigned to the same key. One product may have up to 500 product_labels. Notice that the total number of distinct product_labels over all products in one ProductSet cannot exceed 1M, otherwise the product search pipeline will refuse to work for that ProductSet.
func (o LookupProductResultOutput) ProductLabels() KeyValueResponseArrayOutput {
	return o.ApplyT(func(v LookupProductResult) []KeyValueResponse { return v.ProductLabels }).(KeyValueResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProductResultOutput{})
}
