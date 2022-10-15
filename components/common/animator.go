package common

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"path"
	"time"
)

type Animator struct {
	Container       *Actor
	Sequences       map[string]*Sequence
	Current         string
	LastFrameChange time.Time
	Finished        bool
}

func NewAnimator(container *Actor, sequences map[string]*Sequence, defaultSequence string) *Animator {
	var an Animator
	an.Container = container
	an.Sequences = sequences
	an.Current = defaultSequence
	an.LastFrameChange = time.Now()
	return &an
}

func (an *Animator) OnDraw(renderer *sdl.Renderer) error {
	tex := an.Sequences[an.Current].texture()
	return DrawTexture(tex, an.Container.Position, an.Container.Angle, an.Container.DrawWidth, an.Container.DrawHeight, renderer)
}

func (an *Animator) OnUpdate() error {
	seq := an.Sequences[an.Current]
	frameInterval := float64(time.Second) / seq.SampleRate
	if time.Since(an.LastFrameChange) >= time.Duration(frameInterval) {
		an.Finished = seq.nextFrame()
		an.LastFrameChange = time.Now()
	}
	return nil
}

func (an *Animator) SetSequence(name string) {
	an.Current = name
	an.LastFrameChange = time.Now()
}

func (an *Animator) OnCollision(other *Actor) error {
	return nil
}

type Sequence struct {
	Textures   []*sdl.Texture
	Frame      int
	SampleRate float64
	Loop       bool
}

func NewSequence(filePath string, sampleRate float64, loop bool, renderer *sdl.Renderer) (*Sequence, error) {
	files, err := os.ReadDir(filePath)
	var seq Sequence
	if err != nil {
		return nil, fmt.Errorf("reading directory failed: %s \n %v", filePath, err)
	}
	for _, file := range files {
		filename := path.Join(filePath, file.Name())
		tex, err := LoadTextureFromBMP(renderer, filename)
		if err != nil {
			return nil, fmt.Errorf("loading Sequence Frame: %v", err)
		}
		seq.Textures = append(seq.Textures, tex)
	}

	seq.SampleRate = sampleRate
	seq.Loop = loop
	return &seq, nil
}

func (seq *Sequence) texture() *sdl.Texture {
	return seq.Textures[seq.Frame]
}

func (seq *Sequence) nextFrame() bool {
	if seq.Frame == len(seq.Textures)-1 {
		if seq.Loop {
			seq.Frame = 0
		} else {
			return true
		}
	} else {
		seq.Frame++
	}
	return false
}
