package repositories

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/models"
)

type UserEducationRepository struct {
	DB *sql.DB
}

func (r *UserEducationRepository) Create(edu *models.UserEducation) error {
	query := `
        INSERT INTO user_education (
            id, user_id, degree, institution_name,
            field_of_study, grade, year
        ) VALUES (
            $1, $2, $3, $4,
            $5, $6, $7
        )
        RETURNING created_at, updated_at
    `

	edu.ID = uuid.New()
	return r.DB.QueryRow(
		query,
		edu.ID, edu.UserID, edu.Degree, edu.InstitutionName,
		edu.FieldOfStudy, edu.Grade, edu.Year,
	).Scan(&edu.CreatedAt, &edu.UpdatedAt)
}

func (r *UserEducationRepository) GetByUserID(userID uuid.UUID) ([]*models.UserEducation, error) {
	query := `SELECT id, user_id, degree, institution_name,
                     field_of_study, grade, year,
                     created_at, updated_at
              FROM user_education WHERE user_id = $1`

	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var educations []*models.UserEducation
	for rows.Next() {
		var edu models.UserEducation
		err := rows.Scan(
			&edu.ID, &edu.UserID, &edu.Degree, &edu.InstitutionName,
			&edu.FieldOfStudy, &edu.Grade, &edu.Year,
			&edu.CreatedAt, &edu.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		educations = append(educations, &edu)
	}

	return educations, nil
}

func (r *UserEducationRepository) Update(edu *models.UserEducation) error {
	query := `
        UPDATE user_education
        SET degree = $1,
            institution_name = $2,
            field_of_study = $3,
            grade = $4,
            year = $5,
            updated_at = NOW()
        WHERE id = $6
    `
	_, err := r.DB.Exec(
		query,
		edu.Degree, edu.InstitutionName, edu.FieldOfStudy,
		edu.Grade, edu.Year, edu.ID,
	)
	return err
}

func (r *UserEducationRepository) Delete(eduID uuid.UUID) error {
	query := `DELETE FROM user_education WHERE id = $1`
	_, err := r.DB.Exec(query, eduID)
	return err
}
