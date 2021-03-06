package main

import (
	"context"

	giconfig "github.com/b2wdigital/goignite/v2/config"
	gilog "github.com/b2wdigital/goignite/v2/log"
	gilogrus "github.com/b2wdigital/goignite/v2/log/logrus/v1"
	giresty "github.com/b2wdigital/goignite/v2/resty/v2"
	health "github.com/b2wdigital/goignite/v2/resty/v2/ext/health"
)

func main() {

	var err error

	giconfig.Load()

	ctx := context.Background()

	gilogrus.NewLogger()

	logger := gilog.FromContext(ctx)

	// call google

	googleopt := new(giresty.Options)

	err = giconfig.UnmarshalWithPath("app.client.resty.google", googleopt)
	if err != nil {
		logger.Errorf(err.Error())
	}

	healthIntegrator := health.NewDefaultIntegrator()

	cligoogle := giresty.NewClient(ctx, googleopt, healthIntegrator.Register)
	reqgoogle := cligoogle.R()

	respgoogle, err := reqgoogle.Get("/")
	if err != nil {
		logger.Fatalf(err.Error())
	}

	logger.Infof(respgoogle.String())

	// call acom

	acomopt := new(giresty.Options)

	err = giconfig.UnmarshalWithPath("app.client.resty.acom", acomopt)
	if err != nil {
		logger.Errorf(err.Error())
	}

	cliacom := giresty.NewClient(ctx, acomopt)
	reqacom := cliacom.R()

	respacom, err := reqacom.Get("/")
	if err != nil {
		logger.Fatalf(err.Error())
	}

	logger.Infof(respacom.String())
}
