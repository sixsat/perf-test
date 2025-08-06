package gorestyv2v3

import (
	"fmt"
	"io"
	"math/rand/v2"
	"testing"

	restyV2 "github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
	"github.com/labstack/echo/v4"
	restyV3 "resty.dev/v3"
)

func BenchmarkRestyClientV2MarshalSmall(b *testing.B) {
	client := restyV2.
		New().
		SetBaseURL(startEchoServer(b, 0)).
		SetJSONMarshaler(json.Marshal).
		SetDebug(false).
		DisableTrace()
	defer client.GetClient().CloseIdleConnections()
	body := genBody(10)

	for b.Loop() {
		_, err := client.R().SetBody(body).Post("/post")
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
	}
}

func BenchmarkRestyClientV3MarshalSmall(b *testing.B) {
	client := restyV3.
		New().
		SetBaseURL(startEchoServer(b, 0)).
		AddContentTypeEncoder("application/json", func(w io.Writer, v any) error {
			return json.NewEncoder(w).Encode(v)
		}).
		SetDebug(false).
		DisableTrace()
	defer client.Close()
	body := genBody(10)

	for b.Loop() {
		_, err := client.R().SetBody(body).Post("/post")
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
	}
}

func BenchmarkRestyClientV2MarshalMedium(b *testing.B) {
	client := restyV2.
		New().
		SetBaseURL(startEchoServer(b, 0)).
		SetJSONMarshaler(json.Marshal).
		SetDebug(false).
		DisableTrace()
	defer client.GetClient().CloseIdleConnections()
	body := genBody(100)

	for b.Loop() {
		_, err := client.R().SetBody(body).Post("/post")
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
	}
}

func BenchmarkRestyClientV3MarshalMedium(b *testing.B) {
	client := restyV3.
		New().
		SetBaseURL(startEchoServer(b, 0)).
		AddContentTypeEncoder("application/json", func(w io.Writer, v any) error {
			return json.NewEncoder(w).Encode(v)
		}).
		SetDebug(false).
		DisableTrace()
	defer client.Close()
	body := genBody(100)

	for b.Loop() {
		_, err := client.R().SetBody(body).Post("/post")
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
	}
}

func BenchmarkRestyClientV2MarshalLarge(b *testing.B) {
	client := restyV2.
		New().
		SetBaseURL(startEchoServer(b, 0)).
		SetJSONMarshaler(json.Marshal).
		SetDebug(false).
		DisableTrace()
	defer client.GetClient().CloseIdleConnections()
	body := genBody(1000)

	for b.Loop() {
		_, err := client.R().SetBody(body).Post("/post")
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
	}
}

func BenchmarkRestyClientV3MarshalLarge(b *testing.B) {
	client := restyV3.
		New().
		SetBaseURL(startEchoServer(b, 0)).
		AddContentTypeEncoder("application/json", func(w io.Writer, v any) error {
			return json.NewEncoder(w).Encode(v)
		}).
		SetDebug(false).
		DisableTrace()
	defer client.Close()
	body := genBody(1000)

	for b.Loop() {
		_, err := client.R().SetBody(body).Post("/post")
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
	}
}

func BenchmarkRestyClientV2UnmarshalSmall(b *testing.B) {
	client := restyV2.
		New().
		SetBaseURL(startEchoServer(b, 10)).
		SetJSONUnmarshaler(json.Unmarshal).
		SetDebug(false).
		DisableTrace()
	defer client.GetClient().CloseIdleConnections()

	for b.Loop() {
		res := result{Users: make([]body, 10)}
		_, err := client.R().SetResult(&res).Get("/get")
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
	}
}

func BenchmarkRestyClientV3UnmarshalSmall(b *testing.B) {
	client := restyV3.
		New().
		SetBaseURL(startEchoServer(b, 10)).
		AddContentTypeDecoder("application/json", func(r io.Reader, v any) error {
			return json.NewDecoder(r).Decode(v)
		}).
		SetDebug(false).
		DisableTrace()
	defer client.Close()

	for b.Loop() {
		res := result{Users: make([]body, 10)}
		_, err := client.R().SetResult(&res).Get("/get")
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
	}
}

func BenchmarkRestyClientV2UnmarshalMedium(b *testing.B) {
	client := restyV2.
		New().
		SetBaseURL(startEchoServer(b, 100)).
		SetJSONUnmarshaler(json.Unmarshal).
		SetDebug(false).
		DisableTrace()
	defer client.GetClient().CloseIdleConnections()

	for b.Loop() {
		res := result{Users: make([]body, 100)}
		_, err := client.R().SetResult(&res).Get("/get")
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
	}
}

