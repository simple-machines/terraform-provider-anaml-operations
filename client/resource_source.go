package anaml

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const sourceDescription = `# Sources

A Source is the physical configuration for the location of root tables.
Sources are therefore specific to the underlying storage technology.

Multiple different types of sources are supported:

- Amazon S3
- Google Cloud Storage
- Google BigQuery
- Hive
- HDFS
- JDBC
`

func ResourceSource() *schema.Resource {
	return &schema.Resource{
		Description: sourceDescription,
		Create:      resourceSourceCreate,
		Read:        resourceSourceRead,
		Update:      resourceSourceUpdate,
		Delete:      resourceSourceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAnamlName(),
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"s3": {
				Type:         schema.TypeList,
				Optional:     true,
				MaxItems:     1,
				Elem:         s3SourceDestinationSchema(),
				ExactlyOneOf: []string{"s3", "s3a", "jdbc", "hive", "big_query", "gcs", "local", "hdfs", "kafka", "snowflake"},
			},
			"s3a": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     s3aSourceDestinationSchema(),
			},
			"jdbc": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     jdbcSourceDestinationSchema(),
			},
			"hive": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     hiveSourceDestinationSchema(),
			},
			"big_query": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     bigQuerySourceSchema(),
			},
			"gcs": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     gcsSourceDestinationSchema(),
			},
			"local": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     localSourceDestinationSchema(),
			},
			"hdfs": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     hdfsSourceDestinationSchema(),
			},
			"kafka": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     kafkaSourceDestinationSchema(),
			},
			"snowflake": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     snowflakeSourceDestinationSchema(),
			},
			"labels": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Labels to attach to the object",
				Elem:        labelSchema(),
			},
			"attribute": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Attributes (key value pairs) to attach to the object",
				Elem:        attributeSchema(),
			},
			"access_rule": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Access rules to attach to the object",
				Elem:        accessRuleSchema(),
			},
		},
	}
}

func s3SourceDestinationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bucket": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
			"path": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
			"file_format": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateFileFormat(),
			},
			"field_separator": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"quote_all": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"include_header": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"empty_value": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ignore_leading_whitespace": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ignore_trailing_whitespace": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"compression": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"date_format": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"timestamp_format": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"line_separator": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func s3aSourceDestinationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bucket": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
			"path": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
			"endpoint": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
			"access_key": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
			"secret_key": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
			"file_format": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateFileFormat(),
			},
			"field_separator": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"quote_all": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"include_header": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"empty_value": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ignore_leading_whitespace": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ignore_trailing_whitespace": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"compression": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"date_format": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"timestamp_format": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"line_separator": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func jdbcSourceDestinationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
			"schema": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
			"credentials_provider": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     loginCredentialsProviderConfigSchema(),
			},
		},
	}
}

func hiveSourceDestinationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"database": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
		},
	}
}

func bigQuerySourceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"path": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
		},
	}
}

func gcsSourceDestinationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bucket": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
			"path": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
			"file_format": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateFileFormat(),
			},
			"field_separator": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"quote_all": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"include_header": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"empty_value": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ignore_leading_whitespace": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ignore_trailing_whitespace": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"compression": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"date_format": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"timestamp_format": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"line_separator": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func localSourceDestinationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"path": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
			"file_format": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateFileFormat(),
			},
			"field_separator": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"quote_all": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"include_header": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"empty_value": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ignore_leading_whitespace": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ignore_trailing_whitespace": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"compression": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"date_format": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"timestamp_format": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"line_separator": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func hdfsSourceDestinationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"path": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
			"file_format": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateFileFormat(),
			},
			"field_separator": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"quote_all": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"include_header": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"empty_value": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ignore_leading_whitespace": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ignore_trailing_whitespace": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"compression": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"date_format": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"timestamp_format": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"line_separator": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func kafkaSourceDestinationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"bootstrap_servers": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
			"schema_registry_url": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
			"property": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     sensitiveAttributeSchema(),
			},
		},
	}
}

func onlineDestinationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
			"schema": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
			"credentials_provider": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     loginCredentialsProviderConfigSchema(),
			},
		},
	}
}

func bigtableDestinationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"project": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
			"instance": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
		},
	}
}

func snowflakeSourceDestinationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
			"warehouse": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
			"database": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
			"schema": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
			"credentials_provider": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     loginCredentialsProviderConfigSchema(),
			},
		},
	}
}

func accessRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"resource": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
			"principals": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     principalIdSchema(),
			},
			"masking_rule": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     maskingRuleSchema(),
			},
		},
	}
}

func maskingRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"filter": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     filterMaskingRuleSchema(),
			},
			"mask": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem:     maskMaskingRuleSchema(),
			},
		},
	}
}

func filterMaskingRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"expression": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
		},
	}
}

func maskMaskingRuleSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"column": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
			"expression": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
			},
		},
	}
}

func resourceSourceRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*Client)
	sourceID := d.Id()

	source, err := c.GetSource(sourceID)
	if err != nil {
		return err
	}
	if source == nil {
		d.SetId("")
		return nil
	}

	if err := d.Set("name", source.Name); err != nil {
		return err
	}
	if err := d.Set("description", source.Description); err != nil {
		return err
	}

	if source.Type == "s3" {
		s3, err := parseS3Source(source)
		if err != nil {
			return err
		}
		if err := d.Set("s3", s3); err != nil {
			return err
		}
	}

	if source.Type == "s3a" {
		s3a, err := parseS3ASource(source)
		if err != nil {
			return err
		}
		if err := d.Set("s3a", s3a); err != nil {
			return err
		}
	}

	if source.Type == "gcs" {
		gcs, err := parseS3Source(source)
		if err != nil {
			return err
		}
		if err := d.Set("gcs", gcs); err != nil {
			return err
		}
	}

	if source.Type == "local" {
		local, err := parseLocalSource(source)
		if err != nil {
			return err
		}
		if err := d.Set("local", local); err != nil {
			return err
		}
	}

	if source.Type == "hdfs" {
		hdfs, err := parseLocalSource(source)
		if err != nil {
			return err
		}
		if err := d.Set("hdfs", hdfs); err != nil {
			return err
		}
	}

	if source.Type == "jdbc" {
		jdbc, err := parseJDBCSource(source)
		if err != nil {
			return err
		}
		if err := d.Set("jdbc", jdbc); err != nil {
			return err
		}
	}

	if source.Type == "hive" {
		hive, err := parseHiveSource(source)
		if err != nil {
			return err
		}
		if err := d.Set("hive", hive); err != nil {
			return err
		}
	}

	if source.Type == "bigquery" {
		bigQuery, err := parseBigQuerySource(source)
		if err != nil {
			return err
		}
		if err := d.Set("big_query", bigQuery); err != nil {
			return err
		}
	}

	if source.Type == "kafka" {
		kafka, err := parseKafkaSource(source)
		if err != nil {
			return err
		}
		if err := d.Set("kafka", kafka); err != nil {
			return err
		}
	}

	if source.Type == "snowflake" {
		snowflake, err := parseSnowflakeSource(source)
		if err != nil {
			return err
		}
		if err := d.Set("snowflake", snowflake); err != nil {
			return err
		}
	}

	if err := d.Set("labels", source.Labels); err != nil {
		return err
	}
	if err := d.Set("attribute", flattenAttributes(source.Attributes)); err != nil {
		return err
	}
	if err := d.Set("access_rule", flattenAccessRules(source.AccessRules)); err != nil {
		return err
	}
	return err
}

func resourceSourceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*Client)
	source, err := composeSource(d)
	if source == nil || err != nil {
		return err
	}

	e, err := c.CreateSource(*source)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(e.ID))
	return err
}

func resourceSourceUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*Client)
	sourceID := d.Id()
	source, err := composeSource(d)
	if source == nil || err != nil {
		return err
	}

	err = c.UpdateSource(sourceID, *source)
	if err != nil {
		return err
	}

	return nil
}

func resourceSourceDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*Client)
	sourceID := d.Id()

	err := c.DeleteSource(sourceID)
	if err != nil {
		return err
	}

	return nil
}

// Used for both S3 and GCS sources
func parseS3Source(source *Source) ([]map[string]interface{}, error) {
	if source == nil {
		return nil, errors.New("Source is null")
	}

	s3 := make(map[string]interface{})
	s3["bucket"] = source.Bucket
	s3["path"] = source.Path

	fileFormat := parseFileFormat(source.FileFormat)
	for k, v := range fileFormat {
		s3[k] = v
	}

	s3s := make([]map[string]interface{}, 0, 1)
	s3s = append(s3s, s3)
	return s3s, nil
}

