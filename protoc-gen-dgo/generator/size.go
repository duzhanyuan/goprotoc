// Copyright (c) 2013, Vastech SA (PTY) LTD. All rights reserved.
// http://code.google.com/p/gogoprotobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

/*
This file generates a Size method for each message.
This is useful with the MarshalTo method generated by the marshalto

The size code also generates a test given it is enabled using one of the following extensions:

  - testgen
  - testgen_all

And a benchmark given it is enabled using one of the following extensions:

  - benchgen
  - benchgen_all

Let us look at:

  code.google.com/p/gogoprotobuf/test/example/example.proto

Btw all the output can be seen at:

  code.google.com/p/gogoprotobuf/test/example/*

The following message:

  message B {
    option (gogoproto.description) = true;
    optional string A = 1 [(gogoproto.embed) = true];
    repeated int64 G = 2 [(gogoproto.customtype) = "github.com/dropbox/goprotoc/test.Id"];
  }

will generate the following code:

  func (m *B) Size() (n int) {
    var l int
    _ = l
    if m.xxx_IsASet {
        l = len(m.a)
        n += 1 + l + sovCustom(uint64(l))
    }
    if m.xxx_LenG > 0 {
        for i := 0; i < m.xxx_LenG; i++ {
            e := m.g[i]
            n += 1 + sovCustom(uint64(e))
        }
    }
    if m.XXX_unrecognized != nil {
        n += len(m.XXX_unrecognized)
    }
    m.xxx_sizeCached = n
    return n
  }

and the following test code:

  func TestBSize(t *testing5.T) {
    popr := math_rand5.New(math_rand5.NewSource(time5.Now().UnixNano()))
    p := NewPopulatedB(popr, true)
    data, err := dropbox_gogoprotobuf_proto2.Marshal(p)
    if err != nil {
      panic(err)
    }
    size := g.Size()
    if len(data) != size {
      t.Fatalf("size %v != marshalled size %v", size, len(data))
    }
  }

  func BenchmarkBSize(b *testing5.B) {
    popr := math_rand5.New(math_rand5.NewSource(616))
    total := 0
    pops := make([]*B, 1000)
    for i := 0; i < 1000; i++ {
      pops[i] = NewPopulatedB(popr, false)
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
      total += pops[i%1000].Size()
    }
    b.SetBytes(int64(total / b.N))
  }

The sovExample function is a size of varint function for the example.pb.go file.

*/
package generator

import (
    "fmt"
    "github.com/dropbox/goprotoc/gogoproto"
    "github.com/dropbox/goprotoc/proto"
    descriptor "github.com/dropbox/goprotoc/protoc-gen-dgo/descriptor"
    "strconv"
)

func wireToType(wire string) int {
    switch wire {
    case "fixed64":
        return proto.WireFixed64
    case "fixed32":
        return proto.WireFixed32
    case "varint":
        return proto.WireVarint
    case "bytes":
        return proto.WireBytes
    case "group":
        return proto.WireBytes
    case "zigzag32":
        return proto.WireVarint
    case "zigzag64":
        return proto.WireVarint
    }
    panic("unreachable")
}

func keySize(fieldNumber int32, wireType int) int {
    x := uint32(fieldNumber)<<3 | uint32(wireType)
    size := 0
    for size = 0; x > 127; size++ {
        x >>= 7
    }
    size++
    return size
}

func (g *Generator) sizeVarint() {
    g.P(`
	func sov`, g.localName, `(x uint64) (n int) {
		for {
			n++
			x >>= 7
			if x == 0 {
				break
			}
		}
		return n
	}`)
}

func (g *Generator) sizeZigZag() {
    g.P(`func soz`, g.localName, `(x uint64) (n int) {
 		return sov`, g.localName, `(uint64((x << 1) ^ uint64((int64(x) >> 63)))) 	
	}`)
}

