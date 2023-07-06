# Stim-Core

This service is the primary means for communication between the web platform ([Stim-Web](https://github.com/danielvolchek/stim-web)) and every other service 

All microservices communicate back and forth between Stim-Core as their dispatcher

Stim-Web communicates with Stim-Core

Stim-Core communicates with Stim-Web
Stim-Core communicates with DB
Stim-Core communicates with Email

All microservices need a secret server key generated by Stim-Core (stored on request by Stim-DB)
Stim-Core needs a key generated by each service (except Stim-Web)