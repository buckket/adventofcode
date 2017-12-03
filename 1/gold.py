#!/usr/bin/env python3

indata = "12131415"
outdata = 0

for i in range(0, len(indata)):
    if indata[i] == indata[(i+len(indata)//2)%len(indata)]:
        outdata += int(indata[i])

print(outdata)

