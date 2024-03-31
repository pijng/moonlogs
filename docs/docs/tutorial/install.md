# Installation

There are currently two options for installing Moonlogs:
* With [APT Repository](#using-apt-repository)
* With [Docker Compose](#using-docker-compose)

## Using APT Repository

1. **Add APT Repository**

```bash
sudo add-apt-repository "deb [trusted=yes] https://apt.fury.io/pijng/ /"
sudo apt-get update
```

This will add the Moonlogs repository to your `/etc/apt/sources.list.d/`.


2. **Install moonlogs package**

```bash
apt-get install moonlogs
```
Moonlogs will be immediately available as a service managed by systemd.

3. **Check the status of the service**

```bash
service moonlogs status
```

Navigate to [http://localhost:4200](http://localhost:4200). You should see the moonlogs UI. Follow the instructions for creating the initial administrator there.


## Using Docker Compose

::: info
Make sure you have [Docker](https://docs.docker.com/engine/install/) and [Docker Compose](https://docs.docker.com/compose/install/) installed on your system.

:::

**1. Create a Docker Compose file**

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
      - moonlogs-config:/etc/moonlogs
      - moonlogs-data:/var/lib/moonlogs
    command: --port=4200
volumes:
  moonlogs-config:
  moonlogs-data:
```

**2. Build and run the Docker containers**

Run the following command in the directory with `docker-compose.yml`:

```bash
docker-compose up -d
```

This will download the moonlogs image and start the moonlogs container in detached mode.

**3. Access moonlogs**

Navigate to [http://localhost:4200](http://localhost:4200). You should see the moonlogs UI. Follow the instructions for creating the initial administrator there.

## What's next?

* Please refer to the [Configuration](/tutorial/configuration) section to understand the configuration options supported by Moonlogs and adjust them to suit your requirements.

* Refer to the guide [Choosing DB Adapter](/tutorial/choosing-db-adapter) to better understand which database suits your requirements.

* Or go straight to [Basic usage guide](/usage/basics)