package sigint

// TODO: Reduce the number of these to specific cases we want to test. I.e. not
// all tests!
var testPatterns = `[
{
  "Signature": "ACDC1",
  "Encoding": "Hexadecimal",
  "Relativity": "beginning of file",
  "Comment": "This should fail as it is an uneven length HEX string",
  "Fail": true
}, {
  "Signature": "424D{4}00000000{4}28000000{8}0100(01|04|08|18|20)00(00|01|02)000000",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "00 61 73 6d",
  "Encoding": "",
  "Relativity": ""
},{
  "Signature": "7573746172",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "CFAD12FE",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "526172211A0700",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "786D6C6E733A7064666169643D(22|27)687474703A2F2F7777772E6169696D2E6F72672F706466612F6E732F6964*7064666169643A70617274(3D22|3D27|3E)33(22|27|3C2F7064666169643A706172743E){0-11}7064666169643A636F6E666F726D616E6365(3E|3D22|3D27)42(22|27|3C2F7064666169643A636F6E666F726D616E63653E)",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "255044462D312E[30:37]",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "63616666",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "4D415243",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "53494D504C4520203D202020202020202020202020202020202020202054",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "0000000C6A5020200D0A870A",
  "Encoding": "hexadecimal",
  "Relativity": "beginning of file"
},{
  "Signature": "255044462D312E33",
  "Encoding": "hexadecimal",
  "Relativity": "end of file"
},{
  "Signature": "2525454F46",
  "Encoding": "hexadecimal",
  "Relativity": "end of file"
},{
  "Signature": "736600FF",
  "Encoding": "hexadecimal",
  "Relativity": "beginning of file"
},{
  "Signature": "664C614300000022",
  "Encoding": "hexadecimal",
  "Relativity": "beginning of file"
},{
  "Signature": "52494646",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "736F6C6964",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "EE0C0D0A",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "424D{4}00000000{4}6C000000{8}0100(01|04|08|10|18|20)00(00|01|02|03)00000000",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "89504E470D0A1A0A0000000D49484452*73524742",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "89504E470D0A1A0A0000000D49484452*73504C54",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "89504E470D0A1A0A0000000D49484452*69434350",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "0000000049454E44AE426082",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "3BF20D0A",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "545A6966",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "GIF89a",
  "Encoding": "ASCII",
  "Relativity": "beginning of file"
},{
  "Signature": "3C3F786D6C2076657273696F6E3D(22|27){8-256}687474703A2F2F7777772E746F706F6772616669782E636F6D2F4750582F312F3122",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "B297E169",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "BLENDER_",
  "Encoding": "",
  "Relativity": ""
},{
  "Signature": "4D41523100",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "255044462D322E30",
  "Encoding": "hexadecimal",
  "Relativity": "end of file"
},{
  "Signature": "2525454F46",
  "Encoding": "hexadecimal",
  "Relativity": "end of file"
},{
  "Signature": "C5D0D3C6",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "252150532D41646F62652D332E3020455053462D332030",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "474946383961",
  "Encoding": "hexadecimal",
  "Relativity": "beginning of file"
},{
  "Signature": "000008A800000000000000940020565732302E30{0-3}565732302E30{0-29}566563746F72576F726B73",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "974A42320D0A1A0A",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "545A6966",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "56657273696F6E20",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "B3F20D0A",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "78617221",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "6C0C0D0A",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "4D534346",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "89504E470D0A1A0A0000000D49484452",
  "Encoding": "hexadecimal",
  "Relativity": "beginning of file"
},{
  "Signature": "0000000049454E44AE426082",
  "Encoding": "hexadecimal",
  "Relativity": "beginning of file"
},{
  "Signature": "494433",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "545A6966",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "2AEB0D0A",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "A0461DF0",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "0F00E803",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "006E1EF0",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "667479704D344120",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "00000020667479704D344120",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "774F4632",
  "Encoding": "hexadecimal",
  "Relativity": "beginning of file"
},{
  "Signature": "ý7zXZ",
  "Encoding": "hexadecimal",
  "Relativity": "beginning of file"
},{
  "Signature": "FD377A585A00",
  "Encoding": "hexadecimal",
  "Relativity": "beginning of file"
},{
  "Signature": "FD 37 7A 58 5A 00",
  "Encoding": "hexadecimal",
  "Relativity": "beginning of file"
},{
  "Signature": "595A",
  "Encoding": "hexadecimal",
  "Relativity": "beginning of file"
},{
  "Signature": "FEEDFACE",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "CEFAEDFE",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "802A5FD7",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "RIFF.{4}WEBPVP8L",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "52494646{4}574542505650384C",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "FFD8FFE0",
  "Encoding": "hexadecimal",
  "Relativity": "beginning of file"
},{
  "Signature": "D7CDC69A",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "3B0C0D0A",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "0000{6}01(01|04|08)",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "9E0C0D0A",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "774F4646",
  "Encoding": "hexadecimal",
  "Relativity": "beginning of file"
},{
  "Signature": "425047FB",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "786D6C6E733A7064666169643D(22|27)687474703A2F2F7777772E6169696D2E6F72672F706466612F6E732F6964*7064666169643A70617274(3D22|3D27|3E)32(22|27|3C2F7064666169643A706172743E){0-11}7064666169643A636F6E666F726D616E6365(3E|3D22|3D27)42(22|27|3C2F7064666169643A636F6E666F726D616E63653E)",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "255044462D312E[30:37]",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "EDABEEDB",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "426C696E6B20627920442E542E53",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "325E1010",
  "Encoding": "",
  "Relativity": ""
},{
  "Signature": "B297E169",
  "Encoding": "",
  "Relativity": ""
},{
  "Signature": "52494646",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "255044462D312E30",
  "Encoding": "hexadecimal",
  "Relativity": "end of file"
},{
  "Signature": "2525454F46",
  "Encoding": "hexadecimal",
  "Relativity": "end of file"
},{
  "Signature": "255044462D312E37",
  "Encoding": "PRONOM internal signature",
  "Relativity": "end of file"
},{
  "Signature": "2525454F(46|460A|460D|460D0A|460D00)",
  "Encoding": "PRONOM internal signature",
  "Relativity": "end of file"
},{
  "Signature": "49491A00000048454150434344520200",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "0011AF",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "2D6C68",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "RIFF.{4}WEBPVP8X",
  "Encoding": "Perl Compatible Regular Expressions",
  "Relativity": "beginning of file"
},{
  "Signature": "87C60D0A",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "7F20575332303030FF312E3030",
  "Encoding": "hexadecimal",
  "Relativity": "beginning of file"
},{
  "Signature": "464C5601",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "000000146674797071742020",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "RIFF.{4}WEBPVP8X",
  "Encoding": "Perl Compatible Regular Expressions",
  "Relativity": "beginning of file"
},{
  "Signature": "255044462D312E32",
  "Encoding": "hexadecimal",
  "Relativity": "end of file"
},{
  "Signature": "2525454F46",
  "Encoding": "hexadecimal",
  "Relativity": "end of file"
},{
  "Signature": "330D0D0A",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "7E424B00",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "786D6C6E733A7064666169643D(22|27)687474703A2F2F7777772E6169696D2E6F72672F706466612F6E732F6964*7064666169643A70617274(3D22|3D27|3E)31(22|27|3C2F7064666169643A706172743E){0-11}7064666169643A636F6E666F726D616E6365(3E|3D22|3D27)41(22|27|3C2F7064666169643A636F6E666F726D616E63653E)",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "255044462D312E[30:37]",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "RIFF.{4}WEBPVP8\\x20",
  "Encoding": "Perl Compatible Regular Expressions",
  "Relativity": "beginning of file"
},{
  "Signature": "52494646{4}5745425056503820",
  "Encoding": "Perl Compatible Regular Expressions",
  "Relativity": "beginning of file"
},{
  "Signature": "256269746D6170",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "424F4F4B4D4F4249",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "D1F20D0A",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "255044462D312E34",
  "Encoding": "hexadecimal",
  "Relativity": "beginning of file"
},{
  "Signature": "2525454F46",
  "Encoding": "hexadecimal",
  "Relativity": "beginning of file"
},{
  "Signature": "786D6C6E733A7064666169643D(22|27)687474703A2F2F7777772E6169696D2E6F72672F706466612F6E732F6964*7064666169643A70617274(3D22|3D27|3E)33(22|27|3C2F7064666169643A706172743E){0-11}7064666169643A636F6E666F726D616E6365(3E|3D22|3D27)41(22|27|3C2F7064666169643A636F6E666F726D616E63653E)",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "255044462D312E[30:37]",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "03F30D0A",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "786D6C6E733A7064666169643D(22|27)687474703A2F2F7777772E6169696D2E6F72672F706466612F6E732F6964*7064666169643A70617274(3D22|3D27|3E)32(22|27|3C2F7064666169643A706172743E){0-11}7064666169643A636F6E666F726D616E6365(3E|3D22|3D27)41(22|27|3C2F7064666169643A636F6E666F726D616E63653E)",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "255044462D312E[30:37]",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "7F454C46",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "FEEDFACF",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "FEEDFACE",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "CFFAEDFE",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "CEFAEDFE",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "D0CF11E0A1B11AE1",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "62706C697374",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "REPROZIP VERSION 2",
  "Encoding": "",
  "Relativity": ""
},{
  "Signature": "526172211A070100",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "52494646",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "444F53",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "160D0D0A",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "43505446494C45",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "4350543746494C45",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "4D4D002B",
  "Encoding": "hexadecimal",
  "Relativity": "beginning of file"
},{
  "Signature": "49492B00",
  "Encoding": "hexadecimal",
  "Relativity": "beginning of file"
},{
  "Signature": "EB3C902A",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "444D5321",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "7801730D626260",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "89504E470D0A1A0A",
  "Encoding": "hexadecimal",
  "Relativity": "beginning of file"
},{
  "Signature": "255044462D312E36",
  "Encoding": "hexadecimal",
  "Relativity": "beginning of file"
},{
  "Signature": "2525454F46",
  "Encoding": "hexadecimal",
  "Relativity": "beginning of file"
},{
  "Signature": "2E736E64",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "060E2B34020501010D0102010102",
  "Encoding": "hexadecimal",
  "Relativity": "beginning of file"
},{
  "Signature": "5F27A889",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "424D{4}00000000{4}0C000000{4}0100(01|04|08|18)00",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "DICM",
  "Encoding": "ASCII",
  "Relativity": ""
},{
  "Signature": "424D{4}00000000{4}7C000000{8}0100(01|04|08|10|18|20)00(00|01|02|03|04|05)00000000",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "4F0C0D0A",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "786D6C6E733A7064666169643D(22|27)687474703A2F2F7777772E6169696D2E6F72672F706466612F6E732F6964*7064666169643A70617274(3D22|3D27|3E)32(22|27|3C2F7064666169643A706172743E){0-11}7064666169643A636F6E666F726D616E6365(3E|3D22|3D27)55(22|27|3C2F7064666169643A636F6E666F726D616E63653E)",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "255044462D312E[30:37]",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "GIF87a",
  "Encoding": "ASCII",
  "Relativity": "beginning of file"
},{
  "Signature": "4D5A*4E45",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "89504E470D0A1A0A0000000D49484452*69545874",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "0000000049454E44AE426082",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "1F8B08",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "437265617469766520566F6963652046",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "FFF1",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "25504446",
  "Encoding": "hexadecimal",
  "Relativity": "beginning of file"
},{
  "Signature": "425A68",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "E310000100000000",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "4C00000001140200",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "4F67675300020000000000000000",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "50350A",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "762F3101",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "FEA1",
  "Encoding": "",
  "Relativity": ""
},{
  "Signature": "786D6C6E733A7064666169643D(22|27)687474703A2F2F7777772E6169696D2E6F72672F706466612F6E732F6964*7064666169643A70617274(3D22|3D27|3E)33(22|27|3C2F7064666169643A706172743E){0-11}7064666169643A636F6E666F726D616E6365(3E|3D22|3D27)55(22|27|3C2F7064666169643A636F6E666F726D616E63653E)",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "255044462D312E[30:37]",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "1A45DFA3{0-32}4282847765626D4287",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "1A45DFA3",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "CAFEBABE",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "AC9EBD8F0000??00",
  "Encoding": "hexadecimal",
  "Relativity": "beginning of file"
},{
  "Signature": "255044462D312E31",
  "Encoding": "hexadecimal",
  "Relativity": "beginning of file"
},{
  "Signature": "2525454F46",
  "Encoding": "hexadecimal",
  "Relativity": "beginning of file"
},{
  "Signature": "52494646{4}57415645",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "60EA",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "67696D702078636620",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "255044462D312E35",
  "Encoding": "hexadecimal",
  "Relativity": "end of file"
},{
  "Signature": "2525454F46",
  "Encoding": "hexadecimal",
  "Relativity": "end of file"
},{
  "Signature": "RIFF.{4}WEBP",
  "Encoding": "Perl Compatible Regular Expressions",
  "Relativity": "beginning of file"
},{
  "Signature": "FEEF",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "6DF20D0A",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "52494646",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "D0 CF 11 E0 A1 B1 1A E1 00",
  "Encoding": "",
  "Relativity": ""
},{
  "Signature": "0x01DA",
  "Encoding": "",
  "Relativity": ""
},{
  "Signature": "000000146674797069736F6D",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "786D6C6E733A7064666169643D(22|27)687474703A2F2F7777772E6169696D2E6F72672F706466612F6E732F6964*7064666169643A636F6E666F726D616E6365(3E|3D22|3D27)42(22|27|3C2F7064666169643A636F6E666F726D616E63653E){0-20}7064666169643A70617274(3D22|3D27|3E)31(22|27|3C2F7064666169643A706172743E)",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "786D6C6E733A7064666169643D(22|27)687474703A2F2F7777772E6169696D2E6F72672F706466612F6E732F6964*7064666169643A70617274(3D22|3D27|3E)31(22|27|3C2F7064666169643A706172743E){0-13}7064666169643A636F6E666F726D616E6365(3E|3D22|3D27)42(22|27|3C2F7064666169643A636F6E666F726D616E63653E)",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "255044462D312E[30:37]",
  "Encoding": "PRONOM internal signature",
  "Relativity": "beginning of file"
},{
  "Signature": "mimetypeapplication/epub+zip",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "504B03040A000200",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "0000FFFFFFFF",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "0001000000",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "464F524D",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "2E524D460000001200",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "636F6E6563746978",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "545A6966",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "5a 49 4d",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "6465780A",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "4D6963726F736F66742057696E646F7773204D6564696120506C61796572202D2D20",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "667479706D703432",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "00000020667479704D345620",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "00000018667479706D703432",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "49492A00100000004352",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "4E45534D1A01",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "4D5A*50450000",
  "Encoding": "hexadecimal",
  "Relativity": "beginning of file"
},{
  "Signature": "4D5A",
  "Encoding": "hexadecimal",
  "Relativity": "beginning of file"
},{
  "Signature": "2DED0D0A",
  "Encoding": "hexadecimal",
  "Relativity": ""
},{
  "Signature": "ECA5C100",
  "Encoding": "hexadecimal",
  "Relativity": ""
}]
`
