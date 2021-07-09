// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v2alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Product thumbnail/detail image.
type GoogleCloudRetailV2alphaImage struct {
	// Height of the image in number of pixels. This field must be nonnegative. Otherwise, an INVALID_ARGUMENT error is returned.
	Height *int `pulumi:"height"`
	// URI of the image. This field must be a valid UTF-8 encoded URI with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [image_link](https://support.google.com/merchants/answer/6324350). Schema.org property [Product.image](https://schema.org/image).
	Uri string `pulumi:"uri"`
	// Width of the image in number of pixels. This field must be nonnegative. Otherwise, an INVALID_ARGUMENT error is returned.
	Width *int `pulumi:"width"`
}

// GoogleCloudRetailV2alphaImageInput is an input type that accepts GoogleCloudRetailV2alphaImageArgs and GoogleCloudRetailV2alphaImageOutput values.
// You can construct a concrete instance of `GoogleCloudRetailV2alphaImageInput` via:
//
//          GoogleCloudRetailV2alphaImageArgs{...}
type GoogleCloudRetailV2alphaImageInput interface {
	pulumi.Input

	ToGoogleCloudRetailV2alphaImageOutput() GoogleCloudRetailV2alphaImageOutput
	ToGoogleCloudRetailV2alphaImageOutputWithContext(context.Context) GoogleCloudRetailV2alphaImageOutput
}

// Product thumbnail/detail image.
type GoogleCloudRetailV2alphaImageArgs struct {
	// Height of the image in number of pixels. This field must be nonnegative. Otherwise, an INVALID_ARGUMENT error is returned.
	Height pulumi.IntPtrInput `pulumi:"height"`
	// URI of the image. This field must be a valid UTF-8 encoded URI with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [image_link](https://support.google.com/merchants/answer/6324350). Schema.org property [Product.image](https://schema.org/image).
	Uri pulumi.StringInput `pulumi:"uri"`
	// Width of the image in number of pixels. This field must be nonnegative. Otherwise, an INVALID_ARGUMENT error is returned.
	Width pulumi.IntPtrInput `pulumi:"width"`
}

func (GoogleCloudRetailV2alphaImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleCloudRetailV2alphaImage)(nil)).Elem()
}

func (i GoogleCloudRetailV2alphaImageArgs) ToGoogleCloudRetailV2alphaImageOutput() GoogleCloudRetailV2alphaImageOutput {
	return i.ToGoogleCloudRetailV2alphaImageOutputWithContext(context.Background())
}

func (i GoogleCloudRetailV2alphaImageArgs) ToGoogleCloudRetailV2alphaImageOutputWithContext(ctx context.Context) GoogleCloudRetailV2alphaImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GoogleCloudRetailV2alphaImageOutput)
}

// GoogleCloudRetailV2alphaImageArrayInput is an input type that accepts GoogleCloudRetailV2alphaImageArray and GoogleCloudRetailV2alphaImageArrayOutput values.
// You can construct a concrete instance of `GoogleCloudRetailV2alphaImageArrayInput` via:
//
//          GoogleCloudRetailV2alphaImageArray{ GoogleCloudRetailV2alphaImageArgs{...} }
type GoogleCloudRetailV2alphaImageArrayInput interface {
	pulumi.Input

	ToGoogleCloudRetailV2alphaImageArrayOutput() GoogleCloudRetailV2alphaImageArrayOutput
	ToGoogleCloudRetailV2alphaImageArrayOutputWithContext(context.Context) GoogleCloudRetailV2alphaImageArrayOutput
}

type GoogleCloudRetailV2alphaImageArray []GoogleCloudRetailV2alphaImageInput

func (GoogleCloudRetailV2alphaImageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GoogleCloudRetailV2alphaImage)(nil)).Elem()
}

func (i GoogleCloudRetailV2alphaImageArray) ToGoogleCloudRetailV2alphaImageArrayOutput() GoogleCloudRetailV2alphaImageArrayOutput {
	return i.ToGoogleCloudRetailV2alphaImageArrayOutputWithContext(context.Background())
}

func (i GoogleCloudRetailV2alphaImageArray) ToGoogleCloudRetailV2alphaImageArrayOutputWithContext(ctx context.Context) GoogleCloudRetailV2alphaImageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GoogleCloudRetailV2alphaImageArrayOutput)
}

// Product thumbnail/detail image.
type GoogleCloudRetailV2alphaImageOutput struct{ *pulumi.OutputState }

func (GoogleCloudRetailV2alphaImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleCloudRetailV2alphaImage)(nil)).Elem()
}

func (o GoogleCloudRetailV2alphaImageOutput) ToGoogleCloudRetailV2alphaImageOutput() GoogleCloudRetailV2alphaImageOutput {
	return o
}

func (o GoogleCloudRetailV2alphaImageOutput) ToGoogleCloudRetailV2alphaImageOutputWithContext(ctx context.Context) GoogleCloudRetailV2alphaImageOutput {
	return o
}

