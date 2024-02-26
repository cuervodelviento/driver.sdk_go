# Get admin user config

## Descripcion

Get data del usuario admin 

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "getAdminUser",
        "value": "{}"  // No necesita
    }
}
```



## Response Message
```json
{
    ...,
    "data": "[ {
    "username": "username",
    }]"
}
```

Solo el nombre de usuario para no revelear el password del admin

| Campo | Tipo | Descripcion |
| --- | --- | --- |
| userName | string | nombre del usuario  |
