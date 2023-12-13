// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pulumi_one_password_native_unofficial

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DocumentItem struct {
	pulumi.CustomResourceState

	Attachments OutputAttachmentMapOutput `pulumi:"attachments"`
	Category    pulumi.StringOutput       `pulumi:"category"`
	Fields      OutputFieldMapOutput      `pulumi:"fields"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Id         pulumi.StringOutput        `pulumi:"id"`
	Notes      pulumi.StringPtrOutput     `pulumi:"notes"`
	References OutputReferenceArrayOutput `pulumi:"references"`
	Sections   OutputSectionMapOutput     `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The title of the item.
	Title pulumi.StringOutput    `pulumi:"title"`
	Urls  OutputUrlArrayOutput   `pulumi:"urls"`
	Vault pulumi.StringMapOutput `pulumi:"vault"`
}

// NewDocumentItem registers a new resource with the given unique name, arguments, and options.
func NewDocumentItem(ctx *pulumi.Context,
	name string, args *DocumentItemArgs, opts ...pulumi.ResourceOption) (*DocumentItem, error) {
	if args == nil {
		args = &DocumentItemArgs{}
	}

	args.Category = pulumi.StringPtr("Document")
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"attachments",
		"fields",
		"sections",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource DocumentItem
	err := ctx.RegisterResource("one-password-native-unofficial:index:DocumentItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDocumentItem gets an existing DocumentItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDocumentItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DocumentItemState, opts ...pulumi.ResourceOption) (*DocumentItem, error) {
	var resource DocumentItem
	err := ctx.ReadResource("one-password-native-unofficial:index:DocumentItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DocumentItem resources.
type documentItemState struct {
	// The UUID of the vault the item is in.
	Vault *string `pulumi:"vault"`
}

type DocumentItemState struct {
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput
}

func (DocumentItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*documentItemState)(nil)).Elem()
}

type documentItemArgs struct {
	Attachments map[string]pulumi.AssetOrArchive `pulumi:"attachments"`
	// The category of the vault the item is in.
	Category   *string                `pulumi:"category"`
	Fields     map[string]interface{} `pulumi:"fields"`
	Notes      *string                `pulumi:"notes"`
	References []string               `pulumi:"references"`
	Sections   map[string]Section     `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title *string       `pulumi:"title"`
	Urls  []interface{} `pulumi:"urls"`
	// The UUID of the vault the item is in.
	Vault *string `pulumi:"vault"`
}

// The set of arguments for constructing a DocumentItem resource.
type DocumentItemArgs struct {
	Attachments pulumi.AssetOrArchiveMapInput
	// The category of the vault the item is in.
	Category   pulumi.StringPtrInput
	Fields     pulumi.MapInput
	Notes      pulumi.StringPtrInput
	References pulumi.StringArrayInput
	Sections   SectionMapInput
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayInput
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title pulumi.StringPtrInput
	Urls  pulumi.ArrayInput
	// The UUID of the vault the item is in.
	Vault pulumi.StringPtrInput
}

func (DocumentItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*documentItemArgs)(nil)).Elem()
}

type DocumentItemInput interface {
	pulumi.Input

	ToDocumentItemOutput() DocumentItemOutput
	ToDocumentItemOutputWithContext(ctx context.Context) DocumentItemOutput
}

func (*DocumentItem) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentItem)(nil)).Elem()
}

func (i *DocumentItem) ToDocumentItemOutput() DocumentItemOutput {
	return i.ToDocumentItemOutputWithContext(context.Background())
}

func (i *DocumentItem) ToDocumentItemOutputWithContext(ctx context.Context) DocumentItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentItemOutput)
}

// DocumentItemArrayInput is an input type that accepts DocumentItemArray and DocumentItemArrayOutput values.
// You can construct a concrete instance of `DocumentItemArrayInput` via:
//
//	DocumentItemArray{ DocumentItemArgs{...} }
type DocumentItemArrayInput interface {
	pulumi.Input

	ToDocumentItemArrayOutput() DocumentItemArrayOutput
	ToDocumentItemArrayOutputWithContext(context.Context) DocumentItemArrayOutput
}

type DocumentItemArray []DocumentItemInput

func (DocumentItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DocumentItem)(nil)).Elem()
}

func (i DocumentItemArray) ToDocumentItemArrayOutput() DocumentItemArrayOutput {
	return i.ToDocumentItemArrayOutputWithContext(context.Background())
}

func (i DocumentItemArray) ToDocumentItemArrayOutputWithContext(ctx context.Context) DocumentItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentItemArrayOutput)
}

// DocumentItemMapInput is an input type that accepts DocumentItemMap and DocumentItemMapOutput values.
// You can construct a concrete instance of `DocumentItemMapInput` via:
//
//	DocumentItemMap{ "key": DocumentItemArgs{...} }
type DocumentItemMapInput interface {
	pulumi.Input

	ToDocumentItemMapOutput() DocumentItemMapOutput
	ToDocumentItemMapOutputWithContext(context.Context) DocumentItemMapOutput
}

type DocumentItemMap map[string]DocumentItemInput

func (DocumentItemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DocumentItem)(nil)).Elem()
}

func (i DocumentItemMap) ToDocumentItemMapOutput() DocumentItemMapOutput {
	return i.ToDocumentItemMapOutputWithContext(context.Background())
}

func (i DocumentItemMap) ToDocumentItemMapOutputWithContext(ctx context.Context) DocumentItemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentItemMapOutput)
}

type DocumentItemOutput struct{ *pulumi.OutputState }

func (DocumentItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentItem)(nil)).Elem()
}

func (o DocumentItemOutput) ToDocumentItemOutput() DocumentItemOutput {
	return o
}

func (o DocumentItemOutput) ToDocumentItemOutputWithContext(ctx context.Context) DocumentItemOutput {
	return o
}

type DocumentItemArrayOutput struct{ *pulumi.OutputState }

func (DocumentItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DocumentItem)(nil)).Elem()
}

func (o DocumentItemArrayOutput) ToDocumentItemArrayOutput() DocumentItemArrayOutput {
	return o
}

func (o DocumentItemArrayOutput) ToDocumentItemArrayOutputWithContext(ctx context.Context) DocumentItemArrayOutput {
	return o
}

func (o DocumentItemArrayOutput) Index(i pulumi.IntInput) DocumentItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DocumentItem {
		return vs[0].([]*DocumentItem)[vs[1].(int)]
	}).(DocumentItemOutput)
}

type DocumentItemMapOutput struct{ *pulumi.OutputState }

func (DocumentItemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DocumentItem)(nil)).Elem()
}

func (o DocumentItemMapOutput) ToDocumentItemMapOutput() DocumentItemMapOutput {
	return o
}

func (o DocumentItemMapOutput) ToDocumentItemMapOutputWithContext(ctx context.Context) DocumentItemMapOutput {
	return o
}

func (o DocumentItemMapOutput) MapIndex(k pulumi.StringInput) DocumentItemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DocumentItem {
		return vs[0].(map[string]*DocumentItem)[vs[1].(string)]
	}).(DocumentItemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentItemInput)(nil)).Elem(), &DocumentItem{})
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentItemArrayInput)(nil)).Elem(), DocumentItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentItemMapInput)(nil)).Elem(), DocumentItemMap{})
	pulumi.RegisterOutputType(DocumentItemOutput{})
	pulumi.RegisterOutputType(DocumentItemArrayOutput{})
	pulumi.RegisterOutputType(DocumentItemMapOutput{})
}
