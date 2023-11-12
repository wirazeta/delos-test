package unittest

import (
	"bytes"
	"delos-test/controllers/farmcontroller"
	"delos-test/controllers/pondcontroller"
	"delos-test/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateFarmOK(t *testing.T) {
	models.ConnectDatabase()
	r := gin.Default()

	r.POST("/api/farms", farmcontroller.Create)

	jsonStr := []byte(`{"name": "Farm 1"}`)
	req, err := http.NewRequest("POST", "/api/farms", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetAllFarmOK(t *testing.T) {
	models.ConnectDatabase()
	r := gin.Default()
	r.GET("/api/farms", farmcontroller.Index)

	req, err := http.NewRequest("GET", "/api/farms", nil)

	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, response["farms"])
	assert.NoError(t, err)
}

func TestGetFarmByIdOK(t *testing.T) {
	models.ConnectDatabase()
	r := gin.Default()

	r.GET("/api/farms/:id", farmcontroller.GetById)

	req, err := http.NewRequest("GET", "/api/farms/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, err)
}

func TestGetFarmByIdNotFound(t *testing.T) {
	models.ConnectDatabase()
	r := gin.Default()

	r.GET("/api/farms/:id", farmcontroller.GetById)

	req, err := http.NewRequest("GET", "/api/farms/999", nil)

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCreateFarmBadRequest(t *testing.T) {
	models.ConnectDatabase()
	r := gin.Default()

	r.POST("/api/farms", farmcontroller.Create)

	jsonStr := []byte(`{"name": "Farm 1"}`)
	req, err := http.NewRequest("POST", "/api/farms", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateFarmOK(t *testing.T) {
	models.ConnectDatabase()
	r := gin.Default()

	r.PUT("/api/farms/:id", farmcontroller.Update)

	jsonStr := []byte(`{"name": "Farm 15"}`)
	req, err := http.NewRequest("PUT", "/api/farms/1", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateFarmButCreateObject(t *testing.T) {
	models.ConnectDatabase()
	r := gin.Default()

	r.PUT("/api/farms/:id", farmcontroller.Update)

	jsonStr := []byte(`{"name": "Farm 7"}`)
	req, err := http.NewRequest("PUT", "/api/farms/2", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteFarmOK(t *testing.T) {
	models.ConnectDatabase()
	r := gin.Default()

	r.DELETE("/api/farms/:id", farmcontroller.Delete)

	req, err := http.NewRequest("DELETE", "/api/farms/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteFarmBadRequest(t *testing.T) {
	models.ConnectDatabase()
	r := gin.Default()

	r.DELETE("/api/farms/:id", farmcontroller.Delete)

	req, err := http.NewRequest("DELETE", "/api/farms/100", nil)

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreatePondOK(t *testing.T) {
	models.ConnectDatabase()
	r := gin.Default()

	r.POST("/api/ponds", pondcontroller.Create)

	jsonStr := []byte(`{"name": "Pond 1", "FarmID": 2}`)
	req, err := http.NewRequest("POST", "/api/ponds", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetAllPondOK(t *testing.T) {
	models.ConnectDatabase()
	r := gin.Default()
	r.GET("/api/ponds", pondcontroller.Index)

	req, err := http.NewRequest("GET", "/api/ponds", nil)

	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, response["ponds"])
	assert.NoError(t, err)
}

func TestGetPondByIdOK(t *testing.T) {
	models.ConnectDatabase()
	r := gin.Default()

	r.GET("/api/ponds/:id", pondcontroller.GetById)

	req, err := http.NewRequest("GET", "/api/ponds/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, err)
}

func TestGetPondByIdNotFound(t *testing.T) {
	models.ConnectDatabase()
	r := gin.Default()

	r.GET("/api/ponds/:id", pondcontroller.GetById)

	req, err := http.NewRequest("GET", "/api/ponds/999", nil)

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestCreatePondBadRequest(t *testing.T) {
	models.ConnectDatabase()
	r := gin.Default()

	r.POST("/api/ponds", pondcontroller.Create)

	jsonStr := []byte(`{"name": "Pond 1", "FarmID": 2}`)
	req, err := http.NewRequest("POST", "/api/ponds", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdatePondOK(t *testing.T) {
	models.ConnectDatabase()
	r := gin.Default()

	r.PUT("/api/ponds/:id", pondcontroller.Update)

	jsonStr := []byte(`{"name": "Pond 15"}`)
	req, err := http.NewRequest("PUT", "/api/ponds/1", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdatePondButCreateObject(t *testing.T) {
	models.ConnectDatabase()
	r := gin.Default()

	r.PUT("/api/ponds/:id", pondcontroller.Update)

	jsonStr := []byte(`{"name": "Pond 7"}`)
	req, err := http.NewRequest("PUT", "/api/ponds/2", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeletePondOK(t *testing.T) {
	models.ConnectDatabase()
	r := gin.Default()

	r.DELETE("/api/ponds/:id", pondcontroller.Delete)

	req, err := http.NewRequest("DELETE", "/api/ponds/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeletePondBadRequest(t *testing.T) {
	models.ConnectDatabase()
	r := gin.Default()

	r.DELETE("/api/ponds/:id", pondcontroller.Delete)

	req, err := http.NewRequest("DELETE", "/api/ponds/100", nil)

	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
