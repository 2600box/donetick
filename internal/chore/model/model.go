package model

import (
	"time"

	tModel "donetick.com/core/internal/thing/model"
)

type Chore struct {
	ID                   int                `json:"id" gorm:"primary_key"`
	Name                 string             `json:"name" gorm:"column:name"`                                      // Chore description
	FrequencyType        string             `json:"frequencyType" gorm:"column:frequency_type"`                   // "daily", "weekly", "monthly", "yearly", "adaptive",or "custom"
	Frequency            int                `json:"frequency" gorm:"column:frequency"`                            // Number of days, weeks, months, or years between chores
	FrequencyMetadata    *string            `json:"frequencyMetadata" gorm:"column:frequency_meta"`               // Additional frequency information
	NextDueDate          *time.Time         `json:"nextDueDate" gorm:"column:next_due_date;index"`                // When the chore is due
	IsRolling            bool               `json:"isRolling" gorm:"column:is_rolling"`                           // Whether the chore is rolling
	AssignedTo           int                `json:"assignedTo" gorm:"column:assigned_to"`                         // Who the chore is assigned to
	Assignees            []ChoreAssignees   `json:"assignees" gorm:"foreignkey:ChoreID;references:ID"`            // Assignees of the chore
	AssignStrategy       string             `json:"assignStrategy" gorm:"column:assign_strategy"`                 // How the chore is assigned
	IsActive             bool               `json:"isActive" gorm:"column:is_active"`                             // Whether the chore is active
	Notification         bool               `json:"notification" gorm:"column:notification"`                      // Whether the chore has notification
	NotificationMetadata *string            `json:"notificationMetadata" gorm:"column:notification_meta"`         // Additional notification information
	Labels               *string            `json:"labels" gorm:"column:labels"`                                  // Labels for the chore
	CircleID             int                `json:"circleId" gorm:"column:circle_id;index"`                       // The circle this chore is in
	CreatedAt            time.Time          `json:"createdAt" gorm:"column:created_at"`                           // When the chore was created
	UpdatedAt            time.Time          `json:"updatedAt" gorm:"column:updated_at"`                           // When the chore was last updated
	CreatedBy            int                `json:"createdBy" gorm:"column:created_by"`                           // Who created the chore
	UpdatedBy            int                `json:"updatedBy" gorm:"column:updated_by"`                           // Who last updated the chore
	ThingChore           *tModel.ThingChore `json:"thingChore" gorm:"foreignkey:chore_id;references:id;<-:false"` // ThingChore relationship
}
type ChoreAssignees struct {
	ID      int `json:"-" gorm:"primary_key"`
	ChoreID int `json:"-" gorm:"column:chore_id;uniqueIndex:idx_chore_user"`     // The chore this assignee is for
	UserID  int `json:"userId" gorm:"column:user_id;uniqueIndex:idx_chore_user"` // The user this assignee is for
}
type ChoreHistory struct {
	ID          int        `json:"id" gorm:"primary_key"`                  // Unique identifier
	ChoreID     int        `json:"choreId" gorm:"column:chore_id"`         // The chore this history is for
	CompletedAt time.Time  `json:"completedAt" gorm:"column:completed_at"` // When the chore was completed
	CompletedBy int        `json:"completedBy" gorm:"column:completed_by"` // Who completed the chore
	AssignedTo  int        `json:"assignedTo" gorm:"column:assigned_to"`   // Who the chore was assigned to
	Note        *string    `json:"notes" gorm:"column:notes"`              // Notes about the chore
	DueDate     *time.Time `json:"dueDate" gorm:"column:due_date"`         // When the chore was due
}

type FrequencyMetadata struct {
	Days   []*string `json:"days,omitempty"`
	Months []*string `json:"months,omitempty"`
	Unit   *string   `json:"unit,omitempty"`
}

type NotificationMetadata struct {
	DueDate    bool `json:"dueDate,omitempty"`
	Completion bool `json:"completion,omitempty"`
	Nagging    bool `json:"nagging,omitempty"`
	PreDue     bool `json:"predue,omitempty"`
}

type Tag struct {
	ID   int    `json:"-" gorm:"primary_key"`
	Name string `json:"name" gorm:"column:name;unique"`
}

// type ChoreTag struct {
// 	ChoreID int `json:"choreId" gorm:"primaryKey;autoIncrement:false"`
// 	TagID   int `json:"tagId" gorm:"primaryKey;autoIncrement:false"`
// }

// type CircleTag struct {
// 	CircleID int `json:"circleId" gorm:"primaryKey;autoIncrement:false"`
// 	TagID    int `json:"tagId" gorm:"primaryKey;autoIncrement:false"`
// }

type ChoreDetail struct {
	ID                  int        `json:"id" gorm:"column:id"`
	Name                string     `json:"name" gorm:"column:name"`
	FrequencyType       string     `json:"frequencyType" gorm:"column:frequency_type"`
	NextDueDate         *time.Time `json:"nextDueDate" gorm:"column:next_due_date"`
	AssignedTo          int        `json:"assignedTo" gorm:"column:assigned_to"`
	LastCompletedDate   *time.Time `json:"lastCompletedDate" gorm:"column:last_completed_date"`
	LastCompletedBy     *int       `json:"lastCompletedBy" gorm:"column:last_completed_by"`
	TotalCompletedCount int        `json:"totalCompletedCount" gorm:"column:total_completed"`
	CreatedBy           int        `json:"createdBy" gorm:"column:created_by"`
}
