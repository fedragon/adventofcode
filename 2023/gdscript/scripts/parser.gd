extends Node

class_name LineParser

const numbers = {
	"one": 1,
	"two": 2,
	"three": 3,
	"four": 4,
	"five": 5,
	"six": 6,
	"seven": 7,
	"eight": 8,
	"nine": 9,
}

enum state {READY, PARSING}

func parse(line: String) -> int:
	if line.length() == 0:
		return 0

	var indexes = {}
	for num in numbers:
		var int_value = numbers[num]

		# left->right
		var ix = line.find(num)
		if ix > -1:
			indexes[ix] = int_value
		ix = line.find(String.num(int_value))
		if ix > -1:
			indexes[ix] = int_value

		# right->left
		ix = line.rfind(num)
		if ix > -1:
			indexes[ix] = int_value
		ix = line.rfind(String.num(int_value))
		if ix > -1:
			indexes[ix] = int_value
	
	var min = 999999
	var max = -1
	for key in indexes:
		if key < min:
			min = key
		if key > max:
			max = key
	
	if min == 99999 or max == -1:
		return 0
	return indexes[min] * 10 + indexes[max]