// Height of the image in number of pixels. This field must be nonnegative. Otherwise, an INVALID_ARGUMENT error is returned.
func (o GoogleCloudRetailV2alphaImageOutput) Height() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GoogleCloudRetailV2alphaImage) *int { return v.Height }).(pulumi.IntPtrOutput)
}

// URI of the image. This field must be a valid UTF-8 encoded URI with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [image_link](https://support.google.com/merchants/answer/6324350). Schema.org property [Product.image](https://schema.org/image).
func (o GoogleCloudRetailV2alphaImageOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v GoogleCloudRetailV2alphaImage) string { return v.Uri }).(pulumi.StringOutput)
}

// Width of the image in number of pixels. This field must be nonnegative. Otherwise, an INVALID_ARGUMENT error is returned.
func (o GoogleCloudRetailV2alphaImageOutput) Width() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GoogleCloudRetailV2alphaImage) *int { return v.Width }).(pulumi.IntPtrOutput)
}

type GoogleCloudRetailV2alphaImageArrayOutput struct{ *pulumi.OutputState }

func (GoogleCloudRetailV2alphaImageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GoogleCloudRetailV2alphaImage)(nil)).Elem()
}

func (o GoogleCloudRetailV2alphaImageArrayOutput) ToGoogleCloudRetailV2alphaImageArrayOutput() GoogleCloudRetailV2alphaImageArrayOutput {
	return o
}

func (o GoogleCloudRetailV2alphaImageArrayOutput) ToGoogleCloudRetailV2alphaImageArrayOutputWithContext(ctx context.Context) GoogleCloudRetailV2alphaImageArrayOutput {
	return o
}

func (o GoogleCloudRetailV2alphaImageArrayOutput) Index(i pulumi.IntInput) GoogleCloudRetailV2alphaImageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GoogleCloudRetailV2alphaImage {
		return vs[0].([]GoogleCloudRetailV2alphaImage)[vs[1].(int)]
	}).(GoogleCloudRetailV2alphaImageOutput)
}

// Product thumbnail/detail image.
type GoogleCloudRetailV2alphaImageResponse struct {
	// Height of the image in number of pixels. This field must be nonnegative. Otherwise, an INVALID_ARGUMENT error is returned.
	Height int `pulumi:"height"`
	// URI of the image. This field must be a valid UTF-8 encoded URI with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [image_link](https://support.google.com/merchants/answer/6324350). Schema.org property [Product.image](https://schema.org/image).
	Uri string `pulumi:"uri"`
	// Width of the image in number of pixels. This field must be nonnegative. Otherwise, an INVALID_ARGUMENT error is returned.
	Width int `pulumi:"width"`
}

// GoogleCloudRetailV2alphaImageResponseInput is an input type that accepts GoogleCloudRetailV2alphaImageResponseArgs and GoogleCloudRetailV2alphaImageResponseOutput values.
// You can construct a concrete instance of `GoogleCloudRetailV2alphaImageResponseInput` via:
//
//          GoogleCloudRetailV2alphaImageResponseArgs{...}
type GoogleCloudRetailV2alphaImageResponseInput interface {
	pulumi.Input

	ToGoogleCloudRetailV2alphaImageResponseOutput() GoogleCloudRetailV2alphaImageResponseOutput
	ToGoogleCloudRetailV2alphaImageResponseOutputWithContext(context.Context) GoogleCloudRetailV2alphaImageResponseOutput
}

// Product thumbnail/detail image.
type GoogleCloudRetailV2alphaImageResponseArgs struct {
	// Height of the image in number of pixels. This field must be nonnegative. Otherwise, an INVALID_ARGUMENT error is returned.
	Height pulumi.IntInput `pulumi:"height"`
	// URI of the image. This field must be a valid UTF-8 encoded URI with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [image_link](https://support.google.com/merchants/answer/6324350). Schema.org property [Product.image](https://schema.org/image).
	Uri pulumi.StringInput `pulumi:"uri"`
	// Width of the image in number of pixels. This field must be nonnegative. Otherwise, an INVALID_ARGUMENT error is returned.
	Width pulumi.IntInput `pulumi:"width"`
}

func (GoogleCloudRetailV2alphaImageResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleCloudRetailV2alphaImageResponse)(nil)).Elem()
}

func (i GoogleCloudRetailV2alphaImageResponseArgs) ToGoogleCloudRetailV2alphaImageResponseOutput() GoogleCloudRetailV2alphaImageResponseOutput {
	return i.ToGoogleCloudRetailV2alphaImageResponseOutputWithContext(context.Background())
}

func (i GoogleCloudRetailV2alphaImageResponseArgs) ToGoogleCloudRetailV2alphaImageResponseOutputWithContext(ctx context.Context) GoogleCloudRetailV2alphaImageResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GoogleCloudRetailV2alphaImageResponseOutput)
}

