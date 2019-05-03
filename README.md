# Peso API

#### Calculo de IMC
`GET` /imc?height=1.70&weight=70
```json
{
  "data": {
    "imc": "24.22"
  },
  "message": "IMC calculated with success"
}
```

#### Calculo de peso ideal min & max
`GET` /weights?height=1.70
```json
{
  "data": {
    "max_weight": "72.22",
    "min_weight": "53.46"
  },
  "message": "Weights calculated with success"
}
```
