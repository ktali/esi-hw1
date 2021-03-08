# Task 2
***implementer(s):*** Kaarel Loide, Karl Vaba

***reviewer:*** Kert Tali

---

***Was the stated problem/task solved in an acceptable manner?***
<!-- (Would this count as a delivery for a potential client?) - if this is a no, 
the task is rejected and can be redone (this option can be taken four times in 
total - has to be discussed with instructors) -->

###### Certainly!


***What has been done well and why?***

* The project structure is good: functionalities reside in a separate `/cmd` subpackage, each command having its own appropriately named file.
* Expected JSON structures from the server are made into go `struct`s.
* The Cobra package is put into good use, as there is minimal boilerplate and each command file has a uniform structure.
* The CLI is verbose, notifying of any errors or faulty input, also providing with command help descriptions.


***What is not well implemented and why?***

* The `-h` parameter could use a better description of the required input, e.g. the help prompt `todo-cli add -h` could tell you to follow with the description of your todo, `complete` to input the corresponding id etc.
* For usability, I would've formatted the ToDo items as multiple lines, rather than a single line per item containing the ID, Content and Status.
* My IDE warns me that there are a couple of `Unhandled error`s -- those are all due from the json parser, when giving it some input which doesn't match the structure it should use. I think it's okay here as at that point we know we have got some valid response from the server anyways and we can expect it to not be something completely different.


***How easy is it to understand?***
<!--what is easy to understand/well documented - what is hard to understand/lacks documentation-->

The code is very simple and therefore needs no extensive documentation.
Nevertheless there are some helpful comments and documentation comments for some functions.


***Recommendations***
<!--Are there any recommendations to simplify some of this task that the reviewer would like to share?-->

No further recommendations on my end.


***Did the implementers use anything special that was notable for the reviewer and should eventually be shared with a bigger audience?***
<!--Note these specialties down with a positive remark to encourage sharing.-->

Maybe the usage of environment variables for the server's URL to easily incorporate usage in various environments (like the compose file).


***Anything notable, encouraging, or funny***

_No easter eggs found._
