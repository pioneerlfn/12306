/*
@Time : 2019-12-22 16:12
@Author : lfn
@File : type
*/

package session

type GoLogin struct{}


type LoginConf struct {
	ValidateMessagesShowID string `json:"validateMessagesShowId"`
	Status                 bool   `json:"status"`
	Httpstatus             int    `json:"httpstatus"`
	Data                   struct {
		IsstudentDate   bool     `json:"isstudentDate"`
		IsLoginPassCode string   `json:"is_login_passCode"`
		IsSweepLogin    string   `json:"is_sweep_login"`
		PsrQrCodeResult string   `json:"psr_qr_code_result"`
		LoginURL        string   `json:"login_url"`
		StudentDate     []string `json:"studentDate"`
		StuControl      int      `json:"stu_control"`
		IsUamLogin      string   `json:"is_uam_login"`
		IsLogin         string   `json:"is_login"`
		OtherControl    int      `json:"other_control"`
	} `json:"data"`
	Messages         []interface{} `json:"messages"`
	ValidateMessages struct {
	} `json:"validateMessages"`
}

type CaptchaRes struct {
	Image         []byte `json:"image"`
	ResultMessage string `json:"result_message"`
	ResultCode    string `json:"result_code"`
}