func (g *Generator) generateSize(file *FileDescriptor) {
    for _, message := range file.Messages() {
        ccTypeName := CamelCaseSlice(message.TypeName())
        g.P(`func (m *`, ccTypeName, `) Size() (n int) {`)
        g.In()
        g.P(`var l int`)
        g.P(`_ = l`)
        for _, field := range message.Field {
            fieldname := g.GetFieldName(message, field)
            repeated := field.IsRepeated()
            sizerName := ""
            if repeated {
                sizerName = SizerName(fieldname)
                g.P(`if m.`, sizerName, ` > 0 {`)
                g.In()
            } else {
                g.P(`if m.`, SetterName(fieldname), ` {`)
                g.In()
            }
            packed := field.IsPacked()
            _, wire := g.GoType(message, field)
            wireType := wireToType(wire)
            fieldNumber := field.GetNumber()
            if packed {
                wireType = proto.WireBytes
            }
            key := keySize(fieldNumber, wireType)
            switch *field.Type {
            case descriptor.FieldDescriptorProto_TYPE_DOUBLE,
                descriptor.FieldDescriptorProto_TYPE_FIXED64,
                descriptor.FieldDescriptorProto_TYPE_SFIXED64:
                if packed {
                    g.P(`n+=`, strconv.Itoa(key), `+sov`, g.localName, `(uint64(m.`, sizerName,
                        `*8))`, `+m.`, sizerName, `*8`)
                } else if repeated {
                    g.P(`n+=`, strconv.Itoa(key+8), `*m.`, sizerName)
                } else {
                    g.P(`n+=`, strconv.Itoa(key+8))
                }
            case descriptor.FieldDescriptorProto_TYPE_FLOAT,
                descriptor.FieldDescriptorProto_TYPE_FIXED32,
                descriptor.FieldDescriptorProto_TYPE_SFIXED32:
                if packed {
                    g.P(`n+=`, strconv.Itoa(key), `+sov`, g.localName, `(uint64(m.`,
                        sizerName, `*4))`, `+m.`, sizerName, `*4`)
                } else if repeated {
                    g.P(`n+=`, strconv.Itoa(key+4), `*m.`, sizerName, ``)
                } else {
                    g.P(`n+=`, strconv.Itoa(key+4))
                }
            case descriptor.FieldDescriptorProto_TYPE_INT64,
                descriptor.FieldDescriptorProto_TYPE_UINT64,
                descriptor.FieldDescriptorProto_TYPE_UINT32,
                descriptor.FieldDescriptorProto_TYPE_ENUM:
                if packed {
                    g.P(`l = 0`)
                    g.P(`for i := 0; i < m.`, sizerName, `; i++ {`)
                    g.In()
                    g.P(`e := m.`, fieldname, `[i]`)
                    g.P(`l+=sov`, g.localName, `(uint64(e))`)
                    g.Out()
                    g.P(`}`)
                    g.P(`n+=`, strconv.Itoa(key), `+sov`, g.localName, `(uint64(l))+l`)
                } else if repeated {
                    g.P(`for i := 0; i < m.`, sizerName, `; i++ {`)
                    g.In()
                    g.P(`e := m.`, fieldname, `[i]`)
                    g.P(`n+=`, strconv.Itoa(key), `+sov`, g.localName, `(uint64(e))`)
                    g.Out()
                    g.P(`}`)
                } else {
                    g.P(`n+=`, strconv.Itoa(key), `+sov`, g.localName, `(uint64(m.`, fieldname, `))`)
                }
            case descriptor.FieldDescriptorProto_TYPE_INT32:
                if packed {
                    g.P(`l = 0`)
                    g.P(`for i := 0; i < m.`, sizerName, `; i++ {`)
                    g.In()
                    g.P(`e := m.`, fieldname, `[i]`)
                    g.P(`l+=sov`, g.localName, `(uint64(uint32(e)))`)
                    g.Out()
                    g.P(`}`)
                    g.P(`n+=`, strconv.Itoa(key), `+sov`, g.localName, `(uint64(l))+l`)
                } else if repeated {
                    g.P(`for i := 0; i < m.`, sizerName, `; i++ {`)
                    g.In()
                    g.P(`e := m.`, fieldname, `[i]`)
                    g.P(`n+=`, strconv.Itoa(key), `+sov`, g.localName, `(uint64(uint32(e)))`)
                    g.Out()
                    g.P(`}`)
                } else {
                    g.P(`n+=`, strconv.Itoa(key), `+sov`, g.localName, `(uint64(uint32(m.`, fieldname, `)))`)
                }
            case descriptor.FieldDescriptorProto_TYPE_BOOL:
                if packed {
                    g.P(`n+=`, strconv.Itoa(key), `+sov`, g.localName, `(uint64(m.`,
                        sizerName, `))`, `+m.`, sizerName, `*1`)
                } else if repeated {
                    g.P(`n+=`, strconv.Itoa(key+1), `*m.`, sizerName, ``)
                } else {
                    g.P(`n+=`, strconv.Itoa(key+1))
                }
            case descriptor.FieldDescriptorProto_TYPE_STRING:
                if repeated {
                    g.P(`for i := 0; i < m.`, sizerName, `; i++ {`)
                    g.In()
                    g.P(`s := m.`, fieldname, `[i]`)
                    g.P(`l = len(s)`)
                    g.P(`n+=`, strconv.Itoa(key), `+l+sov`, g.localName, `(uint64(l))`)
                    g.Out()
                    g.P(`}`)
                } else {
                    g.P(`l=len(m.`, fieldname, `)`)
                    g.P(`n+=`, strconv.Itoa(key), `+l+sov`, g.localName, `(uint64(l))`)
                }
            case descriptor.FieldDescriptorProto_TYPE_GROUP:
                panic(fmt.Errorf("size does not support group %v", fieldname))
            case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
                if repeated {
                    g.P(`for i := 0; i < m.`, sizerName, `; i++ {`)
                    g.In()
                    g.P(`e := m.`, fieldname, `[i]`)
                    g.P(`l=e.Size()`)
                    g.P(`n+=`, strconv.Itoa(key), `+l+sov`, g.localName, `(uint64(l))`)
                    g.Out()
                    g.P(`}`)
                } else {
                    g.P(`l=m.`, fieldname, `.Size()`)
                    g.P(`n+=`, strconv.Itoa(key), `+l+sov`, g.localName, `(uint64(l))`)
                }
            case descriptor.FieldDescriptorProto_TYPE_BYTES:
                if !gogoproto.IsCustomType(field) {
                    if repeated {
                        g.P(`for i := 0; i < m.`, sizerName, `; i++ {`)
                        g.In()
                        g.P(`b := m.`, fieldname, `[i]`)
                        g.P(`l = len(b)`)
                        g.P(`n+=`, strconv.Itoa(key), `+l+sov`, g.localName, `(uint64(l))`)
                        g.Out()
                        g.P(`}`)
                    } else {
                        g.P(`l=len(m.`, fieldname, `)`)
                        g.P(`n+=`, strconv.Itoa(key), `+l+sov`, g.localName, `(uint64(l))`)
                    }
                } else {
                    if repeated {
                        g.P(`for i := 0; i < m.`, sizerName, `; i++ {`)
                        g.In()
                        g.P(`e := m.`, fieldname, `[i]`)
                        g.P(`l=e.Size()`)
                        g.P(`n+=`, strconv.Itoa(key), `+l+sov`, g.localName, `(uint64(l))`)
                        g.Out()
                        g.P(`}`)
                    } else {
                        g.P(`l=m.`, fieldname, `.Size()`)
                        g.P(`n+=`, strconv.Itoa(key), `+l+sov`, g.localName, `(uint64(l))`)
                    }
                }
            case descriptor.FieldDescriptorProto_TYPE_SINT32,
                descriptor.FieldDescriptorProto_TYPE_SINT64:
                if packed {
                    g.P(`l = 0`)
                    g.P(`for i := 0; i < m.`, sizerName, `; i++ {`)
                    g.In()
                    g.P(`e := m.`, fieldname, `[i]`)
                    g.P(`l+=soz`, g.localName, `(uint64(e))`)
                    g.Out()
                    g.P(`}`)
                    g.P(`n+=`, strconv.Itoa(key), `+sov`, g.localName, `(uint64(l))+l`)
                } else if repeated {
                    g.P(`for i := 0; i < m.`, sizerName, `; i++ {`)
                    g.In()
                    g.P(`e := m.`, fieldname, `[i]`)
                    g.P(`n+=`, strconv.Itoa(key), `+soz`, g.localName, `(uint64(e))`)
                    g.Out()
                    g.P(`}`)
                } else {
                    g.P(`n+=`, strconv.Itoa(key), `+soz`, g.localName, `(uint64(m.`, fieldname, `))`)
                }
            default:
                panic("not implemented")
            }
            g.Out()
            g.P(`}`)
        }
        if message.DescriptorProto.HasExtension() {
            g.P(`if m.XXX_extensions != nil {`)
            g.In()
            if gogoproto.HasExtensionsMap(file.FileDescriptorProto, message.DescriptorProto) {
                g.P(`n += `, g.Pkg["proto"], `.SizeOfExtensionMap(m.XXX_extensions)`)
            } else {
                g.P(`n+=len(m.XXX_extensions)`)
            }
            g.Out()
            g.P(`}`)
        }
        g.P(`if m.XXX_unrecognized != nil {`)
        g.In()
        g.P(`n+=len(m.XXX_unrecognized)`)
        g.Out()
        g.P(`}`)
        g.P(`m.xxx_sizeCached = n`)
        g.P(`return n`)
        g.Out()
        g.P(`}`)
    }

    g.sizeVarint()
    g.sizeZigZag()
}
