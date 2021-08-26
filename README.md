# golang-channels-demo

Brief demo of golang channels I used to present on TDC Transformation 2021.

### **Running examples:**

`make write`

Shows how deadlocks happen if no routine reads from Channel after write.

`make read`

Shows how deadlocks happen if reading from Channel is never reached in the same routine that writes to it

`make routine`

Shows how deadlocks are resolved when different routines communicate through channels