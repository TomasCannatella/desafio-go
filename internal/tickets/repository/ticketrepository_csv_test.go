package repository_test

import (
	"errors"
	"testing"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
	"github.com/bootcamp-go/desafio-go-bases/internal/tickets/repository"
	"github.com/stretchr/testify/require"
)

func TestGetTotalTickets(t *testing.T) {
	t.Run("success - Total tickets from Argentina are 3", func(t *testing.T) {
		ticketRepository := repository.TicketRepository{}

		ticketRepository.Tickets = append(ticketRepository.Tickets,
			tickets.Ticket{
				Id:          01,
				Name:        "John",
				Destination: "Argentina",
				Time:        "20:19",
				Price:       200,
			}, tickets.Ticket{
				Id:          01,
				Name:        "Jane",
				Destination: "Argentina",
				Time:        "6:25",
				Price:       200,
			}, tickets.Ticket{
				Id:          01,
				Name:        "Joe",
				Destination: "Argentina",
				Time:        "1:30",
				Price:       200,
			})

		result, err := ticketRepository.GetTotalTickets("Argentina")

		if err != nil {
			t.Fatal("Error: ", err)
		}

		expectedResult := 3

		require.Equal(t, result, expectedResult)
	})

	t.Run("error - The destination entered is default value ", func(t *testing.T) {
		ticketRepository := repository.TicketRepository{}

		result, err := ticketRepository.GetTotalTickets("")

		if err != nil {
			if !errors.Is(err, repository.ErrInvalidDestination) {
				t.Fatal("Error: ", err)
			}
		}

		expectedResult := 0

		require.Equal(t, result, expectedResult)
	})
}

func TestGetCountByPeriod(t *testing.T) {
	t.Run("success - the time and period returned is correct", func(t *testing.T) {
		ticketRepository := repository.TicketRepository{}

		ticketRepository.Tickets = append(ticketRepository.Tickets,
			tickets.Ticket{
				Id:          01,
				Name:        "John",
				Destination: "Argentina",
				Time:        "00:19",
				Price:       200,
			}, tickets.Ticket{
				Id:          01,
				Name:        "Jane",
				Destination: "Argentina",
				Time:        "6:25",
				Price:       200,
			}, tickets.Ticket{
				Id:          01,
				Name:        "Joe",
				Destination: "Argentina",
				Time:        "1:30",
				Price:       200,
			},
		)

		quantityPerson, period, err := ticketRepository.GetCountByPeriod("04:00")

		if err != nil {
			t.Fatal("Error:", err)
		}
		quantityPersonExpected, periodExpected := 3, "early morning"

		require.Equal(t, quantityPerson, quantityPersonExpected)
		require.Equal(t, period, periodExpected)
	})

	t.Run("error - Time entered has an invalid format", func(t *testing.T) {
		ticketRepository := repository.TicketRepository{}

		ticketRepository.Tickets = append(ticketRepository.Tickets,
			tickets.Ticket{
				Id:          01,
				Name:        "John",
				Destination: "Argentina",
				Time:        "00:19",
				Price:       200,
			}, tickets.Ticket{
				Id:          01,
				Name:        "Jane",
				Destination: "Argentina",
				Time:        "6:25",
				Price:       200,
			}, tickets.Ticket{
				Id:          01,
				Name:        "Joe",
				Destination: "Argentina",
				Time:        "1:30",
				Price:       200,
			},
		)

		quantityPerson, period, err := ticketRepository.GetCountByPeriod("04:0O")

		quantityPersonExpected, periodExpected, errExpected := 0, "", repository.ErrInvalidFormatTime

		if err != nil {
			if !errors.Is(err, errExpected) {
				t.Fatal("Error: ", err)
			}
		}

		require.Equal(t, quantityPerson, quantityPersonExpected)
		require.Equal(t, period, periodExpected)

	})
}

func TestAverageDestination(t *testing.T) {

	t.Run("success - Ticket repository is not empty and average is 0.5", func(t *testing.T) {
		ticketRepository := repository.TicketRepository{}
		ticketRepository.Tickets = append(ticketRepository.Tickets,
			tickets.Ticket{
				Id:          01,
				Name:        "John",
				Destination: "Argentina",
				Time:        "00:19",
				Price:       200,
			}, tickets.Ticket{
				Id:          02,
				Name:        "Jane",
				Destination: "Argentina",
				Time:        "6:25",
				Price:       200,
			}, tickets.Ticket{
				Id:          03,
				Name:        "Joe",
				Destination: "Argentina",
				Time:        "1:30",
				Price:       200,
			}, tickets.Ticket{
				Id:          04,
				Name:        "John",
				Destination: "Finland",
				Time:        "00:19",
				Price:       200,
			}, tickets.Ticket{
				Id:          05,
				Name:        "Jane",
				Destination: "Poland",
				Time:        "6:25",
				Price:       200,
			}, tickets.Ticket{
				Id:          06,
				Name:        "Joe",
				Destination: "Ukranie",
				Time:        "1:30",
				Price:       200,
			},
		)

		result, err := ticketRepository.AverageDestination("Argentina")
		if err != nil {
			t.Fatal("Error: ", err)
		}
		expectedResult := 0.5
		require.Equal(t, result, expectedResult)
	})

	t.Run("error - TicketRepository is empty", func(t *testing.T) {
		ticketRepository := repository.TicketRepository{}

		result, err := ticketRepository.AverageDestination("Argentina")
		if err != nil {
			if !errors.Is(err, repository.ErrEmptyTickets) {
				t.Fatal("Error: ", err)
			}
		}

		expectedResult := 0.0

		require.Equal(t, result, expectedResult)
	})

}
