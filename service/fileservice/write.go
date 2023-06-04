package fileservice

import "mediaStorer/param/storage"

func (s Service) WriteFile(file storage.File) error {
	err := s.repository.WriteFileToDatabase(file)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) DeleteFile(file storage.File) error {
	//check access
	return nil
}
