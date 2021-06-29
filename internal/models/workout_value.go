package models

import (
	"time"
)

type WorkoutValues []WorkoutValue

type WorkoutValue struct {
	ID            int        `json:"id" db:"id" binding:"required"`
	WorkoutID     int        `json:"workout_id" db:"workout_id"`
	WorkoutTypeID int        `json:"workout_type_id" db:"workout_type_id"`
	Value         *float64   `json:"value" db:"value"`
	Unit          int8       `json:"unit" db:"unit"`
	StartedAt     *time.Time `json:"started_at" db:"started_at"`
	EndedAt       *time.Time `json:"ended_at" db:"ended_at"`
}
