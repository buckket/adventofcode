#!/usr/bin/env python3

import itertools

valid_passwords = 0

def is_valid(data):
    for k, g in itertools.groupby(sorted(data)):
        if len(list(g)) > 1:
            return False
    return True

with open("input.txt") as indata:
    for line in indata:
        if is_valid(line.rstrip("\n").split(" ")):
            print(line.split(" "))
            valid_passwords += 1

print(valid_passwords)
