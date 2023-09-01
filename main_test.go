package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"stock_market_simulation/m/controllers"
	"stock_market_simulation/m/initializers"
	"stock_market_simulation/m/models"
	"strings"
	"testing"
)

func init() {
	initializers.LoadEnviromentalVariables()
	initializers.ConnectToTestDatabase()
}

func TestUserRegister(t *testing.T) {
	// Setup router
	router := gin.Default()
	router.POST("/users", controllers.ResgisterUser)
	// Create a new HTTP request
	payload := `{"username": "testuser", "password": "testpass"}`
	req := httptest.NewRequest("POST", "/users", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	res := w.Body
	body, _ := ioutil.ReadAll(res)
	bodyString := string(body)
	fmt.Println(bodyString)
	var result struct {
		Error string `json:"error"`
	}
	if err := json.Unmarshal([]byte(bodyString), &result); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	if result.Error != "" {
		t.Log("Exception: ", result.Error)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	} else {
		assert.Equal(t, http.StatusOK, w.Code)
	}
}

func TestLogin(t *testing.T) {
	router := gin.Default()
	router.POST("/login", controllers.LoginUser)
	payload := `{"username": "testusr", "password": "testpass"}`
	req := httptest.NewRequest("POST", "/login", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	res := w.Body
	body, _ := ioutil.ReadAll(res)
	bodyString := string(body)
	fmt.Println(bodyString)
	var result struct {
		Error string `json:"error"`
	}
	if err := json.Unmarshal([]byte(bodyString), &result); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	if result.Error != "" {
		t.Log("Exception: ", result.Error)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	} else {
		assert.Equal(t, http.StatusOK, w.Code)
	}
}

func TestGetAllUsers(t *testing.T) {
	router := gin.Default()
	router.GET("/allUsers", controllers.GetAllUsers)

	req := httptest.NewRequest("GET", "/allUsers", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	res := w.Body
	body, _ := ioutil.ReadAll(res)
	bodyString := string(body)
	fmt.Println(bodyString)
	var result struct {
		Error string `json:"error"`
	}
	if err := json.Unmarshal([]byte(bodyString), &result); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	if result.Error != "" {
		t.Log("Exception: ", result.Error)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	} else {
		assert.Equal(t, http.StatusOK, w.Code)
	}
}

func TestGetOneUser(t *testing.T) {
	router := gin.Default()
	router.GET("/username/:username", controllers.GetUser)
	req := httptest.NewRequest("GET", "/username/testuser", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	res := w.Body
	body, _ := ioutil.ReadAll(res)
	bodyString := string(body)
	fmt.Println(bodyString)
	var result struct {
		Error string `json:"error"`
	}
	if err := json.Unmarshal([]byte(bodyString), &result); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	if result.Error != "" {
		t.Log("Exception: ", result.Error)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	} else {
		assert.Equal(t, http.StatusOK, w.Code)
	}
}

func TestAddStocks(t *testing.T) {
	router := gin.Default()
	router.POST("/stocks", controllers.AddStocks)
	payload := `{
    "ticker":"GOOGL",
    "openprice":10,
    "closeprice":20,
    "high":50,
    "low":20,
    "volume":800
}`
	req := httptest.NewRequest("POST", "/stocks", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	res := w.Body
	body, _ := ioutil.ReadAll(res)
	bodyString := string(body)
	fmt.Println(bodyString)
	var result struct {
		Error string `json:"error"`
	}
	if err := json.Unmarshal([]byte(bodyString), &result); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	if result.Error != "" {
		t.Log("Exception: ", result.Error)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	} else {
		assert.Equal(t, http.StatusOK, w.Code)
	}

}

func TestGetAllStocks(t *testing.T) {
	router := gin.Default()
	router.GET("/stocks", controllers.GetAllStocks)
	req := httptest.NewRequest("GET", "/stocks", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	res := w.Body
	body, _ := ioutil.ReadAll(res)
	bodyString := string(body)
	fmt.Println(bodyString)
	var result struct {
		Error string `json:"error"`
	}
	if err := json.Unmarshal([]byte(bodyString), &result); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	if result.Error != "" {
		t.Log("Exception: ", result.Error)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	} else {
		assert.Equal(t, http.StatusOK, w.Code)
	}

}

func TestGetOneStock(t *testing.T) {
	router := gin.Default()
	router.GET("/stocks/:ticker", controllers.GetOneStock)
	req := httptest.NewRequest("GET", "/stocks/GOGL", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	res := w.Body
	body, _ := ioutil.ReadAll(res)
	bodyString := string(body)
	fmt.Println(bodyString)
	var result struct {
		Error string `json:"error"`
	}
	if err := json.Unmarshal([]byte(bodyString), &result); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	if result.Error != "" {
		t.Log("Exception: ", result.Error)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	} else {
		assert.Equal(t, http.StatusOK, w.Code)
	}
}

func TestDoTransaction(t *testing.T) {
	router := gin.Default()
	taskCh := make(chan controllers.TransactionTask)
	go controllers.TransactionWorker(taskCh)
	router.POST("/transactions", func(c *gin.Context) {
		var transaction models.TransactionData
		if err := c.ShouldBindJSON(&transaction); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Send task to background worker
		taskCh <- controllers.TransactionTask{Context: c.Copy(), Transaction: transaction}

		c.JSON(http.StatusOK, gin.H{"message": "Transaction started"})
	})
	payload := `{
    "transactionid":"123123",
    "ticker":"GOOGL",
    "transactiontype":"buy",
    "transactionvolume":1,
    "userid":2
}`
	req := httptest.NewRequest("POST", "/transactions", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	res := w.Body
	body, _ := ioutil.ReadAll(res)
	bodyString := string(body)
	fmt.Println(bodyString)
	var result struct {
		Error string `json:"error"`
	}
	if err := json.Unmarshal([]byte(bodyString), &result); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	if result.Error != "" {
		t.Log("Exception: ", result.Error)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	} else {
		assert.Equal(t, http.StatusOK, w.Code)
	}
}

func TestGetAllTransactionOfUSer(t *testing.T) {
	router := gin.Default()
	router.GET("/transactions/:user_id", controllers.GetAllTransactionOFUser)
	req := httptest.NewRequest("GET", "/transactions/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	res := w.Body
	body, _ := ioutil.ReadAll(res)
	bodyString := string(body)
	fmt.Println(bodyString)
	var result struct {
		Error string `json:"error"`
	}
	if err := json.Unmarshal([]byte(bodyString), &result); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	if result.Error != "" {
		t.Log("Exception: ", result.Error)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	} else {
		assert.Equal(t, http.StatusOK, w.Code)
	}
}

func TestGetTransactionDataBetweenTime(t *testing.T) {
	router := gin.Default()
	router.GET("/transactions/:user_id/:start_time/:end_time", controllers.GetTransactionDataBetweenTime)
	req := httptest.NewRequest("GET", "/transactions/2/2023-08-28/2023-08-29", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	res := w.Body
	body, _ := ioutil.ReadAll(res)
	bodyString := string(body)
	fmt.Println(bodyString)
	var result struct {
		Error string `json:"error"`
	}
	if err := json.Unmarshal([]byte(bodyString), &result); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	if result.Error != "" {
		t.Log("Exception: ", result.Error)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	} else {
		assert.Equal(t, http.StatusOK, w.Code)
	}

}
