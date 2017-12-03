#!/usr/bin/env python3

with open("input.txt") as indata:
    row = []
    for line in indata:
        low = None
        high = None
        for d in line.split("\t"):
            d = int(d)
            if high is None or d > high:
                high = d
            if  low is None or d < low:
                low = d
        row.append(high - low)

print(sum(row))
