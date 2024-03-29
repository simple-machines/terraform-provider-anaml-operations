package anaml

import (
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const entityDescription = `# Entities

An Entity is an item in the business domain. Common examples of Entities are:

- Customers
- Accounts
- Products
- Orders

Anything that has a unique identifier and would be useful to report or predict on could be an Entity.

In a relational database, the identifiers for Entities will often be used for primary keys.

Tables need to specify one or more columns with entity identifiers in order to be used for Feature definitions.

Features will be generated for a specific Entity. This means the aggregation will be grouped by each Entity identitifer.
`

func ResourceEntity() *schema.Resource {
	return &schema.Resource{
		Description: entityDescription,
		Create:      resourceEntityCreate,
		Read:        resourceEntityRead,
		Update:      resourceEntityUpdate,
		Delete:      resourceEntityDelete,
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
			"default_column": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"default_column", "entities"},
			},
			"required_type": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"entities"},
				ValidateFunc: validation.StringInSlice([]string{
					"string", "integer", "long", "binary",
				}, false),
				Description: "The data type the entity is encoded as. If set, tables' entity columns must be of this type. One of string, integer, long, or binary",
			},
			"entities": {
				Type:        schema.TypeList,
				Description: "Entities from which this composite entity is derived",
				Optional:    true,

				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validateAnamlIdentifier(),
				},
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
		},
	}
}

func resourceEntityRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*Client)
	entityID := d.Id()

	entity, err := c.GetEntity(entityID)
	if err != nil {
		return err
	}
	if entity == nil {
		d.SetId("")
		return nil
	}

	if err := d.Set("name", entity.Name); err != nil {
		return err
	}
	if err := d.Set("description", entity.Description); err != nil {
		return err
	}
	if entity.DefaultColumn != nil {
		if err := d.Set("default_column", entity.DefaultColumn); err != nil {
			return err
		}

		if entity.RequiredType != nil {
			derefed := *(entity.RequiredType)
			requiredTypeString, ok := derefed.(string)

			if ok {
				if err := d.Set("required_type", requiredTypeString); err != nil {
					return err
				}
			} else {
				if err := d.Set("required_type", "Complex Type"); err != nil {
					return err
				}
			}
		} else {
			if err := d.Set("required_type", nil); err != nil {
				return err
			}
		}
		if err := d.Set("entities", nil); err != nil {
			return err
		}
	}
	if entity.Entities != nil {
		if err := d.Set("default_column", nil); err != nil {
			return err
		}
		if err := d.Set("required_type", nil); err != nil {
			return err
		}
		if err := d.Set("entities", identifierList(*entity.Entities)); err != nil {
			return err
		}
	}
	if err := d.Set("labels", entity.Labels); err != nil {
		return err
	}
	if err := d.Set("attribute", flattenAttributes(entity.Attributes)); err != nil {
		return err
	}
	return err
}

func buildEntity(d *schema.ResourceData) Entity {
	entity := Entity{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Labels:      expandLabels(d),
		Attributes:  expandAttributes(d),
	}

	if default_column := d.Get("default_column").(string); default_column != "" {
		entity.Type = "base"
		entity.DefaultColumn = &default_column
		if required_type, set := d.GetOk("required_type"); set {
			required_type := required_type
			entity.RequiredType = &required_type
		}
	} else {
		entities := expandIdentifierList(d.Get("entities").([]interface{}))
		entity.Type = "composite"
		entity.Entities = &entities
	}

	return entity
}

func resourceEntityCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*Client)
	entity := buildEntity(d)
	e, err := c.CreateEntity(entity)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(e.ID))
	return err
}

func resourceEntityUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*Client)
	entityID := d.Id()
	entity := buildEntity(d)
	err := c.UpdateEntity(entityID, entity)
	if err != nil {
		return err
	}

	return nil
}

func resourceEntityDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*Client)
	entityID := d.Id()

	err := c.DeleteEntity(entityID)
	if err != nil {
		return err
	}

	return nil
}
