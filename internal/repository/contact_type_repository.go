package repository

import (
	"active-gym-manager/internal/domain/model"
	"database/sql"
)

type ContactTypeRepository struct {
	connection *sql.DB
}

func NewContactTypeRepository(db *sql.DB) ContactTypeRepository {
	return ContactTypeRepository{
		connection: db,
	}
}

func (ctr *ContactTypeRepository) FindAll() ([]model.ContactType, error) {
	rows, err := ctr.connection.Query("SELECT * FROM contact_type")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	contactTypes := make([]model.ContactType, 0)

	for rows.Next() {
		var contactType model.ContactType
		if err := rows.Scan(&contactType.ID, &contactType.Name); err != nil {
			return nil, err
		}
		contactTypes = append(contactTypes, contactType)
	}

	return contactTypes, nil
}
