package nethgateapi

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NethGateService struct {
	db *gorm.DB // todo: what db are we actually using?
}

var _ NethGateService = NethGateServiceInterface{} //todo: complete the implementation

func NewNethGateService(db *gorm.DB) *NethGateService {
	return &NethGateService{db: db}
}

func (service *NethGateService) CreateUser(db *gorm.DB, newUser User) (*User, error) {
	// Generate new UUID for the user
	newUser.UserID = uuid.New()
	newUser.CreatedDate = time.Now()
	newUser.UpdatedDate = time.Now()

	// Insert newUser into the database
	result := db.Create(&newUser)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newUser, nil
}

// SearchUser retrieves a user from the database by login.
func (service *NethGateService) SearchUser(db *gorm.DB, login string) (*User, error) {
	var user User
	result := db.Where("login = ?", login).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// UpdateUser updates a user's information in the database.
func (service *NethGateService) UpdateUser(db *gorm.DB, userID uuid.UUID, updatedFields map[string]interface{}) (*User, error) {
	updatedFields["UpdatedDate"] = time.Now()
	result := db.Model(&User{}).Where("user_id = ?", userID).Updates(updatedFields)
	if result.Error != nil {
		return nil, result.Error
	}
	var user User
	db.First(&user, userID)
	return &user, nil
}

// DeleteUser removes a user from the database.
func (service *NethGateService) DeleteUser(db *gorm.DB, userID uuid.UUID) error {
	result := db.Delete(&User{}, userID)
	return result.Error
}
