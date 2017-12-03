#!/usr/bin/env python3

import itertools

with open("input.txt") as indata:
    row = []
    for line in indata:
        numbers = [int(x) for x in line.split("\t")]
        for a, b in itertools.permutations(numbers, 2):
            if a % b == 0:
                row.append(a // b)

print(sum(row))