// GoogleCloudRetailV2alphaImageResponseArrayInput is an input type that accepts GoogleCloudRetailV2alphaImageResponseArray and GoogleCloudRetailV2alphaImageResponseArrayOutput values.
// You can construct a concrete instance of `GoogleCloudRetailV2alphaImageResponseArrayInput` via:
//
//          GoogleCloudRetailV2alphaImageResponseArray{ GoogleCloudRetailV2alphaImageResponseArgs{...} }
type GoogleCloudRetailV2alphaImageResponseArrayInput interface {
	pulumi.Input

	ToGoogleCloudRetailV2alphaImageResponseArrayOutput() GoogleCloudRetailV2alphaImageResponseArrayOutput
	ToGoogleCloudRetailV2alphaImageResponseArrayOutputWithContext(context.Context) GoogleCloudRetailV2alphaImageResponseArrayOutput
}

type GoogleCloudRetailV2alphaImageResponseArray []GoogleCloudRetailV2alphaImageResponseInput

func (GoogleCloudRetailV2alphaImageResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GoogleCloudRetailV2alphaImageResponse)(nil)).Elem()
}

func (i GoogleCloudRetailV2alphaImageResponseArray) ToGoogleCloudRetailV2alphaImageResponseArrayOutput() GoogleCloudRetailV2alphaImageResponseArrayOutput {
	return i.ToGoogleCloudRetailV2alphaImageResponseArrayOutputWithContext(context.Background())
}

func (i GoogleCloudRetailV2alphaImageResponseArray) ToGoogleCloudRetailV2alphaImageResponseArrayOutputWithContext(ctx context.Context) GoogleCloudRetailV2alphaImageResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GoogleCloudRetailV2alphaImageResponseArrayOutput)
}

// Product thumbnail/detail image.
type GoogleCloudRetailV2alphaImageResponseOutput struct{ *pulumi.OutputState }

func (GoogleCloudRetailV2alphaImageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleCloudRetailV2alphaImageResponse)(nil)).Elem()
}

func (o GoogleCloudRetailV2alphaImageResponseOutput) ToGoogleCloudRetailV2alphaImageResponseOutput() GoogleCloudRetailV2alphaImageResponseOutput {
	return o
}

func (o GoogleCloudRetailV2alphaImageResponseOutput) ToGoogleCloudRetailV2alphaImageResponseOutputWithContext(ctx context.Context) GoogleCloudRetailV2alphaImageResponseOutput {
	return o
}

// Height of the image in number of pixels. This field must be nonnegative. Otherwise, an INVALID_ARGUMENT error is returned.
func (o GoogleCloudRetailV2alphaImageResponseOutput) Height() pulumi.IntOutput {
	return o.ApplyT(func(v GoogleCloudRetailV2alphaImageResponse) int { return v.Height }).(pulumi.IntOutput)
}

// URI of the image. This field must be a valid UTF-8 encoded URI with a length limit of 5,000 characters. Otherwise, an INVALID_ARGUMENT error is returned. Google Merchant Center property [image_link](https://support.google.com/merchants/answer/6324350). Schema.org property [Product.image](https://schema.org/image).
func (o GoogleCloudRetailV2alphaImageResponseOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v GoogleCloudRetailV2alphaImageResponse) string { return v.Uri }).(pulumi.StringOutput)
}

// Width of the image in number of pixels. This field must be nonnegative. Otherwise, an INVALID_ARGUMENT error is returned.
func (o GoogleCloudRetailV2alphaImageResponseOutput) Width() pulumi.IntOutput {
	return o.ApplyT(func(v GoogleCloudRetailV2alphaImageResponse) int { return v.Width }).(pulumi.IntOutput)
}

type GoogleCloudRetailV2alphaImageResponseArrayOutput struct{ *pulumi.OutputState }

func (GoogleCloudRetailV2alphaImageResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GoogleCloudRetailV2alphaImageResponse)(nil)).Elem()
}

func (o GoogleCloudRetailV2alphaImageResponseArrayOutput) ToGoogleCloudRetailV2alphaImageResponseArrayOutput() GoogleCloudRetailV2alphaImageResponseArrayOutput {
	return o
}

func (o GoogleCloudRetailV2alphaImageResponseArrayOutput) ToGoogleCloudRetailV2alphaImageResponseArrayOutputWithContext(ctx context.Context) GoogleCloudRetailV2alphaImageResponseArrayOutput {
	return o
}

func (o GoogleCloudRetailV2alphaImageResponseArrayOutput) Index(i pulumi.IntInput) GoogleCloudRetailV2alphaImageResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GoogleCloudRetailV2alphaImageResponse {
		return vs[0].([]GoogleCloudRetailV2alphaImageResponse)[vs[1].(int)]
	}).(GoogleCloudRetailV2alphaImageResponseOutput)
}

// The price information of a Product.
type GoogleCloudRetailV2alphaPriceInfo struct {
	// The costs associated with the sale of a particular product. Used for gross profit reporting. * Profit = price - cost Google Merchant Center property [cost_of_goods_sold](https://support.google.com/merchants/answer/9017895).
	Cost *float64 `pulumi:"cost"`
	// The 3-letter currency code defined in [ISO 4217](https://www.iso.org/iso-4217-currency-codes.html). If this field is an unrecognizable currency code, an INVALID_ARGUMENT error is returned.
	CurrencyCode *string `pulumi:"currencyCode"`
	// Price of the product without any discount. If zero, by default set to be the price.
	OriginalPrice *float64 `pulumi:"originalPrice"`
	// Price of the product. Google Merchant Center property [price](https://support.google.com/merchants/answer/6324371). Schema.org property [Offer.priceSpecification](https://schema.org/priceSpecification).
	Price *float64 `pulumi:"price"`
}

