/* TODO: 
	1. Add function that greets user with command '!greet'
	2. Add function to greet specific user with command '!greet <user>'
 */

package main 

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var (
	token string
	BotID string
)


func main(){
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	
	u, err := dg.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}
	
	BotID = u.ID
	
	err = dg.Open()
	if err != nil {
		fmt.Println(err.Error())
	}
	
	fmt.Println("Bonfire lit. Press CRTL-C to exit.")
	
	<-make(chan struct{})
	return 
}


