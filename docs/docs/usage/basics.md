# Basic usage

In this section you will learn how to start sending events to Moonlogs.

## Prerequisites

It is highly recommended that you use an off-the-shelf SDK to interact with Moonlogs.
Moonlogs currently provides the following SDKs:
  * [Ruby on Rails SDK](/usage/ruby-on-rails)
  * [Typescript SDK](/usage/typescript)


:::tip
If your language is not represented in the list of ready SDKs - you can always fill in a new issue and we will add it!
:::

Otherwise, you can use the specification described by Swagger if there is no SDK for your language yet or you want to work with the API directly:

📁 [Link to Swagger spec](https://raw.githubusercontent.com/pijng/moonlogs/master/internal/api/swagger.yaml)

## Basic steps

At its core, in order to start submitting events to Moonlogs you need:
1. Create an API token to access the API
2. Install the SDK for your language or work with the API directly
4. Create domain schemas
5. Send events

## 1. Create an API token

In order to create a new API token, you need to go to Moonlogs' Web interface and:

1. Click the `API token` tab on the left:

<p align="left">
  <img src="/usage/api_token_navigation.png" alt="API tokens tab" width="40%" style=""/>
</p>
<hr>

2. Then click the plus icon right to the `API tokens` header:

<p align="left">
  <img src="/usage/api_token_plus.png" alt="API tokens plus icon" width="40%" style=""/>
</p>
<hr>

3. Fill out the API token name and press `Create` button:

<p align="left">
  <img src="/usage/api_token_form.png" alt="New API token form" style=""/>
</p>
<hr>

4. After generating the API token, make sure to copy it immediately, as it won't be shown again:

<p align="left">
  <img src="/usage/api_token_created.png" alt="New API token created" style=""/>
</p>

## 2. Install SDK or work with the API directly

It is highly recommended that you use ready SDK to interact with Moonlogs:
  * [Ruby on Rails SDK](/usage/ruby-on-rails)
  * [Typescript SDK](/usage/typescript)

You can follow the link of your language for a description of the next steps, or continue reading on for a general description of working with a raw API.

## 3. Create domain schemas

Now you need to create a `schema` to which you will send events.
A `schema` is a kind of container for aggregating events of a specific business scenario or functional block.

Examples of schemas could include: Online payments, User change history, SMS authentication, Export reports by email, etc.

Each schema must consist of the following attributes:

* `name`: textual identifier for the schema used in most operations. It must be specified in Latin, in lowercase, and with underscores as separators.
* `title`: human-readable name of the schema in the web interface.
* `description`: human-readable description of the schema details in the web interface.
* `retention_days`: the number of days during which logs will be available after their creation. After the specified number of days elapses, the logs will be deleted. To set an infinite lifespan, specify 0 or leave the field empty.

* `fields`: an array of fields by which log grouping will occur.
Each field consists of:
  * `title`: human-readable name of the field in the web interface for log filtering.
  * `name`: textual identifier for the field. It must be specified in Latin, in lowercase, and with underscores as separators.

To create the first schema, you need to send the following request:

```bash
curl --location --request POST '<host:port>/api/schemas' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <api_token>' \
--data '{
    "title": "Online Payments",
    "name": "online_payments",
    "description": "History of online payments",
    "fields": [
        {
            "title": "Organization ID",
            "name": "organization_id"
        },
        {
            "title": "Customer ID",
            "name": "customer ID"
        },
        {
            "title": "Payment ID",
            "name": "payment_id"
        }
    ]
}'
```

:::info
Please replace `<host:port>` with host:port Moonlogs is running on and `<api_token>` with the API token you generated in step one.
:::

In response, you should receive something like this:

```json
{
    "code": 200,
    "success": true,
    "error": "",
    "data": {
        "id": 1,
        "title": "Online Payments",
        "description": "History of online payments",
        "name": "online_payments",
        "fields": [
            {
                "title": "Organization ID",
                "name": "organization_id"
            },
            {
                "title": "Customer ID",
                "name": "customer_id"
            },
            {
                "title": "Payment ID",
                "name": "payment_id"
            }
        ],
        "kinds": null
    },
    "meta": {
        "page": 0,
        "count": 0,
        "pages": 0
    }
}
```

This means that the schema was successfully created and assigned an ID.

:::tip
If you send a duplicate request to create a schema with an existing `name`, then instead of creating a new schema, the existing one will be updated. This simplifies the creation/update of schemas when starting your application.
:::

## 4. Send events

After creating the `schema`, you can send your first `event` to this schema. To do this, you need to prepare the `event` payload, which must consist of the following attributes:

* `schema_name`: the textual identifier of the existing schema (`schema.name`)
* `text`: text of the event that will be shown is web ui
* `query`: a set of parameters from the schema fields (`schema.fields[]`). This set determines the grouping of events.

For example, the basic request to create an event with simple payload would look like the following:

```bash
curl --location --request POST '<host:port>/api/logs' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <api_token>' \
--data '{
    "schema_name": "online_payment",
    "text": "The customer paid $5 for a subscription to the «Basic» tariff",
    "query": {
        "organization_id": 34,
        "customer_id": 891,
        "payment_id": 217596
    }
}'
```

In response, you should receive something like this:

```json
{
    "code": 200,
    "success": true,
    "error": "",
    "data": {
        "id": 1,
        "text": "The customer paid $5 for a subscription to the «Basic» tariff",
        "created_at": "2024-03-31T19:16:45+03:00",
        "schema_name": "online_payments",
        "schema_id": 23,
        "query": {
            "customer_id": "891",
            "organization_id": "34",
            "payment_id": "217596"
        },
        "group_hash": "7792753873920415191",
        "level": "Info"
    },
    "meta": {
        "page": 0,
        "count": 0,
        "pages": 0
    }
}
```

This means that the event was successfully created and assigned an ID.

### Verify the presence of the created event in the web interface

Now you can verify that the event was succesfully created by checking it in Web UI:

1. Click the `Log groups` tab on the left and click on `Online Payments` card:

<p align="left">
  <img src="/usage/existing_schema.png" alt="Existing schema" style=""/>
</p>
<hr>

2. Here you can see the event was successfully created with all the data we specified: text and query fields:

<p align="left">
  <img src="/usage/new_event.png" alt="Newly created event" style=""/>
</p>
<hr>

### Test the grouping

Now try to send another event with different set of query fields, for example:

```bash
curl --location --request POST '<host:port>/api/logs' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <api_token>' \
--data '{
    "schema_name": "online_payment",
    "text": "The customer paid $25 for a subscription to the «Essential» tariff",
    "query": {
        "organization_id": 35,
        "customer_id": 540,
        "payment_id": 217597
    }
}'
```

After reloading the page, you will see that the second event was successfully created. But more importantly, it was assigned to a different event group based on its query attributes:

<p align="left">
  <img src="/usage/event_grouping.png" alt="Test event grouping" style=""/>
</p>
<hr>

As a result, you have learned how to create events in Moonlogs in the most basic form.

## What's next?

* For a more detailed description to schema and event formats, please refer to our [Entities definitions](/definitions/introduction)

* Check out [Introduction to the Web UI](/web-ui/introduction) section to familiarize yourself with the Moonlogs built-in web-interface
