package anaml

// Entity ..
type Entity struct {
	ID            int          `json:"id,omitempty"`
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	Type          string       `json:"adt_type"`
	DefaultColumn *string      `json:"defaultColumn,omitempty"`
	RequiredType  *interface{} `json:"requiredType,omitempty"`
	Entities      *[]int       `json:"entities,omitempty"`
	Labels        []string     `json:"labels"`
	Attributes    []Attribute  `json:"attributes"`
}

// EntityMapping ..
type EntityMapping struct {
	ID        int   `json:"id,omitempty"`
	From      int   `json:"from"`
	To        int   `json:"to"`
	Mapping   int   `json:"mapping"`
	OneToMany *bool `json:"oneToMany,omitempty"`
}

// EntityPopulation ..
type EntityPopulation struct {
	ID          int         `json:"id,omitempty"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Labels      []string    `json:"labels"`
	Attributes  []Attribute `json:"attributes"`
	Entity      int         `json:"entity"`
	Sources     []int       `json:"sources"`
	Expression  string      `json:"expression"`
}

// TimestampInfo ..
type TimestampInfo struct {
	Column string `json:"timestampColumn"`
	Zone   string `json:"timezone,omitempty"`
}

// EventDescription ..
type EventDescription struct {
	Entities      map[string]string `json:"entities"`
	TimestampInfo *TimestampInfo    `json:"timestampInfo"`
}

// Table ...
type Table struct {
	ID            int               `json:"id,omitempty"`
	Name          string            `json:"name"`
	Description   string            `json:"description"`
	Type          string            `json:"adt_type"`
	Sources       []int             `json:"sources,omitempty"`
	Source        *SourceReference  `json:"source,omitempty"`
	Expression    string            `json:"expression,omitempty"`
	EventInfo     *EventDescription `json:"eventDescription,omitempty"`
	EntityMapping int               `json:"entityMapping,omitempty"`
	ExtraFeatures []int             `json:"extraFeatures,omitempty"`
	Labels        []string          `json:"labels"`
	Attributes    []Attribute       `json:"attributes"`
}

// EventWindow ...
type EventWindow struct {
	Type   string `json:"adt_type"`
	Days   int    `json:"days,omitempty"`
	Hours  int    `json:"hours,omitempty"`
	Months int    `json:"months,omitempty"`
	Rows   int    `json:"rows,omitempty"`
}

// SQLExpression ...
type SQLExpression struct {
	SQL string `json:"sql"`
}

// AggregateExpression ...
type AggregateExpression struct {
	Type string `json:"adt_type"`
}

// Feature ... again, completely normalised.
// Note
// Go is a bad language, We can't use omitempty for over, because both [] and 'nil' are empty.
// Empty list is appropriate, especially for templates. But unfortunately, we will be sending
// a really dumb `null` where it doesn't make sense to do so.
type Feature struct {
	ID          int                  `json:"id,omitempty"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Type        string               `json:"adt_type"`
	Table       int                  `json:"table,omitempty"`
	Window      *EventWindow         `json:"window,omitempty"`
	Select      SQLExpression        `json:"select"`
	Filter      *SQLExpression       `json:"filter"`
	Aggregate   *AggregateExpression `json:"aggregate,omitempty"`
	PostAggExpr *SQLExpression       `json:"postAggregateExpr,omitempty"`
	EntityRestr *[]int               `json:"entityRestrictions,omitempty"`
	Over        []int                `json:"over"`
	EntityID    int                  `json:"entityId,omitempty"`
	TemplateID  *int                 `json:"template,omitempty"`
	Labels      []string             `json:"labels"`
	Attributes  []Attribute          `json:"attributes"`
}

// FeatureTemplate ... again, completely normalised.
type FeatureTemplate struct {
	ID          int                  `json:"id,omitempty"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Type        string               `json:"adt_type"`
	Table       int                  `json:"table"`
	Window      *EventWindow         `json:"window,omitempty"`
	Select      SQLExpression        `json:"select"`
	Filter      *SQLExpression       `json:"filter"`
	Aggregate   *AggregateExpression `json:"aggregate,omitempty"`
	PostAggExpr *SQLExpression       `json:"postAggregateExpr"`
	EntityRestr *[]int               `json:"entityRestrictions,omitempty"`
	Over        []int                `json:"over"`
	EntityID    int                  `json:"entityId,omitempty"`
	Labels      []string             `json:"labels"`
	Attributes  []Attribute          `json:"attributes"`
}

// FeatureSet ...
type FeatureSet struct {
	ID          int         `json:"id,omitempty"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	EntityID    int         `json:"entity,omitempty"`
	Features    []int       `json:"features"`
	Labels      []string    `json:"labels"`
	Attributes  []Attribute `json:"attributes"`
}

