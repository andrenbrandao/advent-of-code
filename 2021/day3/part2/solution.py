import sys


def life_support_rating(data):
    oxygen_rating = oxygen_generator_rating(data)
    co2_rating = co2_scrubber_rating(data)
    return oxygen_rating * co2_rating


def oxygen_generator_rating(data):
    remaining_data = data
    bit_pos = 0

    while len(remaining_data) > 1:
        count1 = 0
        count0 = 0

        for bit_data in remaining_data:
            if bit_data[bit_pos] == "1":
                count1 += 1
            else:
                count0 += 1

        if count1 >= count0:
            remaining_data = [
                bit_data for bit_data in remaining_data if bit_data[bit_pos] == "1"
            ]
        else:
            remaining_data = [
                bit_data for bit_data in remaining_data if bit_data[bit_pos] == "0"
            ]

        bit_pos += 1

    return int(remaining_data[0], 2)


def co2_scrubber_rating(data):
    remaining_data = data
    bit_pos = 0

    while len(remaining_data) > 1:
        count1 = 0
        count0 = 0

        for bit_data in remaining_data:
            if bit_data[bit_pos] == "1":
                count1 += 1
            else:
                count0 += 1

        if count1 >= count0:
            remaining_data = [
                bit_data for bit_data in remaining_data if bit_data[bit_pos] == "0"
            ]
        else:
            remaining_data = [
                bit_data for bit_data in remaining_data if bit_data[bit_pos] == "1"
            ]

        bit_pos += 1

    return int(remaining_data[0], 2)


if __name__ == "__main__":
    input_file = sys.argv[1]

    lines = []
    with open(input_file) as file:
        lines = [line.strip() for line in file]

    print(life_support_rating(lines))
