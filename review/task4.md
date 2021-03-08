# Task 4
***implementer(s):*** Kert Tali

***reviewer:*** Kaarel Loide

---

***Was the stated problem/task solved in an acceptable manner?***

Yes the task was done well.


***What has been done well and why?***

Everything asked to be done has been completed. `The docker-compose.yml` is simple but completes all the requirements.
The implementer of the task also provided instructions on how to update the images is the cloud registry to the teammates on discord, by following these
the registry seems to work as expected. The `docker-compose.yml` works as expected and uses the registry to pull needed images.
There are also some comments in the `rest/Dockerfile` and `todo-cli/Dockerfile` that explain what techniques were used to optimize the size of the Docker images.


***What is not well implemented and why?***

Implementation wise everything seems great.



***How easy is it to understand?***

Maybe there could be some instructions in the project root `README.md` on how to upload images to the cloud registry,
and the changes needed to the `/etc/docker/daemon.json` file for the registry to work locally.
These instructions were given in a discord message but would be nice to have here as well.


***Recommendations***

Nothing else then the documentation recommendation.


***Did the implementers use anything special that was notable for the reviewer and should eventually be shared with a bigger audience?***

For me it was new information that using a single RUN directive to install all packages in the Dockerfile reduces the size of the docker image so maybe this is not known to others as well and is worth sharing.


***Anything notable, encouraging, or funny***

The cloud registry part of this homework was the core supporting all other things because the images were all stored there eventually. 
So it was notable that the implementer of this task
was always quick to help with any questions/problems with integrating other parts of the homework with the cloud setup.
