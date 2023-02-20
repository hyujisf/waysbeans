package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"waysbeans/dto"
	"waysbeans/models"
	"waysbeans/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"gopkg.in/gomail.v2"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(transactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{transactionRepository}
}

// mengambil seluruh data transaksi
func (h *handlerTransaction) FindTransactions(c echo.Context) error {
	transactions, err := h.TransactionRepository.FindTransactions()
	if err != nil {
		res := dto.ErrorResult{Status: "error", Message: err.Error()}
		return c.JSON(http.StatusNotFound, res)
	}

	res := dto.SuccessResult{Status: "success", Data: convertMultipleTransactionResponse(transactions)}
	return c.JSON(http.StatusOK, res)
}

// mengambil seluruh data transaksi milik user tertentu
func (h *handlerTransaction) FindTransactionsByUser(c echo.Context) error {
	userInfo := c.Get("userInfo").(map[string]interface{})
	idUser := int(userInfo["id"].(float64))

	transactions, err := h.TransactionRepository.FindTransactionsByUser(idUser)
	if err != nil {
		res := dto.ErrorResult{Status: "error", Message: err.Error()}
		return c.JSON(http.StatusNotFound, res)
	}

	res := dto.SuccessResult{Status: "success", Data: convertMultipleTransactionResponse(transactions)}
	return c.JSON(http.StatusOK, res)
}

// mengambil 1 data transaksi
func (h *handlerTransaction) GetDetailTransaction(c echo.Context) error {
	id := c.Param("id")

	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		res := dto.ErrorResult{Status: "error", Message: err.Error()}
		return c.JSON(http.StatusNotFound, res)
	}

	res := dto.SuccessResult{Status: "success", Data: convertTransactionResponse(transaction)}
	return c.JSON(http.StatusOK, res)
}

