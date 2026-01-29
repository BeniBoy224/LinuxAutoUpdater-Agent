/* Steps to SUCCESS

 - Check if agent json file is located in /etc/linux-auto-updater/agent.json
 
 - If json config found, send heartbead to controller server

		- if this is the first time the server has seen this hostname and id. create a new UUID on api and send it to agent (along with token)
		- send basic info about server (hardware, software running, versions)


*/
