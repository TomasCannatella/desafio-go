package repository

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

var (
	ErrTicketExist        error = errors.New("this ticket already exist")
	ErrInvalidDestination error = errors.New("invalid destination")
	ErrInvalidPeriod      error = errors.New("invalid period")
	ErrInvalidFormatTime  error = errors.New("invalid time format")
	ErrEmptyTickets       error = errors.New("no ticket record")
	ErrInvalidFormatHours error = errors.New("invalid hour format")
)

type TicketRepository struct {
	Tickets []tickets.Ticket
}

func NewTicketRepository() *TicketRepository {
	return &TicketRepository{}
}

func (t *TicketRepository) LoadTickets() error {

	//Abro el csv
	f, err := os.Open("./tickets.csv")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer f.Close()

	//leo el csv
	rd := csv.NewReader(f)

	rd.Read()

	for {
		record, err := rd.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println(err)
			return err
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			fmt.Println(err)
			return err
		}

		price, err := strconv.Atoi(record[5])
		if err != nil {
			fmt.Println(err)
			return err
		}

		ticket := tickets.Ticket{
			Id:          id,
			Name:        record[1],
			Email:       record[2],
			Destination: record[3],
			Time:        record[4],
			Price:       price,
		}

		if err := tickets.CheckTicketContent(ticket); err != nil {
			return fmt.Errorf("%w: Error when loading ticket", err)
		}

		t.Tickets = append(t.Tickets, ticket)
	}

	return nil
}

// ejemplo 1
func (t *TicketRepository) GetTotalTickets(destination string) (int, error) {
	if destination == "" {
		return 0, fmt.Errorf("%w: This destination is null", ErrInvalidDestination)
	}

	countDestination := 0
	for _, d := range t.Tickets {
		if !(d.Destination == destination) {
			continue
		}
		countDestination++
	}
	return countDestination, nil
}

// ejemplo 2
func getMornings(hour int) (period string, startTime int, finishTime int, err error) {

	if hour >= 0 && hour <= 6 {
		period = "early morning"
		startTime = 0
		finishTime = 6
	} else if hour >= 7 && hour <= 12 {
		period = "morning"
		startTime = 7
		finishTime = 12
	} else if hour >= 13 && hour <= 19 {
		period = "afternoon"
		startTime = 13
		finishTime = 19
	} else if hour >= 20 && hour <= 23 {
		period = "night"
		startTime = 20
		finishTime = 23
	} else {
		err = fmt.Errorf("%w: The hour is out range", ErrInvalidPeriod)
	}

	return
}

func getHours(time string) (hour int, err error) {
	match, _ := regexp.MatchString("^([01]?[0-9]|2[0-3]):[0-5][0-9]$", time)
	if !match {
		err = ErrInvalidFormatTime
		hour = 0
		return
	}

	fullTime := strings.Split(time, ":")
	hour, err = strconv.Atoi(fullTime[0])
	if err != nil {
		err = fmt.Errorf("%w: Cannot format the time", ErrInvalidFormatHours)
		return
	}
	return
}

func (t *TicketRepository) GetCountByPeriod(time string) (quantityPerson int, period string, err error) {

	hour, err := getHours(time)
	if err != nil {
		if errors.Is(err, ErrInvalidFormatHours) {
			return 0, "", err
		} else if errors.Is(err, ErrInvalidFormatTime) {
			return 0, "", err
		} else {
			return 0, "", errors.New("unexpected error")
		}
	}

	period, startTime, finishTime, err := getMornings(hour)

	for _, ticket := range t.Tickets {
		hourTicket, err := getHours(ticket.Time)
		if err != nil {
			return 0, "", ErrInvalidFormatHours
		}
		if hourTicket >= startTime && hourTicket <= finishTime {
			quantityPerson++
		}
	}

	return
}

// ejemplo 3
func (t *TicketRepository) AverageDestination(destination string) (float64, error) {

	totalTicket := len(t.Tickets)

	if totalTicket == 0 {
		return 0, ErrEmptyTickets
	}

	totalDestinationTicket, err := t.GetTotalTickets(destination)
	if err != nil {
		return -2, err
	}

	average := float64(totalDestinationTicket) / float64(totalTicket)
	return average, nil

}
