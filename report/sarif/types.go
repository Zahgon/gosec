package sarif

type Address struct {
	AbsoluteAddress int `json:"absoluteAddress,omitempty"`

	FullyQualifiedName string `json:"fullyQualifiedName,omitempty"`

	Index int `json:"index,omitempty"`

	Kind string `json:"kind,omitempty"`

	Length int `json:"length,omitempty"`

	Name string `json:"name,omitempty"`

	OffsetFromParent int `json:"offsetFromParent,omitempty"`

	ParentIndex int `json:"parentIndex,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	RelativeAddress int `json:"relativeAddress,omitempty"`
}

type Artifact struct {
	Contents *ArtifactContent `json:"contents,omitempty"`

	Description *Message `json:"description,omitempty"`

	Encoding string `json:"encoding,omitempty"`

	Hashes map[string]string `json:"hashes,omitempty"`

	LastModifiedTimeUtc string `json:"lastModifiedTimeUtc,omitempty"`

	Length int `json:"length,omitempty"`

	Location *ArtifactLocation `json:"location,omitempty"`

	MimeType string `json:"mimeType,omitempty"`

	Offset int `json:"offset,omitempty"`

	ParentIndex int `json:"parentIndex,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	Roles []interface{} `json:"roles,omitempty"`

	SourceLanguage string `json:"sourceLanguage,omitempty"`
}

type ArtifactChange struct {
	ArtifactLocation *ArtifactLocation `json:"artifactLocation"`

	Properties *PropertyBag `json:"properties,omitempty"`

	Replacements []*Replacement `json:"replacements"`
}

type ArtifactContent struct {
	Binary string `json:"binary,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	Rendered *MultiformatMessageString `json:"rendered,omitempty"`

	Text string `json:"text,omitempty"`
}

type ArtifactLocation struct {
	Description *Message `json:"description,omitempty"`

	Index int `json:"index,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	URI string `json:"uri,omitempty"`

	UriBaseID string `json:"uriBaseId,omitempty"`
}

type Attachment struct {
	ArtifactLocation *ArtifactLocation `json:"artifactLocation"`

	Description *Message `json:"description,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	Rectangles []*Rectangle `json:"rectangles,omitempty"`

	Regions []*Region `json:"regions,omitempty"`
}

type CodeFlow struct {
	Message *Message `json:"message,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	ThreadFlows []*ThreadFlow `json:"threadFlows"`
}

type ConfigurationOverride struct {
	Configuration *ReportingConfiguration `json:"configuration"`

	Descriptor *ReportingDescriptorReference `json:"descriptor"`

	Properties *PropertyBag `json:"properties,omitempty"`
}

type Conversion struct {
	AnalysisToolLogFiles []*ArtifactLocation `json:"analysisToolLogFiles,omitempty"`

	Invocation *Invocation `json:"invocation,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	Tool *Tool `json:"tool"`
}

type Edge struct {
	ID string `json:"id"`

	Label *Message `json:"label,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	SourceNodeID string `json:"sourceNodeId"`

	TargetNodeID string `json:"targetNodeId"`
}

type EdgeTraversal struct {
	EdgeID string `json:"edgeId"`

	FinalState map[string]*MultiformatMessageString `json:"finalState,omitempty"`

	Message *Message `json:"message,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	StepOverEdgeCount int `json:"stepOverEdgeCount,omitempty"`
}

type Exception struct {
	InnerExceptions []*Exception `json:"innerExceptions,omitempty"`

	Kind string `json:"kind,omitempty"`

	Message string `json:"message,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	Stack *Stack `json:"stack,omitempty"`
}

