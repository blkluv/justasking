package controller

import (
	accountcontroller "github.com/chande/justasking/api/controller/account"
	answerboxcontroller "github.com/chande/justasking/api/controller/boxes/answerbox"
	baseboxcontroller "github.com/chande/justasking/api/controller/boxes/basebox"
	questionboxcontroller "github.com/chande/justasking/api/controller/boxes/questionbox"
	votesboxcontroller "github.com/chande/justasking/api/controller/boxes/votesbox"
	wordcloudcontroller "github.com/chande/justasking/api/controller/boxes/wordcloud"
	featurerequestcontroller "github.com/chande/justasking/api/controller/featurerequest"
	idpjustaskingcontroller "github.com/chande/justasking/api/controller/idpjustasking"
	priceplancontroller "github.com/chande/justasking/api/controller/priceplan"
	smscontroller "github.com/chande/justasking/api/controller/sms"
	stripecontroller "github.com/chande/justasking/api/controller/stripe"
	supportcontroller "github.com/chande/justasking/api/controller/support"
	themecontroller "github.com/chande/justasking/api/controller/theme"
	tokencontroller "github.com/chande/justasking/api/controller/token"
	twiliocontroller "github.com/chande/justasking/api/controller/twilio"
	usercontroller "github.com/chande/justasking/api/controller/user"
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
