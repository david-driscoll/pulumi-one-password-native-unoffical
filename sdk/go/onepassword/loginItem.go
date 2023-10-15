// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package onepassword

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LoginItem struct {
	pulumi.CustomResourceState

	Category pulumi.StringOutput    `pulumi:"category"`
	Fields   GetFieldMapOutput      `pulumi:"fields"`
	Id       pulumi.StringOutput    `pulumi:"id"`
	Notes    pulumi.StringPtrOutput `pulumi:"notes"`
	Password pulumi.StringPtrOutput `pulumi:"password"`
	Sections GetSectionMapOutput    `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The title of the item.
	Title    pulumi.StringOutput    `pulumi:"title"`
	Username pulumi.StringPtrOutput `pulumi:"username"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
	// The UUID of the vault the item is in.
	Vault pulumi.StringOutput `pulumi:"vault"`
}

// NewLoginItem registers a new resource with the given unique name, arguments, and options.
func NewLoginItem(ctx *pulumi.Context,
	name string, args *LoginItemArgs, opts ...pulumi.ResourceOption) (*LoginItem, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Vault == nil {
		return nil, errors.New("invalid value for required argument 'Vault'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringPtrOutput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	var resource LoginItem
	err := ctx.RegisterResource("onepassword:index:LoginItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoginItem gets an existing LoginItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoginItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoginItemState, opts ...pulumi.ResourceOption) (*LoginItem, error) {
	var resource LoginItem
	err := ctx.ReadResource("onepassword:index:LoginItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoginItem resources.
type loginItemState struct {
}

type LoginItemState struct {
}

func (LoginItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*loginItemState)(nil)).Elem()
}

type loginItemArgs struct {
	Fields   map[string]Field   `pulumi:"fields"`
	Notes    *string            `pulumi:"notes"`
	Password *string            `pulumi:"password"`
	Sections map[string]Section `pulumi:"sections"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title    *string `pulumi:"title"`
	Username *string `pulumi:"username"`
	// The UUID of the vault the item is in.
	Vault string `pulumi:"vault"`
}

// The set of arguments for constructing a LoginItem resource.
type LoginItemArgs struct {
	Fields   FieldMapInput
	Notes    pulumi.StringPtrInput
	Password pulumi.StringPtrInput
	Sections SectionMapInput
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayInput
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title    pulumi.StringPtrInput
	Username pulumi.StringPtrInput
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput
}

func (LoginItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loginItemArgs)(nil)).Elem()
}

type LoginItemInput interface {
	pulumi.Input

	ToLoginItemOutput() LoginItemOutput
	ToLoginItemOutputWithContext(ctx context.Context) LoginItemOutput
}

func (*LoginItem) ElementType() reflect.Type {
	return reflect.TypeOf((**LoginItem)(nil)).Elem()
}

func (i *LoginItem) ToLoginItemOutput() LoginItemOutput {
	return i.ToLoginItemOutputWithContext(context.Background())
}

func (i *LoginItem) ToLoginItemOutputWithContext(ctx context.Context) LoginItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoginItemOutput)
}

// LoginItemArrayInput is an input type that accepts LoginItemArray and LoginItemArrayOutput values.
// You can construct a concrete instance of `LoginItemArrayInput` via:
//
//	LoginItemArray{ LoginItemArgs{...} }
type LoginItemArrayInput interface {
	pulumi.Input

	ToLoginItemArrayOutput() LoginItemArrayOutput
	ToLoginItemArrayOutputWithContext(context.Context) LoginItemArrayOutput
}

type LoginItemArray []LoginItemInput

func (LoginItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LoginItem)(nil)).Elem()
}

func (i LoginItemArray) ToLoginItemArrayOutput() LoginItemArrayOutput {
	return i.ToLoginItemArrayOutputWithContext(context.Background())
}

func (i LoginItemArray) ToLoginItemArrayOutputWithContext(ctx context.Context) LoginItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoginItemArrayOutput)
}

// LoginItemMapInput is an input type that accepts LoginItemMap and LoginItemMapOutput values.
// You can construct a concrete instance of `LoginItemMapInput` via:
//
//	LoginItemMap{ "key": LoginItemArgs{...} }
type LoginItemMapInput interface {
	pulumi.Input

	ToLoginItemMapOutput() LoginItemMapOutput
	ToLoginItemMapOutputWithContext(context.Context) LoginItemMapOutput
}

type LoginItemMap map[string]LoginItemInput

func (LoginItemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LoginItem)(nil)).Elem()
}

func (i LoginItemMap) ToLoginItemMapOutput() LoginItemMapOutput {
	return i.ToLoginItemMapOutputWithContext(context.Background())
}

func (i LoginItemMap) ToLoginItemMapOutputWithContext(ctx context.Context) LoginItemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoginItemMapOutput)
}

type LoginItemOutput struct{ *pulumi.OutputState }

func (LoginItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoginItem)(nil)).Elem()
}

func (o LoginItemOutput) ToLoginItemOutput() LoginItemOutput {
	return o
}

func (o LoginItemOutput) ToLoginItemOutputWithContext(ctx context.Context) LoginItemOutput {
	return o
}

type LoginItemArrayOutput struct{ *pulumi.OutputState }

func (LoginItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LoginItem)(nil)).Elem()
}

func (o LoginItemArrayOutput) ToLoginItemArrayOutput() LoginItemArrayOutput {
	return o
}

func (o LoginItemArrayOutput) ToLoginItemArrayOutputWithContext(ctx context.Context) LoginItemArrayOutput {
	return o
}

func (o LoginItemArrayOutput) Index(i pulumi.IntInput) LoginItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LoginItem {
		return vs[0].([]*LoginItem)[vs[1].(int)]
	}).(LoginItemOutput)
}

type LoginItemMapOutput struct{ *pulumi.OutputState }

func (LoginItemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LoginItem)(nil)).Elem()
}

func (o LoginItemMapOutput) ToLoginItemMapOutput() LoginItemMapOutput {
	return o
}

func (o LoginItemMapOutput) ToLoginItemMapOutputWithContext(ctx context.Context) LoginItemMapOutput {
	return o
}

func (o LoginItemMapOutput) MapIndex(k pulumi.StringInput) LoginItemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LoginItem {
		return vs[0].(map[string]*LoginItem)[vs[1].(string)]
	}).(LoginItemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LoginItemInput)(nil)).Elem(), &LoginItem{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoginItemArrayInput)(nil)).Elem(), LoginItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoginItemMapInput)(nil)).Elem(), LoginItemMap{})
	pulumi.RegisterOutputType(LoginItemOutput{})
	pulumi.RegisterOutputType(LoginItemArrayOutput{})
	pulumi.RegisterOutputType(LoginItemMapOutput{})
}
