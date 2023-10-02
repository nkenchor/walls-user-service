package api

import (
	"fmt"
	"walls-user-service/internal/adapter/extensions"
	"walls-user-service/internal/core/domain/dto"
	errorhelper "walls-user-service/internal/core/helper/error-helper"

	"github.com/gin-gonic/gin"
)

// @Summary Create User
// @Description Create a User
// @Tags User
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Success 200 {string} entity.UserReference "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param requestBody body dto.CreateUserDto true "Create User request body"
// @Router /api/user [post]
func (hdl *HTTPHandler) CreateUser(c *gin.Context) {
	body := dto.CreateUserDto{}
	_ = c.BindJSON(&body)

	currentUser := extensions.GetCurrentUser(c)
	if !extensions.ValidateBody(c, &body) || !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}

	user, err := hdl.userService.CreateUser(c.Request.Context(), body, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(201, gin.H{"user_reference:": user})
}

// @Summary Create Otp Request
// @Description Create Otp
// @Tags Otp
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Success 200 {string} entity.UserReference "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param requestBody body dto.CreateOtpDto true "Create Otp Request Body"
// @Router /api/user/send-otp [post]
func (hdl *HTTPHandler) CreateOtpRequest(c *gin.Context) {
	body := dto.CreateOtpDto{}
	_ = c.BindJSON(&body)
	fmt.Println(body)
	currentUser := extensions.GetCurrentUser(c)
	fmt.Println(currentUser)
	if !extensions.ValidateBody(c, &body) || !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}

	user, err := hdl.userService.CreateOtpRequest(c.Request.Context(), body, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(201, gin.H{"request_reference:": user})
}

// @Summary Validate Otp
// @Description Validate Otp
// @Tags Validate Otp
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Success 200 {string} entity.UserReference "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param requestBody body dto.ValidateOtpDto true "Validate Otp Request Body"
// @Router /api/user/validate-otp [post]
func (hdl *HTTPHandler) ValidateOtpRequest(c *gin.Context) {
	body := dto.ValidateOtpDto{}
	_ = c.BindJSON(&body)
	userReference := c.Param("user_reference")

	currentUser := extensions.GetCurrentUser(c)
	if !extensions.ValidateBody(c, &body) || !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}

	identity, err := hdl.userService.ValidateOtpRequest(c.Request.Context(), userReference, body, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}

	c.JSON(201, gin.H{"user_reference:": identity})
}

// @Summary Create Identity
// @Description Create Identity
// @Tags default
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Success 200 {string} entity.UserReference "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param requestBody body dto.IdentityDto true "Validate Otp Request Body"
// @Router /api/user/create-identity [post]
func (hdl *HTTPHandler) CreateIdentityRequest(c *gin.Context) {
	body := dto.IdentityDto{}
	_ = c.BindJSON(&body)
	userReference := c.Param("user_reference")

	currentUser := extensions.GetCurrentUser(c)
	if !extensions.ValidateBody(c, &body) || !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}

	user, err := hdl.userService.CreateIdentityRequest(c.Request.Context(), userReference, body, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(201, gin.H{"request_reference:": user})
}

// @Summary Upgrade Tier
// @Description Upgrade Tier
// @Tags default
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Success 200 {string} entity.UserReference "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param requestBody body dto.TierUpgradeRequestDto true "Upgrade Tier Request Body"
// @Router /api/user/upgrade-tier [post]
func (hdl *HTTPHandler) UpgradeTierRequest(c *gin.Context) {
	body := dto.TierUpgradeRequestDto{}
	_ = c.BindJSON(&body)
	userReference := c.Param("user_reference")

	currentUser := extensions.GetCurrentUser(c)
	if !extensions.ValidateBody(c, &body) || !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}

	user, err := hdl.userService.UpgradeTierRequest(c.Request.Context(), userReference, body, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(201, gin.H{"request_reference:": user})
}

// @Summary Create transaction request
// @Description Create a transaction request
// @Tags default
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Success 200 {string} entity.UserReference "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param requestBody body dto.CreateTransactionDto true "Create Transaction Request Body"
// @Router /api/user/transaction [post]
func (hdl *HTTPHandler) CreateTransactionRequest(c *gin.Context) {
	body := dto.CreateTransactionDto{}
	_ = c.BindJSON(&body)
	userReference := c.Param("user_reference")

	currentUser := extensions.GetCurrentUser(c)
	if !extensions.ValidateBody(c, &body) || !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}

	user, err := hdl.userService.CreateTransactionRequest(c.Request.Context(), userReference, body, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(201, gin.H{"request_reference:": user})
}

