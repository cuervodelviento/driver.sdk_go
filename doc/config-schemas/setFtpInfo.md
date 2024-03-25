# Get ftp info

## Descripcion

Modificar la informacion de la configuracion ftp del dispositivo

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "setFtpInfo",
        "value": {
            "id": "abc123",
            "host": "1.1.1.1",
            "user": "admin",
            "password": "admin",
            "directory": "/directory",
            "port": 21,
        } 
    }
}
```

| Campo | Tipo | Descripcion |
| --- | --- | --- |
| id | string | uuid |
| host | string | ip del servidor ftp |
| user | string | usuario ftp |
| password | string | contraseña ftp |
| directory | string | directorio en cual se desea almacenar guardar |
| port | int | puerto del servidor ftp |




## Response Message
```json
{
    ...,
    "data": {
        "error" : false,
        "msg" : ""
    }
}
```

| Campo | Tipo | Descripcion |
| --- | --- | --- |
| error | boolean | Error al hacer backup |
| msg | string | Mensaje de error o log interno|