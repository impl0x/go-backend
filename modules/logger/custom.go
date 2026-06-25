package logger

var mo = Level{
	Color: Colors["light_blue"],
	Label: "Mo",
}

func Mo(v ...string) {
	log(mo, v)
}
