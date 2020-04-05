# Learn It



### Setup

* Clone this repo
* [Get a credentials.json from Google for the Google Calendar API](https://developers.google.com/calendar/quickstart/go)

### Running

```bash
$ cat input.dict # You should populate this yourself
것:A thing or  an object
하다:To do
있다:To be
수:way, method, Number
하다:To do
$ go build *.go
$ ./main input.dict
```

### References

Most of the code in this repo is taken from [Google's Go
Quickstart](https://developers.google.com/calendar/quickstart/go).