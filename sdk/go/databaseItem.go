// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pulumi_one_password_native_unofficial

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DatabaseItem struct {
	pulumi.CustomResourceState

	Alias             pulumi.StringPtrOutput    `pulumi:"alias"`
	Attachments       OutputAttachmentMapOutput `pulumi:"attachments"`
	Category          pulumi.StringOutput       `pulumi:"category"`
	ConnectionOptions pulumi.StringPtrOutput    `pulumi:"connectionOptions"`
	Database          pulumi.StringPtrOutput    `pulumi:"database"`
	Fields            OutputFieldMapOutput      `pulumi:"fields"`
	// The UUID of the item to retrieve. This field will be populated with the UUID of the item if the item it looked up by its title.
	Id         pulumi.StringOutput        `pulumi:"id"`
	Notes      pulumi.StringPtrOutput     `pulumi:"notes"`
	Password   pulumi.StringPtrOutput     `pulumi:"password"`
	Port       pulumi.StringPtrOutput     `pulumi:"port"`
	References OutputReferenceArrayOutput `pulumi:"references"`
	Sections   OutputSectionMapOutput     `pulumi:"sections"`
	Server     pulumi.StringPtrOutput     `pulumi:"server"`
	Sid        pulumi.StringPtrOutput     `pulumi:"sid"`
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The title of the item.
	Title    pulumi.StringOutput    `pulumi:"title"`
	Type     pulumi.StringPtrOutput `pulumi:"type"`
	Urls     OutputUrlArrayOutput   `pulumi:"urls"`
	Username pulumi.StringPtrOutput `pulumi:"username"`
	Vault    pulumi.StringMapOutput `pulumi:"vault"`
}

// NewDatabaseItem registers a new resource with the given unique name, arguments, and options.
func NewDatabaseItem(ctx *pulumi.Context,
	name string, args *DatabaseItemArgs, opts ...pulumi.ResourceOption) (*DatabaseItem, error) {
	if args == nil {
		args = &DatabaseItemArgs{}
	}

	args.Category = pulumi.StringPtr("Database")
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
	var resource DatabaseItem
	err := ctx.RegisterResource("one-password-native-unofficial:index:DatabaseItem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseItem gets an existing DatabaseItem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseItem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseItemState, opts ...pulumi.ResourceOption) (*DatabaseItem, error) {
	var resource DatabaseItem
	err := ctx.ReadResource("one-password-native-unofficial:index:DatabaseItem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseItem resources.
type databaseItemState struct {
	// The UUID of the vault the item is in.
	Vault *string `pulumi:"vault"`
}

type DatabaseItemState struct {
	// The UUID of the vault the item is in.
	Vault pulumi.StringInput
}

func (DatabaseItemState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseItemState)(nil)).Elem()
}

type databaseItemArgs struct {
	Alias       *string                          `pulumi:"alias"`
	Attachments map[string]pulumi.AssetOrArchive `pulumi:"attachments"`
	// The category of the vault the item is in.
	Category          *string                `pulumi:"category"`
	ConnectionOptions *string                `pulumi:"connectionOptions"`
	Database          *string                `pulumi:"database"`
	Fields            map[string]interface{} `pulumi:"fields"`
	Notes             *string                `pulumi:"notes"`
	Password          *string                `pulumi:"password"`
	Port              *string                `pulumi:"port"`
	References        []string               `pulumi:"references"`
	Sections          map[string]Section     `pulumi:"sections"`
	Server            *string                `pulumi:"server"`
	Sid               *string                `pulumi:"sid"`
	// An array of strings of the tags assigned to the item.
	Tags []string `pulumi:"tags"`
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title    *string       `pulumi:"title"`
	Type     *string       `pulumi:"type"`
	Urls     []interface{} `pulumi:"urls"`
	Username *string       `pulumi:"username"`
	// The UUID of the vault the item is in.
	Vault *string `pulumi:"vault"`
}

// The set of arguments for constructing a DatabaseItem resource.
type DatabaseItemArgs struct {
	Alias       pulumi.StringPtrInput
	Attachments pulumi.AssetOrArchiveMapInput
	// The category of the vault the item is in.
	Category          pulumi.StringPtrInput
	ConnectionOptions pulumi.StringPtrInput
	Database          pulumi.StringPtrInput
	Fields            pulumi.MapInput
	Notes             pulumi.StringPtrInput
	Password          pulumi.StringPtrInput
	Port              pulumi.StringPtrInput
	References        pulumi.StringArrayInput
	Sections          SectionMapInput
	Server            pulumi.StringPtrInput
	Sid               pulumi.StringPtrInput
	// An array of strings of the tags assigned to the item.
	Tags pulumi.StringArrayInput
	// The title of the item to retrieve. This field will be populated with the title of the item if the item it looked up by its UUID.
	Title    pulumi.StringPtrInput
	Type     pulumi.StringPtrInput
	Urls     pulumi.ArrayInput
	Username pulumi.StringPtrInput
	// The UUID of the vault the item is in.
	Vault pulumi.StringPtrInput
}

func (DatabaseItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseItemArgs)(nil)).Elem()
}

