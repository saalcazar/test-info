package storage

import (
	"database/sql"
	"fmt"

	"github.com/saalcazar/ceadlbk-info/model"
)

type scannerProyect interface {
	Scan(dest ...any) error
}

const (
	psqlCreateProyect  = `SELECT create_proyect ($1, $2, $3, $4, $5)`
	psqlUpdateProyect  = `SELECT update_proyect ($1, $2, $3, $4, $5, $6)`
	psqlDeleteProyect  = `SELECT delete_proyect ($1)`
	psqlGetProyectByID = `SELECT id_proyect, cod_proyect, name_proyect, objetive, cod_founder, nick_user FROM proyects WHERE id_proyect = $1`
	psqlGetAllProyect  = `SELECT id_proyect, cod_proyect, name_proyect, objetive, cod_founder, nick_user FROM proyects`
)

type PsqlProyect struct {
	db *sql.DB
}

func NewPsqlProyect(db *sql.DB) *PsqlProyect {
	return &PsqlProyect{db}
}

// Create
func (f *PsqlProyect) Create(m *model.Proyect) error {
	stmt, err := f.db.Prepare(psqlCreateProyect)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		m.CodProyect,
		m.NameProyect,
		m.Objetive,
		m.CodFounder,
		m.NickUser,
	).Scan(&m.CodProyect)
	if err != nil {
		return err
	}
	fmt.Println("El proyecto fue creado exitosamente")
	return nil
}

// Update
func (f *PsqlProyect) Update(m *model.Proyect) error {
	stmt, err := f.db.Prepare(psqlUpdateProyect)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		&m.ID,
		&m.CodProyect,
		&m.NameProyect,
		&m.Objetive,
		&m.CodFounder,
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
		return fmt.Errorf("no existe el proyecto con el id: %d", m.ID)
	}

	fmt.Println("El proyecto se actualizó exitosamente")
	return nil
}

// Delete
func (f *PsqlProyect) Delete(id uint) error {
	stmt, err := f.db.Prepare(psqlDeleteProyect)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	fmt.Println("El proyecto se eliminó correctamente")
	return nil
}

func (p *PsqlProyect) GetByID(id uint) (*model.Proyect, error) {
	stmt, err := p.db.Prepare(psqlGetProyectByID)
	if err != nil {
		return &model.Proyect{}, err
	}
	defer stmt.Close()
	return scanRowProyects(stmt.QueryRow(id))
}

func (p *PsqlProyect) GetAll() (model.Proyects, error) {
	stmt, err := p.db.Prepare(psqlGetAllProyect)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(model.Proyects, 0)
	for rows.Next() {
		p, err := scanRowProyects(rows)
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

func scanRowProyects(s scannerProyect) (*model.Proyect, error) {
	m := &model.Proyect{}
	err := s.Scan(
		&m.ID,
		&m.CodProyect,
		&m.NameProyect,
		&m.Objetive,
		&m.CodFounder,
		&m.NickUser,
	)
	if err != nil {
		return &model.Proyect{}, err
	}
	return m, nil
}