// @Summary Create User Reference
// @Description Validate Otp
// @Tags User Reference
// @Accept json
// @Produce json
// @Success 200 {string} entity.UserReference "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Router /api/user/user-reference [get]
func (hdl *HTTPHandler) CreateUserReference(c *gin.Context) {

	user, err := hdl.userService.CreateUserReference(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.CreateError, err.Error()))
		return
	}
	c.JSON(201, gin.H{"user_reference:": user})
}

// @Summary Create User Reference
// @Description Validate Otp
// @Tags User Reference
// @Accept json
// @Produce json
// @Success 200 {string} entity.UserReference "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Router /api/user/document-reference [get]
func (hdl *HTTPHandler) CreateDocumentReference(c *gin.Context) {

	user, err := hdl.userService.CreateDocumentReference(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.CreateError, err.Error()))
		return
	}
	c.JSON(201, gin.H{"document_reference:": user})
}

// @Summary Create Company Profile
// @Description Create a Company Profile
// @Tags Company
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Param user_reference path string true "User reference"
// @Success 200 {string} entity.UserReference "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param requestBody body dto.CreateCompanyProfileDto true "Create Company Profile request body"
// @Router /api/user/{user_reference}/company [post]
func (hdl *HTTPHandler) CreateCompanyProfile(c *gin.Context) {
	reference := c.Param("user_reference")
	body := dto.CreateCompanyProfileDto{}
	_ = c.BindJSON(&body)

	currentUser := extensions.GetCurrentUser(c)
	if !extensions.ValidateBody(c, &body) || !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}

	user, err := hdl.userService.CreateCompanyProfile(c.Request.Context(), reference, body, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(201, gin.H{"user_reference:": user})
}

// @Summary Create Company Walls Badge
// @Description Create a Company Walls Badge
// @Tags Company
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Param user_reference path string true "User reference"
// @Success 200 {string} string "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param requestBody body dto.CompanyWallsBadgeDto true "Walls Badge request body"
// @Router /api/user/{user_reference}/company/walls-badge [post]
// Need to add a channel and contact field to the swagger docs
func (hdl *HTTPHandler) CreateCompanyWallsBadge(c *gin.Context) {
	reference := c.Param("user_reference")
	body := dto.CompanyWallsBadgeDto{}
	_ = c.BindJSON(&body)

	currentUser := extensions.GetCurrentUser(c)
	if !extensions.ValidateBody(c, &body) || !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}
	user, err := hdl.userService.CreateCompanyWallsBadge(c.Request.Context(), reference, body, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(201, gin.H{"user_reference:": user})
}

// @Summary Create User Walls Badge
// @Description Create a User Walls Badge
// @Tags User
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Param user_reference path string true "User reference"
// @Success 200 {string} string "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param requestBody body dto.UserWallsBadgeDto true "Create User Walls Badge request body"
// @Router /api/user/{user_reference}/user/walls-badge [post]
// Need to add a channel and contact field to the swagger docs
func (hdl *HTTPHandler) CreateUserWallsBadge(c *gin.Context) {
	reference := c.Param("user_reference")
	body := dto.UserWallsBadgeDto{}
	_ = c.BindJSON(&body)

	currentUser := extensions.GetCurrentUser(c)
	if !extensions.ValidateBody(c, &body) || !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}
	user, err := hdl.userService.CreateUserWallsBadge(c.Request.Context(), reference, body, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(201, gin.H{"user_reference:": user})
}

// @Summary Update Comapny Profile
// @Description Update a Comapny Profile
// @Tags Company
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Success 200 {string} entity.UserReference "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param user_reference path string true "User reference"
// @Param company_reference path string true "Company reference"
// @Param requestBody body dto.UpdateCompanyProfileDto true "Update company profile request body"
// @Router /api/user/{user_reference}/company/{company_reference} [put]
// Need to add a channel and contact field to the swagger docs
func (hdl *HTTPHandler) UpdateCompanyProfile(c *gin.Context) {
	reference := c.Param("user_reference")
	coompanyReference := c.Param("company_reference")
	body := dto.UpdateCompanyProfileDto{}
	_ = c.BindJSON(&body)

	currentUser := extensions.GetCurrentUser(c)
	if !extensions.ValidateBody(c, &body) || !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}
	user, err := hdl.userService.UpdateCompanyProfile(c.Request.Context(), reference, coompanyReference, body, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(201, gin.H{"user_reference:": user})
}

