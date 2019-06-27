const config = require("./config");
const Discord = require('discord.js');
const client = new Discord.Client();

client.on("ready", () =>{
  client.user.setPresence({
    game: {
      name: "for Hi\'s",
      type: 'WATCHING'
    },
    status: 'online'
  })
})

client.on("message", message => {
  if(message.author.bot) return;
  if(message.channel.id !== config.chan.hi) return;
  if(message.content.includes("hi ") || message.content.includes("Hi ") || message.content.includes("hi!") || message.content.includes("Hi!") || message.content.includes("hi.") || message.content.includes("Hi.") || message.content.length < 3){
    return;
  } else {
    console.log(`${message.author.tag} said ${message.content}. This is not allowed.`)
    message.delete();
    message.reply("this channel is used for saying Hi. :slight_smile:");
  }
})

client.login(config.token);
