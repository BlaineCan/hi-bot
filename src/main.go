/* TODO: 
	1. Add function that greets user with command '!greet'
	2. Add function to greet specific user with command '!greet <user>'
 */

package main 

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"syscall"
	"os/signal"
)

var (
	token string
	BotID string
)

func main(){
	/* This is just a quick sanity check to make sure that a token is present. */
	if token == "" {
		fmt.Println("No token provided. Please provide a valid token to continue.")
		return
	}

	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println(err.Error())
		return 
	}
	err = discord.Open()
	if err != nil {
		fmt.Println(err.Error())
	}
	/* Hang out here until CTRL-C */
	fmt.Println("Bonfire lit. Press CRTL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}