// membuat transaksi baru
func (h *handlerTransaction) CreateTransaction(c echo.Context) error {
	var request dto.CreateTransactionRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	userInfo := c.Get("userInfo").(jwt.MapClaims)
	request.UserID = int(userInfo["id"].(float64))

	validation := validator.New()
	if err := validation.Struct(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	newTransaction := models.Transaction{
		ID:        fmt.Sprintf("TRX-%d-%d", request.UserID, timeIn("Asia/Jakarta").UnixNano()),
		OrderDate: timeIn("Asia/Jakarta"),
		Total:     request.Total,
		Status:    "new",
		UserID:    request.UserID,
	}

	for _, order := range request.Products {
		newTransaction.Order = append(newTransaction.Order, models.OrderResponseForTransaction{
			ID:        order.ID,
			ProductID: order.ProductID,
			OrderQty:  order.OrderQty,
		})
	}

	transactionAdded, err := h.TransactionRepository.CreateTransaction(newTransaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	transaction, err := h.TransactionRepository.GetTransaction(transactionAdded.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	s := snap.Client{}
	s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transaction.ID,
			GrossAmt: int64(transaction.Total),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: transaction.User.Name,
			Phone: transaction.User.Phone,
			BillAddr: &midtrans.CustomerAddress{
				FName:    transaction.User.Name,
				Phone:    transaction.User.Phone,
				Address:  transaction.User.Address,
				Postcode: transaction.User.PostCode,
			},
			ShipAddr: &midtrans.CustomerAddress{
				FName:    transaction.User.Name,
				Phone:    transaction.User.Phone,
				Address:  transaction.User.Address,
				Postcode: transaction.User.PostCode,
			},
		},
	}

	snapResp, _ := s.CreateTransactionToken(req)

	transaction, err = h.TransactionRepository.UpdateTokenTransaction(snapResp, transaction.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, dto.SuccessResult{Status: "success", Data: convertTransactionResponse(transaction)})
}
func (h *handlerTransaction) UpdateTransactionStatus(c echo.Context) error {
	id := c.Param("id")
	var request dto.UpdateTransactionRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	fmt.Println("ID : ", id)
	fmt.Println("Status : ", request.Status)

	// memeriksa transaksi yang ingin diupdate
	_, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	// mengupdate status transaksi
	transaction, err := h.TransactionRepository.UpdateTransaction(request.Status, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	// mengambil data transaksi yang sudah diupdate
	transaction, err = h.TransactionRepository.GetTransaction(transaction.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	var emailType string
	switch request.Status {
	case "rejected":
		emailType = "Rejected"
	case "sent":
		emailType = "Success, Product On Delivery"
	case "done":
		emailType = "Success, Product Received"
	default:
		emailType = "Undefined"
	}

	go SendTransactionMail(emailType, transaction)

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: convertTransactionResponse(transaction)})
}

func (h *handlerTransaction) Notification(c echo.Context) error {
	// Initialize empty map for notification payload
	var notificationPayload map[string]interface{}

	// Parse JSON request body and set to payload
	err := c.Bind(&notificationPayload)
	if err != nil {
		// Handle error when decoding payload
		response := dto.ErrorResult{Status: "error", Message: err.Error()}
		return c.JSON(http.StatusBadRequest, response)
	}

	// Get order ID from payload
	orderID, exists := notificationPayload["order_id"].(string)
	if !exists {
		// Handle case when order ID key not found
		return c.NoContent(http.StatusBadRequest)
	}

	// Get transaction from database using order ID
	transaction, err := h.TransactionRepository.GetTransaction(orderID)
	if err != nil {
		// Handle case when transaction not found in database
		fmt.Println("Transaction not found")
		return c.NoContent(http.StatusOK)
	}

	transactionStatus, exists := notificationPayload["transaction_status"].(string)
	if !exists {
		// Handle case when transaction status key not found
		return c.NoContent(http.StatusBadRequest)
	}

	fraudStatus, _ := notificationPayload["fraud_status"].(string)

	// Set transaction status based on response from check transaction status
	switch transactionStatus {
	case "capture":
		switch fraudStatus {
		case "challenge":
			h.TransactionRepository.UpdateTransaction("pending", transaction.ID)
			// TODO: Set transaction status on your database to 'challenge'
			// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal'
		case "accept":
			h.TransactionRepository.UpdateTransaction("success", transaction.ID)
			SendTransactionMail("Success", transaction)
			// TODO: Set transaction status on your database to 'success'
		}
	case "settlement":
		h.TransactionRepository.UpdateTransaction("success", transaction.ID)
		SendTransactionMail("Success", transaction)
		// TODO: Set transaction status on your database to 'success'
	case "deny":
		h.TransactionRepository.UpdateTransaction("failed", transaction.ID)
		// TODO: You can ignore 'deny', because most of the time it allows payment retries
		// and later can become success
	case "cancel", "expire":
		h.TransactionRepository.UpdateTransaction("failed", transaction.ID)
		SendTransactionMail("Failed", transaction)
		// TODO: Set transaction status on your database to 'failure'
	case "pending":
		h.TransactionRepository.UpdateTransaction("pending", transaction.ID)
		// TODO: Set transaction status on your database to 'pending' / waiting payment
	}

	return c.NoContent(http.StatusOK)
}

func convertMultipleTransactionResponse(transactions []models.Transaction) []dto.TransactionResponse {
	var transactionsResponse []dto.TransactionResponse

	for _, trx := range transactions {
		var trxResponse = dto.TransactionResponse{
			ID:         trx.ID,
			MidtransID: trx.MidtransID,
			OrderDate:  trx.OrderDate.Format("Monday, 2 January 2006"),
			Total:      trx.Total,
			Status:     trx.Status,
			User:       trx.User,
		}

		for _, order := range trx.Order {
			trxResponse.Products = append(trxResponse.Products, dto.ProductResponseForTransaction{
				ID:          order.Product.ID,
				Name:        order.Product.Name,
				Price:       order.Product.Price,
				Description: order.Product.Description,
				Image:       order.Product.Image,
				OrderQty:    order.OrderQty,
			})
		}

		transactionsResponse = append(transactionsResponse, trxResponse)
	}

	return transactionsResponse
}

func convertTransactionResponse(transaction models.Transaction) dto.TransactionResponse {
	var transactionResponse = dto.TransactionResponse{
		ID:         transaction.ID,
		MidtransID: transaction.MidtransID,
		OrderDate:  transaction.OrderDate.Format("Monday, 2 January 2006"),
		Total:      transaction.Total,
		Status:     transaction.Status,
		User:       transaction.User,
	}

	for _, order := range transaction.Order {
		transactionResponse.Products = append(transactionResponse.Products, dto.ProductResponseForTransaction{
			ID:          order.Product.ID,
			Name:        order.Product.Name,
			Price:       order.Product.Price,
			Description: order.Product.Description,
			Image:       order.Product.Image,
			OrderQty:    order.OrderQty,
		})
	}

	return transactionResponse
}

// fungsi untuk kirim email transaksi
func SendTransactionMail(status string, transaction models.Transaction) {

	var CONFIG_SMTP_HOST = os.Getenv("CONFIG_SMTP_HOST")
	var CONFIG_SMTP_PORT, _ = strconv.Atoi(os.Getenv("CONFIG_SMTP_PORT"))
	var CONFIG_SENDER_NAME = os.Getenv("CONFIG_SENDER_NAME")
	var CONFIG_AUTH_EMAIL = os.Getenv("CONFIG_AUTH_EMAIL")
	var CONFIG_AUTH_PASSWORD = os.Getenv("CONFIG_AUTH_PASSWORD")

	var products []map[string]interface{}

	for _, order := range transaction.Order {
		products = append(products, map[string]interface{}{
			"ProductName": order.Product.Name,
			"Price":       order.Product.Price,
			"Qty":         order.OrderQty,
			"SubTotal":    float64(order.OrderQty * int(order.Product.Price)),
		})
	}

	data := map[string]interface{}{
		"TransactionID":     transaction.ID,
		"TransactionStatus": status,
		"UserName":          transaction.User.Name,
		"OrderDate":         timeIn("Asia/Jakarta").Format("Monday, 2 January 2006"),
		"Total":             float64(transaction.Total),
		"Products":          products,
	}

	// mengambil file template
	t, err := template.ParseFiles("view/notification_email.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bodyMail := new(bytes.Buffer)

	// mengeksekusi template, dan memparse "data" ke template
	t.Execute(bodyMail, data)

	// create new message
	trxMail := gomail.NewMessage()
	trxMail.SetHeader("From", CONFIG_SENDER_NAME)
	trxMail.SetHeader("To", transaction.User.Email)
	trxMail.SetHeader("Subject", "WAYSBEANS ORDER NOTIFICATION")
	trxMail.SetBody("text/html", bodyMail.String())

	trxDialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	err = trxDialer.DialAndSend(trxMail)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Pesan terkirim!")
}

// fungsi untuk mendapatkan waktu sesuai zona indonesia
func timeIn(name string) time.Time {
	loc, err := time.LoadLocation(name)
	if err != nil {
		panic(err)
	}
	return time.Now().In(loc)
}
