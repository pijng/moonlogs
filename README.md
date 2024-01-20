
# Moonlogs

Moonlogs is a business event logging tool with a built-in user-friendly web interface for easy access to events.


## Features

- **Ability to create separate meta-groups to divide logs by domain areas (schemas)**

    For example, you can create a separate schema for the checkout process, a schema for the change history of user access settings, and a separate schema for the Uber Eats integration module.

    Later, each log can be recorded in a separate independent schema, which makes it easy to find the necessary events.

- **Additional grouping of logs by a given query within a schema to maintain information integrity**

    For example, when logs are written for a client with ID 4 and a client with ID 5, they will be recorded in different subgroups in the overall schema.

    This not only allows logs with different set of queries to be separated into separate groups – which increases the ease of searching for specific events – but also increases the integrity of the information because unrelated events will not be mixed together even if they are in the same schema.

- **Convenient filters for specific schemes**

    You can define a unique static set of filters for specific schemes. For example, a scheme for Uber Eats will have the filter with fields like: `organization_id`, `order_id`, `external_order_id` and `restaurant_id`. And a scheme for loyalty programs will have a filter with fields: `organization_id`, `kind`, `program_id`, `bonus_provider`.

    Based on each such set of parameters, a convenient filter will be generated on the web interface side, allowing you to find the desired events by simply filling in the values. This eliminates the need to manually compose a query with a unique DSL with an undefined set of parameters, which can be confusing and difficult, especially for non-technical personnel.

- **Flexible log retention time**

    You can specify different retention times for each schema, depending on your needs.

    For example, for a «Glovo integration» module schema, you can specify a retention time of 7 days. As a result, each individual log in this group will be deleted 7 days after its creation.
    And for the «User's rights change history» schema you can leave the retention time empty - thus, the logs from this schema will be stored indefinitely.

    Of course, this setting can be changed at any time as business requirements change.

