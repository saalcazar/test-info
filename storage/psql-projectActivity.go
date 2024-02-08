package storage

import (
	"database/sql"
	"fmt"

	"github.com/saalcazar/ceadlbk-info/model"
)

type scannerProjectActivity interface {
	Scan(dest ...any) error
}

const (
	psqlCreateProjectActivity    = `SELECT create_project_activity ($1, $2, $3, $4, $5)`
	psqlUpdateProjectActivity    = `SELECT update_project_activity ($1, $2, $3, $4, $5, $6)`
	psqlDeleteProjectActivity    = `SELECT delete_project_activity ($1)`
	pslqGetbyNameProyectActivity = `SELECT id_project_activity, num_project_activity, project_activity, category, name_proyect, nick_user FROM project_activity WHERE name_proyect = $1`
	psqlGetAllProjectActivity    = `SELECT id_project_activity, num_project_activity, project_activity, category, name_proyect, nick_user FROM project_activities`
)

type PsqlProjectActivity struct {
	db *sql.DB
}

func NewPsqlProjectActivity(db *sql.DB) *PsqlProjectActivity {
	return &PsqlProjectActivity{db}
}

// Create
func (f *PsqlProjectActivity) Create(m model.ProjectActivities) error {

	tx, err := f.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	stmt, err := f.db.Prepare(psqlCreateProjectActivity)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, r := range m {
		err = stmt.QueryRow(
			r.NumActivity,
			r.ProjectActivity,
			r.Category,
			r.NameProyect,
			r.NickUser,
		).Scan(&r.ProjectActivity)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	fmt.Println("La actividad fue creada exitosamente")
	return nil
}

// Update
func (f *PsqlProjectActivity) Update(m *model.ProjectActivity) error {
	stmt, err := f.db.Prepare(psqlUpdateProjectActivity)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		&m.ID,
		&m.NumActivity,
		&m.ProjectActivity,
		&m.Category,
		&m.NameProyect,
		&m.NickUser,
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
func (f *PsqlProjectActivity) Delete(id uint) error {
	stmt, err := f.db.Prepare(psqlDeleteProjectActivity)
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

func (p *PsqlProjectActivity) GetByNameProyect(nameProyect string) (model.ProjectActivities, error) {
	stmt, err := p.db.Prepare(pslqGetbyNameProyectActivity)
	if err != nil {
		return model.ProjectActivities{}, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(nameProyect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ms := make(model.ProjectActivities, 0)
	for rows.Next() {
		p, err := scanRowProjectActivity(rows)
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

func (p *PsqlProjectActivity) GetAll() (model.ProjectActivities, error) {
	stmt, err := p.db.Prepare(psqlGetAllProjectActivity)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(model.ProjectActivities, 0)
	for rows.Next() {
		p, err := scanRowProjectActivity(rows)
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

func scanRowProjectActivity(s scannerProjectActivity) (*model.ProjectActivity, error) {
	m := &model.ProjectActivity{}
	err := s.Scan(
		&m.ID,
		&m.NumActivity,
		&m.ProjectActivity,
		&m.Category,
		&m.NameProyect,
		&m.NickUser,
	)
	if err != nil {
		return &model.ProjectActivity{}, err
	}
	return m, nil
}
