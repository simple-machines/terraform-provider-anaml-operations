---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "anaml-operations_attribute_restriction Resource - terraform-provider-anaml-operations"
subcategory: ""
description: |-
  Attribute Restrictions
  An Attribute is a key/value pair for user-defined metadata. Restrictions limit the attributes
  that can be applied to a given object, and what values they can take.
  Multiple different types of attributes are supported:
  Enum ("Choice"), Free Text, Boolean, Integer
---

# anaml-operations_attribute_restriction (Resource)

# Attribute Restrictions

An Attribute is a key/value pair for user-defined metadata. Restrictions limit the attributes
that can be applied to a given object, and what values they can take.

Multiple different types of attributes are supported:

- Enum ("Choice")
- Free Text
- Boolean
- Integer



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `applies_to` (Set of String)
- `description` (String)
- `key` (String)

### Optional

- `boolean` (Block List, Max: 1) (see [below for nested schema](#nestedblock--boolean))
- `enum` (Block List, Max: 1) (see [below for nested schema](#nestedblock--enum))
- `freetext` (Block List, Max: 1) (see [below for nested schema](#nestedblock--freetext))
- `integer` (Block List, Max: 1) (see [below for nested schema](#nestedblock--integer))

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--boolean"></a>
### Nested Schema for `boolean`


<a id="nestedblock--enum"></a>
### Nested Schema for `enum`

Required:

- `choice` (Block Set, Min: 1) (see [below for nested schema](#nestedblock--enum--choice))

<a id="nestedblock--enum--choice"></a>
### Nested Schema for `enum.choice`

Required:

- `value` (String)

Optional:

- `display_colour` (String)
- `display_emoji` (String)



<a id="nestedblock--freetext"></a>
### Nested Schema for `freetext`


<a id="nestedblock--integer"></a>
### Nested Schema for `integer`


