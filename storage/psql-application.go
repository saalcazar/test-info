package storage

import (
	"database/sql"
	"fmt"

	"github.com/saalcazar/ceadlbk-info/model"
)

type scannerApplication interface {
	Scan(dest ...any) error
}

const (
	psqlCreateApplication   = `SELECT create_application ($1, $2, $3, $4, $5, $6, $7)`
	psqlUpdateApplication   = `SELECT update_application ($1, $2, $3, $4, $5, $6, $7, $8)`
	psqlDeleteApplication   = `SELECT delete_application ($1)`
	psqlGetApplicationByID  = `SELECT id_application, presentation, amount, approved, name_proyect, signature, name_user, nick_user, id_activity FROM applications WHERE id_application = $1`
	psqlGetAllApplication   = `SELECT id_application, presentation, amount, approved, name_proyect, signature, name_user, nick_user, id_activity FROM applications`
	psqlApprovedApplication = `SELECT update_approved_application ($1, $2)`
)

type PsqlApplication struct {
	db *sql.DB
}

func NewPsqlApplication(db *sql.DB) *PsqlApplication {
	return &PsqlApplication{db}
}

// Create
func (f *PsqlApplication) Create(m *model.Application) error {
	stmt, err := f.db.Prepare(psqlCreateApplication)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		m.Amount,
		m.Approved,
		m.NameProyect,
		m.Signature,
		m.NameUser,
		m.NickUser,
		m.IdActivity,
	).Scan(&m.NameProyect)
	if err != nil {
		return err
	}
	fmt.Println("La solicitud fue creada exitosamente")
	return nil
}

// Update
func (f *PsqlApplication) Update(m *model.Application) error {
	stmt, err := f.db.Prepare(psqlUpdateApplication)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		&m.ID,
		&m.Amount,
		&m.Approved,
		&m.NameProyect,
		&m.Signature,
		&m.NameUser,
		&m.NickUser,
		&m.IdActivity,
	)
	if err != nil {
		return err
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAfected == 0 {
		return fmt.Errorf("no existe la solicitud con el id: %d", m.ID)
	}

	fmt.Println("La solicitud se actualizó exitosamente")
	return nil
}

func (f *PsqlApplication) Approved(m *model.Application) error {
	stmt, err := f.db.Prepare(psqlApprovedApplication)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		&m.ID,
		&m.Approved,
	)
	if err != nil {
		return err
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAfected == 0 {
		return fmt.Errorf("no existe la solicitud con el id: %d", m.ID)
	}

	fmt.Println("La solicitud se aprobó exitosamente")
	return nil
}

// Delete
func (f *PsqlApplication) Delete(id uint) error {
	stmt, err := f.db.Prepare(psqlDeleteApplication)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	fmt.Println("La solicitud se eliminó correctamente")
	return nil
}

func (p *PsqlApplication) GetByID(id uint) (*model.Application, error) {
	stmt, err := p.db.Prepare(psqlGetApplicationByID)
	if err != nil {
		return &model.Application{}, err
	}
	defer stmt.Close()
	return scanRowApplications(stmt.QueryRow(id))
}

func (p *PsqlApplication) GetAll() (model.Applications, error) {
	stmt, err := p.db.Prepare(psqlGetAllApplication)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(model.Applications, 0)
	for rows.Next() {
		p, err := scanRowApplications(rows)
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

func scanRowApplications(s scannerApplication) (*model.Application, error) {
	m := &model.Application{}
	err := s.Scan(
		&m.ID,
		&m.Presentation,
		&m.Amount,
		&m.Approved,
		&m.NameProyect,
		&m.Signature,
		&m.NameUser,
		&m.NickUser,
		&m.IdActivity,
	)
	if err != nil {
		return &model.Application{}, err
	}
	return m, nil
}