type DatabaseItemInput interface {
	pulumi.Input

	ToDatabaseItemOutput() DatabaseItemOutput
	ToDatabaseItemOutputWithContext(ctx context.Context) DatabaseItemOutput
}

func (*DatabaseItem) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseItem)(nil)).Elem()
}

func (i *DatabaseItem) ToDatabaseItemOutput() DatabaseItemOutput {
	return i.ToDatabaseItemOutputWithContext(context.Background())
}

func (i *DatabaseItem) ToDatabaseItemOutputWithContext(ctx context.Context) DatabaseItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseItemOutput)
}

// DatabaseItemArrayInput is an input type that accepts DatabaseItemArray and DatabaseItemArrayOutput values.
// You can construct a concrete instance of `DatabaseItemArrayInput` via:
//
//	DatabaseItemArray{ DatabaseItemArgs{...} }
type DatabaseItemArrayInput interface {
	pulumi.Input

	ToDatabaseItemArrayOutput() DatabaseItemArrayOutput
	ToDatabaseItemArrayOutputWithContext(context.Context) DatabaseItemArrayOutput
}

type DatabaseItemArray []DatabaseItemInput

func (DatabaseItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseItem)(nil)).Elem()
}

func (i DatabaseItemArray) ToDatabaseItemArrayOutput() DatabaseItemArrayOutput {
	return i.ToDatabaseItemArrayOutputWithContext(context.Background())
}

func (i DatabaseItemArray) ToDatabaseItemArrayOutputWithContext(ctx context.Context) DatabaseItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseItemArrayOutput)
}

// DatabaseItemMapInput is an input type that accepts DatabaseItemMap and DatabaseItemMapOutput values.
// You can construct a concrete instance of `DatabaseItemMapInput` via:
//
//	DatabaseItemMap{ "key": DatabaseItemArgs{...} }
type DatabaseItemMapInput interface {
	pulumi.Input

	ToDatabaseItemMapOutput() DatabaseItemMapOutput
	ToDatabaseItemMapOutputWithContext(context.Context) DatabaseItemMapOutput
}

type DatabaseItemMap map[string]DatabaseItemInput

func (DatabaseItemMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseItem)(nil)).Elem()
}

func (i DatabaseItemMap) ToDatabaseItemMapOutput() DatabaseItemMapOutput {
	return i.ToDatabaseItemMapOutputWithContext(context.Background())
}

func (i DatabaseItemMap) ToDatabaseItemMapOutputWithContext(ctx context.Context) DatabaseItemMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseItemMapOutput)
}

type DatabaseItemOutput struct{ *pulumi.OutputState }

func (DatabaseItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseItem)(nil)).Elem()
}

func (o DatabaseItemOutput) ToDatabaseItemOutput() DatabaseItemOutput {
	return o
}

func (o DatabaseItemOutput) ToDatabaseItemOutputWithContext(ctx context.Context) DatabaseItemOutput {
	return o
}

type DatabaseItemArrayOutput struct{ *pulumi.OutputState }

func (DatabaseItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseItem)(nil)).Elem()
}

func (o DatabaseItemArrayOutput) ToDatabaseItemArrayOutput() DatabaseItemArrayOutput {
	return o
}

func (o DatabaseItemArrayOutput) ToDatabaseItemArrayOutputWithContext(ctx context.Context) DatabaseItemArrayOutput {
	return o
}

func (o DatabaseItemArrayOutput) Index(i pulumi.IntInput) DatabaseItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatabaseItem {
		return vs[0].([]*DatabaseItem)[vs[1].(int)]
	}).(DatabaseItemOutput)
}

type DatabaseItemMapOutput struct{ *pulumi.OutputState }

func (DatabaseItemMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseItem)(nil)).Elem()
}

func (o DatabaseItemMapOutput) ToDatabaseItemMapOutput() DatabaseItemMapOutput {
	return o
}

func (o DatabaseItemMapOutput) ToDatabaseItemMapOutputWithContext(ctx context.Context) DatabaseItemMapOutput {
	return o
}

func (o DatabaseItemMapOutput) MapIndex(k pulumi.StringInput) DatabaseItemOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatabaseItem {
		return vs[0].(map[string]*DatabaseItem)[vs[1].(string)]
	}).(DatabaseItemOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseItemInput)(nil)).Elem(), &DatabaseItem{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseItemArrayInput)(nil)).Elem(), DatabaseItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseItemMapInput)(nil)).Elem(), DatabaseItemMap{})
	pulumi.RegisterOutputType(DatabaseItemOutput{})
	pulumi.RegisterOutputType(DatabaseItemArrayOutput{})
	pulumi.RegisterOutputType(DatabaseItemMapOutput{})
}