// @Summary Disable Company Wallsbadge
// @Description Disable a Company Wallsbadge
// @Tags Company
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Param user_reference path string true "User reference"
// @Param company_reference path string true "Company reference"
// @Param walls_badge_reference path string true "Walls Badge reference"
// @Success 200 {string} entity.UserReference "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Router /api/user/{user_reference}/company/{company_reference}/walls-badge/{walls_badge_reference}/disable [put]
func (hdl *HTTPHandler) DisableCompanyWallsBadge(c *gin.Context) {
	reference := c.Param("user_reference")
	companyReference := c.Param("company_reference")
	companyWallsBadgeReference := c.Param("walls_badge_reference")

	currentUser := extensions.GetCurrentUser(c)
	if !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}

	result, err := hdl.userService.DisableCompanyWallsBadge(c.Request.Context(), reference, companyReference, companyWallsBadgeReference, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}

	c.JSON(200, gin.H{"user_reference": result})
}

// @Summary Disable User Wallsbadge
// @Description Disable a User Wallsbadge
// @Tags User
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Param user_reference path string true "User reference"
// @Param walls_badge_reference path string true "User Walls Badge reference"
// @Success 200 {string} entity.UserReference "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Router /api/user/{user_reference}/walls-badge/{walls_badge_reference}/disable [put]
func (hdl *HTTPHandler) DisableUserWallsBadge(c *gin.Context) {
	reference := c.Param("user_reference")
	wallsBadgeReference := c.Param("walls_badge_reference")
	currentUser := extensions.GetCurrentUser(c)

	if !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}
	result, err := hdl.userService.DisableUserWallsBadge(c.Request.Context(), reference, wallsBadgeReference, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}

	c.JSON(200, gin.H{"user_reference": result})
}

// @Summary Get company Walls Badge List
// @Description Get all wallsbadges for a company
// @Tags Company
// @Accept json
// @Produce json
// @Param user_reference path string true "User reference"
// @Param company_reference path string true "Company reference"
// @Success 200 {array} interface{} "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Router /api/user/{user_reference}/company/{company_reference}/walls-badges [get]
func (hdl *HTTPHandler) GetCompanyWallsBadgeList(c *gin.Context) {
	reference := c.Param("user_reference")
	companyReference := c.Param("company_reference")

	result, err := hdl.userService.GetCompanyWallsBadgeList(c.Request.Context(), reference, companyReference)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}

	c.JSON(200, result)
}

// @Summary Get user Walls Badge List
// @Description Get all wallsbadges for a user
// @Tags User
// @Accept json
// @Produce json
// @Param user_reference path string true "User reference"
// @Success 200 {string} interface{} "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Router /api/user/{user_reference}/walls-badges [get]
func (hdl *HTTPHandler) GetUserWallsBadgeList(c *gin.Context) {
	reference := c.Param("user_reference")

	result, err := hdl.userService.GetUserWallsBadgeList(c.Request.Context(), reference)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}

	c.JSON(200, result)
}

// @Summary Get default company walls badge
// @Description Get the default wallsbadge for a company
// @Tags Company
// @Accept json
// @Produce json
// @Param user_reference path string true "User reference"
// @Param company_reference path string true "Company reference"
// @Success 200 {string} interface{} "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Router /api/user/{user_reference}/company/{company_reference}/walls-badge/default [get]
func (hdl *HTTPHandler) GetDefaultCompanyWallsBadge(c *gin.Context) {
	reference := c.Param("user_reference")
	companyReference := c.Param("company_reference")

	result, err := hdl.userService.GetDefaultCompanyWallsBadge(c.Request.Context(), reference, companyReference)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}

	c.JSON(200, result)
}

// @Summary Get default user walls badge
// @Description Get the default wallsbadge for a user
// @Tags User
// @Accept json
// @Produce json
// @Param user_reference path string true "User reference"
// @Success 200 {string} interface{} "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Router /api/user/{user_reference}/walls-badge/default [get]
func (hdl *HTTPHandler) GetDefaultUserWallsBadge(c *gin.Context) {
	reference := c.Param("user_reference")

	result, err := hdl.userService.GetDefaultUserWallsBadge(c.Request.Context(), reference)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}

	c.JSON(200, result)
}

