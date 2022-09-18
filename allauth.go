package golang_gorm_allauth

import (
	"gorm.io/gorm"
	"log"
)

// DefaultAccountAuthenticationMethod /* The default authentication method which is email*/
const DefaultAccountAuthenticationMethod = "EMAIL"

// AllauthConfig /* AllauthConfig enables setting the default variables for the user models
type AllauthConfig struct {
	UsernameRequired            bool
	AccountAuthenticationMethod string
	Database                    *gorm.DB
}

func (a *AllauthConfig) GetAllauthConfig() AllauthConfig {
	/* Get the default configurations set in the config*/

	if a.AccountAuthenticationMethod == "" {
		/* set the authentication method to the default */
		a.AccountAuthenticationMethod = DefaultAccountAuthenticationMethod
	}
	return AllauthConfig{
		UsernameRequired:            a.UsernameRequired,
		AccountAuthenticationMethod: a.AccountAuthenticationMethod,
		Database:                    a.Database,
	}
}

func (a *AllauthConfig) SetUsernameRequired(required bool) {
	a.UsernameRequired = required
}

func (a *AllauthConfig) SetDatabase(Database *gorm.DB) {
	/* set the database instance gotten from the createdb or opendb from gorm*/
	a.Database = Database
}

func (a *AllauthConfig) SetAccountAuthenticationMethod(AccountAuthenticationMethod string) {
	/* set the AccountAuthenticationMethod which is either email or username in capital letters*/
	if AccountAuthenticationMethod == "USERNAME" || AccountAuthenticationMethod == "EMAIL" {
		log.Panicln("AccountAuthenticationMethod must be USERNAME or EMAIL, not ", AccountAuthenticationMethod)
	}
	a.AccountAuthenticationMethod = AccountAuthenticationMethod
}