func parseS3ASource(source *Source) ([]map[string]interface{}, error) {
	if source == nil {
		return nil, errors.New("Source is null")
	}

	s3a := make(map[string]interface{})
	s3a["bucket"] = source.Bucket
	s3a["path"] = source.Path
	s3a["endpoint"] = source.Endpoint
	s3a["access_key"] = source.AccessKey
	s3a["secret_key"] = source.SecretKey

	fileFormat := parseFileFormat(source.FileFormat)
	for k, v := range fileFormat {
		s3a[k] = v
	}

	s3as := make([]map[string]interface{}, 0, 1)
	s3as = append(s3as, s3a)
	return s3as, nil
}

// Used for local and HDFS sources
func parseLocalSource(source *Source) ([]map[string]interface{}, error) {
	if source == nil {
		return nil, errors.New("Source is null")
	}

	local := make(map[string]interface{})
	local["path"] = source.Path

	fileFormat := parseFileFormat(source.FileFormat)
	for k, v := range fileFormat {
		local[k] = v
	}

	locals := make([]map[string]interface{}, 0, 1)
	locals = append(locals, local)
	return locals, nil
}

func parseJDBCSource(source *Source) ([]map[string]interface{}, error) {
	if source == nil {
		return nil, errors.New("Source is null")
	}

	jdbc := make(map[string]interface{})
	jdbc["url"] = source.URL
	jdbc["schema"] = source.Schema

	credentialsProvider, err := parseLoginCredentialsProviderConfig(source.CredentialsProvider)
	if err != nil {
		return nil, err
	}
	jdbc["credentials_provider"] = []map[string]interface{}{credentialsProvider}

	jdbcs := make([]map[string]interface{}, 0, 1)
	jdbcs = append(jdbcs, jdbc)
	return jdbcs, nil
}

func parseBigQuerySource(source *Source) ([]map[string]interface{}, error) {
	if source == nil {
		return nil, errors.New("Source is null")
	}

	bigQuery := make(map[string]interface{})
	bigQuery["path"] = source.Path

	bigQueries := make([]map[string]interface{}, 0, 1)
	bigQueries = append(bigQueries, bigQuery)
	return bigQueries, nil
}

func parseHiveSource(source *Source) ([]map[string]interface{}, error) {
	if source == nil {
		return nil, errors.New("Source is null")
	}

	hive := make(map[string]interface{})
	hive["database"] = source.Database

	hives := make([]map[string]interface{}, 0, 1)
	hives = append(hives, hive)
	return hives, nil
}

func parseKafkaSource(source *Source) ([]map[string]interface{}, error) {
	if source == nil {
		return nil, errors.New("Source is null")
	}

	kafka := make(map[string]interface{})
	kafka["bootstrap_servers"] = source.BootstrapServers
	kafka["schema_registry_url"] = source.SchemaRegistryURL

	sensitives := make([]map[string]interface{}, len(source.KafkaProperties))
	for i, v := range source.KafkaProperties {
		sa, err := parseSensitiveAttribute(&v)
		if err != nil {
			return nil, err
		}

		sensitives[i] = sa
	}

	kafka["property"] = sensitives

	kafkas := make([]map[string]interface{}, 0, 1)
	kafkas = append(kafkas, kafka)
	return kafkas, nil
}

func parseSnowflakeSource(source *Source) ([]map[string]interface{}, error) {
	if source == nil {
		return nil, errors.New("Source is null")
	}

	snowflake := make(map[string]interface{})
	snowflake["url"] = source.URL
	snowflake["warehouse"] = source.Warehouse
	snowflake["database"] = source.Database
	snowflake["schema"] = source.Schema

	credentialsProvider, err := parseLoginCredentialsProviderConfig(source.CredentialsProvider)
	if err != nil {
		return nil, err
	}
	snowflake["credentials_provider"] = []map[string]interface{}{credentialsProvider}

	snowflakes := make([]map[string]interface{}, 0, 1)
	snowflakes = append(snowflakes, snowflake)
	return snowflakes, nil
}

