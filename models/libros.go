package models

import "errors"

type Libro struct {
	ID     int    `json:"id"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
	Stock  int    `json:"stock"`
}

func (l *Libro) ReducirStock() error {
	if l.Stock <= 0 {
		return errors.New("no hay stock disponible")
	}
	l.Stock--
	return nil
}
