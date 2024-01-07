package tickets

import (
	"errors"
	"fmt"
)

type Ticket struct {
	Id          int
	Name        string
	Email       string
	Destination string
	Time        string
	Price       int
}

var ErrorInvalidField error = errors.New("invalid field")

func CheckTicketContent(ticket Ticket) (err error) {
	if ticket.Id == 0 {
		err = fmt.Errorf("id %w", ErrorInvalidField)
	}
	if ticket.Name == "" {
		err = fmt.Errorf("name %w", ErrorInvalidField)
	}
	if ticket.Email == "" {
		err = fmt.Errorf("email %w", ErrorInvalidField)
	}
	if ticket.Destination == "" {
		err = fmt.Errorf("destination %w", ErrorInvalidField)
	}
	if ticket.Time == "" {
		err = fmt.Errorf("time %w", ErrorInvalidField)
	}
	if ticket.Price == 0 {
		err = fmt.Errorf("price %w", ErrorInvalidField)
	}
	return
}