// GoogleCloudRetailV2alphaPriceInfoInput is an input type that accepts GoogleCloudRetailV2alphaPriceInfoArgs and GoogleCloudRetailV2alphaPriceInfoOutput values.
// You can construct a concrete instance of `GoogleCloudRetailV2alphaPriceInfoInput` via:
//
//          GoogleCloudRetailV2alphaPriceInfoArgs{...}
type GoogleCloudRetailV2alphaPriceInfoInput interface {
	pulumi.Input

	ToGoogleCloudRetailV2alphaPriceInfoOutput() GoogleCloudRetailV2alphaPriceInfoOutput
	ToGoogleCloudRetailV2alphaPriceInfoOutputWithContext(context.Context) GoogleCloudRetailV2alphaPriceInfoOutput
}

// The price information of a Product.
type GoogleCloudRetailV2alphaPriceInfoArgs struct {
	// The costs associated with the sale of a particular product. Used for gross profit reporting. * Profit = price - cost Google Merchant Center property [cost_of_goods_sold](https://support.google.com/merchants/answer/9017895).
	Cost pulumi.Float64PtrInput `pulumi:"cost"`
	// The 3-letter currency code defined in [ISO 4217](https://www.iso.org/iso-4217-currency-codes.html). If this field is an unrecognizable currency code, an INVALID_ARGUMENT error is returned.
	CurrencyCode pulumi.StringPtrInput `pulumi:"currencyCode"`
	// Price of the product without any discount. If zero, by default set to be the price.
	OriginalPrice pulumi.Float64PtrInput `pulumi:"originalPrice"`
	// Price of the product. Google Merchant Center property [price](https://support.google.com/merchants/answer/6324371). Schema.org property [Offer.priceSpecification](https://schema.org/priceSpecification).
	Price pulumi.Float64PtrInput `pulumi:"price"`
}

func (GoogleCloudRetailV2alphaPriceInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleCloudRetailV2alphaPriceInfo)(nil)).Elem()
}

func (i GoogleCloudRetailV2alphaPriceInfoArgs) ToGoogleCloudRetailV2alphaPriceInfoOutput() GoogleCloudRetailV2alphaPriceInfoOutput {
	return i.ToGoogleCloudRetailV2alphaPriceInfoOutputWithContext(context.Background())
}

func (i GoogleCloudRetailV2alphaPriceInfoArgs) ToGoogleCloudRetailV2alphaPriceInfoOutputWithContext(ctx context.Context) GoogleCloudRetailV2alphaPriceInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GoogleCloudRetailV2alphaPriceInfoOutput)
}

func (i GoogleCloudRetailV2alphaPriceInfoArgs) ToGoogleCloudRetailV2alphaPriceInfoPtrOutput() GoogleCloudRetailV2alphaPriceInfoPtrOutput {
	return i.ToGoogleCloudRetailV2alphaPriceInfoPtrOutputWithContext(context.Background())
}

func (i GoogleCloudRetailV2alphaPriceInfoArgs) ToGoogleCloudRetailV2alphaPriceInfoPtrOutputWithContext(ctx context.Context) GoogleCloudRetailV2alphaPriceInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GoogleCloudRetailV2alphaPriceInfoOutput).ToGoogleCloudRetailV2alphaPriceInfoPtrOutputWithContext(ctx)
}

// GoogleCloudRetailV2alphaPriceInfoPtrInput is an input type that accepts GoogleCloudRetailV2alphaPriceInfoArgs, GoogleCloudRetailV2alphaPriceInfoPtr and GoogleCloudRetailV2alphaPriceInfoPtrOutput values.
// You can construct a concrete instance of `GoogleCloudRetailV2alphaPriceInfoPtrInput` via:
//
//          GoogleCloudRetailV2alphaPriceInfoArgs{...}
//
//  or:
//
//          nil
type GoogleCloudRetailV2alphaPriceInfoPtrInput interface {
	pulumi.Input

	ToGoogleCloudRetailV2alphaPriceInfoPtrOutput() GoogleCloudRetailV2alphaPriceInfoPtrOutput
	ToGoogleCloudRetailV2alphaPriceInfoPtrOutputWithContext(context.Context) GoogleCloudRetailV2alphaPriceInfoPtrOutput
}

type googleCloudRetailV2alphaPriceInfoPtrType GoogleCloudRetailV2alphaPriceInfoArgs

func GoogleCloudRetailV2alphaPriceInfoPtr(v *GoogleCloudRetailV2alphaPriceInfoArgs) GoogleCloudRetailV2alphaPriceInfoPtrInput {
	return (*googleCloudRetailV2alphaPriceInfoPtrType)(v)
}

func (*googleCloudRetailV2alphaPriceInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GoogleCloudRetailV2alphaPriceInfo)(nil)).Elem()
}

