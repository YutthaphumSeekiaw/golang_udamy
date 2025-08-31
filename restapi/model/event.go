package model

import (
	"fmt"
	"restapi/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `json:"time"`
	UserId      int64     `json:"userid"`
}

var events = []Event{}

func (e *Event) Save() error {
	query := `
	  insert into events(name,description,location,dateTime,userId) values
(?,?,?,?,?)
	`
	smtm, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer smtm.Close()
	res, err := smtm.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()

	e.ID = id

	events = append(events, *e)

	return err
}

func GetAllEvents() ([]Event, error) {
	query := `select * from events`
	row, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var events []Event

	fmt.Println(row)
	for row.Next() {
		var event Event
		err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		fmt.Println(event)
		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := `select * from events where id = ?`
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (e Event) UpdateEvents() error {
	query := `update events set
	  name = ?,
	  description = ?,
	  location = ?,
	  dateTime = ?,
	  userId = ?
	  where id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)

	return err
}

func (e Event) DeleteEvent(id int64) error {
	query := `delete from events where id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	return err
}

func (e Event) Register(userId int64) error {
	query := `insert into registrations(eventId,userId) values (?,?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)
	if err != nil {
		return err
	}

	return nil
}

func (e Event) CancelRegis(userId int64) error {
	query := `delete from registrations where eventId = ? and  userId = ?`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)
	if err != nil {
		return err
	}

	return nil

}
