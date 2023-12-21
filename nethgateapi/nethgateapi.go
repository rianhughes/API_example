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

func (service *NethGateService) CreateUser(newUser User) (*User, error) {
	newUser.UserID = uuid.New()
	newUser.CreatedDate = time.Now()
	newUser.UpdatedDate = time.Now()

	// Insert newUser into the database.
	result := service.db.Create(&newUser)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newUser, nil
}

// SearchUser retrieves a user from the database by login.
func (service *NethGateService) SearchUser(login string) (*User, error) {
	var user User
	result := service.db.Where("login = ?", login).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// UpdateUser updates a user's information in the database.
func (service *NethGateService) UpdateUser(userID uuid.UUID, updatedFields map[string]interface{}) (*User, error) {
	updatedFields["UpdatedDate"] = time.Now()
	result := service.db.Model(&User{}).Where("user_id = ?", userID).Updates(updatedFields)
	if result.Error != nil {
		return nil, result.Error
	}
	var user User
	service.db.First(&user, userID)
	return &user, nil
}

// DeleteUser removes a user from the database.
func (service *NethGateService) DeleteUser(userID uuid.UUID) error {
	result := service.db.Delete(&User{}, userID)
	return result.Error
}