type ExternalProperties struct {
	Addresses []*Address `json:"addresses,omitempty"`

	Artifacts []*Artifact `json:"artifacts,omitempty"`

	Conversion *Conversion `json:"conversion,omitempty"`

	Driver *ToolComponent `json:"driver,omitempty"`

	Extensions []*ToolComponent `json:"extensions,omitempty"`

	ExternalizedProperties *PropertyBag `json:"externalizedProperties,omitempty"`

	Graphs []*Graph `json:"graphs,omitempty"`

	GUID string `json:"guid,omitempty"`

	Invocations []*Invocation `json:"invocations,omitempty"`

	LogicalLocations []*LogicalLocation `json:"logicalLocations,omitempty"`

	Policies []*ToolComponent `json:"policies,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	Results []*Result `json:"results,omitempty"`

	RunGUID string `json:"runGuid,omitempty"`

	Schema string `json:"schema,omitempty"`

	Taxonomies []*ToolComponent `json:"taxonomies,omitempty"`

	ThreadFlowLocations []*ThreadFlowLocation `json:"threadFlowLocations,omitempty"`

	Translations []*ToolComponent `json:"translations,omitempty"`

	Version interface{} `json:"version,omitempty"`

	WebRequests []*WebRequest `json:"webRequests,omitempty"`

	WebResponses []*WebResponse `json:"webResponses,omitempty"`
}

type ExternalPropertyFileReference struct {
	GUID string `json:"guid,omitempty"`

	ItemCount int `json:"itemCount,omitempty"`

	Location *ArtifactLocation `json:"location,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`
}

type ExternalPropertyFileReferences struct {
	Addresses []*ExternalPropertyFileReference `json:"addresses,omitempty"`

	Artifacts []*ExternalPropertyFileReference `json:"artifacts,omitempty"`

	Conversion *ExternalPropertyFileReference `json:"conversion,omitempty"`

	Driver *ExternalPropertyFileReference `json:"driver,omitempty"`

	Extensions []*ExternalPropertyFileReference `json:"extensions,omitempty"`

	ExternalizedProperties *ExternalPropertyFileReference `json:"externalizedProperties,omitempty"`

	Graphs []*ExternalPropertyFileReference `json:"graphs,omitempty"`

	Invocations []*ExternalPropertyFileReference `json:"invocations,omitempty"`

	LogicalLocations []*ExternalPropertyFileReference `json:"logicalLocations,omitempty"`

	Policies []*ExternalPropertyFileReference `json:"policies,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	Results []*ExternalPropertyFileReference `json:"results,omitempty"`

	Taxonomies []*ExternalPropertyFileReference `json:"taxonomies,omitempty"`

	ThreadFlowLocations []*ExternalPropertyFileReference `json:"threadFlowLocations,omitempty"`

	Translations []*ExternalPropertyFileReference `json:"translations,omitempty"`

	WebRequests []*ExternalPropertyFileReference `json:"webRequests,omitempty"`

	WebResponses []*ExternalPropertyFileReference `json:"webResponses,omitempty"`
}

type Fix struct {
	ArtifactChanges []*ArtifactChange `json:"artifactChanges"`

	Description *Message `json:"description,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`
}

type Graph struct {
	Description *Message `json:"description,omitempty"`

	Edges []*Edge `json:"edges,omitempty"`

	Nodes []*Node `json:"nodes,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`
}

type GraphTraversal struct {
	Description *Message `json:"description,omitempty"`

	EdgeTraversals []*EdgeTraversal `json:"edgeTraversals,omitempty"`

	ImmutableState map[string]*MultiformatMessageString `json:"immutableState,omitempty"`

	InitialState map[string]*MultiformatMessageString `json:"initialState,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	ResultGraphIndex int `json:"resultGraphIndex,omitempty"`

	RunGraphIndex int `json:"runGraphIndex,omitempty"`
}

type Invocation struct {
	Account string `json:"account,omitempty"`

	Arguments []string `json:"arguments,omitempty"`

	CommandLine string `json:"commandLine,omitempty"`

	EndTimeUtc string `json:"endTimeUtc,omitempty"`

	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty"`

	ExecutableLocation *ArtifactLocation `json:"executableLocation,omitempty"`

	ExecutionSuccessful bool `json:"executionSuccessful"`

	ExitCode int `json:"exitCode,omitempty"`

	ExitCodeDescription string `json:"exitCodeDescription,omitempty"`

	ExitSignalName string `json:"exitSignalName,omitempty"`

	ExitSignalNumber int `json:"exitSignalNumber,omitempty"`

	Machine string `json:"machine,omitempty"`

	NotificationConfigurationOverrides []*ConfigurationOverride `json:"notificationConfigurationOverrides,omitempty"`

	ProcessId int `json:"processId,omitempty"`

	ProcessStartFailureMessage string `json:"processStartFailureMessage,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	ResponseFiles []*ArtifactLocation `json:"responseFiles,omitempty"`

	RuleConfigurationOverrides []*ConfigurationOverride `json:"ruleConfigurationOverrides,omitempty"`

	StartTimeUtc string `json:"startTimeUtc,omitempty"`

	Stderr *ArtifactLocation `json:"stderr,omitempty"`

	Stdin *ArtifactLocation `json:"stdin,omitempty"`

	Stdout *ArtifactLocation `json:"stdout,omitempty"`

	StdoutStderr *ArtifactLocation `json:"stdoutStderr,omitempty"`

	ToolConfigurationNotifications []*Notification `json:"toolConfigurationNotifications,omitempty"`

	ToolExecutionNotifications []*Notification `json:"toolExecutionNotifications,omitempty"`

	WorkingDirectory *ArtifactLocation `json:"workingDirectory,omitempty"`
}

type Location struct {
	Annotations []*Region `json:"annotations,omitempty"`

	Id int `json:"id,omitempty"`

	LogicalLocations []*LogicalLocation `json:"logicalLocations,omitempty"`

	Message *Message `json:"message,omitempty"`

	PhysicalLocation *PhysicalLocation `json:"physicalLocation,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	Relationships []*LocationRelationship `json:"relationships,omitempty"`
}