// VersionTarget ...
type VersionTarget struct {
	Type   string  `json:"adt_type"`
	Commit *string `json:"commitId,omitempty"`
	Branch *string `json:"branchName,omitempty"`
}

// FeatureStore ...
type FeatureStore struct {
	ID                        int                    `json:"id,omitempty"`
	Type                      string                 `json:"adt_type"`
	Name                      string                 `json:"name"`
	Description               string                 `json:"description"`
	Labels                    []string               `json:"labels"`
	Attributes                []Attribute            `json:"attributes"`
	FeatureSet                int                    `json:"featureSet"`
	Enabled                   bool                   `json:"enabled"`
	Schedule                  *Schedule              `json:"schedule"`
	Destinations              []DestinationReference `json:"destinations"`
	Cluster                   int                    `json:"cluster"`
	ClusterPropertySets       []int                  `json:"clusterPropertySets"`
	AdditionalSparkProperties map[string]string      `json:"additionalSparkProperties,omitempty"`
	RunDateOffset             *int                   `json:"runDateOffset,omitempty"`
	Principal                 *int                   `json:"principal,omitempty"`
	Population                *int                   `json:"entityPopulation,omitempty"`
	StartDate                 *string                `json:"startDate,omitempty"`
	EndDate                   *string                `json:"endDate,omitempty"`
	Table                     *int                   `json:"table,omitempty"`
	IncludeMetadata           bool                   `json:"includeMetadata"`
	VersionTarget             *VersionTarget         `json:"versionTarget,omitempty"`
}

type Schedule struct {
	Type           string       `json:"adt_type"`
	StartTimeOfDay *string      `json:"startTimeOfDay,omitempty"`
	CronString     string       `json:"cronString,omitempty"`
	RetryPolicy    *RetryPolicy `json:"retryPolicy,omitempty"`
}

type RetryPolicy struct {
	Type        string `json:"adt_type"`
	Backoff     string `json:"backoff,omitempty"`
	MaxAttempts int    `json:"maxAttempts,omitempty"`
}

type SensitiveAttribute struct {
	Key         string             `json:"key"`
	ValueConfig *SecretValueConfig `json:"valueConfig"`
}

type SecretValueConfig struct {
	Type          string `json:"adt_type"`
	Secret        string `json:"secret,omitempty"`
	FilePath      string `json:"filepath,omitempty"`
	SecretProject string `json:"secretProject,omitempty"`
	SecretId      string `json:"secretId,omitempty"`
}

type ViewMaterialisationSpec struct {
	Table       int                  `json:"table"`
	Destination DestinationReference `json:"destination"`
}

// ViewMaterialisation ...
type ViewMaterialisationJob struct {
	ID                        int                       `json:"id,omitempty"`
	Type                      string                    `json:"adt_type"`
	Name                      string                    `json:"name"`
	Description               string                    `json:"description"`
	Labels                    []string                  `json:"labels"`
	Attributes                []Attribute               `json:"attributes"`
	Views                     []ViewMaterialisationSpec `json:"views"`
	Schedule                  *Schedule                 `json:"schedule"`
	Cluster                   int                       `json:"cluster"`
	ClusterPropertySets       []int                     `json:"clusterPropertySets"`
	AdditionalSparkProperties map[string]string         `json:"additionalSparkProperties,omitempty"`
	Principal                 *int                      `json:"principal,omitempty"`
	UsageTTL                  *string                   `json:"usageTTL,omitempty"`
	IncludeMetadata           bool                      `json:"includeMetadata"`
	VersionTarget             *VersionTarget            `json:"versionTarget,omitempty"`
}

