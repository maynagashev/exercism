package erratum

import "fmt"

// Use opens a resource, calls Frob(input) on the resource and then closes that resource (in all cases).
func Use(o ResourceOpener, input string) (err error) {
	resource, err := o()
	for err != nil {
		if _, ok := err.(TransientError); ok {
			resource, err = o()
			continue
		}
		return err
	}
	defer resource.Close()
	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case FrobError:
				resource.Defrob(e.defrobTag)
				err = e
			case error:
				err = e
			default:
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	resource.Frob(input)

	return nil
}
