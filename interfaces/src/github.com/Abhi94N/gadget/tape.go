package gadget

import (
	"fmt"
)

type Player interface {
	Play(string)
	Stop()
}

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopped")
}

type TapeRecorder struct {
	Microphones int
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapeRecorder) Record() {
	fmt.Println("Recording")
}

func (t TapeRecorder) Stop() {
	fmt.Println("Stopped")
}

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	//asserting the value to confirm its a TapeRecorder
	recorder, ok := player.(TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("This aint no recorder")
	}

}
