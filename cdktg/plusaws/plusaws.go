package plusaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hupe1980/cdk-threagile-go/cdktg/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hupe1980/cdk-threagile-go/cdktg"
	"github.com/hupe1980/cdk-threagile-go/cdktg/plusaws/internal"
)

type ApplicationLoadBalancer interface {
	cdktg.TechnicalAsset
	CiaTriad() cdktg.CIATriad
	CustomDevelopedParts() *bool
	DataFormatsAccepted() *[]cdktg.DataFormat
	Description() *string
	Encryption() cdktg.Encryption
	HighestAvailability() cdktg.Availability
	HighestIntegrity() cdktg.Integrity
	HumanUse() *bool
	Id() *string
	Internet() *bool
	Machine() cdktg.Machine
	MultiTenant() *bool
	// The tree node.
	Node() constructs.Node
	Owner() *string
	Redundant() *bool
	Scope() cdktg.Scope
	SecurityGroup() SecurityGroup
	Size() cdktg.Size
	Tags() *[]*string
	Technology() cdktg.Technology
	Title() *string
	TrustBoundary() cdktg.TrustBoundary
	SetTrustBoundary(val cdktg.TrustBoundary)
	Type() cdktg.TechnicalAssetType
	Usage() cdktg.Usage
	CommunicatesWith(id *string, target cdktg.TechnicalAsset, options *cdktg.CommunicationOptions) cdktg.Communication
	IsInScope() *bool
	IsTrafficForwarding() *bool
	IsWebApplication() *bool
	IsWebService() *bool
	Processes(assets ...cdktg.DataAsset)
	Stores(assets ...cdktg.DataAsset)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for ApplicationLoadBalancer
type jsiiProxy_ApplicationLoadBalancer struct {
	internal.Type__cdktgTechnicalAsset
}

func (j *jsiiProxy_ApplicationLoadBalancer) CiaTriad() cdktg.CIATriad {
	var returns cdktg.CIATriad
	_jsii_.Get(
		j,
		"ciaTriad",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) CustomDevelopedParts() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"customDevelopedParts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) DataFormatsAccepted() *[]cdktg.DataFormat {
	var returns *[]cdktg.DataFormat
	_jsii_.Get(
		j,
		"dataFormatsAccepted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) Encryption() cdktg.Encryption {
	var returns cdktg.Encryption
	_jsii_.Get(
		j,
		"encryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) HighestAvailability() cdktg.Availability {
	var returns cdktg.Availability
	_jsii_.Get(
		j,
		"highestAvailability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) HighestIntegrity() cdktg.Integrity {
	var returns cdktg.Integrity
	_jsii_.Get(
		j,
		"highestIntegrity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) HumanUse() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"humanUse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) Internet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"internet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) Machine() cdktg.Machine {
	var returns cdktg.Machine
	_jsii_.Get(
		j,
		"machine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) MultiTenant() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"multiTenant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) Redundant() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"redundant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) Scope() cdktg.Scope {
	var returns cdktg.Scope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) SecurityGroup() SecurityGroup {
	var returns SecurityGroup
	_jsii_.Get(
		j,
		"securityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) Size() cdktg.Size {
	var returns cdktg.Size
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) Technology() cdktg.Technology {
	var returns cdktg.Technology
	_jsii_.Get(
		j,
		"technology",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) TrustBoundary() cdktg.TrustBoundary {
	var returns cdktg.TrustBoundary
	_jsii_.Get(
		j,
		"trustBoundary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) Type() cdktg.TechnicalAssetType {
	var returns cdktg.TechnicalAssetType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancer) Usage() cdktg.Usage {
	var returns cdktg.Usage
	_jsii_.Get(
		j,
		"usage",
		&returns,
	)
	return returns
}


func NewApplicationLoadBalancer(scope constructs.Construct, id *string, props *ApplicationLoadBalancerProps) ApplicationLoadBalancer {
	_init_.Initialize()

	j := jsiiProxy_ApplicationLoadBalancer{}

	_jsii_.Create(
		"cdktg.plus_aws.ApplicationLoadBalancer",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewApplicationLoadBalancer_Override(a ApplicationLoadBalancer, scope constructs.Construct, id *string, props *ApplicationLoadBalancerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktg.plus_aws.ApplicationLoadBalancer",
		[]interface{}{scope, id, props},
		a,
	)
}

func (j *jsiiProxy_ApplicationLoadBalancer) SetTrustBoundary(val cdktg.TrustBoundary) {
	_jsii_.Set(
		j,
		"trustBoundary",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func ApplicationLoadBalancer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktg.plus_aws.ApplicationLoadBalancer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) CommunicatesWith(id *string, target cdktg.TechnicalAsset, options *cdktg.CommunicationOptions) cdktg.Communication {
	var returns cdktg.Communication

	_jsii_.Invoke(
		a,
		"communicatesWith",
		[]interface{}{id, target, options},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) IsInScope() *bool {
	var returns *bool

	_jsii_.Invoke(
		a,
		"isInScope",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) IsTrafficForwarding() *bool {
	var returns *bool

	_jsii_.Invoke(
		a,
		"isTrafficForwarding",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) IsWebApplication() *bool {
	var returns *bool

	_jsii_.Invoke(
		a,
		"isWebApplication",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) IsWebService() *bool {
	var returns *bool

	_jsii_.Invoke(
		a,
		"isWebService",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancer) Processes(assets ...cdktg.DataAsset) {
	args := []interface{}{}
	for _, a := range assets {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"processes",
		args,
	)
}

func (a *jsiiProxy_ApplicationLoadBalancer) Stores(assets ...cdktg.DataAsset) {
	args := []interface{}{}
	for _, a := range assets {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"stores",
		args,
	)
}

func (a *jsiiProxy_ApplicationLoadBalancer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type ApplicationLoadBalancerProps struct {
	CiaTriad cdktg.CIATriad `field:"required" json:"ciaTriad" yaml:"ciaTriad"`
	Description *string `field:"optional" json:"description" yaml:"description"`
	SecurityGroup SecurityGroup `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	Waf *bool `field:"optional" json:"waf" yaml:"waf"`
}

type Cloud interface {
	cdktg.TrustBoundary
	Description() *string
	Id() *string
	// The tree node.
	Node() constructs.Node
	Tags() *[]*string
	Title() *string
	Type() cdktg.TrustBoundaryType
	AddTechnicalAssets(assets ...cdktg.TechnicalAsset)
	AddTrustBoundary(boundary cdktg.TrustBoundary)
	IsNetworkBoundary() *bool
	IsWithinCloud() *bool
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Cloud
type jsiiProxy_Cloud struct {
	internal.Type__cdktgTrustBoundary
}

func (j *jsiiProxy_Cloud) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloud) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloud) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloud) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloud) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloud) Type() cdktg.TrustBoundaryType {
	var returns cdktg.TrustBoundaryType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func NewCloud(scope constructs.Construct, id *string, props *CloudProps) Cloud {
	_init_.Initialize()

	j := jsiiProxy_Cloud{}

	_jsii_.Create(
		"cdktg.plus_aws.Cloud",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCloud_Override(c Cloud, scope constructs.Construct, id *string, props *CloudProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktg.plus_aws.Cloud",
		[]interface{}{scope, id, props},
		c,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Cloud_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktg.plus_aws.Cloud",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cloud) AddTechnicalAssets(assets ...cdktg.TechnicalAsset) {
	args := []interface{}{}
	for _, a := range assets {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"addTechnicalAssets",
		args,
	)
}

func (c *jsiiProxy_Cloud) AddTrustBoundary(boundary cdktg.TrustBoundary) {
	_jsii_.InvokeVoid(
		c,
		"addTrustBoundary",
		[]interface{}{boundary},
	)
}

func (c *jsiiProxy_Cloud) IsNetworkBoundary() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"isNetworkBoundary",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cloud) IsWithinCloud() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"isWithinCloud",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cloud) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type CloudProps struct {
	Description *string `field:"optional" json:"description" yaml:"description"`
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

type SecurityGroup interface {
	cdktg.TrustBoundary
	Description() *string
	Id() *string
	// The tree node.
	Node() constructs.Node
	Tags() *[]*string
	Title() *string
	Type() cdktg.TrustBoundaryType
	AddTechnicalAssets(assets ...cdktg.TechnicalAsset)
	AddTrustBoundary(boundary cdktg.TrustBoundary)
	IsNetworkBoundary() *bool
	IsWithinCloud() *bool
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for SecurityGroup
type jsiiProxy_SecurityGroup struct {
	internal.Type__cdktgTrustBoundary
}

func (j *jsiiProxy_SecurityGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroup) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroup) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroup) Type() cdktg.TrustBoundaryType {
	var returns cdktg.TrustBoundaryType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func NewSecurityGroup(scope constructs.Construct, id *string, props *SecurityGroupProps) SecurityGroup {
	_init_.Initialize()

	j := jsiiProxy_SecurityGroup{}

	_jsii_.Create(
		"cdktg.plus_aws.SecurityGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSecurityGroup_Override(s SecurityGroup, scope constructs.Construct, id *string, props *SecurityGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdktg.plus_aws.SecurityGroup",
		[]interface{}{scope, id, props},
		s,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func SecurityGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"cdktg.plus_aws.SecurityGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityGroup) AddTechnicalAssets(assets ...cdktg.TechnicalAsset) {
	args := []interface{}{}
	for _, a := range assets {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		s,
		"addTechnicalAssets",
		args,
	)
}

func (s *jsiiProxy_SecurityGroup) AddTrustBoundary(boundary cdktg.TrustBoundary) {
	_jsii_.InvokeVoid(
		s,
		"addTrustBoundary",
		[]interface{}{boundary},
	)
}

func (s *jsiiProxy_SecurityGroup) IsNetworkBoundary() *bool {
	var returns *bool

	_jsii_.Invoke(
		s,
		"isNetworkBoundary",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityGroup) IsWithinCloud() *bool {
	var returns *bool

	_jsii_.Invoke(
		s,
		"isWithinCloud",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SecurityGroupProps struct {
	Description *string `field:"optional" json:"description" yaml:"description"`
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}

