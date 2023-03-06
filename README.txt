The RequestID is just the order that the requests were made.

I unfortunately have not been able to implement concurrency yet, but hopefully
I'll be able to by the end of the week.

When you try to test the program with large.txt, many requests fail. The errors
are logged and those requests are then skipped. The final large_output.csv file 
just contains the results of the successful requests.

When you run plots.py, check the console for the average byte size and the 
two plots will appear sequentially (you have to kill the first plot to see the
second).