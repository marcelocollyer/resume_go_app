# resume_go_app
Resume back-end application developed with GOlang and mongoDB.

LIVE DEMO:
http://www.marcelocollyer.com

This simple application has been developed to be used with a front-end project wrote using Ionic Framework (Angular).
You might want to setup the front-end later on to have the full experience.
Please check for more information here: https://github.com/marcelocollyer/resume_angular_app 

To setup this project, please check if you have all the following pre-requirements:

Pre-requirements:

* You MUST have GO installed.
* You MUST have mongoDB installed.

Then you can go forward with those steps:  

1) Clone this repo into a local folder
2) Run 'go get -u github.com/marcelocollyer/resume_go_app'
3) Run 'go build' command.
4) Have mongoDB running
5) Run './resume_go_app'
6) Execute an http POST request to local server.
   Something like: http://localhost:8000/resume
   Use content of the file /dao/db_scripts/marcelo_resume_script.sql as the POST request body.
  

Have fun!
