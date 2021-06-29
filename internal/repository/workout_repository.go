package repository

import (
	"github.com/ZmaximillianZ/stskp_sport_api/internal/logging"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/models"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/utils"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/jmoiron/sqlx"
)

type WorkoutRepository struct {
	BaseRepository
}

// NewWorkoutRepository creates new instance of WorkoutRepository
func NewWorkoutRepository(db *sqlx.DB, queryBuilder goqu.DialectWrapper) *WorkoutRepository {
	table := `workout`
	fields := utils.GetTagValue(models.Workout{}, tagName)
	baseQuery := queryBuilder.From(table).Select(fields...).Prepared(true)

	return &WorkoutRepository{BaseRepository{
		db:           db,
		table:        table,
		baseQuery:    baseQuery,
		queryBuilder: queryBuilder,
	}}
}

func (repo *WorkoutRepository) GetByID(id int) (models.Workout, error) {
	workout := models.Workout{}
	sql, params, err := repo.baseQuery.Where(exp.Ex{"id": id}).ToSQL()
	if err != nil {
		logging.Error(err)

		return workout, err
	}
	err = repo.db.Get(&workout, sql, params...)
	if err != nil {
		logging.Error(err)

		return workout, err
	}

	return workout, nil
}

func (repo *WorkoutRepository) GetAll() (models.Workouts, error) {
	var workouts = models.Workouts{}
	query := repo.baseQuery.Limit(maxItemsPerPage)
	sql, p, err := query.ToSQL()
	if err != nil {
		return workouts, err
	}

	err = repo.db.Select(&workouts, sql, p...)

	return workouts, err
}

func (repo *WorkoutRepository) Create(workout *models.Workout) error {
	query := repo.
		baseQuery.
		Insert().
		Into("workout").
		Cols("user_id", "description", "created_at").
		Vals(goqu.Vals{workout.UserID, workout.Description, workout.CreatedAt})

	return repo.execInsert(query)
}

func (repo *WorkoutRepository) Update(workout *models.Workout) error {
	expr := repo.baseQuery.Update().Set(workout).Where(exp.Ex{"id": workout.ID})
	return repo.execUpdate(expr)
}

func (repo *WorkoutRepository) Delete(workout *models.Workout) error {
	expr := repo.baseQuery.Delete().Where(exp.Ex{"id": workout.ID})
	return repo.execDelete(expr)
}
