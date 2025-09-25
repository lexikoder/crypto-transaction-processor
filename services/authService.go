package service

import (
	"context"
	"crypto-transaction-processor/database"
	"crypto-transaction-processor/dto"
	"crypto-transaction-processor/models"
	"crypto-transaction-processor/utils"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OtpEntry struct {
	OTP       string
	ExpiresAt int64
	Verified  bool
}
type UserInfo struct {
    ID       uuid.UUID
    Email    string
    Role     models.Role
}


type AuthService struct {
	Db *database.Repository
}
var otpCache = make(map[string]OtpEntry)
var expiryMinutes = 10
var AccessTokenExpiryTime = 10   // in minutes
var RefreshTokenExpiryTime = 24   // in minutes

func (authService *AuthService) ReqOtpService(ctx context.Context,dto dto.BaseEmailDTO)(error){
	
    otp, err := utils.GenerateOTP(6)
    if err != nil {
    return utils.NewAppError("failed to generate otp", http.StatusInternalServerError)
    }

	otpCache[dto.Email] = OtpEntry{
		OTP:       otp,
		ExpiresAt: time.Now().Add(time.Minute *  time.Duration(expiryMinutes)).Unix(), // 10 minutes from now
		Verified:  false,
	}
  
	sendotperr := utils.SendOtp(dto.Email,otp, expiryMinutes)

	if sendotperr != nil {
    return utils.NewAppError("failed to send otp", http.StatusInternalServerError)
    }
	return nil
}



func (authService *AuthService) VerifyOtpService(ctx context.Context,dto dto.VerifyOtpDTO)(error){
	_, ok := otpCache[dto.Email]

	if !ok {
       return utils.NewAppError("otp not yet requested", http.StatusBadRequest)
	}

	if time.Now().Unix() > otpCache[dto.Email].ExpiresAt {
       return utils.NewAppError("OTP expired", http.StatusBadRequest)
	}
	if otpCache[dto.Email].Verified {
       return utils.NewAppError("OTP already verified", http.StatusBadRequest)
	}


   if (otpCache[dto.Email].OTP !=  dto.Otp ){
	   return utils.NewAppError("invalid otp", http.StatusBadRequest)
	}
	if (otpCache[dto.Email].OTP ==  dto.Otp ){
	   	otp := otpCache[dto.Email]  // copy the struct
        otp.Verified = true         // modify the copy
        otpCache[dto.Email] = otp  
       return nil 
	}
	return utils.NewAppError("something went wrong", http.StatusBadRequest)
}

func (authService *AuthService) SignupService(ctx context.Context,userdto dto.SignupRequestDTO)(error){
	
    var user models.User
	err := authService.Db.DB.WithContext(ctx).Where("email = ?", userdto.Email).First(&user).Error
		if err == nil {
			return utils.NewAppError("user already exist", http.StatusBadRequest)
		}

		_, ok := otpCache[userdto.Email]

	    if !ok || !otpCache[userdto.Email].Verified {
           return utils.NewAppError("Email not verified with OTP", http.StatusBadRequest)
	    }

		hashedpassword, err := utils.HashPassword(userdto.Password)
		if err != nil {
			log.Println("Password hashing error:", err) // internal log
			return utils.NewAppError("server error", http.StatusInternalServerError)
		}
		// VerificationType := models.VerificationType(userdto.VerificationType)
		user = models.User{
			ID:       uuid.New(),
			Email:    &userdto.Email,
			Age:      &userdto.Age,
		    Phone:    &userdto.Phone,
			Password: &hashedpassword,
			Role: models.Role(userdto.Role),
			Verified: userdto.Verified,
	        // VerificationType: &VerificationType,
	        VerificationId:&userdto.VerificationId,   

		}

		if userdto.VerificationType != "" {
           vt := models.VerificationType(userdto.VerificationType)
           user.VerificationType = &vt
        }

		err = authService.Db.DB.WithContext(ctx).Create(&user).Error

		if err != nil {
			log.Println("Error creating user:", err)
			return utils.NewAppError("error creating user", http.StatusInternalServerError)
			// c.JSON(http.StatusBadRequest, gin.H{"error": "error creating user"})
			// return
		}
	return nil
}

func (authService *AuthService) LoginService(ctx context.Context,userdto dto.LoginRequestDTO)(string,string,error){
	
    var user models.User
	var refreshTokendata models.RefreshToken
	err := authService.Db.DB.WithContext(ctx).Where("email = ?", userdto.Email).First(&user).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
        return "", "",  utils.NewAppError("User not found", http.StatusNotFound)
       }

	 _ ,  verifyerr := utils.VerifyPassword(*user.Password,userdto.Password)
     
	 if verifyerr != nil  {
			return "","", utils.NewAppError("Invalid credentials", http.StatusNotFound)
	 }

    userInfo := UserInfo{
    ID:       user.ID,
    Email: *user.Email,
    Role:  user.Role,
    }

  accessToken,JwtSignerr := utils.JwtSignMinutes(userInfo,AccessTokenExpiryTime)   // expires in 10 minutes
  if JwtSignerr != nil {
		return "","", utils.NewAppError("something went wrong", 500)
  } 

  refreshToken,JwtSignerr := utils.JwtSignHour(userInfo,RefreshTokenExpiryTime)   // expires in 10 hours
  if JwtSignerr != nil {
		return "","", utils.NewAppError("something went wrong", 500)
  } 

   expiry := time.Now().Add(time.Hour * time.Duration(RefreshTokenExpiryTime)).Unix()
   refreshTokendata = models.RefreshToken{
	ID:       uuid.New(),
	Token:    &refreshToken,
	ExpiresAt : &expiry,
	UserID: (*uuid.UUID)(&user.ID),
    }
 
  err = authService.Db.DB.WithContext(ctx).Create(&refreshTokendata).Error

		if err != nil {
			return "","", utils.NewAppError("something went wrong", http.StatusInternalServerError)
		}


  return accessToken , refreshToken, nil
}