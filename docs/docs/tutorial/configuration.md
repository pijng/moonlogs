# Configuration

Moonlogs uses configuration options to define the underlying behavior. These options can be changed to adjust the service to your requirements.

## List of configuration options

| Name                                            | Flag                    | Default value                     |
|-------------------------------------------------|-------------------------|-----------------------------------|
|                  [port](#port)                  | --port                  | 4200                              |
|            [db_adapter](#db_adapter)            | --db-adapter            | sqlite                            |
|               [db_path](#db_path)               | --db-path               | /var/lib/moonlogs/database.sqlite |
| [read_timeout](#read_timeout)                   | --read-timeout          | 5s                                |
| [write_timeout](#write_timeout)                 | --write-timeout         | 1s                                |
| [async_record_creation](#async_record_creation) | --async-record-creation | false                             |

::: warning
Configuration options will always use the default value if left unspecified.
:::

## Usage

You can specify configuration options either in the `config.yaml` file or by passing a command-line flag when executing Moonlogs.

### Using `config.yaml`

By default, the `config.yaml` file is located at the `/etc/moonlogs/config.yaml` path.

Typical `config.yaml` file may look like this:

```yaml
port: 4200
db_adapter: mongodb
db_path: 'mongodb://127.0.0.1:27017'
async_record_creation: false
read_timeout: 1s
write_timeout: 15s
```

After making appropriate changes to the config file, don't forget to restart the service:

```bash
service moonlogs restart
```

### Using command-line flags

If you're running Moonlogs via Docker Compose, it's preferable to use command-line flags to specify configuration options. This allows you to customize the behavior of Moonlogs directly from the `docker-compose.yml` file.

You can specify command-line flags in the `services.moonlogs.command` section of your `docker-compose.yml` file. This ensures that Moonlogs is started with the desired configuration options each time it's launched through Docker Compose.

For example, to specify a custom port for Moonlogs, you can add the `--port` flag followed by the desired port number to the `command` section in your `docker-compose.yml`:

```yaml
services:
  moonlogs:
    command: "--port=7171"
```

Specifying the MongoDB adapter with a connection URL can be done as follows:

```yaml
services:
  moonlogs:
    command: "--db-adapter=mongodb --db-path=mongodb://127.0.0.1:27017"
```


## Explanation of options

### port

* Type: `int`
* Default: `4200`

The `port` parameter specifies the port number on which the Moonlogs service will listen for incoming connections. By default, Moonlogs uses port 4200.

### db_adapter

* Type: `sqlite | mongodb`
* Default: `sqlite`

The `db_adapter` parameter determines the type of database adapter or driver used by Moonlogs to interact with the underlying database. It supports two options: `sqlite` for SQLite databases and `mongodb` for MongoDB databases.

### db_path

* Type: `string`
* Default: `/var/lib/moonlogs/database.sqlite`

The `db_path` parameter specifies the file path or connection URL for the database used by Moonlogs.

* For SQLite the path should be an absolute path to database file, for example:
  * `/var/lib/moonlogs/database.sqlite`
  * `/etc/moonlogs/moonlogs.db`

* For MongoDB, however, the path should be a connection URL to the MongoDB instance, for example:
  * `mongodb://127.0.0.1:27017`
  * `mongodb://192.168.101.2:27017`

### read_timeout

* Type: `duration`
* Default: `1s`

The `read_timeout` parameter sets the maximum duration that Moonlogs waits for reading data from incoming request.

A `duration` type is a string representation of possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as "300ms", "3s" or "1h15m".

Valid time units are:
* `ns`
* `us` (or `µs`)
* `ms`
* `s`
* `m`
* `h`

### write_timeout

* Type: `duration`
* Default: `10s`

The `write_timeout` parameter sets the maximum duration that Moonlogs waits for writing data to the client.

A `duration` type is a string representation of possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as "300ms", "3s" or "1h15m".

Valid time units are:
* `ns`
* `us` (or `µs`)
* `ms`
* `s`
* `m`
* `h`

::: warning
Keep in mind that `write_timeout` also serves as a deadline for returning a response to the Moonlogs Web UI. Therefore, setting this value too low might result in timeouts when querying logs from the web interface. It's essential to balance the `write_timeout` value to ensure timely responses without sacrificing usability.

Additionally, consider controlling the time you want to wait for a response on your client or SDK side for optimal performance.
:::

### async_record_creation

* Type: `bool`
* Default: `false`

The `async_record_creation` parameter determines whether Moonlogs asynchronously creates records after receiving request from a client.

When set to `true`, record creation operations are performed asynchronously, potentially improving performance. If left unspecified or set to `false`, record creation occurs synchronously.

This is particularly useful when you don't have the ability to make asynchronous requests or delegate requests for background job processing, and you don't want to block your code execution while sending logs.

::: tip
When `async_record_creation` is set to `true`, it is highly recommended to specify the `created_at` parameter for your logs on your side to maintain a consistent order of logs, since `async_record_creation` won't guarantee the order of logs to be the same as you send it.
:::