# gittools
gittools helps us analyze the storage in the git repository.

## idx-format
`idx-format` is used to format the git packfile index for human reading.

`Usage: gittools idx-format <idx file>`

Output example:
```
Version: 2
Fanout: 256
FanoutMapping: 256
Names: 3
Offset32: 3
Offset64: 0
CRC32: 3
PackfileChecksum: 6e53ca9a7f574b96a477e23cabe74ab4808fb063
IdxChecksum: 191baebcfa818863eceecb835d0b49cd6b8a04b5
Fanout[00] = 0
..
Fanout[ff] = 3
FanoutMapping[40] = 0
FanoutMapping[49] = 1
FanoutMapping[e6] = 2
Names[0,0] = 40594824d1cdf9d3683d69fc86b85ee7f3de630d
Names[1,0] = 496d6428b9cf92981dc9495211e6e1120fb6f2ba
Names[2,0] = e69de29bb2d1d6434b8b29ae775ad8c2e48c5391
Offset[0,0] = 12
Offset[1,0] = 804
Offset[2,0] = 795
CRC32[0] = 60d443bf
CRC32[1] = 0f49d649
CRC32[2] = 6e760029
```

## read-pack
`read-pack` helps us read a piece of git packfile binary data with the given location.

`Usage: gittools read-pack <packfile> <offset> <length>`

Output examples:
```
Type = commit
Size = 788
Content bytes length = 62
Content bytes =
00000000  74 72 65 65 20 33 38 34  31 64 30 65 61 35 38 36  |tree 3841d0ea586|
00000010  39 66 35 61 61 36 34 37  32 32 37 39 65 32 35 64  |9f5aa6472279e25d|
00000020  63 64 65 34 38 66 62 34  35 34 61 61 66 0a 70 61  |cde48fb454aaf.pa|
00000030  72 65 6e 74 20 66 64 62  64 62 36 34 66 34        |rent fdbdb64f4|
00000040
Content =
tree 3841d0ea5869f5aa6472279e25dcde48fb454aaf
parent fdbdb64f4
```

```
Type = ofs-delta
Size = 35
Content bytes length = 35
Content bytes =
00000000  02 a7 6a 78 9c 7b d2 cd  f8 a4 9b 71 43 a5 8c c8  |..jx.{.....qC...|
00000010  b1 e3 fb cc 55 8e a6 ca  cd e5 60 50 8c 9d b1 e3  |....U.....`P....|
00000020  e8 3e 25                                          |.>%|
00000030
```