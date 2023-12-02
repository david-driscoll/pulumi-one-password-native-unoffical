// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pulumi_one_password_native_unoffical

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetAttachment(ctx *pulumi.Context, args *GetAttachmentArgs, opts ...pulumi.InvokeOption) (*GetAttachmentResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetAttachmentResult
	err := ctx.Invoke("one-password-native-unoffical:index:GetAttachment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetAttachmentArgs struct {
	// The 1Password secret reference path to the item.  eg: op://vault/item/[section]/file
	Reference string `pulumi:"reference"`
}

// The resolved reference value
type GetAttachmentResult struct {
	Value *string `pulumi:"value"`
}

func GetAttachmentOutput(ctx *pulumi.Context, args GetAttachmentOutputArgs, opts ...pulumi.InvokeOption) GetAttachmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAttachmentResult, error) {
			args := v.(GetAttachmentArgs)
			r, err := GetAttachment(ctx, &args, opts...)
			var s GetAttachmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAttachmentResultOutput)
}

type GetAttachmentOutputArgs struct {
	// The 1Password secret reference path to the item.  eg: op://vault/item/[section]/file
	Reference pulumi.StringInput `pulumi:"reference"`
}

func (GetAttachmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAttachmentArgs)(nil)).Elem()
}

// The resolved reference value
type GetAttachmentResultOutput struct{ *pulumi.OutputState }

func (GetAttachmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAttachmentResult)(nil)).Elem()
}

func (o GetAttachmentResultOutput) ToGetAttachmentResultOutput() GetAttachmentResultOutput {
	return o
}

func (o GetAttachmentResultOutput) ToGetAttachmentResultOutputWithContext(ctx context.Context) GetAttachmentResultOutput {
	return o
}

func (o GetAttachmentResultOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAttachmentResult) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAttachmentResultOutput{})
}
