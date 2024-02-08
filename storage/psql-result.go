package storage

import (
	"database/sql"
	"fmt"

	"github.com/saalcazar/ceadlbk-info/model"
)

type scannerResult interface {
	Scan(dest ...any) error
}

const (
	psqlCreateResult            = `SELECT create_project_result ($1, $2, $3, $4)`
	psqlUpdateResult            = `SELECT update_project_result ($1, $2, $3, $4, $5)`
	psqlDeleteResult            = `SELECT delete_project_result ($1)`
	pslqGetbyNameProyectResults = `SELECT id_project_result, num_project_result, project_result, name_proyect, nick_user FROM project_results WHERE name_proyect = $1`
	psqlGetAllResults           = `SELECT id_project_result, num_project_result, project_result, name_proyect, nick_user FROM project_results`
)

type PsqlResult struct {
	db *sql.DB
}

func NewPsqlResult(db *sql.DB) *PsqlResult {
	return &PsqlResult{db}
}

// Create
func (f *PsqlResult) Create(m model.ProjectResults) error {

	tx, err := f.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	stmt, err := f.db.Prepare(psqlCreateResult)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, r := range m {
		err = stmt.QueryRow(
			r.NumResult,
			r.ProjectResult,
			r.NameProyect,
			r.NickUser,
		).Scan(&r.ProjectResult)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	fmt.Println("El resultado esperado fue creado exitosamente")
	return nil
}

// Update
func (f *PsqlResult) Update(m *model.ProjectResult) error {
	stmt, err := f.db.Prepare(psqlUpdateResult)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		&m.ID,
		&m.NumResult,
		&m.ProjectResult,
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
		return fmt.Errorf("no existe el resultado esperado con el id: %d", m.ID)
	}

	fmt.Println("El resultado esperado se actualizó exitosamente")
	return nil
}

// Delete
func (f *PsqlResult) Delete(id uint) error {
	stmt, err := f.db.Prepare(psqlDeleteResult)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	fmt.Println("El resultado esperado se eliminó correctamente")
	return nil
}

func (p *PsqlResult) GetByNameProyect(nameProyect string) (model.ProjectResults, error) {
	stmt, err := p.db.Prepare(pslqGetbyNameProyectResults)
	if err != nil {
		return model.ProjectResults{}, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(nameProyect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ms := make(model.ProjectResults, 0)
	for rows.Next() {
		p, err := scanRowResults(rows)
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

func (p *PsqlResult) GetAll() (model.ProjectResults, error) {
	stmt, err := p.db.Prepare(psqlGetAllResults)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(model.ProjectResults, 0)
	for rows.Next() {
		p, err := scanRowResults(rows)
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

func scanRowResults(s scannerResult) (*model.ProjectResult, error) {
	m := &model.ProjectResult{}
	err := s.Scan(
		&m.ID,
		&m.NumResult,
		&m.ProjectResult,
		&m.NameProyect,
		&m.NickUser,
	)
	if err != nil {
		return &model.ProjectResult{}, err
	}
	return m, nil
}
