package client

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	time "time"
	codec1978 "github.com/ugorji/go/codec"
)

const (
	codecSelferC_UTF87612			= 1
	codecSelferC_RAW7612			= 0
	codecSelferValueTypeArray7612		= 10
	codecSelferValueTypeMap7612		= 9
	codecSelfer_containerMapKey7612		= 2
	codecSelfer_containerMapValue7612	= 3
	codecSelfer_containerMapEnd7612		= 4
	codecSelfer_containerArrayElem7612	= 6
	codecSelfer_containerArrayEnd7612	= 7
)

var (
	codecSelferBitsize7612				= uint8(reflect.TypeOf(uint(0)).Bits())
	codecSelferOnlyMapOrArrayEncodeToStructErr7612	= errors.New(`only encoded map or array can be decoded into a struct`)
)

type codecSelfer7612 struct{}

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if codec1978.GenVersion != 8 {
		_, file, _, _ := runtime.Caller(0)
		err := fmt.Errorf("codecgen version mismatch: current: %v, need %v. Re-generate file: %v", 8, codec1978.GenVersion, file)
		panic(err)
	}
	if false {
		var v0 time.Duration
		_ = v0
	}
}
func (x *Error) CodecEncodeSelf(e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			_, _ = yysep2, yy2arr2
			const yyr2 bool = false
			if yyr2 || yy2arr2 {
				r.WriteArrayStart(4)
			} else {
				r.WriteMapStart(4)
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym4 := z.EncBinary()
				_ = yym4
				if false {
				} else {
					r.EncodeInt(int64(x.Code))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("errorCode"))
				r.WriteMapElemValue()
				yym5 := z.EncBinary()
				_ = yym5
				if false {
				} else {
					r.EncodeInt(int64(x.Code))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym7 := z.EncBinary()
				_ = yym7
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Message))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("message"))
				r.WriteMapElemValue()
				yym8 := z.EncBinary()
				_ = yym8
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Message))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym10 := z.EncBinary()
				_ = yym10
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Cause))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("cause"))
				r.WriteMapElemValue()
				yym11 := z.EncBinary()
				_ = yym11
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Cause))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym13 := z.EncBinary()
				_ = yym13
				if false {
				} else {
					r.EncodeUint(uint64(x.Index))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("index"))
				r.WriteMapElemValue()
				yym14 := z.EncBinary()
				_ = yym14
				if false {
				} else {
					r.EncodeUint(uint64(x.Index))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayEnd()
			} else {
				r.WriteMapEnd()
			}
		}
	}
}
func (x *Error) CodecDecodeSelf(d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap7612 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray7612 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr7612)
		}
	}
}
func (x *Error) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer()
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		r.ReadMapElemKey()
		yys3Slc = r.DecodeStringAsBytes()
		yys3 := string(yys3Slc)
		r.ReadMapElemValue()
		switch yys3 {
		case "errorCode":
			if r.TryDecodeAsNil() {
				x.Code = 0
			} else {
				yyv4 := &x.Code
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else {
					*((*int)(yyv4)) = int(r.DecodeInt(codecSelferBitsize7612))
				}
			}
		case "message":
			if r.TryDecodeAsNil() {
				x.Message = ""
			} else {
				yyv6 := &x.Message
				yym7 := z.DecBinary()
				_ = yym7
				if false {
				} else {
					*((*string)(yyv6)) = r.DecodeString()
				}
			}
		case "cause":
			if r.TryDecodeAsNil() {
				x.Cause = ""
			} else {
				yyv8 := &x.Cause
				yym9 := z.DecBinary()
				_ = yym9
				if false {
				} else {
					*((*string)(yyv8)) = r.DecodeString()
				}
			}
		case "index":
			if r.TryDecodeAsNil() {
				x.Index = 0
			} else {
				yyv10 := &x.Index
				yym11 := z.DecBinary()
				_ = yym11
				if false {
				} else {
					*((*uint64)(yyv10)) = uint64(r.DecodeUint(64))
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		}
	}
	r.ReadMapEnd()
}
func (x *Error) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj12 int
	var yyb12 bool
	var yyhl12 bool = l >= 0
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Code = 0
	} else {
		yyv13 := &x.Code
		yym14 := z.DecBinary()
		_ = yym14
		if false {
		} else {
			*((*int)(yyv13)) = int(r.DecodeInt(codecSelferBitsize7612))
		}
	}
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Message = ""
	} else {
		yyv15 := &x.Message
		yym16 := z.DecBinary()
		_ = yym16
		if false {
		} else {
			*((*string)(yyv15)) = r.DecodeString()
		}
	}
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Cause = ""
	} else {
		yyv17 := &x.Cause
		yym18 := z.DecBinary()
		_ = yym18
		if false {
		} else {
			*((*string)(yyv17)) = r.DecodeString()
		}
	}
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Index = 0
	} else {
		yyv19 := &x.Index
		yym20 := z.DecBinary()
		_ = yym20
		if false {
		} else {
			*((*uint64)(yyv19)) = uint64(r.DecodeUint(64))
		}
	}
	for {
		yyj12++
		if yyhl12 {
			yyb12 = yyj12 > l
		} else {
			yyb12 = r.CheckBreak()
		}
		if yyb12 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj12-1, "")
	}
	r.ReadArrayEnd()
}
func (x PrevExistType) CodecEncodeSelf(e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	yym1 := z.EncBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.EncExt(x) {
	} else {
		r.EncodeString(codecSelferC_UTF87612, string(x))
	}
}
func (x *PrevExistType) CodecDecodeSelf(d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		*((*string)(x)) = r.DecodeString()
	}
}
func (x *WatcherOptions) CodecEncodeSelf(e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			_, _ = yysep2, yy2arr2
			const yyr2 bool = false
			if yyr2 || yy2arr2 {
				r.WriteArrayStart(2)
			} else {
				r.WriteMapStart(2)
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym4 := z.EncBinary()
				_ = yym4
				if false {
				} else {
					r.EncodeUint(uint64(x.AfterIndex))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("AfterIndex"))
				r.WriteMapElemValue()
				yym5 := z.EncBinary()
				_ = yym5
				if false {
				} else {
					r.EncodeUint(uint64(x.AfterIndex))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym7 := z.EncBinary()
				_ = yym7
				if false {
				} else {
					r.EncodeBool(bool(x.Recursive))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("Recursive"))
				r.WriteMapElemValue()
				yym8 := z.EncBinary()
				_ = yym8
				if false {
				} else {
					r.EncodeBool(bool(x.Recursive))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayEnd()
			} else {
				r.WriteMapEnd()
			}
		}
	}
}
func (x *WatcherOptions) CodecDecodeSelf(d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap7612 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray7612 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr7612)
		}
	}
}
func (x *WatcherOptions) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer()
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		r.ReadMapElemKey()
		yys3Slc = r.DecodeStringAsBytes()
		yys3 := string(yys3Slc)
		r.ReadMapElemValue()
		switch yys3 {
		case "AfterIndex":
			if r.TryDecodeAsNil() {
				x.AfterIndex = 0
			} else {
				yyv4 := &x.AfterIndex
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else {
					*((*uint64)(yyv4)) = uint64(r.DecodeUint(64))
				}
			}
		case "Recursive":
			if r.TryDecodeAsNil() {
				x.Recursive = false
			} else {
				yyv6 := &x.Recursive
				yym7 := z.DecBinary()
				_ = yym7
				if false {
				} else {
					*((*bool)(yyv6)) = r.DecodeBool()
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		}
	}
	r.ReadMapEnd()
}
func (x *WatcherOptions) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj8 int
	var yyb8 bool
	var yyhl8 bool = l >= 0
	yyj8++
	if yyhl8 {
		yyb8 = yyj8 > l
	} else {
		yyb8 = r.CheckBreak()
	}
	if yyb8 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.AfterIndex = 0
	} else {
		yyv9 := &x.AfterIndex
		yym10 := z.DecBinary()
		_ = yym10
		if false {
		} else {
			*((*uint64)(yyv9)) = uint64(r.DecodeUint(64))
		}
	}
	yyj8++
	if yyhl8 {
		yyb8 = yyj8 > l
	} else {
		yyb8 = r.CheckBreak()
	}
	if yyb8 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Recursive = false
	} else {
		yyv11 := &x.Recursive
		yym12 := z.DecBinary()
		_ = yym12
		if false {
		} else {
			*((*bool)(yyv11)) = r.DecodeBool()
		}
	}
	for {
		yyj8++
		if yyhl8 {
			yyb8 = yyj8 > l
		} else {
			yyb8 = r.CheckBreak()
		}
		if yyb8 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj8-1, "")
	}
	r.ReadArrayEnd()
}
func (x *CreateInOrderOptions) CodecEncodeSelf(e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			_, _ = yysep2, yy2arr2
			const yyr2 bool = false
			if yyr2 || yy2arr2 {
				r.WriteArrayStart(1)
			} else {
				r.WriteMapStart(1)
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym4 := z.EncBinary()
				_ = yym4
				if false {
				} else if z.HasExtensions() && z.EncExt(x.TTL) {
				} else {
					r.EncodeInt(int64(x.TTL))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("TTL"))
				r.WriteMapElemValue()
				yym5 := z.EncBinary()
				_ = yym5
				if false {
				} else if z.HasExtensions() && z.EncExt(x.TTL) {
				} else {
					r.EncodeInt(int64(x.TTL))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayEnd()
			} else {
				r.WriteMapEnd()
			}
		}
	}
}
func (x *CreateInOrderOptions) CodecDecodeSelf(d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap7612 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray7612 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr7612)
		}
	}
}
func (x *CreateInOrderOptions) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer()
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		r.ReadMapElemKey()
		yys3Slc = r.DecodeStringAsBytes()
		yys3 := string(yys3Slc)
		r.ReadMapElemValue()
		switch yys3 {
		case "TTL":
			if r.TryDecodeAsNil() {
				x.TTL = 0
			} else {
				yyv4 := &x.TTL
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else if z.HasExtensions() && z.DecExt(yyv4) {
				} else {
					*((*int64)(yyv4)) = int64(r.DecodeInt(64))
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		}
	}
	r.ReadMapEnd()
}
func (x *CreateInOrderOptions) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj6 int
	var yyb6 bool
	var yyhl6 bool = l >= 0
	yyj6++
	if yyhl6 {
		yyb6 = yyj6 > l
	} else {
		yyb6 = r.CheckBreak()
	}
	if yyb6 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.TTL = 0
	} else {
		yyv7 := &x.TTL
		yym8 := z.DecBinary()
		_ = yym8
		if false {
		} else if z.HasExtensions() && z.DecExt(yyv7) {
		} else {
			*((*int64)(yyv7)) = int64(r.DecodeInt(64))
		}
	}
	for {
		yyj6++
		if yyhl6 {
			yyb6 = yyj6 > l
		} else {
			yyb6 = r.CheckBreak()
		}
		if yyb6 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj6-1, "")
	}
	r.ReadArrayEnd()
}
func (x *SetOptions) CodecEncodeSelf(e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			_, _ = yysep2, yy2arr2
			const yyr2 bool = false
			if yyr2 || yy2arr2 {
				r.WriteArrayStart(7)
			} else {
				r.WriteMapStart(7)
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym4 := z.EncBinary()
				_ = yym4
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.PrevValue))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("PrevValue"))
				r.WriteMapElemValue()
				yym5 := z.EncBinary()
				_ = yym5
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.PrevValue))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym7 := z.EncBinary()
				_ = yym7
				if false {
				} else {
					r.EncodeUint(uint64(x.PrevIndex))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("PrevIndex"))
				r.WriteMapElemValue()
				yym8 := z.EncBinary()
				_ = yym8
				if false {
				} else {
					r.EncodeUint(uint64(x.PrevIndex))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				x.PrevExist.CodecEncodeSelf(e)
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("PrevExist"))
				r.WriteMapElemValue()
				x.PrevExist.CodecEncodeSelf(e)
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym13 := z.EncBinary()
				_ = yym13
				if false {
				} else if z.HasExtensions() && z.EncExt(x.TTL) {
				} else {
					r.EncodeInt(int64(x.TTL))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("TTL"))
				r.WriteMapElemValue()
				yym14 := z.EncBinary()
				_ = yym14
				if false {
				} else if z.HasExtensions() && z.EncExt(x.TTL) {
				} else {
					r.EncodeInt(int64(x.TTL))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym16 := z.EncBinary()
				_ = yym16
				if false {
				} else {
					r.EncodeBool(bool(x.Refresh))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("Refresh"))
				r.WriteMapElemValue()
				yym17 := z.EncBinary()
				_ = yym17
				if false {
				} else {
					r.EncodeBool(bool(x.Refresh))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym19 := z.EncBinary()
				_ = yym19
				if false {
				} else {
					r.EncodeBool(bool(x.Dir))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("Dir"))
				r.WriteMapElemValue()
				yym20 := z.EncBinary()
				_ = yym20
				if false {
				} else {
					r.EncodeBool(bool(x.Dir))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym22 := z.EncBinary()
				_ = yym22
				if false {
				} else {
					r.EncodeBool(bool(x.NoValueOnSuccess))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("NoValueOnSuccess"))
				r.WriteMapElemValue()
				yym23 := z.EncBinary()
				_ = yym23
				if false {
				} else {
					r.EncodeBool(bool(x.NoValueOnSuccess))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayEnd()
			} else {
				r.WriteMapEnd()
			}
		}
	}
}
func (x *SetOptions) CodecDecodeSelf(d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap7612 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray7612 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr7612)
		}
	}
}
func (x *SetOptions) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer()
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		r.ReadMapElemKey()
		yys3Slc = r.DecodeStringAsBytes()
		yys3 := string(yys3Slc)
		r.ReadMapElemValue()
		switch yys3 {
		case "PrevValue":
			if r.TryDecodeAsNil() {
				x.PrevValue = ""
			} else {
				yyv4 := &x.PrevValue
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else {
					*((*string)(yyv4)) = r.DecodeString()
				}
			}
		case "PrevIndex":
			if r.TryDecodeAsNil() {
				x.PrevIndex = 0
			} else {
				yyv6 := &x.PrevIndex
				yym7 := z.DecBinary()
				_ = yym7
				if false {
				} else {
					*((*uint64)(yyv6)) = uint64(r.DecodeUint(64))
				}
			}
		case "PrevExist":
			if r.TryDecodeAsNil() {
				x.PrevExist = ""
			} else {
				yyv8 := &x.PrevExist
				yyv8.CodecDecodeSelf(d)
			}
		case "TTL":
			if r.TryDecodeAsNil() {
				x.TTL = 0
			} else {
				yyv9 := &x.TTL
				yym10 := z.DecBinary()
				_ = yym10
				if false {
				} else if z.HasExtensions() && z.DecExt(yyv9) {
				} else {
					*((*int64)(yyv9)) = int64(r.DecodeInt(64))
				}
			}
		case "Refresh":
			if r.TryDecodeAsNil() {
				x.Refresh = false
			} else {
				yyv11 := &x.Refresh
				yym12 := z.DecBinary()
				_ = yym12
				if false {
				} else {
					*((*bool)(yyv11)) = r.DecodeBool()
				}
			}
		case "Dir":
			if r.TryDecodeAsNil() {
				x.Dir = false
			} else {
				yyv13 := &x.Dir
				yym14 := z.DecBinary()
				_ = yym14
				if false {
				} else {
					*((*bool)(yyv13)) = r.DecodeBool()
				}
			}
		case "NoValueOnSuccess":
			if r.TryDecodeAsNil() {
				x.NoValueOnSuccess = false
			} else {
				yyv15 := &x.NoValueOnSuccess
				yym16 := z.DecBinary()
				_ = yym16
				if false {
				} else {
					*((*bool)(yyv15)) = r.DecodeBool()
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		}
	}
	r.ReadMapEnd()
}
func (x *SetOptions) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj17 int
	var yyb17 bool
	var yyhl17 bool = l >= 0
	yyj17++
	if yyhl17 {
		yyb17 = yyj17 > l
	} else {
		yyb17 = r.CheckBreak()
	}
	if yyb17 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.PrevValue = ""
	} else {
		yyv18 := &x.PrevValue
		yym19 := z.DecBinary()
		_ = yym19
		if false {
		} else {
			*((*string)(yyv18)) = r.DecodeString()
		}
	}
	yyj17++
	if yyhl17 {
		yyb17 = yyj17 > l
	} else {
		yyb17 = r.CheckBreak()
	}
	if yyb17 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.PrevIndex = 0
	} else {
		yyv20 := &x.PrevIndex
		yym21 := z.DecBinary()
		_ = yym21
		if false {
		} else {
			*((*uint64)(yyv20)) = uint64(r.DecodeUint(64))
		}
	}
	yyj17++
	if yyhl17 {
		yyb17 = yyj17 > l
	} else {
		yyb17 = r.CheckBreak()
	}
	if yyb17 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.PrevExist = ""
	} else {
		yyv22 := &x.PrevExist
		yyv22.CodecDecodeSelf(d)
	}
	yyj17++
	if yyhl17 {
		yyb17 = yyj17 > l
	} else {
		yyb17 = r.CheckBreak()
	}
	if yyb17 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.TTL = 0
	} else {
		yyv23 := &x.TTL
		yym24 := z.DecBinary()
		_ = yym24
		if false {
		} else if z.HasExtensions() && z.DecExt(yyv23) {
		} else {
			*((*int64)(yyv23)) = int64(r.DecodeInt(64))
		}
	}
	yyj17++
	if yyhl17 {
		yyb17 = yyj17 > l
	} else {
		yyb17 = r.CheckBreak()
	}
	if yyb17 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Refresh = false
	} else {
		yyv25 := &x.Refresh
		yym26 := z.DecBinary()
		_ = yym26
		if false {
		} else {
			*((*bool)(yyv25)) = r.DecodeBool()
		}
	}
	yyj17++
	if yyhl17 {
		yyb17 = yyj17 > l
	} else {
		yyb17 = r.CheckBreak()
	}
	if yyb17 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Dir = false
	} else {
		yyv27 := &x.Dir
		yym28 := z.DecBinary()
		_ = yym28
		if false {
		} else {
			*((*bool)(yyv27)) = r.DecodeBool()
		}
	}
	yyj17++
	if yyhl17 {
		yyb17 = yyj17 > l
	} else {
		yyb17 = r.CheckBreak()
	}
	if yyb17 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.NoValueOnSuccess = false
	} else {
		yyv29 := &x.NoValueOnSuccess
		yym30 := z.DecBinary()
		_ = yym30
		if false {
		} else {
			*((*bool)(yyv29)) = r.DecodeBool()
		}
	}
	for {
		yyj17++
		if yyhl17 {
			yyb17 = yyj17 > l
		} else {
			yyb17 = r.CheckBreak()
		}
		if yyb17 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj17-1, "")
	}
	r.ReadArrayEnd()
}
func (x *GetOptions) CodecEncodeSelf(e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			_, _ = yysep2, yy2arr2
			const yyr2 bool = false
			if yyr2 || yy2arr2 {
				r.WriteArrayStart(3)
			} else {
				r.WriteMapStart(3)
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym4 := z.EncBinary()
				_ = yym4
				if false {
				} else {
					r.EncodeBool(bool(x.Recursive))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("Recursive"))
				r.WriteMapElemValue()
				yym5 := z.EncBinary()
				_ = yym5
				if false {
				} else {
					r.EncodeBool(bool(x.Recursive))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym7 := z.EncBinary()
				_ = yym7
				if false {
				} else {
					r.EncodeBool(bool(x.Sort))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("Sort"))
				r.WriteMapElemValue()
				yym8 := z.EncBinary()
				_ = yym8
				if false {
				} else {
					r.EncodeBool(bool(x.Sort))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym10 := z.EncBinary()
				_ = yym10
				if false {
				} else {
					r.EncodeBool(bool(x.Quorum))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("Quorum"))
				r.WriteMapElemValue()
				yym11 := z.EncBinary()
				_ = yym11
				if false {
				} else {
					r.EncodeBool(bool(x.Quorum))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayEnd()
			} else {
				r.WriteMapEnd()
			}
		}
	}
}
func (x *GetOptions) CodecDecodeSelf(d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap7612 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray7612 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr7612)
		}
	}
}
func (x *GetOptions) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer()
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		r.ReadMapElemKey()
		yys3Slc = r.DecodeStringAsBytes()
		yys3 := string(yys3Slc)
		r.ReadMapElemValue()
		switch yys3 {
		case "Recursive":
			if r.TryDecodeAsNil() {
				x.Recursive = false
			} else {
				yyv4 := &x.Recursive
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else {
					*((*bool)(yyv4)) = r.DecodeBool()
				}
			}
		case "Sort":
			if r.TryDecodeAsNil() {
				x.Sort = false
			} else {
				yyv6 := &x.Sort
				yym7 := z.DecBinary()
				_ = yym7
				if false {
				} else {
					*((*bool)(yyv6)) = r.DecodeBool()
				}
			}
		case "Quorum":
			if r.TryDecodeAsNil() {
				x.Quorum = false
			} else {
				yyv8 := &x.Quorum
				yym9 := z.DecBinary()
				_ = yym9
				if false {
				} else {
					*((*bool)(yyv8)) = r.DecodeBool()
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		}
	}
	r.ReadMapEnd()
}
func (x *GetOptions) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj10 int
	var yyb10 bool
	var yyhl10 bool = l >= 0
	yyj10++
	if yyhl10 {
		yyb10 = yyj10 > l
	} else {
		yyb10 = r.CheckBreak()
	}
	if yyb10 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Recursive = false
	} else {
		yyv11 := &x.Recursive
		yym12 := z.DecBinary()
		_ = yym12
		if false {
		} else {
			*((*bool)(yyv11)) = r.DecodeBool()
		}
	}
	yyj10++
	if yyhl10 {
		yyb10 = yyj10 > l
	} else {
		yyb10 = r.CheckBreak()
	}
	if yyb10 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Sort = false
	} else {
		yyv13 := &x.Sort
		yym14 := z.DecBinary()
		_ = yym14
		if false {
		} else {
			*((*bool)(yyv13)) = r.DecodeBool()
		}
	}
	yyj10++
	if yyhl10 {
		yyb10 = yyj10 > l
	} else {
		yyb10 = r.CheckBreak()
	}
	if yyb10 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Quorum = false
	} else {
		yyv15 := &x.Quorum
		yym16 := z.DecBinary()
		_ = yym16
		if false {
		} else {
			*((*bool)(yyv15)) = r.DecodeBool()
		}
	}
	for {
		yyj10++
		if yyhl10 {
			yyb10 = yyj10 > l
		} else {
			yyb10 = r.CheckBreak()
		}
		if yyb10 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj10-1, "")
	}
	r.ReadArrayEnd()
}
func (x *DeleteOptions) CodecEncodeSelf(e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			_, _ = yysep2, yy2arr2
			const yyr2 bool = false
			if yyr2 || yy2arr2 {
				r.WriteArrayStart(4)
			} else {
				r.WriteMapStart(4)
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym4 := z.EncBinary()
				_ = yym4
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.PrevValue))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("PrevValue"))
				r.WriteMapElemValue()
				yym5 := z.EncBinary()
				_ = yym5
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.PrevValue))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym7 := z.EncBinary()
				_ = yym7
				if false {
				} else {
					r.EncodeUint(uint64(x.PrevIndex))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("PrevIndex"))
				r.WriteMapElemValue()
				yym8 := z.EncBinary()
				_ = yym8
				if false {
				} else {
					r.EncodeUint(uint64(x.PrevIndex))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym10 := z.EncBinary()
				_ = yym10
				if false {
				} else {
					r.EncodeBool(bool(x.Recursive))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("Recursive"))
				r.WriteMapElemValue()
				yym11 := z.EncBinary()
				_ = yym11
				if false {
				} else {
					r.EncodeBool(bool(x.Recursive))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym13 := z.EncBinary()
				_ = yym13
				if false {
				} else {
					r.EncodeBool(bool(x.Dir))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("Dir"))
				r.WriteMapElemValue()
				yym14 := z.EncBinary()
				_ = yym14
				if false {
				} else {
					r.EncodeBool(bool(x.Dir))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayEnd()
			} else {
				r.WriteMapEnd()
			}
		}
	}
}
func (x *DeleteOptions) CodecDecodeSelf(d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap7612 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray7612 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr7612)
		}
	}
}
func (x *DeleteOptions) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer()
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		r.ReadMapElemKey()
		yys3Slc = r.DecodeStringAsBytes()
		yys3 := string(yys3Slc)
		r.ReadMapElemValue()
		switch yys3 {
		case "PrevValue":
			if r.TryDecodeAsNil() {
				x.PrevValue = ""
			} else {
				yyv4 := &x.PrevValue
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else {
					*((*string)(yyv4)) = r.DecodeString()
				}
			}
		case "PrevIndex":
			if r.TryDecodeAsNil() {
				x.PrevIndex = 0
			} else {
				yyv6 := &x.PrevIndex
				yym7 := z.DecBinary()
				_ = yym7
				if false {
				} else {
					*((*uint64)(yyv6)) = uint64(r.DecodeUint(64))
				}
			}
		case "Recursive":
			if r.TryDecodeAsNil() {
				x.Recursive = false
			} else {
				yyv8 := &x.Recursive
				yym9 := z.DecBinary()
				_ = yym9
				if false {
				} else {
					*((*bool)(yyv8)) = r.DecodeBool()
				}
			}
		case "Dir":
			if r.TryDecodeAsNil() {
				x.Dir = false
			} else {
				yyv10 := &x.Dir
				yym11 := z.DecBinary()
				_ = yym11
				if false {
				} else {
					*((*bool)(yyv10)) = r.DecodeBool()
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		}
	}
	r.ReadMapEnd()
}
func (x *DeleteOptions) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj12 int
	var yyb12 bool
	var yyhl12 bool = l >= 0
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.PrevValue = ""
	} else {
		yyv13 := &x.PrevValue
		yym14 := z.DecBinary()
		_ = yym14
		if false {
		} else {
			*((*string)(yyv13)) = r.DecodeString()
		}
	}
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.PrevIndex = 0
	} else {
		yyv15 := &x.PrevIndex
		yym16 := z.DecBinary()
		_ = yym16
		if false {
		} else {
			*((*uint64)(yyv15)) = uint64(r.DecodeUint(64))
		}
	}
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Recursive = false
	} else {
		yyv17 := &x.Recursive
		yym18 := z.DecBinary()
		_ = yym18
		if false {
		} else {
			*((*bool)(yyv17)) = r.DecodeBool()
		}
	}
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Dir = false
	} else {
		yyv19 := &x.Dir
		yym20 := z.DecBinary()
		_ = yym20
		if false {
		} else {
			*((*bool)(yyv19)) = r.DecodeBool()
		}
	}
	for {
		yyj12++
		if yyhl12 {
			yyb12 = yyj12 > l
		} else {
			yyb12 = r.CheckBreak()
		}
		if yyb12 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj12-1, "")
	}
	r.ReadArrayEnd()
}
func (x *Response) CodecEncodeSelf(e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			_, _ = yysep2, yy2arr2
			const yyr2 bool = false
			if yyr2 || yy2arr2 {
				r.WriteArrayStart(3)
			} else {
				r.WriteMapStart(3)
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym4 := z.EncBinary()
				_ = yym4
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Action))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("action"))
				r.WriteMapElemValue()
				yym5 := z.EncBinary()
				_ = yym5
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Action))
				}
			}
			var yyn6 bool
			if x.Node == nil {
				yyn6 = true
				goto LABEL6
			}
		LABEL6:
			if yyr2 || yy2arr2 {
				if yyn6 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.Node == nil {
						r.EncodeNil()
					} else {
						x.Node.CodecEncodeSelf(e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("node"))
				r.WriteMapElemValue()
				if yyn6 {
					r.EncodeNil()
				} else {
					if x.Node == nil {
						r.EncodeNil()
					} else {
						x.Node.CodecEncodeSelf(e)
					}
				}
			}
			var yyn9 bool
			if x.PrevNode == nil {
				yyn9 = true
				goto LABEL9
			}
		LABEL9:
			if yyr2 || yy2arr2 {
				if yyn9 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if x.PrevNode == nil {
						r.EncodeNil()
					} else {
						x.PrevNode.CodecEncodeSelf(e)
					}
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("prevNode"))
				r.WriteMapElemValue()
				if yyn9 {
					r.EncodeNil()
				} else {
					if x.PrevNode == nil {
						r.EncodeNil()
					} else {
						x.PrevNode.CodecEncodeSelf(e)
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayEnd()
			} else {
				r.WriteMapEnd()
			}
		}
	}
}
func (x *Response) CodecDecodeSelf(d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap7612 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray7612 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr7612)
		}
	}
}
func (x *Response) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer()
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		r.ReadMapElemKey()
		yys3Slc = r.DecodeStringAsBytes()
		yys3 := string(yys3Slc)
		r.ReadMapElemValue()
		switch yys3 {
		case "action":
			if r.TryDecodeAsNil() {
				x.Action = ""
			} else {
				yyv4 := &x.Action
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else {
					*((*string)(yyv4)) = r.DecodeString()
				}
			}
		case "node":
			if x.Node == nil {
				x.Node = new(Node)
			}
			if r.TryDecodeAsNil() {
				if x.Node != nil {
					x.Node = nil
				}
			} else {
				if x.Node == nil {
					x.Node = new(Node)
				}
				x.Node.CodecDecodeSelf(d)
			}
		case "prevNode":
			if x.PrevNode == nil {
				x.PrevNode = new(Node)
			}
			if r.TryDecodeAsNil() {
				if x.PrevNode != nil {
					x.PrevNode = nil
				}
			} else {
				if x.PrevNode == nil {
					x.PrevNode = new(Node)
				}
				x.PrevNode.CodecDecodeSelf(d)
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		}
	}
	r.ReadMapEnd()
}
func (x *Response) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj8 int
	var yyb8 bool
	var yyhl8 bool = l >= 0
	yyj8++
	if yyhl8 {
		yyb8 = yyj8 > l
	} else {
		yyb8 = r.CheckBreak()
	}
	if yyb8 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Action = ""
	} else {
		yyv9 := &x.Action
		yym10 := z.DecBinary()
		_ = yym10
		if false {
		} else {
			*((*string)(yyv9)) = r.DecodeString()
		}
	}
	if x.Node == nil {
		x.Node = new(Node)
	}
	yyj8++
	if yyhl8 {
		yyb8 = yyj8 > l
	} else {
		yyb8 = r.CheckBreak()
	}
	if yyb8 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.Node != nil {
			x.Node = nil
		}
	} else {
		if x.Node == nil {
			x.Node = new(Node)
		}
		x.Node.CodecDecodeSelf(d)
	}
	if x.PrevNode == nil {
		x.PrevNode = new(Node)
	}
	yyj8++
	if yyhl8 {
		yyb8 = yyj8 > l
	} else {
		yyb8 = r.CheckBreak()
	}
	if yyb8 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.PrevNode != nil {
			x.PrevNode = nil
		}
	} else {
		if x.PrevNode == nil {
			x.PrevNode = new(Node)
		}
		x.PrevNode.CodecDecodeSelf(d)
	}
	for {
		yyj8++
		if yyhl8 {
			yyb8 = yyj8 > l
		} else {
			yyb8 = r.CheckBreak()
		}
		if yyb8 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj8-1, "")
	}
	r.ReadArrayEnd()
}
func (x *Node) CodecEncodeSelf(e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			var yyq2 [8]bool
			_ = yyq2
			_, _ = yysep2, yy2arr2
			const yyr2 bool = false
			yyq2[1] = x.Dir != false
			yyq2[6] = x.Expiration != nil
			yyq2[7] = x.TTL != 0
			if yyr2 || yy2arr2 {
				r.WriteArrayStart(8)
			} else {
				var yynn2 = 5
				for _, b := range yyq2 {
					if b {
						yynn2++
					}
				}
				r.WriteMapStart(yynn2)
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym4 := z.EncBinary()
				_ = yym4
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Key))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("key"))
				r.WriteMapElemValue()
				yym5 := z.EncBinary()
				_ = yym5
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Key))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if yyq2[1] {
					yym7 := z.EncBinary()
					_ = yym7
					if false {
					} else {
						r.EncodeBool(bool(x.Dir))
					}
				} else {
					r.EncodeBool(false)
				}
			} else {
				if yyq2[1] {
					r.WriteMapElemKey()
					r.EncodeString(codecSelferC_UTF87612, string("dir"))
					r.WriteMapElemValue()
					yym8 := z.EncBinary()
					_ = yym8
					if false {
					} else {
						r.EncodeBool(bool(x.Dir))
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym10 := z.EncBinary()
				_ = yym10
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Value))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("value"))
				r.WriteMapElemValue()
				yym11 := z.EncBinary()
				_ = yym11
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Value))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if x.Nodes == nil {
					r.EncodeNil()
				} else {
					x.Nodes.CodecEncodeSelf(e)
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("nodes"))
				r.WriteMapElemValue()
				if x.Nodes == nil {
					r.EncodeNil()
				} else {
					x.Nodes.CodecEncodeSelf(e)
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym16 := z.EncBinary()
				_ = yym16
				if false {
				} else {
					r.EncodeUint(uint64(x.CreatedIndex))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("createdIndex"))
				r.WriteMapElemValue()
				yym17 := z.EncBinary()
				_ = yym17
				if false {
				} else {
					r.EncodeUint(uint64(x.CreatedIndex))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym19 := z.EncBinary()
				_ = yym19
				if false {
				} else {
					r.EncodeUint(uint64(x.ModifiedIndex))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("modifiedIndex"))
				r.WriteMapElemValue()
				yym20 := z.EncBinary()
				_ = yym20
				if false {
				} else {
					r.EncodeUint(uint64(x.ModifiedIndex))
				}
			}
			var yyn21 bool
			if x.Expiration == nil {
				yyn21 = true
				goto LABEL21
			}
		LABEL21:
			if yyr2 || yy2arr2 {
				if yyn21 {
					r.WriteArrayElem()
					r.EncodeNil()
				} else {
					r.WriteArrayElem()
					if yyq2[6] {
						if x.Expiration == nil {
							r.EncodeNil()
						} else {
							yym22 := z.EncBinary()
							_ = yym22
							if false {
							} else if yym23 := z.TimeRtidIfBinc(); yym23 != 0 {
								r.EncodeBuiltin(yym23, x.Expiration)
							} else if z.HasExtensions() && z.EncExt(x.Expiration) {
							} else if yym22 {
								z.EncBinaryMarshal(x.Expiration)
							} else if !yym22 && z.IsJSONHandle() {
								z.EncJSONMarshal(x.Expiration)
							} else {
								z.EncFallback(x.Expiration)
							}
						}
					} else {
						r.EncodeNil()
					}
				}
			} else {
				if yyq2[6] {
					r.WriteMapElemKey()
					r.EncodeString(codecSelferC_UTF87612, string("expiration"))
					r.WriteMapElemValue()
					if yyn21 {
						r.EncodeNil()
					} else {
						if x.Expiration == nil {
							r.EncodeNil()
						} else {
							yym24 := z.EncBinary()
							_ = yym24
							if false {
							} else if yym25 := z.TimeRtidIfBinc(); yym25 != 0 {
								r.EncodeBuiltin(yym25, x.Expiration)
							} else if z.HasExtensions() && z.EncExt(x.Expiration) {
							} else if yym24 {
								z.EncBinaryMarshal(x.Expiration)
							} else if !yym24 && z.IsJSONHandle() {
								z.EncJSONMarshal(x.Expiration)
							} else {
								z.EncFallback(x.Expiration)
							}
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if yyq2[7] {
					yym27 := z.EncBinary()
					_ = yym27
					if false {
					} else {
						r.EncodeInt(int64(x.TTL))
					}
				} else {
					r.EncodeInt(0)
				}
			} else {
				if yyq2[7] {
					r.WriteMapElemKey()
					r.EncodeString(codecSelferC_UTF87612, string("ttl"))
					r.WriteMapElemValue()
					yym28 := z.EncBinary()
					_ = yym28
					if false {
					} else {
						r.EncodeInt(int64(x.TTL))
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayEnd()
			} else {
				r.WriteMapEnd()
			}
		}
	}
}
func (x *Node) CodecDecodeSelf(d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap7612 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray7612 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr7612)
		}
	}
}
func (x *Node) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer()
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		r.ReadMapElemKey()
		yys3Slc = r.DecodeStringAsBytes()
		yys3 := string(yys3Slc)
		r.ReadMapElemValue()
		switch yys3 {
		case "key":
			if r.TryDecodeAsNil() {
				x.Key = ""
			} else {
				yyv4 := &x.Key
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else {
					*((*string)(yyv4)) = r.DecodeString()
				}
			}
		case "dir":
			if r.TryDecodeAsNil() {
				x.Dir = false
			} else {
				yyv6 := &x.Dir
				yym7 := z.DecBinary()
				_ = yym7
				if false {
				} else {
					*((*bool)(yyv6)) = r.DecodeBool()
				}
			}
		case "value":
			if r.TryDecodeAsNil() {
				x.Value = ""
			} else {
				yyv8 := &x.Value
				yym9 := z.DecBinary()
				_ = yym9
				if false {
				} else {
					*((*string)(yyv8)) = r.DecodeString()
				}
			}
		case "nodes":
			if r.TryDecodeAsNil() {
				x.Nodes = nil
			} else {
				yyv10 := &x.Nodes
				yyv10.CodecDecodeSelf(d)
			}
		case "createdIndex":
			if r.TryDecodeAsNil() {
				x.CreatedIndex = 0
			} else {
				yyv11 := &x.CreatedIndex
				yym12 := z.DecBinary()
				_ = yym12
				if false {
				} else {
					*((*uint64)(yyv11)) = uint64(r.DecodeUint(64))
				}
			}
		case "modifiedIndex":
			if r.TryDecodeAsNil() {
				x.ModifiedIndex = 0
			} else {
				yyv13 := &x.ModifiedIndex
				yym14 := z.DecBinary()
				_ = yym14
				if false {
				} else {
					*((*uint64)(yyv13)) = uint64(r.DecodeUint(64))
				}
			}
		case "expiration":
			if x.Expiration == nil {
				x.Expiration = new(time.Time)
			}
			if r.TryDecodeAsNil() {
				if x.Expiration != nil {
					x.Expiration = nil
				}
			} else {
				if x.Expiration == nil {
					x.Expiration = new(time.Time)
				}
				yym16 := z.DecBinary()
				_ = yym16
				if false {
				} else if yym17 := z.TimeRtidIfBinc(); yym17 != 0 {
					r.DecodeBuiltin(yym17, x.Expiration)
				} else if z.HasExtensions() && z.DecExt(x.Expiration) {
				} else if yym16 {
					z.DecBinaryUnmarshal(x.Expiration)
				} else if !yym16 && z.IsJSONHandle() {
					z.DecJSONUnmarshal(x.Expiration)
				} else {
					z.DecFallback(x.Expiration, false)
				}
			}
		case "ttl":
			if r.TryDecodeAsNil() {
				x.TTL = 0
			} else {
				yyv18 := &x.TTL
				yym19 := z.DecBinary()
				_ = yym19
				if false {
				} else {
					*((*int64)(yyv18)) = int64(r.DecodeInt(64))
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		}
	}
	r.ReadMapEnd()
}
func (x *Node) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj20 int
	var yyb20 bool
	var yyhl20 bool = l >= 0
	yyj20++
	if yyhl20 {
		yyb20 = yyj20 > l
	} else {
		yyb20 = r.CheckBreak()
	}
	if yyb20 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Key = ""
	} else {
		yyv21 := &x.Key
		yym22 := z.DecBinary()
		_ = yym22
		if false {
		} else {
			*((*string)(yyv21)) = r.DecodeString()
		}
	}
	yyj20++
	if yyhl20 {
		yyb20 = yyj20 > l
	} else {
		yyb20 = r.CheckBreak()
	}
	if yyb20 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Dir = false
	} else {
		yyv23 := &x.Dir
		yym24 := z.DecBinary()
		_ = yym24
		if false {
		} else {
			*((*bool)(yyv23)) = r.DecodeBool()
		}
	}
	yyj20++
	if yyhl20 {
		yyb20 = yyj20 > l
	} else {
		yyb20 = r.CheckBreak()
	}
	if yyb20 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Value = ""
	} else {
		yyv25 := &x.Value
		yym26 := z.DecBinary()
		_ = yym26
		if false {
		} else {
			*((*string)(yyv25)) = r.DecodeString()
		}
	}
	yyj20++
	if yyhl20 {
		yyb20 = yyj20 > l
	} else {
		yyb20 = r.CheckBreak()
	}
	if yyb20 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Nodes = nil
	} else {
		yyv27 := &x.Nodes
		yyv27.CodecDecodeSelf(d)
	}
	yyj20++
	if yyhl20 {
		yyb20 = yyj20 > l
	} else {
		yyb20 = r.CheckBreak()
	}
	if yyb20 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.CreatedIndex = 0
	} else {
		yyv28 := &x.CreatedIndex
		yym29 := z.DecBinary()
		_ = yym29
		if false {
		} else {
			*((*uint64)(yyv28)) = uint64(r.DecodeUint(64))
		}
	}
	yyj20++
	if yyhl20 {
		yyb20 = yyj20 > l
	} else {
		yyb20 = r.CheckBreak()
	}
	if yyb20 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.ModifiedIndex = 0
	} else {
		yyv30 := &x.ModifiedIndex
		yym31 := z.DecBinary()
		_ = yym31
		if false {
		} else {
			*((*uint64)(yyv30)) = uint64(r.DecodeUint(64))
		}
	}
	if x.Expiration == nil {
		x.Expiration = new(time.Time)
	}
	yyj20++
	if yyhl20 {
		yyb20 = yyj20 > l
	} else {
		yyb20 = r.CheckBreak()
	}
	if yyb20 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if x.Expiration != nil {
			x.Expiration = nil
		}
	} else {
		if x.Expiration == nil {
			x.Expiration = new(time.Time)
		}
		yym33 := z.DecBinary()
		_ = yym33
		if false {
		} else if yym34 := z.TimeRtidIfBinc(); yym34 != 0 {
			r.DecodeBuiltin(yym34, x.Expiration)
		} else if z.HasExtensions() && z.DecExt(x.Expiration) {
		} else if yym33 {
			z.DecBinaryUnmarshal(x.Expiration)
		} else if !yym33 && z.IsJSONHandle() {
			z.DecJSONUnmarshal(x.Expiration)
		} else {
			z.DecFallback(x.Expiration, false)
		}
	}
	yyj20++
	if yyhl20 {
		yyb20 = yyj20 > l
	} else {
		yyb20 = r.CheckBreak()
	}
	if yyb20 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.TTL = 0
	} else {
		yyv35 := &x.TTL
		yym36 := z.DecBinary()
		_ = yym36
		if false {
		} else {
			*((*int64)(yyv35)) = int64(r.DecodeInt(64))
		}
	}
	for {
		yyj20++
		if yyhl20 {
			yyb20 = yyj20 > l
		} else {
			yyb20 = r.CheckBreak()
		}
		if yyb20 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj20-1, "")
	}
	r.ReadArrayEnd()
}
func (x Nodes) CodecEncodeSelf(e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			h.encNodes((Nodes)(x), e)
		}
	}
}
func (x *Nodes) CodecDecodeSelf(d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		h.decNodes((*Nodes)(x), d)
	}
}
func (x *httpKeysAPI) CodecEncodeSelf(e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			_, _ = yysep2, yy2arr2
			const yyr2 bool = false
			if yyr2 || yy2arr2 {
				r.WriteArrayStart(0)
			} else {
				r.WriteMapStart(0)
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayEnd()
			} else {
				r.WriteMapEnd()
			}
		}
	}
}
func (x *httpKeysAPI) CodecDecodeSelf(d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap7612 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray7612 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr7612)
		}
	}
}
func (x *httpKeysAPI) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer()
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		r.ReadMapElemKey()
		yys3Slc = r.DecodeStringAsBytes()
		yys3 := string(yys3Slc)
		r.ReadMapElemValue()
		switch yys3 {
		default:
			z.DecStructFieldNotFound(-1, yys3)
		}
	}
	r.ReadMapEnd()
}
func (x *httpKeysAPI) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj4 int
	var yyb4 bool
	var yyhl4 bool = l >= 0
	for {
		yyj4++
		if yyhl4 {
			yyb4 = yyj4 > l
		} else {
			yyb4 = r.CheckBreak()
		}
		if yyb4 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj4-1, "")
	}
	r.ReadArrayEnd()
}
func (x *httpWatcher) CodecEncodeSelf(e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			_, _ = yysep2, yy2arr2
			const yyr2 bool = false
			if yyr2 || yy2arr2 {
				r.WriteArrayStart(0)
			} else {
				r.WriteMapStart(0)
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayEnd()
			} else {
				r.WriteMapEnd()
			}
		}
	}
}
func (x *httpWatcher) CodecDecodeSelf(d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap7612 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray7612 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr7612)
		}
	}
}
func (x *httpWatcher) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer()
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		r.ReadMapElemKey()
		yys3Slc = r.DecodeStringAsBytes()
		yys3 := string(yys3Slc)
		r.ReadMapElemValue()
		switch yys3 {
		default:
			z.DecStructFieldNotFound(-1, yys3)
		}
	}
	r.ReadMapEnd()
}
func (x *httpWatcher) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj4 int
	var yyb4 bool
	var yyhl4 bool = l >= 0
	for {
		yyj4++
		if yyhl4 {
			yyb4 = yyj4 > l
		} else {
			yyb4 = r.CheckBreak()
		}
		if yyb4 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj4-1, "")
	}
	r.ReadArrayEnd()
}
func (x *getAction) CodecEncodeSelf(e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			_, _ = yysep2, yy2arr2
			const yyr2 bool = false
			if yyr2 || yy2arr2 {
				r.WriteArrayStart(5)
			} else {
				r.WriteMapStart(5)
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym4 := z.EncBinary()
				_ = yym4
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Prefix))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("Prefix"))
				r.WriteMapElemValue()
				yym5 := z.EncBinary()
				_ = yym5
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Prefix))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym7 := z.EncBinary()
				_ = yym7
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Key))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("Key"))
				r.WriteMapElemValue()
				yym8 := z.EncBinary()
				_ = yym8
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Key))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym10 := z.EncBinary()
				_ = yym10
				if false {
				} else {
					r.EncodeBool(bool(x.Recursive))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("Recursive"))
				r.WriteMapElemValue()
				yym11 := z.EncBinary()
				_ = yym11
				if false {
				} else {
					r.EncodeBool(bool(x.Recursive))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym13 := z.EncBinary()
				_ = yym13
				if false {
				} else {
					r.EncodeBool(bool(x.Sorted))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("Sorted"))
				r.WriteMapElemValue()
				yym14 := z.EncBinary()
				_ = yym14
				if false {
				} else {
					r.EncodeBool(bool(x.Sorted))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym16 := z.EncBinary()
				_ = yym16
				if false {
				} else {
					r.EncodeBool(bool(x.Quorum))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("Quorum"))
				r.WriteMapElemValue()
				yym17 := z.EncBinary()
				_ = yym17
				if false {
				} else {
					r.EncodeBool(bool(x.Quorum))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayEnd()
			} else {
				r.WriteMapEnd()
			}
		}
	}
}
func (x *getAction) CodecDecodeSelf(d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap7612 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray7612 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr7612)
		}
	}
}
func (x *getAction) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer()
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		r.ReadMapElemKey()
		yys3Slc = r.DecodeStringAsBytes()
		yys3 := string(yys3Slc)
		r.ReadMapElemValue()
		switch yys3 {
		case "Prefix":
			if r.TryDecodeAsNil() {
				x.Prefix = ""
			} else {
				yyv4 := &x.Prefix
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else {
					*((*string)(yyv4)) = r.DecodeString()
				}
			}
		case "Key":
			if r.TryDecodeAsNil() {
				x.Key = ""
			} else {
				yyv6 := &x.Key
				yym7 := z.DecBinary()
				_ = yym7
				if false {
				} else {
					*((*string)(yyv6)) = r.DecodeString()
				}
			}
		case "Recursive":
			if r.TryDecodeAsNil() {
				x.Recursive = false
			} else {
				yyv8 := &x.Recursive
				yym9 := z.DecBinary()
				_ = yym9
				if false {
				} else {
					*((*bool)(yyv8)) = r.DecodeBool()
				}
			}
		case "Sorted":
			if r.TryDecodeAsNil() {
				x.Sorted = false
			} else {
				yyv10 := &x.Sorted
				yym11 := z.DecBinary()
				_ = yym11
				if false {
				} else {
					*((*bool)(yyv10)) = r.DecodeBool()
				}
			}
		case "Quorum":
			if r.TryDecodeAsNil() {
				x.Quorum = false
			} else {
				yyv12 := &x.Quorum
				yym13 := z.DecBinary()
				_ = yym13
				if false {
				} else {
					*((*bool)(yyv12)) = r.DecodeBool()
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		}
	}
	r.ReadMapEnd()
}
func (x *getAction) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj14 int
	var yyb14 bool
	var yyhl14 bool = l >= 0
	yyj14++
	if yyhl14 {
		yyb14 = yyj14 > l
	} else {
		yyb14 = r.CheckBreak()
	}
	if yyb14 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Prefix = ""
	} else {
		yyv15 := &x.Prefix
		yym16 := z.DecBinary()
		_ = yym16
		if false {
		} else {
			*((*string)(yyv15)) = r.DecodeString()
		}
	}
	yyj14++
	if yyhl14 {
		yyb14 = yyj14 > l
	} else {
		yyb14 = r.CheckBreak()
	}
	if yyb14 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Key = ""
	} else {
		yyv17 := &x.Key
		yym18 := z.DecBinary()
		_ = yym18
		if false {
		} else {
			*((*string)(yyv17)) = r.DecodeString()
		}
	}
	yyj14++
	if yyhl14 {
		yyb14 = yyj14 > l
	} else {
		yyb14 = r.CheckBreak()
	}
	if yyb14 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Recursive = false
	} else {
		yyv19 := &x.Recursive
		yym20 := z.DecBinary()
		_ = yym20
		if false {
		} else {
			*((*bool)(yyv19)) = r.DecodeBool()
		}
	}
	yyj14++
	if yyhl14 {
		yyb14 = yyj14 > l
	} else {
		yyb14 = r.CheckBreak()
	}
	if yyb14 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Sorted = false
	} else {
		yyv21 := &x.Sorted
		yym22 := z.DecBinary()
		_ = yym22
		if false {
		} else {
			*((*bool)(yyv21)) = r.DecodeBool()
		}
	}
	yyj14++
	if yyhl14 {
		yyb14 = yyj14 > l
	} else {
		yyb14 = r.CheckBreak()
	}
	if yyb14 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Quorum = false
	} else {
		yyv23 := &x.Quorum
		yym24 := z.DecBinary()
		_ = yym24
		if false {
		} else {
			*((*bool)(yyv23)) = r.DecodeBool()
		}
	}
	for {
		yyj14++
		if yyhl14 {
			yyb14 = yyj14 > l
		} else {
			yyb14 = r.CheckBreak()
		}
		if yyb14 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj14-1, "")
	}
	r.ReadArrayEnd()
}
func (x *waitAction) CodecEncodeSelf(e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			_, _ = yysep2, yy2arr2
			const yyr2 bool = false
			if yyr2 || yy2arr2 {
				r.WriteArrayStart(4)
			} else {
				r.WriteMapStart(4)
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym4 := z.EncBinary()
				_ = yym4
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Prefix))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("Prefix"))
				r.WriteMapElemValue()
				yym5 := z.EncBinary()
				_ = yym5
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Prefix))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym7 := z.EncBinary()
				_ = yym7
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Key))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("Key"))
				r.WriteMapElemValue()
				yym8 := z.EncBinary()
				_ = yym8
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Key))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym10 := z.EncBinary()
				_ = yym10
				if false {
				} else {
					r.EncodeUint(uint64(x.WaitIndex))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("WaitIndex"))
				r.WriteMapElemValue()
				yym11 := z.EncBinary()
				_ = yym11
				if false {
				} else {
					r.EncodeUint(uint64(x.WaitIndex))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym13 := z.EncBinary()
				_ = yym13
				if false {
				} else {
					r.EncodeBool(bool(x.Recursive))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("Recursive"))
				r.WriteMapElemValue()
				yym14 := z.EncBinary()
				_ = yym14
				if false {
				} else {
					r.EncodeBool(bool(x.Recursive))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayEnd()
			} else {
				r.WriteMapEnd()
			}
		}
	}
}
func (x *waitAction) CodecDecodeSelf(d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap7612 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray7612 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr7612)
		}
	}
}
func (x *waitAction) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer()
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		r.ReadMapElemKey()
		yys3Slc = r.DecodeStringAsBytes()
		yys3 := string(yys3Slc)
		r.ReadMapElemValue()
		switch yys3 {
		case "Prefix":
			if r.TryDecodeAsNil() {
				x.Prefix = ""
			} else {
				yyv4 := &x.Prefix
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else {
					*((*string)(yyv4)) = r.DecodeString()
				}
			}
		case "Key":
			if r.TryDecodeAsNil() {
				x.Key = ""
			} else {
				yyv6 := &x.Key
				yym7 := z.DecBinary()
				_ = yym7
				if false {
				} else {
					*((*string)(yyv6)) = r.DecodeString()
				}
			}
		case "WaitIndex":
			if r.TryDecodeAsNil() {
				x.WaitIndex = 0
			} else {
				yyv8 := &x.WaitIndex
				yym9 := z.DecBinary()
				_ = yym9
				if false {
				} else {
					*((*uint64)(yyv8)) = uint64(r.DecodeUint(64))
				}
			}
		case "Recursive":
			if r.TryDecodeAsNil() {
				x.Recursive = false
			} else {
				yyv10 := &x.Recursive
				yym11 := z.DecBinary()
				_ = yym11
				if false {
				} else {
					*((*bool)(yyv10)) = r.DecodeBool()
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		}
	}
	r.ReadMapEnd()
}
func (x *waitAction) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj12 int
	var yyb12 bool
	var yyhl12 bool = l >= 0
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Prefix = ""
	} else {
		yyv13 := &x.Prefix
		yym14 := z.DecBinary()
		_ = yym14
		if false {
		} else {
			*((*string)(yyv13)) = r.DecodeString()
		}
	}
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Key = ""
	} else {
		yyv15 := &x.Key
		yym16 := z.DecBinary()
		_ = yym16
		if false {
		} else {
			*((*string)(yyv15)) = r.DecodeString()
		}
	}
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.WaitIndex = 0
	} else {
		yyv17 := &x.WaitIndex
		yym18 := z.DecBinary()
		_ = yym18
		if false {
		} else {
			*((*uint64)(yyv17)) = uint64(r.DecodeUint(64))
		}
	}
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Recursive = false
	} else {
		yyv19 := &x.Recursive
		yym20 := z.DecBinary()
		_ = yym20
		if false {
		} else {
			*((*bool)(yyv19)) = r.DecodeBool()
		}
	}
	for {
		yyj12++
		if yyhl12 {
			yyb12 = yyj12 > l
		} else {
			yyb12 = r.CheckBreak()
		}
		if yyb12 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj12-1, "")
	}
	r.ReadArrayEnd()
}
func (x *setAction) CodecEncodeSelf(e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			_, _ = yysep2, yy2arr2
			const yyr2 bool = false
			if yyr2 || yy2arr2 {
				r.WriteArrayStart(10)
			} else {
				r.WriteMapStart(10)
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym4 := z.EncBinary()
				_ = yym4
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Prefix))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("Prefix"))
				r.WriteMapElemValue()
				yym5 := z.EncBinary()
				_ = yym5
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Prefix))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym7 := z.EncBinary()
				_ = yym7
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Key))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("Key"))
				r.WriteMapElemValue()
				yym8 := z.EncBinary()
				_ = yym8
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Key))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym10 := z.EncBinary()
				_ = yym10
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Value))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("Value"))
				r.WriteMapElemValue()
				yym11 := z.EncBinary()
				_ = yym11
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Value))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym13 := z.EncBinary()
				_ = yym13
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.PrevValue))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("PrevValue"))
				r.WriteMapElemValue()
				yym14 := z.EncBinary()
				_ = yym14
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.PrevValue))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym16 := z.EncBinary()
				_ = yym16
				if false {
				} else {
					r.EncodeUint(uint64(x.PrevIndex))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("PrevIndex"))
				r.WriteMapElemValue()
				yym17 := z.EncBinary()
				_ = yym17
				if false {
				} else {
					r.EncodeUint(uint64(x.PrevIndex))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				x.PrevExist.CodecEncodeSelf(e)
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("PrevExist"))
				r.WriteMapElemValue()
				x.PrevExist.CodecEncodeSelf(e)
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym22 := z.EncBinary()
				_ = yym22
				if false {
				} else if z.HasExtensions() && z.EncExt(x.TTL) {
				} else {
					r.EncodeInt(int64(x.TTL))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("TTL"))
				r.WriteMapElemValue()
				yym23 := z.EncBinary()
				_ = yym23
				if false {
				} else if z.HasExtensions() && z.EncExt(x.TTL) {
				} else {
					r.EncodeInt(int64(x.TTL))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym25 := z.EncBinary()
				_ = yym25
				if false {
				} else {
					r.EncodeBool(bool(x.Refresh))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("Refresh"))
				r.WriteMapElemValue()
				yym26 := z.EncBinary()
				_ = yym26
				if false {
				} else {
					r.EncodeBool(bool(x.Refresh))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym28 := z.EncBinary()
				_ = yym28
				if false {
				} else {
					r.EncodeBool(bool(x.Dir))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("Dir"))
				r.WriteMapElemValue()
				yym29 := z.EncBinary()
				_ = yym29
				if false {
				} else {
					r.EncodeBool(bool(x.Dir))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym31 := z.EncBinary()
				_ = yym31
				if false {
				} else {
					r.EncodeBool(bool(x.NoValueOnSuccess))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("NoValueOnSuccess"))
				r.WriteMapElemValue()
				yym32 := z.EncBinary()
				_ = yym32
				if false {
				} else {
					r.EncodeBool(bool(x.NoValueOnSuccess))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayEnd()
			} else {
				r.WriteMapEnd()
			}
		}
	}
}
func (x *setAction) CodecDecodeSelf(d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap7612 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray7612 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr7612)
		}
	}
}
func (x *setAction) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer()
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		r.ReadMapElemKey()
		yys3Slc = r.DecodeStringAsBytes()
		yys3 := string(yys3Slc)
		r.ReadMapElemValue()
		switch yys3 {
		case "Prefix":
			if r.TryDecodeAsNil() {
				x.Prefix = ""
			} else {
				yyv4 := &x.Prefix
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else {
					*((*string)(yyv4)) = r.DecodeString()
				}
			}
		case "Key":
			if r.TryDecodeAsNil() {
				x.Key = ""
			} else {
				yyv6 := &x.Key
				yym7 := z.DecBinary()
				_ = yym7
				if false {
				} else {
					*((*string)(yyv6)) = r.DecodeString()
				}
			}
		case "Value":
			if r.TryDecodeAsNil() {
				x.Value = ""
			} else {
				yyv8 := &x.Value
				yym9 := z.DecBinary()
				_ = yym9
				if false {
				} else {
					*((*string)(yyv8)) = r.DecodeString()
				}
			}
		case "PrevValue":
			if r.TryDecodeAsNil() {
				x.PrevValue = ""
			} else {
				yyv10 := &x.PrevValue
				yym11 := z.DecBinary()
				_ = yym11
				if false {
				} else {
					*((*string)(yyv10)) = r.DecodeString()
				}
			}
		case "PrevIndex":
			if r.TryDecodeAsNil() {
				x.PrevIndex = 0
			} else {
				yyv12 := &x.PrevIndex
				yym13 := z.DecBinary()
				_ = yym13
				if false {
				} else {
					*((*uint64)(yyv12)) = uint64(r.DecodeUint(64))
				}
			}
		case "PrevExist":
			if r.TryDecodeAsNil() {
				x.PrevExist = ""
			} else {
				yyv14 := &x.PrevExist
				yyv14.CodecDecodeSelf(d)
			}
		case "TTL":
			if r.TryDecodeAsNil() {
				x.TTL = 0
			} else {
				yyv15 := &x.TTL
				yym16 := z.DecBinary()
				_ = yym16
				if false {
				} else if z.HasExtensions() && z.DecExt(yyv15) {
				} else {
					*((*int64)(yyv15)) = int64(r.DecodeInt(64))
				}
			}
		case "Refresh":
			if r.TryDecodeAsNil() {
				x.Refresh = false
			} else {
				yyv17 := &x.Refresh
				yym18 := z.DecBinary()
				_ = yym18
				if false {
				} else {
					*((*bool)(yyv17)) = r.DecodeBool()
				}
			}
		case "Dir":
			if r.TryDecodeAsNil() {
				x.Dir = false
			} else {
				yyv19 := &x.Dir
				yym20 := z.DecBinary()
				_ = yym20
				if false {
				} else {
					*((*bool)(yyv19)) = r.DecodeBool()
				}
			}
		case "NoValueOnSuccess":
			if r.TryDecodeAsNil() {
				x.NoValueOnSuccess = false
			} else {
				yyv21 := &x.NoValueOnSuccess
				yym22 := z.DecBinary()
				_ = yym22
				if false {
				} else {
					*((*bool)(yyv21)) = r.DecodeBool()
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		}
	}
	r.ReadMapEnd()
}
func (x *setAction) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj23 int
	var yyb23 bool
	var yyhl23 bool = l >= 0
	yyj23++
	if yyhl23 {
		yyb23 = yyj23 > l
	} else {
		yyb23 = r.CheckBreak()
	}
	if yyb23 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Prefix = ""
	} else {
		yyv24 := &x.Prefix
		yym25 := z.DecBinary()
		_ = yym25
		if false {
		} else {
			*((*string)(yyv24)) = r.DecodeString()
		}
	}
	yyj23++
	if yyhl23 {
		yyb23 = yyj23 > l
	} else {
		yyb23 = r.CheckBreak()
	}
	if yyb23 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Key = ""
	} else {
		yyv26 := &x.Key
		yym27 := z.DecBinary()
		_ = yym27
		if false {
		} else {
			*((*string)(yyv26)) = r.DecodeString()
		}
	}
	yyj23++
	if yyhl23 {
		yyb23 = yyj23 > l
	} else {
		yyb23 = r.CheckBreak()
	}
	if yyb23 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Value = ""
	} else {
		yyv28 := &x.Value
		yym29 := z.DecBinary()
		_ = yym29
		if false {
		} else {
			*((*string)(yyv28)) = r.DecodeString()
		}
	}
	yyj23++
	if yyhl23 {
		yyb23 = yyj23 > l
	} else {
		yyb23 = r.CheckBreak()
	}
	if yyb23 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.PrevValue = ""
	} else {
		yyv30 := &x.PrevValue
		yym31 := z.DecBinary()
		_ = yym31
		if false {
		} else {
			*((*string)(yyv30)) = r.DecodeString()
		}
	}
	yyj23++
	if yyhl23 {
		yyb23 = yyj23 > l
	} else {
		yyb23 = r.CheckBreak()
	}
	if yyb23 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.PrevIndex = 0
	} else {
		yyv32 := &x.PrevIndex
		yym33 := z.DecBinary()
		_ = yym33
		if false {
		} else {
			*((*uint64)(yyv32)) = uint64(r.DecodeUint(64))
		}
	}
	yyj23++
	if yyhl23 {
		yyb23 = yyj23 > l
	} else {
		yyb23 = r.CheckBreak()
	}
	if yyb23 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.PrevExist = ""
	} else {
		yyv34 := &x.PrevExist
		yyv34.CodecDecodeSelf(d)
	}
	yyj23++
	if yyhl23 {
		yyb23 = yyj23 > l
	} else {
		yyb23 = r.CheckBreak()
	}
	if yyb23 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.TTL = 0
	} else {
		yyv35 := &x.TTL
		yym36 := z.DecBinary()
		_ = yym36
		if false {
		} else if z.HasExtensions() && z.DecExt(yyv35) {
		} else {
			*((*int64)(yyv35)) = int64(r.DecodeInt(64))
		}
	}
	yyj23++
	if yyhl23 {
		yyb23 = yyj23 > l
	} else {
		yyb23 = r.CheckBreak()
	}
	if yyb23 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Refresh = false
	} else {
		yyv37 := &x.Refresh
		yym38 := z.DecBinary()
		_ = yym38
		if false {
		} else {
			*((*bool)(yyv37)) = r.DecodeBool()
		}
	}
	yyj23++
	if yyhl23 {
		yyb23 = yyj23 > l
	} else {
		yyb23 = r.CheckBreak()
	}
	if yyb23 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Dir = false
	} else {
		yyv39 := &x.Dir
		yym40 := z.DecBinary()
		_ = yym40
		if false {
		} else {
			*((*bool)(yyv39)) = r.DecodeBool()
		}
	}
	yyj23++
	if yyhl23 {
		yyb23 = yyj23 > l
	} else {
		yyb23 = r.CheckBreak()
	}
	if yyb23 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.NoValueOnSuccess = false
	} else {
		yyv41 := &x.NoValueOnSuccess
		yym42 := z.DecBinary()
		_ = yym42
		if false {
		} else {
			*((*bool)(yyv41)) = r.DecodeBool()
		}
	}
	for {
		yyj23++
		if yyhl23 {
			yyb23 = yyj23 > l
		} else {
			yyb23 = r.CheckBreak()
		}
		if yyb23 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj23-1, "")
	}
	r.ReadArrayEnd()
}
func (x *deleteAction) CodecEncodeSelf(e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			_, _ = yysep2, yy2arr2
			const yyr2 bool = false
			if yyr2 || yy2arr2 {
				r.WriteArrayStart(6)
			} else {
				r.WriteMapStart(6)
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym4 := z.EncBinary()
				_ = yym4
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Prefix))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("Prefix"))
				r.WriteMapElemValue()
				yym5 := z.EncBinary()
				_ = yym5
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Prefix))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym7 := z.EncBinary()
				_ = yym7
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Key))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("Key"))
				r.WriteMapElemValue()
				yym8 := z.EncBinary()
				_ = yym8
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Key))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym10 := z.EncBinary()
				_ = yym10
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.PrevValue))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("PrevValue"))
				r.WriteMapElemValue()
				yym11 := z.EncBinary()
				_ = yym11
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.PrevValue))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym13 := z.EncBinary()
				_ = yym13
				if false {
				} else {
					r.EncodeUint(uint64(x.PrevIndex))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("PrevIndex"))
				r.WriteMapElemValue()
				yym14 := z.EncBinary()
				_ = yym14
				if false {
				} else {
					r.EncodeUint(uint64(x.PrevIndex))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym16 := z.EncBinary()
				_ = yym16
				if false {
				} else {
					r.EncodeBool(bool(x.Dir))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("Dir"))
				r.WriteMapElemValue()
				yym17 := z.EncBinary()
				_ = yym17
				if false {
				} else {
					r.EncodeBool(bool(x.Dir))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym19 := z.EncBinary()
				_ = yym19
				if false {
				} else {
					r.EncodeBool(bool(x.Recursive))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("Recursive"))
				r.WriteMapElemValue()
				yym20 := z.EncBinary()
				_ = yym20
				if false {
				} else {
					r.EncodeBool(bool(x.Recursive))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayEnd()
			} else {
				r.WriteMapEnd()
			}
		}
	}
}
func (x *deleteAction) CodecDecodeSelf(d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap7612 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray7612 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr7612)
		}
	}
}
func (x *deleteAction) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer()
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		r.ReadMapElemKey()
		yys3Slc = r.DecodeStringAsBytes()
		yys3 := string(yys3Slc)
		r.ReadMapElemValue()
		switch yys3 {
		case "Prefix":
			if r.TryDecodeAsNil() {
				x.Prefix = ""
			} else {
				yyv4 := &x.Prefix
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else {
					*((*string)(yyv4)) = r.DecodeString()
				}
			}
		case "Key":
			if r.TryDecodeAsNil() {
				x.Key = ""
			} else {
				yyv6 := &x.Key
				yym7 := z.DecBinary()
				_ = yym7
				if false {
				} else {
					*((*string)(yyv6)) = r.DecodeString()
				}
			}
		case "PrevValue":
			if r.TryDecodeAsNil() {
				x.PrevValue = ""
			} else {
				yyv8 := &x.PrevValue
				yym9 := z.DecBinary()
				_ = yym9
				if false {
				} else {
					*((*string)(yyv8)) = r.DecodeString()
				}
			}
		case "PrevIndex":
			if r.TryDecodeAsNil() {
				x.PrevIndex = 0
			} else {
				yyv10 := &x.PrevIndex
				yym11 := z.DecBinary()
				_ = yym11
				if false {
				} else {
					*((*uint64)(yyv10)) = uint64(r.DecodeUint(64))
				}
			}
		case "Dir":
			if r.TryDecodeAsNil() {
				x.Dir = false
			} else {
				yyv12 := &x.Dir
				yym13 := z.DecBinary()
				_ = yym13
				if false {
				} else {
					*((*bool)(yyv12)) = r.DecodeBool()
				}
			}
		case "Recursive":
			if r.TryDecodeAsNil() {
				x.Recursive = false
			} else {
				yyv14 := &x.Recursive
				yym15 := z.DecBinary()
				_ = yym15
				if false {
				} else {
					*((*bool)(yyv14)) = r.DecodeBool()
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		}
	}
	r.ReadMapEnd()
}
func (x *deleteAction) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj16 int
	var yyb16 bool
	var yyhl16 bool = l >= 0
	yyj16++
	if yyhl16 {
		yyb16 = yyj16 > l
	} else {
		yyb16 = r.CheckBreak()
	}
	if yyb16 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Prefix = ""
	} else {
		yyv17 := &x.Prefix
		yym18 := z.DecBinary()
		_ = yym18
		if false {
		} else {
			*((*string)(yyv17)) = r.DecodeString()
		}
	}
	yyj16++
	if yyhl16 {
		yyb16 = yyj16 > l
	} else {
		yyb16 = r.CheckBreak()
	}
	if yyb16 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Key = ""
	} else {
		yyv19 := &x.Key
		yym20 := z.DecBinary()
		_ = yym20
		if false {
		} else {
			*((*string)(yyv19)) = r.DecodeString()
		}
	}
	yyj16++
	if yyhl16 {
		yyb16 = yyj16 > l
	} else {
		yyb16 = r.CheckBreak()
	}
	if yyb16 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.PrevValue = ""
	} else {
		yyv21 := &x.PrevValue
		yym22 := z.DecBinary()
		_ = yym22
		if false {
		} else {
			*((*string)(yyv21)) = r.DecodeString()
		}
	}
	yyj16++
	if yyhl16 {
		yyb16 = yyj16 > l
	} else {
		yyb16 = r.CheckBreak()
	}
	if yyb16 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.PrevIndex = 0
	} else {
		yyv23 := &x.PrevIndex
		yym24 := z.DecBinary()
		_ = yym24
		if false {
		} else {
			*((*uint64)(yyv23)) = uint64(r.DecodeUint(64))
		}
	}
	yyj16++
	if yyhl16 {
		yyb16 = yyj16 > l
	} else {
		yyb16 = r.CheckBreak()
	}
	if yyb16 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Dir = false
	} else {
		yyv25 := &x.Dir
		yym26 := z.DecBinary()
		_ = yym26
		if false {
		} else {
			*((*bool)(yyv25)) = r.DecodeBool()
		}
	}
	yyj16++
	if yyhl16 {
		yyb16 = yyj16 > l
	} else {
		yyb16 = r.CheckBreak()
	}
	if yyb16 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Recursive = false
	} else {
		yyv27 := &x.Recursive
		yym28 := z.DecBinary()
		_ = yym28
		if false {
		} else {
			*((*bool)(yyv27)) = r.DecodeBool()
		}
	}
	for {
		yyj16++
		if yyhl16 {
			yyb16 = yyj16 > l
		} else {
			yyb16 = r.CheckBreak()
		}
		if yyb16 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj16-1, "")
	}
	r.ReadArrayEnd()
}
func (x *createInOrderAction) CodecEncodeSelf(e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			_, _ = yysep2, yy2arr2
			const yyr2 bool = false
			if yyr2 || yy2arr2 {
				r.WriteArrayStart(4)
			} else {
				r.WriteMapStart(4)
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym4 := z.EncBinary()
				_ = yym4
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Prefix))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("Prefix"))
				r.WriteMapElemValue()
				yym5 := z.EncBinary()
				_ = yym5
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Prefix))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym7 := z.EncBinary()
				_ = yym7
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Dir))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("Dir"))
				r.WriteMapElemValue()
				yym8 := z.EncBinary()
				_ = yym8
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Dir))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym10 := z.EncBinary()
				_ = yym10
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Value))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("Value"))
				r.WriteMapElemValue()
				yym11 := z.EncBinary()
				_ = yym11
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF87612, string(x.Value))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				yym13 := z.EncBinary()
				_ = yym13
				if false {
				} else if z.HasExtensions() && z.EncExt(x.TTL) {
				} else {
					r.EncodeInt(int64(x.TTL))
				}
			} else {
				r.WriteMapElemKey()
				r.EncodeString(codecSelferC_UTF87612, string("TTL"))
				r.WriteMapElemValue()
				yym14 := z.EncBinary()
				_ = yym14
				if false {
				} else if z.HasExtensions() && z.EncExt(x.TTL) {
				} else {
					r.EncodeInt(int64(x.TTL))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayEnd()
			} else {
				r.WriteMapEnd()
			}
		}
	}
}
func (x *createInOrderAction) CodecDecodeSelf(d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap7612 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray7612 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr7612)
		}
	}
}
func (x *createInOrderAction) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer()
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		r.ReadMapElemKey()
		yys3Slc = r.DecodeStringAsBytes()
		yys3 := string(yys3Slc)
		r.ReadMapElemValue()
		switch yys3 {
		case "Prefix":
			if r.TryDecodeAsNil() {
				x.Prefix = ""
			} else {
				yyv4 := &x.Prefix
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else {
					*((*string)(yyv4)) = r.DecodeString()
				}
			}
		case "Dir":
			if r.TryDecodeAsNil() {
				x.Dir = ""
			} else {
				yyv6 := &x.Dir
				yym7 := z.DecBinary()
				_ = yym7
				if false {
				} else {
					*((*string)(yyv6)) = r.DecodeString()
				}
			}
		case "Value":
			if r.TryDecodeAsNil() {
				x.Value = ""
			} else {
				yyv8 := &x.Value
				yym9 := z.DecBinary()
				_ = yym9
				if false {
				} else {
					*((*string)(yyv8)) = r.DecodeString()
				}
			}
		case "TTL":
			if r.TryDecodeAsNil() {
				x.TTL = 0
			} else {
				yyv10 := &x.TTL
				yym11 := z.DecBinary()
				_ = yym11
				if false {
				} else if z.HasExtensions() && z.DecExt(yyv10) {
				} else {
					*((*int64)(yyv10)) = int64(r.DecodeInt(64))
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		}
	}
	r.ReadMapEnd()
}
func (x *createInOrderAction) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj12 int
	var yyb12 bool
	var yyhl12 bool = l >= 0
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Prefix = ""
	} else {
		yyv13 := &x.Prefix
		yym14 := z.DecBinary()
		_ = yym14
		if false {
		} else {
			*((*string)(yyv13)) = r.DecodeString()
		}
	}
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Dir = ""
	} else {
		yyv15 := &x.Dir
		yym16 := z.DecBinary()
		_ = yym16
		if false {
		} else {
			*((*string)(yyv15)) = r.DecodeString()
		}
	}
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Value = ""
	} else {
		yyv17 := &x.Value
		yym18 := z.DecBinary()
		_ = yym18
		if false {
		} else {
			*((*string)(yyv17)) = r.DecodeString()
		}
	}
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.TTL = 0
	} else {
		yyv19 := &x.TTL
		yym20 := z.DecBinary()
		_ = yym20
		if false {
		} else if z.HasExtensions() && z.DecExt(yyv19) {
		} else {
			*((*int64)(yyv19)) = int64(r.DecodeInt(64))
		}
	}
	for {
		yyj12++
		if yyhl12 {
			yyb12 = yyj12 > l
		} else {
			yyb12 = r.CheckBreak()
		}
		if yyb12 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj12-1, "")
	}
	r.ReadArrayEnd()
}
func (x codecSelfer7612) encNodes(v Nodes, e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.WriteArrayStart(len(v))
	for _, yyv1 := range v {
		r.WriteArrayElem()
		if yyv1 == nil {
			r.EncodeNil()
		} else {
			yyv1.CodecEncodeSelf(e)
		}
	}
	r.WriteArrayEnd()
}
func (x codecSelfer7612) decNodes(v *Nodes, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer7612
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yyv1 := *v
	yyh1, yyl1 := z.DecSliceHelperStart()
	var yyc1 bool
	_ = yyc1
	if yyl1 == 0 {
		if yyv1 == nil {
			yyv1 = []*Node{}
			yyc1 = true
		} else if len(yyv1) != 0 {
			yyv1 = yyv1[:0]
			yyc1 = true
		}
	} else {
		yyhl1 := yyl1 > 0
		var yyrl1 int
		_ = yyrl1
		if yyhl1 {
			if yyl1 > cap(yyv1) {
				yyrl1 = z.DecInferLen(yyl1, z.DecBasicHandle().MaxInitLen, 8)
				if yyrl1 <= cap(yyv1) {
					yyv1 = yyv1[:yyrl1]
				} else {
					yyv1 = make([]*Node, yyrl1)
				}
				yyc1 = true
			} else if yyl1 != len(yyv1) {
				yyv1 = yyv1[:yyl1]
				yyc1 = true
			}
		}
		var yyj1 int
		for ; (yyhl1 && yyj1 < yyl1) || !(yyhl1 || r.CheckBreak()); yyj1++ {
			if yyj1 == 0 && len(yyv1) == 0 {
				if yyhl1 {
					yyrl1 = z.DecInferLen(yyl1, z.DecBasicHandle().MaxInitLen, 8)
				} else {
					yyrl1 = 8
				}
				yyv1 = make([]*Node, yyrl1)
				yyc1 = true
			}
			yyh1.ElemContainerState(yyj1)
			var yydb1 bool
			if yyj1 >= len(yyv1) {
				yyv1 = append(yyv1, nil)
				yyc1 = true
			}
			if yydb1 {
				z.DecSwallow()
			} else {
				if r.TryDecodeAsNil() {
					if yyv1[yyj1] != nil {
						*yyv1[yyj1] = Node{}
					}
				} else {
					if yyv1[yyj1] == nil {
						yyv1[yyj1] = new(Node)
					}
					yyw2 := yyv1[yyj1]
					yyw2.CodecDecodeSelf(d)
				}
			}
		}
		if yyj1 < len(yyv1) {
			yyv1 = yyv1[:yyj1]
			yyc1 = true
		} else if yyj1 == 0 && yyv1 == nil {
			yyv1 = make([]*Node, 0)
			yyc1 = true
		}
	}
	yyh1.End()
	if yyc1 {
		*v = yyv1
	}
}
