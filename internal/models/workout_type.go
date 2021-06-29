package models

type WorkoutTypes []WorkoutType

type WorkoutType struct {
	ID       int    `json:"id" db:"id"`
	ParentID *int   `json:"parent_id" db:"parent_id"`
	Name     string `json:"name" db:"name"`
	Type     uint16 `json:"type" db:"type"`
}
