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
		if len(response) == 0 {
			fmt.Println("Action failed")
			return
		}
		var notExist bool = storage.StoreAndCheckResponse(applet.ID, response, applet.If)
		if notExist {
			fmt.Println("Action succes :", response)
			if reAction, exists := reActions[applet.ThatType]; exists {
				reAction(applet.ID_User.Hex(), applet.That, response)
			} else {
				fmt.Println("ReAction non trouvée :", applet.ThatType)
			}
		} else {
			fmt.Println("Action already have been executed")
		}
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
