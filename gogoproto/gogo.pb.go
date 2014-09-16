// Code generated by protoc-gen-dgo.
// source: gogo.proto
// DO NOT EDIT!

/*
Package gogoproto is a generated protocol buffer package.

It is generated from these files:
	gogo.proto

It has these top-level messages:
*/
package gogoproto

import proto "github.com/dropbox/goprotoc/proto"
import math "math"

// renamed import google/protobuf/descriptor to code.google.com/p/gogoprotobuf/protoc-gen-dgo/descriptor
import google_protobuf "github.com/dropbox/goprotoc/protoc-gen-dgo/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

var E_GoprotoEnumPrefix = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.EnumOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         62001,
    Name:          "gogoproto.goproto_enum_prefix",
    Tag:           "varint,62001,opt,name=goproto_enum_prefix",
}

var E_GoprotoEnumStringer = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.EnumOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         62021,
    Name:          "gogoproto.goproto_enum_stringer",
    Tag:           "varint,62021,opt,name=goproto_enum_stringer",
}

var E_EnumStringer = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.EnumOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         62022,
    Name:          "gogoproto.enum_stringer",
    Tag:           "varint,62022,opt,name=enum_stringer",
}

var E_GoprotoEnumPrefixAll = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.FileOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         63002,
    Name:          "gogoproto.goproto_enum_prefix_all",
    Tag:           "varint,63002,opt,name=goproto_enum_prefix_all",
}

var E_VerboseEqualAll = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.FileOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         63004,
    Name:          "gogoproto.verbose_equal_all",
    Tag:           "varint,63004,opt,name=verbose_equal_all",
}

var E_FaceAll = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.FileOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         63005,
    Name:          "gogoproto.face_all",
    Tag:           "varint,63005,opt,name=face_all",
}

var E_PopulateAll = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.FileOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         63007,
    Name:          "gogoproto.populate_all",
    Tag:           "varint,63007,opt,name=populate_all",
}

var E_StringerAll = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.FileOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         63008,
    Name:          "gogoproto.stringer_all",
    Tag:           "varint,63008,opt,name=stringer_all",
}

var E_OnlyoneAll = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.FileOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         63009,
    Name:          "gogoproto.onlyone_all",
    Tag:           "varint,63009,opt,name=onlyone_all",
}

var E_EqualAll = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.FileOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         63013,
    Name:          "gogoproto.equal_all",
    Tag:           "varint,63013,opt,name=equal_all",
}

var E_DescriptionAll = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.FileOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         63014,
    Name:          "gogoproto.description_all",
    Tag:           "varint,63014,opt,name=description_all",
}

var E_TestgenAll = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.FileOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         63015,
    Name:          "gogoproto.testgen_all",
    Tag:           "varint,63015,opt,name=testgen_all",
}

var E_BenchgenAll = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.FileOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         63016,
    Name:          "gogoproto.benchgen_all",
    Tag:           "varint,63016,opt,name=benchgen_all",
}

var E_MarshalerAll = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.FileOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         63017,
    Name:          "gogoproto.marshaler_all",
    Tag:           "varint,63017,opt,name=marshaler_all",
}

var E_UnmarshalerAll = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.FileOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         63018,
    Name:          "gogoproto.unmarshaler_all",
    Tag:           "varint,63018,opt,name=unmarshaler_all",
}

var E_BuffertoAll = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.FileOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         63019,
    Name:          "gogoproto.bufferto_all",
    Tag:           "varint,63019,opt,name=bufferto_all",
}

var E_SetterAll = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.FieldOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         63020,
    Name:          "gogoproto.setter_all",
    Tag:           "varint,63020,opt,name=setter_all",
}

var E_SizerAll = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.FileOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         63020,
    Name:          "gogoproto.sizer_all",
    Tag:           "varint,63020,opt,name=sizer_all",
}

var E_GoprotoEnumStringerAll = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.FileOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         63021,
    Name:          "gogoproto.goproto_enum_stringer_all",
    Tag:           "varint,63021,opt,name=goproto_enum_stringer_all",
}

var E_EnumStringerAll = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.FileOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         63022,
    Name:          "gogoproto.enum_stringer_all",
    Tag:           "varint,63022,opt,name=enum_stringer_all",
}

var E_GoprotoExtensionsMapAll = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.FileOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         63025,
    Name:          "gogoproto.goproto_extensions_map_all",
    Tag:           "varint,63025,opt,name=goproto_extensions_map_all",
}

var E_VerboseEqual = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.MessageOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         64004,
    Name:          "gogoproto.verbose_equal",
    Tag:           "varint,64004,opt,name=verbose_equal",
}

var E_Face = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.MessageOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         64005,
    Name:          "gogoproto.face",
    Tag:           "varint,64005,opt,name=face",
}

var E_Populate = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.MessageOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         64007,
    Name:          "gogoproto.populate",
    Tag:           "varint,64007,opt,name=populate",
}

var E_Stringer = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.MessageOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         67008,
    Name:          "gogoproto.stringer",
    Tag:           "varint,67008,opt,name=stringer",
}

var E_Onlyone = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.MessageOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         64009,
    Name:          "gogoproto.onlyone",
    Tag:           "varint,64009,opt,name=onlyone",
}

