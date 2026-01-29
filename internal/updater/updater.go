/*
Main heart of the process

 - check that there are usefull updates to be made
 - make sure the server is safe to update (disk space being free, cpu not too high, not crashing appliations of server itself)
 - execute dummy commands to make sure update can happen
 - execute the update commands
 - post update, check updates were successfull
 - tell the server what happend (success, error, where it errored)
 - If an error happens during an update STOP and alert the controller
 - If reboot requried let the controller know and see what to do
*/