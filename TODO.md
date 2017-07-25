* Figure out how remote authentication would work for things like Amazon ECS
* Should be an easy way to run a series of commands that aid in deployment without actually typing them out
    * For instance to deploy docker we pull the new images, spin the environment down, and then spin the environment back up with new containers
    * This could all be solved by a series of aliases embedded into the shell we drop the user into (Not quite sure how to implement) or someway to embed custom commands
* Check ssh connection is actually valid by printing docker version or something silent?
    * Currently this fails on first try of a docker command but acts normally until you try
* Remember to document(both in app and not) loading the key first errors