func composeSource(d *schema.ResourceData) (*Source, error) {
	accessRules, err := expandAccessRules(d.Get("access_rule").([]interface{}))
	if err != nil {
		return nil, err
	}

	if s3, _ := expandSingleMap(d.Get("s3")); s3 != nil {
		fileFormat := composeFileFormat(s3)
		source := Source{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Type:        "s3",
			Bucket:      s3["bucket"].(string),
			Path:        s3["path"].(string),
			FileFormat:  fileFormat,
			Labels:      expandLabels(d),
			Attributes:  expandAttributes(d),
			AccessRules: accessRules,
		}
		return &source, nil
	}

	if s3a, _ := expandSingleMap(d.Get("s3a")); s3a != nil {
		fileFormat := composeFileFormat(s3a)
		source := Source{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Type:        "s3a",
			Bucket:      s3a["bucket"].(string),
			Path:        s3a["path"].(string),
			Endpoint:    s3a["endpoint"].(string),
			AccessKey:   s3a["access_key"].(string),
			SecretKey:   s3a["secret_key"].(string),
			FileFormat:  fileFormat,
			Labels:      expandLabels(d),
			Attributes:  expandAttributes(d),
			AccessRules: accessRules,
		}
		return &source, nil
	}

	if jdbc, _ := expandSingleMap(d.Get("jdbc")); jdbc != nil {
		credentialsProviderMap, err := expandSingleMap(jdbc["credentials_provider"])
		if err != nil {
			return nil, err
		}

		credentialsProvider, err := composeLoginCredentialsProviderConfig(credentialsProviderMap)
		if err != nil {
			return nil, err
		}

		source := Source{
			Name:                d.Get("name").(string),
			Description:         d.Get("description").(string),
			Type:                "jdbc",
			URL:                 jdbc["url"].(string),
			Schema:              jdbc["schema"].(string),
			CredentialsProvider: credentialsProvider,
			Labels:              expandLabels(d),
			Attributes:          expandAttributes(d),
			AccessRules:         accessRules,
		}
		return &source, nil
	}

	if hive, _ := expandSingleMap(d.Get("hive")); hive != nil {
		source := Source{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Type:        "hive",
			Database:    hive["database"].(string),
			Labels:      expandLabels(d),
			Attributes:  expandAttributes(d),
			AccessRules: accessRules,
		}
		return &source, nil
	}

	if bigQuery, _ := expandSingleMap(d.Get("big_query")); bigQuery != nil {
		source := Source{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Type:        "bigquery",
			Path:        bigQuery["path"].(string),
			Labels:      expandLabels(d),
			Attributes:  expandAttributes(d),
			AccessRules: accessRules,
		}
		return &source, nil
	}

	if gcs, _ := expandSingleMap(d.Get("gcs")); gcs != nil {
		fileFormat := composeFileFormat(gcs)
		source := Source{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Type:        "gcs",
			Bucket:      gcs["bucket"].(string),
			Path:        gcs["path"].(string),
			FileFormat:  fileFormat,
			Labels:      expandLabels(d),
			Attributes:  expandAttributes(d),
			AccessRules: accessRules,
		}
		return &source, nil
	}

	if local, _ := expandSingleMap(d.Get("local")); local != nil {
		fileFormat := composeFileFormat(local)
		source := Source{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Type:        "local",
			Path:        local["path"].(string),
			FileFormat:  fileFormat,
			Labels:      expandLabels(d),
			Attributes:  expandAttributes(d),
			AccessRules: accessRules,
		}
		return &source, nil
	}

	if hdfs, _ := expandSingleMap(d.Get("hdfs")); hdfs != nil {
		fileFormat := composeFileFormat(hdfs)
		source := Source{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Type:        "hdfs",
			Path:        hdfs["path"].(string),
			FileFormat:  fileFormat,
			Labels:      expandLabels(d),
			Attributes:  expandAttributes(d),
			AccessRules: accessRules,
		}
		return &source, nil
	}

	if kafka, _ := expandSingleMap(d.Get("kafka")); kafka != nil {
		value := kafka["property"]

		array, ok := kafka["property"].([]interface{})
		if !ok {
			return nil, fmt.Errorf("Kafka Properties Value is not an array. Value: %v", value)
		}

		sensitives := make([]SensitiveAttribute, len(array))
		for i, v := range array {

			prop, ok := v.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("Kafka Properties Value is not a map interfaces. Value: %v.", v)
			}
			sa, err := composeSensitiveAttribute(prop)
			if err != nil {
				return nil, err
			}
			sensitives[i] = *sa
		}

		source := Source{
			Name:              d.Get("name").(string),
			Description:       d.Get("description").(string),
			Type:              "kafka",
			BootstrapServers:  kafka["bootstrap_servers"].(string),
			SchemaRegistryURL: kafka["schema_registry_url"].(string),
			KafkaProperties:   sensitives,
			Labels:            expandLabels(d),
			Attributes:        expandAttributes(d),
			AccessRules:       accessRules,
		}
		return &source, nil
	}

	if snowflake, _ := expandSingleMap(d.Get("snowflake")); snowflake != nil {
		credentialsProviderMap, err := expandSingleMap(snowflake["credentials_provider"])
		if err != nil {
			return nil, err
		}

		credentialsProvider, err := composeLoginCredentialsProviderConfig(credentialsProviderMap)
		if err != nil {
			return nil, err
		}

		source := Source{
			Name:                d.Get("name").(string),
			Description:         d.Get("description").(string),
			Type:                "snowflake",
			URL:                 snowflake["url"].(string),
			Schema:              snowflake["schema"].(string),
			Warehouse:           snowflake["warehouse"].(string),
			Database:            snowflake["database"].(string),
			CredentialsProvider: credentialsProvider,
			Labels:              expandLabels(d),
			Attributes:          expandAttributes(d),
			AccessRules:         accessRules,
		}
		return &source, nil
	}

	return nil, errors.New("Invalid source type")
}

