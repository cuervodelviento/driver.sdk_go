# Gets users from an alarm

## Descripcion

Devuelve una lista de usuarios de una alarma

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "getAlarmUsers",
        "value":{}//not needed
    }
}

```

## Response Message

```json
{
    ...,
    "data": {
       "partitionId":"xxx-xxxx",
        "userName": "Em-32",
        "userId:"xxx-xxxx",
    }
}
```

| Campo       | Tipo   | Descripcion        |
| ----------- | ------ | ------------------ |
| partitionId | string | Id de la particion |
| userId      | string | Id de el usuario   |
| userName    | string | Nombre del usuario |
