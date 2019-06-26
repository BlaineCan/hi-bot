const config = require("./config");
const Discord = require('discord.js');
const client = new Discord.Client();

let greetings = ["Hi", "Hi!", "Hi :slight_smile:", "Hi! :slight_smile:"]
let greet = greetings[Math.floor(Math.random() * greetings.length)];

//setInterval(sayHi(), 60000);

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
  if(message.channel.id !== config.chan) return;
  if(message.content.includes("hi") || message.content.includes("Hi")){
    return;
  } else {
    message.reply("this channel is used for saying Hi. :slight_smile:");
  }
})

client.login(config.token);