func parseFileFormat(fileFormat *FileFormat) map[string]interface{} {
	fileFormatMap := make(map[string]interface{})
	fileFormatMap["file_format"] = fileFormat.Type
	if fileFormat.Type == "csv" {
		if fileFormat.Compression != nil {
			fileFormatMap["compression"] = fileFormat.Compression
		} else {
			fileFormatMap["compression"] = nil
		}
		if fileFormat.DateFormat != nil {
			fileFormatMap["date_format"] = fileFormat.DateFormat
		} else {
			fileFormatMap["date_format"] = nil
		}
		if fileFormat.EmptyValue != nil {
			fileFormatMap["empty_value"] = fileFormat.EmptyValue
		} else {
			fileFormatMap["empty_value"] = nil
		}
		if fileFormat.Sep != nil {
			fileFormatMap["field_separator"] = fileFormat.Sep
		} else {
			fileFormatMap["field_separator"] = nil
		}
		if fileFormat.IgnoreLeadingWhiteSpace != nil {
			fileFormatMap["ignore_leading_whitespace"] = fileFormat.IgnoreLeadingWhiteSpace
		} else {
			fileFormatMap["ignore_leading_whitespace"] = nil
		}
		if fileFormat.IgnoreTrailingWhiteSpace != nil {
			fileFormatMap["ignore_trailing_whitespace"] = fileFormat.IgnoreTrailingWhiteSpace
		} else {
			fileFormatMap["ignore_trailing_whitespace"] = nil
		}
		if fileFormat.IncludeHeader != nil {
			fileFormatMap["include_header"] = fileFormat.IncludeHeader
		} else {
			fileFormatMap["include_header"] = nil
		}
		if fileFormat.QuoteAll != nil {
			fileFormatMap["quote_all"] = fileFormat.QuoteAll
		} else {
			fileFormatMap["quote_all"] = nil
		}
		if fileFormat.TimestampFormat != nil {
			fileFormatMap["timestamp_format"] = fileFormat.TimestampFormat
		} else {
			fileFormatMap["timestamp_format"] = nil
		}
		if fileFormat.LineSep != nil {
			fileFormatMap["line_separator"] = fileFormat.LineSep
		} else {
			fileFormatMap["line_separator"] = nil
		}
	}
	return fileFormatMap
}

