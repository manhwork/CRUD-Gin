package services

import (
	"CRUD_Gin/database"
	"CRUD_Gin/models"
	"log"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

const (
	FIND_ALL_USERS = `SELECT id, cccd, full_name, date_of_birth, gender, nationality, place_of_birth, 
              issue_date, expiry_date, old_cccd, ethnicity, religion, personal_id, 
              permanent_addr, father_name, mother_name, spouse_name
              FROM users `

	FIND_USER_BY_CCCD = `SELECT id, cccd, full_name, date_of_birth, gender, nationality, place_of_birth, 
              issue_date, expiry_date, old_cccd, ethnicity, religion, personal_id, 
              permanent_addr, father_name, mother_name, spouse_name
              FROM users  WHERE cccd = ?`

	CREATE_USER = `INSERT INTO users (cccd, full_name, date_of_birth, gender, nationality, place_of_birth, 
			  issue_date, expiry_date, old_cccd, ethnicity, religion, personal_id, 
			  permanent_addr, father_name, mother_name, spouse_name) 
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
)

func (us *UserService) FindAllUsers() ([]models.User, error) {
	db := database.DBConn()
	defer db.Close()

	row, err := db.Query(FIND_ALL_USERS)

	if err != nil {
		log.Println("Query user error")
		return nil, err
	}

	var users []models.User

	for row.Next() {
		var user models.User
		err := row.Scan(&user.ID, &user.CCCD, &user.FullName, &user.DateOfBirth, &user.Gender, &user.Nationality,
			&user.PlaceOfBirth, &user.IssueDate, &user.ExpiryDate, &user.OldCCCD, &user.Ethnicity,
			&user.Religion, &user.PersonalID, &user.PermanentAddr, &user.FatherName, &user.MotherName, &user.SpouseName)

		if err != nil {
			log.Fatalln("Scan user failed")
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (us *UserService) FindUserById(cccd string) (*models.User, error) {
	db := database.DBConn()
	defer db.Close()

	row := db.QueryRow(FIND_USER_BY_CCCD, cccd)

	var user models.User

	errScan := row.Scan(&user.ID, &user.CCCD, &user.FullName, &user.DateOfBirth, &user.Gender, &user.Nationality,
		&user.PlaceOfBirth, &user.IssueDate, &user.ExpiryDate, &user.OldCCCD, &user.Ethnicity,
		&user.Religion, &user.PersonalID, &user.PermanentAddr, &user.FatherName, &user.MotherName, &user.SpouseName)

	if errScan != nil {
		log.Printf("error scan user: %v", errScan)

		return nil, errScan
	}

	return &user, nil
}

func (us *UserService) Create(user *models.User) error {
	db := database.DBConn()
	defer db.Close()

	stmt, err := db.Prepare(CREATE_USER)
	if err != nil {
		log.Println("Prepare statement error:", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.CCCD, user.FullName, user.DateOfBirth, user.Gender, user.Nationality,
		user.PlaceOfBirth, user.IssueDate, user.ExpiryDate, user.OldCCCD, user.Ethnicity,
		user.Religion, user.PersonalID, user.PermanentAddr, user.FatherName, user.MotherName, user.SpouseName)

	if err != nil {
		log.Println("Execute statement error:", err)
		return err
	}

	return nil
}