// Source ...
type Source struct {
	ID                  int                             `json:"id,omitempty"`
	Name                string                          `json:"name"`
	Description         string                          `json:"description"`
	Type                string                          `json:"adt_type"`
	Bucket              string                          `json:"bucket,omitempty"`
	Path                string                          `json:"path,omitempty"`
	FileFormat          *FileFormat                     `json:"fileFormat,omitempty"`
	Endpoint            string                          `json:"endpoint,omitempty"`
	AccessKey           string                          `json:"accessKey,omitempty"`
	SecretKey           string                          `json:"secretKey,omitempty"`
	URL                 string                          `json:"url,omitempty"`
	Schema              string                          `json:"schema,omitempty"`
	CredentialsProvider *LoginCredentialsProviderConfig `json:"credentialsProvider,omitempty"`
	Database            string                          `json:"database,omitempty"`
	BootstrapServers    string                          `json:"bootstrapServers,omitempty"`
	SchemaRegistryURL   string                          `json:"schemaRegistryUrl,omitempty"`
	KafkaProperties     []SensitiveAttribute            `json:"kafkaPropertiesProviders"`
	Labels              []string                        `json:"labels"`
	Attributes          []Attribute                     `json:"attributes"`
	Warehouse           string                          `json:"warehouse,omitempty"`
	AccessRules         []AccessRule                    `json:"accessRules"`
}

type FileFormat struct {
	Type                     string  `json:"adt_type"`
	Sep                      *string `json:"sep,omitempty"`
	QuoteAll                 *bool   `json:"quoteAll,omitempty"`
	IncludeHeader            *bool   `json:"includeHeader,omitempty"`
	EmptyValue               *string `json:"emptyValue,omitempty"`
	Compression              *string `json:"compression,omitempty"`
	DateFormat               *string `json:"dateFormat,omitempty"`
	TimestampFormat          *string `json:"timestampFormat,omitempty"`
	IgnoreLeadingWhiteSpace  *bool   `json:"ignoreLeadingWhiteSpace,omitempty"`
	IgnoreTrailingWhiteSpace *bool   `json:"ignoreTrailingWhiteSpace,omitempty"`
	LineSep                  *string `json:"lineSep,omitempty"`
}

type KafkaFormat struct {
	Type string `json:"adt_type"`
}

// SourceReference ...
// Doubles up with EventStoreReference as the json name is
// the same and we normalised everything.
type SourceReference struct {
	Type      string `json:"adt_type,omitempty"`
	SourceID  int    `json:"sourceId,omitempty"`
	Folder    string `json:"folder,omitempty"`
	TableName string `json:"tableName,omitempty"`
	Topic     string `json:"topic,omitempty"`

	EventStoreId int `json:"eventStoreId,omitempty"`
	Entity       int `json:"entity,omitempty"`
}

// AccessRule ...
type AccessRule struct {
	Resource     string        `json:"resource"`
	Principals   []PrincipalId `json:"principals"`
	MaskingRules []MaskingRule `json:"maskingRules"`
}

// MaskingRule ...
type MaskingRule struct {
	Type       string `json:"adt_type"`
	Expression string `json:"expression"`
	Column     string `json:"column,omitempty"`
}

// Destination ...
type Destination struct {
	ID                  int                             `json:"id,omitempty"`
	Name                string                          `json:"name"`
	Description         string                          `json:"description"`
	Labels              []string                        `json:"labels"`
	Attributes          []Attribute                     `json:"attributes"`
	Type                string                          `json:"adt_type"`
	Bucket              string                          `json:"bucket,omitempty"`
	Path                string                          `json:"path,omitempty"`
	FileFormat          *FileFormat                     `json:"fileFormat,omitempty"`
	Endpoint            string                          `json:"endpoint,omitempty"`
	AccessKey           string                          `json:"accessKey,omitempty"`
	SecretKey           string                          `json:"secretKey,omitempty"`
	URL                 string                          `json:"url,omitempty"`
	Schema              string                          `json:"schema,omitempty"`
	CredentialsProvider *LoginCredentialsProviderConfig `json:"credentialsProvider,omitempty"`
	Database            string                          `json:"database,omitempty"`
	BootstrapServers    string                          `json:"bootstrapServers,omitempty"`
	SchemaRegistryURL   string                          `json:"schemaRegistryUrl,omitempty"`
	KafkaProperties     []SensitiveAttribute            `json:"kafkaPropertiesProviders"`
	StagingArea         *GCSStagingArea                 `json:"stagingArea,omitempty"`
	Warehouse           string                          `json:"warehouse,omitempty"`
	Project             string                          `json:"project,omitempty"`
	Instance            string                          `json:"instance,omitempty"`
}

// GCSStagingArea ...
type GCSStagingArea struct {
	Type   string `json:"adt_type"`
	Bucket string `json:"bucket"`
	Path   string `json:"path,omitempty"`
}

