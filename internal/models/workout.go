package models

import "time"

type Workouts []Workout

type Workout struct {
	ID          int       `json:"id" db:"id"`
	UserID      int       `json:"user_id" db:"user_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	Description string    `json:"description" db:"description"`
}
