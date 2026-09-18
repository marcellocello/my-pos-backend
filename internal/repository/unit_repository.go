package repository

import (
	"database/sql"
	"mypos-backend/internal/model"
)

type UnitRepository struct {
	db *sql.DB
}

func NewUnitRepository(db *sql.DB) *UnitRepository {
	return &UnitRepository{db: db}
}

func (r *UnitRepository) InsertUnit(unit *model.CreateUnitRequest) error {
	query := `INSERT INTO units (name, description) VALUES (?, ?)`
	_, err := r.db.Exec(query, unit.Name, unit.Description)
	return err
}

func (r *UnitRepository) UpdateUnit(id int, unit *model.UpdateUnitRequest) error {
	query := `UPDATE units SET name = ?, description = ? WHERE id = ?`
	_, err := r.db.Exec(query, unit.Name, unit.Description, id)
	return err
}

func (r *UnitRepository) DeleteUnit(id int) error {
	query := `DELETE FROM units WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *UnitRepository) GetUnitByID(id int) (*model.Unit, error) {
	var unit model.Unit
	query := `SELECT id, name, description FROM units WHERE id = ?`
	err := r.db.QueryRow(query, id).Scan(&unit.ID, &unit.Name, &unit.Description)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &unit, err
}

func (r *UnitRepository) GetAllUnits() ([]model.Unit, error) {
	var unit []model.Unit

	query := `SELECT id, name, description FROM units`
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u model.Unit
		err := rows.Scan(&u.ID, &u.Name, &u.Description)
		if err != nil {
			return nil, err
		}
		unit = append(unit, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return unit, err
}