// DestinationReference ...
type DestinationReference struct {
	Type                      string       `json:"adt_type"`
	DestinationID             int          `json:"destinationId"`
	Folder                    string       `json:"folder,omitempty"`
	FolderPartitioningEnabled *bool        `json:"folderPartitioningEnabled,omitempty"`
	TableName                 string       `json:"tableName,omitempty"`
	Topic                     string       `json:"topic,omitempty"`
	Format                    *KafkaFormat `json:"format,omitempty"`
	Mode                      string       `json:"saveMode,omitempty"`
	Options                   []Attribute  `json:"options,omitempty"`
}

// Cluster ...
type Cluster struct {
	ID                  int                             `json:"id,omitempty"`
	Name                string                          `json:"name"`
	Description         string                          `json:"description"`
	Type                string                          `json:"adt_type"`
	IsPreviewCluster    bool                            `json:"isPreviewCluster"`
	AnamlServerURL      string                          `json:"anamlServerUrl,omitempty"`
	SparkServerURL      string                          `json:"sparkServerUrl,omitempty"`
	CredentialsProvider *LoginCredentialsProviderConfig `json:"credentialsProvider,omitempty"`
	SparkConfig         *SparkConfig                    `json:"sparkConfig,omitempty"`
	PropertySet         []PropertySet                   `json:"propertySets"`
	Labels              []string                        `json:"labels"`
	Attributes          []Attribute                     `json:"attributes"`
}

// LoginCredentialsProviderConfig  ...
type LoginCredentialsProviderConfig struct {
	Type                  string `json:"adt_type"`
	Username              string `json:"username"`
	Password              string `json:"password,omitempty"`
	FilePath              string `json:"filepath,omitempty"`
	PasswordSecretProject string `json:"passwordSecretProject,omitempty"`
	PasswordSecretId      string `json:"passwordSecretId,omitempty"`
}

// SparkConfig ...
type SparkConfig struct {
	EnableHiveSupport         bool              `json:"enableHiveSupport"`
	HiveMetastoreURL          string            `json:"hiveMetastoreUrl,omitempty"`
	AdditionalSparkProperties map[string]string `json:"additionalSparkProperties"`
}

// PropertySet ...
type PropertySet struct {
	ID                        *int              `json:"id,omitempty"`
	Name                      string            `json:"name"`
	AdditionalSparkProperties map[string]string `json:"additionalSparkProperties"`
}

// User ...
type Role struct {
	Type string `json:"adt_type"`
}

type User struct {
	ID        int     `json:"id,omitempty"`
	Name      string  `json:"name"`
	Email     *string `json:"email,omitempty"`
	GivenName *string `json:"givenName,omitempty"`
	Surname   *string `json:"surname,omitempty"`
	Password  *string `json:"password,omitempty"`
	Roles     []Role  `json:"roles"`
}

// Access token and creation request.
type AccessToken struct {
	ID          string `json:"id,omitempty"`
	Secret      string `json:"secret,omitempty"`
	Owner       *int   `json:"owner,omitempty"`
	Description string `json:"description,omitempty"`
	Roles       []Role `json:"roles"`
}

type ChangeOtherPasswordRequest struct {
	Password *string `json:"password"`
}

type UserGroupMemberSource struct {
	Type string `json:"adt_type"`
}

type UserGroupMember struct {
	UserID int                   `json:"userId,omitempty"`
	Source UserGroupMemberSource `json:"source"`
}

// UserGroup ..
type UserGroup struct {
	ID              int               `json:"id,omitempty"`
	Name            string            `json:"name"`
	Description     string            `json:"description"`
	Roles           []Role            `json:"roles"`
	Members         []UserGroupMember `json:"members"`
	ExternalGroupID *string           `json:"externalGroupId,omitempty"`
}

// BranchProtection
type BranchProtection struct {
	ID                  int            `json:"id,omitempty"`
	ProtectionPattern   string         `json:"protectionPattern"`
	MergeApprovalRules  []ApprovalRule `json:"mergeApprovalRules"`
	PushWhitelist       []PrincipalId  `json:"pushWhitelist"`
	ApplyToAdmins       bool           `json:"applyToAdmins"`
	AllowBranchDeletion bool           `json:"allowBranchDeletion"`
}

// ApprovalRule
type ApprovalRule struct {
	Approvers            []PrincipalId `json:"approvers,omitempty"`
	NumRequiredApprovals int           `json:"numRequiredApprovals"`
	Type                 string        `json:"adt_type"`
}

