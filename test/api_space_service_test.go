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
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(routes)

}
