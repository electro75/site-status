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