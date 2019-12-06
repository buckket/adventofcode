#!/usr/bin/env python3

import itertools

indata = "91212129"
outdata = 0

for k, g in itertools.groupby(indata + indata[0]):
    outdata += (int(k) * (len(list(g)) - 1))

print(outdata)

