# Task: Add another command-line flag **-b** to count the number of bytes. Assumptions:

* update *count()* function to accept another param. **countBytes**
* when this input parameter is set to **true**, the function should count bytes
* check all methods for the type *bufio.Scanner* in **golang.org/pkg/bufio


## Task completed: How?

After some googling around, it appeared that the official docu was the most helpful. At first I tried to split the STDIN to **words**, **bytes** and then use len to change []byte return into integer. But I just only neeed to scan the file and then use the len to count the bytes in the STDIN. After that tests was easy.


