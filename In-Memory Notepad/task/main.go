package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Inicialização dos Componentes com Dependências
	reader := bufio.NewReader(os.Stdin)
	receiver := NewInputReceiver(reader)
	parser := &CommandParser{}
	executor := newExecutor()

	// Criação da Instância de Notepad
	notepad := &Notepad{
		receiver: receiver,
		parser:   parser,
		executor: executor,
	}

	// Loop Principal para Receber e Processar Comandos
	receiveMax()
	fmt.Printf("Enter a command and data:")
	for {

		notepad.Init()
		fmt.Printf("Enter a command and data:")
	}
}

var maxNotes int

func receiveMax() {
	for maxNotes < 1 {
		fmt.Println("Enter the maximum number of notes:")
		fmt.Scan(&maxNotes)
	}

}

// Notepad struct e seu método
type Notepad struct {
	receiver Receiver
	parser   Parser
	executor Executor
}

func (np *Notepad) Init() error {

	input, err := np.receiver.Receive()

	if err != nil {
		return fmt.Errorf("erro ao receber entrada: %v", err)
	}
	cmd, err := np.parser.Parse(input)
	if err != nil {
		return fmt.Errorf("[ERROR] %v", err)
	}

	err = np.executor.Execute(cmd)
	if err != nil {
		return fmt.Errorf("erro ao executar comando: %v", err)
	}
	return nil
}

// Receiver interface e InputReceiver struct
type Receiver interface {
	Receive() (string, error)
}

type InputReceiver struct {
	reader   *bufio.Reader
	maxNotes int
}

func NewInputReceiver(r *bufio.Reader) *InputReceiver {

	return &InputReceiver{reader: r}
}

func (ir *InputReceiver) Receive() (string, error) {

	input, err := ir.reader.ReadString('\n')

	if err != nil {
		return "Invalid Input", err
	}
	return strings.TrimSpace(input), nil
}

// Parser interface e CommandParser struct
type Parser interface {
	Parse(input string) (Command, error)
}

type CommandParser struct{}

func (cp *CommandParser) Parse(input string) (Command, error) {
	parts := strings.Fields(input)

	switch {
	case len(parts) == 0:
		fmt.Println("[Error] Missing note argument")
		return Command{}, errors.New("Missing note argument")
	case len(parts) == 1 && parts[0] == "create":
		fmt.Println("[Error] Missing note argument")
		return Command{}, errors.New("Missing note argument")

	}

	return Command{
		Name: parts[0],
		Args: parts[1:],
	}, nil
}

// Command struct
type Command struct {
	Name string
	Args []string
}

// Executor interface e CommandExecutor struct
type Executor interface {
	Execute(cmd Command) error
}

type CommandFunc func(args []string) error

type CommandExecutor struct {
	commands map[string]CommandFunc
	sdb      *StorageDB
}

func newExecutor() *CommandExecutor {
	sdb := &StorageDB{
		data: []string{},
	}

	ce := &CommandExecutor{
		sdb: sdb,
	}

	ce.commands = map[string]CommandFunc{
		"create": ce.create,
		"list":   ce.read,
		"update": ce.update,
		"delete": ce.delete,
		"clear":  ce.clear,
		"exit":   exit,
	}
	return ce
}
func (ce *CommandExecutor) Execute(cmd Command) error {
	if commandFunc, exists := ce.commands[cmd.Name]; exists {
		//fmt.Println(cmd.Args)
		return commandFunc(cmd.Args)
	} else {
		fmt.Println("[Error] Unknown command")
	}
	return nil
}

type StorageDB struct {
	data []string
}

func (sdb *StorageDB) Store(data string) {

	sdb.data = append(sdb.data, data)

}

// Funções de comando
func (ce *CommandExecutor) create(args []string) error {

	if len(ce.sdb.data) < maxNotes {

		data := strings.Join(args, " ")

		ce.sdb.Store(data)

		fmt.Println("[OK] The note was successfully created.")
	} else {
		fmt.Println("[Error] Notepad is full.")
	}
	return nil
}

func (ce *CommandExecutor) read(args []string) error {
	if len(ce.sdb.data) == 0 {
		fmt.Println("[Info] Notepad is empty")
	} else {
		for key, value := range ce.sdb.data {
			if value != "" {
				fmt.Printf("[Info] %d: %s\n", key+1, value)
			}
		}
	}
	return nil
}

func (ce *CommandExecutor) update(args []string) error {

	var pos int
	var err error
	if len(args) > 0 {
		pos, err = strconv.Atoi(args[0])

	}
	var usrNotes string

	switch {
	case err != nil:
		fmt.Printf("[Error] Invalid position: %s\n", strings.Join(args, " "))
	case len(args) == 1:
		fmt.Printf("[Error] Missing note argument.\n")
	case len(ce.sdb.data) == 0:
		fmt.Printf("[Error] There is nothing to update\n\n")

	case pos+1 > maxNotes:
		fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", pos, maxNotes)

	case len(args) == 0:
		fmt.Printf("[Error] Missing position argument\n")

	default:
		fmt.Printf("[OK] The note at position %d was successfully updated\n", pos)
		usrNotes = strings.Join(args[1:], " ")
		ce.sdb.data[pos-1] = usrNotes
	}

	return nil
}

func (ce *CommandExecutor) clear(args []string) error {
	ce.sdb.data = []string{}
	fmt.Println("[OK] All notes were successfully deleted")
	return nil
}

func (ce *CommandExecutor) delete(args []string) error {
	var pos int
	var err error
	if len(args) > 0 {
		pos, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("[Error] Invalid position: %s\n", strings.Join(args, " "))
			return err
		}
		pos -= 1
	}

	switch {
	case pos+1 > maxNotes || pos < 0:
		fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", pos+1, maxNotes)
	case len(ce.sdb.data) == 0 || pos > len(ce.sdb.data):
		fmt.Printf("[Error] There is nothing to delete\n")
	case len(args) == 0:
		fmt.Printf("[Error] Missing position argument\n")
	default:
		fmt.Printf("[OK] The note at position %d was successfully deleted\n", pos+1)
		ce.sdb.data = append(ce.sdb.data[:pos], ce.sdb.data[pos+1:]...)

	}

	return nil
}

func exit(args []string) error {
	fmt.Println("[Info] Bye!")
	os.Exit(0)
	return nil
}
