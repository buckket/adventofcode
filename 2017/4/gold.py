#!/usr/bin/env python3

import itertools
import collections

valid_passwords = 0

def is_valid(data):
    c = []
    for word in data:
        c.append(collections.Counter(word))
    if any([x == y for x in c for y in c if x is not y]):
        return False
    return True


with open("input.txt") as indata:
    for line in indata:
        if is_valid(line.rstrip("\n").split(" ")):
            valid_passwords += 1

print(valid_passwords)
