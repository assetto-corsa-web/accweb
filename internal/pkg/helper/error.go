package helper

type wrapError struct {
	err error
}

func (e *wrapError) Error() string {
	return e.err.Error()
}

func (e *wrapError) Unwrap() error {
	return e.err
}

func WrapErrors(errs ...error) error {
	if len(errs) == 0 {
		return nil
	}

	if len(errs) == 1 {
		return errs[0]
	}

	var e error
	for i := len(errs) - 1; i >= 0; i-- {
		if e == nil {
			e = errs[i]
		} else {
			e = &wrapError{err: e}
		}
	}

	return e
}