type LocationRelationship struct {
	Description *Message `json:"description,omitempty"`

	Kinds []string `json:"kinds,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	Target int `json:"target"`
}

type LogicalLocation struct {
	DecoratedName string `json:"decoratedName,omitempty"`

	FullyQualifiedName string `json:"fullyQualifiedName,omitempty"`

	Index int `json:"index,omitempty"`

	Kind string `json:"kind,omitempty"`

	Name string `json:"name,omitempty"`

	ParentIndex int `json:"parentIndex,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`
}

type Message struct {
	Arguments []string `json:"arguments,omitempty"`

	ID string `json:"id,omitempty"`

	Markdown string `json:"markdown,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	Text string `json:"text,omitempty"`
}

type MultiformatMessageString struct {
	Markdown string `json:"markdown,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	Text string `json:"text"`
}

type Node struct {
	Children []*Node `json:"children,omitempty"`

	ID string `json:"id"`

	Label *Message `json:"label,omitempty"`

	Location *Location `json:"location,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`
}

type Notification struct {
	AssociatedRule *ReportingDescriptorReference `json:"associatedRule,omitempty"`

	Descriptor *ReportingDescriptorReference `json:"descriptor,omitempty"`

	Exception *Exception `json:"exception,omitempty"`

	Level interface{} `json:"level,omitempty"`

	Locations []*Location `json:"locations,omitempty"`

	Message *Message `json:"message"`

	Properties *PropertyBag `json:"properties,omitempty"`

	ThreadID int `json:"threadId,omitempty"`

	TimeUtc string `json:"timeUtc,omitempty"`
}

type PhysicalLocation struct {
	Address *Address `json:"address,omitempty"`

	ArtifactLocation *ArtifactLocation `json:"artifactLocation,omitempty"`

	ContextRegion *Region `json:"contextRegion,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	Region *Region `json:"region,omitempty"`
}

type PropertyBag map[string]interface{}

type Rectangle struct {
	Bottom float64 `json:"bottom,omitempty"`

	Left float64 `json:"left,omitempty"`

	Message *Message `json:"message,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	Right float64 `json:"right,omitempty"`

	Top float64 `json:"top,omitempty"`
}

type Region struct {
	ByteLength int `json:"byteLength,omitempty"`

	ByteOffset int `json:"byteOffset,omitempty"`

	CharLength int `json:"charLength,omitempty"`

	CharOffset int `json:"charOffset,omitempty"`

	EndColumn int `json:"endColumn,omitempty"`

	EndLine int `json:"endLine,omitempty"`

	Message *Message `json:"message,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	Snippet *ArtifactContent `json:"snippet,omitempty"`

	SourceLanguage string `json:"sourceLanguage,omitempty"`

	StartColumn int `json:"startColumn,omitempty"`

	StartLine int `json:"startLine,omitempty"`
}

type Replacement struct {
	DeletedRegion *Region `json:"deletedRegion"`

	InsertedContent *ArtifactContent `json:"insertedContent,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`
}

