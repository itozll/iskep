package iflag

import (
	"net"
	"time"
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

	handlerNewFlag[T any] func(set *pflag.FlagSet, p *T, name, shorthand string, value T, usage string)

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

func newFlag[T any](pf handlerNewFlag[T], p *T, name, shorthand string, value T, usage string) *flag[T] {
	if p != nil {
		if binders[unsafe.Pointer(p)] {
			panic("重复绑定变量: " + name)
		} else {
			binders[unsafe.Pointer(p)] = true
		}
	}

	return &flag[T]{
		Name:    name,
		Alias:   shorthand,
		Usage:   usage,
		Binder:  p,
		pf:      pf,
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

func NewBytes(p *[]byte, name, shorthand string, value []byte, usage string) Value[[]byte] {
	return newFlag((*pflag.FlagSet).BytesBase64VarP, p, name, shorthand, value, usage)
}

func NewBool(p *bool, name, shorthand string, value bool, usage string) Value[bool] {
	return newFlag((*pflag.FlagSet).BoolVarP, p, name, shorthand, value, usage)
}

func NewBoolSlice(p *[]bool, name, shorthand string, value []bool, usage string) Value[[]bool] {
	return newFlag((*pflag.FlagSet).BoolSliceVarP, p, name, shorthand, value, usage)
}

func NewString(p *string, name, shorthand string, value string, usage string) Value[string] {
	return newFlag((*pflag.FlagSet).StringVarP, p, name, shorthand, value, usage)
}

func NewStringSlice(p *[]string, name, shorthand string, value []string, usage string) Value[[]string] {
	return newFlag((*pflag.FlagSet).StringSliceVarP, p, name, shorthand, value, usage)
}

func NewFloat32(p *float32, name, shorthand string, value float32, usage string) Value[float32] {
	return newFlag((*pflag.FlagSet).Float32VarP, p, name, shorthand, value, usage)
}

func NewFloat32Slice(p *[]float32, name, shorthand string, value []float32, usage string) Value[[]float32] {
	return newFlag((*pflag.FlagSet).Float32SliceVarP, p, name, shorthand, value, usage)
}

func NewFloat64(p *float64, name, shorthand string, value float64, usage string) Value[float64] {
	return newFlag((*pflag.FlagSet).Float64VarP, p, name, shorthand, value, usage)
}

func NewFloat64Slice(p *[]float64, name, shorthand string, value []float64, usage string) Value[[]float64] {
	return newFlag((*pflag.FlagSet).Float64SliceVarP, p, name, shorthand, value, usage)
}

func NewDuration(p *time.Duration, name, shorthand string, value time.Duration, usage string) Value[time.Duration] {
	return newFlag((*pflag.FlagSet).DurationVarP, p, name, shorthand, value, usage)
}

func NewDurationSlice(p *[]time.Duration, name, shorthand string, value []time.Duration, usage string) Value[[]time.Duration] {
	return newFlag((*pflag.FlagSet).DurationSliceVarP, p, name, shorthand, value, usage)
}

func countVarP(pf *pflag.FlagSet, p *int, name, shorthand string, _ int, usage string) {
	(*pflag.FlagSet).CountVarP(pf, p, name, shorthand, usage)
}
func NewCount(p *int, name, shorthand string, value int, usage string) Value[int] {
	return newFlag(countVarP, p, name, shorthand, value, usage)
}

func NewInt(p *int, name, shorthand string, value int, usage string) Value[int] {
	return newFlag((*pflag.FlagSet).IntVarP, p, name, shorthand, value, usage)
}

func NewIntSlice(p *[]int, name, shorthand string, value []int, usage string) Value[[]int] {
	return newFlag((*pflag.FlagSet).IntSliceVarP, p, name, shorthand, value, usage)
}

func NewInt8(p *int8, name, shorthand string, value int8, usage string) Value[int8] {
	return newFlag((*pflag.FlagSet).Int8VarP, p, name, shorthand, value, usage)
}

func NewInt16(p *int16, name, shorthand string, value int16, usage string) Value[int16] {
	return newFlag((*pflag.FlagSet).Int16VarP, p, name, shorthand, value, usage)
}

func NewInt32(p *int32, name, shorthand string, value int32, usage string) Value[int32] {
	return newFlag((*pflag.FlagSet).Int32VarP, p, name, shorthand, value, usage)
}

func NewInt32Slice(p *[]int32, name, shorthand string, value []int32, usage string) Value[[]int32] {
	return newFlag((*pflag.FlagSet).Int32SliceVarP, p, name, shorthand, value, usage)
}

func NewInt64(p *int64, name, shorthand string, value int64, usage string) Value[int64] {
	return newFlag((*pflag.FlagSet).Int64VarP, p, name, shorthand, value, usage)
}

func NewInt64Slice(p *[]int64, name, shorthand string, value []int64, usage string) Value[[]int64] {
	return newFlag((*pflag.FlagSet).Int64SliceVarP, p, name, shorthand, value, usage)
}

func NewUint(p *uint, name, shorthand string, value uint, usage string) Value[uint] {
	return newFlag((*pflag.FlagSet).UintVarP, p, name, shorthand, value, usage)
}

func NewUintSlice(p *[]uint, name, shorthand string, value []uint, usage string) Value[[]uint] {
	return newFlag((*pflag.FlagSet).UintSliceVarP, p, name, shorthand, value, usage)
}

func NewUint8(p *uint8, name, shorthand string, value uint8, usage string) Value[uint8] {
	return newFlag((*pflag.FlagSet).Uint8VarP, p, name, shorthand, value, usage)
}

func NewUint16(p *uint16, name, shorthand string, value uint16, usage string) Value[uint16] {
	return newFlag((*pflag.FlagSet).Uint16VarP, p, name, shorthand, value, usage)
}

func NewUint32(p *uint32, name, shorthand string, value uint32, usage string) Value[uint32] {
	return newFlag((*pflag.FlagSet).Uint32VarP, p, name, shorthand, value, usage)
}

func NewUint64(p *uint64, name, shorthand string, value uint64, usage string) Value[uint64] {
	return newFlag((*pflag.FlagSet).Uint64VarP, p, name, shorthand, value, usage)
}

func NewIP(p *net.IP, name, shorthand string, value net.IP, usage string) Value[net.IP] {
	return newFlag((*pflag.FlagSet).IPVarP, p, name, shorthand, value, usage)
}

func NewIPSlice(p *[]net.IP, name, shorthand string, value []net.IP, usage string) Value[[]net.IP] {
	return newFlag((*pflag.FlagSet).IPSliceVarP, p, name, shorthand, value, usage)
}

func NewIPNet(p *net.IPNet, name, shorthand string, value net.IPNet, usage string) Value[net.IPNet] {
	return newFlag((*pflag.FlagSet).IPNetVarP, p, name, shorthand, value, usage)
}

func NewStringToInt(p *map[string]int, name, shorthand string, value map[string]int, usage string) Value[map[string]int] {
	return newFlag((*pflag.FlagSet).StringToIntVarP, p, name, shorthand, value, usage)
}

func NewStringToInt64(p *map[string]int64, name, shorthand string, value map[string]int64, usage string) Value[map[string]int64] {
	return newFlag((*pflag.FlagSet).StringToInt64VarP, p, name, shorthand, value, usage)
}

func NewStringToString(p *map[string]string, name, shorthand string, value map[string]string, usage string) Value[map[string]string] {
	return newFlag((*pflag.FlagSet).StringToStringVarP, p, name, shorthand, value, usage)
}
