/*
@Time : 2019-12-21 15:39
@Author : lfn
@File : select
*/

package session

import (
	"net/http"

	"github.com/pioneerlfn/12306/pkg/login"
)

type Session struct {
	StationSeat []string
	AutoCodeType int
	Client *http.Client
	Urls map[string]map[string]interface{}
	Login login.GoLogin
	CDNList []string
	Cookies string
	QueryURL string
	PassengerTicketStrList string
	PassengerTicketStrByAfterLate string
	OldPassengerStr string
	SetType int
	Flag bool
}

func NewSelect() *Session {
	return &Session{
		Client:new(http.Client),
	}
}

func (s *Session) Config() error {

	return nil
}

func (s *Session) Run() {
	err := s.LogIn()
	if err != nil {
		panic(err)
	}
	s.QueryTickets()
	s.Order()
}

func (s *Session) LogIn() error {
	err := s.login()
	if err != nil {
		return err
	}
	return nil
}

func (s *Session) login() error {
	var err error
	err = s.setCookies()
	if err != nil {
		return err
	}
	err = s.setBaseAuth()
	if err != nil {
		return err
	}
	err = s.setCaptha()
	if err != nil {
		return err
	}
	err = s.send(request)
	if err != nil {
		return err
	}

	return nil
}

func get_ticket_info() {

}