## Welcome!
This is a GoLang command line application I created which takes a URL or a
or a file containing multiple URLs as arguments. The program makes HTTP get requests and saves the following response information:

- Request ID
- URL Queried
- Query Response Status / Code [HTTP response codes]
- Start Request Time
- End Request Time
- Total Size of Response bytes


## Linux Terminal Commands

### `./getinfo --url <url name>`

This prints a response to the terminal in the following format:

`{<URL queried> <HTTP Status Code> <start request time (nanoseconds)> <end request time (nanoseconds)> <total size of response bytes> <RequestID> <Error>}`

### `./getinfo --file <filename.txt>`

For a file that contains a list of URLs, this logs the same response as above into a CSV file for 
each request to be made in the file.

You can see the results of running `small.txt` and `large.txt` saved in `small_output.csv` and 
`large_output.csv` respectively.

(Make sure if you want to test different input files that you change the output file in line 64 of `getinfo.go`. And make sure to build it again after modifying.)

DISCLAIMER: 
When you try to test the program with large.txt, many requests fail. The errors
are logged and those requests are then skipped. The final large_output.csv file 
just contains the results of the successful requests.

## Python Script

I also made a python script called `plots.py` that can read one of the csv files and

1. print out the average size of response bytes
2. Create a line plot of each request and the time taken to get the response
3. Create a cumulative distribution function (CDF) of the response times and find out the 
[50, 70, 75, 80, 85, 90, 95, 99] percentiles of the response times.

When you run plots.py, check the console for the average byte size and the 
two plots will appear sequentially (you have to kill the first plot to see the
second).

## Concurrency
The last step of this project was to implement `getinfo` using concurrency.

The `RequestID` is the order that URL was in the original text files. But the order they are in the csv file is the order that the request completed, so the numbers will appear to be all over the place.