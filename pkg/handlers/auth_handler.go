// handlers.go

package handlers

import (
	"log/slog"
	"net/http"
	"time"

	"git.nothr.in/nothr/g-electra/pkg/database"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type SignupRequest struct {
	Password        string `json:"password" form:"password"`
	Email           string `json:"email" form:"email"`
	FullName        string `json:"full_name" form:"full_name"`
	RegistrationNum string `json:"registration_number" form:"registration_number"`
	MembershipLevel string `json:"membership_level" form:"membership_level"`
}

func SignUp(c echo.Context) error {
	// Bind request body to SignupRequest struct
	req := new(SignupRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	// Validate input (basic validation, can be expanded)
	if req.Email == "" || req.Password == "" || req.RegistrationNum == "" || req.MembershipLevel == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "All fields are required"})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to hash password"})
	}

	// Create a User struct
	newUser := database.User{
		PasswordHash: string(hashedPassword),
		Email:        req.Email,
		FullName:     req.FullName,
		JoinDate:     time.Now(),
	}

	// Create a Member struct
	newMember := database.Member{
		RegistrationNum: req.RegistrationNum,
		MembershipLevel: req.MembershipLevel,
		JoinDate:        time.Now(),
	}

	// Save user to database
	if err := database.CreateUser(newUser, slog.Logger{}); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	// Link user to member (assuming user_id is populated after CreateUser)
	newMember.UserID = newUser.UserID // Assuming newUser.UserID is set by CreateUser function

	// Save member to database
	if err := database.CreateMember(newMember, slog.Logger{}); err != nil {
		// Rollback user creation if member creation fails (optional, depending on your application logic)
		// You can implement a rollback mechanism here
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create member"})
	}

	// Return success response
	return c.JSON(http.StatusOK, map[string]string{"message": "User and member created successfully"})
}
