// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: components/mil_view.proto

package components

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Disposition int32

const (
	Disposition_DISPOSITION_UNKNOWN          Disposition = 0
	Disposition_DISPOSITION_FRIENDLY         Disposition = 1
	Disposition_DISPOSITION_HOSTILE          Disposition = 2
	Disposition_DISPOSITION_SUSPICIOUS       Disposition = 3
	Disposition_DISPOSITION_ASSUMED_FRIENDLY Disposition = 4
	Disposition_DISPOSITION_NEUTRAL          Disposition = 5
	Disposition_DISPOSITION_PENDING          Disposition = 6
)

// Enum value maps for Disposition.
var (
	Disposition_name = map[int32]string{
		0: "DISPOSITION_UNKNOWN",
		1: "DISPOSITION_FRIENDLY",
		2: "DISPOSITION_HOSTILE",
		3: "DISPOSITION_SUSPICIOUS",
		4: "DISPOSITION_ASSUMED_FRIENDLY",
		5: "DISPOSITION_NEUTRAL",
		6: "DISPOSITION_PENDING",
	}
	Disposition_value = map[string]int32{
		"DISPOSITION_UNKNOWN":          0,
		"DISPOSITION_FRIENDLY":         1,
		"DISPOSITION_HOSTILE":          2,
		"DISPOSITION_SUSPICIOUS":       3,
		"DISPOSITION_ASSUMED_FRIENDLY": 4,
		"DISPOSITION_NEUTRAL":          5,
		"DISPOSITION_PENDING":          6,
	}
)

func (x Disposition) Enum() *Disposition {
	p := new(Disposition)
	*p = x
	return p
}

func (x Disposition) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Disposition) Descriptor() protoreflect.EnumDescriptor {
	return file_components_mil_view_proto_enumTypes[0].Descriptor()
}

func (Disposition) Type() protoreflect.EnumType {
	return &file_components_mil_view_proto_enumTypes[0]
}

func (x Disposition) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Disposition.Descriptor instead.
func (Disposition) EnumDescriptor() ([]byte, []int) {
	return file_components_mil_view_proto_rawDescGZIP(), []int{0}
}

type Environment int32

const (
	Environment_ENVIRONMENT_UNKNOWN     Environment = 0
	Environment_ENVIRONMENT_AIR         Environment = 1
	Environment_ENVIRONMENT_SURFACE     Environment = 2
	Environment_ENVIRONMENT_SUB_SURFACE Environment = 3
	Environment_ENVIRONMENT_LAND        Environment = 4
	Environment_ENVIRONMENT_SPA         Environment = 5
)

// Enum value maps for Environment.
var (
	Environment_name = map[int32]string{
		0: "ENVIRONMENT_UNKNOWN",
		1: "ENVIRONMENT_AIR",
		2: "ENVIRONMENT_SURFACE",
		3: "ENVIRONMENT_SUB_SURFACE",
		4: "ENVIRONMENT_LAND",
		5: "ENVIRONMENT_SPA",
	}
	Environment_value = map[string]int32{
		"ENVIRONMENT_UNKNOWN":     0,
		"ENVIRONMENT_AIR":         1,
		"ENVIRONMENT_SURFACE":     2,
		"ENVIRONMENT_SUB_SURFACE": 3,
		"ENVIRONMENT_LAND":        4,
		"ENVIRONMENT_SPA":         5,
	}
)

func (x Environment) Enum() *Environment {
	p := new(Environment)
	*p = x
	return p
}

func (x Environment) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Environment) Descriptor() protoreflect.EnumDescriptor {
	return file_components_mil_view_proto_enumTypes[1].Descriptor()
}

func (Environment) Type() protoreflect.EnumType {
	return &file_components_mil_view_proto_enumTypes[1]
}

func (x Environment) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Environment.Descriptor instead.
func (Environment) EnumDescriptor() ([]byte, []int) {
	return file_components_mil_view_proto_rawDescGZIP(), []int{1}
}

type Nationality int32