func (i *googleCloudRetailV2alphaPriceInfoPtrType) ToGoogleCloudRetailV2alphaPriceInfoPtrOutput() GoogleCloudRetailV2alphaPriceInfoPtrOutput {
	return i.ToGoogleCloudRetailV2alphaPriceInfoPtrOutputWithContext(context.Background())
}

func (i *googleCloudRetailV2alphaPriceInfoPtrType) ToGoogleCloudRetailV2alphaPriceInfoPtrOutputWithContext(ctx context.Context) GoogleCloudRetailV2alphaPriceInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GoogleCloudRetailV2alphaPriceInfoPtrOutput)
}

// The price information of a Product.
type GoogleCloudRetailV2alphaPriceInfoOutput struct{ *pulumi.OutputState }

func (GoogleCloudRetailV2alphaPriceInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleCloudRetailV2alphaPriceInfo)(nil)).Elem()
}

func (o GoogleCloudRetailV2alphaPriceInfoOutput) ToGoogleCloudRetailV2alphaPriceInfoOutput() GoogleCloudRetailV2alphaPriceInfoOutput {
	return o
}

func (o GoogleCloudRetailV2alphaPriceInfoOutput) ToGoogleCloudRetailV2alphaPriceInfoOutputWithContext(ctx context.Context) GoogleCloudRetailV2alphaPriceInfoOutput {
	return o
}

func (o GoogleCloudRetailV2alphaPriceInfoOutput) ToGoogleCloudRetailV2alphaPriceInfoPtrOutput() GoogleCloudRetailV2alphaPriceInfoPtrOutput {
	return o.ToGoogleCloudRetailV2alphaPriceInfoPtrOutputWithContext(context.Background())
}

func (o GoogleCloudRetailV2alphaPriceInfoOutput) ToGoogleCloudRetailV2alphaPriceInfoPtrOutputWithContext(ctx context.Context) GoogleCloudRetailV2alphaPriceInfoPtrOutput {
	return o.ApplyT(func(v GoogleCloudRetailV2alphaPriceInfo) *GoogleCloudRetailV2alphaPriceInfo {
		return &v
	}).(GoogleCloudRetailV2alphaPriceInfoPtrOutput)
}

// The costs associated with the sale of a particular product. Used for gross profit reporting. * Profit = price - cost Google Merchant Center property [cost_of_goods_sold](https://support.google.com/merchants/answer/9017895).
func (o GoogleCloudRetailV2alphaPriceInfoOutput) Cost() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v GoogleCloudRetailV2alphaPriceInfo) *float64 { return v.Cost }).(pulumi.Float64PtrOutput)
}

// The 3-letter currency code defined in [ISO 4217](https://www.iso.org/iso-4217-currency-codes.html). If this field is an unrecognizable currency code, an INVALID_ARGUMENT error is returned.
func (o GoogleCloudRetailV2alphaPriceInfoOutput) CurrencyCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GoogleCloudRetailV2alphaPriceInfo) *string { return v.CurrencyCode }).(pulumi.StringPtrOutput)
}

// Price of the product without any discount. If zero, by default set to be the price.
func (o GoogleCloudRetailV2alphaPriceInfoOutput) OriginalPrice() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v GoogleCloudRetailV2alphaPriceInfo) *float64 { return v.OriginalPrice }).(pulumi.Float64PtrOutput)
}

// Price of the product. Google Merchant Center property [price](https://support.google.com/merchants/answer/6324371). Schema.org property [Offer.priceSpecification](https://schema.org/priceSpecification).
func (o GoogleCloudRetailV2alphaPriceInfoOutput) Price() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v GoogleCloudRetailV2alphaPriceInfo) *float64 { return v.Price }).(pulumi.Float64PtrOutput)
}

type GoogleCloudRetailV2alphaPriceInfoPtrOutput struct{ *pulumi.OutputState }

func (GoogleCloudRetailV2alphaPriceInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GoogleCloudRetailV2alphaPriceInfo)(nil)).Elem()
}

func (o GoogleCloudRetailV2alphaPriceInfoPtrOutput) ToGoogleCloudRetailV2alphaPriceInfoPtrOutput() GoogleCloudRetailV2alphaPriceInfoPtrOutput {
	return o
}

func (o GoogleCloudRetailV2alphaPriceInfoPtrOutput) ToGoogleCloudRetailV2alphaPriceInfoPtrOutputWithContext(ctx context.Context) GoogleCloudRetailV2alphaPriceInfoPtrOutput {
	return o
}

func (o GoogleCloudRetailV2alphaPriceInfoPtrOutput) Elem() GoogleCloudRetailV2alphaPriceInfoOutput {
	return o.ApplyT(func(v *GoogleCloudRetailV2alphaPriceInfo) GoogleCloudRetailV2alphaPriceInfo { return *v }).(GoogleCloudRetailV2alphaPriceInfoOutput)
}

// The costs associated with the sale of a particular product. Used for gross profit reporting. * Profit = price - cost Google Merchant Center property [cost_of_goods_sold](https://support.google.com/merchants/answer/9017895).
func (o GoogleCloudRetailV2alphaPriceInfoPtrOutput) Cost() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *GoogleCloudRetailV2alphaPriceInfo) *float64 {
		if v == nil {
			return nil
		}
		return v.Cost
	}).(pulumi.Float64PtrOutput)
}

