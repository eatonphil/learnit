# Learn It

A simple script for generating one Google Calendar event per day from
each line of a local file.

### Running

[Get a credentials.json from Google for the Google Calendar
API.](https://developers.google.com/calendar/quickstart/go) Place it
in this directory.

```bash
$ cat examples/input.dict
것:A thing or  an object
하다:To do
있다:To be
수:way, method, Number
하다:To do
...
$ go build *.go
$ ./main examples/input.dict
... events are created
```

### References

Most of the code in this repo is taken from [Google's Go
Quickstart](https://developers.google.com/calendar/quickstart/go).