extends GutTest

func test_assert_returns_zero_on_empty_line():
	assert_eq(autofree(LineParser.new()).parse(""), 0)
	
func test_assert_returns_zero_in_absence_of_numbers():
	assert_eq(autofree(LineParser.new()).parse("abc"), 0)

func test_assert_concatenates_single_number_twice():
	assert_eq(autofree(LineParser.new()).parse("1"), 11)

func test_assert_concatenates_two_numbers():
	assert_eq(autofree(LineParser.new()).parse("13"), 13)

func test_assert_concatenates_first_and_last_number():
	assert_eq(autofree(LineParser.new()).parse("123"), 13)
	
func test_assert_concatenates_first_and_last_number_interleaved_by_string():
	assert_eq(autofree(LineParser.new()).parse("a1bc2d3efg"), 13)
	
func test_assert_parses_number_from_string():
	assert_eq(autofree(LineParser.new()).parse("one"), 11)

func test_assert_parses_two_numbers_from_string():
	assert_eq(autofree(LineParser.new()).parse("onethree"), 13)

func test_assert_parses_two_numbers_from_string_interleaved_with_integers():
	assert_eq(autofree(LineParser.new()).parse("one3two"), 12)

func test_assert_parses_real_world_input():
	var test_cases = {
		"two1nine": 29,
		"eightwothree": 83,
		"abcone2threexyz": 13,
		"xtwone3four": 24,
		"4nineeightseven2": 42,
		"zoneight234": 14,
		"7pqrstsixteen": 76,
		"gseightwo6": 86,
		"nine15hjn3nbrteightwoxst": 92,
	}
	
	for line in test_cases:
		var expected = test_cases[line]
		var got = autofree(LineParser.new()).parse(line)
		assert_eq(got, expected, "expected %d, but got %d. line: %s" % [expected, got, line])