func composeFileFormat(d map[string]interface{}) *FileFormat {
	fileFormat := FileFormat{
		Type: d["file_format"].(string),
	}

	if d["file_format"] == "csv" {
		if compression, ok := d["compression"].(string); ok {
			fileFormat.Compression = &compression
		}
		if dateFormat, ok := d["date_format"].(string); ok {
			fileFormat.DateFormat = &dateFormat
		}
		if emptyValue, ok := d["empty_value"].(string); ok {
			fileFormat.EmptyValue = &emptyValue
		}
		if ignoreLeadingWhiteSpace, ok := d["ignore_leading_whitespace"].(bool); ok {
			fileFormat.IgnoreLeadingWhiteSpace = &ignoreLeadingWhiteSpace
		}
		if ignoreTrailingWhiteSpace, ok := d["ignore_trailing_whitespace"].(bool); ok {
			fileFormat.IgnoreTrailingWhiteSpace = &ignoreTrailingWhiteSpace
		}
		if includeHeader, ok := d["include_header"].(bool); ok {
			fileFormat.IncludeHeader = &includeHeader
		}
		if quoteAll, ok := d["quote_all"].(bool); ok {
			fileFormat.QuoteAll = &quoteAll
		}
		if sep, ok := d["field_separator"].(string); ok {
			fileFormat.Sep = &sep
		}
		if timestampFormat, ok := d["timestamp_format"].(string); ok {
			fileFormat.TimestampFormat = &timestampFormat
		}
		if lineSep, ok := d["line_separator"].(string); ok {
			fileFormat.LineSep = &lineSep
		}
	}

	return &fileFormat
}

func expandAccessRules(accessRules []interface{}) ([]AccessRule, error) {
	res := make([]AccessRule, 0, len(accessRules))

	for _, accessRule := range accessRules {
		val, _ := accessRule.(map[string]interface{})

		principals, err := expandPrincipalIds(val["principals"].([]interface{}))
		if err != nil {
			return nil, err
		}

		maskingRules, err := expandMaskingRules(val["masking_rule"].([]interface{}))
		if err != nil {
			return nil, err
		}

		parsed := AccessRule{
			Resource:     val["resource"].(string),
			Principals:   principals,
			MaskingRules: maskingRules,
		}
		res = append(res, parsed)
	}

	return res, nil
}

func expandMaskingRules(maskingRules []interface{}) ([]MaskingRule, error) {
	res := make([]MaskingRule, 0, len(maskingRules))

	for _, maskingRule := range maskingRules {
		val, _ := maskingRule.(map[string]interface{})

		if filterMaskingRule, _ := expandSingleMap(val["filter"]); filterMaskingRule != nil {
			parsed, err := composeFilterMaskingRule(filterMaskingRule)
			if err != nil {
				return nil, err
			}
			res = append(res, *parsed)
		}
		if maskMaskingRule, _ := expandSingleMap(val["mask"]); maskMaskingRule != nil {
			parsed, err := composeMaskMaskingRule(maskMaskingRule)
			if err != nil {
				return nil, err
			}
			res = append(res, *parsed)
		}
	}

	return res, nil
}

func composeFilterMaskingRule(d map[string]interface{}) (*MaskingRule, error) {
	return &MaskingRule{
		Type:       "filter",
		Expression: d["expression"].(string),
	}, nil
}
func composeMaskMaskingRule(d map[string]interface{}) (*MaskingRule, error) {
	return &MaskingRule{
		Type:       "mask",
		Column:     d["column"].(string),
		Expression: d["expression"].(string),
	}, nil
}

func validateFileFormat() schema.SchemaValidateFunc {
	return validation.StringInSlice([]string{"csv", "orc", "parquet"}, false)
}

func flattenAccessRules(accessRules []AccessRule) []map[string]interface{} {
	res := make([]map[string]interface{}, 0, len(accessRules))
	for _, accessRule := range accessRules {
		single := make(map[string]interface{})
		single["resource"] = accessRule.Resource
		single["principals"] = flattenPrincipalIds(accessRule.Principals)
		single["masking_rule"] = flatternMaskingRules(accessRule.MaskingRules)
		res = append(res, single)
	}
	return res
}

func flatternMaskingRules(maskingRules []MaskingRule) []map[string]([]map[string]interface{}) {
	res := make([]map[string]([]map[string]interface{}), 0, len(maskingRules))
	for _, maskingRule := range maskingRules {
		single := make(map[string]([]map[string]interface{}))
		if maskingRule.Type == "filter" {
			nest := make(map[string]interface{})
			nest["expression"] = maskingRule.Expression
			single["filter"] = []map[string]interface{}{nest}
		}
		if maskingRule.Type == "mask" {
			nest := make(map[string]interface{})
			nest["column"] = maskingRule.Column
			nest["expression"] = maskingRule.Expression
			single["mask"] = []map[string]interface{}{nest}
		}
		res = append(res, single)
	}
	return res
}
