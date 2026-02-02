package handlers

import (
	"orch-go/internal/domain/customer"
	"orch-go/internal/domain/customer_address"
	"orch-go/internal/services"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/sync/errgroup"

	"orch-go/internal/transport/midleware"
)

type customerRest struct {
	Id        int32   `json:"id"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	Phone     *string `json:"phone"`
}

type customerAddressRest struct {
	CustomerId int32  `json:"customer_id"`
	Country    string `json:"country"`
	City       string `json:"city"`
	Street     string `json:"street"`
	ZipCode    string `json:"zip_code"`
	IsPrimary  bool   `json:"is_primary"`
}

func toRest(customer *customer.Customer) customerRest {
	restModel := customerRest{
		Id:        customer.Id,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Phone:     customer.Phone,
	}
	return restModel
}

func GetCustomerByIdHandler(s services.CustomerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		cust, err := s.GetCustomerById(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		restModel := toRest(cust)

		c.JSON(200, restModel)
		return
	}
}

func GetBySubstringHandler(s services.CustomerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		substring := c.Param("substr")
		if substring == "" {
			c.JSON(400, gin.H{"error": "substring cannot be empty"})
			return
		}
		customers, err := s.GetCustomersBySubstring(c.Request.Context(), customer.GetBySubStrRequest{
			SubStr:   substring,
			PageN:    0,
			PageSize: 0,
			Order:    "",
			Desk:     false,
		})
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		restModels := make([]customerRest, len(customers))
		for i, cus := range customers {
			restModels[i] = toRest(&cus)
		}
		c.JSON(200, customers)
		return
	}
}

func GetSelfHandler(s services.CustomerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, done := getIdFromRequest(c)
		if done {
			return
		}
		cust, err := s.GetCustomerById(c.Request.Context(), id)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		restModel := struct {
			Id        int32     `json:"id"`
			FirstName string    `json:"first_name"`
			LastName  string    `json:"last_name"`
			Email     string    `json:"email"`
			Phone     string    `json:"phone"`
			BirthDate time.Time `json:"birth_date"`
			CreatedAt time.Time `json:"created_at"`
		}{
			Id:        cust.Id,
			FirstName: cust.FirstName,
			LastName:  cust.LastName,
			Email:     cust.Email,
			Phone:     *cust.Phone,
			BirthDate: *cust.BirthDate,
			CreatedAt: *cust.CreatedAt,
		}
		c.JSON(200, restModel)
		return
	}
}

func getIdFromRequest(c *gin.Context) (int32, bool) {
	claims, ok := c.Get("claims")
	if !ok {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return 0, true
	}
	claimsMap, ok := claims.(jwt.MapClaims)
	if !ok {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return 0, true
	}

	id := claimsMap["id"].(int32)
	return id, false
}

func RegisterCustomerHandler(customerService services.CustomerService,
	credService services.UserCredentialService,
	addressService services.CustomerAddressService) gin.HandlerFunc {
	return func(c *gin.Context) {
		type RegisterCustomerModel struct {
			customerRest
			Login     string              `json:"login"`
			Password  string              `json:"password"`
			BirthDate time.Time           `json:"birth_date"`
			Address   customerAddressRest `json:"address"`
		}

		var model RegisterCustomerModel
		if err := c.ShouldBindJSON(&model); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		user, err := customerService.CreateCustomer(c.Request.Context(), &customer.Customer{
			Id:        0,
			FirstName: model.FirstName,
			LastName:  model.LastName,
			Email:     model.Email,
			Phone:     model.Phone,
			BirthDate: &model.BirthDate,
			CreatedAt: nil,
		})
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		g, gCtx := errgroup.WithContext(c.Request.Context())

		g.Go(func() error {
			_, err = addressService.CreateCustomerAddress(gCtx, &customer_address.CustomerAddress{
				Id:         nil,
				CustomerId: user.Id,
				Country:    model.Address.Country,
				City:       model.Address.City,
				Street:     model.Address.Street,
				ZipCode:    model.Address.ZipCode,
				IsPrimary:  model.Address.IsPrimary,
			})
			return err
		})

		g.Go(func() error {
			_, err = credService.CreateUserCredential(gCtx, user.Id, model.Login, model.Password)
			return err
		})

		if err := g.Wait(); err != nil {
			c.JSON(500, gin.H{"error": err.Error()}) // Use 500 for internal server errors from concurrent operations
			return
		}

		secretKey := os.Getenv("JWT_SECRET_KEY")
		authMiddleware := midleware.NewAuthMiddleware(secretKey)
		token, err := authMiddleware.GenerateToken(user.Id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"token": token})
		return

	}
}

func UpdateUserCredHandler(customerService services.CustomerService, credService services.UserCredentialService) gin.HandlerFunc {
	return func(c *gin.Context) {
		type UpdateCredsModel struct {
			Login    string `json:"login"`
			Password string `json:"password"`
		}
		var model UpdateCredsModel

		if err := c.ShouldBindJSON(&model); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if model.Login == "" || model.Password == "" {
			c.JSON(400, gin.H{"error": "login or password cannot be empty"})
		}
		id, done := getIdFromRequest(c)
		if done {
			return
		}

		cred, err := credService.GetUserCredentialById(c.Request.Context(), id)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		cred.Username = model.Login
		cred.PasswordHash = model.Password
		err = credService.UpdateUserCredential(c.Request.Context(), cred)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "credentials updated"})
		return
	}
}

func LoginHandler(credService services.UserCredentialService) gin.HandlerFunc {
	return func(c *gin.Context) {
		type LoginModel struct {
			Login    string `json:"login"`
			Password string `json:"password"`
		}

		var model LoginModel
		if err := c.ShouldBindJSON(&model); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		creds, err := credService.GetUserCredentialByUsername(c.Request.Context(), model.Login)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(creds.PasswordHash), []byte(model.Password))
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid password"})
			return
		}
		secretKey := os.Getenv("JWT_SECRET_KEY")
		authMiddleware := midleware.NewAuthMiddleware(secretKey)
		token, err := authMiddleware.GenerateToken(*creds.CustomerId)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, token)
		return
	}
}
func AddCustomerAddressHandler(s services.CustomerAddressService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var model customerAddressRest
		if err := c.ShouldBindJSON(&model); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		id, done := getIdFromRequest(c)
		if done {
			return
		}
		if id != model.CustomerId {
			c.JSON(400, gin.H{"error": "customer id does not match"})
			return
		}
		addressResult, err := s.CreateCustomerAddress(c.Request.Context(), &customer_address.CustomerAddress{
			Id:         nil,
			CustomerId: model.CustomerId,
			Country:    model.Country,
			City:       model.City,
			Street:     model.Street,
			ZipCode:    model.ZipCode,
			IsPrimary:  model.IsPrimary,
		})
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, addressResult)
		return
	}

}

func UpdateCustomerAddressHandler(s services.CustomerAddressService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var model struct {
			Id int32 `json:"id"`
			customerAddressRest
		}
		if err := c.ShouldBindJSON(&model); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		id, done := getIdFromRequest(c)
		if done {
			return
		}
		if id != model.CustomerId {
			c.JSON(400, gin.H{"error": "customer id does not match"})
			return
		}
		err := s.UpdateCustomerAddress(c.Request.Context(), &customer_address.CustomerAddress{
			Id:         &model.Id,
			CustomerId: model.CustomerId,
			Country:    model.Country,
			City:       model.City,
			Street:     model.Street,
			ZipCode:    model.ZipCode,
			IsPrimary:  model.IsPrimary,
		})
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "address updated"})
		return
	}
}

func DeleteCustomerAddressHandler(s services.CustomerAddressService) gin.HandlerFunc {
	return func(c *gin.Context) {
		customerId, done := getIdFromRequest(c)
		if done {
			return
		}
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		address, err := s.GetCustomerAddressById(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if address.CustomerId != customerId {
			c.JSON(400, gin.H{"error": "customer id does not match"})
			return
		}
		err = s.DeleteCustomerAddress(c.Request.Context(), int32(id))
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "address deleted"})
		return
	}
}

func GetAddressesByCustomerHandler(s services.CustomerAddressService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, done := getIdFromRequest(c)
		if done {
			return
		}
		addresses, err := s.GetCustomerAddressesByCustomerId(c.Request.Context(), id)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, addresses)
		return
	}
}

func InitCustomerRouter(r *gin.Engine, customerService services.CustomerService,
	credService services.UserCredentialService, addressService services.CustomerAddressService) {

	customerGroup := r.Group("/customer")
	{
		// Публичные маршруты (без аутентификации)
		customerGroup.POST("/register", RegisterCustomerHandler(customerService, credService, addressService))
		customerGroup.POST("/login", LoginHandler(credService))
		customerGroup.GET("/:id", GetCustomerByIdHandler(customerService))
		customerGroup.GET("/search/:substr", GetBySubstringHandler(customerService))

		// Приватные маршруты (требуют аутентификации)
		privateGroup := customerGroup.Group("/")
		authMiddleware := midleware.NewAuthMiddleware(os.Getenv("JWT_SECRET_KEY"))
		privateGroup.Use(authMiddleware.AuthRequired())
		{
			privateGroup.GET("/self", GetSelfHandler(customerService))
			privateGroup.PUT("/credentials", UpdateUserCredHandler(customerService, credService))
			privateGroup.POST("/address", AddCustomerAddressHandler(addressService))
			privateGroup.PUT("/address", UpdateCustomerAddressHandler(addressService))
			privateGroup.DELETE("/address/:id", DeleteCustomerAddressHandler(addressService))
			privateGroup.GET("/addresses", GetAddressesByCustomerHandler(addressService))
		}
	}
}
