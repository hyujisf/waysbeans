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

	"github.com/go-mail/mail"
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

	res := dto.SuccessResult{Status: "success", Data: transactions}
	return c.JSON(http.StatusOK, res)
}

// mengambil 1 data transaksi
func (h *handlerTransaction) GetTransaction(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		res := dto.ErrorResult{Status: "error", Message: err.Error()}
		return c.JSON(http.StatusNotFound, res)
	}

	res := dto.SuccessResult{Status: "success", Data: transaction}
	return c.JSON(http.StatusOK, res)
}

// mengambil seluruh data transaksi milik user tertentu
func (h *handlerTransaction) GetUserTransaction(c echo.Context) error {
	userLogin := c.Get("userLogin").(jwt.MapClaims)
	id := int(userLogin["id"].(float64))

	transaction, err := h.TransactionRepository.GetUserTransaction(id)
	if err != nil {
		res := dto.ErrorResult{Status: "error", Message: err.Error()}
		return c.JSON(http.StatusNotFound, res)
	}

	res := dto.SuccessResult{Status: "success", Data: transaction}
	return c.JSON(http.StatusOK, res)
}
func (h *handlerTransaction) CreateTransaction(c echo.Context) error {
	userLogin := c.Get("userLogin").(jwt.MapClaims)
	id := int(userLogin["id"].(float64))

	request := new(dto.CreateTransaction)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	order, err := h.TransactionRepository.FindOrdersTransactions(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	var transactionIsMatch = false
	var transactionId int
	for !transactionIsMatch {
		transactionId = int(time.Now().Unix())
		transactionData, _ := h.TransactionRepository.GetTransaction(transactionId)
		if transactionData.ID == 0 {
			transactionIsMatch = true
		}
	}

	transaction := models.Transaction{
		ID:     int64(transactionId),
		UserID: id,
		Total:  request.Total,
		Status: "pending",
		Order:  order,
	}

	data, err := h.TransactionRepository.CreateTransaction(transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	dataTransactions, err := h.TransactionRepository.GetTransactions(data.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Request payment token from midtrans ...
	// 1. Initiate Snap client
	var s = snap.Client{}
	s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)
	// Use to midtrans.Production if you want Production Environment (accept real transaction).

	// 2. Initiate Snap request param
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(int(dataTransactions.ID)),
			GrossAmt: int64(dataTransactions.Total),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: dataTransactions.User.Name,
			Email: dataTransactions.User.Email,
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, _ := s.CreateTransaction(req)

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: snapResp})
}
func (h *handlerTransaction) UpdateTransaction(c echo.Context) error {
	request := new(dto.UpdateTransaction)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))
	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	if request.UserID != 0 {
		transaction.UserID = request.UserID
	}

	if request.Total != 0 {
		transaction.Total = request.Total
	}

	if request.Status != "" {
		transaction.Status = request.Status
	}

	data, err := h.TransactionRepository.UpdateTransaction(transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: data})
}
func (h *handlerTransaction) DeleteTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transaction, err := h.TransactionRepository.GetTransaction(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	data, err := h.TransactionRepository.DeleteTransaction(transaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Status: "success", Data: data})
}
func (h *handlerTransaction) Notification(c echo.Context) error {
	var notificationPayload map[string]interface{}

	err := c.Bind(&notificationPayload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Status: "error", Message: err.Error()})
	}

	transactionStatus := notificationPayload["transaction_status"].(string)
	fraudStatus := notificationPayload["fraud_status"].(string)
	orderId := notificationPayload["order_id"].(string)

	transaction, _ := h.TransactionRepository.GetOneTransaction(orderId)

	switch transactionStatus {
	case "capture":
		switch fraudStatus {
		case "challenge":
			// TODO set transaction status on your database to 'challenge'
			// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
			h.TransactionRepository.UpdateTransactions("pending", orderId)
		case "accept":
			// TODO set transaction status on your database to 'success'
			SendMail("success", transaction)
			h.TransactionRepository.UpdateTransactions("success", orderId)
		}
	case "settlement":
		// TODO set transaction status on your databaase to 'success'
		SendMail("success", transaction)
		h.TransactionRepository.UpdateTransactions("success", orderId)
	case "deny":
		// TODO you can ignore 'deny', because most of the time it allows payment retries
		// and later can become success
		SendMail("failed", transaction)
		h.TransactionRepository.UpdateTransactions("failed", orderId)
	case "cancel", "expire":
		// TODO set transaction status on your databaase to 'failure'
		SendMail("failed", transaction)
		h.TransactionRepository.UpdateTransactions("failed", orderId)
	case "pending":
		// TODO set transaction status on your databaase to 'pending' / waiting payment
		h.TransactionRepository.UpdateTransactions("pending", orderId)
	}

	return c.NoContent(http.StatusOK)
}

// fungsi untuk kirim email transaksi

func SendMail(status string, transaction models.Transaction) {

	if status != transaction.Status && (status == "success") {
		var CONFIG_SMTP_HOST = "smtp.gmail.com"
		var CONFIG_SMTP_PORT = 587
		var CONFIG_SENDER_NAME = "WaysBeans <adefitryana007@gmail.com>"
		var CONFIG_AUTH_EMAIL = os.Getenv("EMAIL_SYSTEM")
		var CONFIG_AUTH_PASSWORD = os.Getenv("PASSWORD_SYSTEM")

		var productName = "Thank For Purchasing Our Product"
		var price = strconv.Itoa(transaction.Total)

		dialer := mail.NewDialer(CONFIG_SMTP_HOST, CONFIG_SMTP_PORT, CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD)
		dialer.StartTLSPolicy = mail.MandatoryStartTLS

		mailer := mail.NewMessage()
		mailer.SetHeader("From", CONFIG_SENDER_NAME)
		mailer.SetHeader("To", transaction.User.Email)
		mailer.SetHeader("Subject", "Transaction Status")
		mailer.SetBody("text/html", fmt.Sprintf(`<!DOCTYPE html>
		  <html lang="en">
			<head>
			<meta charset="UTF-8" />
			<meta http-equiv="X-UA-Compatible" content="IE=edge" />
			<meta name="viewport" content="width=device-width, initial-scale=1.0" />
			<title>Document</title>
			<style>
			  h1 {
			  color: brown;
			  }
			</style>
			</head>
			<body>
			<h2>Product payment :</h2>
			<ul style="list-style-type:none;">
			  <li> %s</li>
			  <li>Total payment: Rp.%s</li>
			  <li>Status : <b>%s</b></li>
			</ul>
			</body>
		  </html>`, productName, price, status))

		err := dialer.DialAndSend(mailer)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Println("Mail sent! to " + transaction.User.Email)
	}
}
func SendTransactionMail(status string, transaction models.Transaction) {

	var CONFIG_SMTP_HOST = "smtp.gmail.com"
	var CONFIG_SMTP_PORT = 587
	var CONFIG_SENDER_NAME = "Waysbeans <canqalloid.me@gmail.com>"
	var CONFIG_AUTH_EMAIL = os.Getenv("EMAIL_SYSTEM")
	var CONFIG_AUTH_PASSWORD = os.Getenv("PASSWORD_SYSTEM")

	var products []map[string]interface{}

	for _, order := range transaction.Order {
		products = append(products, map[string]interface{}{
			"ProductName": order.Product.Name,
			"Price":       order.Product.Price,
			"Qty":         order.QTY,
			"SubTotal":    float64(order.QTY * int(order.Product.Price)),
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
