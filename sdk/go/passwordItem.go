// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pulumi_one_password_native_unofficial

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PasswordItem struct {
	pulumi.CustomResourceState

	Attachments OutputAttachmentMapOutput `pulumi:"attachments"`
	Category    pulumi.StringOutput       `pulumi:"category"`
	Fields      OutputFieldMapOutput      `pulumi:"fields"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Id         pulumi.StringOutput        `pulumi:"id"`
	Notes      pulumi.StringPtrOutput     `pulumi:"notes"`
	Password   pulumi.StringPtrOutput     `pulumi:"password"`
	References OutputReferenceArrayOutput `pulumi:"references"`
	Sections   OutputSectionMapOutput     `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The title of the item.
	Title pulumi.StringOutput    `pulumi:"title"`
	Urls  OutputUrlArrayOutput   `pulumi:"urls"`
	Vault pulumi.StringMapOutput `pulumi:"vault"`
}

// NewPasswordItem registers a new resource with the given unique name, arguments, and options.
func NewPasswordItem(ctx *pulumi.Context,
	name string, args *PasswordItemArgs, opts ...pulumi.ResourceOption) (*PasswordItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Vault == nil {
		return nil, errors.New("invalid value for required argument 'Vault'")
	}
	args.Category = pulumi.StringPtr("Password")
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrOutput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"attachments",
		"fields",
		"password",
		"sections",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource PasswordItem
	err := ctx.RegisterResource("one-password-native-unofficial:index:PasswordItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPasswordItem gets an existing PasswordItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPasswordItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PasswordItemState, opts ...pulumi.ResourceOption) (*PasswordItem, error) {
	var resource PasswordItem
	err := ctx.ReadResource("one-password-native-unofficial:index:PasswordItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PasswordItem resources.
type passwordItemState struct {
	// The UUID of the vault the item is in.
	Vault *string `pulumi:"vault"`
}

type PasswordItemState struct {
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput
}

func (PasswordItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*passwordItemState)(nil)).Elem()
}

type passwordItemArgs struct {
	Attachments map[string]pulumi.AssetOrArchive `pulumi:"attachments"`
	// The category of the vault the item is in.
	Category         *string            `pulumi:"category"`
	Fields           map[string]Field   `pulumi:"fields"`
	GeneratePassword interface{}        `pulumi:"generatePassword"`
	Notes            *string            `pulumi:"notes"`
	Password         *string            `pulumi:"password"`
	Sections         map[string]Section `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title *string `pulumi:"title"`
	Urls  []Url   `pulumi:"urls"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

// The set of arguments for constructing a PasswordItem resource.
type PasswordItemArgs struct {
	Attachments pulumi.AssetOrArchiveMapInput
	// The category of the vault the item is in.
	Category         pulumi.StringPtrInput
	Fields           FieldMapInput
	GeneratePassword pulumi.Input
	Notes            pulumi.StringPtrInput
	Password         pulumi.StringPtrInput
	Sections         SectionMapInput
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayInput
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title pulumi.StringPtrInput
	Urls  UrlArrayInput
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput
}

func (PasswordItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*passwordItemArgs)(nil)).Elem()
}

type PasswordItemInput interface {
	pulumi.Input

	ToPasswordItemOutput() PasswordItemOutput
	ToPasswordItemOutputWithContext(ctx context.Context) PasswordItemOutput
}

func (*PasswordItem) ElementType() reflect.Type {
	return reflect.TypeOf((**PasswordItem)(nil)).Elem()
}

func (i *PasswordItem) ToPasswordItemOutput() PasswordItemOutput {
	return i.ToPasswordItemOutputWithContext(context.Background())
}

func (i *PasswordItem) ToPasswordItemOutputWithContext(ctx context.Context) PasswordItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PasswordItemOutput)
}

