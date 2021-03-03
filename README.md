# Dummy API written in Go and secured by a traefik reverse proxy

This docker-compose contains a dummy API (written in Go) that responds to POST and GET requests. The API is externally accessible via a reverse proxy (traefik) on port 443 (self signed cert). The reverse proxy allows from outside only a GET request, from inside (docker-compose network) also a POST request which can be sent from the traefik container or any other container within the docker-compose network.

                           ┌───────────────────────────────────────────────────────────────┐
                           │ docker-compose local network "proxy"                          │
                           │                                                               │
                           │       ┌────────────────┐                 ┌────────────────┐   │
                           │       │                │                 │                │   │
          only GET allowed │       │                │    GET, POST    │                │   │
    iNet ──────────────────┼──────►│     traefik    ├────────────────►│    dummyapi    │   │
                           │  :443 │  reverse proxy │            :6543│                │   │
                           │       │                │                 │                │   │
                           │       │                │                 │                │   │
                           │       └────────────────┘                 └────────────────┘   │
                           │                                                               │
                           │                                                               │
                           └───────────────────────────────────────────────────────────────┘

## Prerequisite

- a useful OS with command line
- curl
- docker (20.10.5 tested on Linux)
- docker-compose (1.27.4 tested on Linux)

## Usage

In addition to the traefik dashboard, the container management portainer UI is also provided.

### docker-compose

Start everything. Ignore the `dummyapi was built` Warning.

```bash
docker-compose up
````

### Dummy API

#### Request from outside/internet

Remember, GET is allowed from outside, but POST is not.

##### GET

Will return `"GetMethod called"`

```bash
curl --insecure -X GET https://dummyapi.docker.localhost
````

##### POST

Won't return anything

```bash
curl --insecure -X POST https://dummyapi.docker.localhost
````

#### Request from the docker-compose network "proxy"

We send the requests from the traefik container, but we need to install curl first.

##### Install curl

Because of the self signed certificate `./traefik/cert`, we have to call curl with `--insecure` 

```bash
docker exec -it traefik apk add --update curl
```

##### GET

Will return `"GetMethod called"`

```bash
docker exec -it traefik curl --insecure -X GET dummyapi.docker.localhost:6543
```

##### POST

Will return `"PostMethod called"`

```bash
docker exec -it traefik curl --insecure -X POST dummyapi.docker.localhost:6543
```

## traefik

Use the test login credentials `testuser`, `testpassword` which are provided by the line 32 in `./traefik/dynamic.yml` . You can create new credentials with `echo $(htpasswd -nb <new_username> <new_password>)`

The traefik dashboard can be found here: https://traefik.docker.localhost/

You have to ignore/accept the certificate warning.

## portainer

The Portainer dashboard can be found here: https://portainer.docker.localhost/

You have to ignore/accept the certificate warning.

At the first call you have to assign a new password and then select `Docker` as the environment you want to manage.

