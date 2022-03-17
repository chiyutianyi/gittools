# gittools
git tools

## idx-format
format git packfile index:

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
read packfile with given offset and length:

```
Type = tree
Size = 29
Content bytes length = 29
Content bytes=
[49 48 48 54 52 52 32 97 0 230 157 226 155 178 209 214 67 75 139 41 174 119 90 216 194 228 140 83 145]
Content =
100644 ...
```