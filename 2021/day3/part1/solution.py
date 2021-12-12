import sys


def bin_to_int(bin_num_str):
    result = 0

    exponent = 0
    for i in range(len(bin_num_str) - 1, -1, -1):
        result += (2 ** exponent) * int(bin_num_str[i])
        exponent += 1

    return result


def power_consumption(data):
    width = len(data[0])

    gamma_rate = ["X"] * width
    epsilon_rate = ["X"] * width
    for i in range(width):
        count0 = 0
        count1 = 0
        for value in data:
            if value[i] == "1":
                count1 += 1
            else:
                count0 += 1

        if count0 >= count1:
            gamma_rate[i] = "0"
            epsilon_rate[i] = "1"
        else:
            gamma_rate[i] = "1"
            epsilon_rate[i] = "0"

    decimal_gamma_rate = bin_to_int("".join(gamma_rate))
    decimal_epsilon_rate = bin_to_int("".join(epsilon_rate))

    return decimal_gamma_rate * decimal_epsilon_rate


if __name__ == "__main__":
    input_file = sys.argv[1]

    lines = []
    with open(input_file) as file:
        lines = [line.strip() for line in file]

    print(power_consumption(lines))
