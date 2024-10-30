package services

import (
	"area/models"
	"area/storage"
	"fmt"
)

func runApplet(applet models.Applet) {
	var response []string
	print("Running applet\n")
	print("\tENTRY : ", applet.IfType, "PARAMETERS : ", applet.If, "\n")
	print("\tREACTION : ", applet.ThatType, "PARAMETERS : ", applet.That, "\n\n")

	if action, exists := actions[applet.IfType]; exists {
		response = action(applet.ID_User.Hex(), applet.If)
		fmt.Println("Action succes :", response)
	} else {
		fmt.Println("Action non trouvée :", applet.IfType)
	}
}

func RunApplets() {
	var applets []models.Applet
	var users []models.User
	var applets_user []models.Applet

	users = storage.GetAllUsers()
	print("users collected : ", len(users), "\n")

	for _, user := range users {
		applets_user = storage.GetApplets(user.ID)
		print("applets_user collected : ", len(applets_user), "\n")
		applets = append(applets, applets_user...)
	}
	fmt.Println("applets collected : ", len(applets))
	for _, applet := range applets {
		runApplet(applet)
	}

}
