package config

import (
	"gorm.io/gorm"
	"log"
)

// AllauthConfig /* AllauthConfig enables setting the default variables for the user models
type AllauthConfig struct {
	UsernameRequired            bool
	AccountAuthenticationMethod string
	Database                    *gorm.DB
	AuthorizationMethod         string
	UserModel                   struct{}
}

func (a *AllauthConfig) GetAllauthConfig() AllauthConfig {
	/* Get the default configurations set in the config*/

	if a.AccountAuthenticationMethod == "" {
		/* set the authentication method to the default */
		a.AccountAuthenticationMethod = DefaultAccountAuthenticationMethod
	}
	if a.AuthorizationMethod == "" {
		/* set the authorization method to the default */
		a.AuthorizationMethod = DefaultAuthorizationMethod
	}
	/* Return all the configurations which was set*/
	return AllauthConfig{
		UsernameRequired:            a.UsernameRequired,
		AccountAuthenticationMethod: a.AccountAuthenticationMethod,
		AuthorizationMethod:         a.AuthorizationMethod,
		Database:                    a.Database,
		UserModel:                   a.UserModel,
	}
}

func (a *AllauthConfig) SetUsernameRequired(required bool) {
	/*Set if the username will be required when registering a user*/
	a.UsernameRequired = required
}

func (a *AllauthConfig) SetDatabase(Database *gorm.DB) {
	/* set the database instance gotten from the createdb or opendb from gorm in the individual project to be used*/
	a.Database = Database
}

func (a *AllauthConfig) SetAccountAuthenticationMethod(AccountAuthenticationMethod string) {
	/* set the AccountAuthenticationMethod which is either email or username in capital letters*/
	if AccountAuthenticationMethod == "USERNAME" || AccountAuthenticationMethod == "EMAIL" {
		log.Panicln("AccountAuthenticationMethod must be USERNAME or EMAIL, not ", AccountAuthenticationMethod)
	}
	a.AccountAuthenticationMethod = AccountAuthenticationMethod
}

func (a *AllauthConfig) SetAuthorizationMethod(AuthorizationMethod string) {
	/* The set the authorization method  which is Either JWT or TOKEN*/
	if AuthorizationMethod == "JWT" || AuthorizationMethod == "TOKEN" {
		log.Panicln("AuthorizationMethod must be JWT or TOKEN, not ", AuthorizationMethod)
	}
	a.AuthorizationMethod = AuthorizationMethod
}

func (a *AllauthConfig) SetUserModel(UserModel *struct{}) {
	/* This set the userModel if he has other values he would like to set on the user models*/
	if UserModel != nil {
		a.UserModel = *UserModel
	}
}
