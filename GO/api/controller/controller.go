package controller

import (
	"justasking/GO/api/controller/account"
	"justasking/GO/api/controller/boxes/answerbox"
	"justasking/GO/api/controller/boxes/basebox"
	"justasking/GO/api/controller/boxes/questionbox"
	"justasking/GO/api/controller/boxes/votesbox"
	"justasking/GO/api/controller/boxes/wordcloud"
	"justasking/GO/api/controller/featurerequest"
	"justasking/GO/api/controller/idpjustasking"
	"justasking/GO/api/controller/priceplan"
	"justasking/GO/api/controller/sms"
	"justasking/GO/api/controller/stripe"
	"justasking/GO/api/controller/support"
	"justasking/GO/api/controller/theme"
	"justasking/GO/api/controller/token"
	"justasking/GO/api/controller/twilio"
	"justasking/GO/api/controller/user"
)

// LoadRoutes loads the routes for the controllers
func LoadRoutes() {
	tokencontroller.Load()
	usercontroller.Load()
	wordcloudcontroller.Load()
	questionboxcontroller.Load()
	answerboxcontroller.Load()
	votesboxcontroller.Load()
	baseboxcontroller.Load()
	themecontroller.Load()
	smscontroller.Load()
	twiliocontroller.Load()
	priceplancontroller.Load()
	accountcontroller.Load()
	stripecontroller.Load()
	supportcontroller.Load()
	featurerequestcontroller.Load()
	idpjustaskingcontroller.Load()
}
