# Get user config

## Descripcion

Lista de usuarios con credenciales 

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "getUsers",
        "value": "{}"  // No necesita
    }
}
```


## Response Message
```json
{
    ...,
    "data": [ {
    "id" : "68e5cc48-e23f-46b5-a491-2a623fccf658",
    "username": "username",
    "status": 0,
    "created_at": "2021-08-12T00:00:00Z",
    "last_login": "2021-08-12T00:00:00Z",
    "last_logout": "2021-08-12T00:00:00Z",
    "last_ip": "127.0.0.1",
    "is_admin_user": true
    }]
}
```

Solo el nombre de usuario para no revelear el password del admin

| Campo | Tipo | Descripcion |
| --- | --- | --- |
| id | string | uuid |
| userName | string | nombre del usuario  |
| status | int | Estado del usuario, -1: Unknown, 0: Online, 1: Offline |
| created_at | string or null | Fecha de creacion del usuario en formato ISO 8601 |
| last_login | string or null | Fecha del ultimo login del usuario en formato ISO 8601 |
| last_logout | string or null | Fecha del ultimo logout del usuario en formato ISO 8601 |
| last_ip | string or null | IP del ultimo login del usuario |
