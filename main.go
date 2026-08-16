package main
import "github.com/Nikos35/Gator/internal/config"

func main() {
	print(config.Read().DataBaseUrl)
}
