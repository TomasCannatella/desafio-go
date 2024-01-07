/*
-------------------------------------- PLANTEO --------------------------------------
Una aerolínea pequeña tiene un sistema de reservas de pasajes a diferentes países, este retorna un archivo con la información de los pasajes sacados en las últimas 24 horas.
La aerolínea necesita un programa para extraer información de las ventas del día y así, analizar las tendencias de compra.
El archivo en cuestión es del tipo valores separados por coma (csv por su siglas en inglés),
donde los campos están compuestos por: id, name, email, destination country, flight time y price.
-------------------------------------- DESAFIO --------------------------------------
Realizar un programa que sirva como herramienta para calcular diferentes datos estadísticos.
Para lograrlo, debes clonar este repositorio que contiene un archivo .csv con datos generados
y un esqueleto del proyecto.

Requerimiento 1:

	Una función que calcule cuántas personas viajan a un país determinado:
	func GetTotalTickets(destination string) (int, error) {}

Requerimiento 2:
	Una o varias funciones que calcule cuántas personas viajan en madrugada (0 → 6), mañana (7 → 12), tarde (13 → 19) y noche (20 → 23):
	func GetCountByPeriod(time string) (int, error) {}

Requerimiento 3:
	Calcular el porcentaje de personas que viajan a un país determinado en un dia, con respecto al resto:
	func AverageDestination(destination string, total int) (float64, error) {}

Requerimiento 4:
	Crear test unitarios para cada uno de los requerimientos anteriores, mínimo 2 casos por requerimiento:
	import "testing"
	func TestGetTotalTickets(t *testing.T) {}
*/

package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets/repository"
)

func main() {
	var ticketRepository repository.TicketRepository = repository.TicketRepository{}
	ticketRepository.LoadTickets()

	destination := "Finland"
	time := "10:30:00"

	cantDestination, err := ticketRepository.GetTotalTickets(destination)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Number of destinations to ", destination, ": ", cantDestination)

	quantityPerson, period, err := ticketRepository.GetCountByPeriod(time)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("The number of people traveling in the ", period, " is: ", quantityPerson)

	average, err := ticketRepository.AverageDestination(destination)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("The average number of people traveling to  ", destination, "with respect to others are: ", average)

}
