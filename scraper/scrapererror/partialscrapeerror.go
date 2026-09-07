package scrapererror

type PartialScrapeError struct {
	error
	Failed int
}

func NewPartialScrapeError(err error, failed int) PartialScrapeError {
	_ = "STUB: not implemented"
	return *new(PartialScrapeError)
}

func IsPartialScrapeError(err error) bool { _ = "STUB: not implemented"; return false }
