package storage

import (
	"database/sql"
	"fmt"

	"github.com/saalcazar/ceadlbk-info/model"
)

type scannerEspecific interface {
	Scan(dest ...any) error
}

const (
	psqlCreateEspecific  = `SELECT create_especific ($1, $2, $3, $4)`
	psqlUpdateEspecific  = `SELECT update_especific ($1, $2, $3, $4, $5)`
	psqlDeleteEspecific  = `SELECT delete_especific ($1)`
	pslqGetbyNameProyect = `SELECT id_especific, num_especific, especific, nick_user, name_proyect FROM especifics WHERE name_proyect = $1`
	psqlGetAllEspecifics = `SELECT id_especific, num_especific, especific, nick_user, name_proyect FROM especifics`
)

type PsqlEspecific struct {
	db *sql.DB
}

func NewPsqlEspecific(db *sql.DB) *PsqlEspecific {
	return &PsqlEspecific{db}
}

// Create
func (f *PsqlEspecific) Create(m model.Especifics) error {

	tx, err := f.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	stmt, err := f.db.Prepare(psqlCreateEspecific)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, es := range m {
		err = stmt.QueryRow(
			es.NumEspecific,
			es.Especific,
			es.NickUser,
			es.NameProyect,
		).Scan(&es.Especific)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	fmt.Println("El objetivo especifico fue creado exitosamente")
	return nil
}

// Update
func (f *PsqlEspecific) Update(m *model.Especific) error {
	stmt, err := f.db.Prepare(psqlUpdateEspecific)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		&m.ID,
		&m.NumEspecific,
		&m.Especific,
		&m.NickUser,
		&m.NameProyect,
	)
	if err != nil {
		return err
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAfected == 0 {
		return fmt.Errorf("no existe el objetivo especifico con el id: %d", m.ID)
	}

	fmt.Println("El objetivo especifico se actualizó exitosamente")
	return nil
}

// Delete
func (f *PsqlEspecific) Delete(id uint) error {
	stmt, err := f.db.Prepare(psqlDeleteEspecific)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	fmt.Println("El objetivo especifico se eliminó correctamente")
	return nil
}

func (p *PsqlEspecific) GetByNameProyect(nameProyect string) (model.Especifics, error) {
	stmt, err := p.db.Prepare(pslqGetbyNameProyect)
	if err != nil {
		return model.Especifics{}, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(nameProyect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ms := make(model.Especifics, 0)
	for rows.Next() {
		p, err := scanRowEspecifics(rows)
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

func (p *PsqlEspecific) GetAll() (model.Especifics, error) {
	stmt, err := p.db.Prepare(psqlGetAllEspecifics)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(model.Especifics, 0)
	for rows.Next() {
		p, err := scanRowEspecifics(rows)
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

func scanRowEspecifics(s scannerEspecific) (*model.Especific, error) {
	m := &model.Especific{}
	err := s.Scan(
		&m.ID,
		&m.NumEspecific,
		&m.Especific,
		&m.NickUser,
		&m.NameProyect,
	)
	if err != nil {
		return &model.Especific{}, err
	}
	return m, nil
}
