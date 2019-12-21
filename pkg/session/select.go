/*
@Time : 2019-12-21 15:39
@Author : lfn
@File : select
*/

package session

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Session struct {
	StationSeat                   []string
	AutoCodeType                  int
	Client                        *http.Client
	Urls                          map[string]map[string]interface{}
	Login                         login.GoLogin
	CDNList                       []string
	Cookies                       string
	QueryURL                      string
	PassengerTicketStrList        string
	PassengerTicketStrByAfterLate string
	OldPassengerStr               string
	SetType                       int
	Flag                          bool
}

func NewSelect() *Session {
	return &Session{
		Client: new(http.Client),
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


/*func (s *Session) login() error {
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
*/
func get_ticket_info() {

}

func (s *Session) SetCookies() error {
	cookie, err := GetCookies()
	log.Println("len(cookies):", len(cookie))
	if err != nil {
		return fmt.Errorf("login:%w", err)
	}
	var cookieVal []string
	cookieVal = append(cookieVal, "BIGipServerpool_statistics=635503114.44582.0000")
	for _, ck := range cookie {
		fmt.Println(ck.Name, ck.Value, ck.Expiry, ck.Domain)
		if ck.Name == railExpiration {
			cookieVal = append(cookieVal, strings.Join([]string{ck.Name, ck.Value}, "="))
		}
	}
	return nil
}