// PrincipalId
type PrincipalId struct {
	ID   int    `json:"id"`
	Type string `json:"adt_type"`
}

type MonitoringPlan struct {
	Type     string `json:"adt_type"`
	Tables   []int  `json:"tables,omitempty"`
	Excluded []int  `json:"excluded"`
}

// TableMonitoring ...
type TableMonitoring struct {
	ID                  int             `json:"id,omitempty"`
	Name                string          `json:"name"`
	Description         string          `json:"description"`
	Plan                *MonitoringPlan `json:"plan"`
	Schedule            *Schedule       `json:"schedule"`
	Cluster             int             `json:"cluster"`
	ClusterPropertySets []int           `json:"clusterPropertySets"`
	Enabled             bool            `json:"enabled"`
	Principal           *int            `json:"principal,omitempty"`
}

type CachingPlan struct {
	Type     string             `json:"adt_type"`
	Specs    []TableCachingSpec `json:"specs"`
	Excluded []TableCachingSpec `json:"excluded"`
}

// TableCaching ...
type TableCaching struct {
	ID                  int          `json:"id,omitempty"`
	Name                string       `json:"name"`
	Description         string       `json:"description"`
	Principal           *int         `json:"principal,omitempty"`
	Plan                *CachingPlan `json:"plan"`
	Retainement         *string      `json:"retainment"`
	PrefixURI           string       `json:"prefixURI"`
	Schedule            *Schedule    `json:"schedule"`
	Cluster             int          `json:"cluster"`
	ClusterPropertySets []int        `json:"clusterPropertySets"`
}

type TableCachingSpec struct {
	Table  int `json:"table"`
	Entity int `json:"entity"`
}

type Webhook struct {
	ID                   int       `json:"id,omitempty"`
	Name                 string    `json:"name"`
	Description          string    `json:"description"`
	URL                  string    `json:"url"`
	MergeRequests        *struct{} `json:"mergeRequests,omitempty"`
	MergeRequestComments *struct{} `json:"mergeRequestComments,omitempty"`
	Commits              *struct{} `json:"commits,omitempty"`
	FeatureStoreRuns     *struct{} `json:"featureStoreRuns,omitempty"`
	MonitoringRuns       *struct{} `json:"monitoringRuns,omitempty"`
	CachingRuns          *struct{} `json:"cachingRuns,omitempty"`
	MaterialisationRuns  *struct{} `json:"materialisationRuns,omitempty"`
	EventStoreRuns       *struct{} `json:"eventStoreRuns,omitempty"`
}

type Attribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// LabelRestriction ...
type LabelRestriction struct {
	ID     int     `json:"id,omitempty"`
	Text   string  `json:"text"`
	Emoji  *string `json:"emoji,omitempty"`
	Colour *string `json:"colour,omitempty"`
}

// AttributeRestriction ...
type AttributeRestriction struct {
	ID           int                    `json:"id,omitempty"`
	Key          string                 `json:"key"`
	Description  string                 `json:"description"`
	Type         string                 `json:"adt_type"`
	Mandatory    bool                   `json:"mandatory"`
	DefaultValue *string                `json:"defaultValue,omitempty"`
	Choices      *[]EnumAttributeChoice `json:"choices,omitempty"`
	AppliesTo    []AttributeTarget      `json:"appliesTo"`
}

type AttributeTarget struct {
	Type string `json:"adt_type"`
}

type EnumAttributeChoice struct {
	Value   string                `json:"value"`
	Display *EnumAttributeDisplay `json:"display,omitempty"`
}

type EnumAttributeDisplay struct {
	Emoji  string `json:"emoji,omitempty"`
	Colour string `json:"colour,omitempty"`
}

// EventDescription ..
type EventStoreTopicColumns struct {
	Entity        string         `json:"entity"`
	TimestampInfo *TimestampInfo `json:"timestampInfo"`
	HasStreaming  bool           `json:"hasStreaming"`
}