func BenchmarkRestyClientV3UnmarshalMedium(b *testing.B) {
	client := restyV3.
		New().
		SetBaseURL(startEchoServer(b, 100)).
		AddContentTypeDecoder("application/json", func(r io.Reader, v any) error {
			return json.NewDecoder(r).Decode(v)
		}).
		SetDebug(false).
		DisableTrace()
	defer client.Close()

	for b.Loop() {
		res := result{Users: make([]body, 100)}
		_, err := client.R().SetResult(&res).Get("/get")
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
	}
}

func BenchmarkRestyClientV2UnmarshalLarge(b *testing.B) {
	client := restyV2.
		New().
		SetBaseURL(startEchoServer(b, 1000)).
		SetJSONUnmarshaler(json.Unmarshal).
		SetDebug(false).
		DisableTrace()
	defer client.GetClient().CloseIdleConnections()

	for b.Loop() {
		res := result{Users: make([]body, 1000)}
		_, err := client.R().SetResult(&res).Get("/get")
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
	}
}

func BenchmarkRestyClientV3UnmarshalLarge(b *testing.B) {
	client := restyV3.
		New().
		SetBaseURL(startEchoServer(b, 1000)).
		AddContentTypeDecoder("application/json", func(r io.Reader, v any) error {
			return json.NewDecoder(r).Decode(v)
		}).
		SetDebug(false).
		DisableTrace()
	defer client.Close()

	for b.Loop() {
		res := result{Users: make([]body, 1000)}
		_, err := client.R().SetResult(&res).Get("/get")
		if err != nil {
			b.Error(err)
			b.FailNow()
		}
	}
}

func startEchoServer(b *testing.B, numElements int) string {
	b.Helper()
	e := echo.New()
	e.GET("/get", func(c echo.Context) error {
		return c.JSON(200, echo.Map{
			"users": genBody(numElements),
		})
	})
	e.POST("/post", func(c echo.Context) error {
		return c.String(200, "")
	})
	e.HideBanner = true
	e.HidePort = true

	go func() {
		e.Start(":3000")
	}()
	b.Cleanup(func() { e.Close() })

	return "http://localhost:3000"
}

type result struct {
	Users []body `json:"users"`
}

type body struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	Age       int     `json:"age"`
	MobileNo  string  `json:"mobile_no"`
	Height    float64 `json:"height"`
	Weight    float64 `json:"weight"`
	Address   address `json:"address"`
	IsMarried bool    `json:"is_married"`
}

type address struct {
	AddressNo   string `json:"address_no"`
	VillageNo   string `json:"village_no"`
	Alley       string `json:"alley"`
	Street      string `json:"street"`
	District    string `json:"district"`
	SubDistrict string `json:"sub_district"`
	Province    string `json:"province"`
	Country     string `json:"country"`
	ZipCode     string `json:"zip_code"`
	PhoneNo     string `json:"phone_no"`
}

func genBody(numElements int) []body {
	bodies := make([]body, numElements)
	for i := range bodies {
		bodies[i] = body{
			FirstName: fmt.Sprintf("FirstName%d", i),
			LastName:  fmt.Sprintf("LastName%d", i),
			Email:     fmt.Sprintf("email%d@example.com", i),
			Age:       20 + rand.IntN(30),
			MobileNo:  fmt.Sprintf("MobileNo%d", i),
			Height:    155.0 + rand.Float64()*3,
			Weight:    50.0 + rand.Float64()*3,
			Address: address{
				AddressNo:   fmt.Sprintf("AddressNo%d", i),
				VillageNo:   fmt.Sprintf("VillageNo%d", i),
				Alley:       fmt.Sprintf("Alley%d", i),
				Street:      fmt.Sprintf("Street%d", i),
				District:    fmt.Sprintf("District%d", i),
				SubDistrict: fmt.Sprintf("SubDistrict%d", i),
				Province:    fmt.Sprintf("Province%d", i),
				Country:     fmt.Sprintf("Country%d", i),
				ZipCode:     fmt.Sprintf("ZipCode%d", i),
				PhoneNo:     fmt.Sprintf("PhoneNo%d", i),
			},
			IsMarried: i&1 == 0,
		}
	}
	return bodies
}