// The 3-letter currency code defined in [ISO 4217](https://www.iso.org/iso-4217-currency-codes.html). If this field is an unrecognizable currency code, an INVALID_ARGUMENT error is returned.
func (o GoogleCloudRetailV2alphaPriceInfoPtrOutput) CurrencyCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GoogleCloudRetailV2alphaPriceInfo) *string {
		if v == nil {
			return nil
		}
		return v.CurrencyCode
	}).(pulumi.StringPtrOutput)
}

// Price of the product without any discount. If zero, by default set to be the price.
func (o GoogleCloudRetailV2alphaPriceInfoPtrOutput) OriginalPrice() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *GoogleCloudRetailV2alphaPriceInfo) *float64 {
		if v == nil {
			return nil
		}
		return v.OriginalPrice
	}).(pulumi.Float64PtrOutput)
}

// Price of the product. Google Merchant Center property [price](https://support.google.com/merchants/answer/6324371). Schema.org property [Offer.priceSpecification](https://schema.org/priceSpecification).
func (o GoogleCloudRetailV2alphaPriceInfoPtrOutput) Price() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *GoogleCloudRetailV2alphaPriceInfo) *float64 {
		if v == nil {
			return nil
		}
		return v.Price
	}).(pulumi.Float64PtrOutput)
}

// The price information of a Product.
type GoogleCloudRetailV2alphaPriceInfoResponse struct {
	// The costs associated with the sale of a particular product. Used for gross profit reporting. * Profit = price - cost Google Merchant Center property [cost_of_goods_sold](https://support.google.com/merchants/answer/9017895).
	Cost float64 `pulumi:"cost"`
	// The 3-letter currency code defined in [ISO 4217](https://www.iso.org/iso-4217-currency-codes.html). If this field is an unrecognizable currency code, an INVALID_ARGUMENT error is returned.
	CurrencyCode string `pulumi:"currencyCode"`
	// Price of the product without any discount. If zero, by default set to be the price.
	OriginalPrice float64 `pulumi:"originalPrice"`
	// Price of the product. Google Merchant Center property [price](https://support.google.com/merchants/answer/6324371). Schema.org property [Offer.priceSpecification](https://schema.org/priceSpecification).
	Price float64 `pulumi:"price"`
}

// GoogleCloudRetailV2alphaPriceInfoResponseInput is an input type that accepts GoogleCloudRetailV2alphaPriceInfoResponseArgs and GoogleCloudRetailV2alphaPriceInfoResponseOutput values.
// You can construct a concrete instance of `GoogleCloudRetailV2alphaPriceInfoResponseInput` via:
//
//          GoogleCloudRetailV2alphaPriceInfoResponseArgs{...}
type GoogleCloudRetailV2alphaPriceInfoResponseInput interface {
	pulumi.Input

	ToGoogleCloudRetailV2alphaPriceInfoResponseOutput() GoogleCloudRetailV2alphaPriceInfoResponseOutput
	ToGoogleCloudRetailV2alphaPriceInfoResponseOutputWithContext(context.Context) GoogleCloudRetailV2alphaPriceInfoResponseOutput
}

// The price information of a Product.
type GoogleCloudRetailV2alphaPriceInfoResponseArgs struct {
	// The costs associated with the sale of a particular product. Used for gross profit reporting. * Profit = price - cost Google Merchant Center property [cost_of_goods_sold](https://support.google.com/merchants/answer/9017895).
	Cost pulumi.Float64Input `pulumi:"cost"`
	// The 3-letter currency code defined in [ISO 4217](https://www.iso.org/iso-4217-currency-codes.html). If this field is an unrecognizable currency code, an INVALID_ARGUMENT error is returned.
	CurrencyCode pulumi.StringInput `pulumi:"currencyCode"`
	// Price of the product without any discount. If zero, by default set to be the price.
	OriginalPrice pulumi.Float64Input `pulumi:"originalPrice"`
	// Price of the product. Google Merchant Center property [price](https://support.google.com/merchants/answer/6324371). Schema.org property [Offer.priceSpecification](https://schema.org/priceSpecification).
	Price pulumi.Float64Input `pulumi:"price"`
}

func (GoogleCloudRetailV2alphaPriceInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleCloudRetailV2alphaPriceInfoResponse)(nil)).Elem()
}

func (i GoogleCloudRetailV2alphaPriceInfoResponseArgs) ToGoogleCloudRetailV2alphaPriceInfoResponseOutput() GoogleCloudRetailV2alphaPriceInfoResponseOutput {
	return i.ToGoogleCloudRetailV2alphaPriceInfoResponseOutputWithContext(context.Background())
}

func (i GoogleCloudRetailV2alphaPriceInfoResponseArgs) ToGoogleCloudRetailV2alphaPriceInfoResponseOutputWithContext(ctx context.Context) GoogleCloudRetailV2alphaPriceInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GoogleCloudRetailV2alphaPriceInfoResponseOutput)
}

