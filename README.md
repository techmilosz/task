# UI
Documentation along with swagger ui is available at `/docs`, while raw documentation (oapi.yaml) is available at `/docs/oapi`

# Running
## Running locally - without other tools
```
make run
```

## Running locally - with docker
```
make run-docker
```

# Endpoints
## Packs
### Get all packs
URL:
```
GET /packs
```

### Add pack
URL:
```
POST /packs
```

Payload:
```
{
    "value": [value]
}
```

### Delete pack
URL:
```
DELETE /packs/{value}
```

## Order
URL:
```
GET /order/{order}
```




