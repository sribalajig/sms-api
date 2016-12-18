package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"plivo/plivo.sms/model"
	"plivo/plivo.sms/service"
	"plivo/plivo.sms/validator"
)

type SMSController struct {
	beego.Controller
}

func (smsController *SMSController) Inbound() {
	var sms model.SMS

	json.Unmarshal(smsController.Ctx.Input.RequestBody, &sms)

	err := service.ProcessInboundSMS(sms)

	if err == nil {
		smsController.Data["json"], _ = json.Marshal(validator.ValidationResult{Message: "inbound sms ok”"})
	} else {
		smsController.Data["json"], _ = json.Marshal(validator.ValidationResult{Error: "unknown failure””"})
	}

	smsController.ServeJSON()
}

func (smsController *SMSController) Outbound() {
	var sms model.SMS

	json.Unmarshal(smsController.Ctx.Input.RequestBody, &sms)

	err := service.ProcessOutboundSMS(sms)

	if err == nil {
		smsController.Data["json"], _ = json.Marshal(validator.ValidationResult{Message: "outbound sms ok”"})
	} else {
		smsController.Data["json"], _ = json.Marshal(validator.ValidationResult{Error: "unknown failure””"})
	}

	smsController.ServeJSON()
}
