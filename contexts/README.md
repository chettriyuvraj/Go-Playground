# README

Exploring some basics of the context package with some sort-of believable use cases.


## Values

So this a simple wrapping of a value in a context and printing it - doesn't teach us much. 

TODO: The documentation mentions _lightweight_ values can be sent using this, have a look at an actual example and tag it here.


## Timeouts or Deadlines / Cancellations

Sort of a real-world situation: A request that must be completed within a certain amount of time. If request times out, server can free up its resources.

The general set of steps would always be along these lines:

- Create a context that times out within a specific duration.
- Trigger the job
- Monitor whichever happens first: expiry of context timeout or job completion and return response accordingly.

This would hold true for cancellations as well.

### Code

- The _TestHandler_ struct accepts a context and a _jobDuration_

- It will execute a job that takes _jobDuration_ to complete (+ monitor it), while also monitoring the parent context for completion.

- Whichever happens first determines its response: currently timeout is 3 seconds and job duration is 5 seconds so context will timeout before job completion.

- You can change these values in _main_ and run the program to see the diff in output.

