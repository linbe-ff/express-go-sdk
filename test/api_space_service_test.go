package test

import (
	"context"
	"fmt"
	"github.com/linbe-ff/express-go-sdk"
	"testing"
)

func TestApiSpaceService(t *testing.T) {
	asCient := express.NewAPISpaceService("xxxxxxxxx")
	routes, err := asCient.SearchRoutes(context.Background(), &express.APISpaceReq{
		CpCode: "SF",
		MailNo: "SFxxxxxxxxxxxxxx",
		Tel:    "0714",
		Key:    "xxxx",
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(routes)

}

func TestApiSpaceServiceDiscern(t *testing.T) {
	asCient := express.NewAPISpaceService("xxxxxxxxxxxxxxxxxxxxx")
	routes, err := asCient.MailDiscern(context.Background(), "dfSF316322510125", "key")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(routes)

}
