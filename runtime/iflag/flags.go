package iflag

import (
	"unsafe"

	"github.com/spf13/pflag"
)

type (
	Argument interface {
		Bind(set *pflag.FlagSet)
	}

	Value[T any] interface {
		Value() T

		Argument
	}

	handlerNewFlag[T any]  func(set *pflag.FlagSet, p *T, name, shorthand string, value T, usage string)
	handlerNewFlag2[T any] func(set *pflag.FlagSet, p *T, name, shorthand string, usage string)

	flag[T any] struct {
		Name    string
		Alias   string
		Usage   string
		Binder  *T
		pf      handlerNewFlag[T]
		Default T
	}
)

var binders = map[unsafe.Pointer]bool{}

func newFlag[T any](pf handlerNewFlag[T]) *flag[T] { return &flag[T]{pf: pf} }
func newFlag2[T any](pf handlerNewFlag2[T]) *flag[T] {
	return &flag[T]{pf: func(set *pflag.FlagSet, p *T, name, shorthand string, _ T, usage string) {
		pf(set, p, name, shorthand, usage)
	}}
}

func (f *flag[T]) call(p *T, name, shorthand string, value T, usage string) *flag[T] {
	if p != nil {
		if binders[unsafe.Pointer(p)] {
			panic("重复绑定变量: " + name)
		}

		binders[unsafe.Pointer(p)] = true
	}

	return &flag[T]{
		Name:    name,
		Alias:   shorthand,
		Usage:   usage,
		Binder:  p,
		pf:      f.pf,
		Default: value,
	}
}

func (f *flag[T]) Value() T { return *f.Binder }

func (arg *flag[T]) Bind(set *pflag.FlagSet) {
	if arg.Binder == nil {
		arg.Binder = new(T)
	}

	arg.pf(
		set,
		arg.Binder,
		arg.Name,
		arg.Alias,
		arg.Default,
		arg.Usage,
	)
}

var (
	NewBytes          = newFlag((*pflag.FlagSet).BytesBase64VarP).call
	NewBool           = newFlag((*pflag.FlagSet).BoolVarP).call
	NewBoolSlice      = newFlag((*pflag.FlagSet).BoolSliceVarP).call
	NewString         = newFlag((*pflag.FlagSet).StringVarP).call
	NewStringSlice    = newFlag((*pflag.FlagSet).StringSliceVarP).call
	NewFloat32        = newFlag((*pflag.FlagSet).Float32VarP).call
	NewFloat32Slice   = newFlag((*pflag.FlagSet).Float32SliceVarP).call
	NewFloat64        = newFlag((*pflag.FlagSet).Float64VarP).call
	NewFloat64Slice   = newFlag((*pflag.FlagSet).Float64SliceVarP).call
	NewDuration       = newFlag((*pflag.FlagSet).DurationVarP).call
	NewDurationSlice  = newFlag((*pflag.FlagSet).DurationSliceVarP).call
	NewInt            = newFlag((*pflag.FlagSet).IntVarP).call
	NewIntSlice       = newFlag((*pflag.FlagSet).IntSliceVarP).call
	NewInt8           = newFlag((*pflag.FlagSet).Int8VarP).call
	NewInt16          = newFlag((*pflag.FlagSet).Int16VarP).call
	NewInt32          = newFlag((*pflag.FlagSet).Int32VarP).call
	NewInt32Slice     = newFlag((*pflag.FlagSet).Int32SliceVarP).call
	NewInt64          = newFlag((*pflag.FlagSet).Int64VarP).call
	NewInt64Slice     = newFlag((*pflag.FlagSet).Int64SliceVarP).call
	NewUint           = newFlag((*pflag.FlagSet).UintVarP).call
	NewUintSlice      = newFlag((*pflag.FlagSet).UintSliceVarP).call
	NewUint8          = newFlag((*pflag.FlagSet).Uint8VarP).call
	NewUint16         = newFlag((*pflag.FlagSet).Uint16VarP).call
	NewUint32         = newFlag((*pflag.FlagSet).Uint32VarP).call
	NewUint64         = newFlag((*pflag.FlagSet).Uint64VarP).call
	NewIP             = newFlag((*pflag.FlagSet).IPVarP).call
	NewIPSlice        = newFlag((*pflag.FlagSet).IPSliceVarP).call
	NewIPNet          = newFlag((*pflag.FlagSet).IPNetVarP).call
	NewStringToInt    = newFlag((*pflag.FlagSet).StringToIntVarP).call
	NewStringToInt64  = newFlag((*pflag.FlagSet).StringToInt64VarP).call
	NewStringToString = newFlag((*pflag.FlagSet).StringToStringVarP).call

	NewCount = newFlag2((*pflag.FlagSet).CountVarP).call
)
