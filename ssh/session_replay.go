package ssh

import (
	"errors"
	"strings"
)

const commandNotFound = "Command not found"

type replayMultiCommandSession struct {
	shellPrompt string
	system      string
	commands    *ReplayCommands
}

func (s *replayMultiCommandSession) Run(command string, timeoutMs int, terminators ...string) (string, error) {
	if !strings.HasSuffix(command, "\n") {
		command = command + "\n"
	}

	replay, ok := s.commands.Commands[command]
	if !ok {
		return commandNotFound, nil
	}
	if replay.Error != "" {
		return "", errors.New(replay.Error)
	}
	return s.commands.Next(command), nil
}

func (s *replayMultiCommandSession) ShellPrompt() string {
	return s.shellPrompt
}

func (s *replayMultiCommandSession) System() string {
	return s.system
}

func (s *replayMultiCommandSession) Close() {

}

func NewReplayMultiCommandSession(shellPrompt, system string, commands *ReplayCommands) MultiCommandSession {
	return &replayMultiCommandSession{
		shellPrompt: shellPrompt,
		system:      system,
		commands:    commands,
	}
}
