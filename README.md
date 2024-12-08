## Prerequisites

- docker-compose
- docker
- vegeta (https://github.com/tsenart/vegeta)

## How to run the applications

1. Run docker-compose command

```
docker-compose up -d
```

2. To generate http load, run these following command

```
bash generate_http_load.sh
```

3. Grafana URL

```
Localhost:3000
```

4. Cleanup

```
docker-compose down
```

## Rate limit configuration

- MAX_REQUEST: maximum overall request rate
- TOKEN_BUCKET_RATE_PER_SECOND: define the token bucket refill rate (r tokens per second)
