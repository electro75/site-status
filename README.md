## Site checker App

- In this small application we learn about concurrency in golang, why it is important and how it is implemented.

- This app makes a simple request to a few well known websites and checks whether they are running or not

### Initial implementation (Without channels/Routines)

- problems: 
    - Time is wasted when the program is waiting for a request to be fulfilled.
    - Everytime we need to check the site status again, we have to check it for all the sites.

### Routines

- Every Go program runs on a main **Routine**
- Whenever a routine encounters a blocking call (eg. http request) it stops and cannot do anything else.
- This is when another routine can be launched to handle different part of the program.
- __Concurrency is not parallelism__ : in a single core CPU, even goroutines are executed in a somewhat serial order.
- only making a new routine does not solve the problem: once new go routines are created, the main routine exits before the child routines are done with execution due to lack of communication between routines.
- To resolve the above issue, __Channels__ are implemented.

### Channels
- They are used to communicate in between routines
- Channels are typed. The type of data shared through a channel **must** be the same as the type of the channel.
- Channel calls are blocking operations.
- A channel should be called everytime a go routine is done.