const (
	Nationality_NATIONALITY_INVALID                                   Nationality = 0
	Nationality_NATIONALITY_ALBANIA                                   Nationality = 1
	Nationality_NATIONALITY_ALGERIA                                   Nationality = 2
	Nationality_NATIONALITY_ARGENTINA                                 Nationality = 3
	Nationality_NATIONALITY_ARMENIA                                   Nationality = 4
	Nationality_NATIONALITY_AUSTRALIA                                 Nationality = 5
	Nationality_NATIONALITY_AUSTRIA                                   Nationality = 6
	Nationality_NATIONALITY_AZERBAIJAN                                Nationality = 7
	Nationality_NATIONALITY_BELARUS                                   Nationality = 8
	Nationality_NATIONALITY_BELGIUM                                   Nationality = 9
	Nationality_NATIONALITY_BOLIVIA                                   Nationality = 10
	Nationality_NATIONALITY_BOSNIA_AND_HERZEGOVINA                    Nationality = 11
	Nationality_NATIONALITY_BRAZIL                                    Nationality = 12
	Nationality_NATIONALITY_BULGARIA                                  Nationality = 13
	Nationality_NATIONALITY_CAMBODIA                                  Nationality = 14
	Nationality_NATIONALITY_CANADA                                    Nationality = 15
	Nationality_NATIONALITY_CHILE                                     Nationality = 16
	Nationality_NATIONALITY_CHINA                                     Nationality = 17
	Nationality_NATIONALITY_COLOMBIA                                  Nationality = 18
	Nationality_NATIONALITY_CROATIA                                   Nationality = 19
	Nationality_NATIONALITY_CUBA                                      Nationality = 20
	Nationality_NATIONALITY_CYPRUS                                    Nationality = 21
	Nationality_NATIONALITY_CZECH_REPUBLIC                            Nationality = 22
	Nationality_NATIONALITY_DEMOCRATIC_PEOPLES_REPUBLIC_OF_KOREA      Nationality = 23
	Nationality_NATIONALITY_DENMARK                                   Nationality = 24
	Nationality_NATIONALITY_DOMINICAN_REPUBLIC                        Nationality = 25
	Nationality_NATIONALITY_ECUADOR                                   Nationality = 26
	Nationality_NATIONALITY_EGYPT                                     Nationality = 27
	Nationality_NATIONALITY_ESTONIA                                   Nationality = 28
	Nationality_NATIONALITY_ETHIOPIA                                  Nationality = 29
	Nationality_NATIONALITY_FINLAND                                   Nationality = 30
	Nationality_NATIONALITY_FRANCE                                    Nationality = 31
	Nationality_NATIONALITY_GEORGIA                                   Nationality = 32
	Nationality_NATIONALITY_GERMANY                                   Nationality = 33
	Nationality_NATIONALITY_GREECE                                    Nationality = 34
	Nationality_NATIONALITY_GUATEMALA                                 Nationality = 35
	Nationality_NATIONALITY_GUINEA                                    Nationality = 36
	Nationality_NATIONALITY_HUNGARY                                   Nationality = 37
	Nationality_NATIONALITY_ICELAND                                   Nationality = 38
	Nationality_NATIONALITY_INDIA                                     Nationality = 39
	Nationality_NATIONALITY_INDONESIA                                 Nationality = 40
	Nationality_NATIONALITY_INTERNATIONAL_RED_CROSS                   Nationality = 41
	Nationality_NATIONALITY_IRAQ                                      Nationality = 42
	Nationality_NATIONALITY_IRELAND                                   Nationality = 43
	Nationality_NATIONALITY_ISLAMIC_REPUBLIC_OF_IRAN                  Nationality = 44
	Nationality_NATIONALITY_ISRAEL                                    Nationality = 45
	Nationality_NATIONALITY_ITALY                                     Nationality = 46
	Nationality_NATIONALITY_JAMAICA                                   Nationality = 47
	Nationality_NATIONALITY_JAPAN                                     Nationality = 48
	Nationality_NATIONALITY_JORDAN                                    Nationality = 49
	Nationality_NATIONALITY_KAZAKHSTAN                                Nationality = 50
	Nationality_NATIONALITY_KUWAIT                                    Nationality = 51
	Nationality_NATIONALITY_KYRGHYZ_REPUBLIC                          Nationality = 52
	Nationality_NATIONALITY_LAO_PEOPLES_DEMOCRATIC_REPUBLIC           Nationality = 53
	Nationality_NATIONALITY_LATVIA                                    Nationality = 54
	Nationality_NATIONALITY_LEBANON                                   Nationality = 55
	Nationality_NATIONALITY_LIBERIA                                   Nationality = 56
	Nationality_NATIONALITY_LITHUANIA                                 Nationality = 57
	Nationality_NATIONALITY_LUXEMBOURG                                Nationality = 58
	Nationality_NATIONALITY_MADAGASCAR                                Nationality = 59
	Nationality_NATIONALITY_MALAYSIA                                  Nationality = 60
	Nationality_NATIONALITY_MALTA                                     Nationality = 61
	Nationality_NATIONALITY_MEXICO                                    Nationality = 62
	Nationality_NATIONALITY_MOLDOVA                                   Nationality = 63
	Nationality_NATIONALITY_MONTENEGRO                                Nationality = 64
	Nationality_NATIONALITY_MOROCCO                                   Nationality = 65
	Nationality_NATIONALITY_MYANMAR                                   Nationality = 66
	Nationality_NATIONALITY_NATO                                      Nationality = 67
	Nationality_NATIONALITY_NETHERLANDS                               Nationality = 68
	Nationality_NATIONALITY_NEW_ZEALAND                               Nationality = 69
	Nationality_NATIONALITY_NICARAGUA                                 Nationality = 70
	Nationality_NATIONALITY_NIGERIA                                   Nationality = 71
	Nationality_NATIONALITY_NORWAY                                    Nationality = 72
	Nationality_NATIONALITY_PAKISTAN                                  Nationality = 73
	Nationality_NATIONALITY_PANAMA                                    Nationality = 74
	Nationality_NATIONALITY_PARAGUAY                                  Nationality = 75
	Nationality_NATIONALITY_PERU                                      Nationality = 76
	Nationality_NATIONALITY_PHILIPPINES                               Nationality = 77
	Nationality_NATIONALITY_POLAND                                    Nationality = 78
	Nationality_NATIONALITY_PORTUGAL                                  Nationality = 79
	Nationality_NATIONALITY_REPUBLIC_OF_KOREA                         Nationality = 80
	Nationality_NATIONALITY_ROMANIA                                   Nationality = 81
	Nationality_NATIONALITY_RUSSIA                                    Nationality = 82
	Nationality_NATIONALITY_SAUDI_ARABIA                              Nationality = 83
	Nationality_NATIONALITY_SENEGAL                                   Nationality = 84
	Nationality_NATIONALITY_SERBIA                                    Nationality = 85
	Nationality_NATIONALITY_SINGAPORE                                 Nationality = 86
	Nationality_NATIONALITY_SLOVAKIA                                  Nationality = 87
	Nationality_NATIONALITY_SLOVENIA                                  Nationality = 88
	Nationality_NATIONALITY_SOUTH_AFRICA                              Nationality = 89
	Nationality_NATIONALITY_SPAIN                                     Nationality = 90
	Nationality_NATIONALITY_SUDAN                                     Nationality = 91
	Nationality_NATIONALITY_SWEDEN                                    Nationality = 92
	Nationality_NATIONALITY_SWITZERLAND                               Nationality = 93
	Nationality_NATIONALITY_SYRIAN_ARAB_REPUBLIC                      Nationality = 94
	Nationality_NATIONALITY_TAIWAN                                    Nationality = 95
	Nationality_NATIONALITY_TAJIKISTAN                                Nationality = 96
	Nationality_NATIONALITY_THAILAND                                  Nationality = 97
	Nationality_NATIONALITY_THE_FORMER_YUGOSLAV_REPUBLIC_OF_MACEDONIA Nationality = 98
	Nationality_NATIONALITY_TUNISIA                                   Nationality = 99
	Nationality_NATIONALITY_TURKEY                                    Nationality = 100
	Nationality_NATIONALITY_TURKMENISTAN                              Nationality = 101
	Nationality_NATIONALITY_UGANDA                                    Nationality = 102
	Nationality_NATIONALITY_UKRAINE                                   Nationality = 103
	Nationality_NATIONALITY_UNITED_KINGDOM                            Nationality = 104
	Nationality_NATIONALITY_UNITED_NATIONS                            Nationality = 105
	Nationality_NATIONALITY_UNITED_REPUBLIC_OF_TANZANIA               Nationality = 106
	Nationality_NATIONALITY_UNITED_STATES_OF_AMERICA                  Nationality = 107
	Nationality_NATIONALITY_URUGUAY                                   Nationality = 108
	Nationality_NATIONALITY_UZBEKISTAN                                Nationality = 109
	Nationality_NATIONALITY_VENEZUELA                                 Nationality = 110
	Nationality_NATIONALITY_VIETNAM                                   Nationality = 111
	Nationality_NATIONALITY_YEMEN                                     Nationality = 112
	Nationality_NATIONALITY_ZIMBABWE                                  Nationality = 113
)

