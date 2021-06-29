package repository

import (
	"github.com/ZmaximillianZ/stskp_sport_api/internal/logging"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/models"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/utils"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/jmoiron/sqlx"
)

// WorkoutValueRepository struct
type WorkoutValueRepository struct {
	BaseRepository
}

// NewWorkoutValueRepository creates new instance of WorkoutValueRepository
func NewWorkoutValueRepository(db *sqlx.DB, queryBuilder goqu.DialectWrapper) *WorkoutValueRepository {
	table := `workout_value`
	fields := utils.GetTagValue(models.WorkoutValue{}, tagName)
	baseQuery := queryBuilder.From(table).Select(fields...).Prepared(true)

	return &WorkoutValueRepository{BaseRepository{
		db:           db,
		table:        table,
		baseQuery:    baseQuery,
		queryBuilder: queryBuilder,
	}}
}

func (repo *WorkoutValueRepository) GetByID(id int) (models.WorkoutValue, error) {
	WorkoutValue := models.WorkoutValue{}
	sql, params, err := repo.baseQuery.Where(exp.Ex{"id": id}).ToSQL()
	if err != nil {
		logging.Error(err)

		return WorkoutValue, err
	}
	err = repo.db.Get(&WorkoutValue, sql, params...)
	if err != nil {
		logging.Error(err)

		return WorkoutValue, err
	}

	return WorkoutValue, nil
}

func (repo *WorkoutValueRepository) GetAll() (models.WorkoutValues, error) {
	var WorkoutValues = models.WorkoutValues{}
	query := repo.baseQuery.Limit(maxItemsPerPage)
	sql, p, err := query.ToSQL()
	if err != nil {
		return WorkoutValues, err
	}

	err = repo.db.Select(&WorkoutValues, sql, p...)

	return WorkoutValues, err
}

func (repo *WorkoutValueRepository) Create(workoutValue *models.WorkoutValue) error {
	query := repo.
		baseQuery.
		Insert().
		Into("workout_value").
		Cols("workout_id", "workout_type_id", "value", "unit", "started_at", "ended_at").
		Vals(goqu.Vals{
			workoutValue.WorkoutID,
			workoutValue.WorkoutTypeID,
			workoutValue.Value,
			workoutValue.Unit,
			workoutValue.StartedAt,
			workoutValue.EndedAt,
		})

	return repo.execInsert(query)
}

func (repo *WorkoutValueRepository) Update(workoutValue *models.WorkoutValue) error {
	expr := repo.baseQuery.Update().Set(workoutValue).Where(exp.Ex{"id": workoutValue.ID})
	return repo.execUpdate(expr)
}

func (repo *WorkoutValueRepository) Delete(workout *models.WorkoutValue) error {
	expr := repo.baseQuery.Delete().Where(exp.Ex{"id": workout.ID})
	return repo.execDelete(expr)
}
