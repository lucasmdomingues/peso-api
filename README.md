# Peso API

### Como usar

`GET` /imc?height=1.70&weight=70 - Calcula o IMC
```json
{
  "data": {
    "imc": "24.22"
  },
  "message": "IMC calculated with success"
}
```


`GET` /weights?height=1.70 - Calcula o peso min & max
```json
{
  "data": {
    "max_weight": "72.22",
    "min_weight": "53.46"
  },
  "message": "Weights calculated with success"
}
```
