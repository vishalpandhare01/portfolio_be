package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vishalpandhare01/portfolio_be/initializer"
	"github.com/vishalpandhare01/portfolio_be/internal/model"
	"github.com/vishalpandhare01/portfolio_be/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

// create user data
func CreateUser(C *fiber.Ctx) error {
	var body *model.UserModel
	var userNameExist *model.UserModel
	if err := C.BodyParser(&body); err != nil {
		return C.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if body.Role == "" {
		return C.Status(400).JSON(fiber.Map{
			"message": "Role required",
		})
	}

	if body.Role == "admin" {
		if err := initializer.DB.Where("role = ?", body.Role).First(&body).Error; err == nil {
			return C.Status(400).JSON(fiber.Map{
				"message": "admin role already exist what are you doing bro 😒F",
			})
		}
	}

	if body.UserName == "" {
		return C.Status(400).JSON(fiber.Map{
			"message": "UserName required",
		})
	}

	if err := initializer.DB.Where("user_name = ?", body.UserName).First(&userNameExist).Error; err != nil {
		if err.Error() != "record not found" {
			return C.Status(500).JSON(fiber.Map{
				"message": "username Db error: " + err.Error(),
			})
		}
	}
	fmt.Println("userNameExist", userNameExist)
	if userNameExist.UserName != "" {
		return C.Status(400).JSON(fiber.Map{
			"message": "This user name not allow choose new  one",
		})
	}

	iSValidUserName := utils.ValidateUsername(body.UserName)
	if !iSValidUserName {
		return C.Status(400).JSON(fiber.Map{
			"message": "This username is not valid",
		})
	}

	if body.Password == "" {
		return C.Status(400).JSON(fiber.Map{
			"message": "Password required",
		})
	}

	if err := initializer.DB.Create(body).Error; err != nil {
		return C.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return C.Status(201).JSON(fiber.Map{
		"message": "Success",
		"data":    body,
	})

}

type LoginBody struct {
	UserName string
	Password string
}

// loginuser
func LoginUser(C *fiber.Ctx) error {
	var body LoginBody
	var data model.UserModel

	if err := C.BodyParser(&body); err != nil {
		return C.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := initializer.DB.Where("user_name = ?", body.UserName).First(&data).Error; err != nil {
		if err.Error() == "record not found" {
			return C.Status(404).JSON(fiber.Map{
				"message": "user dose not exist create acount",
			})
		}
		return C.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(body.Password)); err != nil {
		return C.Status(400).JSON(fiber.Map{
			"message": "password not match",
		})
	}

	token, err := utils.GenerateJwtToken(data.ID, data.Role)
	if err != nil {
		return C.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return C.Status(200).JSON(fiber.Map{
		"message": "Success",
		"data":    token,
	})
}

// get all users
func GetAllUsers(C *fiber.Ctx) error {
	var data []model.UserModel
	if err := initializer.DB.Find(&data).Error; err != nil {
		return C.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return C.Status(200).JSON(fiber.Map{
		"message": "Success",
		"data":    data,
	})
}

// create user profile
func CreateUserProfile(C *fiber.Ctx) error {
	var body model.UserProfile
	var checkExist model.UserProfile
	id := C.Locals("userId")
	if id == nil {
		return C.Status(400).JSON(fiber.Map{
			"message": "Login required",
		})
	}
	body.UserID = fmt.Sprint(id)

	if err := C.BodyParser(&body); err != nil {
		return C.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := initializer.DB.Where("user_id = ?", body.UserID).First(&checkExist).Error; err != nil {
		fmt.Println("error in check exist", err.Error())
	}
	fmt.Println("is here errror", id)

	if checkExist.UserID == id {
		return C.Status(400).JSON(fiber.Map{
			"message": "profile already exist delete it and create new if any problem we cant afford edit feature",
		})
	}

	if err := initializer.DB.Create(&body).Error; err != nil {
		return C.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return C.Status(201).JSON(fiber.Map{
		"message": "Success",
		"data":    body,
	})

}

// delete user profile by admin
func DeleteUserByAdmin(C *fiber.Ctx) error {
	var data *model.UserModel
	id := C.Params("id")

	if err := initializer.DB.Where("id = ?", id).Delete(&data).Error; err != nil {
		return C.Status(500).JSON(fiber.Map{
			"message": "User Profile: " + err.Error(),
		})
	}

	return C.Status(200).JSON(fiber.Map{
		"message": "Delete Successfully",
		"data":    data,
	})

}

// get login user profile
func GetUserProfileByUserName(C *fiber.Ctx) error {
	var data *model.UserProfile
	var user *model.UserModel
	userName := C.Params("userName")

	if err := initializer.DB.Where("user_name = ?", userName).First(&user).Error; err != nil {
		return C.Status(500).JSON(fiber.Map{
			"message": "User: " + err.Error(),
		})
	}

	if err := initializer.DB.Where("user_id = ?", user.ID).First(&data).Error; err != nil {
		return C.Status(500).JSON(fiber.Map{
			"message": "User Profile: " + err.Error(),
		})
	}

	return C.Status(200).JSON(fiber.Map{
		"message": "Success",
		"data":    data,
	})

}

func GetLogInUserProfile(C *fiber.Ctx) error {
	var data *model.UserModel
	id := C.Locals("userId")

	// Use Select to limit the columns retrieved
	if err := initializer.DB.Select("id, user_name, role").Where("id = ?", id).First(&data).Error; err != nil {
		return C.Status(500).JSON(fiber.Map{
			"message": "User: " + err.Error(),
		})
	}

	// Return the selected data
	return C.Status(200).JSON(fiber.Map{
		"message": "Success",
		"data":    data,
	})
}

// delete user profile by id
func DeleteUserProfileById(C *fiber.Ctx) error {
	var data *model.UserProfile
	id := C.Locals("userId")

	if id == nil {
		return C.Status(400).JSON(fiber.Map{
			"message": "Login required",
		})
	}

	if err := initializer.DB.Where("user_id = ?", id).Delete(&data).Error; err != nil {
		return C.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return C.Status(200).JSON(fiber.Map{
		"message": "Success",
		"data":    data,
	})

}

// create contact
func CreateContact(C *fiber.Ctx) error {
	var body model.Contacts
	id := C.Params("id")

	body.UserID = fmt.Sprint(id)

	if err := C.BodyParser(&body); err != nil {
		return C.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := initializer.DB.Create(&body).Error; err != nil {
		return C.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return C.Status(201).JSON(fiber.Map{
		"message": "Success",
		"data":    body,
	})

}

// get contact
func GetUserContacts(C *fiber.Ctx) error {
	var data *[]model.Contacts
	id := C.Locals("userId")

	fmt.Println("is here errror", id)
	if err := initializer.DB.Where("user_id = ?", id).Find(&data).Error; err != nil {
		return C.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return C.Status(200).JSON(fiber.Map{
		"message": "Success",
		"data":    data,
	})

}

// delete user contact
func DeleteUserContacts(C *fiber.Ctx) error {
	var data *model.Contacts
	userId := C.Locals("userId")
	id := C.Params("id")

	if userId == nil {
		return C.Status(400).JSON(fiber.Map{
			"message": "Login required",
		})
	}

	if err := initializer.DB.Where("user_id = ? and  id  =  ?", userId, id).Delete(&data).Error; err != nil {
		return C.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return C.Status(200).JSON(fiber.Map{
		"message": "Success",
		"data":    data,
	})

}