// @Summary Disable Company Profile
// @Description Disable a Company Profile
// @Tags Company
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Param user_reference path string true "User reference"
// @Param company_reference path string true "Company reference"
// @Success 200 {string} interface{} "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Router /api/user/{user_reference}/company/{company_reference}/disable [put]
func (hdl *HTTPHandler) DisableCompanyProfile(c *gin.Context) {
	reference := c.Param("user_reference")
	companyReference := c.Param("company_reference")

	result, err := hdl.userService.DisableCompanyProfile(c.Request.Context(), reference, companyReference)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}

	c.JSON(200, gin.H{"user_reference": result})
}

// @Summary Update Comapny Logo
// @Description Update a Comapny Logo
// @Tags Company
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Success 200 {string} string "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param user_reference path string true "User reference"
// @Param company_reference path string true "Company reference"
// @Param requestBody body dto.UpdateCompanyLogo true "Walls Badge request body"
// @Router /api/user/{user_reference}/company/{company_reference}/logo [put]
// Need to add a channel and contact field to the swagger docs
func (hdl *HTTPHandler) UpdateCompanyLogo(c *gin.Context) {
	reference := c.Param("user_reference")
	coompanyReference := c.Param("company_reference")
	body := dto.UpdateCompanyLogo{}
	_ = c.BindJSON(&body)

	currentUser := extensions.GetCurrentUser(c)
	if !extensions.ValidateBody(c, &body) || !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}
	user, err := hdl.userService.UpdateCompanyLogo(c.Request.Context(), reference, coompanyReference, body, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(201, gin.H{"user_reference:": user})
}

// @Summary Update User Email Status
// @Description Update user email verification status
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {string} string "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param user_reference path string true "User reference"
// @Router /api/user/{user_reference}/email/update-status [put]
func (hdl *HTTPHandler) UpdateUserProfileEmailStatus(c *gin.Context) {
	reference := c.Param("user_reference")

	response, err := hdl.userService.UpdateUserProfileEmailStatus(c.Request.Context(), reference)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(200, gin.H{"user_reference:": response})
}

// @Summary Update User Phone Status
// @Description Update user Phone number verification status
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {string} string "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param user_reference path string true "User reference"
// @Router /api/user/{user_reference}/phone/update-status [put]
// func (hdl *HTTPHandler) UpdateUserProfilePhoneStatus(c *gin.Context) {
// 	reference := c.Param("user_reference")

// 	response, err := hdl.userService.UpdateUserProfilePhoneStatus(c.Request.Context(), reference)
// 	if err != nil {
// 		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
// 		return
// 	}
// 	c.JSON(200, gin.H{"user_reference:": response})
// }

// @Summary Update Company Email Status
// @Description Update company email verification status
// @Tags Company
// @Accept json
// @Produce json
// @Success 200 {string} string "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param user_reference path string true "User reference"
// @Param company_reference path string true "Company reference"
// @Router /api/user/{user_reference}/company/{company_reference}/email/update-status [put]
func (hdl *HTTPHandler) UpdateCompanyProfileEmailStatus(c *gin.Context) {
	reference := c.Param("user_reference")
	companyReference := c.Param("company_reference")

	response, err := hdl.userService.UpdateCompanyProfileEmailStatus(c.Request.Context(), reference, companyReference)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(200, gin.H{"user_reference:": response})
}



// @Summary Set Default bank for user
// @Description Set user's default bank
// @Tags User
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Success 200 {string} string "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param user_reference path string true "User reference"
// @Param bank_reference path string true "Bank reference"
// @Router /api/user/{user_reference}/bank/{bank_reference}/set-default [put]
func (hdl *HTTPHandler) SetDefaultBank(c *gin.Context) {
	reference := c.Param("user_reference")
	bankReference := c.Param("bank_reference")

	currentUser := extensions.GetCurrentUser(c)
	if !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}

	response, err := hdl.userService.SetDefaultBank(c.Request.Context(), reference, bankReference, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(200, gin.H{"user_reference:": response})
}

// @Summary Set Default Card for user
// @Description Set user default Card
// @Tags User
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Success 200 {string} string "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param user_reference path string true "User reference"
// @Param card_reference path string true "Card reference"
// @Router /api/user/{user_reference}/card/{card_reference}/set-default [put]
func (hdl *HTTPHandler) SetDefaultCard(c *gin.Context) {
	reference := c.Param("user_reference")
	cardReference := c.Param("card_reference")

	currentUser := extensions.GetCurrentUser(c)
	if !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}

	response, err := hdl.userService.SetDefaultCard(c.Request.Context(), reference, cardReference, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(200, gin.H{"user_reference:": response})
}

