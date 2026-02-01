package core

import (
	"strings"
)

type CommandInput struct {
	Command string
	Arg     string
}

type parseConfig struct {
	keepBackSlash bool
}

const QUOTE = '\''
const DOUBLE_QUOTE = '"'
const BACK_SLASH = '\\'
const DOLLAR_SIGN = '$'
const BACKTICK = '`'

func isQuote(char byte) bool {
	return char == QUOTE || char == DOUBLE_QUOTE
}

func parse(arg string, cfg *parseConfig) []string {
	var keepBackSlash = false
	if cfg != nil {
		keepBackSlash = cfg.keepBackSlash
	}

	quoteType := byte(QUOTE)
	insideQuotes := false
	groups := make([]string, 0)

	group := ""

	index := 0

	for index < len(arg) {
		if isQuote(arg[index]) && index+1 != len(arg) && arg[index+1] == arg[index] {
			index += 2
			continue
		}

		if arg[index] == BACK_SLASH && !keepBackSlash {
			if !insideQuotes && index+1 < len(arg) {
				group += string(arg[index+1])
				index += 2
				continue
			}

			if insideQuotes && quoteType == DOUBLE_QUOTE && index+1 < len(arg) {
				group += string(arg[index+1])
				index += 2
				continue
			}
		}

		if !insideQuotes && isQuote(arg[index]) {
			quoteType = arg[index]
			insideQuotes = true
		} else if insideQuotes && isQuote(arg[index]) && arg[index] == quoteType {
			if group != "" && !(index+1 < len(arg) && arg[index+1] != ' ') {
				groups = append(groups, group)
				group = ""
			}
			insideQuotes = false
		} else if insideQuotes || arg[index] != ' ' {
			group += string(arg[index])
		} else if !insideQuotes && arg[index] == ' ' && len(group) > 0 {
			groups = append(groups, group)
			group = ""
		}

		index += 1
	}

	if group != " " && group != "" {
		groups = append(groups, group)
	}

	if insideQuotes {
		// has only opening quote
		return []string{arg}
	}

	return groups
}

func ParseArg(arg string) []string {
	return parse(arg, nil)
}

func ParseInput(input string) *CommandInput {
	var command string
	var arg string

	if len(input) <= 0 {
		return nil
	}

	if isQuote(input[0]) {
		groups := parse(input, &parseConfig{
			keepBackSlash: true,
		})
		command = groups[0]
		idx := strings.LastIndex(input, command)
		arg = input[idx+len(command)+1:]
	} else {
		inputArray := strings.Split(input, " ")

		if len(inputArray) > 0 {
			command = inputArray[0]
			if command == "" {
				return nil
			}
			if len(inputArray) == 0 {
				return nil
			}
			if len(inputArray) > 1 {
				arg = strings.Join(inputArray[1:], " ")
			}
		}
	}

	return &CommandInput{
		command, arg,
	}
}
