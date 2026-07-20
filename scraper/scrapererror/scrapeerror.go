package scrapererror

type ScrapeErrors struct {
	errs              []error
	failedScrapeCount int
}

func (s *ScrapeErrors) AddPartial(failed int, err error) { _ = "STUB: not implemented"; return }

func (s *ScrapeErrors) Add(err error) { _ = "STUB: not implemented"; return }

func (s *ScrapeErrors) Combine() error { _ = "STUB: not implemented"; return nil }
