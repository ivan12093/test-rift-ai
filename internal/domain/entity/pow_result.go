package entity

type POWResult struct {
	Valid bool
	Error error
}

func NewPOWResult(valid bool, err error) *POWResult {
	return &POWResult{
		Valid: valid,
		Error: err,
	}
}
