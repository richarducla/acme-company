package setuptesting

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func ServerTest(method string, endPoint string, payload string, params map[string]string, headers map[string]string, cookies map[string]string) (*echo.Echo, *http.Request, *httptest.ResponseRecorder, echo.Context) {
	server := echo.New()

	req := httptest.NewRequest(
		method,
		endPoint,
		strings.NewReader(payload),
	)

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("Content-Type", "application/json")

	for k, v := range cookies {
		cookie := new(http.Cookie)
		cookie.Name = k
		cookie.Value = v
		cookie.Expires = time.Now().Add(24 * time.Hour)
		req.AddCookie(cookie)
	}

	rec := httptest.NewRecorder()

	c := server.NewContext(req, rec)
	for k, v := range params {
		c.SetParamNames(k)
		c.SetParamValues(v)
	}

	return server, req, rec, c
}

func BuildBody(body interface{}) string {
	out, _ := json.Marshal(body)
	return string(out)
}