// @Summary Get User by Reference
// @Description Get user details by reference
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} entity.User "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param user_reference path string true "User reference"
// @Router /api/user/{user_reference} [get]
func (hdl *HTTPHandler) GetUserByReference(c *gin.Context) {
	user, err := hdl.userService.GetUserByReference(c.Request.Context(), c.Param("user_reference"))
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(200, user)
}

// @Summary Update User Name
// @Description Update user name
// @Tags User
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Success 200 {string} string "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param user_reference path string true "User reference"
// @Param requestBody body dto.UserNameDto true "User name data"
// @Router /api/user/{user_reference}/name [put]
func (hdl *HTTPHandler) UpdateUserName(c *gin.Context) {
	reference := c.Param("user_reference")
	body := dto.UserNameDto{}
	_ = c.BindJSON(&body)

	currentUser := extensions.GetCurrentUser(c)
	if !extensions.ValidateBody(c, &body) || !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}
	response, err := hdl.userService.UpdateUserName(c.Request.Context(), reference, body, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(200, gin.H{"user_reference:": response})
}

// @Summary Update User Email
// @Description Update user email
// @Tags User
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Success 200 {string} string "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param user_reference path string true "User reference"
// @Param requestBody body dto.EmailDto true "User email data"
// @Router /api/user/{user_reference}/email [put]
func (hdl *HTTPHandler) UpdateEmail(c *gin.Context) {
	reference := c.Param("user_reference")
	body := dto.EmailDto{}
	_ = c.BindJSON(&body)

	fmt.Println(body)
	currentUser := extensions.GetCurrentUser(c)
	fmt.Println(currentUser)
	if !extensions.ValidateBody(c, &body) || !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}
	response, err := hdl.userService.UpdateEmail(c.Request.Context(), reference, body, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(200, gin.H{"user_reference:": response})
}

// @Summary Update User Date of Birth
// @Description Update user's date of birth
// @Tags User
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Success 200 {string} string "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param user_reference path string true "User reference"
// @Param requestBody body dto.DobDto true "User date of birth data"
// @Router /api/user/{user_reference}/dob [put]
func (hdl *HTTPHandler) UpdateDateOfBirth(c *gin.Context) {
	reference := c.Param("user_reference")
	body := dto.DobDto{}
	_ = c.BindJSON(&body)

	currentUser := extensions.GetCurrentUser(c)
	if !extensions.ValidateBody(c, &body) || !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}
	response, err := hdl.userService.UpdateDateOfBirth(c.Request.Context(), reference, body, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(200, gin.H{"user_reference:": response})
}

// @Summary Update User Address
// @Description Update user's address
// @Tags User
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Success 200 {string} string "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param user_reference path string true "User reference"
// @Param requestBody body dto.AddressDto true "User address data"
// @Router /api/user/{user_reference}/address [put]
func (hdl *HTTPHandler) UpdateAddress(c *gin.Context) {
	reference := c.Param("user_reference")
	body := dto.AddressDto{}
	_ = c.BindJSON(&body)

	currentUser := extensions.GetCurrentUser(c)
	if !extensions.ValidateBody(c, &body) || !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}
	response, err := hdl.userService.UpdateAddress(c.Request.Context(), reference, body, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(200, gin.H{"user_reference:": response})
}

// @Summary Get User by WallsTag
// @Description Get user details by WallsTag
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} entity.User "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param wallsTag path string true "WallsTag"
// @Router /api/user/walls-tag/{wallsTag} [get]
func (hdl *HTTPHandler) GetUserByWallsTag(c *gin.Context) {
	user, err := hdl.userService.GetUserByWallsTag(c.Request.Context(), c.Param("wallsTag"))
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(200, user)
}

// @Summary Get User by Walls Badge Reference
// @Description Get user details by Badge Reference
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} entity.User "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param walls_badge_reference path string true "Walls Badge Reference"
// @Router /api/user/walls-badge-reference/{walls_badge_reference} [get]
func (hdl *HTTPHandler) GetUserByWallsBagdeReference(c *gin.Context) {
	user, err := hdl.userService.GetUserByWallsBagdeReference(c.Request.Context(), c.Param("walls_badge_reference"))
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(200, user)
}

// @Summary Get User by Phone
// @Description Get user details by Phone
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} entity.User "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param phone path string true "Phone"
// @Router /api/user/phone/{phone} [get]
func (hdl *HTTPHandler) GetUserByPhone(c *gin.Context) {
	user, err := hdl.userService.GetUserByPhone(c.Request.Context(), c.Param("phone"))
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(200, user)
}