var E_Equal = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.MessageOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         64013,
    Name:          "gogoproto.equal",
    Tag:           "varint,64013,opt,name=equal",
}

var E_Description = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.MessageOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         64014,
    Name:          "gogoproto.description",
    Tag:           "varint,64014,opt,name=description",
}

var E_Testgen = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.MessageOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         64015,
    Name:          "gogoproto.testgen",
    Tag:           "varint,64015,opt,name=testgen",
}

var E_Benchgen = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.MessageOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         64016,
    Name:          "gogoproto.benchgen",
    Tag:           "varint,64016,opt,name=benchgen",
}

var E_Marshaler = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.MessageOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         64017,
    Name:          "gogoproto.marshaler",
    Tag:           "varint,64017,opt,name=marshaler",
}

var E_Unmarshaler = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.MessageOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         64018,
    Name:          "gogoproto.unmarshaler",
    Tag:           "varint,64018,opt,name=unmarshaler",
}

var E_Bufferto = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.MessageOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         64019,
    Name:          "gogoproto.bufferto",
    Tag:           "varint,64019,opt,name=bufferto",
}

var E_Sizer = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.MessageOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         64020,
    Name:          "gogoproto.sizer",
    Tag:           "varint,64020,opt,name=sizer",
}

var E_GoprotoExtensionsMap = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.MessageOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         64025,
    Name:          "gogoproto.goproto_extensions_map",
    Tag:           "varint,64025,opt,name=goproto_extensions_map",
}

var E_Nullable = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.FieldOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         65001,
    Name:          "gogoproto.nullable",
    Tag:           "varint,65001,opt,name=nullable",
}

var E_Embed = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.FieldOptions)(nil),
    ExtensionType: (*bool)(nil),
    Field:         65002,
    Name:          "gogoproto.embed",
    Tag:           "varint,65002,opt,name=embed",
}

var E_Customtype = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.FieldOptions)(nil),
    ExtensionType: (*string)(nil),
    Field:         65003,
    Name:          "gogoproto.customtype",
    Tag:           "bytes,65003,opt,name=customtype",
}

var E_Customname = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.FieldOptions)(nil),
    ExtensionType: (*string)(nil),
    Field:         65004,
    Name:          "gogoproto.customname",
    Tag:           "bytes,65004,opt,name=customname",
}

var E_Jsontag = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.FieldOptions)(nil),
    ExtensionType: (*string)(nil),
    Field:         65005,
    Name:          "gogoproto.jsontag",
    Tag:           "bytes,65005,opt,name=jsontag",
}

var E_Moretags = &proto.ExtensionDesc{
    ExtendedType:  (*google_protobuf.FieldOptions)(nil),
    ExtensionType: (*string)(nil),
    Field:         65006,
    Name:          "gogoproto.moretags",
    Tag:           "bytes,65006,opt,name=moretags",
}

func init() {
    proto.RegisterExtension(E_GoprotoEnumPrefix)
    proto.RegisterExtension(E_GoprotoEnumStringer)
    proto.RegisterExtension(E_EnumStringer)
    proto.RegisterExtension(E_GoprotoEnumPrefixAll)
    proto.RegisterExtension(E_VerboseEqualAll)
    proto.RegisterExtension(E_FaceAll)
    proto.RegisterExtension(E_PopulateAll)
    proto.RegisterExtension(E_StringerAll)
    proto.RegisterExtension(E_OnlyoneAll)
    proto.RegisterExtension(E_EqualAll)
    proto.RegisterExtension(E_DescriptionAll)
    proto.RegisterExtension(E_TestgenAll)
    proto.RegisterExtension(E_BenchgenAll)
    proto.RegisterExtension(E_MarshalerAll)
    proto.RegisterExtension(E_UnmarshalerAll)
    proto.RegisterExtension(E_BuffertoAll)
    proto.RegisterExtension(E_SizerAll)
    proto.RegisterExtension(E_GoprotoEnumStringerAll)
    proto.RegisterExtension(E_EnumStringerAll)
    proto.RegisterExtension(E_GoprotoExtensionsMapAll)
    proto.RegisterExtension(E_VerboseEqual)
    proto.RegisterExtension(E_Face)
    proto.RegisterExtension(E_Populate)
    proto.RegisterExtension(E_Stringer)
    proto.RegisterExtension(E_Onlyone)
    proto.RegisterExtension(E_Equal)
    proto.RegisterExtension(E_Description)
    proto.RegisterExtension(E_Testgen)
    proto.RegisterExtension(E_Benchgen)
    proto.RegisterExtension(E_Marshaler)
    proto.RegisterExtension(E_Unmarshaler)
    proto.RegisterExtension(E_Bufferto)
    proto.RegisterExtension(E_Sizer)
    proto.RegisterExtension(E_GoprotoExtensionsMap)
    proto.RegisterExtension(E_Nullable)
    proto.RegisterExtension(E_Embed)
    proto.RegisterExtension(E_Customtype)
    proto.RegisterExtension(E_Customname)
    proto.RegisterExtension(E_Jsontag)
    proto.RegisterExtension(E_Moretags)
}