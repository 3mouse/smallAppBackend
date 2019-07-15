package mysql

import (
	"database/sql"
	"fmt"
)

const (
	username = "admin"
	password = "pass"
	tcpAddress = "127.0.0.1:3306"
)

type Operator struct {
	db *sql.DB
}

func (p *Operator) Insert(sql string, args ...string) error {
	stmt, err := p.db.Prepare(sql)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(args)
	if err != nil {
		return err
	}

	return nil
}

func (p *Operator) Delete(sql string, args ...string) error {
	stmt, err := p.db.Prepare(sql)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(args)
	if err != nil {
		return err
	}

	return nil
}

func (p *Operator) Update(sql string, args ...string) error {
	stmt, err := p.db.Prepare(sql)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(args)
	if err != nil {
		return err
	}

	return nil
}

func (p *Operator) Query(sql string, args ...string) ([]string, error) {
	res := make([]string, 0)
	rows, err := p.db.Query(sql, args)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var r string
		err := rows.Scan(&r)
		if err != nil {
			return nil, err
		}

		res = append(res, r)
	}

	return res, nil
}

func NewOperator(dbname string) (*Operator, error) {
	ds := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, tcpAddress, dbname)
	db, err := sql.Open("mysql", ds)
	if err != nil {
		return nil, err
	}

	return &Operator{db}, nil
}

func (p *Operator) CloseDB() error {
	return p.db.Close()
}