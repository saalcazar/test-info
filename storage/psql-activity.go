package storage

import (
	"database/sql"
	"fmt"

	"github.com/saalcazar/ceadlbk-info/model"
)

type scannerActivity interface {
	Scan(dest ...any) error
}

const (
	psqlCreateActivity  = `SELECT create_activity ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`
	psqlUpdateActivity  = `SELECT update_activity ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`
	psqlDeleteActivity  = `SELECT delete_activity ($1)`
	psqlGetActivityByID = `SELECT id_activity, activity, date_start, date_end, place, expected, objetive, result_expected, description, name_proyect, cod_founder, especific, nick_user, project_result, project_activity FROM activities WHERE id_activity = $1`
	psqlGetAllActivity  = `SELECT id_activity, activity, date_start, date_end, place, expected, objetive, result_expected, description, name_proyect, cod_founder, especific, nick_user, project_result, project_activity FROM activities`
)

type PsqlActivity struct {
	db *sql.DB
}

func NewPsqlActivity(db *sql.DB) *PsqlActivity {
	return &PsqlActivity{db}
}

// Create
func (f *PsqlActivity) Create(m *model.Activity) error {
	stmt, err := f.db.Prepare(psqlCreateActivity)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		m.Activity,
		m.DateStar,
		m.DateEnd,
		m.Place,
		m.Expected,
		m.Objetive,
		m.ResultExpected,
		m.Description,
		m.NameProyect,
		m.CodFounder,
		m.Especific,
		m.NickUser,
		m.ProjectResult,
		m.ProjectActivity,
	).Scan(&m.Activity)
	if err != nil {
		return err
	}
	fmt.Println("La actividad fue creada exitosamente")
	return nil
}

// Update
func (f *PsqlActivity) Update(m *model.Activity) error {
	stmt, err := f.db.Prepare(psqlUpdateActivity)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		&m.ID,
		&m.Activity,
		&m.DateStar,
		&m.DateEnd,
		&m.Place,
		&m.Expected,
		&m.Objetive,
		&m.ResultExpected,
		&m.Description,
		&m.NameProyect,
		&m.CodFounder,
		&m.Especific,
		&m.NickUser,
		&m.ProjectResult,
		&m.ProjectActivity,
	)
	if err != nil {
		return err
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAfected == 0 {
		return fmt.Errorf("no existe la actividad con el id: %d", m.ID)
	}

	fmt.Println("La actividad se actualizó exitosamente")
	return nil
}

// Delete
func (f *PsqlActivity) Delete(id uint) error {
	stmt, err := f.db.Prepare(psqlDeleteActivity)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	fmt.Println("La actividad se eliminó correctamente")
	return nil
}

func (p *PsqlActivity) GetByID(id uint) (*model.Activity, error) {
	stmt, err := p.db.Prepare(psqlGetActivityByID)
	if err != nil {
		return &model.Activity{}, err
	}
	defer stmt.Close()
	return scanRowActivities(stmt.QueryRow(id))

}

func (p *PsqlActivity) GetAll() (model.Activities, error) {
	stmt, err := p.db.Prepare(psqlGetAllActivity)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(model.Activities, 0)
	for rows.Next() {
		p, err := scanRowActivities(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, *p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ms, nil
}

func scanRowActivities(s scannerActivity) (*model.Activity, error) {
	m := &model.Activity{}
	err := s.Scan(
		&m.ID,
		&m.Activity,
		&m.DateStar,
		&m.DateEnd,
		&m.Place,
		&m.Expected,
		&m.Objetive,
		&m.ResultExpected,
		&m.Description,
		&m.NameProyect,
		&m.CodFounder,
		&m.Especific,
		&m.NickUser,
		&m.ProjectResult,
		&m.ProjectActivity,
	)
	if err != nil {
		return &model.Activity{}, err
	}
	return m, nil
}
