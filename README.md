# golang-channels-demo

Brief demo of golang channels I used to present on TDC Transformation 2021.

#

### **Running examples:**

#

#### `make write`

Shows how deadlocks happen if no routine reads from Channel after write.

#

#### `make read`

Shows how deadlocks happen if reading from Channel is never reached in the same routine that writes to it

#

#### `make routine`

Shows how deadlocks are resolved when different routines communicate through channels

#

#### `make seq`

Runs sequential program.

Optional:

> `NUMBER_OF_REQS` : Number of requests to be performed

#

#### `make cur`

Runs concurrent program.

Optional:

> `NUMBER_OF_REQS` : Number of requests to be performed

#

#### `make time-seq`
#### `make time-cur`

This commands will measure time of sequential and concurrent executions, respectively. They also accept the same variable `NUMBER_OF_REQS`.