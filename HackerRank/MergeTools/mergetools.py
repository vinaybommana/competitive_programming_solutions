def merge_the_tools(string, k):
    tools = []
    for index in range(0, len(string), k):
        tools.append(string[index : index + k])

    for tool in tools:
        print(unique_tool(tool))


def unique_tool(string):
    output_string = ""
    occurence_list = []
    for i in string:
        if i not in occurence_list:
            output_string += i
            occurence_list.append(i)

    return output_string


if __name__ == "__main__":
    string, k = input(), int(input())
    merge_the_tools(string, k)