type ReportingConfiguration struct {
	Enabled bool `json:"enabled,omitempty"`

	Level interface{} `json:"level,omitempty"`

	Parameters *PropertyBag `json:"parameters,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	Rank float64 `json:"rank,omitempty"`
}

type ReportingDescriptor struct {
	DefaultConfiguration *ReportingConfiguration `json:"defaultConfiguration,omitempty"`

	DeprecatedGuids []string `json:"deprecatedGuids,omitempty"`

	DeprecatedIds []string `json:"deprecatedIds,omitempty"`

	DeprecatedNames []string `json:"deprecatedNames,omitempty"`

	FullDescription *MultiformatMessageString `json:"fullDescription,omitempty"`

	GUID string `json:"guid,omitempty"`

	Help *MultiformatMessageString `json:"help,omitempty"`

	HelpURI string `json:"helpUri,omitempty"`

	ID string `json:"id"`

	MessageStrings map[string]*MultiformatMessageString `json:"messageStrings,omitempty"`

	Name string `json:"name,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	Relationships []*ReportingDescriptorRelationship `json:"relationships,omitempty"`

	ShortDescription *MultiformatMessageString `json:"shortDescription,omitempty"`
}

type ReportingDescriptorReference struct {
	GUID string `json:"guid,omitempty"`

	ID string `json:"id,omitempty"`

	Index int `json:"index,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	ToolComponent *ToolComponentReference `json:"toolComponent,omitempty"`
}

type ReportingDescriptorRelationship struct {
	Description *Message `json:"description,omitempty"`

	Kinds []string `json:"kinds,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	Target *ReportingDescriptorReference `json:"target"`
}

type Result struct {
	AnalysisTarget *ArtifactLocation `json:"analysisTarget,omitempty"`

	Attachments []*Attachment `json:"attachments,omitempty"`

	BaselineState interface{} `json:"baselineState,omitempty"`

	CodeFlows []*CodeFlow `json:"codeFlows,omitempty"`

	CorrelationGUID string `json:"correlationGuid,omitempty"`

	Fingerprints map[string]string `json:"fingerprints,omitempty"`

	Fixes []*Fix `json:"fixes,omitempty"`

	GraphTraversals []*GraphTraversal `json:"graphTraversals,omitempty"`

	Graphs []*Graph `json:"graphs,omitempty"`

	GUID string `json:"guid,omitempty"`

	HostedViewerURI string `json:"hostedViewerUri,omitempty"`

	Kind interface{} `json:"kind,omitempty"`

	Level interface{} `json:"level,omitempty"`

	Locations []*Location `json:"locations,omitempty"`

	Message *Message `json:"message"`

	OccurrenceCount int `json:"occurrenceCount,omitempty"`

	PartialFingerprints map[string]string `json:"partialFingerprints,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	Provenance *ResultProvenance `json:"provenance,omitempty"`

	Rank float64 `json:"rank,omitempty"`

	RelatedLocations []*Location `json:"relatedLocations,omitempty"`

	Rule *ReportingDescriptorReference `json:"rule,omitempty"`

	RuleID string `json:"ruleId,omitempty"`

	RuleIndex int `json:"ruleIndex,omitempty"`

	Stacks []*Stack `json:"stacks,omitempty"`

	Suppressions []*Suppression `json:"suppressions,omitempty"`

	Taxa []*ReportingDescriptorReference `json:"taxa,omitempty"`

	WebRequest *WebRequest `json:"webRequest,omitempty"`

	WebResponse *WebResponse `json:"webResponse,omitempty"`

	WorkItemUris []string `json:"workItemUris,omitempty"`
}

type ResultProvenance struct {
	ConversionSources []*PhysicalLocation `json:"conversionSources,omitempty"`

	FirstDetectionRunGUID string `json:"firstDetectionRunGuid,omitempty"`

	FirstDetectionTimeUtc string `json:"firstDetectionTimeUtc,omitempty"`

	InvocationIndex int `json:"invocationIndex,omitempty"`

	LastDetectionRunGUID string `json:"lastDetectionRunGuid,omitempty"`

	LastDetectionTimeUtc string `json:"lastDetectionTimeUtc,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`
}

