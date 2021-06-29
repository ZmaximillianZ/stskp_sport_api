package repository

import (
	"github.com/ZmaximillianZ/stskp_sport_api/internal/logging"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/models"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/utils"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/jmoiron/sqlx"
)

// WorkoutTypeRepository struct
type WorkoutTypeRepository struct {
	BaseRepository
}

// NewWorkoutTypeRepository creates new instance of WorkoutTypeRepository
func NewWorkoutTypeRepository(db *sqlx.DB, queryBuilder goqu.DialectWrapper) *WorkoutTypeRepository {
	table := `workout_type`
	fields := utils.GetTagValue(models.WorkoutType{}, tagName)
	baseQuery := queryBuilder.From(table).Select(fields...).Prepared(true)

	return &WorkoutTypeRepository{BaseRepository{
		db:           db,
		table:        table,
		baseQuery:    baseQuery,
		queryBuilder: queryBuilder,
	}}
}

func (repo *WorkoutTypeRepository) GetByID(id int) (models.WorkoutType, error) {
	workoutType := models.WorkoutType{}
	sql, params, err := repo.baseQuery.Where(exp.Ex{"id": id}).ToSQL()
	if err != nil {
		logging.Error(err)

		return workoutType, err
	}
	err = repo.db.Get(&workoutType, sql, params...)
	if err != nil {
		logging.Error(err)

		return workoutType, err
	}

	return workoutType, nil
}

func (repo *WorkoutTypeRepository) GetAll() (models.WorkoutTypes, error) {
	var workoutTypes = models.WorkoutTypes{}
	query := repo.baseQuery.Limit(maxItemsPerPage)
	sql, p, err := query.ToSQL()
	if err != nil {
		return workoutTypes, err
	}

	err = repo.db.Select(&workoutTypes, sql, p...)

	return workoutTypes, err
}

func (repo *WorkoutTypeRepository) Create(workoutType *models.WorkoutType) error {
	query := repo.
		baseQuery.
		Insert().
		Into("workout_type").
		Cols("parent_id", "name", "type").
		Vals(goqu.Vals{workoutType.ParentID, workoutType.Name, workoutType.Type})

	return repo.execInsert(query)
}

func (repo *WorkoutTypeRepository) Update(workoutType *models.WorkoutType) error {
	expr := repo.baseQuery.Update().Set(workoutType).Where(exp.Ex{"id": workoutType.ID})
	return repo.execUpdate(expr)
}

func (repo *WorkoutTypeRepository) Delete(workout *models.WorkoutType) error {
	expr := repo.baseQuery.Delete().Where(exp.Ex{"id": workout.ID})
	return repo.execDelete(expr)
}
