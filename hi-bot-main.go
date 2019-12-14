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

	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	
	err = discord.Open()
	if err != nil {
		fmt.Println(err.Error())
	}
	
	fmt.Println("Bonfire lit. Press CRTL-C to exit.")
	
	<-make(chan struct{})
	return 
}




