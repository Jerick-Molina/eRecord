package security

import "fmt"

var roleQueue = []string{"Owner", "Admin", "Manager", "Developer"}

func RoleAuthorization(validRoles []string, userRole string) error {

	for i := 0; i < len(roleQueue); i++ {


		// Owner gets access to everything || otherwise what ever goes
		if roleQueue[i] == userRole || "Owner" == userRole{
			//Authorized
			return nil
		}
	}
	//Unauthorized
	return fmt.Errorf("Unauthorized  access")
}



func Authorization(token string, roles []string) (db.AccesstokenClaims,error) {

	claims, err :=  jwt.ReadTokender(token)
	if err != nil {
		return _, nil
	}

	role := fmt.Sprint("%s",claims["role"])

	if err != nil {
		return _, nil
	}

	return  
}