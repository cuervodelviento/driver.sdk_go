> _[!NOTE]_

```mermaid
graph TD;
    A-->B;
    A-->C;
    B-->D;
    C-->D;
```

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
    "id" : 0,
    "enabled": true,
    "userName": "test",
    }]
}
```

| Campo | Tipo | Descripcion |
| --- | --- | --- |
| id | int | index del item en el driver  |
| userName | string | nombre del usuario  |
| enabled | boolean | está o no habilitado |
