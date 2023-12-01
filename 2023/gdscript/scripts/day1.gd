extends Node

@onready var label = $CenterContainer/RichTextLabel


func _ready():
	var file = FileAccess.open("res://inputs/day1.txt", FileAccess.READ)
	var total = 0
	while not file.eof_reached():
		var line = file.get_line()
		if line.length() > 0:
			var line_total = LineParser.new().parse(line)
			total += line_total
	label.text = String.num(total, 0)
