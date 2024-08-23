# Obtener subdispositivos

## Descripcion

Devolver una lista de dispositivos que estan conectados al dispositivo principal.

## Request Message

```json
{
    ...,
    "data": {
        "deviceData": {...},
        "configKey": "getSubdevices",
        "value": "{}"
    }
}
```

## Response Message

```json
{
    ...,
    "data": [{
        "id_brand": 9,
        "id_model": 48,
        "id_manufacturer": 13,
        "id_sub_system": 2,
        "id_device_group": 3,
        "child_id": "83df4a9f-e278-4b11-a014-7b2499450002",
        "child_name":"Backyard PIR",
        "child_description":"PIR set to avoid pets under 25Kg",
        "child_ip":"127.0.0.5",
        "child_port":9098
    }]
}
```

| Campo             | Tipo   | Descripcion                       |
| ----------------- | ------ | --------------------------------- |
| id_brand          | int    | Id de la marca del dispositivo    |
| id_model          | int    | Id del modelo del dispositivo     |
| id_manufacturer   | int    | Id del fabricante del dispositivo |
| id_sub_system     | int    | Id del subsistema del dispositivo |
| id_device_group   | int    | Id del grupo de dispositivos      |
| child_id          | string | Id del subdispositivo             |
| child_name        | string | Nombre del dispositivo            |
| child_description | string | Descripicon del dispositivo       |
| child_ip          | string | ip del dispositivo                |
| child_port        | int    | puerto del dispositivo            |