type Run struct {
	Addresses []*Address `json:"addresses,omitempty"`

	Artifacts []*Artifact `json:"artifacts,omitempty"`

	AutomationDetails *RunAutomationDetails `json:"automationDetails,omitempty"`

	BaselineGUID string `json:"baselineGuid,omitempty"`

	ColumnKind interface{} `json:"columnKind,omitempty"`

	Conversion *Conversion `json:"conversion,omitempty"`

	DefaultEncoding string `json:"defaultEncoding,omitempty"`

	DefaultSourceLanguage string `json:"defaultSourceLanguage,omitempty"`

	ExternalPropertyFileReferences *ExternalPropertyFileReferences `json:"externalPropertyFileReferences,omitempty"`

	Graphs []*Graph `json:"graphs,omitempty"`

	Invocations []*Invocation `json:"invocations,omitempty"`

	Language string `json:"language,omitempty"`

	LogicalLocations []*LogicalLocation `json:"logicalLocations,omitempty"`

	NewlineSequences []string `json:"newlineSequences,omitempty"`

	OriginalUriBaseIds map[string]*ArtifactLocation `json:"originalUriBaseIds,omitempty"`

	Policies []*ToolComponent `json:"policies,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	RedactionTokens []string `json:"redactionTokens,omitempty"`

	Results []*Result `json:"results"`

	RunAggregates []*RunAutomationDetails `json:"runAggregates,omitempty"`

	SpecialLocations *SpecialLocations `json:"specialLocations,omitempty"`

	Taxonomies []*ToolComponent `json:"taxonomies,omitempty"`

	ThreadFlowLocations []*ThreadFlowLocation `json:"threadFlowLocations,omitempty"`

	Tool *Tool `json:"tool"`

	Translations []*ToolComponent `json:"translations,omitempty"`

	VersionControlProvenance []*VersionControlDetails `json:"versionControlProvenance,omitempty"`

	WebRequests []*WebRequest `json:"webRequests,omitempty"`

	WebResponses []*WebResponse `json:"webResponses,omitempty"`
}

type RunAutomationDetails struct {
	CorrelationGUID string `json:"correlationGuid,omitempty"`

	Description *Message `json:"description,omitempty"`

	GUID string `json:"guid,omitempty"`

	ID string `json:"id,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`
}

type SpecialLocations struct {
	DisplayBase *ArtifactLocation `json:"displayBase,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`
}

type Stack struct {
	Frames []*StackFrame `json:"frames"`

	Message *Message `json:"message,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`
}

type StackFrame struct {
	Location *Location `json:"location,omitempty"`

	Module string `json:"module,omitempty"`

	Parameters []string `json:"parameters,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	ThreadID int `json:"threadId,omitempty"`
}

type Report struct {
	InlineExternalProperties []*ExternalProperties `json:"inlineExternalProperties,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	Runs []*Run `json:"runs"`

	Schema string `json:"$schema,omitempty"`

	Version interface{} `json:"version"`
}

type Suppression struct {
	GUID string `json:"guid,omitempty"`

	Justification string `json:"justification,omitempty"`

	Kind interface{} `json:"kind"`

	Location *Location `json:"location,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	Status interface{} `json:"status,omitempty"`
}

type ThreadFlow struct {
	ID string `json:"id,omitempty"`

	ImmutableState map[string]*MultiformatMessageString `json:"immutableState,omitempty"`

	InitialState map[string]*MultiformatMessageString `json:"initialState,omitempty"`

	Locations []*ThreadFlowLocation `json:"locations"`

	Message *Message `json:"message,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`
}

type ThreadFlowLocation struct {
	ExecutionOrder int `json:"executionOrder,omitempty"`

	ExecutionTimeUtc string `json:"executionTimeUtc,omitempty"`

	Importance interface{} `json:"importance,omitempty"`

	Index int `json:"index,omitempty"`

	Kinds []string `json:"kinds,omitempty"`

	Location *Location `json:"location,omitempty"`

	Module string `json:"module,omitempty"`

	NestingLevel int `json:"nestingLevel,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	Stack *Stack `json:"stack,omitempty"`

	State map[string]*MultiformatMessageString `json:"state,omitempty"`

	Taxa []*ReportingDescriptorReference `json:"taxa,omitempty"`

	WebRequest *WebRequest `json:"webRequest,omitempty"`

	WebResponse *WebResponse `json:"webResponse,omitempty"`
}

type Tool struct {
	Driver *ToolComponent `json:"driver"`

	Extensions []*ToolComponent `json:"extensions,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`
}

