def parse_string(s):
    even_numbers = []
    odd_numbers = []
    lowercaseletters = []
    uppercaseletters = []

    for i in s:
        if i.isdigit():
            if int(i) % 2 == 0:
                even_numbers.append(i)
            else:
                odd_numbers.append(i)
        else:
            if i.islower():
                lowercaseletters.append(i)
            elif i.isupper():
                uppercaseletters.append(i)

    lowercaseletters = sorted(lowercaseletters)
    uppercaseletters = sorted(uppercaseletters)
    odd_numbers = sorted(odd_numbers)
    even_numbers = sorted(even_numbers)

    output_string = ""
    for i in lowercaseletters:
        output_string += i
    for i in uppercaseletters:
        output_string += i
    for i in odd_numbers:
        output_string += i
    for i in even_numbers:
        output_string += i

    print(output_string)


if __name__ == "__main__":
    s = input()
    parse_string(s)
