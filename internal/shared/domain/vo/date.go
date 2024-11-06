package vo

import "time"

//No hi ha comprovacions perque les comprovacions es fan a traves dels validators de gin-gonic

type Date time.Time

func NewDate(t *time.Time) Date {
	d := Date(*t)
	return d
}

func (this *Date) Value() *time.Time {
	t := time.Time(*this)
	return &t
}

func (this *Date) ToString() string {
	t := time.Time(*this).String()
	return t
}