type ToolComponent struct {
	AssociatedComponent *ToolComponentReference `json:"associatedComponent,omitempty"`

	Contents []interface{} `json:"contents,omitempty"`

	DottedQuadFileVersion string `json:"dottedQuadFileVersion,omitempty"`

	DownloadURI string `json:"downloadUri,omitempty"`

	FullDescription *MultiformatMessageString `json:"fullDescription,omitempty"`

	FullName string `json:"fullName,omitempty"`

	GlobalMessageStrings map[string]*MultiformatMessageString `json:"globalMessageStrings,omitempty"`

	GUID string `json:"guid,omitempty"`

	InformationURI string `json:"informationUri,omitempty"`

	IsComprehensive bool `json:"isComprehensive,omitempty"`

	Language string `json:"language,omitempty"`

	LocalizedDataSemanticVersion string `json:"localizedDataSemanticVersion,omitempty"`

	Locations []*ArtifactLocation `json:"locations,omitempty"`

	MinimumRequiredLocalizedDataSemanticVersion string `json:"minimumRequiredLocalizedDataSemanticVersion,omitempty"`

	Name string `json:"name"`

	Notifications []*ReportingDescriptor `json:"notifications,omitempty"`

	Organization string `json:"organization,omitempty"`

	Product string `json:"product,omitempty"`

	ProductSuite string `json:"productSuite,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	ReleaseDateUtc string `json:"releaseDateUtc,omitempty"`

	Rules []*ReportingDescriptor `json:"rules,omitempty"`

	SemanticVersion string `json:"semanticVersion,omitempty"`

	ShortDescription *MultiformatMessageString `json:"shortDescription,omitempty"`

	SupportedTaxonomies []*ToolComponentReference `json:"supportedTaxonomies,omitempty"`

	Taxa []*ReportingDescriptor `json:"taxa,omitempty"`

	TranslationMetadata *TranslationMetadata `json:"translationMetadata,omitempty"`

	Version string `json:"version,omitempty"`
}

type ToolComponentReference struct {
	GUID string `json:"guid,omitempty"`

	Index int `json:"index,omitempty"`

	Name string `json:"name,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`
}

type TranslationMetadata struct {
	DownloadURI string `json:"downloadUri,omitempty"`

	FullDescription *MultiformatMessageString `json:"fullDescription,omitempty"`

	FullName string `json:"fullName,omitempty"`

	InformationURI string `json:"informationUri,omitempty"`

	Name string `json:"name"`

	Properties *PropertyBag `json:"properties,omitempty"`

	ShortDescription *MultiformatMessageString `json:"shortDescription,omitempty"`
}

type VersionControlDetails struct {
	AsOfTimeUtc string `json:"asOfTimeUtc,omitempty"`

	Branch string `json:"branch,omitempty"`

	MappedTo *ArtifactLocation `json:"mappedTo,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	RepositoryURI string `json:"repositoryUri"`

	RevisionID string `json:"revisionId,omitempty"`

	RevisionTag string `json:"revisionTag,omitempty"`
}

type WebRequest struct {
	Body *ArtifactContent `json:"body,omitempty"`

	Headers map[string]string `json:"headers,omitempty"`

	Index int `json:"index,omitempty"`

	Method string `json:"method,omitempty"`

	Parameters map[string]string `json:"parameters,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	Protocol string `json:"protocol,omitempty"`

	Target string `json:"target,omitempty"`

	Version string `json:"version,omitempty"`
}

type WebResponse struct {
	Body *ArtifactContent `json:"body,omitempty"`

	Headers map[string]string `json:"headers,omitempty"`

	Index int `json:"index,omitempty"`

	NoResponseReceived bool `json:"noResponseReceived,omitempty"`

	Properties *PropertyBag `json:"properties,omitempty"`

	Protocol string `json:"protocol,omitempty"`

	ReasonPhrase string `json:"reasonPhrase,omitempty"`

	StatusCode int `json:"statusCode,omitempty"`

	Version string `json:"version,omitempty"`
}
