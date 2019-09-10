#!/bin/python3

import math
import os
import random
import re
import sys


def parse_string(string: str):
    occurence_dict = {}
    for char in string:
        if char in occurence_dict:
            occurence_dict[char] += 1
        else:
            occurence_dict[char] = 1

    occurence_list = []
    for char, times in occurence_dict.items():
        occurence_list.append((char, times))

    occurence_list = sorted(occurence_list, key=lambda x: (-x[1], x[0]))

    occurence_list = occurence_list[:3]
    for item, times in occurence_list:
        print(item, times)


if __name__ == "__main__":
    s = input()