// @Summary Get User by Device
// @Description Get user details by device
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} entity.User "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param requestBody body dto.DeviceDto true "Device data"
// @Router /api/user/device [post]
func (hdl *HTTPHandler) GetUserByDevice(c *gin.Context) {
	body := dto.DeviceDto{}
	_ = c.BindJSON(&body)

	if !extensions.ValidateBody(c, &body) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}

	user, err := hdl.userService.GetUserByDevice(c.Request.Context(), body)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(200, user)
}

// @Summary Update User's Photos
// @Description Update a User's Photos
// @Tags User
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Success 200 {string} string "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param user_reference path string true "User reference"
// @Param requestBody body dto.PhotoDto true "Photo request body"
// @Router /api/user/{user_reference}/photos [put]
func (hdl *HTTPHandler) UpdatePhoto(c *gin.Context) {
	reference := c.Param("user_reference")
	body := dto.PhotoDto{}
	_ = c.BindJSON(&body)

	currentUser := extensions.GetCurrentUser(c)
	if !extensions.ValidateBody(c, &body) || !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}
	user, err := hdl.userService.UpdatePhoto(c.Request.Context(), reference, body, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(200, gin.H{"user_reference": user})
}

// @Summary Update User Wallet
// @Description Update a User's Wallet
// @Tags User
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Success 200 {string} string "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param user_reference path string true "User reference"
// @Param requestBody body dto.UpdateWalletDto true "Wallet request body"
// @Router /api/user/{user_reference}/wallet [put]
func (hdl *HTTPHandler) UpdateWallet(c *gin.Context) {
	reference := c.Param("user_reference")
	body := dto.UpdateWalletDto{}
	_ = c.BindJSON(&body)

	currentUser := extensions.GetCurrentUser(c)
	if !extensions.ValidateBody(c, &body) || !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}
	user, err := hdl.userService.UpdateWallet(c.Request.Context(), reference, body, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(200, gin.H{"user_reference": user})
}

// @Summary Add Bank to User
// @Description Add a Bank to a User
// @Tags User
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Success 200 {string} string "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param user_reference path string true "User reference"
// @Param requestBody body dto.BankDto true "Bank request body"
// @Router /api/user/{user_reference}/bank [post]
func (hdl *HTTPHandler) AddBank(c *gin.Context) {
	reference := c.Param("user_reference")
	body := dto.BankDto{}
	_ = c.BindJSON(&body)

	currentUser := extensions.GetCurrentUser(c)
	if !extensions.ValidateBody(c, &body) || !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}
	user, err := hdl.userService.AddBank(c.Request.Context(), reference, body, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(201, gin.H{"user_reference": user})
}

// @Summary Update User's Bank
// @Description Update a User's Bank Details
// @Tags User
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Success 200 {string} string "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param user_reference path string true "User reference"
// @Param bank_reference path string true "Bank reference"
// @Param requestBody body dto.BankDto true "Bank request body"
// @Router /api/user/{user_reference}/bank/{bank_reference} [put]
func (hdl *HTTPHandler) UpdateBank(c *gin.Context) {
	reference := c.Param("user_reference")
	bank_reference := c.Param("bank_reference")
	body := dto.BankDto{}
	_ = c.BindJSON(&body)

	currentUser := extensions.GetCurrentUser(c)
	if !extensions.ValidateBody(c, &body) || !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}
	user, err := hdl.userService.UpdateBank(c.Request.Context(), reference, bank_reference, body, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(200, gin.H{"user_reference": user})
}

// @Summary Add Card to User
// @Description Add a Card to a User
// @Tags User
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Success 200 {string} string "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param user_reference path string true "User reference"
// @Param requestBody body dto.CardDto true "Card request body"
// @Router /api/user/{user_reference}/card [post]
func (hdl *HTTPHandler) AddCard(c *gin.Context) {
	reference := c.Param("user_reference")
	body := dto.CardDto{}
	_ = c.BindJSON(&body)

	currentUser := extensions.GetCurrentUser(c)
	if !extensions.ValidateBody(c, &body) || !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}
	user, err := hdl.userService.AddCard(c.Request.Context(), reference, body, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(201, gin.H{"user_reference": user})
}

// @Summary Update User's Card
// @Description Update a User's Card
// @Tags User
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Success 200 {string} string "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param user_reference path string true "User reference"
// @Param card_reference path string true "Card reference"
// @Param requestBody body dto.CardDto true "Card request body"
// @Router /api/user/{user_reference}/card/{card_reference} [put]
func (hdl *HTTPHandler) UpdateCard(c *gin.Context) {
	reference := c.Param("user_reference")
	card_reference := c.Param("card_reference")
	body := dto.CardDto{}
	_ = c.BindJSON(&body)

	currentUser := extensions.GetCurrentUser(c)
	if !extensions.ValidateBody(c, &body) || !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}
	user, err := hdl.userService.UpdateCard(c.Request.Context(), reference, card_reference, body, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(200, gin.H{"user_reference": user})
}

