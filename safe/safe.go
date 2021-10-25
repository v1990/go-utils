package safe

// Executor executor
type Executor interface {
	Execute(f func())
}

// GoExecutor go executor
type GoExecutor interface {
	Go(f func())
}

// PanicHandler  handle panic Executor and GoExecutor
type PanicHandler func(recovered interface{})

// PanicRetryHandler handle panic and retry Executor and GoExecutor
type PanicRetryHandler func(recovered interface{}, attempt uint) (retry bool)

var (
	_ Executor = PanicHandler(nil)
	_ Executor = PanicRetryHandler(nil)

	_ GoExecutor = PanicHandler(nil)
	_ GoExecutor = PanicRetryHandler(nil)
)

func (t PanicHandler) Execute(f func()) {
	defer func() {
		if r := recover(); r != nil {
			t(r)
		}
	}()

	f()
}

func (t PanicHandler) Go(f func()) { go t.Execute(f) }

func (t PanicRetryHandler) Execute(f func()) {
	ff := func(attempt uint) (retry bool) {
		defer func() {
			if r := recover(); r != nil {
				retry = t(r, attempt)
			}
		}()

		f()
		return false
	}

	for attempt := uint(0); ff(attempt); attempt++ {
	}
}

func (t PanicRetryHandler) Go(f func()) { go t.Execute(f) }
