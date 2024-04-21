# Entities definitions

This section describes the current types of public Moonlogs entities available for operations via the token API.

## List of entities

  * [Schema](#schema)
  * [Record](#record)

## Schema

A `schema` is a container for aggregating records of a specific business scenario or functional block.

| Name                                     | Required | Type                                   |
|------------------------------------------|----------|----------------------------------------|
|          [title](#schema-title)          | Yes      | `string`                               |
|    [description](#schema-description)    | Yes      | `string`                               |
|           [name](#schema-name)           | Yes      | `string`                               |
| [fields](#schema-fields)                 | Yes      | Array of [SchemaField](#schema-field) |
| [kinds](#schema-kinds)                   | –        | Array of [SchemaKind](#schema-kind)   |
| [tag_id](#schema-tag_id)                 | –        | `integer`                              |
| [tag_name](#schema-tag_name)             | –        | `string`                               |
| [retention_days](#schema-retention_days) | –        | `integer`                              |

### title {#schema-title}

* Type: `string`
* Required: `yes`

`title` is a the human-readable name of the schema in the web interface. Schema search will also search for schemas based on this characteristic.

### description {#schema-description}

* Type: `string`
* Required: `yes`

`description` is a human-readable description of schema details in the web interface. Schema search will also search for schemas based on this characteristic.

### name {#schema-name}

* Type: `string`
* Required: `yes`

`name` is a textual identifier for the schema. Must be specified in Latin, in lowercase, and with underscores as separators.

### fields {#schema-fields}

* Type: Array of [SchemaField](#schema-field)
* Required: `yes`

`fields` is a set of fields by which log grouping inside a schema will occur.

### kinds {#schema-kinds}

* Type: Array of [SchemaKind](#schema-kind)
* Required: —

`kinds` is a set of select options by which log grouping will occur.

### tag_id {#schema-tag_id}

* Type: `integer`
* Required: —

`tag_id` is and ID of Tag, applied to schema. Tags are used to group multiple schemas belonging to the same business area into logical blocks. If a schema includes a `tag_id`, and a user possesses a non-empty list of assigned `tag_ids`, access will be granted to schemas that contain any of the corresponding tags.

### tag_name {#schema-tag_name}

* Type: `string`
* Required: —

`tag_name` is a a human-readable name of the tag. Schema search will also search for schema based on this characteristic.

### retention_days {#schema-retention_days}

* Type: `integer`
* Required: —

`retention_days` represents the duration logs will remain available after creation. Once this duration expires, the logs will be automatically deleted. To allow logs to persist indefinitely, either specify 0 or leave this field empty.

## SchemaField {#schema-field}

| Name                         | Required | Type     |
|------------------------------|----------|----------|
| [title](#schema-field-title) | Yes      | `string` |
|  [name](#schema-field-name)  | Yes      | `string` |

### title {#schema-field-title}

* Type: `string`
* Required: `yes`

`title` is a human-readable name of the field in the web interface for log filtering.

### name {#schema-field-name}

* Type: `string`
* Required: `yes`

`name` is a textual identifier of the field. Must be specified in Latin, in lowercase, and with underscores as separators.

## SchemaKind {#schema-kind}

| Name                        | Required | Type     |
|-----------------------------|----------|----------|
| [title](#schema-kind-title) | Yes      | `string` |
|  [name](#schema-kind-name)  | Yes      | `string` |

### title {#schema-kind-title}

* Type: `string`
* Required: `yes`

`title` is a human-readable name of the kind in the web interface for log filtering.

### name {#schema-kind-name}

* Type: `string`
* Required: `yes`

`name` is a textual identifier for the kind. Must be specified in Latin, in lowercase, and with underscores as separators.
