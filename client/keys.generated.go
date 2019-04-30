package client

import (
	"errors"
	"runtime"
	"strconv"
	"time"
	codec1978 "github.com/ugorji/go/codec"
)

const (
	codecSelferCcUTF89381		= 1
	codecSelferCcRAW9381		= 255
	codecSelferValueTypeArray9381	= 10
	codecSelferValueTypeMap9381	= 9
	codecSelferValueTypeString9381	= 6
	codecSelferValueTypeInt9381	= 2
	codecSelferValueTypeUint9381	= 3
	codecSelferValueTypeFloat9381	= 4
	codecSelferBitsize9381		= uint8(32 << (^uint(0) >> 63))
)

var (
	errCodecSelferOnlyMapOrArrayEncodeToStruct9381 = errors.New(`only encoded map or array can be decoded into a struct`)
)

type codecSelfer9381 struct{}

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if codec1978.GenVersion != 10 {
		_, file, _, _ := runtime.Caller(0)
		panic("codecgen version mismatch: current: 10, need " + strconv.FormatInt(int64(codec1978.GenVersion), 10) + ". Re-generate file: " + file)
	}
	if false {
		var _ byte = 0
		var v0 time.Duration
		_ = v0
	}
}
func (x *Error) CodecEncodeSelf(e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		if false {
		} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
			z.EncExtension(x, yyxt1)
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
				if false {
				} else {
					r.EncodeInt(int64(x.Code))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"errorCode\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `errorCode`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeInt(int64(x.Code))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Message))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"message\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `message`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Message))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Cause))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"cause\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `cause`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Cause))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeUint(uint64(x.Index))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"index\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `index`)
				}
				r.WriteMapElemValue()
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
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.DecExtension(x, yyxt1)
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap9381 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray9381 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(errCodecSelferOnlyMapOrArrayEncodeToStruct9381)
		}
	}
}
func (x *Error) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
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
		yys3 := z.StringView(r.DecodeStringAsBytes())
		r.ReadMapElemValue()
		switch yys3 {
		case "errorCode":
			if r.TryDecodeAsNil() {
				x.Code = 0
			} else {
				x.Code = (int)(z.C.IntV(r.DecodeInt64(), codecSelferBitsize9381))
			}
		case "message":
			if r.TryDecodeAsNil() {
				x.Message = ""
			} else {
				x.Message = (string)(r.DecodeString())
			}
		case "cause":
			if r.TryDecodeAsNil() {
				x.Cause = ""
			} else {
				x.Cause = (string)(r.DecodeString())
			}
		case "index":
			if r.TryDecodeAsNil() {
				x.Index = 0
			} else {
				x.Index = (uint64)(r.DecodeUint64())
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
	var h codecSelfer9381
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
		x.Code = 0
	} else {
		x.Code = (int)(z.C.IntV(r.DecodeInt64(), codecSelferBitsize9381))
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
		x.Message = ""
	} else {
		x.Message = (string)(r.DecodeString())
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
		x.Cause = ""
	} else {
		x.Cause = (string)(r.DecodeString())
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
		x.Index = 0
	} else {
		x.Index = (uint64)(r.DecodeUint64())
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
func (x PrevExistType) CodecEncodeSelf(e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.EncExtension(x, yyxt1)
	} else {
		r.EncodeStringEnc(codecSelferCcUTF89381, string(x))
	}
}
func (x *PrevExistType) CodecDecodeSelf(d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.DecExtension(x, yyxt1)
	} else {
		*x = (PrevExistType)(r.DecodeString())
	}
}
func (x *WatcherOptions) CodecEncodeSelf(e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		if false {
		} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
			z.EncExtension(x, yyxt1)
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
				if false {
				} else {
					r.EncodeUint(uint64(x.AfterIndex))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"AfterIndex\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `AfterIndex`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeUint(uint64(x.AfterIndex))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeBool(bool(x.Recursive))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"Recursive\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `Recursive`)
				}
				r.WriteMapElemValue()
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
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.DecExtension(x, yyxt1)
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap9381 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray9381 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(errCodecSelferOnlyMapOrArrayEncodeToStruct9381)
		}
	}
}
func (x *WatcherOptions) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
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
		yys3 := z.StringView(r.DecodeStringAsBytes())
		r.ReadMapElemValue()
		switch yys3 {
		case "AfterIndex":
			if r.TryDecodeAsNil() {
				x.AfterIndex = 0
			} else {
				x.AfterIndex = (uint64)(r.DecodeUint64())
			}
		case "Recursive":
			if r.TryDecodeAsNil() {
				x.Recursive = false
			} else {
				x.Recursive = (bool)(r.DecodeBool())
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
	var h codecSelfer9381
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
		x.AfterIndex = 0
	} else {
		x.AfterIndex = (uint64)(r.DecodeUint64())
	}
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
		x.Recursive = false
	} else {
		x.Recursive = (bool)(r.DecodeBool())
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
func (x *CreateInOrderOptions) CodecEncodeSelf(e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		if false {
		} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
			z.EncExtension(x, yyxt1)
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
				if false {
				} else if yyxt4 := z.Extension(z.I2Rtid(x.TTL)); yyxt4 != nil {
					z.EncExtension(x.TTL, yyxt4)
				} else {
					r.EncodeInt(int64(x.TTL))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"TTL\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `TTL`)
				}
				r.WriteMapElemValue()
				if false {
				} else if yyxt5 := z.Extension(z.I2Rtid(x.TTL)); yyxt5 != nil {
					z.EncExtension(x.TTL, yyxt5)
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
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.DecExtension(x, yyxt1)
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap9381 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray9381 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(errCodecSelferOnlyMapOrArrayEncodeToStruct9381)
		}
	}
}
func (x *CreateInOrderOptions) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
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
		yys3 := z.StringView(r.DecodeStringAsBytes())
		r.ReadMapElemValue()
		switch yys3 {
		case "TTL":
			if r.TryDecodeAsNil() {
				x.TTL = 0
			} else {
				if false {
				} else if yyxt5 := z.Extension(z.I2Rtid(x.TTL)); yyxt5 != nil {
					z.DecExtension(x.TTL, yyxt5)
				} else {
					x.TTL = (time.Duration)(r.DecodeInt64())
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
	var h codecSelfer9381
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
		if false {
		} else if yyxt8 := z.Extension(z.I2Rtid(x.TTL)); yyxt8 != nil {
			z.DecExtension(x.TTL, yyxt8)
		} else {
			x.TTL = (time.Duration)(r.DecodeInt64())
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
	var h codecSelfer9381
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		if false {
		} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
			z.EncExtension(x, yyxt1)
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
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.PrevValue))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"PrevValue\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `PrevValue`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.PrevValue))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeUint(uint64(x.PrevIndex))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"PrevIndex\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `PrevIndex`)
				}
				r.WriteMapElemValue()
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
				if z.IsJSONHandle() {
					z.WriteStr("\"PrevExist\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `PrevExist`)
				}
				r.WriteMapElemValue()
				x.PrevExist.CodecEncodeSelf(e)
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else if yyxt13 := z.Extension(z.I2Rtid(x.TTL)); yyxt13 != nil {
					z.EncExtension(x.TTL, yyxt13)
				} else {
					r.EncodeInt(int64(x.TTL))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"TTL\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `TTL`)
				}
				r.WriteMapElemValue()
				if false {
				} else if yyxt14 := z.Extension(z.I2Rtid(x.TTL)); yyxt14 != nil {
					z.EncExtension(x.TTL, yyxt14)
				} else {
					r.EncodeInt(int64(x.TTL))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeBool(bool(x.Refresh))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"Refresh\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `Refresh`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeBool(bool(x.Refresh))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeBool(bool(x.Dir))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"Dir\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `Dir`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeBool(bool(x.Dir))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeBool(bool(x.NoValueOnSuccess))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"NoValueOnSuccess\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `NoValueOnSuccess`)
				}
				r.WriteMapElemValue()
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
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.DecExtension(x, yyxt1)
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap9381 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray9381 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(errCodecSelferOnlyMapOrArrayEncodeToStruct9381)
		}
	}
}
func (x *SetOptions) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
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
		yys3 := z.StringView(r.DecodeStringAsBytes())
		r.ReadMapElemValue()
		switch yys3 {
		case "PrevValue":
			if r.TryDecodeAsNil() {
				x.PrevValue = ""
			} else {
				x.PrevValue = (string)(r.DecodeString())
			}
		case "PrevIndex":
			if r.TryDecodeAsNil() {
				x.PrevIndex = 0
			} else {
				x.PrevIndex = (uint64)(r.DecodeUint64())
			}
		case "PrevExist":
			if r.TryDecodeAsNil() {
				x.PrevExist = ""
			} else {
				x.PrevExist.CodecDecodeSelf(d)
			}
		case "TTL":
			if r.TryDecodeAsNil() {
				x.TTL = 0
			} else {
				if false {
				} else if yyxt8 := z.Extension(z.I2Rtid(x.TTL)); yyxt8 != nil {
					z.DecExtension(x.TTL, yyxt8)
				} else {
					x.TTL = (time.Duration)(r.DecodeInt64())
				}
			}
		case "Refresh":
			if r.TryDecodeAsNil() {
				x.Refresh = false
			} else {
				x.Refresh = (bool)(r.DecodeBool())
			}
		case "Dir":
			if r.TryDecodeAsNil() {
				x.Dir = false
			} else {
				x.Dir = (bool)(r.DecodeBool())
			}
		case "NoValueOnSuccess":
			if r.TryDecodeAsNil() {
				x.NoValueOnSuccess = false
			} else {
				x.NoValueOnSuccess = (bool)(r.DecodeBool())
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
	var h codecSelfer9381
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
		x.PrevValue = (string)(r.DecodeString())
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
		x.PrevIndex = (uint64)(r.DecodeUint64())
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
		x.PrevExist = ""
	} else {
		x.PrevExist.CodecDecodeSelf(d)
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
		if false {
		} else if yyxt17 := z.Extension(z.I2Rtid(x.TTL)); yyxt17 != nil {
			z.DecExtension(x.TTL, yyxt17)
		} else {
			x.TTL = (time.Duration)(r.DecodeInt64())
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
		x.Refresh = false
	} else {
		x.Refresh = (bool)(r.DecodeBool())
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
		x.Dir = (bool)(r.DecodeBool())
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
		x.NoValueOnSuccess = false
	} else {
		x.NoValueOnSuccess = (bool)(r.DecodeBool())
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
func (x *GetOptions) CodecEncodeSelf(e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		if false {
		} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
			z.EncExtension(x, yyxt1)
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
				if false {
				} else {
					r.EncodeBool(bool(x.Recursive))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"Recursive\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `Recursive`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeBool(bool(x.Recursive))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeBool(bool(x.Sort))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"Sort\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `Sort`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeBool(bool(x.Sort))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeBool(bool(x.Quorum))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"Quorum\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `Quorum`)
				}
				r.WriteMapElemValue()
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
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.DecExtension(x, yyxt1)
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap9381 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray9381 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(errCodecSelferOnlyMapOrArrayEncodeToStruct9381)
		}
	}
}
func (x *GetOptions) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
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
		yys3 := z.StringView(r.DecodeStringAsBytes())
		r.ReadMapElemValue()
		switch yys3 {
		case "Recursive":
			if r.TryDecodeAsNil() {
				x.Recursive = false
			} else {
				x.Recursive = (bool)(r.DecodeBool())
			}
		case "Sort":
			if r.TryDecodeAsNil() {
				x.Sort = false
			} else {
				x.Sort = (bool)(r.DecodeBool())
			}
		case "Quorum":
			if r.TryDecodeAsNil() {
				x.Quorum = false
			} else {
				x.Quorum = (bool)(r.DecodeBool())
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
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj7 int
	var yyb7 bool
	var yyhl7 bool = l >= 0
	yyj7++
	if yyhl7 {
		yyb7 = yyj7 > l
	} else {
		yyb7 = r.CheckBreak()
	}
	if yyb7 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Recursive = false
	} else {
		x.Recursive = (bool)(r.DecodeBool())
	}
	yyj7++
	if yyhl7 {
		yyb7 = yyj7 > l
	} else {
		yyb7 = r.CheckBreak()
	}
	if yyb7 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Sort = false
	} else {
		x.Sort = (bool)(r.DecodeBool())
	}
	yyj7++
	if yyhl7 {
		yyb7 = yyj7 > l
	} else {
		yyb7 = r.CheckBreak()
	}
	if yyb7 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Quorum = false
	} else {
		x.Quorum = (bool)(r.DecodeBool())
	}
	for {
		yyj7++
		if yyhl7 {
			yyb7 = yyj7 > l
		} else {
			yyb7 = r.CheckBreak()
		}
		if yyb7 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj7-1, "")
	}
	r.ReadArrayEnd()
}
func (x *DeleteOptions) CodecEncodeSelf(e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		if false {
		} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
			z.EncExtension(x, yyxt1)
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
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.PrevValue))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"PrevValue\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `PrevValue`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.PrevValue))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeUint(uint64(x.PrevIndex))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"PrevIndex\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `PrevIndex`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeUint(uint64(x.PrevIndex))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeBool(bool(x.Recursive))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"Recursive\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `Recursive`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeBool(bool(x.Recursive))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeBool(bool(x.Dir))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"Dir\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `Dir`)
				}
				r.WriteMapElemValue()
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
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.DecExtension(x, yyxt1)
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap9381 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray9381 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(errCodecSelferOnlyMapOrArrayEncodeToStruct9381)
		}
	}
}
func (x *DeleteOptions) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
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
		yys3 := z.StringView(r.DecodeStringAsBytes())
		r.ReadMapElemValue()
		switch yys3 {
		case "PrevValue":
			if r.TryDecodeAsNil() {
				x.PrevValue = ""
			} else {
				x.PrevValue = (string)(r.DecodeString())
			}
		case "PrevIndex":
			if r.TryDecodeAsNil() {
				x.PrevIndex = 0
			} else {
				x.PrevIndex = (uint64)(r.DecodeUint64())
			}
		case "Recursive":
			if r.TryDecodeAsNil() {
				x.Recursive = false
			} else {
				x.Recursive = (bool)(r.DecodeBool())
			}
		case "Dir":
			if r.TryDecodeAsNil() {
				x.Dir = false
			} else {
				x.Dir = (bool)(r.DecodeBool())
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
	var h codecSelfer9381
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
		x.PrevValue = ""
	} else {
		x.PrevValue = (string)(r.DecodeString())
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
		x.PrevIndex = 0
	} else {
		x.PrevIndex = (uint64)(r.DecodeUint64())
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
		x.Recursive = (bool)(r.DecodeBool())
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
		x.Dir = false
	} else {
		x.Dir = (bool)(r.DecodeBool())
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
func (x *Response) CodecEncodeSelf(e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		if false {
		} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
			z.EncExtension(x, yyxt1)
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
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Action))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"action\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `action`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Action))
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
				if z.IsJSONHandle() {
					z.WriteStr("\"node\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `node`)
				}
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
				if z.IsJSONHandle() {
					z.WriteStr("\"prevNode\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `prevNode`)
				}
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
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.DecExtension(x, yyxt1)
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap9381 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray9381 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(errCodecSelferOnlyMapOrArrayEncodeToStruct9381)
		}
	}
}
func (x *Response) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
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
		yys3 := z.StringView(r.DecodeStringAsBytes())
		r.ReadMapElemValue()
		switch yys3 {
		case "action":
			if r.TryDecodeAsNil() {
				x.Action = ""
			} else {
				x.Action = (string)(r.DecodeString())
			}
		case "node":
			if r.TryDecodeAsNil() {
				if true && x.Node != nil {
					x.Node = nil
				}
			} else {
				if x.Node == nil {
					x.Node = new(Node)
				}
				x.Node.CodecDecodeSelf(d)
			}
		case "prevNode":
			if r.TryDecodeAsNil() {
				if true && x.PrevNode != nil {
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
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj7 int
	var yyb7 bool
	var yyhl7 bool = l >= 0
	yyj7++
	if yyhl7 {
		yyb7 = yyj7 > l
	} else {
		yyb7 = r.CheckBreak()
	}
	if yyb7 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Action = ""
	} else {
		x.Action = (string)(r.DecodeString())
	}
	yyj7++
	if yyhl7 {
		yyb7 = yyj7 > l
	} else {
		yyb7 = r.CheckBreak()
	}
	if yyb7 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if true && x.Node != nil {
			x.Node = nil
		}
	} else {
		if x.Node == nil {
			x.Node = new(Node)
		}
		x.Node.CodecDecodeSelf(d)
	}
	yyj7++
	if yyhl7 {
		yyb7 = yyj7 > l
	} else {
		yyb7 = r.CheckBreak()
	}
	if yyb7 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if true && x.PrevNode != nil {
			x.PrevNode = nil
		}
	} else {
		if x.PrevNode == nil {
			x.PrevNode = new(Node)
		}
		x.PrevNode.CodecDecodeSelf(d)
	}
	for {
		yyj7++
		if yyhl7 {
			yyb7 = yyj7 > l
		} else {
			yyb7 = r.CheckBreak()
		}
		if yyb7 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj7-1, "")
	}
	r.ReadArrayEnd()
}
func (x *Node) CodecEncodeSelf(e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		if false {
		} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
			z.EncExtension(x, yyxt1)
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			_, _ = yysep2, yy2arr2
			const yyr2 bool = false
			var yyq2 = [8]bool{true, x.Dir, true, true, true, true, x.Expiration != nil, x.TTL != 0}
			_ = yyq2
			if yyr2 || yy2arr2 {
				r.WriteArrayStart(8)
			} else {
				var yynn2 int
				for _, b := range yyq2 {
					if b {
						yynn2++
					}
				}
				r.WriteMapStart(yynn2)
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Key))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"key\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `key`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Key))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if yyq2[1] {
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
					if z.IsJSONHandle() {
						z.WriteStr("\"dir\"")
					} else {
						r.EncodeStringEnc(codecSelferCcUTF89381, `dir`)
					}
					r.WriteMapElemValue()
					if false {
					} else {
						r.EncodeBool(bool(x.Dir))
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Value))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"value\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `value`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Value))
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
				if z.IsJSONHandle() {
					z.WriteStr("\"nodes\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `nodes`)
				}
				r.WriteMapElemValue()
				if x.Nodes == nil {
					r.EncodeNil()
				} else {
					x.Nodes.CodecEncodeSelf(e)
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeUint(uint64(x.CreatedIndex))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"createdIndex\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `createdIndex`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeUint(uint64(x.CreatedIndex))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeUint(uint64(x.ModifiedIndex))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"modifiedIndex\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `modifiedIndex`)
				}
				r.WriteMapElemValue()
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
							yy22 := *x.Expiration
							if false {
							} else if !z.EncBasicHandle().TimeNotBuiltin {
								r.EncodeTime(yy22)
							} else if yyxt23 := z.Extension(z.I2Rtid(yy22)); yyxt23 != nil {
								z.EncExtension(yy22, yyxt23)
							} else if z.EncBinary() {
								z.EncBinaryMarshal(yy22)
							} else if !z.EncBinary() && z.IsJSONHandle() {
								z.EncJSONMarshal(yy22)
							} else {
								z.EncFallback(yy22)
							}
						}
					} else {
						r.EncodeNil()
					}
				}
			} else {
				if yyq2[6] {
					r.WriteMapElemKey()
					if z.IsJSONHandle() {
						z.WriteStr("\"expiration\"")
					} else {
						r.EncodeStringEnc(codecSelferCcUTF89381, `expiration`)
					}
					r.WriteMapElemValue()
					if yyn21 {
						r.EncodeNil()
					} else {
						if x.Expiration == nil {
							r.EncodeNil()
						} else {
							yy24 := *x.Expiration
							if false {
							} else if !z.EncBasicHandle().TimeNotBuiltin {
								r.EncodeTime(yy24)
							} else if yyxt25 := z.Extension(z.I2Rtid(yy24)); yyxt25 != nil {
								z.EncExtension(yy24, yyxt25)
							} else if z.EncBinary() {
								z.EncBinaryMarshal(yy24)
							} else if !z.EncBinary() && z.IsJSONHandle() {
								z.EncJSONMarshal(yy24)
							} else {
								z.EncFallback(yy24)
							}
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if yyq2[7] {
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
					if z.IsJSONHandle() {
						z.WriteStr("\"ttl\"")
					} else {
						r.EncodeStringEnc(codecSelferCcUTF89381, `ttl`)
					}
					r.WriteMapElemValue()
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
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.DecExtension(x, yyxt1)
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap9381 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray9381 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(errCodecSelferOnlyMapOrArrayEncodeToStruct9381)
		}
	}
}
func (x *Node) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
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
		yys3 := z.StringView(r.DecodeStringAsBytes())
		r.ReadMapElemValue()
		switch yys3 {
		case "key":
			if r.TryDecodeAsNil() {
				x.Key = ""
			} else {
				x.Key = (string)(r.DecodeString())
			}
		case "dir":
			if r.TryDecodeAsNil() {
				x.Dir = false
			} else {
				x.Dir = (bool)(r.DecodeBool())
			}
		case "value":
			if r.TryDecodeAsNil() {
				x.Value = ""
			} else {
				x.Value = (string)(r.DecodeString())
			}
		case "nodes":
			if r.TryDecodeAsNil() {
				x.Nodes = nil
			} else {
				x.Nodes.CodecDecodeSelf(d)
			}
		case "createdIndex":
			if r.TryDecodeAsNil() {
				x.CreatedIndex = 0
			} else {
				x.CreatedIndex = (uint64)(r.DecodeUint64())
			}
		case "modifiedIndex":
			if r.TryDecodeAsNil() {
				x.ModifiedIndex = 0
			} else {
				x.ModifiedIndex = (uint64)(r.DecodeUint64())
			}
		case "expiration":
			if r.TryDecodeAsNil() {
				if true && x.Expiration != nil {
					x.Expiration = nil
				}
			} else {
				if x.Expiration == nil {
					x.Expiration = new(time.Time)
				}
				if false {
				} else if !z.DecBasicHandle().TimeNotBuiltin {
					*x.Expiration = r.DecodeTime()
				} else if yyxt11 := z.Extension(z.I2Rtid(x.Expiration)); yyxt11 != nil {
					z.DecExtension(x.Expiration, yyxt11)
				} else if z.DecBinary() {
					z.DecBinaryUnmarshal(x.Expiration)
				} else if !z.DecBinary() && z.IsJSONHandle() {
					z.DecJSONUnmarshal(x.Expiration)
				} else {
					z.DecFallback(x.Expiration, false)
				}
			}
		case "ttl":
			if r.TryDecodeAsNil() {
				x.TTL = 0
			} else {
				x.TTL = (int64)(r.DecodeInt64())
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
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj13 int
	var yyb13 bool
	var yyhl13 bool = l >= 0
	yyj13++
	if yyhl13 {
		yyb13 = yyj13 > l
	} else {
		yyb13 = r.CheckBreak()
	}
	if yyb13 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Key = ""
	} else {
		x.Key = (string)(r.DecodeString())
	}
	yyj13++
	if yyhl13 {
		yyb13 = yyj13 > l
	} else {
		yyb13 = r.CheckBreak()
	}
	if yyb13 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Dir = false
	} else {
		x.Dir = (bool)(r.DecodeBool())
	}
	yyj13++
	if yyhl13 {
		yyb13 = yyj13 > l
	} else {
		yyb13 = r.CheckBreak()
	}
	if yyb13 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Value = ""
	} else {
		x.Value = (string)(r.DecodeString())
	}
	yyj13++
	if yyhl13 {
		yyb13 = yyj13 > l
	} else {
		yyb13 = r.CheckBreak()
	}
	if yyb13 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Nodes = nil
	} else {
		x.Nodes.CodecDecodeSelf(d)
	}
	yyj13++
	if yyhl13 {
		yyb13 = yyj13 > l
	} else {
		yyb13 = r.CheckBreak()
	}
	if yyb13 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.CreatedIndex = 0
	} else {
		x.CreatedIndex = (uint64)(r.DecodeUint64())
	}
	yyj13++
	if yyhl13 {
		yyb13 = yyj13 > l
	} else {
		yyb13 = r.CheckBreak()
	}
	if yyb13 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.ModifiedIndex = 0
	} else {
		x.ModifiedIndex = (uint64)(r.DecodeUint64())
	}
	yyj13++
	if yyhl13 {
		yyb13 = yyj13 > l
	} else {
		yyb13 = r.CheckBreak()
	}
	if yyb13 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		if true && x.Expiration != nil {
			x.Expiration = nil
		}
	} else {
		if x.Expiration == nil {
			x.Expiration = new(time.Time)
		}
		if false {
		} else if !z.DecBasicHandle().TimeNotBuiltin {
			*x.Expiration = r.DecodeTime()
		} else if yyxt21 := z.Extension(z.I2Rtid(x.Expiration)); yyxt21 != nil {
			z.DecExtension(x.Expiration, yyxt21)
		} else if z.DecBinary() {
			z.DecBinaryUnmarshal(x.Expiration)
		} else if !z.DecBinary() && z.IsJSONHandle() {
			z.DecJSONUnmarshal(x.Expiration)
		} else {
			z.DecFallback(x.Expiration, false)
		}
	}
	yyj13++
	if yyhl13 {
		yyb13 = yyj13 > l
	} else {
		yyb13 = r.CheckBreak()
	}
	if yyb13 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.TTL = 0
	} else {
		x.TTL = (int64)(r.DecodeInt64())
	}
	for {
		yyj13++
		if yyhl13 {
			yyb13 = yyj13 > l
		} else {
			yyb13 = r.CheckBreak()
		}
		if yyb13 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj13-1, "")
	}
	r.ReadArrayEnd()
}
func (x Nodes) CodecEncodeSelf(e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		if false {
		} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
			z.EncExtension(x, yyxt1)
		} else {
			h.encNodes((Nodes)(x), e)
		}
	}
}
func (x *Nodes) CodecDecodeSelf(d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.DecExtension(x, yyxt1)
	} else {
		h.decNodes((*Nodes)(x), d)
	}
}
func (x *httpKeysAPI) CodecEncodeSelf(e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		if false {
		} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
			z.EncExtension(x, yyxt1)
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
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.DecExtension(x, yyxt1)
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap9381 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray9381 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(errCodecSelferOnlyMapOrArrayEncodeToStruct9381)
		}
	}
}
func (x *httpKeysAPI) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
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
		yys3 := z.StringView(r.DecodeStringAsBytes())
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
	var h codecSelfer9381
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
	var h codecSelfer9381
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		if false {
		} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
			z.EncExtension(x, yyxt1)
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
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.DecExtension(x, yyxt1)
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap9381 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray9381 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(errCodecSelferOnlyMapOrArrayEncodeToStruct9381)
		}
	}
}
func (x *httpWatcher) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
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
		yys3 := z.StringView(r.DecodeStringAsBytes())
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
	var h codecSelfer9381
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
	var h codecSelfer9381
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		if false {
		} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
			z.EncExtension(x, yyxt1)
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
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Prefix))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"Prefix\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `Prefix`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Prefix))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Key))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"Key\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `Key`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Key))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeBool(bool(x.Recursive))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"Recursive\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `Recursive`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeBool(bool(x.Recursive))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeBool(bool(x.Sorted))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"Sorted\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `Sorted`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeBool(bool(x.Sorted))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeBool(bool(x.Quorum))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"Quorum\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `Quorum`)
				}
				r.WriteMapElemValue()
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
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.DecExtension(x, yyxt1)
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap9381 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray9381 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(errCodecSelferOnlyMapOrArrayEncodeToStruct9381)
		}
	}
}
func (x *getAction) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
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
		yys3 := z.StringView(r.DecodeStringAsBytes())
		r.ReadMapElemValue()
		switch yys3 {
		case "Prefix":
			if r.TryDecodeAsNil() {
				x.Prefix = ""
			} else {
				x.Prefix = (string)(r.DecodeString())
			}
		case "Key":
			if r.TryDecodeAsNil() {
				x.Key = ""
			} else {
				x.Key = (string)(r.DecodeString())
			}
		case "Recursive":
			if r.TryDecodeAsNil() {
				x.Recursive = false
			} else {
				x.Recursive = (bool)(r.DecodeBool())
			}
		case "Sorted":
			if r.TryDecodeAsNil() {
				x.Sorted = false
			} else {
				x.Sorted = (bool)(r.DecodeBool())
			}
		case "Quorum":
			if r.TryDecodeAsNil() {
				x.Quorum = false
			} else {
				x.Quorum = (bool)(r.DecodeBool())
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
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj9 int
	var yyb9 bool
	var yyhl9 bool = l >= 0
	yyj9++
	if yyhl9 {
		yyb9 = yyj9 > l
	} else {
		yyb9 = r.CheckBreak()
	}
	if yyb9 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Prefix = ""
	} else {
		x.Prefix = (string)(r.DecodeString())
	}
	yyj9++
	if yyhl9 {
		yyb9 = yyj9 > l
	} else {
		yyb9 = r.CheckBreak()
	}
	if yyb9 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Key = ""
	} else {
		x.Key = (string)(r.DecodeString())
	}
	yyj9++
	if yyhl9 {
		yyb9 = yyj9 > l
	} else {
		yyb9 = r.CheckBreak()
	}
	if yyb9 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Recursive = false
	} else {
		x.Recursive = (bool)(r.DecodeBool())
	}
	yyj9++
	if yyhl9 {
		yyb9 = yyj9 > l
	} else {
		yyb9 = r.CheckBreak()
	}
	if yyb9 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Sorted = false
	} else {
		x.Sorted = (bool)(r.DecodeBool())
	}
	yyj9++
	if yyhl9 {
		yyb9 = yyj9 > l
	} else {
		yyb9 = r.CheckBreak()
	}
	if yyb9 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Quorum = false
	} else {
		x.Quorum = (bool)(r.DecodeBool())
	}
	for {
		yyj9++
		if yyhl9 {
			yyb9 = yyj9 > l
		} else {
			yyb9 = r.CheckBreak()
		}
		if yyb9 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj9-1, "")
	}
	r.ReadArrayEnd()
}
func (x *waitAction) CodecEncodeSelf(e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		if false {
		} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
			z.EncExtension(x, yyxt1)
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
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Prefix))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"Prefix\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `Prefix`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Prefix))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Key))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"Key\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `Key`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Key))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeUint(uint64(x.WaitIndex))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"WaitIndex\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `WaitIndex`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeUint(uint64(x.WaitIndex))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeBool(bool(x.Recursive))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"Recursive\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `Recursive`)
				}
				r.WriteMapElemValue()
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
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.DecExtension(x, yyxt1)
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap9381 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray9381 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(errCodecSelferOnlyMapOrArrayEncodeToStruct9381)
		}
	}
}
func (x *waitAction) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
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
		yys3 := z.StringView(r.DecodeStringAsBytes())
		r.ReadMapElemValue()
		switch yys3 {
		case "Prefix":
			if r.TryDecodeAsNil() {
				x.Prefix = ""
			} else {
				x.Prefix = (string)(r.DecodeString())
			}
		case "Key":
			if r.TryDecodeAsNil() {
				x.Key = ""
			} else {
				x.Key = (string)(r.DecodeString())
			}
		case "WaitIndex":
			if r.TryDecodeAsNil() {
				x.WaitIndex = 0
			} else {
				x.WaitIndex = (uint64)(r.DecodeUint64())
			}
		case "Recursive":
			if r.TryDecodeAsNil() {
				x.Recursive = false
			} else {
				x.Recursive = (bool)(r.DecodeBool())
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
	var h codecSelfer9381
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
		x.Prefix = ""
	} else {
		x.Prefix = (string)(r.DecodeString())
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
		x.Key = ""
	} else {
		x.Key = (string)(r.DecodeString())
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
		x.WaitIndex = 0
	} else {
		x.WaitIndex = (uint64)(r.DecodeUint64())
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
		x.Recursive = (bool)(r.DecodeBool())
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
func (x *setAction) CodecEncodeSelf(e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		if false {
		} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
			z.EncExtension(x, yyxt1)
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
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Prefix))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"Prefix\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `Prefix`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Prefix))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Key))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"Key\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `Key`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Key))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Value))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"Value\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `Value`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Value))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.PrevValue))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"PrevValue\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `PrevValue`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.PrevValue))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeUint(uint64(x.PrevIndex))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"PrevIndex\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `PrevIndex`)
				}
				r.WriteMapElemValue()
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
				if z.IsJSONHandle() {
					z.WriteStr("\"PrevExist\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `PrevExist`)
				}
				r.WriteMapElemValue()
				x.PrevExist.CodecEncodeSelf(e)
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else if yyxt22 := z.Extension(z.I2Rtid(x.TTL)); yyxt22 != nil {
					z.EncExtension(x.TTL, yyxt22)
				} else {
					r.EncodeInt(int64(x.TTL))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"TTL\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `TTL`)
				}
				r.WriteMapElemValue()
				if false {
				} else if yyxt23 := z.Extension(z.I2Rtid(x.TTL)); yyxt23 != nil {
					z.EncExtension(x.TTL, yyxt23)
				} else {
					r.EncodeInt(int64(x.TTL))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeBool(bool(x.Refresh))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"Refresh\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `Refresh`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeBool(bool(x.Refresh))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeBool(bool(x.Dir))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"Dir\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `Dir`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeBool(bool(x.Dir))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeBool(bool(x.NoValueOnSuccess))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"NoValueOnSuccess\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `NoValueOnSuccess`)
				}
				r.WriteMapElemValue()
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
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.DecExtension(x, yyxt1)
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap9381 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray9381 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(errCodecSelferOnlyMapOrArrayEncodeToStruct9381)
		}
	}
}
func (x *setAction) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
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
		yys3 := z.StringView(r.DecodeStringAsBytes())
		r.ReadMapElemValue()
		switch yys3 {
		case "Prefix":
			if r.TryDecodeAsNil() {
				x.Prefix = ""
			} else {
				x.Prefix = (string)(r.DecodeString())
			}
		case "Key":
			if r.TryDecodeAsNil() {
				x.Key = ""
			} else {
				x.Key = (string)(r.DecodeString())
			}
		case "Value":
			if r.TryDecodeAsNil() {
				x.Value = ""
			} else {
				x.Value = (string)(r.DecodeString())
			}
		case "PrevValue":
			if r.TryDecodeAsNil() {
				x.PrevValue = ""
			} else {
				x.PrevValue = (string)(r.DecodeString())
			}
		case "PrevIndex":
			if r.TryDecodeAsNil() {
				x.PrevIndex = 0
			} else {
				x.PrevIndex = (uint64)(r.DecodeUint64())
			}
		case "PrevExist":
			if r.TryDecodeAsNil() {
				x.PrevExist = ""
			} else {
				x.PrevExist.CodecDecodeSelf(d)
			}
		case "TTL":
			if r.TryDecodeAsNil() {
				x.TTL = 0
			} else {
				if false {
				} else if yyxt11 := z.Extension(z.I2Rtid(x.TTL)); yyxt11 != nil {
					z.DecExtension(x.TTL, yyxt11)
				} else {
					x.TTL = (time.Duration)(r.DecodeInt64())
				}
			}
		case "Refresh":
			if r.TryDecodeAsNil() {
				x.Refresh = false
			} else {
				x.Refresh = (bool)(r.DecodeBool())
			}
		case "Dir":
			if r.TryDecodeAsNil() {
				x.Dir = false
			} else {
				x.Dir = (bool)(r.DecodeBool())
			}
		case "NoValueOnSuccess":
			if r.TryDecodeAsNil() {
				x.NoValueOnSuccess = false
			} else {
				x.NoValueOnSuccess = (bool)(r.DecodeBool())
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
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj15 int
	var yyb15 bool
	var yyhl15 bool = l >= 0
	yyj15++
	if yyhl15 {
		yyb15 = yyj15 > l
	} else {
		yyb15 = r.CheckBreak()
	}
	if yyb15 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Prefix = ""
	} else {
		x.Prefix = (string)(r.DecodeString())
	}
	yyj15++
	if yyhl15 {
		yyb15 = yyj15 > l
	} else {
		yyb15 = r.CheckBreak()
	}
	if yyb15 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Key = ""
	} else {
		x.Key = (string)(r.DecodeString())
	}
	yyj15++
	if yyhl15 {
		yyb15 = yyj15 > l
	} else {
		yyb15 = r.CheckBreak()
	}
	if yyb15 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Value = ""
	} else {
		x.Value = (string)(r.DecodeString())
	}
	yyj15++
	if yyhl15 {
		yyb15 = yyj15 > l
	} else {
		yyb15 = r.CheckBreak()
	}
	if yyb15 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.PrevValue = ""
	} else {
		x.PrevValue = (string)(r.DecodeString())
	}
	yyj15++
	if yyhl15 {
		yyb15 = yyj15 > l
	} else {
		yyb15 = r.CheckBreak()
	}
	if yyb15 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.PrevIndex = 0
	} else {
		x.PrevIndex = (uint64)(r.DecodeUint64())
	}
	yyj15++
	if yyhl15 {
		yyb15 = yyj15 > l
	} else {
		yyb15 = r.CheckBreak()
	}
	if yyb15 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.PrevExist = ""
	} else {
		x.PrevExist.CodecDecodeSelf(d)
	}
	yyj15++
	if yyhl15 {
		yyb15 = yyj15 > l
	} else {
		yyb15 = r.CheckBreak()
	}
	if yyb15 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.TTL = 0
	} else {
		if false {
		} else if yyxt23 := z.Extension(z.I2Rtid(x.TTL)); yyxt23 != nil {
			z.DecExtension(x.TTL, yyxt23)
		} else {
			x.TTL = (time.Duration)(r.DecodeInt64())
		}
	}
	yyj15++
	if yyhl15 {
		yyb15 = yyj15 > l
	} else {
		yyb15 = r.CheckBreak()
	}
	if yyb15 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Refresh = false
	} else {
		x.Refresh = (bool)(r.DecodeBool())
	}
	yyj15++
	if yyhl15 {
		yyb15 = yyj15 > l
	} else {
		yyb15 = r.CheckBreak()
	}
	if yyb15 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Dir = false
	} else {
		x.Dir = (bool)(r.DecodeBool())
	}
	yyj15++
	if yyhl15 {
		yyb15 = yyj15 > l
	} else {
		yyb15 = r.CheckBreak()
	}
	if yyb15 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.NoValueOnSuccess = false
	} else {
		x.NoValueOnSuccess = (bool)(r.DecodeBool())
	}
	for {
		yyj15++
		if yyhl15 {
			yyb15 = yyj15 > l
		} else {
			yyb15 = r.CheckBreak()
		}
		if yyb15 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj15-1, "")
	}
	r.ReadArrayEnd()
}
func (x *deleteAction) CodecEncodeSelf(e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		if false {
		} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
			z.EncExtension(x, yyxt1)
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
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Prefix))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"Prefix\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `Prefix`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Prefix))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Key))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"Key\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `Key`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Key))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.PrevValue))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"PrevValue\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `PrevValue`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.PrevValue))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeUint(uint64(x.PrevIndex))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"PrevIndex\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `PrevIndex`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeUint(uint64(x.PrevIndex))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeBool(bool(x.Dir))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"Dir\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `Dir`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeBool(bool(x.Dir))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeBool(bool(x.Recursive))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"Recursive\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `Recursive`)
				}
				r.WriteMapElemValue()
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
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.DecExtension(x, yyxt1)
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap9381 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray9381 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(errCodecSelferOnlyMapOrArrayEncodeToStruct9381)
		}
	}
}
func (x *deleteAction) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
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
		yys3 := z.StringView(r.DecodeStringAsBytes())
		r.ReadMapElemValue()
		switch yys3 {
		case "Prefix":
			if r.TryDecodeAsNil() {
				x.Prefix = ""
			} else {
				x.Prefix = (string)(r.DecodeString())
			}
		case "Key":
			if r.TryDecodeAsNil() {
				x.Key = ""
			} else {
				x.Key = (string)(r.DecodeString())
			}
		case "PrevValue":
			if r.TryDecodeAsNil() {
				x.PrevValue = ""
			} else {
				x.PrevValue = (string)(r.DecodeString())
			}
		case "PrevIndex":
			if r.TryDecodeAsNil() {
				x.PrevIndex = 0
			} else {
				x.PrevIndex = (uint64)(r.DecodeUint64())
			}
		case "Dir":
			if r.TryDecodeAsNil() {
				x.Dir = false
			} else {
				x.Dir = (bool)(r.DecodeBool())
			}
		case "Recursive":
			if r.TryDecodeAsNil() {
				x.Recursive = false
			} else {
				x.Recursive = (bool)(r.DecodeBool())
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
	var h codecSelfer9381
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
		x.Prefix = ""
	} else {
		x.Prefix = (string)(r.DecodeString())
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
		x.Key = ""
	} else {
		x.Key = (string)(r.DecodeString())
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
		x.PrevValue = ""
	} else {
		x.PrevValue = (string)(r.DecodeString())
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
		x.PrevIndex = 0
	} else {
		x.PrevIndex = (uint64)(r.DecodeUint64())
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
		x.Dir = false
	} else {
		x.Dir = (bool)(r.DecodeBool())
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
		x.Recursive = false
	} else {
		x.Recursive = (bool)(r.DecodeBool())
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
func (x *createInOrderAction) CodecEncodeSelf(e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		if false {
		} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
			z.EncExtension(x, yyxt1)
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
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Prefix))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"Prefix\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `Prefix`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Prefix))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Dir))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"Dir\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `Dir`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Dir))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Value))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"Value\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `Value`)
				}
				r.WriteMapElemValue()
				if false {
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, string(x.Value))
				}
			}
			if yyr2 || yy2arr2 {
				r.WriteArrayElem()
				if false {
				} else if yyxt13 := z.Extension(z.I2Rtid(x.TTL)); yyxt13 != nil {
					z.EncExtension(x.TTL, yyxt13)
				} else {
					r.EncodeInt(int64(x.TTL))
				}
			} else {
				r.WriteMapElemKey()
				if z.IsJSONHandle() {
					z.WriteStr("\"TTL\"")
				} else {
					r.EncodeStringEnc(codecSelferCcUTF89381, `TTL`)
				}
				r.WriteMapElemValue()
				if false {
				} else if yyxt14 := z.Extension(z.I2Rtid(x.TTL)); yyxt14 != nil {
					z.EncExtension(x.TTL, yyxt14)
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
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	if false {
	} else if yyxt1 := z.Extension(z.I2Rtid(x)); yyxt1 != nil {
		z.DecExtension(x, yyxt1)
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap9381 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				r.ReadMapEnd()
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray9381 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				r.ReadArrayEnd()
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(errCodecSelferOnlyMapOrArrayEncodeToStruct9381)
		}
	}
}
func (x *createInOrderAction) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
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
		yys3 := z.StringView(r.DecodeStringAsBytes())
		r.ReadMapElemValue()
		switch yys3 {
		case "Prefix":
			if r.TryDecodeAsNil() {
				x.Prefix = ""
			} else {
				x.Prefix = (string)(r.DecodeString())
			}
		case "Dir":
			if r.TryDecodeAsNil() {
				x.Dir = ""
			} else {
				x.Dir = (string)(r.DecodeString())
			}
		case "Value":
			if r.TryDecodeAsNil() {
				x.Value = ""
			} else {
				x.Value = (string)(r.DecodeString())
			}
		case "TTL":
			if r.TryDecodeAsNil() {
				x.TTL = 0
			} else {
				if false {
				} else if yyxt8 := z.Extension(z.I2Rtid(x.TTL)); yyxt8 != nil {
					z.DecExtension(x.TTL, yyxt8)
				} else {
					x.TTL = (time.Duration)(r.DecodeInt64())
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
	var h codecSelfer9381
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj9 int
	var yyb9 bool
	var yyhl9 bool = l >= 0
	yyj9++
	if yyhl9 {
		yyb9 = yyj9 > l
	} else {
		yyb9 = r.CheckBreak()
	}
	if yyb9 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Prefix = ""
	} else {
		x.Prefix = (string)(r.DecodeString())
	}
	yyj9++
	if yyhl9 {
		yyb9 = yyj9 > l
	} else {
		yyb9 = r.CheckBreak()
	}
	if yyb9 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Dir = ""
	} else {
		x.Dir = (string)(r.DecodeString())
	}
	yyj9++
	if yyhl9 {
		yyb9 = yyj9 > l
	} else {
		yyb9 = r.CheckBreak()
	}
	if yyb9 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.Value = ""
	} else {
		x.Value = (string)(r.DecodeString())
	}
	yyj9++
	if yyhl9 {
		yyb9 = yyj9 > l
	} else {
		yyb9 = r.CheckBreak()
	}
	if yyb9 {
		r.ReadArrayEnd()
		return
	}
	r.ReadArrayElem()
	if r.TryDecodeAsNil() {
		x.TTL = 0
	} else {
		if false {
		} else if yyxt14 := z.Extension(z.I2Rtid(x.TTL)); yyxt14 != nil {
			z.DecExtension(x.TTL, yyxt14)
		} else {
			x.TTL = (time.Duration)(r.DecodeInt64())
		}
	}
	for {
		yyj9++
		if yyhl9 {
			yyb9 = yyj9 > l
		} else {
			yyb9 = r.CheckBreak()
		}
		if yyb9 {
			break
		}
		r.ReadArrayElem()
		z.DecStructFieldNotFound(yyj9-1, "")
	}
	r.ReadArrayEnd()
}
func (x codecSelfer9381) encNodes(v Nodes, e *codec1978.Encoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
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
func (x codecSelfer9381) decNodes(v *Nodes, d *codec1978.Decoder) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	var h codecSelfer9381
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
		for yyj1 = 0; (yyhl1 && yyj1 < yyl1) || !(yyhl1 || r.CheckBreak()); yyj1++ {
			if yyj1 == 0 && yyv1 == nil {
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
					yyv1[yyj1] = nil
				} else {
					if yyv1[yyj1] == nil {
						yyv1[yyj1] = new(Node)
					}
					yyv1[yyj1].CodecDecodeSelf(d)
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
