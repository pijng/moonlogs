
# Moonlogs

Moonlogs is a business event logging tool with a built-in user-friendly web interface for easy access to events.

For a full list of key capabilities, please see the [Features](#features) section.


## Installation

### Prerequisites

- Docker: Make sure you have Docker installed on your system.

### Getting Started

1. **Create a Docker Compose File:**

Create a file named `docker-compose.yml` in your desired directory and copy the following content into it:

```yaml
version: '3'

services:
  moonlogs:
    image: pijng/moonlogs
    restart: always
    ports:
      - "4200:4200"
    volumes:
      - moonlogs-data:/opt/moonlogs
    command: --port=4200
volumes:
  moonlogs-data:
```

Save the file.

2. **Build and Run the Docker Containers:**

Open a terminal in your project's root directory and run the following command:

```bash
docker-compose up -d
```

This will download the Moonlogs image and start the Moonlogs container in detached mode.

3. **Access Moonlogs:**

Navigate to [http://localhost:4200](http://localhost:<your-port>). You should see the Moonlogs UI. Follow the instructions for creating the initial administrator there.


### Configuration Options

- The Moonlogs container is configured to run on port 4200 by default. If you need to change the port, update the `--port` parameter in the `command` section of the `docker-compose.yml` file.

- By default, Moonlogs uses a Docker volume named `moonlogs-data` to store configuration and database files. This ensures persistent data even if the container is stopped or removed.
If you prefer to use your own bind mount for data storage, you can modify the `volumes` section in the `docker-compose.yml` file.

    * Remove the line at the bottom of `docker-compose.yml`:

        ```yaml
        volumes:
          moonlogs-data:
        ```

    * Change the `services.moonlogs.volumes` directive to:

        ```yaml
        volumes:
          - <your-desired-dir-on-host>:/opt/moonlogs
        ```

    * The resulting file may look like the following:

        ```yaml
        version: '3'

        services:
          moonlogs:
            image: pijng/moonlogs
            restart: always
            ports:
              - "4200:4200"
            volumes:
              - /etc/moonlogs:/opt/moonlogs
            command:
              --port=4200
        ```

## Features

#### Meta-groups for log organization

Create separate meta-groups (schemas) to categorize logs by domain areas. For instance, create schemas for the checkout process, user access setting changes, and Uber Eats integration. Logs within each schema are recorded independently, facilitating efficient event retrieval.

#### Query-based log subgrouping

Group logs within a schema based on specified queries to enhance information integrity. Logs for distinct clients, such as those with IDs 4 and 5, will be segregated within the overall schema. This not only simplifies searchability but also ensures unrelated events remain separate even if in the same schema.

#### Custom filters for schemas

Define unique static filters for specific schemas, streamlining event retrieval. For example, a schema for Uber Eats may have filters like `organization_id`, `order_id`, `external_order_id`, and `restaurant_id`. This eliminates the need for manual DSL queries, particularly beneficial for non-technical users.

#### Convenient schema-based filters

Generate convenient filters on the web interface for each schema, simplifying event search by allowing users to input values. This eliminates the complexity of composing queries with an undefined set of parameters, making it user-friendly, especially for non-technical personnel.

#### Flexible log retention time

Specify varying retention times for each schema to align with specific business needs. For instance, set a 7-day retention time for logs in the "Glovo integration" schema, while logs in the "User's rights change history" schema can be stored indefinitely. Adjust these settings dynamically as business requirements evolve.