func (i GoogleCloudRetailV2alphaPriceInfoResponseArgs) ToGoogleCloudRetailV2alphaPriceInfoResponsePtrOutput() GoogleCloudRetailV2alphaPriceInfoResponsePtrOutput {
	return i.ToGoogleCloudRetailV2alphaPriceInfoResponsePtrOutputWithContext(context.Background())
}

func (i GoogleCloudRetailV2alphaPriceInfoResponseArgs) ToGoogleCloudRetailV2alphaPriceInfoResponsePtrOutputWithContext(ctx context.Context) GoogleCloudRetailV2alphaPriceInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GoogleCloudRetailV2alphaPriceInfoResponseOutput).ToGoogleCloudRetailV2alphaPriceInfoResponsePtrOutputWithContext(ctx)
}

// GoogleCloudRetailV2alphaPriceInfoResponsePtrInput is an input type that accepts GoogleCloudRetailV2alphaPriceInfoResponseArgs, GoogleCloudRetailV2alphaPriceInfoResponsePtr and GoogleCloudRetailV2alphaPriceInfoResponsePtrOutput values.
// You can construct a concrete instance of `GoogleCloudRetailV2alphaPriceInfoResponsePtrInput` via:
//
//          GoogleCloudRetailV2alphaPriceInfoResponseArgs{...}
//
//  or:
//
//          nil
type GoogleCloudRetailV2alphaPriceInfoResponsePtrInput interface {
	pulumi.Input

	ToGoogleCloudRetailV2alphaPriceInfoResponsePtrOutput() GoogleCloudRetailV2alphaPriceInfoResponsePtrOutput
	ToGoogleCloudRetailV2alphaPriceInfoResponsePtrOutputWithContext(context.Context) GoogleCloudRetailV2alphaPriceInfoResponsePtrOutput
}

type googleCloudRetailV2alphaPriceInfoResponsePtrType GoogleCloudRetailV2alphaPriceInfoResponseArgs

func GoogleCloudRetailV2alphaPriceInfoResponsePtr(v *GoogleCloudRetailV2alphaPriceInfoResponseArgs) GoogleCloudRetailV2alphaPriceInfoResponsePtrInput {
	return (*googleCloudRetailV2alphaPriceInfoResponsePtrType)(v)
}

func (*googleCloudRetailV2alphaPriceInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GoogleCloudRetailV2alphaPriceInfoResponse)(nil)).Elem()
}

func (i *googleCloudRetailV2alphaPriceInfoResponsePtrType) ToGoogleCloudRetailV2alphaPriceInfoResponsePtrOutput() GoogleCloudRetailV2alphaPriceInfoResponsePtrOutput {
	return i.ToGoogleCloudRetailV2alphaPriceInfoResponsePtrOutputWithContext(context.Background())
}

func (i *googleCloudRetailV2alphaPriceInfoResponsePtrType) ToGoogleCloudRetailV2alphaPriceInfoResponsePtrOutputWithContext(ctx context.Context) GoogleCloudRetailV2alphaPriceInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GoogleCloudRetailV2alphaPriceInfoResponsePtrOutput)
}

// The price information of a Product.
type GoogleCloudRetailV2alphaPriceInfoResponseOutput struct{ *pulumi.OutputState }

func (GoogleCloudRetailV2alphaPriceInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GoogleCloudRetailV2alphaPriceInfoResponse)(nil)).Elem()
}

func (o GoogleCloudRetailV2alphaPriceInfoResponseOutput) ToGoogleCloudRetailV2alphaPriceInfoResponseOutput() GoogleCloudRetailV2alphaPriceInfoResponseOutput {
	return o
}

func (o GoogleCloudRetailV2alphaPriceInfoResponseOutput) ToGoogleCloudRetailV2alphaPriceInfoResponseOutputWithContext(ctx context.Context) GoogleCloudRetailV2alphaPriceInfoResponseOutput {
	return o
}

func (o GoogleCloudRetailV2alphaPriceInfoResponseOutput) ToGoogleCloudRetailV2alphaPriceInfoResponsePtrOutput() GoogleCloudRetailV2alphaPriceInfoResponsePtrOutput {
	return o.ToGoogleCloudRetailV2alphaPriceInfoResponsePtrOutputWithContext(context.Background())
}

func (o GoogleCloudRetailV2alphaPriceInfoResponseOutput) ToGoogleCloudRetailV2alphaPriceInfoResponsePtrOutputWithContext(ctx context.Context) GoogleCloudRetailV2alphaPriceInfoResponsePtrOutput {
	return o.ApplyT(func(v GoogleCloudRetailV2alphaPriceInfoResponse) *GoogleCloudRetailV2alphaPriceInfoResponse {
		return &v
	}).(GoogleCloudRetailV2alphaPriceInfoResponsePtrOutput)
}