// Enum value maps for Nationality.
var (
	Nationality_name = map[int32]string{
		0:   "NATIONALITY_INVALID",
		1:   "NATIONALITY_ALBANIA",
		2:   "NATIONALITY_ALGERIA",
		3:   "NATIONALITY_ARGENTINA",
		4:   "NATIONALITY_ARMENIA",
		5:   "NATIONALITY_AUSTRALIA",
		6:   "NATIONALITY_AUSTRIA",
		7:   "NATIONALITY_AZERBAIJAN",
		8:   "NATIONALITY_BELARUS",
		9:   "NATIONALITY_BELGIUM",
		10:  "NATIONALITY_BOLIVIA",
		11:  "NATIONALITY_BOSNIA_AND_HERZEGOVINA",
		12:  "NATIONALITY_BRAZIL",
		13:  "NATIONALITY_BULGARIA",
		14:  "NATIONALITY_CAMBODIA",
		15:  "NATIONALITY_CANADA",
		16:  "NATIONALITY_CHILE",
		17:  "NATIONALITY_CHINA",
		18:  "NATIONALITY_COLOMBIA",
		19:  "NATIONALITY_CROATIA",
		20:  "NATIONALITY_CUBA",
		21:  "NATIONALITY_CYPRUS",
		22:  "NATIONALITY_CZECH_REPUBLIC",
		23:  "NATIONALITY_DEMOCRATIC_PEOPLES_REPUBLIC_OF_KOREA",
		24:  "NATIONALITY_DENMARK",
		25:  "NATIONALITY_DOMINICAN_REPUBLIC",
		26:  "NATIONALITY_ECUADOR",
		27:  "NATIONALITY_EGYPT",
		28:  "NATIONALITY_ESTONIA",
		29:  "NATIONALITY_ETHIOPIA",
		30:  "NATIONALITY_FINLAND",
		31:  "NATIONALITY_FRANCE",
		32:  "NATIONALITY_GEORGIA",
		33:  "NATIONALITY_GERMANY",
		34:  "NATIONALITY_GREECE",
		35:  "NATIONALITY_GUATEMALA",
		36:  "NATIONALITY_GUINEA",
		37:  "NATIONALITY_HUNGARY",
		38:  "NATIONALITY_ICELAND",
		39:  "NATIONALITY_INDIA",
		40:  "NATIONALITY_INDONESIA",
		41:  "NATIONALITY_INTERNATIONAL_RED_CROSS",
		42:  "NATIONALITY_IRAQ",
		43:  "NATIONALITY_IRELAND",
		44:  "NATIONALITY_ISLAMIC_REPUBLIC_OF_IRAN",
		45:  "NATIONALITY_ISRAEL",
		46:  "NATIONALITY_ITALY",
		47:  "NATIONALITY_JAMAICA",
		48:  "NATIONALITY_JAPAN",
		49:  "NATIONALITY_JORDAN",
		50:  "NATIONALITY_KAZAKHSTAN",
		51:  "NATIONALITY_KUWAIT",
		52:  "NATIONALITY_KYRGHYZ_REPUBLIC",
		53:  "NATIONALITY_LAO_PEOPLES_DEMOCRATIC_REPUBLIC",
		54:  "NATIONALITY_LATVIA",
		55:  "NATIONALITY_LEBANON",
		56:  "NATIONALITY_LIBERIA",
		57:  "NATIONALITY_LITHUANIA",
		58:  "NATIONALITY_LUXEMBOURG",
		59:  "NATIONALITY_MADAGASCAR",
		60:  "NATIONALITY_MALAYSIA",
		61:  "NATIONALITY_MALTA",
		62:  "NATIONALITY_MEXICO",
		63:  "NATIONALITY_MOLDOVA",
		64:  "NATIONALITY_MONTENEGRO",
		65:  "NATIONALITY_MOROCCO",
		66:  "NATIONALITY_MYANMAR",
		67:  "NATIONALITY_NATO",
		68:  "NATIONALITY_NETHERLANDS",
		69:  "NATIONALITY_NEW_ZEALAND",
		70:  "NATIONALITY_NICARAGUA",
		71:  "NATIONALITY_NIGERIA",
		72:  "NATIONALITY_NORWAY",
		73:  "NATIONALITY_PAKISTAN",
		74:  "NATIONALITY_PANAMA",
		75:  "NATIONALITY_PARAGUAY",
		76:  "NATIONALITY_PERU",
		77:  "NATIONALITY_PHILIPPINES",
		78:  "NATIONALITY_POLAND",
		79:  "NATIONALITY_PORTUGAL",
		80:  "NATIONALITY_REPUBLIC_OF_KOREA",
		81:  "NATIONALITY_ROMANIA",
		82:  "NATIONALITY_RUSSIA",
		83:  "NATIONALITY_SAUDI_ARABIA",
		84:  "NATIONALITY_SENEGAL",
		85:  "NATIONALITY_SERBIA",
		86:  "NATIONALITY_SINGAPORE",
		87:  "NATIONALITY_SLOVAKIA",
		88:  "NATIONALITY_SLOVENIA",
		89:  "NATIONALITY_SOUTH_AFRICA",
		90:  "NATIONALITY_SPAIN",
		91:  "NATIONALITY_SUDAN",
		92:  "NATIONALITY_SWEDEN",
		93:  "NATIONALITY_SWITZERLAND",
		94:  "NATIONALITY_SYRIAN_ARAB_REPUBLIC",
		95:  "NATIONALITY_TAIWAN",
		96:  "NATIONALITY_TAJIKISTAN",
		97:  "NATIONALITY_THAILAND",
		98:  "NATIONALITY_THE_FORMER_YUGOSLAV_REPUBLIC_OF_MACEDONIA",
		99:  "NATIONALITY_TUNISIA",
		100: "NATIONALITY_TURKEY",
		101: "NATIONALITY_TURKMENISTAN",
		102: "NATIONALITY_UGANDA",
		103: "NATIONALITY_UKRAINE",
		104: "NATIONALITY_UNITED_KINGDOM",
		105: "NATIONALITY_UNITED_NATIONS",
		106: "NATIONALITY_UNITED_REPUBLIC_OF_TANZANIA",
		107: "NATIONALITY_UNITED_STATES_OF_AMERICA",
		108: "NATIONALITY_URUGUAY",
		109: "NATIONALITY_UZBEKISTAN",
		110: "NATIONALITY_VENEZUELA",
		111: "NATIONALITY_VIETNAM",
		112: "NATIONALITY_YEMEN",
		113: "NATIONALITY_ZIMBABWE",
	}
	Nationality_value = map[string]int32{
		"NATIONALITY_INVALID":                                   0,
		"NATIONALITY_ALBANIA":                                   1,
		"NATIONALITY_ALGERIA":                                   2,
		"NATIONALITY_ARGENTINA":                                 3,
		"NATIONALITY_ARMENIA":                                   4,
		"NATIONALITY_AUSTRALIA":                                 5,
		"NATIONALITY_AUSTRIA":                                   6,
		"NATIONALITY_AZERBAIJAN":                                7,
		"NATIONALITY_BELARUS":                                   8,
		"NATIONALITY_BELGIUM":                                   9,
		"NATIONALITY_BOLIVIA":                                   10,
		"NATIONALITY_BOSNIA_AND_HERZEGOVINA":                    11,
		"NATIONALITY_BRAZIL":                                    12,
		"NATIONALITY_BULGARIA":                                  13,
		"NATIONALITY_CAMBODIA":                                  14,
		"NATIONALITY_CANADA":                                    15,
		"NATIONALITY_CHILE":                                     16,
		"NATIONALITY_CHINA":                                     17,
		"NATIONALITY_COLOMBIA":                                  18,
		"NATIONALITY_CROATIA":                                   19,
		"NATIONALITY_CUBA":                                      20,
		"NATIONALITY_CYPRUS":                                    21,
		"NATIONALITY_CZECH_REPUBLIC":                            22,
		"NATIONALITY_DEMOCRATIC_PEOPLES_REPUBLIC_OF_KOREA":      23,
		"NATIONALITY_DENMARK":                                   24,
		"NATIONALITY_DOMINICAN_REPUBLIC":                        25,
		"NATIONALITY_ECUADOR":                                   26,
		"NATIONALITY_EGYPT":                                     27,
		"NATIONALITY_ESTONIA":                                   28,
		"NATIONALITY_ETHIOPIA":                                  29,
		"NATIONALITY_FINLAND":                                   30,
		"NATIONALITY_FRANCE":                                    31,
		"NATIONALITY_GEORGIA":                                   32,
		"NATIONALITY_GERMANY":                                   33,
		"NATIONALITY_GREECE":                                    34,
		"NATIONALITY_GUATEMALA":                                 35,
		"NATIONALITY_GUINEA":                                    36,
		"NATIONALITY_HUNGARY":                                   37,
		"NATIONALITY_ICELAND":                                   38,
		"NATIONALITY_INDIA":                                     39,
		"NATIONALITY_INDONESIA":                                 40,
		"NATIONALITY_INTERNATIONAL_RED_CROSS":                   41,
		"NATIONALITY_IRAQ":                                      42,
		"NATIONALITY_IRELAND":                                   43,
		"NATIONALITY_ISLAMIC_REPUBLIC_OF_IRAN":                  44,
		"NATIONALITY_ISRAEL":                                    45,
		"NATIONALITY_ITALY":                                     46,
		"NATIONALITY_JAMAICA":                                   47,
		"NATIONALITY_JAPAN":                                     48,
		"NATIONALITY_JORDAN":                                    49,
		"NATIONALITY_KAZAKHSTAN":                                50,
		"NATIONALITY_KUWAIT":                                    51,
		"NATIONALITY_KYRGHYZ_REPUBLIC":                          52,
		"NATIONALITY_LAO_PEOPLES_DEMOCRATIC_REPUBLIC":           53,
		"NATIONALITY_LATVIA":                                    54,
		"NATIONALITY_LEBANON":                                   55,
		"NATIONALITY_LIBERIA":                                   56,
		"NATIONALITY_LITHUANIA":                                 57,
		"NATIONALITY_LUXEMBOURG":                                58,
		"NATIONALITY_MADAGASCAR":                                59,
		"NATIONALITY_MALAYSIA":                                  60,
		"NATIONALITY_MALTA":                                     61,
		"NATIONALITY_MEXICO":                                    62,
		"NATIONALITY_MOLDOVA":                                   63,
		"NATIONALITY_MONTENEGRO":                                64,
		"NATIONALITY_MOROCCO":                                   65,
		"NATIONALITY_MYANMAR":                                   66,
		"NATIONALITY_NATO":                                      67,
		"NATIONALITY_NETHERLANDS":                               68,
		"NATIONALITY_NEW_ZEALAND":                               69,
		"NATIONALITY_NICARAGUA":                                 70,
		"NATIONALITY_NIGERIA":                                   71,
		"NATIONALITY_NORWAY":                                    72,
		"NATIONALITY_PAKISTAN":                                  73,
		"NATIONALITY_PANAMA":                                    74,
		"NATIONALITY_PARAGUAY":                                  75,
		"NATIONALITY_PERU":                                      76,
		"NATIONALITY_PHILIPPINES":                               77,
		"NATIONALITY_POLAND":                                    78,
		"NATIONALITY_PORTUGAL":                                  79,
		"NATIONALITY_REPUBLIC_OF_KOREA":                         80,
		"NATIONALITY_ROMANIA":                                   81,
		"NATIONALITY_RUSSIA":                                    82,
		"NATIONALITY_SAUDI_ARABIA":                              83,
		"NATIONALITY_SENEGAL":                                   84,
		"NATIONALITY_SERBIA":                                    85,
		"NATIONALITY_SINGAPORE":                                 86,
		"NATIONALITY_SLOVAKIA":                                  87,
		"NATIONALITY_SLOVENIA":                                  88,
		"NATIONALITY_SOUTH_AFRICA":                              89,
		"NATIONALITY_SPAIN":                                     90,
		"NATIONALITY_SUDAN":                                     91,
		"NATIONALITY_SWEDEN":                                    92,
		"NATIONALITY_SWITZERLAND":                               93,
		"NATIONALITY_SYRIAN_ARAB_REPUBLIC":                      94,
		"NATIONALITY_TAIWAN":                                    95,
		"NATIONALITY_TAJIKISTAN":                                96,
		"NATIONALITY_THAILAND":                                  97,
		"NATIONALITY_THE_FORMER_YUGOSLAV_REPUBLIC_OF_MACEDONIA": 98,
		"NATIONALITY_TUNISIA":                                   99,
		"NATIONALITY_TURKEY":                                    100,
		"NATIONALITY_TURKMENISTAN":                              101,
		"NATIONALITY_UGANDA":                                    102,
		"NATIONALITY_UKRAINE":                                   103,
		"NATIONALITY_UNITED_KINGDOM":                            104,
		"NATIONALITY_UNITED_NATIONS":                            105,
		"NATIONALITY_UNITED_REPUBLIC_OF_TANZANIA":               106,
		"NATIONALITY_UNITED_STATES_OF_AMERICA":                  107,
		"NATIONALITY_URUGUAY":                                   108,
		"NATIONALITY_UZBEKISTAN":                                109,
		"NATIONALITY_VENEZUELA":                                 110,
		"NATIONALITY_VIETNAM":                                   111,
		"NATIONALITY_YEMEN":                                     112,
		"NATIONALITY_ZIMBABWE":                                  113,
	}
)

