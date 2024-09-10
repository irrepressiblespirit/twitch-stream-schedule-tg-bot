Telegram bot (back-end side) which show you translation schedule of your favorite twitch streamer by his name.

 # Description
Application consist of two parts: 
1. Web server which handle request from twitch API (get parameters from request and redirect new request with this parameters to telegram bot).
2. Back-end code of telegram bot 

 # How to run
All commands describe in Makefile. Befor application run put values in config file. The web server and telegram bot must work in parallel.