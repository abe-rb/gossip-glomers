package main

import (
	"encoding/json"
	"log"
	"os/exec"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
    n := maelstrom.NewNode()

    n.Handle("generate", func(msg maelstrom.Message) error {
        // Unmarshal the message body as an loosely-typed map.
        var body map[string]any
        if err := json.Unmarshal(msg.Body, &body); err != nil {
            return err
        }

        // Update the message type to return back
        body["type"] = "generate_ok"

        // Generate UUID
        uuid, err := exec.Command("uuidgen").Output()
        if err != nil {
            return err
        }

        body["id"] = uuid

        return n.Reply(msg, body)
    })

    if err := n.Run(); err != nil {
        log.Fatal(err)
    }
}
