package codegen

type UseCase struct {
	files     filesGetter
	presenter presenter
	tests     testRepository
}

func NewUseCase(fg filesGetter, p presenter, tr testRepository) *UseCase {
	return &UseCase{
		files:     fg,
		presenter: p,
		tests:     tr,
	}
}

func (u *UseCase) Execute(path string) error {
	files, err := u.files.Get(path)
	if err != nil {
		return err
	}

	tests := u.presenter.Present(files)

	return u.tests.Create(tests)
}
