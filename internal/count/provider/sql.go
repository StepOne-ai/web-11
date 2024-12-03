package provider

import (
	"database/sql"
	"errors"
)

func (p *Provider) FetchCount() (int, error) {
	var msg int

	err := p.conn.QueryRow("SELECT number FROM count").Scan(&msg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, err
	}

	return msg, nil
}

func (p *Provider) CheckCountExist() (bool, error) {
	msg := 1
	err := p.conn.QueryRow("SELECT number FROM count WHERE number_id = $1", msg).Scan(&msg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (p *Provider) UpdateCount(count int) error {
	_, err := p.conn.Exec("UPDATE count SET number = number + $1 WHERE number_id = 1", count)
	if err != nil {
		return err
	}

	return nil
}