type EventStore struct {
	ID                  int                               `json:"id,omitempty"`
	Name                string                            `json:"name"`
	Description         string                            `json:"description"`
	Labels              []string                          `json:"labels"`
	Attributes          []Attribute                       `json:"attributes"`
	BootstrapServers    string                            `json:"bootstrapServers"`
	SchemaRegistryURL   string                            `json:"schemaRegistryUrl"`
	KafkaProperties     []SensitiveAttribute              `json:"kafkaPropertiesProviders"`
	Ingestions          map[string]EventStoreTopicColumns `json:"ingestions"`
	ConnectBaseURI      *string                           `json:"connectBaseURI"`
	BatchIngestBaseURI  *string                           `json:"batchIngestBaseURI"`
	ScatterBaseURI      string                            `json:"scatterBaseURI"`
	GlacierBaseURI      string                            `json:"glacierBaseURI"`
	Schedule            *Schedule                         `json:"schedule"`
	Cluster             int                               `json:"cluster"`
	ClusterPropertySets []int                             `json:"clusterPropertySets"`
	AccessRules         []AccessRule                      `json:"accessRules"`
}

func validRoles() []string {
	return []string{
		"admin_attributes",
		"admin_branch_perms",
		"admin_groups",
		"admin_projects",
		"admin_schedules",
		"admin_system",
		"admin_users",
		"admin_webhooks",
		"author",
		"edit_projects",
		"run_caching",
		"run_event_store",
		"run_featuregen",
		"run_monitoring",
		"super_user",
		"view_reports",
	}
}

func mapRolesToBackend(frontend []string) []Role {
	vs := make([]Role, 0, len(frontend))
	for _, v := range frontend {
		if v == "admin_attributes" {
			vs = append(vs, Role{"adminattributes"})
		} else if v == "admin_branch_perms" {
			vs = append(vs, Role{"adminbranchperms"})
		} else if v == "admin_groups" {
			vs = append(vs, Role{"admingroups"})
		} else if v == "admin_projects" {
			vs = append(vs, Role{"adminprojects"})
		} else if v == "admin_schedules" {
			vs = append(vs, Role{"adminschedules"})
		} else if v == "admin_system" {
			vs = append(vs, Role{"adminsystem"})
		} else if v == "admin_users" {
			vs = append(vs, Role{"adminusers"})
		} else if v == "admin_webhooks" {
			vs = append(vs, Role{"adminwebhooks"})
		} else if v == "author" {
			vs = append(vs, Role{"author"})
		} else if v == "edit_projects" {
			vs = append(vs, Role{"editprojects"})
		} else if v == "run_caching" {
			vs = append(vs, Role{"runcaching"})
		} else if v == "run_event_store" {
			vs = append(vs, Role{"runeventstore"})
		} else if v == "run_featuregen" {
			vs = append(vs, Role{"runfeaturegen"})
		} else if v == "run_monitoring" {
			vs = append(vs, Role{"runmonitoring"})
		} else if v == "super_user" {
			vs = append(vs, Role{"superuser"})
		} else if v == "view_reports" {
			vs = append(vs, Role{"viewreports"})
		}
		// TODO: We should raise an error if we fall through the cases.
	}
	return vs
}

func mapRolesToFrontend(backend []Role) []string {
	vs := make([]string, 0, len(backend))
	for _, v := range backend {
		if v.Type == "adminattributes" {
			vs = append(vs, "admin_attributes")
		} else if v.Type == "adminbranchperms" {
			vs = append(vs, "admin_branch_perms")
		} else if v.Type == "admingroups" {
			vs = append(vs, "admin_groups")
		} else if v.Type == "adminprojects" {
			vs = append(vs, "admin_projects")
		} else if v.Type == "adminschedules" {
			vs = append(vs, "admin_schedules")
		} else if v.Type == "adminsystem" {
			vs = append(vs, "admin_system")
		} else if v.Type == "adminusers" {
			vs = append(vs, "admin_users")
		} else if v.Type == "adminwebhooks" {
			vs = append(vs, "admin_webhooks")
		} else if v.Type == "author" {
			vs = append(vs, "author")
		} else if v.Type == "editprojects" {
			vs = append(vs, "edit_projects")
		} else if v.Type == "runcaching" {
			vs = append(vs, "run_caching")
		} else if v.Type == "runevent_store" {
			vs = append(vs, "run_event_store")
		} else if v.Type == "runfeaturegen" {
			vs = append(vs, "run_featuregen")
		} else if v.Type == "runmonitoring" {
			vs = append(vs, "run_monitoring")
		} else if v.Type == "superuser" {
			vs = append(vs, "super_user")
		} else if v.Type == "viewreports" {
			vs = append(vs, "view_reports")
		}
		// TODO: We should raise an error if we fall through the cases.
	}
	return vs
}

func validGroupMemberSource() []string {
	return []string{
		"anaml", "external",
	}
}
