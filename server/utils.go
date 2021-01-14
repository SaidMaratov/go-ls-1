package server

const welcomeIcon = "" +
	"Welcome to TCP-Chat!\n" +
	orange +
	"         _nnnn_\n" +
	"        dGGGGMMb\n" +
	"       @p~qp~~qMb\n" +
	"       M|@||@) M|\n" +
	"       @,----.JM|\n" +
	"      JS^\\__/  qKL\n" +
	"     dZP        qKRb\n" +
	"    dZP          qKKb\n" +
	"   fZP            SMMb\n" +
	"   HZM            MMMM\n" +
	"   FqM            MMMM\n" +
	" __| \".        |\\dS\"qML\n" +
	" |    `.       | `' \\Zq\n" +
	"_)      \\.___.,|     .'\n" +
	"\\____   )MMMMMP|   .'\n" +
	"     `-'       `--'\n\n" +
	reset +
	"Your default name is anonymous\n" +
	"You are in the default room \"general\"!\n"

const manual = "Manual:\n\n" +

	"/nick [nickname] - create or change nickname\n" +
	"/rooms - to list the available rooms\n" +
	"/join [name of room] - to create a new room or join the available room\n" +
	"/quit - to leave the server\n\n"

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	purple = "\033[35m"
	cyan   = "\033[36m"
	gray   = "\033[37m"
	white  = "\033[97m"
	orange = "\033[38;5;166m"
)
