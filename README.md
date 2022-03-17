# gittools
git tools

## format
format git packfile index like:

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
FanoutMapping[00] = -1
..
FanoutMapping[ff] = -1
Names[0] = 40594824d1cdf9d3683d69fc86b85ee7f3de630d
Names[1] = 496d6428b9cf92981dc9495211e6e1120fb6f2ba
Names[2] = e69de29bb2d1d6434b8b29ae775ad8c2e48c5391
Offset32[0] = 12
Offset32[1] = 804
Offset32[2] = 795
CRC32[0] = 60d443bf
CRC32[1] = 0f49d649
CRC32[2] = 6e760029
```