// The costs associated with the sale of a particular product. Used for gross profit reporting. * Profit = price - cost Google Merchant Center property [cost_of_goods_sold](https://support.google.com/merchants/answer/9017895).
func (o GoogleCloudRetailV2alphaPriceInfoResponseOutput) Cost() pulumi.Float64Output {
	return o.ApplyT(func(v GoogleCloudRetailV2alphaPriceInfoResponse) float64 { return v.Cost }).(pulumi.Float64Output)
}

// The 3-letter currency code defined in [ISO 4217](https://www.iso.org/iso-4217-currency-codes.html). If this field is an unrecognizable currency code, an INVALID_ARGUMENT error is returned.
func (o GoogleCloudRetailV2alphaPriceInfoResponseOutput) CurrencyCode() pulumi.StringOutput {
	return o.ApplyT(func(v GoogleCloudRetailV2alphaPriceInfoResponse) string { return v.CurrencyCode }).(pulumi.StringOutput)
}

// Price of the product without any discount. If zero, by default set to be the price.
func (o GoogleCloudRetailV2alphaPriceInfoResponseOutput) OriginalPrice() pulumi.Float64Output {
	return o.ApplyT(func(v GoogleCloudRetailV2alphaPriceInfoResponse) float64 { return v.OriginalPrice }).(pulumi.Float64Output)
}

// Price of the product. Google Merchant Center property [price](https://support.google.com/merchants/answer/6324371). Schema.org property [Offer.priceSpecification](https://schema.org/priceSpecification).
func (o GoogleCloudRetailV2alphaPriceInfoResponseOutput) Price() pulumi.Float64Output {
	return o.ApplyT(func(v GoogleCloudRetailV2alphaPriceInfoResponse) float64 { return v.Price }).(pulumi.Float64Output)
}

type GoogleCloudRetailV2alphaPriceInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (GoogleCloudRetailV2alphaPriceInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GoogleCloudRetailV2alphaPriceInfoResponse)(nil)).Elem()
}

func (o GoogleCloudRetailV2alphaPriceInfoResponsePtrOutput) ToGoogleCloudRetailV2alphaPriceInfoResponsePtrOutput() GoogleCloudRetailV2alphaPriceInfoResponsePtrOutput {
	return o
}

func (o GoogleCloudRetailV2alphaPriceInfoResponsePtrOutput) ToGoogleCloudRetailV2alphaPriceInfoResponsePtrOutputWithContext(ctx context.Context) GoogleCloudRetailV2alphaPriceInfoResponsePtrOutput {
	return o
}

func (o GoogleCloudRetailV2alphaPriceInfoResponsePtrOutput) Elem() GoogleCloudRetailV2alphaPriceInfoResponseOutput {
	return o.ApplyT(func(v *GoogleCloudRetailV2alphaPriceInfoResponse) GoogleCloudRetailV2alphaPriceInfoResponse {
		return *v
	}).(GoogleCloudRetailV2alphaPriceInfoResponseOutput)
}

// The costs associated with the sale of a particular product. Used for gross profit reporting. * Profit = price - cost Google Merchant Center property [cost_of_goods_sold](https://support.google.com/merchants/answer/9017895).
func (o GoogleCloudRetailV2alphaPriceInfoResponsePtrOutput) Cost() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *GoogleCloudRetailV2alphaPriceInfoResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.Cost
	}).(pulumi.Float64PtrOutput)
}

// The 3-letter currency code defined in [ISO 4217](https://www.iso.org/iso-4217-currency-codes.html). If this field is an unrecognizable currency code, an INVALID_ARGUMENT error is returned.
func (o GoogleCloudRetailV2alphaPriceInfoResponsePtrOutput) CurrencyCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GoogleCloudRetailV2alphaPriceInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CurrencyCode
	}).(pulumi.StringPtrOutput)
}

// Price of the product without any discount. If zero, by default set to be the price.
func (o GoogleCloudRetailV2alphaPriceInfoResponsePtrOutput) OriginalPrice() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *GoogleCloudRetailV2alphaPriceInfoResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.OriginalPrice
	}).(pulumi.Float64PtrOutput)
}

// Price of the product. Google Merchant Center property [price](https://support.google.com/merchants/answer/6324371). Schema.org property [Offer.priceSpecification](https://schema.org/priceSpecification).
func (o GoogleCloudRetailV2alphaPriceInfoResponsePtrOutput) Price() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *GoogleCloudRetailV2alphaPriceInfoResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.Price
	}).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GoogleCloudRetailV2alphaImageOutput{})
	pulumi.RegisterOutputType(GoogleCloudRetailV2alphaImageArrayOutput{})
	pulumi.RegisterOutputType(GoogleCloudRetailV2alphaImageResponseOutput{})
	pulumi.RegisterOutputType(GoogleCloudRetailV2alphaImageResponseArrayOutput{})
	pulumi.RegisterOutputType(GoogleCloudRetailV2alphaPriceInfoOutput{})
	pulumi.RegisterOutputType(GoogleCloudRetailV2alphaPriceInfoPtrOutput{})
	pulumi.RegisterOutputType(GoogleCloudRetailV2alphaPriceInfoResponseOutput{})
	pulumi.RegisterOutputType(GoogleCloudRetailV2alphaPriceInfoResponsePtrOutput{})
}
