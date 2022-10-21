package security

import "fmt"

var roleQueue = []string{"Owner", "Admin", "Manager", "Developer"}

func RoleAuthorization(validRoles []string, userRole string) error {

	for i := 0; i < len(roleQueue); i++ {

		if roleQueue[i] == userRole {
			//Authorized
			return nil
		}
	}
	//Unauthorized
	return fmt.Errorf("Unauthorized  access")
}
