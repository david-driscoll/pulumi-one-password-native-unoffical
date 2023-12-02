// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pulumi_one_password_native_unoffical

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSecretReference(ctx *pulumi.Context, args *GetSecretReferenceArgs, opts ...pulumi.InvokeOption) (*GetSecretReferenceResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSecretReferenceResult
	err := ctx.Invoke("one-password-native-unoffical:index:GetSecretReference", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetSecretReferenceArgs struct {
	// The 1Password secret reference path to the item.  eg: op://vault/item/[section]/field
	Reference string `pulumi:"reference"`
}

// The resolved reference value
type GetSecretReferenceResult struct {
	Value *string `pulumi:"value"`
}

func GetSecretReferenceOutput(ctx *pulumi.Context, args GetSecretReferenceOutputArgs, opts ...pulumi.InvokeOption) GetSecretReferenceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSecretReferenceResult, error) {
			args := v.(GetSecretReferenceArgs)
			r, err := GetSecretReference(ctx, &args, opts...)
			var s GetSecretReferenceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSecretReferenceResultOutput)
}

type GetSecretReferenceOutputArgs struct {
	// The 1Password secret reference path to the item.  eg: op://vault/item/[section]/field
	Reference pulumi.StringInput `pulumi:"reference"`
}

func (GetSecretReferenceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretReferenceArgs)(nil)).Elem()
}

// The resolved reference value
type GetSecretReferenceResultOutput struct{ *pulumi.OutputState }

func (GetSecretReferenceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSecretReferenceResult)(nil)).Elem()
}

func (o GetSecretReferenceResultOutput) ToGetSecretReferenceResultOutput() GetSecretReferenceResultOutput {
	return o
}

func (o GetSecretReferenceResultOutput) ToGetSecretReferenceResultOutputWithContext(ctx context.Context) GetSecretReferenceResultOutput {
	return o
}

func (o GetSecretReferenceResultOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSecretReferenceResult) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSecretReferenceResultOutput{})
}