// @Summary Update User's Notification Options
// @Description Update a User's Notification Options
// @Tags User
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Success 200 {string} string "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param user_reference path string true "User reference"
// @Param requestBody body dto.UpdateNotificationOptionsDto true "Notification options request body"
// @Router /api/user/{user_reference}/notification-options [put]
func (hdl *HTTPHandler) UpdateNotificationOptions(c *gin.Context) {
	reference := c.Param("user_reference")
	body := dto.UpdateNotificationOptionsDto{}
	_ = c.BindJSON(&body)

	currentUser := extensions.GetCurrentUser(c)
	if !extensions.ValidateBody(c, &body) || !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}
	user, err := hdl.userService.UpdateNotificationOptions(c.Request.Context(), reference, body, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(200, gin.H{"user_reference": user})
}

// @Summary Update User's Device
// @Description Update a User's Device
// @Tags User
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Success 200 {string} string "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param user_reference path string true "User reference"
// @Param requestBody body dto.UpdateDeviceDto true "Device request body"
// @Router /api/user/{user_reference}/device [put]
func (hdl *HTTPHandler) UpdateDevice(c *gin.Context) {
	reference := c.Param("user_reference")
	body := dto.UpdateDeviceDto{}
	_ = c.BindJSON(&body)

	currentUser := extensions.GetCurrentUser(c)
	if !extensions.ValidateBody(c, &body) || !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}
	user, err := hdl.userService.UpdateDevice(c.Request.Context(), reference, body, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(200, gin.H{"user_reference": user})
}

// @Summary Add Documentation
// @Description Add an documentation
// @Tags User
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Success 200 {string} string "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param user_reference path string true "User reference"
// @Param requestBody body dto.AddDocumentationDto true "Documentation request body"
// @Router /api/user/{user_reference}/documentation [post]
func (hdl *HTTPHandler) AddDocumentation(c *gin.Context) {
	reference := c.Param("user_reference")
	body := dto.AddDocumentationDto{}
	_ = c.BindJSON(&body)

	currentUser := extensions.GetCurrentUser(c)
	if !extensions.ValidateBody(c, &body) || !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}
	documentation, err := hdl.userService.AddDocumentation(c.Request.Context(), reference, body, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(201, gin.H{"user_reference": documentation})
}

// @Summary Update Documentation
// @Description Update an documentation
// @Tags User
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Success 200 {string} string "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param user_reference path string true "User reference"
// @Param documentation_reference path string true "Documentation reference"
// @Param requestBody body dto.AddDocumentationDto true "Documentation request body"
// @Router /api/user/{user_reference}/documentation/{documentation_reference} [put]
func (hdl *HTTPHandler) UpdateDocumentation(c *gin.Context) {
	reference := c.Param("user_reference")
	documentationReference := c.Param("documentation_reference")
	body := dto.AddDocumentationDto{}
	_ = c.BindJSON(&body)
	currentUser := extensions.GetCurrentUser(c)
	if !extensions.ValidateBody(c, &body) || !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}
	updatedDocumentation, err := hdl.userService.UpdateDocumentation(c.Request.Context(), reference, documentationReference, body, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(200, gin.H{"user_reference": updatedDocumentation})
}

// @Summary Add Contact
// @Description Add a contact
// @Tags User
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Success 200 {string} string "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Param user_reference path string true "User reference"
// @Param requestBody body dto.ContactDto true "Contact request body"
// @Router /api/user/{user_reference}/contact [post]
func (hdl *HTTPHandler) AddContact(c *gin.Context) {
	reference := c.Param("user_reference")
	body := dto.ContactDto{}
	_ = c.BindJSON(&body)
	currentUser := extensions.GetCurrentUser(c)
	if !extensions.ValidateBody(c, &body) || !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}
	contact, err := hdl.userService.AddContact(c.Request.Context(), reference, body, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(201, gin.H{"user_reference": contact})
}

