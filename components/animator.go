package components

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"path"
	"time"
)

type animator struct {
	container       *Actor
	sequences       map[string]*sequence
	current         string
	lastFrameChange time.Time
	finished        bool
}

func newAnimator(container *Actor, sequences map[string]*sequence, defaultSequence string) *animator {
	var an animator
	an.container = container
	an.sequences = sequences
	an.current = defaultSequence
	an.lastFrameChange = time.Now()
	return &an
}

func (an *animator) onDraw(renderer *sdl.Renderer) error {
	tex := an.sequences[an.current].texture()
	return drawTexture(tex, an.container.position, an.container.angle, an.container.drawWidth, an.container.drawHeight, renderer)
}

func (an *animator) onUpdate() error {
	seq := an.sequences[an.current]
	frameInterval := float64(time.Second) / seq.sampleRate
	if time.Since(an.lastFrameChange) >= time.Duration(frameInterval) {
		an.finished = seq.nextFrame()
		an.lastFrameChange = time.Now()
	}
	return nil
}

func (an *animator) setSequence(name string) {
	an.current = name
	an.lastFrameChange = time.Now()
}

func (an *animator) onCollision(other *Actor) error {
	return nil
}

type sequence struct {
	textures   []*sdl.Texture
	frame      int
	sampleRate float64
	loop       bool
}

func newSequence(filePath string, sampleRate float64, loop bool, renderer *sdl.Renderer) (*sequence, error) {
	files, err := os.ReadDir(filePath)
	var seq sequence
	if err != nil {
		return nil, fmt.Errorf("reading directory failed: %s \n %v", filePath, err)
	}
	for _, file := range files {
		filename := path.Join(filePath, file.Name())
		tex, err := loadTextureFromBMP(renderer, filename)
		if err != nil {
			return nil, fmt.Errorf("loading sequence frame: %v", err)
		}
		seq.textures = append(seq.textures, tex)
	}

	seq.sampleRate = sampleRate
	seq.loop = loop
	return &seq, nil
}

func (seq *sequence) texture() *sdl.Texture {
	return seq.textures[seq.frame]
}

func (seq *sequence) nextFrame() bool {
	if seq.frame == len(seq.textures)-1 {
		if seq.loop {
			seq.frame = 0
		} else {
			return true
		}
	} else {
		seq.frame++
	}
	return false
}
