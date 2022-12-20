package app

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"

	"service/internal/service"
)

type request struct {
	method  string
	path    string
	payload string
}

type want struct {
	status int
	body   string
}

func TestHandler(t *testing.T) {
	a := App{
		port:    "8080",
		router:  chi.NewRouter(),
		service: service.Mock{},
	}
	go a.Run()

	tests := map[string]struct {
		request request
		want    want
	}{
		"create_biba": {
			request{
				"POST",
				"/create",
				`{"name":"Biba","age":24}`,
			},
			want{201, `{"id":1}`},
		},
		"create_boba": {
			request{
				"POST",
				"/create",
				`{"name":"Boba","age":24}`,
			},
			want{201, `{"id":2}`},
		},
		"create500": {
			request{
				"POST",
				"/create",
				`{"name":"name500","age":"24"}`,
			},
			want{500, `json: cannot unmarshal string into Go struct field CreateRequest.age of type int`},
		},
		"makeFriends": {
			request{
				"POST",
				"/makeFriends",
				`{"source_id":1,"target_id":2}`,
			},
			want{200, `{"message":"Biba и Boba теперь друзья"}`},
		},
		"DeleteUser": {
			request{
				"DELETE",
				"/user",
				`{"user_id": 3}`,
			},
			want{200, `{"message":"Biba удален"}`},
		},
		"GetAllUsers": {
			request{
				"GET",
				"/getAll",
				"",
			},
			want{200, `{"message":"all users"}`},
		},
		"GetUserFriends": {
			request{
				"GET",
				"/friends/2",
				"",
			},
			want{200, `{"message":"Boba"}`},
		},
		"UpdateAge": {
			request{
				"PUT",
				"/user/3",
				`{"user_age" : 23}`,
			},
			want{200, `{"message":"возраст пользователя успешно обновлён"}`},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			client := &http.Client{}
			r := strings.NewReader(tc.request.payload)
			req, err := http.NewRequest(tc.request.method, "http://localhost:8080"+tc.request.path, r)
			if err != nil {
				t.Fatal(err)
			}
			res, err := client.Do(req)
			if err != nil {
				t.Fatal(err)
			}

			b, _ := io.ReadAll(res.Body)
			res.Body.Close()

			assert.Equal(t, tc.want.status, res.StatusCode)
			assert.Equal(t, tc.want.body, string(b))
		})
	}
}