// @Summary Update Balance
// @Description Update the balance of a user
// @Tags User
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Param user_reference path string true "User reference"
// @Param requestBody body dto.BalanceDto true "Balance request body"
// @Success 200 {string} string "Success"
// @Failure 400 {object} helper.ErrorResponse
// @Failure 500 {object} helper.ErrorResponse
// @Router /api/user/{user_reference}/balance [put]
func (hdl *HTTPHandler) UpdateBalance(c *gin.Context) {
	reference := c.Param("user_reference")
	body := dto.BalanceDto{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(400, errorhelper.ErrorMessage(errorhelper.BadRequestError, err.Error()))
		return
	}

	currentUser := extensions.GetCurrentUser(c)
	if !extensions.ValidateBody(c, &body) || !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}

	result, err := hdl.userService.UpdateBalance(c.Request.Context(), reference, body, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}
	c.JSON(200, gin.H{"user_reference": result})
}

// @Summary Update User Tier
// @Description Update the user's tier
// @Tags User
// @Accept json
// @Produce json
// @Param X-User-Reference header string true "User Reference"
// @Param X-Phone header string true "Phone"
// @Param X-Imei header string true "IMEI"
// @Param X-Device-Type header string true "Device Type"
// @Param X-Device-Brand header string true "Device Brand"
// @Param X-Device-Model header string true "Device Model"
// @Param X-Device-Reference header string true "Device Reference"
// @Param user_reference path string true "User reference"
// @Param requestBody body dto.TierDto true "Tier request body"
// @Success 200 {string} interface{} "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Router /api/user/{user_reference}/tier [put]
func (hdl *HTTPHandler) UpdateTier(c *gin.Context) {
	reference := c.Param("user_reference")
	body := dto.TierDto{}
	_ = c.BindJSON(&body)

	currentUser := extensions.GetCurrentUser(c)
	if !extensions.ValidateBody(c, &body) || !extensions.ValidateHeaders(c, currentUser) {
		c.AbortWithStatusJSON(400, gin.H{"error": "invalid request in request body or request headers"})
		return
	}

	result, err := hdl.userService.UpdateTier(c.Request.Context(), reference, body, currentUser)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}

	c.JSON(200, gin.H{"user_reference": result})
}

// @Summary Add User coupon
// @Description Add a coupon for the user
// @Tags User
// @Accept json
// @Produce json
// @Param user_reference path string true "User reference"
// @Param requestBody body dto.CouponDto true "Coupon request body"
// @Success 200 {string} interface{} "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Router /api/user/{user_reference}/coupon [put]
func (hdl *HTTPHandler) AddCoupon(c *gin.Context) {
	reference := c.Param("user_reference")
	body := dto.CouponDto{}
	_ = c.BindJSON(&body)

	result, err := hdl.userService.AddCoupon(c.Request.Context(), reference, body)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}

	c.JSON(200, gin.H{"user_reference": result})
}

// @Summary Update User Rewards
// @Description Update the user's rewards
// @Tags User
// @Accept json
// @Produce json
// @Param user_reference path string true "User reference"
// @Param requestBody body dto.RewardDto true "Reward request body"
// @Success 200 {string} interface{} "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Router /api/user/{user_reference}/reward [put]
func (hdl *HTTPHandler) UpdateRewards(c *gin.Context) {
	reference := c.Param("user_reference")
	body := dto.RewardDto{}
	_ = c.BindJSON(&body)

	result, err := hdl.userService.UpdateRewards(c.Request.Context(), reference, body)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}

	c.JSON(200, gin.H{"user_reference": result})
}



// @Summary Enable User
// @Description Enable a user
// @Tags User
// @Accept json
// @Param user_reference path string true "User reference"
// @Success 200 {string} interface{} "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Router /api/user/{user_reference}/enable [put]
func (hdl *HTTPHandler) EnableUser(c *gin.Context) {
	reference := c.Param("user_reference")

	result, err := hdl.userService.EnableUser(c.Request.Context(), reference)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}

	c.JSON(200, gin.H{"user_reference": result})
}

// @Summary Disable User
// @Description Disable a user
// @Tags User
// @Accept json
// @Produce json
// @Param user_reference path string true "User reference"
// @Success 200 {string} interface{} "Success"
// @Failure 500 {object} helper.ErrorResponse
// @Router /api/user/{user_reference}/disable [put]
func (hdl *HTTPHandler) DisableUser(c *gin.Context) {
	reference := c.Param("user_reference")

	result, err := hdl.userService.DisableUser(c.Request.Context(), reference)
	if err != nil {
		c.AbortWithStatusJSON(500, errorhelper.ErrorMessage(errorhelper.MongoDBError, err.Error()))
		return
	}

	c.JSON(200, gin.H{"user_reference": result})
}
