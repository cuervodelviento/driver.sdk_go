# Get all people from AC

## Descripcion

Devolver una lista de todas las personas que estan registradas en el dispositivo.

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "getAllPeopleFromAC",
        "value": "{}" 
    }
}
```

El Request Message de `getAllPeopleFromAC` no necesita value. El driver solo con el configKey identificara la operacion que debe realizar.

## Response Message
```json
{
    ...,
    "data": [{
        "personId": "em2191",
        "personName": "John Doe",
        "totalCards": 2,
        "totalFaces": 1,
        "totalFingerprints": 0,
        "totalIris": 0,
        "totalVoiceRecognition": 0,
        "totalRFID": 0,
    }]
}
```

| Campo | Tipo | Descripcion |
| --- | --- | --- |
| personId | string | El id de la persona |
| personName | string | El nombre de la persona |
| totalCards | number | El total de tarjetas registradas |
| totalFaces | number | El total de rostros registrados |
| totalFingerprints | number | El total de huellas registradas |
| totalIris | number | El total de iris registrados |
| totalVoiceRecognition | number | El total de reconocimientos de voz registrados |
| totalRFID | number | El total de RFID registrados |