func (x Nationality) Enum() *Nationality {
	p := new(Nationality)
	*p = x
	return p
}

func (x Nationality) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Nationality) Descriptor() protoreflect.EnumDescriptor {
	return file_components_mil_view_proto_enumTypes[2].Descriptor()
}

func (Nationality) Type() protoreflect.EnumType {
	return &file_components_mil_view_proto_enumTypes[2]
}

func (x Nationality) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Nationality.Descriptor instead.
func (Nationality) EnumDescriptor() ([]byte, []int) {
	return file_components_mil_view_proto_rawDescGZIP(), []int{2}
}

type MilView struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Disposition   *Disposition           `protobuf:"varint,1,opt,name=disposition,proto3,enum=components.Disposition,oneof" json:"disposition,omitempty"`
	Environment   *Environment           `protobuf:"varint,2,opt,name=environment,proto3,enum=components.Environment,oneof" json:"environment,omitempty"`
	Nationality   *Nationality           `protobuf:"varint,3,opt,name=nationality,proto3,enum=components.Nationality,oneof" json:"nationality,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MilView) Reset() {
	*x = MilView{}
	mi := &file_components_mil_view_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MilView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MilView) ProtoMessage() {}

func (x *MilView) ProtoReflect() protoreflect.Message {
	mi := &file_components_mil_view_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MilView.ProtoReflect.Descriptor instead.
func (*MilView) Descriptor() ([]byte, []int) {
	return file_components_mil_view_proto_rawDescGZIP(), []int{0}
}

func (x *MilView) GetDisposition() Disposition {
	if x != nil && x.Disposition != nil {
		return *x.Disposition
	}
	return Disposition_DISPOSITION_UNKNOWN
}

func (x *MilView) GetEnvironment() Environment {
	if x != nil && x.Environment != nil {
		return *x.Environment
	}
	return Environment_ENVIRONMENT_UNKNOWN
}

func (x *MilView) GetNationality() Nationality {
	if x != nil && x.Nationality != nil {
		return *x.Nationality
	}
	return Nationality_NATIONALITY_INVALID
}

var File_components_mil_view_proto protoreflect.FileDescriptor

const file_components_mil_view_proto_rawDesc = "" +
	"\n" +
	"\x19components/mil_view.proto\x12\n" +
	"components\"\xf9\x01\n" +
	"\aMilView\x12>\n" +
	"\vdisposition\x18\x01 \x01(\x0e2\x17.components.DispositionH\x00R\vdisposition\x88\x01\x01\x12>\n" +
	"\venvironment\x18\x02 \x01(\x0e2\x17.components.EnvironmentH\x01R\venvironment\x88\x01\x01\x12>\n" +
	"\vnationality\x18\x03 \x01(\x0e2\x17.components.NationalityH\x02R\vnationality\x88\x01\x01B\x0e\n" +
	"\f_dispositionB\x0e\n" +
	"\f_environmentB\x0e\n" +
	"\f_nationality*\xc9\x01\n" +
	"\vDisposition\x12\x17\n" +
	"\x13DISPOSITION_UNKNOWN\x10\x00\x12\x18\n" +
	"\x14DISPOSITION_FRIENDLY\x10\x01\x12\x17\n" +
	"\x13DISPOSITION_HOSTILE\x10\x02\x12\x1a\n" +
	"\x16DISPOSITION_SUSPICIOUS\x10\x03\x12 \n" +
	"\x1cDISPOSITION_ASSUMED_FRIENDLY\x10\x04\x12\x17\n" +
	"\x13DISPOSITION_NEUTRAL\x10\x05\x12\x17\n" +
	"\x13DISPOSITION_PENDING\x10\x06*\x9c\x01\n" +
	"\vEnvironment\x12\x17\n" +
	"\x13ENVIRONMENT_UNKNOWN\x10\x00\x12\x13\n" +
	"\x0fENVIRONMENT_AIR\x10\x01\x12\x17\n" +
	"\x13ENVIRONMENT_SURFACE\x10\x02\x12\x1b\n" +
	"\x17ENVIRONMENT_SUB_SURFACE\x10\x03\x12\x14\n" +
	"\x10ENVIRONMENT_LAND\x10\x04\x12\x13\n" +
	"\x0fENVIRONMENT_SPA\x10\x05*\xb7\x18\n" +
	"\vNationality\x12\x17\n" +
	"\x13NATIONALITY_INVALID\x10\x00\x12\x17\n" +
	"\x13NATIONALITY_ALBANIA\x10\x01\x12\x17\n" +
	"\x13NATIONALITY_ALGERIA\x10\x02\x12\x19\n" +
	"\x15NATIONALITY_ARGENTINA\x10\x03\x12\x17\n" +
	"\x13NATIONALITY_ARMENIA\x10\x04\x12\x19\n" +
	"\x15NATIONALITY_AUSTRALIA\x10\x05\x12\x17\n" +
	"\x13NATIONALITY_AUSTRIA\x10\x06\x12\x1a\n" +
	"\x16NATIONALITY_AZERBAIJAN\x10\a\x12\x17\n" +
	"\x13NATIONALITY_BELARUS\x10\b\x12\x17\n" +
	"\x13NATIONALITY_BELGIUM\x10\t\x12\x17\n" +
	"\x13NATIONALITY_BOLIVIA\x10\n" +
	"\x12&\n" +
	"\"NATIONALITY_BOSNIA_AND_HERZEGOVINA\x10\v\x12\x16\n" +
	"\x12NATIONALITY_BRAZIL\x10\f\x12\x18\n" +
	"\x14NATIONALITY_BULGARIA\x10\r\x12\x18\n" +
	"\x14NATIONALITY_CAMBODIA\x10\x0e\x12\x16\n" +
	"\x12NATIONALITY_CANADA\x10\x0f\x12\x15\n" +
	"\x11NATIONALITY_CHILE\x10\x10\x12\x15\n" +
	"\x11NATIONALITY_CHINA\x10\x11\x12\x18\n" +
	"\x14NATIONALITY_COLOMBIA\x10\x12\x12\x17\n" +
	"\x13NATIONALITY_CROATIA\x10\x13\x12\x14\n" +
	"\x10NATIONALITY_CUBA\x10\x14\x12\x16\n" +
	"\x12NATIONALITY_CYPRUS\x10\x15\x12\x1e\n" +
	"\x1aNATIONALITY_CZECH_REPUBLIC\x10\x16\x124\n" +
	"0NATIONALITY_DEMOCRATIC_PEOPLES_REPUBLIC_OF_KOREA\x10\x17\x12\x17\n" +
	"\x13NATIONALITY_DENMARK\x10\x18\x12\"\n" +
	"\x1eNATIONALITY_DOMINICAN_REPUBLIC\x10\x19\x12\x17\n" +
	"\x13NATIONALITY_ECUADOR\x10\x1a\x12\x15\n" +
	"\x11NATIONALITY_EGYPT\x10\x1b\x12\x17\n" +
	"\x13NATIONALITY_ESTONIA\x10\x1c\x12\x18\n" +
	"\x14NATIONALITY_ETHIOPIA\x10\x1d\x12\x17\n" +
	"\x13NATIONALITY_FINLAND\x10\x1e\x12\x16\n" +
	"\x12NATIONALITY_FRANCE\x10\x1f\x12\x17\n" +
	"\x13NATIONALITY_GEORGIA\x10 \x12\x17\n" +
	"\x13NATIONALITY_GERMANY\x10!\x12\x16\n" +
	"\x12NATIONALITY_GREECE\x10\"\x12\x19\n" +
	"\x15NATIONALITY_GUATEMALA\x10#\x12\x16\n" +
	"\x12NATIONALITY_GUINEA\x10$\x12\x17\n" +
	"\x13NATIONALITY_HUNGARY\x10%\x12\x17\n" +
	"\x13NATIONALITY_ICELAND\x10&\x12\x15\n" +
	"\x11NATIONALITY_INDIA\x10'\x12\x19\n" +
	"\x15NATIONALITY_INDONESIA\x10(\x12'\n" +
	"#NATIONALITY_INTERNATIONAL_RED_CROSS\x10)\x12\x14\n" +
	"\x10NATIONALITY_IRAQ\x10*\x12\x17\n" +
	"\x13NATIONALITY_IRELAND\x10+\x12(\n" +
	"$NATIONALITY_ISLAMIC_REPUBLIC_OF_IRAN\x10,\x12\x16\n" +
	"\x12NATIONALITY_ISRAEL\x10-\x12\x15\n" +
	"\x11NATIONALITY_ITALY\x10.\x12\x17\n" +
	"\x13NATIONALITY_JAMAICA\x10/\x12\x15\n" +
	"\x11NATIONALITY_JAPAN\x100\x12\x16\n" +
	"\x12NATIONALITY_JORDAN\x101\x12\x1a\n" +
	"\x16NATIONALITY_KAZAKHSTAN\x102\x12\x16\n" +
	"\x12NATIONALITY_KUWAIT\x103\x12 \n" +
	"\x1cNATIONALITY_KYRGHYZ_REPUBLIC\x104\x12/\n" +
	"+NATIONALITY_LAO_PEOPLES_DEMOCRATIC_REPUBLIC\x105\x12\x16\n" +
	"\x12NATIONALITY_LATVIA\x106\x12\x17\n" +
	"\x13NATIONALITY_LEBANON\x107\x12\x17\n" +
	"\x13NATIONALITY_LIBERIA\x108\x12\x19\n" +
	"\x15NATIONALITY_LITHUANIA\x109\x12\x1a\n" +
	"\x16NATIONALITY_LUXEMBOURG\x10:\x12\x1a\n" +
	"\x16NATIONALITY_MADAGASCAR\x10;\x12\x18\n" +
	"\x14NATIONALITY_MALAYSIA\x10<\x12\x15\n" +
	"\x11NATIONALITY_MALTA\x10=\x12\x16\n" +
	"\x12NATIONALITY_MEXICO\x10>\x12\x17\n" +
	"\x13NATIONALITY_MOLDOVA\x10?\x12\x1a\n" +
	"\x16NATIONALITY_MONTENEGRO\x10@\x12\x17\n" +
	"\x13NATIONALITY_MOROCCO\x10A\x12\x17\n" +
	"\x13NATIONALITY_MYANMAR\x10B\x12\x14\n" +
	"\x10NATIONALITY_NATO\x10C\x12\x1b\n" +
	"\x17NATIONALITY_NETHERLANDS\x10D\x12\x1b\n" +
	"\x17NATIONALITY_NEW_ZEALAND\x10E\x12\x19\n" +
	"\x15NATIONALITY_NICARAGUA\x10F\x12\x17\n" +
	"\x13NATIONALITY_NIGERIA\x10G\x12\x16\n" +
	"\x12NATIONALITY_NORWAY\x10H\x12\x18\n" +
	"\x14NATIONALITY_PAKISTAN\x10I\x12\x16\n" +
	"\x12NATIONALITY_PANAMA\x10J\x12\x18\n" +
	"\x14NATIONALITY_PARAGUAY\x10K\x12\x14\n" +
	"\x10NATIONALITY_PERU\x10L\x12\x1b\n" +
	"\x17NATIONALITY_PHILIPPINES\x10M\x12\x16\n" +
	"\x12NATIONALITY_POLAND\x10N\x12\x18\n" +
	"\x14NATIONALITY_PORTUGAL\x10O\x12!\n" +
	"\x1dNATIONALITY_REPUBLIC_OF_KOREA\x10P\x12\x17\n" +
	"\x13NATIONALITY_ROMANIA\x10Q\x12\x16\n" +
	"\x12NATIONALITY_RUSSIA\x10R\x12\x1c\n" +
	"\x18NATIONALITY_SAUDI_ARABIA\x10S\x12\x17\n" +
	"\x13NATIONALITY_SENEGAL\x10T\x12\x16\n" +
	"\x12NATIONALITY_SERBIA\x10U\x12\x19\n" +
	"\x15NATIONALITY_SINGAPORE\x10V\x12\x18\n" +
	"\x14NATIONALITY_SLOVAKIA\x10W\x12\x18\n" +
	"\x14NATIONALITY_SLOVENIA\x10X\x12\x1c\n" +
	"\x18NATIONALITY_SOUTH_AFRICA\x10Y\x12\x15\n" +
	"\x11NATIONALITY_SPAIN\x10Z\x12\x15\n" +
	"\x11NATIONALITY_SUDAN\x10[\x12\x16\n" +
	"\x12NATIONALITY_SWEDEN\x10\\\x12\x1b\n" +
	"\x17NATIONALITY_SWITZERLAND\x10]\x12$\n" +
	" NATIONALITY_SYRIAN_ARAB_REPUBLIC\x10^\x12\x16\n" +
	"\x12NATIONALITY_TAIWAN\x10_\x12\x1a\n" +
	"\x16NATIONALITY_TAJIKISTAN\x10`\x12\x18\n" +
	"\x14NATIONALITY_THAILAND\x10a\x129\n" +
	"5NATIONALITY_THE_FORMER_YUGOSLAV_REPUBLIC_OF_MACEDONIA\x10b\x12\x17\n" +
	"\x13NATIONALITY_TUNISIA\x10c\x12\x16\n" +
	"\x12NATIONALITY_TURKEY\x10d\x12\x1c\n" +
	"\x18NATIONALITY_TURKMENISTAN\x10e\x12\x16\n" +
	"\x12NATIONALITY_UGANDA\x10f\x12\x17\n" +
	"\x13NATIONALITY_UKRAINE\x10g\x12\x1e\n" +
	"\x1aNATIONALITY_UNITED_KINGDOM\x10h\x12\x1e\n" +
	"\x1aNATIONALITY_UNITED_NATIONS\x10i\x12+\n" +
	"'NATIONALITY_UNITED_REPUBLIC_OF_TANZANIA\x10j\x12(\n" +
	"$NATIONALITY_UNITED_STATES_OF_AMERICA\x10k\x12\x17\n" +
	"\x13NATIONALITY_URUGUAY\x10l\x12\x1a\n" +
	"\x16NATIONALITY_UZBEKISTAN\x10m\x12\x19\n" +
	"\x15NATIONALITY_VENEZUELA\x10n\x12\x17\n" +
	"\x13NATIONALITY_VIETNAM\x10o\x12\x15\n" +
	"\x11NATIONALITY_YEMEN\x10p\x12\x18\n" +
	"\x14NATIONALITY_ZIMBABWE\x10qB\rZ\v/componentsb\x06proto3"

