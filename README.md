# desafio-go-bases

## Planteo 
Una aerolínea pequeña tiene un sistema de reservas de pasajes a diferentes países, este retorna un archivo con la información de los pasajes sacados en las últimas 24 horas.
La aerolínea necesita un programa para extraer información de las ventas del día y así, analizar las tendencias de compra.
El archivo en cuestión es del tipo valores separados por coma (csv por su siglas en inglés), donde los campos están compuestos por: <b> id, name, email, destination country, flight time y price. </b>

## Desafio
Realizar un programa que sirva como herramienta para calcular diferentes datos estadísticos. Para lograrlo, debes clonar este repositorio que contiene un archivo .csv con datos generados y un esqueleto del proyecto.

##### Requerimiento 1:

Una función que calcule cuántas personas viajan a un país determinado:
``` go
func GetTotalTickets(destination string) (int, error) {}
```

##### Requerimiento 2:
Una o varias funciones que calcule cuántas personas viajan en madrugada <b>(0 → 6), mañana (7 → 12), tarde (13 → 19) y noche (20 → 23):</b>
```go
func GetCountByPeriod(time string) (int, error) {}
```	

##### Requerimiento 3:
Calcular el porcentaje de personas que viajan a un país determinado en un dia, con respecto al resto:
```go
func AverageDestination(destination string, total int) (float64, error) {}
```

##### Requerimiento 4:
Crear test unitarios para cada uno de los requerimientos anteriores, mínimo 2 casos por requerimiento:
```go 
import "testing"
func TestGetTotalTickets(t *testing.T) {}
```