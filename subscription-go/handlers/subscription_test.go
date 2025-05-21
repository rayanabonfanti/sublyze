package handlers

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"subscription/database"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// MockDatabase cria um mock do banco de dados para testes
func MockDatabase(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	require.NoError(t, err)

	return gormDB, mock
}

func TestCreateSubscription(t *testing.T) {
	// Setup do mock do banco de dados
	gormDB, mock := MockDatabase(t)
	originalDB := database.DB
	database.DB = gormDB
	defer func() { database.DB = originalDB }()

	tests := []struct {
		name         string
		requestBody  string
		mockExpect   func()
		expectedCode int
		expectedBody string
	}{
		{
			name:        "successful creation",
			requestBody: `{"plan":"premium","user_id":"123","status":"active"}`,
			mockExpect: func() {
				mock.ExpectBegin()
				mock.ExpectQuery(`INSERT INTO "subscriptions"`).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
				mock.ExpectCommit()
			},
			expectedCode: http.StatusCreated,
			expectedBody: `"plan":"premium"`,
		},
		{
			name:         "invalid data",
			requestBody:  `{"invalid":"data"`,
			mockExpect:   func() {},
			expectedCode: http.StatusBadRequest,
			expectedBody: `"error":"Dados inválidos"`,
		},
		{
			name:        "database error",
			requestBody: `{"plan":"premium","user_id":"123","status":"active"}`,
			mockExpect: func() {
				mock.ExpectBegin()
				mock.ExpectQuery(`INSERT INTO "subscriptions"`).
					WillReturnError(sql.ErrConnDone)
				mock.ExpectRollback()
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: `"error":"Falha ao criar assinatura"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Configura o mock
			tt.mockExpect()

			// Cria um router Gin para teste
			router := gin.Default()
			router.POST("/subscriptions", CreateSubscription)

			// Cria uma requisição de teste
			req, err := http.NewRequest("POST", "/subscriptions", strings.NewReader(tt.requestBody))
			require.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			// Grava a resposta
			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			// Verifica os resultados
			assert.Equal(t, tt.expectedCode, rr.Code)
			assert.Contains(t, rr.Body.String(), tt.expectedBody)

			// Verifica se todas as expectativas do mock foram atendidas
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
