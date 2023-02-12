# stdslack

send `STDIN` to slack as a file(uses [files.upload](https://api.slack.com/methods/files.upload) endpoint).

## Usage

1. set your slack bot/user token which have [files:write](https://api.slack.com/scopes/files:write) permission and have joined target channel to environ `SLACK_TOKEN`.
2. `cat sample.txt | ./stdslack --channels <CHANNEL_ID1>,<CHANNEL_ID2> (add some options)`

### When some error occurs

prints error message to `STDERR` and dump contents received from `STDIN` to `STDOUT`.

## License

MIT

## Author

walkure at 3pf.jp
