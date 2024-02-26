# Not Found Config

## Descripcion

Cuando se solicita una configuración y el driver no la tiene implementada, se debe devolver una respuesta para que no se interprete como una desconexión del driver o un error de red.

## Request Message

```json
{
    ...
    "data": {...}
}
```



## Response Message
```json
{
    ...,
    "data": {
     "status" : 404,
     "msg": "'action' not found on the driver",
    }
}
```

| Campo | Tipo | Descripcion |
| --- | --- | --- |
| status | int | status de la respuesta  |
| msg | string | mensaje de problema o log sucedido  |
