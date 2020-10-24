package main

import (
	"fmt"
	"strconv"
)

const NOP = -1

type Flow interface {
	// Next returns next portion of fun. Use -1 when there is no choices.
	Next(choice int) (string, []string)
}

type flow struct{}

func (f *flow) Next(choice int) (string, []string) {
	if choice == 0 {
		// without choices test:
		return "Text without choices", nil
	}
	return "Hello, world!", []string{"Ugh.", "Hello!", "Let's talk."}
}

type cli struct {
	flow Flow

	choice  int
	text    string
	choices []string

	counter int
}

func newCLI(f Flow) *cli {
	return &cli{
		flow:   f,
		choice: NOP,
	}
}

func (c *cli) Run() {
	c.hello()
	c.next()
	for {
		c.show()
		next, exit := c.input()
		if exit {
			c.bye()
			break
		}
		if next {
			c.next()
		}
		// show again
	}
}

func (c *cli) hello() {
	fmt.Println("[BEGIN]")
}
func (c *cli) bye() {
	fmt.Println("[GOODBYE]")
}

func (c *cli) next() {
	c.text, c.choices = c.flow.Next(c.choice)
	c.counter++
}

func (c *cli) show() {
	fmt.Println()
	fmt.Printf("[step:%d]\n", c.counter)
	fmt.Println(c.text)
	for i, c := range c.choices {
		fmt.Printf("%d) %s\n", i+1, c)
	}
}

func (c *cli) input() (next bool, exit bool) {
	var line string

	var err error
	if len(c.choices) > 0 {
		fmt.Print("> ")
		_, err = fmt.Scan(&line)
	} else {
		fmt.Println("[continue]")
		_, err = fmt.Scan()
	}

	if err != nil {
		fmt.Printf("[%s]", err.Error())
		return false, false
	}

	if line == "q" || line == "Q" {
		return false, true
	}

	if len(c.choices) == 0 {
		c.choice = NOP
		return true, false
	}

	ch, err := strconv.Atoi(line)
	if err != nil {
		fmt.Printf("[%s]", err.Error())
		return false, false
	}

	ch--
	if ch < 0 || ch >= len(c.choices) {
		return false, false
	}

	c.choice = ch
	return true, false
}

func main() {
	fl := (*flow)(nil)
	g := newCLI(fl)
	g.Run()
}
