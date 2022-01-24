package utils

// Reference Site:
// https://developpaper.com/implementation-of-generating-uuid-unique-id-in-golang/

import "fmt"

const (
	MAXUINT32              = 4294967295
	DEFAULT_UUID_CNT_CACHE = 512
)

type UUIDGenerator struct {
	Prefix       string
	idGen        uint32
	internalChan chan uint32
}

func NewUUIDGenerator(prefix string) *UUIDGenerator {
	gen := &UUIDGenerator{
		Prefix:       prefix,
		idGen:        0,
		internalChan: make(chan uint32, DEFAULT_UUID_CNT_CACHE),
	}
	gen.startGen()
	return gen
}

//Open goroutine and put the generated UUID in digital form into the buffer pipe
func (this *UUIDGenerator) startGen() {
	go func() {
		for {
			if this.idGen == MAXUINT32 {
				this.idGen = 1
			} else {
				this.idGen += 1
			}
			this.internalChan <- this.idGen
		}
	}()
}

//Gets the UUID in the form of a prefixed string
func (this *UUIDGenerator) Get() string {
	idgen := <-this.internalChan
	return fmt.Sprintf("%s%d", this.Prefix, idgen)
}

//Get UUID in uint32 form
func (this *UUIDGenerator) GetUint32() uint32 {
	return <-this.internalChan
}
