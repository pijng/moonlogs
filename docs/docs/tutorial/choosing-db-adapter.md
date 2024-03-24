# Choosing DB Adapter

Moonlogs, as an application, does not store any data itself. Instead, it utilizes an underlying database for persistent data storage and interacts with the database through appropriate adapters.

Unfortunately, there is no "silver bullet" when it comes to choosing a database. Therefore, the choice of database depends entirely on your requirements for various factors. Below, we will discuss the adapters supported by Moonlogs and which one is best suited for your use case.

## Supported adapters

At the moment, Moonlogs offers the following options:

* `SQLite` (default)
* `MongoDB`

## When to choose SQLite

Choosing SQLite as the database for Moonlogs makes sense if:

* You **plan to handle low to moderate loads**, ranging from 0 to tens of thousands of requests per minute.
* You **don't have strict latency requirements** and are satisfied with an average response time in the range of `60-90ms` per write.
* You **don't plan to horizontally scale** the database across one or multiple servers.
* You **expect** to store fast-paced logs with `retention_days` of a few days.
* You **don't expect** burst spikes in writes reaching hundreds of thousands of requests per minute (for example, logging mass mailings/notifications/sms).

## When to choose MongoDB

Choosing MongoDB as the database for Moonlogs makes sense if:

* You **plan to handle high loads**, ranging from hundres of thousands to millions of requests per minute.
* You **have strict latency requirements** and are satisfied with an average response time in the range of `5-15ms` per write.
* You **plan to horizontally scale** the database across one or multiple servers.
* You **expect** a high spike in requests, significantly exceeding the average load (for example, logging mass mailings/notifications/sms).

## Why SQLite as default?

One of the primary reasons for selecting SQLite is its lightweight and embedded nature. Moonlogs utilizes the [pure-Go SQLite adapter provided by glebarez](https://github.com/glebarez/go-sqlite) to accomplish this goal.

This allows for complete embedding of SQLite within the Moonlogs binary file, thereby eliminating the need to install additional dependencies.

This also simplifies working with Moonlogs during local development.

## SQLite is not production-grade solution?

Often, people consider SQLite suitable only for tests, local development, or small embedded applications that need to persistently store small amount of data.

However, in reality, SQLite performs well under moderate production workloads. Moonlogs employs several approaches to achieve high performance with SQLite:

**1. Separate Read/Write Connection Pools**

Moonlogs creates two separate connection pools to SQLite:
* Since SQLite does not support concurrent writes, the write pool utilizes only one active/idle connection. This allows Moonlogs to delegate synchronization to the Golang side using built-in mutexes, reducing the average response time from SQLite without requiring blocking procedures.

* The read pool, on the other hand, uses the number of connections to the database equal to the number of CPU cores on the machine running Moonlogs. This efficiently allows for parallel reading of data from SQLite without overloading it with concurrent access.

**2. Usage of WAL (Write-Ahead Logging)**

`WAL` allows SQLite to perform insertions into the database through a separate journal controlled by SQLite itself. This avoids blocking read or write operations to the database file for `COMMIT`, enabling more concurrent transactions.

**3. Usage of synchronous=NORMAL pragma**

By default, SQLite uses `FULL` synchronization when writing data to the database. This means SQLite will wait for the OS to confirm that the data has been successfully written.

The `NORMAL` setting for synchronous in SQLite means that the database engine will write data to disk at critical moments but not with every transaction commit, which can improve performance compared to `FULL` synchronization. Paired with `WAL` mode, this ensures that transactions are ACID-compliant and data integrity is maintained.

**4. In-Memory Caching**

Moonlogs utilizes SQLite's built-in mechanisms for in-memory caching of database pages with `cache_size` pragma and instructs SQLite to store temporary data in memory rather than on disk with `temp_store=memory`. This further reduces the frequency of disk I/O operations.

**Summary**

All of this allows Moonlogs to use SQLite as a balanced yet efficient solution for production environments.
