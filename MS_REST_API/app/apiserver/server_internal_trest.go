package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/iavealokin/microservices/MS_REST_API/app/store/sqlstore/teststore"
	"github.com/stretchr/testify/assert"
)

//TestServerHandleUsersCreate ...
func TestServerHandleUsersCreate(t *testing.T) {
	s := newServer(teststore.New())
	testCases:=[]struct{
		name string
		payload interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"login":"evgen",
  				"username":"Evgeniy",
  				"surname":"Sheblyakin",
  				"birthday":"13.09.1794",
 				"password":"SanyaIsTheBest",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "invalid payload",
			payload: "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid params",
			payload: map[string]string{
				"email":"invalid",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}


	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T){
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/users",b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})

	}
}

