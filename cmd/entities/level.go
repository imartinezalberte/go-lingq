package entities

import (
	"errors"
	"reflect"
	"strings"

	"github.com/imartinezalberte/go-lingq/internal/utils"
	"github.com/repeale/fp-go"
)

const (
	LevelName                  = "level"
	LevelUsage                 = "Specify here the level of your course."
	LevelDefault ResourceLevel = FirstLevel
)

type (
	ResourcesLevel utils.Set[ResourceLevel]
	ResourceLevel  uint
)

const (
	_ ResourceLevel = iota
	FirstLevel
	SecondLevel
	ThirdLevel
	FourthLevel
	FifthLevel
	SixthLevel
)

var Levels = [...]string{"A1", "A2", "B1", "B2", "C1", "C2"}

func (r *ResourceLevel) Type() string {
	return reflect.Uint.String()
}

func (r ResourceLevel) String() string {
	index := int(r) - 1
	if index >= len(Levels) || index < 0 {
		return "unknown"
	}

	return Levels[index]
}

func (r *ResourceLevel) Set(input string) error {
	if !r.Check(input) {
		return errors.New("resource level is not correct")
	}
	return nil
}

func (r *ResourceLevel) Check(input string) bool {
	switch strings.ToUpper(input) {
	case Levels[0]:
		*r = FirstLevel
	case Levels[1]:
		*r = SecondLevel
	case Levels[2]:
		*r = ThirdLevel
	case Levels[3]:
		*r = FourthLevel
	case Levels[4]:
		*r = FifthLevel
	case Levels[5]:
		*r = SixthLevel
	default:
		return false
	}
	return true
}

func (rr ResourcesLevel) InnerType() utils.Set[ResourceLevel] {
	return utils.Set[ResourceLevel](rr)
}

func (rr *ResourcesLevel) Type() string {
	return "utils.Set[ResourceLevel]"
}

func (rr *ResourcesLevel) String() string {
	return strings.Join(fp.Map(func(r ResourceLevel) string {
		return r.String()
	})(rr.InnerType().ToArr()), utils.Comma)
}

func (rr *ResourcesLevel) Set(input string) error {
	if *rr == nil {
		*rr = ResourcesLevel(utils.NewSet[ResourceLevel]())
	}

	var r ResourceLevel
	if err := r.Set(input); err != nil {
		return err
	}

	if rr.InnerType().Exists(r) {
		return errors.New("resource level already exists")
	}

	rr.InnerType().Add(r)
	return nil
}
