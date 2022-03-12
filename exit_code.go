package main

type ExitCode int

const (
	Success int = iota
	MissingArgument
	ParseError
	NoSuchTweet
)

func (e ExitCode) String() string {
	return [...]string{
		"Success",
		"Missing argument for the tweet to get",
		"Unable to parse tweet URL or ID",
		"Unable to find the tweet",
	}[e]
}