var (
	file_components_mil_view_proto_rawDescOnce sync.Once
	file_components_mil_view_proto_rawDescData []byte
)

func file_components_mil_view_proto_rawDescGZIP() []byte {
	file_components_mil_view_proto_rawDescOnce.Do(func() {
		file_components_mil_view_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_components_mil_view_proto_rawDesc), len(file_components_mil_view_proto_rawDesc)))
	})
	return file_components_mil_view_proto_rawDescData
}

var file_components_mil_view_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_components_mil_view_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_components_mil_view_proto_goTypes = []any{
	(Disposition)(0), // 0: components.Disposition
	(Environment)(0), // 1: components.Environment
	(Nationality)(0), // 2: components.Nationality
	(*MilView)(nil),  // 3: components.MilView
}
var file_components_mil_view_proto_depIdxs = []int32{
	0, // 0: components.MilView.disposition:type_name -> components.Disposition
	1, // 1: components.MilView.environment:type_name -> components.Environment
	2, // 2: components.MilView.nationality:type_name -> components.Nationality
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_components_mil_view_proto_init() }
func file_components_mil_view_proto_init() {
	if File_components_mil_view_proto != nil {
		return
	}
	file_components_mil_view_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_components_mil_view_proto_rawDesc), len(file_components_mil_view_proto_rawDesc)),
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_components_mil_view_proto_goTypes,
		DependencyIndexes: file_components_mil_view_proto_depIdxs,
		EnumInfos:         file_components_mil_view_proto_enumTypes,
		MessageInfos:      file_components_mil_view_proto_msgTypes,
	}.Build()
	File_components_mil_view_proto = out.File
	file_components_mil_view_proto_goTypes = nil
	file_components_mil_view_proto_depIdxs = nil
}
