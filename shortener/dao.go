package shortener

var Shortener shortenerInterface = &shortener{}

type shortenerInterface interface {
	GenerateShortLink(initialLink string, userId string) string
}

type shortener struct{}
