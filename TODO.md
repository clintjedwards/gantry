* Figure out how remote authentication would work for things like Amazon ECS
* Should be an easy way to run a series of commands that aid in deployment without actually typing them out
    * For instance to deploy docker we pull the new images, spin the environment down, and then spin the environment back up with new containers
    * This could all be solved by a series of aliases embedded into the shell we drop the user into (Not quite sure how to implement) or someway to embed custom commands
