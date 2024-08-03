package database

import "time"

// User represents a user entity
type User struct {
	UserID       int       `db:"user_id"`
	Username     string    `db:"username"`
	PasswordHash string    `db:"password_hash"`
	Email        string    `db:"email"`
	FullName     string    `db:"full_name"`
	JoinDate     time.Time `db:"join_date"`
}

// Member represents a member entity
type Member struct {
	MemberID        int       `db:"member_id"`
	UserID          int       `db:"user_id"`
	RegistrationNum string    `db:"reg_num"`
	MembershipLevel string    `db:"membership_level"`
	JoinDate        time.Time `db:"join_date"`
}

// CoreMember represents a core member entity
type CoreMember struct {
	CoreMemberID int       `db:"core_member_id"`
	UserID       int       `db:"user_id"`
	Role         string    `db:"role"`
	JoinDate     time.Time `db:"join_date"`
}

// Staff represents a staff entity
type Staff struct {
	StaffID    int       `db:"staff_id"`
	UserID     int       `db:"user_id"`
	Role       string    `db:"role"`
	Department string    `db:"department"`
	HireDate   time.Time `db:"hire_date"`
}

// Event represents an event entity
type Event struct {
	EventID     int       `db:"event_id"`
	EventName   string    `db:"event_name"`
	EventDate   time.Time `db:"event_date"`
	Description string    `db:"description"`
	Location    string    `db:"location"`
	CreatedAt   time.Time `db:"created_at"`
}

// Attendee represents an attendee entity
type Attendee struct {
	UserID  int `db:"user_id"`
	EventID int `db:"event_id"`
}

// GalleryImage represents a gallery image entity
type GalleryImage struct {
	ImageID     int       `db:"image_id"`
	EventID     int       `db:"event_id"`
	UserID      int       `db:"user_id"`
	ImageURL    string    `db:"image_url"`
	Description string    `db:"description"`
	UploadedAt  time.Time `db:"uploaded_at"`
}

// Comment represents a comment entity
type Comment struct {
	CommentID   int       `db:"comment_id"`
	UserID      int       `db:"user_id"`
	EventID     int       `db:"event_id"`
	CommentText string    `db:"comment_text"`
	CommentedAt time.Time `db:"commented_at"`
}

// Project represents a project entity
type Project struct {
	ProjectID   int       `db:"project_id"`
	ProjectName string    `db:"project_name"`
	Description string    `db:"description"`
	Status      string    `db:"status"`
	CreatedBy   int       `db:"created_by"`
	CreatedAt   time.Time `db:"created_at"`
}

// ProjectRequest represents a project request entity
type ProjectRequest struct {
	RequestID   int       `db:"request_id"`
	ProjectID   int       `db:"project_id"`
	UserID      int       `db:"user_id"`
	Status      string    `db:"status"`
	RequestedAt time.Time `db:"requested_at"`
}

// ProjectMember represents a project member entity
type ProjectMember struct {
	ProjectID int `db:"project_id"`
	UserID    int `db:"user_id"`
}

// Notice represents a notice entity
type Notice struct {
	NoticeID  int       `db:"notice_id"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	CreatedBy int       `db:"created_by"`
	CreatedAt time.Time `db:"created_at"`
}
