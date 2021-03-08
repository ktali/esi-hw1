# Task 5
***implementer(s):*** Kert Tali

***reviewer:*** Kaarel Loide

---

***Was the stated problem/task solved in an acceptable manner?***

Yes the task was solved in an acceptable manner.


***What has been done well and why?***

The configuration files written for this task are easy to understand and do what is needed.
The choice of mystery image is great. **A whole OS running in the browser!** Can it get any better than this? :D


***What is not well implemented and why?***

Everything seems good to me implementation wise.
One thing to note tho. Not sure if this is an issue on my part or if anything can even be done about this, but while trying out the mystery image
an annoying popping sound kept happening because of some connection lost notification displaying in the corner of the screen 
every few seconds. This is not really part of the task, and the image itself seemed to operate correctly despite the notification so this is a minor issue at best.


***How easy is it to understand?***

As with the previous task, maybe a README file would be nice. While reviewing this task I started off by navigating
to the `task5` directory to figure out what is going on. At first, I saw the file `docker-compose.yml` so I assumed
I need to run `docker-compose up` to make this task work. My assumption was correct and that started the containers.
Now I didn't really know what to do next but then noticed that my terminal window was displaying the message

`osjs_1         | ✔ Server listening on http://localhost:8000`

That is what I tried next. I typed the url into my browser but got no response.
Then I looked at the task description and saw the `/api` and `/mystery` prefix requirements.
After some experimenting with different port numbers gotten from reading the config files and remembering some
things about nginx from the Systems Modelling project last semester I figured out that the correct urls were
`localhost/api` and `localhost/mystery`. All of this trouble could have been avoided with a simple README file.


***Recommendations***

I am not that familiar with setting up a proxy and this seemed more like a guide for me on how to do it in the future so no recommendations besides the documentation one for this task.


***Did the implementers use anything special that was notable for the reviewer and should eventually be shared with a bigger audience?***

This was all pretty new to me so don't know anything extra special to bring out here.


***Anything notable, encouraging, or funny***
### A WHOLE OS RUNNING IN THE BROWSER
and it is written in JavaScript. This is a very intriguing image choice!