// PasswordItemArrayInput is an input type that accepts PasswordItemArray and PasswordItemArrayOutput values.
// You can construct a concrete instance of `PasswordItemArrayInput` via:
//
//	PasswordItemArray{ PasswordItemArgs{...} }
type PasswordItemArrayInput interface {
	pulumi.Input

	ToPasswordItemArrayOutput() PasswordItemArrayOutput
	ToPasswordItemArrayOutputWithContext(context.Context) PasswordItemArrayOutput
}

type PasswordItemArray []PasswordItemInput

func (PasswordItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PasswordItem)(nil)).Elem()
}

func (i PasswordItemArray) ToPasswordItemArrayOutput() PasswordItemArrayOutput {
	return i.ToPasswordItemArrayOutputWithContext(context.Background())
}

func (i PasswordItemArray) ToPasswordItemArrayOutputWithContext(ctx context.Context) PasswordItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PasswordItemArrayOutput)
}

// PasswordItemMapInput is an input type that accepts PasswordItemMap and PasswordItemMapOutput values.
// You can construct a concrete instance of `PasswordItemMapInput` via:
//
//	PasswordItemMap{ "key": PasswordItemArgs{...} }
type PasswordItemMapInput interface {
	pulumi.Input

	ToPasswordItemMapOutput() PasswordItemMapOutput
	ToPasswordItemMapOutputWithContext(context.Context) PasswordItemMapOutput
}

type PasswordItemMap map[string]PasswordItemInput

func (PasswordItemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PasswordItem)(nil)).Elem()
}

func (i PasswordItemMap) ToPasswordItemMapOutput() PasswordItemMapOutput {
	return i.ToPasswordItemMapOutputWithContext(context.Background())
}

func (i PasswordItemMap) ToPasswordItemMapOutputWithContext(ctx context.Context) PasswordItemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PasswordItemMapOutput)
}

type PasswordItemOutput struct{ *pulumi.OutputState }

func (PasswordItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PasswordItem)(nil)).Elem()
}

func (o PasswordItemOutput) ToPasswordItemOutput() PasswordItemOutput {
	return o
}

func (o PasswordItemOutput) ToPasswordItemOutputWithContext(ctx context.Context) PasswordItemOutput {
	return o
}

type PasswordItemArrayOutput struct{ *pulumi.OutputState }

func (PasswordItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PasswordItem)(nil)).Elem()
}

func (o PasswordItemArrayOutput) ToPasswordItemArrayOutput() PasswordItemArrayOutput {
	return o
}

func (o PasswordItemArrayOutput) ToPasswordItemArrayOutputWithContext(ctx context.Context) PasswordItemArrayOutput {
	return o
}

func (o PasswordItemArrayOutput) Index(i pulumi.IntInput) PasswordItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PasswordItem {
		return vs[0].([]*PasswordItem)[vs[1].(int)]
	}).(PasswordItemOutput)
}

type PasswordItemMapOutput struct{ *pulumi.OutputState }

func (PasswordItemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PasswordItem)(nil)).Elem()
}

func (o PasswordItemMapOutput) ToPasswordItemMapOutput() PasswordItemMapOutput {
	return o
}

func (o PasswordItemMapOutput) ToPasswordItemMapOutputWithContext(ctx context.Context) PasswordItemMapOutput {
	return o
}

func (o PasswordItemMapOutput) MapIndex(k pulumi.StringInput) PasswordItemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PasswordItem {
		return vs[0].(map[string]*PasswordItem)[vs[1].(string)]
	}).(PasswordItemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PasswordItemInput)(nil)).Elem(), &PasswordItem{})
	pulumi.RegisterInputType(reflect.TypeOf((*PasswordItemArrayInput)(nil)).Elem(), PasswordItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PasswordItemMapInput)(nil)).Elem(), PasswordItemMap{})
	pulumi.RegisterOutputType(PasswordItemOutput{})
	pulumi.RegisterOutputType(PasswordItemArrayOutput{})
	pulumi.RegisterOutputType(PasswordItemMapOutput{})
}
