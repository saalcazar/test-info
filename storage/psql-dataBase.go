package storage

import (
	"database/sql"
	"fmt"

	"github.com/saalcazar/ceadlbk-info/model"
)

type scannerDataBase interface {
	Scan(dest ...any) error
}

const (
	psqlCreateParticipant = `SELECT create_data_base_participants ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	psqlUpdateParticipant = `SELECT update_data_base_participants ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
	psqlDeleteParticipant = `SELECT delete_data_base_participants ($1)`
	psqlGetAllParticipant = `SELECT id_participant, name_participant, gender, age, organization, phone, type_participant, name_proyect, cod_founder, activity, nick_user, municipality, type_organization FROM data_base_participants`
)

type PsqlDataBase struct {
	db *sql.DB
}

func NewPsqlDataBase(db *sql.DB) *PsqlDataBase {
	return &PsqlDataBase{db}
}

// Create
func (f *PsqlDataBase) Create(m model.DataBases) error {
	tx, err := f.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	stmt, err := f.db.Prepare(psqlCreateParticipant)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, es := range m {
		err = stmt.QueryRow(
			es.NameParticipant,
			es.Gender,
			es.Age,
			es.Organization,
			es.Phone,
			es.TypeParticipant,
			es.NameProyect,
			es.CodFounder,
			es.Activity,
			es.NickUser,
			es.Municipality,
			es.TypeOrganization,
		).Scan(&es.NameParticipant)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	fmt.Println("Los participantes se ingresaron a la base de datos")
	return nil
}

// Update
func (f *PsqlDataBase) Update(m *model.DataBase) error {
	stmt, err := f.db.Prepare(psqlUpdateParticipant)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		&m.ID,
		&m.NameParticipant,
		&m.Gender,
		&m.Age,
		&m.Organization,
		&m.Phone,
		&m.TypeParticipant,
		&m.NameProyect,
		&m.CodFounder,
		&m.Activity,
		&m.NickUser,
		&m.Municipality,
		&m.TypeOrganization,
	)
	if err != nil {
		return err
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAfected == 0 {
		return fmt.Errorf("no existe el perticipante con el id: %d", m.ID)
	}

	fmt.Println("El item se actualizo correctamente")
	return nil
}

func (f *PsqlDataBase) Delete(id uint) error {
	stmt, err := f.db.Prepare(psqlDeleteParticipant)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	fmt.Println("El item se elimino correctamente")
	return nil
}

// GETALL
func (f *PsqlDataBase) GetAll() (model.DataBases, error) {
	stmt, err := f.db.Prepare(psqlGetAllParticipant)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(model.DataBases, 0)
	for rows.Next() {
		p, err := scannerRowDataBase(rows)
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

func scannerRowDataBase(s scannerDataBase) (*model.DataBase, error) {
	m := &model.DataBase{}
	err := s.Scan(
		&m.ID,
		&m.NameParticipant,
		&m.Gender,
		&m.Age,
		&m.Organization,
		&m.Phone,
		&m.TypeParticipant,
		&m.NameProyect,
		&m.CodFounder,
		&m.Activity,
		&m.NickUser,
		&m.Municipality,
		&m.TypeOrganization,
	)
	if err != nil {
		return &model.DataBase{}, err
	}
	return m, nil
}
