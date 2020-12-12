package erratum

func Use(o ResourceOpener, input string) error {
	resource, err := o()
	if err != nil {
		return err
	}
	defer resource.Close()

	resource.Frob(input)

	return nil
}
