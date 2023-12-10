// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pulumi_one_password_native_unofficial

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDatabase(ctx *pulumi.Context, args *GetDatabaseArgs, opts ...pulumi.InvokeOption) (*GetDatabaseResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetDatabaseResult
	err := ctx.Invoke("one-password-native-unofficial:index:GetDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetDatabaseArgs struct {
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Id *string `pulumi:"id"`
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title *string `pulumi:"title"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

type GetDatabaseResult struct {
	Alias             *string                     `pulumi:"alias"`
	Attachments       map[string]OutputAttachment `pulumi:"attachments"`
	Category          string                      `pulumi:"category"`
	ConnectionOptions *string                     `pulumi:"connectionOptions"`
	Database          *string                     `pulumi:"database"`
	Fields            map[string]OutputField      `pulumi:"fields"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Id         string                   `pulumi:"id"`
	Notes      *string                  `pulumi:"notes"`
	Password   *string                  `pulumi:"password"`
	Port       *string                  `pulumi:"port"`
	References []OutputReference        `pulumi:"references"`
	Sections   map[string]OutputSection `pulumi:"sections"`
	Server     *string                  `pulumi:"server"`
	Sid        *string                  `pulumi:"sid"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item.
	Title    string            `pulumi:"title"`
	Type     *string           `pulumi:"type"`
	Urls     []OutputUrl       `pulumi:"urls"`
	Username *string           `pulumi:"username"`
	Vault    map[string]string `pulumi:"vault"`
}

func GetDatabaseOutput(ctx *pulumi.Context, args GetDatabaseOutputArgs, opts ...pulumi.InvokeOption) GetDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDatabaseResult, error) {
			args := v.(GetDatabaseArgs)
			r, err := GetDatabase(ctx, &args, opts...)
			var s GetDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDatabaseResultOutput)
}

type GetDatabaseOutputArgs struct {
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title pulumi.StringPtrInput `pulumi:"title"`
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput `pulumi:"vault"`
}

func (GetDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabaseArgs)(nil)).Elem()
}

type GetDatabaseResultOutput struct{ *pulumi.OutputState }

func (GetDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabaseResult)(nil)).Elem()
}

func (o GetDatabaseResultOutput) ToGetDatabaseResultOutput() GetDatabaseResultOutput {
	return o
}

func (o GetDatabaseResultOutput) ToGetDatabaseResultOutputWithContext(ctx context.Context) GetDatabaseResultOutput {
	return o
}

func (o GetDatabaseResultOutput) Alias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDatabaseResult) *string { return v.Alias }).(pulumi.StringPtrOutput)
}

func (o GetDatabaseResultOutput) Attachments() OutputAttachmentMapOutput {
	return o.ApplyT(func(v GetDatabaseResult) map[string]OutputAttachment { return v.Attachments }).(OutputAttachmentMapOutput)
}

func (o GetDatabaseResultOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseResult) string { return v.Category }).(pulumi.StringOutput)
}

func (o GetDatabaseResultOutput) ConnectionOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDatabaseResult) *string { return v.ConnectionOptions }).(pulumi.StringPtrOutput)
}

func (o GetDatabaseResultOutput) Database() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDatabaseResult) *string { return v.Database }).(pulumi.StringPtrOutput)
}

func (o GetDatabaseResultOutput) Fields() OutputFieldMapOutput {
	return o.ApplyT(func(v GetDatabaseResult) map[string]OutputField { return v.Fields }).(OutputFieldMapOutput)
}

// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
func (o GetDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDatabaseResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDatabaseResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o GetDatabaseResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDatabaseResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o GetDatabaseResultOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDatabaseResult) *string { return v.Port }).(pulumi.StringPtrOutput)
}

func (o GetDatabaseResultOutput) References() OutputReferenceArrayOutput {
	return o.ApplyT(func(v GetDatabaseResult) []OutputReference { return v.References }).(OutputReferenceArrayOutput)
}

func (o GetDatabaseResultOutput) Sections() OutputSectionMapOutput {
	return o.ApplyT(func(v GetDatabaseResult) map[string]OutputSection { return v.Sections }).(OutputSectionMapOutput)
}

func (o GetDatabaseResultOutput) Server() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDatabaseResult) *string { return v.Server }).(pulumi.StringPtrOutput)
}

func (o GetDatabaseResultOutput) Sid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDatabaseResult) *string { return v.Sid }).(pulumi.StringPtrOutput)
}

// An array of strings of the tags assigned to the item.
func (o GetDatabaseResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDatabaseResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The title of the item.
func (o GetDatabaseResultOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseResult) string { return v.Title }).(pulumi.StringOutput)
}

func (o GetDatabaseResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDatabaseResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o GetDatabaseResultOutput) Urls() OutputUrlArrayOutput {
	return o.ApplyT(func(v GetDatabaseResult) []OutputUrl { return v.Urls }).(OutputUrlArrayOutput)
}

func (o GetDatabaseResultOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDatabaseResult) *string { return v.Username }).(pulumi.StringPtrOutput)
}

func (o GetDatabaseResultOutput) Vault() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetDatabaseResult) map[string]string { return v.Vault }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDatabaseResultOutput{})
}
